# merkletools-go

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Merkletools-go is a library used by the Chainpoint project to efficiently generate and interact with [Merkle trees](https://en.wikipedia.org/wiki/Merkle_tree). 
It possesses a number of options for creating blockchain-compatible trees and generating proofs of inclusion.

## Install

This package uses Go modules for dependency management.

`go get github.com/chainpoint/merkletools-go`

## Usage

```go
package main

import (
	"fmt"
	"encoding/hex"
	"time"
	merkletools "github.com/chainpoint/merkletools-go"
)

func main() {
	mt := merkletools.MerkleTree{}
	var hash, _ = hex.DecodeString("cb4990b9a8936bbc137ddeb6dcab4620897b099a450ecdc5f3e86ef4b3a7135c")
	leafCount := 100000
	for i := 0; i < leafCount; i++ {
		mt.AddLeaf(hash)
	}

	mt.MakeTree()

	for i := 0; i < leafCount; i++ {
		proof := mt.GetProof(i)
		isValid := merkletools.VerifyProof(proof, mt.GetLeaf(i).Hash, mt.GetMerkleRoot())
		if !isValid {
			panic("Bad Proof!")
		}
	}
}
```