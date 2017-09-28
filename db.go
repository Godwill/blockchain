package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil{
		fmt.Println("Error occured ", err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block);
	if err != nil {
		fmt.Println("An error occurred ", err)
	}

	return &block
}
