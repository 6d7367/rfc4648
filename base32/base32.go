package base32

// imports

import "fmt"
import "strconv"
import "strings"

var base32abc string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

// кодирование строки в base32
func EncodeBase32(str string) string {
	acc := ""
	chunks := []string{}
	r := ""

	for _, c := range str {
		if len(acc) == 5 {
			chunks = append(chunks, acc)
			acc = ""
		}
		acc += string(c)
	}

	if len(acc) > 0 {
		chunks = append(chunks, acc)
	}

	for _, chunk := range chunks {


		acc = ""
		for _, c := range chunk {
			acc += fmt.Sprintf("%08b", c)
		}

		zeroPads, equalPads := helperGetZeroedPadding(len(acc))
		if len(acc) < 40 {
			acc += zeroPads
		}

		currentIndex := 0
		offset := 5
		nextIndex := currentIndex + offset
		base := 2
		bitSize := 8

		for nextIndex <= len(acc) {
			byteVal, _ := strconv.ParseUint(acc[currentIndex:nextIndex], base, bitSize)
			currentIndex = nextIndex
			nextIndex += offset

			r += string(base32abc[byteVal])
		}

		r += equalPads
	}

	return r
}

// вспомогательная функция
// расчитывает необходимое количество нулей для выравнивания битов 
// и необходимое количество "=" для оббщего выравнивания строки
func helperGetZeroedPadding(currLen int) (string, string) {
	
	nextLen := currLen
	zeroPads  := ""
	equalPads := ""
	for true {
		if (nextLen % 10) != 0 {
			nextLen += 1
		} else {
			break
		}
	}

	for len(zeroPads) < (nextLen - currLen) {
		zeroPads += "0"
	}

	equalPadsCurr := nextLen / 5
	equalPadsRes  := 40 / 5

	for equalPadsCurr < equalPadsRes {
		equalPads += "="
		equalPadsCurr += 1
	}

	return zeroPads, equalPads
}

// декодирование из base32
// аналогична функции rfc4648.base64.DecodeBase64(string) string
func DecodeBase32(str string) string {

	str = helperClearNonAbcSymbols(str)

	r := []byte{}

	chunks := []string{}
	acc := ""

	for _, c := range str {
		if len(acc) == 8 {
			chunks = append(chunks, acc)
			acc = ""
		}
		acc += string(c)
	}

	if len(acc) > 0 {
		chunks = append(chunks, acc)
	}

	for _, chunk := range chunks {
		acc = ""

		for _, c := range chunk {
			acc += fmt.Sprintf("%05b", strings.Index(base32abc, string(c)))
		}

		indexOffset  := 8
		currentIndex := 0
		nextIndex    := currentIndex + indexOffset
		
		base := 2
		bitSize := 8

		// обхожу строку и получаю срез, состоящий из 8 нулей и единиц
		// затем превращаю его в обычное число и добавляю в список результатов
		for nextIndex <= len(acc) {
			byteVal, _ := strconv.ParseInt(acc[currentIndex:nextIndex], base, bitSize)
			currentIndex = nextIndex
			nextIndex += indexOffset
			r = append(r, byte(byteVal))
		}
	}

	return string(r)
	
}

// вспомогательная функция
// очищает строку от неалфавитных символов
func helperClearNonAbcSymbols(str string) string {
	r := ""

	for _, c := range str {
		if strings.Contains(base32abc, string(c)){
			r += string(c)
		}
	}

	return r
}