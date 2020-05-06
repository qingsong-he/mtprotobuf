package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/qingsong-he/mtprotobuf/common"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s foobar.tl\n\n", os.Args[0])
		os.Exit(-1)
	}

	tlBin, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	var layerStr string = "layer0"
	var layerNumberStr string = "0"
	var crc32Ids []string
	var inits []string
	var objs []string

	sc := bufio.NewScanner(bytes.NewReader(tlBin))
	for sc.Scan() {
		lineByTrim := strings.TrimSpace(sc.Text())
		if strings.HasPrefix(lineByTrim, "//") {
			sp := strings.Split(lineByTrim, " ")
			if len(sp) == 3 && strings.ToLower(sp[1]) == "layer" {
				layerNumberStr = sp[2]
				layerStr = "layer" + sp[2]
			}
			continue
		}
		if lineByTrim == "" {
			continue
		}

		var lineByTrimSplit []string
		for _, v := range strings.Split(lineByTrim, " ") {
			if v1 := strings.TrimSpace(v); v1 != "" {
				lineByTrimSplit = append(lineByTrimSplit, v1)
			}
		}
		lineByTrim = strings.Join(lineByTrimSplit, " ")
		if strings.Count(lineByTrim, "\t") > 0 {
			panic(lineByTrim)
		}

		cleanLine, crc32ID := common.GetTLCRC32ByLine(lineByTrim)
		if strings.Count(cleanLine, "#") > 1 {
			panic(lineByTrim)
		}
		if strings.Count(lineByTrim, "=") != 1 {
			panic(lineByTrim)
		}

		// debug output
		fmt.Println("'" + lineByTrim + "'")
		fmt.Println("\t'" + cleanLine + "'")
		fmt.Println("\t\t'" + fmt.Sprintf("0x%x", crc32ID) + "'")
		fmt.Println()

		cleanLineBySpli := strings.Split(cleanLine, " ")
		if len(cleanLineBySpli) < 3 {
			panic(lineByTrim)
		}

		// get crc id
		typeName := "TL_" + strings.Replace(cleanLineBySpli[0], ".", "_", -1) + "_" + layerStr
		crcIdName := "CRC32_" + typeName
		crc32Ids = append(crc32Ids, crcIdName+" uint32="+fmt.Sprintf("0x%x", crc32ID))

		// get init
		inits = append(inits, fmt.Sprintf(fmtByInit, crcIdName, crcIdName, typeName))

		// get type
		{
			typeBuf1 := bytes.NewBuffer(nil)
			typeBuf1.WriteString(fmt.Sprintf("// begin of '%s'\n", lineByTrim))

			typeBuf1.WriteString(fmt.Sprintf("type %s struct {\n", typeName))
			var s scanner.Scanner
			for _, v := range lineByTrimSplit[1:] {
				if v == "=" {
					break
				}

				s.Init(strings.NewReader(v))
				var parts []string
				for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
					parts = append(parts, s.TokenText())
				}
				switch len(parts) {
				case 3:
					switch parts[2] {
					case "#":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " int32\n")
					case "int":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " int32\n")
					case "long":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " int64\n")
					case "double":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " float64\n")
					case "string":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " string\n")
					default:
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " mtprotobuf.TL\n")
					}
				case 6:
					switch parts[5] {
					case ">":
						switch parts[4] {
						case "int":
							typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []int32\n")
						case "long":
							typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []int64\n")
						case "double":
							typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []float64\n")
						case "string":
							typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []string\n")
						default:
							typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []mtprotobuf.TL\n")
						}
					case "int":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " int32\n")
					case "long":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " int64\n")
					case "double":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " float64\n")
					case "string":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " string\n")
					case "true":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " bool\n")
					default:
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " mtprotobuf.TL\n")
					}
				case 9:
					switch parts[7] {
					case "int":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []int32\n")
					case "long":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []int64\n")
					case "double":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []float64\n")
					case "string":
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []string\n")
					default:
						typeBuf1.WriteString("\t" + strings.Title(parts[0]) + " []mtprotobuf.TL\n")
					}
				default:
					panic(lineByTrim)
				}
			}
			typeBuf1.WriteString("}\n")

			typeBuf1.WriteString(fmt.Sprintf(fmtByTypeFunc, typeName, typeName, typeName, typeName, crcIdName, typeName, layerNumberStr))

			// get encode and decode func
			{
				var boolsByEncode, boolsByDecode []string
				var notBoolsByEncode, notBoolsByDecode []string
				for _, v := range lineByTrimSplit[1:] {
					if v == "=" {
						break
					}

					s.Init(strings.NewReader(v))
					var parts []string
					for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
						parts = append(parts, s.TokenText())
					}

					switch len(parts) {
					case 3:
						// default_banned_rights:#
						switch parts[2] {
						case "#":
							boolsByDecode = append(boolsByDecode, fmt.Sprintf("t.%s = d.Int()", strings.Title(parts[0])))
						case "int":
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.Int(t.%s)", strings.Join(parts, ""), strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.Int()", strings.Join(parts, ""), strings.Title(parts[0])))
						case "long":
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.Long(t.%s)", strings.Join(parts, ""), strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.Long()", strings.Join(parts, ""), strings.Title(parts[0])))
						case "double":
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.Double(t.%s)", strings.Join(parts, ""), strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.Double()", strings.Join(parts, ""), strings.Title(parts[0])))
						case "string":
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.String(t.%s)", strings.Join(parts, ""), strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.String()", strings.Join(parts, ""), strings.Title(parts[0])))
						default:
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.Bytes(t.%s.Encode())", strings.Join(parts, ""), strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.Object()", strings.Join(parts, ""), strings.Title(parts[0])))
						}
					case 6:
						// available_min_id:flags.9?int
						switch parts[5] {
						// tag_code:Vector<int>
						case ">":
							switch parts[4] {
							case "int":
								notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.VectorInt(t.%s)", strings.Join(parts, ""), strings.Title(parts[0])))
								notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.VectorInt()", strings.Join(parts, ""), strings.Title(parts[0])))
							case "long":
								notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.VectorLong(t.%s)", strings.Join(parts, ""), strings.Title(parts[0])))
								notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.VectorLong()", strings.Join(parts, ""), strings.Title(parts[0])))
							case "double":
								notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.VectorDouble(t.%s)", strings.Join(parts, ""), strings.Title(parts[0])))
								notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.VectorDouble()", strings.Join(parts, ""), strings.Title(parts[0])))
							case "string":
								notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.VectorString(t.%s)", strings.Join(parts, ""), strings.Title(parts[0])))
								notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.VectorString()", strings.Join(parts, ""), strings.Title(parts[0])))
							default:
								notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nx.Vector(t.%s)", strings.Join(parts, ""), strings.Title(parts[0])))
								notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nt.%s = d.Vector()", strings.Join(parts, ""), strings.Title(parts[0])))
							}
						case "int":
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.Int(t.%s)\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.Int()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
						case "long":
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.Long(t.%s)\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.Long()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
						case "double":
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.Double(t.%s)\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.Double()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
						case "string":
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.String(t.%s)\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.String()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
						case "true":
							// available_min_id:flags.9?true
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							boolsByEncode = append(boolsByEncode, fmt.Sprintf("// %s\nif t.%s {\nt.%s |= %d\n} else {\nt.%s &^=%d\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[0]),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[2]),
								1<<bitNum))
							boolsByDecode = append(boolsByDecode, fmt.Sprintf("// %s\nt.%s = (t.%s & %d) != 0",
								strings.Join(parts, ""),
								strings.Title(parts[0]),
								strings.Title(parts[2]),
								1<<bitNum))
						default:
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.Bytes(t.%s.Encode())\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.Object()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
						}
					case 9:
						// order:flags.0?Vector<DialogPeer>
						switch parts[7] {
						case "int":
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.VectorInt(t.%s)\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.VectorInt()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
						case "long":
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.VectorLong(t.%s)\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.VectorLong()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
						case "double":
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.VectorDouble(t.%s)\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.VectorDouble()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
						case "string":
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.VectorString(t.%s)\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.VectorString()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
						default:
							bitNum, err := strconv.Atoi(parts[3][1:])
							if err != nil {
								panic(err)
							}
							notBoolsByEncode = append(notBoolsByEncode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nx.Vector(t.%s)\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))
							notBoolsByDecode = append(notBoolsByDecode, fmt.Sprintf("// %s\nif (t.%s & %d) != 0 {\nt.%s=d.Vector()\n}\n",
								strings.Join(parts, ""),
								strings.Title(parts[2]),
								1<<bitNum,
								strings.Title(parts[0])))

						}
					default:
						panic(lineByTrim)
					}
				}

				typeBuf1.WriteString(fmt.Sprintf(fmtByTypeEncodeFunc,
					typeName, crcIdName, strings.Join(boolsByEncode, "\n"), strings.Join(notBoolsByEncode, "\n")))
				typeBuf1.WriteString(fmt.Sprintf(fmtByTypeDecodeFunc,
					typeName, strings.Join(boolsByDecode, "\n"), strings.Join(notBoolsByDecode, "\n")))
			}

			typeBuf1.WriteString(fmt.Sprintf("// end of '%s'\n\n", lineByTrim))
			objs = append(objs, typeBuf1.String())
		}
	}

	// output all
	{
		output := bytes.NewBuffer(nil)
		output.WriteString("You_should_modify_the_package_name_and_format_the_code_yourself\n\n")
		output.WriteString(fmt.Sprintf("import (\n%s\n)\n\n", "\"github.com/qingsong-he/mtprotobuf\""))
		output.WriteString(fmt.Sprintf("var (\n%s\n)\n\n", strings.Join(crc32Ids, "\n")))
		output.WriteString(fmt.Sprintf("func init() {\n%s\n}\n\n", strings.Join(inits, "\n")))
		for _, v := range objs {
			output.WriteString(v + "\n\n")
		}
		err := ioutil.WriteFile(filepath.Join(filepath.Dir(os.Args[1]), filepath.Base(os.Args[1])+".go"), output.Bytes(), 0600)
		if err != nil {
			panic(err)
		}
	}
}

var fmtByInit = `
	if _, has := mtprotobuf.DefaultDecodeMap[%s]; !has {
		mtprotobuf.DefaultDecodeMap[%s] = func(m *mtprotobuf.DecodeBuf) mtprotobuf.TL {
			return New%s().Decode(m)
		}
	} else {
		panic(mtprotobuf.TLErrByTLRepeatReg)
	}
`
var fmtByTypeFunc = `
func New%s() *%s {
	return new(%s)
}

func (*%s) CRC32() uint32 {
	return %s
}

func (*%s) GetLayer() int32 {
	return %s
}

`

var fmtByTypeEncodeFunc = `
func (t *%s) Encode() []byte {
x := mtprotobuf.NewEncodeBuf(256)
x.UInt(%s)
%s
%s
	return x.GetBuf()
}

`

var fmtByTypeDecodeFunc = `
func (t *%s) Decode(d *mtprotobuf.DecodeBuf) mtprotobuf.TL {
%s
%s
	return t
}

`
