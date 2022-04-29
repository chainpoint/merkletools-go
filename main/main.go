/* Copyright 2018 Tierion
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*     http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

 // This is a sample app which consumes and tests the merkletools package
 
 package main

import (
	"fmt"
	"merkletools"
	"encoding/hex"
	"time"
)

func main() {
	mt := merkletools.MerkleTree{}
	var hash, _ = hex.DecodeString("cb4990b9a8936bbc137ddeb6dcab4620897b099a450ecdc5f3e86ef4b3a7135c")
	leafCount := 100000
	for i := 0; i < leafCount; i++ {
		mt.AddLeaf(hash)
	}

	startAll := time.Now()

	startMake := time.Now()
	mt.MakeTree()
	endMake := time.Now()

	startProofs := time.Now()
	for i := 0; i < leafCount; i++ {
		proof := mt.GetProof(i)
		isValid := merkletools.VerifyProof(proof, mt.GetLeaf(i).Hash, mt.GetMerkleRoot())
		if !isValid {
			panic("Bad Proof!")
		}
	}
	endProofs := time.Now()

	endAll := time.Now()

	elapsedMake := endMake.Sub(startMake)
	elapsedProofs := endProofs.Sub(startProofs)
	elapsedAll := endAll.Sub(startAll)

	fmt.Printf("trial complete over %f seconds\nmake: %f\nproofs: %f", elapsedAll.Seconds(), elapsedMake.Seconds(), elapsedProofs.Seconds())
}
