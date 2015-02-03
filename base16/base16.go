package base16

import "fmt"
import "strconv"

var base16abc string = "0123456789ABCDEF"

func EncodeBase16(str string) string {
	r := ""

	bin := ""

	base := 2
	biteSize := 8

	for _, c := range str {
		bin = fmt.Sprintf("%08b", c)
		byteVal1, _ := strconv.ParseInt(bin[0:4], base, biteSize)
		byteVal2, _ := strconv.ParseInt(bin[4:8], base, biteSize)

		r += fmt.Sprintf("%X%X", byteVal1, byteVal2)
	}

	return r
}

func DecodeBase16(str string) string {
	r := ""
	acc := ""

	for _, c := range str {
		baseDecode1 := 16
		bitSizeDecode1 := 8
		baseDecode2 := 2
		bitSizeDecode2 := 8
		byteVal, _ := strconv.ParseInt(string(c), baseDecode1, bitSizeDecode1)
		
		if len(acc) < 8 {
			acc += fmt.Sprintf("%04b", byteVal)
		} else {
			byteVal2, _ := strconv.ParseInt(acc, baseDecode2, bitSizeDecode2)
			acc = fmt.Sprintf("%04b", byte(byteVal))
			
			r += string(byteVal2)
		}
		
	}

	return r

}