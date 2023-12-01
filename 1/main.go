package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var digits = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var stringDigits = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var stringToDigit = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://adventofcode.com/2023/day/1/input", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Cookie",
		"session="+os.Getenv("AOC_COOKIE"))
	get, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(get.Body)
	body, err := io.ReadAll(get.Body)
	if err != nil {
		fmt.Println(err)
	}
	arr := strings.Split(string(body), "\n")
	fmt.Println(part1(arr))
}

func part1(input []string) int {
	start := time.Now()
	totalChannel := make(chan int)
	total := 0
	for _, i := range input {
		go combineFirstLast(totalChannel, i)
	}
	for range input {
		total += <-totalChannel
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	return total
}

func combineFirstLast(c chan int, input string) {
	if len(input) == 0 {
		c <- 0
		return
	}
	first := findFirstV2(input)
	last := findLastV2(input)

	result := string(first[0]) + string(last[0])
	newResult, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println(err)
	}
	c <- newResult
}

func findFirst(input string) string {
	for i, v := range input {
		for _, d := range digits {
			if string(v) == d {
				return string(input[i])
			}
		}
	}
	panic("No digit found in string: " + input)
}

func findLast(input string) string {
	for i := len(input) - 1; i >= 0; i-- {
		for _, d := range digits {
			if string(input[i]) == d {
				return string(input[i])
			}
		}
	}
	panic("No digit found in string: " + input)
}

func findFirstV2(input string) string {
	numberIndex := 999999999999999999
	solution := "temp"
	solutionFound := false
	for i, v := range input {
		for _, d := range digits {
			if string(v) == d && numberIndex > i {
				solution = string(v)
				numberIndex = i
				solutionFound = true
				fmt.Println(" first found: " + solution + "in: " + input)
			}
			if solutionFound {
				break
			}
		}
		if solutionFound {
			break
		}
	}
	for _, v := range stringDigits {
		stringIndex := strings.Index(input, v)
		if stringIndex == -1 {
			continue
		}
		if stringIndex < numberIndex {
			solution = stringToDigit[v]
			numberIndex = stringIndex
			fmt.Println(" first found: " + solution + "in: " + input)
		}
	}

	if solution == "temp" {
		panic("No digit found in string: " + input)
	}
	return solution

}

func findLastV2(input string) string {
	numberIndex := -1
	solution := "temp"
	solutionFound := false
	for i := len(input) - 1; i >= 0; i-- {
		for _, d := range digits {
			if string(input[i]) == d {
				solution = string(input[i])
				numberIndex = i
				solutionFound = true
			}
			if solutionFound {
				break
			}
		}
		if solutionFound {
			break
		}
	}
	solutionFound = false
	for _, v := range stringDigits {
		stringIndex := strings.LastIndex(input, v)
		if stringIndex > numberIndex {
			solution = stringToDigit[v]
			numberIndex = stringIndex
		}
	}
	fmt.Println(" last found: " + solution + "in: " + input)
	if solution == "temp" {
		panic("No digit found in string: " + input)
	}
	return solution
}
