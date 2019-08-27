package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
)

func parseBlockFile(f *os.File, blockchain *Blockchain, lastHeight *int, lastBlock *string) {
	for true {
		block := Block{}
		header := BlockHeader{}
		rawHeader := ""

		// Magic Number
		f.Read(read4)
		if bytes.Equal(magicBytes, read4) {
			// Size, no es parte del Block Header
			f.Read(read4)

			// Version
			f.Read(read4)
			header.Version = byte2int(reverse(read4))
			rawHeader += string(read4)

			// prevBlockHash
			f.Read(readHash)
			header.PrevBlockHash = hex.EncodeToString(reverse(readHash))
			rawHeader += string(readHash)

			// Merkle Root
			f.Read(readHash)
			header.MerkleRoot = hex.EncodeToString(reverse(readHash))
			rawHeader += string(readHash)

			// Timestamp
			f.Read(read4)
			header.Timestamp = byte2int(reverse(read4))
			rawHeader += string(read4)

			// Bits
			f.Read(read4)
			header.bits = hex.EncodeToString(reverse(read4))
			rawHeader += string(read4)

			// Nonce
			f.Read(read4)
			rawHeader += string(read4)

			header.Raw = hex.EncodeToString([]byte(rawHeader))

			block.Header = header
			block.Hash = doubleSHA256(rawHeader)

			txCount := readVariableInt(f)

			for tx := 0; tx < txCount; tx++ {
				tx := Transaction{}

				f.Read(read4)
				tx.Version = byte2int(reverse(read4))

				insCount := readVariableInt(f)

				for ins := 0; ins < insCount; ins++ {
					input := TxIn{}

					f.Read(readHash)
					f.Read(read4)

					input.TxID = hex.EncodeToString(reverse(readHash))
					input.N = byte2int(reverse(read4))

					scriptLength = readVariableInt(f)
					readScript = make([]byte, scriptLength)
					f.Read(readScript)

					input.SigScript = hex.EncodeToString(readScript)

					f.Read(read4)
					input.Sequence = byte2int(reverse(read4))

					tx.Inputs = append(tx.Inputs, input)
				}

				outsCount := readVariableInt(f)

				for out := 0; out < outsCount; out++ {
					output := TxOut{}

					f.Read(read8)
					output.Value = byte2int(reverse(read8))

					scriptLength = readVariableInt(f)
					readScript = make([]byte, scriptLength)
					f.Read(readScript)

					output.ScriptPubKey = hex.EncodeToString(readScript)

					tx.Outputs = append(tx.Outputs, output)
				}

				f.Read(read4)
				tx.Locktime = byte2int(reverse(read4))
				block.Transactions = append(block.Transactions, tx)
			}

			blockchain.Block = append(blockchain.Block, block)

			for i := len(blockchain.Block) - 1; i >= int(len(blockchain.Block)*8/10); i-- {
				block = blockchain.Block[i]
				if block.Header.PrevBlockHash == *lastBlock {
					*lastHeight++
					*lastBlock = block.Hash
					block.Height = *lastHeight
					fmt.Println(*lastHeight, *lastBlock)
					break
				}
			}

		} else {
			break
		}
	}
}
