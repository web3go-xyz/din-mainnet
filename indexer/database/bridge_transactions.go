package database

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
)

/**
 * Types
 */

type Transaction struct {
	FromAddress common.Address `gorm:"serializer:bytes"`
	ToAddress   common.Address `gorm:"serializer:bytes"`
	Amount      *big.Int       `gorm:"serializer:u256"`
	Data        Bytes          `gorm:"serializer:bytes"`
	Timestamp   uint64
}

type L1TransactionDeposit struct {
	SourceHash           common.Hash `gorm:"serializer:bytes;primaryKey"`
	L2TransactionHash    common.Hash `gorm:"serializer:bytes"`
	InitiatedL1EventGUID uuid.UUID

	Tx       Transaction `gorm:"embedded"`
	GasLimit *big.Int    `gorm:"serializer:u256"`
}

type L2TransactionWithdrawal struct {
	WithdrawalHash       common.Hash `gorm:"serializer:bytes;primaryKey"`
	Nonce                *big.Int    `gorm:"serializer:u256"`
	InitiatedL2EventGUID uuid.UUID

	ProvenL1EventGUID    *uuid.UUID
	FinalizedL1EventGUID *uuid.UUID
	Succeeded            *bool

	Tx       Transaction `gorm:"embedded"`
	GasLimit *big.Int    `gorm:"serializer:u256"`
}

type BridgeTransactionsView interface {
	L1TransactionDeposit(common.Hash) (*L1TransactionDeposit, error)
	L1LatestBlockHeader() (*L1BlockHeader, error)

	L2TransactionWithdrawal(common.Hash) (*L2TransactionWithdrawal, error)
	L2LatestBlockHeader() (*L2BlockHeader, error)
}

type BridgeTransactionsDB interface {
	BridgeTransactionsView

	StoreL1TransactionDeposits([]L1TransactionDeposit) error

	StoreL2TransactionWithdrawals([]L2TransactionWithdrawal) error
	MarkL2TransactionWithdrawalProvenEvent(common.Hash, uuid.UUID) error
	MarkL2TransactionWithdrawalFinalizedEvent(common.Hash, uuid.UUID, bool) error
}

/**
 * Implementation
 */

type bridgeTransactionsDB struct {
	gorm *gorm.DB
}

func newBridgeTransactionsDB(db *gorm.DB) BridgeTransactionsDB {
	return &bridgeTransactionsDB{gorm: db}
}

/**
 * Transactions deposited from L1
 */

func (db *bridgeTransactionsDB) StoreL1TransactionDeposits(deposits []L1TransactionDeposit) error {
	result := db.gorm.Create(&deposits)
	return result.Error
}

func (db *bridgeTransactionsDB) L1TransactionDeposit(sourceHash common.Hash) (*L1TransactionDeposit, error) {
	var deposit L1TransactionDeposit
	result := db.gorm.Where(&L1TransactionDeposit{SourceHash: sourceHash}).Take(&deposit)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &deposit, nil
}

