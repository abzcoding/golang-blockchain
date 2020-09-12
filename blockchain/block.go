package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block represents a data block in the blockchain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock creates a new block based on the data
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis creates the initial block inside the blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize implements the serializable for block object
func (b *Block) Serialize() []byte {
	var res bytes.Buffer

	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}

// Deserialize implements the deserializer for the block object
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	Handle(err)

	return &block
}

// Handle will do the error handling
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
