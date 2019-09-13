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

	fmt.Println("Reading block indexes from", blocksFolder+"index")
	db, err := leveldb.OpenFile(blocksFolder+"/index", nil)
	errorHandler(err)
	defer db.Close()

	files, err := ioutil.ReadDir(blocksFolder)
	errorHandler(err)

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".dat") && strings.HasPrefix(name, "blk") {

			fmt.Println("Procesing file", blocksFolder+name)
			blockFile, err := os.Open(blocksFolder + name)
			defer blockFile.Close()
			errorHandler(err)

			parseBlockFile(blockFile, &blockchain, db)
		}
	}

	db.Close()

	fmt.Println("\nBlockchain info:")
	fmt.Println("\t* Blocks found:", len(blockchain.Block))
	fmt.Println("\t* Longest Chain:", blockchain.Tip)
	fmt.Println("\t* Orphan Blocks:", len(blockchain.Block)-blockchain.Tip)
	fmt.Println("\t* Best block Hash:", blockchain.bestBlockHash)
	fmt.Println("")

}
