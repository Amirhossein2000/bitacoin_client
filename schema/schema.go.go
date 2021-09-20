package schema

type TransactionRequest struct {
	Time       int64
	FromPubKey []byte
	ToPubKey   []byte
	Signature  []byte
	Amount     int
}
