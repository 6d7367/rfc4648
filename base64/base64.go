package base64

// imports

import "fmt"
import "strconv"
import "strings"

// алфавит
var base64abc string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// кодирование строки в base64
func EncodeBase64(str string) string {
	r := ""

	chunks := []string{}
	acc := ""

	// превращаю строку в список нулей и единиц, 
	// каждая подстрока длиной 8 символов
	for _, c := range str {
		if len(acc) >= 24 {
			chunks = append(chunks, acc)
			acc = ""
		}

		acc += fmt.Sprintf("%08b", c)
	}

	// если что-то осталось в аккумуляторе, то добавляю 
	// к уже существующим подстрокам
	if len(acc) > 0 {
		chunks = append(chunks, acc)
	}

	// делаю срезы и получаю новые индексы алфавита для кодировки
	for _, c := range chunks {
		// если полученная строка не кратна 6, то ее необходимо дополнить нулями
		zeroPads, equalPads := helperGetZeroedPadding(len(c))
		if (len(c) % 6) != 0 {
			c += zeroPads
		}

		acc = ""

		indexOffset  := 6
		currentIndex := 0
		nextIndex    := currentIndex + indexOffset
		
		// в каком фрмате находятся числа в строке (bin, oct, dec, hex)
		base := 2
		// сколько бит в итоговом числе
		bitSize := 32


		for nextIndex <= len(c) {
			// конвертирую строку с нулями и единицами в 8-битное число
			// и получаю индекс 
			abcIndex, _ := strconv.ParseInt(c[currentIndex:nextIndex], base, bitSize)
			acc += string(base64abc[abcIndex])
			currentIndex = nextIndex
			nextIndex += indexOffset
		}

		// выравниваю строку до 24 бит (с помощью =)
		acc += equalPads

		r += acc
	}

	return r
}

// вспомошгательная функция
// возвращает строки с необходимым выравниванием
// (0) и (=)
func helperGetZeroedPadding(currLen int) (string, string) {
	
	nextLen := currLen
	zeroPads  := ""
	equalPads := ""
	for true {
		if (nextLen % 6) != 0 {
			nextLen += 1
		} else {
			break
		}
	}

	for len(zeroPads) < (nextLen - currLen) {
		zeroPads += "0"
	}

	equalPadsCurr := currLen / 8
	equalPadsRes  := 3

	for equalPadsCurr < equalPadsRes {
		equalPads += "="
		equalPadsCurr += 1
	}

	return zeroPads, equalPads
}


// декодирует строку из base64
func DecodeBase64(str string) string {
	// очищаю строку от любых неалфавитных символов
	str = helperClearNonAbcSymbols(str)

	// возвращаемый результат, массив байтов
	r := []byte{}

	// временное хранилище подстрок, по 4 символа
	chunks := []string{}
	acc := ""

	for _, c := range str {
		if len(acc) == 4 {
			chunks = append(chunks, acc)
			acc = ""
		}
		acc += string(c)
	}

	if len(acc) != 0 {
		chunks = append(chunks, acc)
	}

	// обхожу полученные подстроки и декодирую их
	for _, chunk := range chunks {

		acc = ""

		// прохожу по полученной подстроке и 
		// выбираю положение закодированного символа в алфавите bas64
		// затем привожу его к виду 6-битного значение
		// и добавляю к аккумулятору
		for _, c := range chunk {
			acc += fmt.Sprintf("%06b", strings.Index(base64abc, string(c)))
		}

		// полученная строка из 4 6-битных символово это на самом деле
		// 3 8-битных символа

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
// очищает строку от всех неалфавитных (для base64) символов
func helperClearNonAbcSymbols(str string) string {
	tmpAbc := base64abc + "="
	r := ""

	for _, c := range str {
		if strings.Contains(tmpAbc, string(c)){
			r += string(c)
		}
	}

	return r
}