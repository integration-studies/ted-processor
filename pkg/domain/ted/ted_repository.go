package ted

type Repository interface {

	StorePreTransaction(t *PreTransaction)(*PreTransaction,error)

	StoreTransactionConfirmation(t *TransactionConfirmation)(*TransactionConfirmation,error)

}