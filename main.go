package main

import (
    "fmt"
    "os"
    "log"
    "bytes"
    "encoding/hex"
)

var magicBytes = []byte{170 ,162, 38, 169}
var readHash = make([]byte, 32)
var read8 = make([]byte, 8)
var read4 = make([]byte, 4)
var read1 = make([]byte, 1)
var read2 = make([]byte, 2)
var scriptLength = 0
var readScript = make([]byte, 0)

func byte2int(slice []byte) int {
  data := uint64(0)
  for _, b := range slice {
      data = (data << 8) | uint64(b)
  }
  return int(data)
}

func reverse(input []byte) []byte {
  if len(input) == 0 { return input }
  return append(reverse(input[1:]), input[0])
}


func readVariableInt(f *os.File) int {
  f.Read(read1)
  txCount := byte2int(read1)
  switch txCount {
    case 253:
      f.Read(read2)
      txCount = byte2int(reverse(read2))
    case 254:
      f.Read(read4)
      txCount = byte2int(reverse(read4))
    case 255:
      f.Read(read8)
      txCount = byte2int(reverse(read8))
  }
  return txCount
}

func main() {
  f, err := os.Open(os.Getenv("HOME") + "/.chaucha/blocks/blk00000.dat")
  if err != nil {
      log.Fatal(err)
  }

  for true {
    // Read Magic Number (4 bytes)
    f.Read(read4)
    if bytes.Equal(magicBytes, read4) {
      fmt.Println("########## START BLOCK ##########")
      f.Read(read4)
      fmt.Println("> Size:", hex.EncodeToString(reverse(read4)))
      f.Read(read4)
      fmt.Println("> Version:", hex.EncodeToString(reverse(read4)))
      f.Read(readHash)
      fmt.Println("> prevBlock:", hex.EncodeToString(reverse(readHash)))
      f.Read(readHash)
      fmt.Println("> MerkleRoot:", hex.EncodeToString(reverse(readHash)))
      f.Read(read4)
      fmt.Println("> Timestamp:", hex.EncodeToString(reverse(read4)))
      f.Read(read4)
      fmt.Println("> Bits:", hex.EncodeToString(reverse(read4)))
      f.Read(read4)
      fmt.Println("> Nonce:", hex.EncodeToString(reverse(read4)))

      txCount := readVariableInt(f)

      for tx := 0; tx < txCount; tx++ {
        fmt.Println("# TX:", tx)
        f.Read(read4)
        fmt.Println("> Tx Version:", hex.EncodeToString(reverse(read4)))

        insCount := readVariableInt(f)
        fmt.Println("\t# INPUTS:", insCount)

        for ins := 0; ins < insCount; ins++ {
          f.Read(readHash)
          fmt.Println("\t> TXID:", hex.EncodeToString(reverse(readHash)))
          f.Read(read4)
          fmt.Println("\t> N:", hex.EncodeToString(reverse(read4)))

          scriptLength = readVariableInt(f)
          readScript = make([]byte, scriptLength)
          f.Read(readScript)
          fmt.Println("\t> SigScript:", hex.EncodeToString(readScript))

          f.Read(read4)
          fmt.Println("\t> Sequence number:", hex.EncodeToString(reverse(read4)))
        }

        outsCount := readVariableInt(f)
        fmt.Println("\t# OUTPUTS:", outsCount)

        for out := 0; out < outsCount; out++ {
          f.Read(read8)
          Chauchas := byte2int(reverse(read8))
          fmt.Println("\t> Value:", Chauchas)

          scriptLength = readVariableInt(f)
          readScript = make([]byte, scriptLength)
          f.Read(readScript)
          fmt.Println("\t> ScriptPubKey:", hex.EncodeToString(readScript))
        }

        f.Read(read4)
        fmt.Println("\t> Locktime:", hex.EncodeToString(reverse(read4)))
      }

      fmt.Println("########## END BLOCK ##########\n")
    
    } else {
      os.Exit(0)
    }
  }
}
