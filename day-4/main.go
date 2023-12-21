package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const INPUT_FILE = "input_1.txt"

type set[T comparable] struct {
	elements []T // all elements are distinct
}

func NewSet[T comparable]() *set[T] {
	s := set[T]{}
	s.elements = make([]T, 0)

	return &s
}

func NewSetFromSlice[T comparable](elems []T) *set[T] {
	s := set[T]{}
	s.elements = make([]T, 0)
	for _, a := range elems {
		s.Add(a)
	}
	return &s
}

func (t *set[T]) Add(a T) {
	i := slices.Index[[]T, T](t.elements, a)
	if i == -1 {
		t.elements = append(t.elements, a)
	}
}
func (t *set[T]) In(a T) bool {
	i := slices.Index[[]T, T](t.elements, a)
	if i == -1 {
		return false
	} else {
		return true
	}
}

func (t *set[T]) Remove(a T) bool {
	i := slices.Index[[]T, T](t.elements, a)
	if i == -1 {
		return false
	} else {
		t.elements = slices.Delete[[]T, T](t.elements, i, i+1)
		return true
	}
}

func (t *set[T]) Intersection(o *set[T]) *set[T] {
	s := NewSet[T]()
	for _, a := range o.elements {
		if t.In(a) {
			s.Add(a)
		}
	}
	for _, b := range t.elements {
		if o.In(b) {
			s.Add(b)
		}
	}
	return s
}
func (t *set[T]) Union(o *set[T]) *set[T] {
	s := NewSetFromSlice[T](t.elements)
	for _, a := range o.elements {
		s.Add(a)
	}
	return s
}

func (t *set[T]) Difference(o *set[T]) *set[T] {
	s := NewSetFromSlice[T](t.elements)
	for _, a := range o.elements {
		s.Remove(a)
	}
	return s
}

func (t *set[T]) Len() int {
	return len(t.elements)
}

func (t *set[T]) Equal(o set[T]) bool {
	// two sets are equal if all the elements are equal (unordered)
	if len(t.elements) != len(o.elements) {
		return false
	}
	for _, a := range t.elements {
		found := false
		for _, b := range o.elements {
			if a == b {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true

}
func FromStringToIntSet(line string) *set[int] {
	sepLine := strings.Split(strings.Trim(line, " "), " ")
	intSet := NewSet[int]()
	for _, strNum := range sepLine {
		if strNum == "" {
			// ignore empty strings
			continue
		}
		num, err := strconv.Atoi(strNum)
		if err != nil {
			log.Fatalln(err)
		}
		intSet.Add(num)
	}
	return intSet
}

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
		// next line is for getting id, (not necessary)
		log.Println(sepLine[0] + ",")
		// I have to find a way to compress multiple black spaces to one, so I can split it easily
		log.Println(len(strings.Split(sepLine[0], " ")))
		log.Println(strings.Trim(strings.Split(strings.Trim(sepLine[0], " "), " ")[1], " "))
		_, err := strconv.Atoi(strings.Trim(strings.Split(sepLine[0], " ")[1], "\t"))
		if err != nil {
			log.Fatalln(err)
		}

		sepNums := strings.Split(sepLine[1], "|")
		winningNums := FromStringToIntSet(sepNums[0])
		myNums := FromStringToIntSet(sepNums[1])
		intersectNums := winningNums.Intersection(myNums)
		if intersectNums.Len() != 0 {
			result += int(math.Pow(2, float64(intersectNums.Len()-1)))
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
