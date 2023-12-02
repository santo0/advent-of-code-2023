package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE = "input_1.txt"

var cubeQty = map[string]int{"red": 12, "green": 13, "blue": 14}

func p1() {
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		sepLine := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.Trim(strings.Split(sepLine[0], " ")[1], " "))
		if err != nil {
			log.Fatalln(err)
		}
		invalidSet := false
		for _, set := range strings.Split(sepLine[1], ";") {
			for _, subset := range strings.Split(set, ",") {
				sepSubset := strings.Split(strings.Trim(subset, " "), " ")
				color := strings.Trim(sepSubset[1], " ")
				qty, err := strconv.Atoi(strings.Trim(sepSubset[0], " "))
				if err != nil {
					log.Fatalln(err)
				}
				if cubeQty[color] < qty {
					invalidSet = true
					break
				}
			}
			if invalidSet {
				break
			}
		}
		if !invalidSet {
			result += id
		}
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
		line := scanner.Text()
		sepLine := strings.Split(line, ":")
		if err != nil {
			log.Fatalln(err)
		}
		var acc = map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, set := range strings.Split(sepLine[1], ";") {
			for _, subset := range strings.Split(set, ",") {
				sepSubset := strings.Split(strings.Trim(subset, " "), " ")
				color := strings.Trim(sepSubset[1], " ")
				qty, err := strconv.Atoi(strings.Trim(sepSubset[0], " "))
				if err != nil {
					log.Fatalln(err)
				}
				if acc[color] < qty {
					acc[color] = qty
				}
			}
		}
		power := 1
		for _, v := range acc {
			power *= v
		}
		result += power
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
