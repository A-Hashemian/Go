package main

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}
