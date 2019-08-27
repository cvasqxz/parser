package main

type Block struct {
	Hash         string
	Header       BlockHeader
	Transactions []Transaction
}

type BlockHeader struct {
	Version       int
	PrevBlockHash string
	MerkleRoot    string
	Timestamp     int
	bits          string
	nonce         string
	Raw           string
}

type TxOut struct {
	Value        int
	ScriptPubKey string
}

type TxIn struct {
	TxID      string
	N         int
	SigScript string
	Sequence  int
}

type Transaction struct {
	Version  int
	Inputs   []TxIn
	Outputs  []TxOut
	Locktime int
}
