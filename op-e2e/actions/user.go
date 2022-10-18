package actions

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
)

type L1Bindings struct {
	// contract bindings
	OptimismPortal *bindings.OptimismPortal

	L2OutputOracle *bindings.L2OutputOracle
}

func NewL1Bindings(t Testing, l1Cl *ethclient.Client, deployments *e2eutils.DeploymentsL1) *L1Bindings {
	optimismPortal, err := bindings.NewOptimismPortal(deployments.OptimismPortalProxy, l1Cl)
	require.NoError(t, err)

	l2OutputOracle, err := bindings.NewL2OutputOracle(deployments.L2OutputOracleProxy, l1Cl)
	require.NoError(t, err)

	return &L1Bindings{
		OptimismPortal: optimismPortal,
		L2OutputOracle: l2OutputOracle,
	}
}

type L2Bindings struct {
	L2ToL1MessagePasser *bindings.L2ToL1MessagePasser

	WithdrawalsClient *withdrawals.Client
}

func NewL2Bindings(t Testing, l2Cl *ethclient.Client, withdrawalsCl *withdrawals.Client) *L2Bindings {
	l2ToL1MessagePasser, err := bindings.NewL2ToL1MessagePasser(predeploys.L2ToL1MessagePasserAddr, l2Cl)
	require.NoError(t, err)

	return &L2Bindings{
		L2ToL1MessagePasser: l2ToL1MessagePasser,
		WithdrawalsClient:   withdrawalsCl,
	}
}

// BasicUserEnv provides access to the eth RPC, signer, and contract bindings for a single ethereum layer.
// This environment can be shared between different BasicUser instances.
type BasicUserEnv[B any] struct {
	EthCl  *ethclient.Client
	Signer types.Signer

	AddressCorpora []common.Address

	Bindings B
}

// BasicUser is an actor on a single ethereum layer, with one account key.
// The user maintains a set of standard txOpts to build its transactions with,
// along with configurable txToAddr and txCallData.
// The user has an RNG source with actions to randomize its transaction building.
type BasicUser[B any] struct {
	log log.Logger
	rng *rand.Rand
	env *BasicUserEnv[B]

	account *ecdsa.PrivateKey
	address common.Address

	txOpts bind.TransactOpts

	txToAddr   *common.Address
	txCallData []byte

	// lastTxHash persists the last transaction,
	// so we can chain together tx sending and tx checking easily.
	// Sending and checking are detached, since txs may not be instantly confirmed.
	lastTxHash common.Hash
}

func NewBasicUser[B any](log log.Logger, priv *ecdsa.PrivateKey, rng *rand.Rand) *BasicUser[B] {
	return &BasicUser[B]{
		log:     log,
		rng:     rng,
		account: priv,
		address: crypto.PubkeyToAddress(priv.PublicKey),
	}
}

// SetUserEnv changes the user environment.
// This way a user can be initialized before being embedded in a genesis allocation,
// and change between different endpoints that may be initialized after the user.
func (s *BasicUser[B]) SetUserEnv(env *BasicUserEnv[B]) {
	s.env = env
}

func (s *BasicUser[B]) signerFn(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
	if address != s.address {
		return nil, bind.ErrNotAuthorized
	}
	signature, err := crypto.Sign(s.env.Signer.Hash(tx).Bytes(), s.account)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(s.env.Signer, signature)
}

// ActResetTxOpts prepares the tx options to default values, based on the current pending block header.
func (s *BasicUser[B]) ActResetTxOpts(t Testing) {
	pendingHeader, err := s.env.EthCl.HeaderByNumber(t.Ctx(), big.NewInt(-1))
	require.NoError(t, err, "need l2 pending header for accurate basefee info")

	gasTipCap := big.NewInt(2 * params.GWei)
	gasFeeCap := new(big.Int).Add(gasTipCap, new(big.Int).Mul(pendingHeader.BaseFee, big.NewInt(2)))

	s.txOpts = bind.TransactOpts{
		From:      s.address,
		Nonce:     nil, // pick nonce based on pending state
		Signer:    s.signerFn,
		Value:     big.NewInt(0),
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		GasLimit:  0,    // a.k.a. estimate
		NoSend:    true, // actions should be explicit about sending
	}
}

