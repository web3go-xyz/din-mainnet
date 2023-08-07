
CREATE DOMAIN UINT256 AS NUMERIC
    CHECK (VALUE >= 0 AND VALUE < 2^256 and SCALE(VALUE) = 0);

/**
 * BLOCK DATA
 */

CREATE TABLE IF NOT EXISTS l1_block_headers (
	hash        VARCHAR NOT NULL PRIMARY KEY,
	parent_hash VARCHAR NOT NULL,
	number      UINT256 NOT NULL,
	timestamp   INTEGER NOT NULL CHECK (timestamp > 0)
);

CREATE TABLE IF NOT EXISTS l2_block_headers (
    -- Block header
	hash                     VARCHAR NOT NULL PRIMARY KEY,
	parent_hash              VARCHAR NOT NULL,
	number                   UINT256 NOT NULL,
	timestamp                INTEGER NOT NULL CHECK (timestamp > 0)
);

/** 
 * EVENT DATA
 */

CREATE TABLE IF NOT EXISTS l1_contract_events (
    guid             VARCHAR NOT NULL PRIMARY KEY,
	block_hash       VARCHAR NOT NULL REFERENCES l1_block_headers(hash),
    transaction_hash VARCHAR NOT NULL,
    event_signature  VARCHAR NOT NULL,
    log_index        INTEGER NOT NULL,
    timestamp        INTEGER NOT NULL CHECK (timestamp > 0)
);

CREATE TABLE IF NOT EXISTS l2_contract_events (
    guid             VARCHAR NOT NULL PRIMARY KEY,
	block_hash       VARCHAR NOT NULL REFERENCES l2_block_headers(hash),
    transaction_hash VARCHAR NOT NULL,
    event_signature  VARCHAR NOT NULL,
    log_index        INTEGER NOT NULL,
    timestamp        INTEGER NOT NULL CHECK (timestamp > 0)
);

-- Tables that index finalization markers for L2 blocks.

CREATE TABLE IF NOT EXISTS legacy_state_batches (
	index         INTEGER NOT NULL PRIMARY KEY,
	root          VARCHAR NOT NULL,
	size          INTEGER NOT NULL,
	prev_total    INTEGER NOT NULL,

    l1_contract_event_guid VARCHAR REFERENCES l1_contract_events(guid)
);

CREATE TABLE IF NOT EXISTS output_proposals (
    output_root     VARCHAR NOT NULL PRIMARY KEY,

    l2_output_index UINT256 NOT NULL,
    l2_block_number UINT256 NOT NULL,

    l1_contract_event_guid VARCHAR REFERENCES l1_contract_events(guid)
);

/**
 * BRIDGING DATA
 */

-- OptimismPortal/L2ToL1MessagePasser
CREATE TABLE IF NOT EXISTS l1_transaction_deposits (
    source_hash         VARCHAR NOT NULL PRIMARY KEY,
    l2_transaction_hash VARCHAR NOT NULL,

    initiated_l1_event_guid VARCHAR NOT NULL REFERENCES l1_contract_events(guid),

    -- OptimismPortal specific
    version     UINT256 NOT NULL,
    opaque_data VARCHAR NOT NULL,

    -- transaction data
    from_address VARCHAR NOT NULL,
    to_address   VARCHAR NOT NULL,
    amount       UINT256 NOT NULL,
    gas_limit    UINT256 NOT NULL,
    data         VARCHAR NOT NULL,
    timestamp    INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE TABLE IF NOT EXISTS l2_transaction_withdrawals (
    withdrawal_hash VARCHAR NOT NULL PRIMARY KEY,

    initiated_l2_event_guid VARCHAR NOT NULL REFERENCES l2_contract_events(guid),

    -- Multistep (bedrock) process of a withdrawal
    proven_l1_event_guid    VARCHAR REFERENCES l1_contract_events(guid),
    finalized_l1_event_guid VARCHAR REFERENCES l1_contract_events(guid),
    succeeded               BOOLEAN,

    -- L2ToL1MessagePasser specific
    nonce UINT256 UNIQUE,

    -- transaction data
    from_address VARCHAR NOT NULL,
    to_address   VARCHAR NOT NULL,
    amount       UINT256 NOT NULL,
    gas_limit    UINT256 NOT NULL,
    data         VARCHAR NOT NULL,
    timestamp    INTEGER NOT NULL CHECK (timestamp > 0)
);

-- StandardBridge
CREATE TABLE IF NOT EXISTS l1_bridge_deposits (
    transaction_source_hash VARCHAR PRIMARY KEY REFERENCES l1_transaction_deposits(source_hash),

    -- We allow the cross_domain_messenger_nonce to be NULL-able to account
    -- for scenarios where ETH is simply sent to the OptimismPortal contract
    cross_domain_messenger_nonce UINT256 UNIQUE,

    -- Deposit information
	from_address     VARCHAR NOT NULL,
	to_address       VARCHAR NOT NULL,
	l1_token_address VARCHAR NOT NULL,
	l2_token_address VARCHAR NOT NULL,
	amount           UINT256 NOT NULL,
	data             VARCHAR NOT NULL,
    timestamp        INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE TABLE IF NOT EXISTS l2_bridge_withdrawals (
    transaction_withdrawal_hash VARCHAR PRIMARY KEY REFERENCES l2_transaction_withdrawals(withdrawal_hash),

    -- We allow the cross_domain_messenger_nonce to be NULL-able to account for
    -- scenarios where ETH is simply sent to the L2ToL1MessagePasser contract
    cross_domain_messenger_nonce UINT256 UNIQUE,

    -- Withdrawal information
	from_address     VARCHAR NOT NULL,
	to_address       VARCHAR NOT NULL,
	l1_token_address VARCHAR NOT NULL,
	l2_token_address VARCHAR NOT NULL,
	amount           UINT256 NOT NULL,
	data             VARCHAR NOT NULL,
    timestamp        INTEGER NOT NULL CHECK (timestamp > 0)
);
