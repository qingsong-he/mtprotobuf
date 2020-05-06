package common

import (
	"hash/crc32"
	"regexp"
	"strings"
)

func GetTLCRC32ByLine(tlSchemaInput string) (cleanLine string, crc32ID uint32) {
	re := regexp.MustCompile(`\s*\w+:flags\.\d+\?true`)
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
	return tlSchemaInputByjoin, crc32.ChecksumIEEE([]byte(tlSchemaInputByjoin))
}
