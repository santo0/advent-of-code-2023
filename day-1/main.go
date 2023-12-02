package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const INPUT_FILE = "input_1.txt"

var digitMap = map[string]byte{"one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'}

func getDigitFromString(word string) (byte, bool) {
	for k := range digitMap {
		if len(k) <= len(word) && word[:len(k)] == k {
			return digitMap[word[:len(k)]], true
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
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		var digits []byte
		for i := 0; i < len(line); i++ {
			if '1' <= line[i] && line[i] <= '9' {
				digits = append(digits, line[i])
			} else if digit, ok := getDigitFromString(line[i:]); ok {
				digits = append(digits, digit)
			}
		}

		value, err := strconv.Atoi(string([]byte{digits[0], digits[len(digits)-1]}))
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
