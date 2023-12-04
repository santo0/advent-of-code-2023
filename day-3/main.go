package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const INPUT_FILE = "input_1.txt"

type position struct {
	char rune
	i    int
	j    int
}

type PositionRange struct {
	positions []position
	number    int
}

func (this PositionRange) equals(other PositionRange) bool {
	if len(this.positions) != len(other.positions) || this.number != other.number {
		return false
	}
	for _, a := range this.positions {
		aEqual := false
		for _, b := range other.positions {
			if a.char == b.char && a.i == b.i && a.j == b.j {
				aEqual = true
			}
		}
		if !aEqual {
			return false
		}
	}
	for _, a := range other.positions {
		aEqual := false
		for _, b := range this.positions {
			if a.char == b.char && a.i == b.i && a.j == b.j {
				aEqual = true
			}
		}
		if !aEqual {
			return false
		}
	}
	return true
}

func getNumberInPositionRange(positions []position) int {
	runeSlice := make([]rune, 0)
	for _, pos := range positions {
		runeSlice = append(runeSlice, pos.char)
	}
	result, err := strconv.Atoi(string(runeSlice))
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func NewPositionRange(positions []position) PositionRange {
	return PositionRange{positions, getNumberInPositionRange(positions)}
}

func isPositionAdjacent(sPos position, tPos position) bool {
	return sPos.i-1 == tPos.i && sPos.j-1 <= tPos.j && sPos.j+1 >= tPos.j ||
		sPos.i+1 == tPos.i && sPos.j-1 <= tPos.j && sPos.j+1 >= tPos.j ||
		sPos.i == tPos.i && (sPos.j+1 == tPos.j || sPos.j-1 == tPos.j)
}

func isRangeAdjacent(posRange PositionRange, trgPositions []position) bool {
	for _, sPos := range posRange.positions {
		for _, tPos := range trgPositions {
			if isPositionAdjacent(sPos, tPos) {
				return true
			}
		}
	}
	return false
}

func p1() {
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	i := 0
	lastLineSyms := make([]position, 0)
	currentLineSyms := make([]position, 0)
	lastLineNums := make([]PositionRange, 0)
	currentLineNums := make([]PositionRange, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numberBuffer := make([]position, 0)
		numDetected := false
		for j, c := range line {
			if c >= '0' && c <= '9' {
				// detect seq numbers
				numberBuffer = append(numberBuffer, position{c, i, j})
				numDetected = true
			} else {
				if c != '.' {
					// keep positions of special chars
					currentLineSyms = append(currentLineSyms, position{c, i, j})
				}
				if numDetected {
					currentLineNums = append(currentLineNums, NewPositionRange(numberBuffer))
					numberBuffer = make([]position, 0)
					numDetected = false
				}
			}

		}

		if numDetected {
			currentLineNums = append(currentLineNums, NewPositionRange(numberBuffer))
			numDetected = false
		}

		// when finished seq numbers, isAdjacent?
		rangeToRemove := make([]PositionRange, 0)
		for _, currentRange := range currentLineNums {
			// optimizations can be done (we know that lastLine can't be in same i or j+1 and currentLine always has same i)
			if isRangeAdjacent(currentRange, lastLineSyms) || isRangeAdjacent(currentRange, currentLineSyms) {
				// calculate number in srcPositions
				rangeToRemove = append(rangeToRemove, currentRange)
				// log number
				log.Println("Number=", currentRange.number)
				// add number in result
				result += currentRange.number
			}
		}
		for _, rangeToDel := range rangeToRemove {
			indexToDel := -1
			for i, aRange := range currentLineNums {
				if &rangeToDel == &aRange {
					indexToDel = i
					break
				}
			}
			if indexToDel >= 0 {
				currentLineNums = append(currentLineNums[:indexToDel], currentLineNums[indexToDel+1:]...)
			}
		}
		for _, lastRange := range lastLineNums {
			if isRangeAdjacent(lastRange, currentLineSyms) {
				// calculate number in srcPositions
				// log number
				log.Println("Number=", lastRange.number)
				// add number in result
				result += lastRange.number
			}
		}

		//lo mateix pero amb syms
		lastLineSyms = currentLineSyms
		lastLineNums = currentLineNums
		currentLineSyms = make([]position, 0)
		currentLineNums = make([]PositionRange, 0)
		i += 1
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
	i := 0
	var gearMap = map[position][]PositionRange{}

	lastLineSyms := make([]position, 0)
	currentLineSyms := make([]position, 0)
	lastLineNums := make([]PositionRange, 0)
	currentLineNums := make([]PositionRange, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numberBuffer := make([]position, 0)
		numDetected := false
		for j, c := range line {
			if c >= '0' && c <= '9' {
				// detect seq numbers
				numberBuffer = append(numberBuffer, position{c, i, j})
				numDetected = true
			} else {
				if c != '.' {
					// keep positions of special chars
					currentLineSyms = append(currentLineSyms, position{c, i, j})
				}
				if numDetected {
					currentLineNums = append(currentLineNums, NewPositionRange(numberBuffer))
					numberBuffer = make([]position, 0)
					numDetected = false
				}
			}

		}

		if numDetected {
			currentLineNums = append(currentLineNums, NewPositionRange(numberBuffer))
			numDetected = false
		}

		for _, symPos := range currentLineSyms {
			if symPos.char == '*' {
				for _, lastRange := range lastLineNums {
					if isRangeAdjacent(lastRange, []position{symPos}) {
						numSlice, ok := gearMap[symPos]
						if ok {
							gearMap[symPos] = append(numSlice, lastRange)
						} else {
							gearMap[symPos] = []PositionRange{lastRange}
						}

					}
				}
				for _, currentRange := range currentLineNums {
					if isRangeAdjacent(currentRange, []position{symPos}) {
						numSlice, ok := gearMap[symPos]
						if ok {
							gearMap[symPos] = append(numSlice, currentRange)
						} else {
							gearMap[symPos] = []PositionRange{currentRange}
						}

					}
				}
			}
		}

		for _, symPos := range lastLineSyms {
			if symPos.char == '*' {
				for _, lastRange := range lastLineNums {
					if isRangeAdjacent(lastRange, []position{symPos}) {
						numSlice, ok := gearMap[symPos]
						if ok {
							gearMap[symPos] = append(numSlice, lastRange)
						} else {
							gearMap[symPos] = []PositionRange{lastRange}
						}

					}
				}
				for _, currentRange := range currentLineNums {
					if isRangeAdjacent(currentRange, []position{symPos}) {
						numSlice, ok := gearMap[symPos]
						if ok {
							gearMap[symPos] = append(numSlice, currentRange)
						} else {
							gearMap[symPos] = []PositionRange{currentRange}
						}

					}
				}
			}
		}

		lastLineSyms = currentLineSyms
		lastLineNums = currentLineNums
		currentLineSyms = make([]position, 0)
		currentLineNums = make([]PositionRange, 0)
		i += 1
	}

	for gear, numSlice := range gearMap {
		var uniqueNums = []PositionRange{}
		for _, num := range numSlice {
			isUnique := true
			for _, eNum := range uniqueNums {
				if num.equals(eNum) {
					isUnique = false
					break
				}
			}
			if isUnique {
				uniqueNums = append(uniqueNums, num)
			}
		}
		log.Println(gear, len(uniqueNums))
		if len(uniqueNums) == 2 {
			result += uniqueNums[0].number * uniqueNums[1].number
		}
	}
	log.Println("RESULT=", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//p1()
	p2()
}
