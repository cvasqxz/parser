package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	blockchain := Blockchain{}
	lastHeight := 0
	lastBlock := genesis

	blocksFolder := os.Getenv("HOME") + "/" + folder + "/blocks/"

	files, err := ioutil.ReadDir(blocksFolder)
	errorHandler(err)

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".dat") && strings.HasPrefix(name, "blk") {

			blockFile, err := os.Open(blocksFolder + name)
			errorHandler(err)

			parseBlockFile(blockFile, &blockchain, &lastHeight, &lastBlock)
		}
	}

}
