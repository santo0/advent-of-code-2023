package main

import (
	"bufio"
	"log"
	"os"
)

const INPUT_FILE = "input_1.txt"

func p1() {
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {

	}
	log.Println("RESULT=", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func p2() {
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {

	}
	log.Println("RESULT=", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	p1()
	p2()
}
