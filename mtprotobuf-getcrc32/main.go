package main

import (
	"fmt"
	"github.com/qingsong-he/mtprotobuf/common"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		println("Must provide a TL-schema definition.")
		os.Exit(-1)
	}

	if strings.Count(os.Args[1], "\t") > 0 {
		panic(os.Args[1])
	}

	cleanLine, crc32ID := common.GetTLCRC32ByLine(os.Args[1])
	fmt.Printf("\"%s\"  ->  0x%x\n", cleanLine, crc32ID)
}
