package main

import (
	"fmt"
	"hash/crc32"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		println("Must provide a TL-schema definition.")
		os.Exit(-1)
	}

	defer func() {
		if errByPanic := recover(); errByPanic != nil {
			println(errByPanic)
			os.Exit(-1)
		}
	}()

	re := regexp.MustCompile(`\s*\w+:flags\.\d+\?true`)
	tlSchemaInput := os.Args[1]
	for _, one := range re.FindAllString(tlSchemaInput, -1) {
		tlSchemaInput = strings.Replace(tlSchemaInput, one, " ", -1)
	}

	tlSchemaInput = strings.TrimSpace(tlSchemaInput)
	tlSchemaInput = strings.TrimSuffix(tlSchemaInput, ";")
	tlSchemaInput = strings.Replace(tlSchemaInput, "<", " ", -1)
	tlSchemaInput = strings.Replace(tlSchemaInput, ">", " ", -1)
	tlSchemaInput = strings.Replace(tlSchemaInput, "{", " ", -1)
	tlSchemaInput = strings.Replace(tlSchemaInput, "}", " ", -1)
	tlSchemaInput = strings.Replace(tlSchemaInput, ":bytes", ":string", -1)
	tlSchemaInput = strings.Replace(tlSchemaInput, "?bytes", "?string", -1)

	for {
		if strings.Contains(tlSchemaInput, "  ") {
			tlSchemaInput = strings.Replace(tlSchemaInput, "  ", " ", -1)
			continue
		}
		break
	}

	tlSchemaInputBySplit := strings.Split(tlSchemaInput, " ")
	if len(tlSchemaInputBySplit) != 0 {
		tlSchemaInputBySplit[0] = strings.Split(tlSchemaInputBySplit[0], "#")[0]
	}
	tlSchemaInputByjoin := strings.Join(tlSchemaInputBySplit, " ")
	fmt.Printf("\"%s\"  ->  0x%x\n", tlSchemaInputByjoin, crc32.ChecksumIEEE([]byte(tlSchemaInputByjoin)))
}
