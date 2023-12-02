package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
)

const INPUT_FILE = "input_1.txt"

var digitMap = map[string]byte{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func getDigitFromString(word string) (byte, bool) {
	for k := range digitMap {
		k_len := len(k)
		if k_len > len(word) {
			continue
		}
		if word[:k_len] == k {
			digit, ok := digitMap[word[:k_len]]
			return digit, ok
		}
	}
	return '0', false
}

func main() {

	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int = 0
	for scanner.Scan() {
		line := scanner.Text()
		var digits []byte = make([]byte, 0)
		for i := 0; i < len(line); i++ {
			if '1' <= line[i] && line[i] <= '9' {
				digits = append(digits, line[i])
			} else {
				digit, ok := getDigitFromString(line[i:])
				if ok {
					digits = append(digits, digit)
				}
			}
		}

		value, err := strconv.Atoi(bytes.NewBuffer([]byte{digits[0], digits[len(digits)-1]}).String())
		if err != nil {
			log.Fatal(err)
		}
		result += value
	}

	log.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
