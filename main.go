package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	blockchain := Blockchain{}
	blockchain.Tip = 0

	blocksFolder := os.Getenv("HOME") + "/" + folder + "/blocks/"

	db, err := leveldb.OpenFile(blocksFolder+"/index", nil)
	errorHandler(err)
	fmt.Println("Reading block indexes from", blocksFolder+"index")

	defer db.Close()

	files, err := ioutil.ReadDir(blocksFolder)
	errorHandler(err)

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".dat") && strings.HasPrefix(name, "blk") {

			blockFile, err := os.Open(blocksFolder + name)
			errorHandler(err)
			fmt.Println("Procesing file", blocksFolder+name)
			parseBlockFile(blockFile, &blockchain, db)
		}
	}

	fmt.Println("\nBlockchain info:")
	fmt.Println("\t* Blocks found:", len(blockchain.Block))
	fmt.Println("\t* Longest Chain:", blockchain.Tip)
	fmt.Println("\t* Orphan Blocks:", len(blockchain.Block)-blockchain.Tip)
	fmt.Println("\t* Best block Hash:", blockchain.bestBlockHash)
	fmt.Println("")
}
