package api

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Success code, usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error response
type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}
