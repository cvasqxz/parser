package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	blocksFolder := os.Getenv("HOME") + "/" + folder + "/blocks/"

	files, err := ioutil.ReadDir(blocksFolder)
	errorHandler(err)

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".dat") && strings.HasPrefix(name, "blk") {

			blockFile, err := os.Open(blocksFolder + name)
			errorHandler(err)
			fmt.Println("Processing", blocksFolder+name)
			parseBlockFile(blockFile)
		}
	}

}
