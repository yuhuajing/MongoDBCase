package table

type Trans struct {
	BlockNumber      int64 `json:"blockNumber"`
	TransactionIndex int64 `json:"transactionIndex"`
	Gas              int64 `json:"gas"`
}

type NFTData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Result  Trans  `json:"result"`
}
