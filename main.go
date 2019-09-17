package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/otiai10/copy"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	blockchain := Blockchain{}
	blockchain.Tip = 0
	blockchain.Blocks = make(map[int]Block)

	var err error

	err = os.RemoveAll(blocksFolder + "index_backup")
	errorHandler(err)
	err = copy.Copy(blocksFolder+"index", blocksFolder+"index_backup")
	errorHandler(err)

	fmt.Println("Reading block indexes from", blocksFolder+"index_backup")

	db, err := leveldb.OpenFile(blocksFolder+"/index_backup", nil)
	errorHandler(err)

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
	os.RemoveAll(blocksFolder + "index_backup")

	fmt.Println("\nBlockchain info:")
	fmt.Println("\t* Blocks found:", len(blockchain.Blocks)-1)
	fmt.Println("\t* Longest Chain:", blockchain.Tip)
	fmt.Println("\t* Orphan Blocks:", len(blockchain.Blocks)-blockchain.Tip)
	fmt.Println("\t* Best block Hash:", blockchain.bestBlockHash)
	fmt.Println("")

}
