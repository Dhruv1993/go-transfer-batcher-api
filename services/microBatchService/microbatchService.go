package microBatchService

// TransferJob represents the details of a money transfer.
type TransferRequest struct {
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Amount      float64 `json:"amount"`
}

type TransferResponse struct {
	Message  string
	Transfer Transfer
	status   string
}

type Transfer struct {
	FromAccount string
	ToAccount   string
	Amount      float64
}
