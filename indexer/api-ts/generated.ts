// Code generated by tygo. DO NOT EDIT.

//////////
// source: models.go

/**
 * DepositItem ... Deposit item model for API responses
 */
export interface DepositItem {
  guid: string;
  from: string;
  to: string;
  timestamp: number /* uint64 */;
  l1BlockHash: string;
  l1TxHash: string;
  l2TxHash: string;
  amount: string;
  l1TokenAddress: string;
  l2TokenAddress: string;
}
/**
 * DepositResponse ... Data model for API JSON response
 */
export interface DepositResponse {
  cursor: string;
  hasNextPage: boolean;
  items: DepositItem[];
}
/**
 * WithdrawalItem ... Data model for API JSON response
 */
export interface WithdrawalItem {
  guid: string;
  from: string;
  to: string;
  transactionHash: string;
  crossDomainMessageHash: string;
  timestamp: number /* uint64 */;
  l2BlockHash: string;
  amount: string;
  l1ProvenTxHash: string;
  l1FinalizedTxHash: string;
  l1TokenAddress: string;
  l2TokenAddress: string;
}
/**
 * WithdrawalResponse ... Data model for API JSON response
 */
export interface WithdrawalResponse {
  cursor: string;
  hasNextPage: boolean;
  items: WithdrawalItem[];
}