func (db *bridgeTransactionsDB) L1LatestBlockHeader() (*L1BlockHeader, error) {
	// Markers for an indexed bridge event
	// L1: Latest Transaction Deposit, Latest Proven/Finalized Withdrawal
	l1DepositQuery := db.gorm.Table("l1_transaction_deposits").Order("l1_transaction_deposits.timestamp DESC").Limit(1)
	l1DepositQuery = l1DepositQuery.Joins("INNER JOIN l1_contract_events ON l1_contract_events.guid = l1_transaction_deposits.initiated_l1_event_guid")
	l1DepositQuery = l1DepositQuery.Select("l1_contract_events.*")

	l1ProvenQuery := db.gorm.Table("l2_transaction_withdrawals")
	l1ProvenQuery = l1ProvenQuery.Joins("INNER JOIN l1_contract_events ON l1_contract_events.guid = l2_transaction_withdrawals.proven_l1_event_guid")
	l1ProvenQuery = l1ProvenQuery.Order("l1_contract_events.timestamp DESC").Select("l1_contract_events.*").Limit(1)

	l1FinalizedQuery := db.gorm.Table("l2_transaction_withdrawals")
	l1FinalizedQuery = l1FinalizedQuery.Joins("INNER JOIN l1_contract_events ON l1_contract_events.guid = l2_transaction_withdrawals.proven_l1_event_guid")
	l1FinalizedQuery = l1FinalizedQuery.Order("l1_contract_events.timestamp DESC").Select("l1_contract_events.*").Limit(1)

	l1Query := db.gorm.Table("((?) UNION (?) UNION (?)) AS latest_bridge_events", l1DepositQuery.Limit(1), l1ProvenQuery, l1FinalizedQuery)
	l1Query = l1Query.Joins("INNER JOIN l1_block_headers ON l1_block_headers.hash = latest_bridge_events.block_hash")
	l1Query = l1Query.Order("l1_block_headers.number DESC").Select("l1_block_headers.*")

	var l1Header L1BlockHeader
	result := l1Query.Take(&l1Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &l1Header, nil
}

/**
 * Transactions withdrawn from L2
 */

func (db *bridgeTransactionsDB) StoreL2TransactionWithdrawals(withdrawals []L2TransactionWithdrawal) error {
	result := db.gorm.Create(&withdrawals)
	return result.Error
}

func (db *bridgeTransactionsDB) L2TransactionWithdrawal(withdrawalHash common.Hash) (*L2TransactionWithdrawal, error) {
	var withdrawal L2TransactionWithdrawal
	result := db.gorm.Where(&L2TransactionWithdrawal{WithdrawalHash: withdrawalHash}).Take(&withdrawal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &withdrawal, nil
}

// MarkL2TransactionWithdrawalProvenEvent links a withdrawn transaction with associated Prove action on L1.
func (db *bridgeTransactionsDB) MarkL2TransactionWithdrawalProvenEvent(withdrawalHash common.Hash, provenL1EventGuid uuid.UUID) error {
	withdrawal, err := db.L2TransactionWithdrawal(withdrawalHash)
	if err != nil {
		return err
	}
	if withdrawal == nil {
		return fmt.Errorf("transaction withdrawal hash %s not found", withdrawalHash)
	}

	withdrawal.ProvenL1EventGUID = &provenL1EventGuid
	result := db.gorm.Save(&withdrawal)
	return result.Error
}

// MarkL2TransactionWithdrawalProvenEvent links a withdrawn transaction in its finalized state
func (db *bridgeTransactionsDB) MarkL2TransactionWithdrawalFinalizedEvent(withdrawalHash common.Hash, finalizedL1EventGuid uuid.UUID, succeeded bool) error {
	withdrawal, err := db.L2TransactionWithdrawal(withdrawalHash)
	if err != nil {
		return err
	}
	if withdrawal == nil {
		return fmt.Errorf("transaction withdrawal hash %s not found", withdrawalHash)
	}
	if withdrawal.ProvenL1EventGUID == nil {
		return fmt.Errorf("cannot mark unproven withdrawal hash %s as finalized", withdrawal.WithdrawalHash)
	}

	withdrawal.FinalizedL1EventGUID = &finalizedL1EventGuid
	withdrawal.Succeeded = &succeeded
	result := db.gorm.Save(&withdrawal)
	return result.Error
}

func (db *bridgeTransactionsDB) L2LatestBlockHeader() (*L2BlockHeader, error) {
	// L2: Inclusion of the latest deposit
	l1DepositQuery := db.gorm.Table("l1_transaction_deposits").Order("l1_transaction_deposits.timestamp DESC")
	l1DepositQuery = l1DepositQuery.Joins("INNER JOIN l1_contract_events ON l1_contract_events.guid = l1_transaction_deposits.initiated_l1_event_guid")
	l1DepositQuery = l1DepositQuery.Select("l1_contract_events.*")

	l2Query := db.gorm.Table("(?) AS l1_deposit_events", l1DepositQuery)
	l2Query = l2Query.Joins("INNER JOIN l2_block_headers ON l2_block_headers.timestamp = l1_deposit_events.timestamp")
	l2Query = l2Query.Order("l2_block_headers.timestamp DESC").Select("l2_block_headers.*")

	var l2Header L2BlockHeader
	result := l2Query.Take(&l2Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &l2Header, nil
}
