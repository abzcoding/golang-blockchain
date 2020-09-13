package blockchain

// TxOutput represents a transaction output
type TxOutput struct {
	Value  int
	PubKey string
}

// TxInput represents a transaction input
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// CanUnlock determines the ability to unlock this transaction
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked will determine if we have the proper key to unlock this transaction or not
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