func (s *BasicUser[B]) ActRandomTxToAddr(t Testing) {
	i := s.rng.Intn(len(s.env.AddressCorpora))
	var to *common.Address
	if i > 0 { // 0 == nil
		to = &s.env.AddressCorpora[i]
	}
	s.txToAddr = to
}

func (s *BasicUser[B]) ActSetTxCalldata(calldata []byte) Action {
	return func(t Testing) {
		require.NotNil(t, calldata)
		s.txCallData = calldata
	}
}

func (s *BasicUser[B]) ActSetTxToAddr(to *common.Address) Action {
	return func(t Testing) {
		s.txToAddr = to
	}
}

func (s *BasicUser[B]) ActRandomTxValue(t Testing) {
	// compute a random portion of balance
	precision := int64(1000)
	bal, err := s.env.EthCl.BalanceAt(t.Ctx(), s.address, nil)
	require.NoError(t, err)
	part := big.NewInt(s.rng.Int63n(precision))
	new(big.Int).Div(new(big.Int).Mul(bal, part), big.NewInt(precision))
	s.txOpts.Value = big.NewInt(s.rng.Int63())
}

func (s *BasicUser[B]) ActSetTxValue(value *big.Int) Action {
	return func(t Testing) {
		s.txOpts.Value = value
	}
}

func (s *BasicUser[B]) ActRandomTxData(t Testing) {
	dataLen := s.rng.Intn(128_000)
	out := make([]byte, dataLen)
	_, err := s.rng.Read(out[:])
	require.NoError(t, err)
	s.txCallData = out
}

func (s *BasicUser[B]) PendingNonce(t Testing) uint64 {
	if s.txOpts.Nonce != nil {
		return s.txOpts.Nonce.Uint64()
	}
	// fetch from pending state
	nonce, err := s.env.EthCl.PendingNonceAt(t.Ctx(), s.address)
	require.NoError(t, err, "failed to get L1 nonce for account %s", s.address)
	return nonce
}

func (s *BasicUser[B]) TxValue() *big.Int {
	if s.txOpts.Value != nil {
		return s.txOpts.Value
	}
	return big.NewInt(0)
}

func (s *BasicUser[B]) LastTxReceipt(t Testing) *types.Receipt {
	require.NotEqual(t, s.lastTxHash, common.Hash{}, "must send tx before getting last receipt")
	receipt, err := s.env.EthCl.TransactionReceipt(t.Ctx(), s.lastTxHash)
	require.NoError(t, err)
	return receipt
}

// ActMakeTx makes a tx with the predetermined contents (see randomization and other actions)
// and sends it to the tx pool
func (s *BasicUser[B]) ActMakeTx(t Testing) {
	gas, err := s.env.EthCl.EstimateGas(t.Ctx(), ethereum.CallMsg{
		From:      s.address,
		To:        s.txToAddr,
		GasFeeCap: s.txOpts.GasFeeCap,
		GasTipCap: s.txOpts.GasTipCap,
		Value:     s.TxValue(),
		Data:      s.txCallData,
	})
	require.NoError(t, err, "gas estimation should pass")
	tx := types.MustSignNewTx(s.account, s.env.Signer, &types.DynamicFeeTx{
		To:        s.txToAddr,
		GasFeeCap: s.txOpts.GasFeeCap,
		GasTipCap: s.txOpts.GasTipCap,
		Value:     s.TxValue(),
		ChainID:   s.env.Signer.ChainID(),
		Nonce:     s.PendingNonce(t),
		Gas:       gas,
		Data:      s.txCallData,
	})
	err = s.env.EthCl.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err, "must send tx")
	s.lastTxHash = tx.Hash()
	// reset the calldata
	s.txCallData = []byte{}
}

func (s *BasicUser[B]) ActCheckReceiptStatusOfLastTx(success bool) func(t Testing) {
	return func(t Testing) {
		s.CheckReceipt(t, success, s.lastTxHash)
	}
}

