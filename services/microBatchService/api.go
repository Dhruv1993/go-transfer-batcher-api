package microBatchService

// API consists of different apis this service has
type API interface {
	processTransfers(transfer Transfer) (Transfer, error)
}