func (s *BasicUser[B]) CheckReceipt(t Testing, success bool, txHash common.Hash) *types.Receipt {
	receipt, err := s.env.EthCl.TransactionReceipt(t.Ctx(), txHash)
	if receipt != nil && err == nil {
		expected := types.ReceiptStatusFailed
		if success {
			expected = types.ReceiptStatusSuccessful
		}
		require.Equal(t, expected, receipt.Status, "expected receipt status to match")
		return receipt
	} else if err != nil && !errors.Is(err, ethereum.NotFound) {
		t.Fatalf("receipt for tx %s was not found", txHash)
	} else {
		t.Fatalf("receipt error: %v", err)
	}
	return nil
}

type L1User struct {
	BasicUser[*L1Bindings]
}

type L2User struct {
	BasicUser[*L2Bindings]
}

// CrossLayerUser represents the same user account on L1 and L2,
// and provides actions to make cross-layer transactions.
type CrossLayerUser struct {
	L1 L1User
	L2 L2User

	// track the last deposit, to easily chain together deposit actions
	lastL1DepositTxHash common.Hash
}

func NewCrossLayerUser(log log.Logger, priv *ecdsa.PrivateKey, rng *rand.Rand) *CrossLayerUser {
	addr := crypto.PubkeyToAddress(priv.PublicKey)
	return &CrossLayerUser{
		L1: L1User{
			BasicUser: BasicUser[*L1Bindings]{
				log:     log,
				rng:     rng,
				account: priv,
				address: addr,
			},
		},
		L2: L2User{
			BasicUser: BasicUser[*L2Bindings]{
				log:     log,
				rng:     rng,
				account: priv,
				address: addr,
			},
		},
	}
}

func (s *CrossLayerUser) ActDeposit(t Testing) {
	isCreation := false
	toAddr := common.Address{}
	if s.L2.txToAddr == nil {
		isCreation = true
	} else {
		toAddr = *s.L2.txToAddr
	}
	depositTransferValue := s.L2.TxValue()
	depositGas := s.L2.txOpts.GasLimit
	if s.L2.txOpts.GasLimit == 0 {
		// estimate gas used by deposit
		gas, err := s.L2.env.EthCl.EstimateGas(t.Ctx(), ethereum.CallMsg{
			From:       s.L2.address,
			To:         s.L2.txToAddr,
			Value:      depositTransferValue, // TODO: estimate gas does not support minting yet
			Data:       s.L2.txCallData,
			AccessList: nil,
		})
		require.NoError(t, err)
		depositGas = gas
	}

	tx, err := s.L1.env.Bindings.OptimismPortal.DepositTransaction(&s.L1.txOpts, toAddr, depositTransferValue, depositGas, isCreation, s.L2.txCallData)
	require.NoError(t, err, "failed to create deposit tx")

	// Send the actual tx (since tx opts don't send by default)
	err = s.L1.env.EthCl.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err, "must send tx")
	s.lastL1DepositTxHash = tx.Hash()
}

func (s *CrossLayerUser) ActCheckDepositStatus(l1Success, l2Success bool) Action {
	return func(t Testing) {
		s.CheckDepositTx(t, s.lastL1DepositTxHash, 0, l1Success, l2Success)
	}
}

func (s *CrossLayerUser) CheckDepositTx(t Testing, l1TxHash common.Hash, index int, l1Success, l2Success bool) {
	depositReceipt := s.L1.CheckReceipt(t, l1Success, l1TxHash)
	if depositReceipt == nil {
		require.False(t, l1Success)
		require.False(t, l2Success)
	} else {
		require.Less(t, index, len(depositReceipt.Logs), "must have enough logs in receipt")
		reconstructedDep, err := derive.UnmarshalDepositLogEvent(depositReceipt.Logs[index])
		require.NoError(t, err, "Could not reconstruct L2 Deposit")
		l2Tx := types.NewTx(reconstructedDep)
		s.L2.CheckReceipt(t, l2Success, l2Tx.Hash())
	}
}

func (s *CrossLayerUser) Address() common.Address {
	return s.L1.address
}
