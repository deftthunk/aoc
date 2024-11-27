package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var pl = fmt.Println

func input() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		pl(err)
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func parse(text []string) []string {
	num_all := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"0":     "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	pairs := make([]string, 0)
	for _, word := range text {
		indexes := make(map[int]string)
		for key, _ := range num_all {
			matches := findAllOccurances(word, key)
			if len(matches) > 0 {
				maps.Copy(indexes, matches)
			}
		}

		firstLastIndicies := getFirstLast(indexes)

		firstNum := num_all[indexes[firstLastIndicies[0]]]
		lastNum := num_all[indexes[firstLastIndicies[1]]]
		concat := strings.Builder{}
		concat.WriteString(firstNum)
		concat.WriteString(lastNum)

		pairs = append(pairs, concat.String())
	}
	return pairs
}

func getFirstLast(indicies map[int]string) []int {
	keys := lo.Keys(indicies)
	slices.Sort(keys)
	return []int{keys[0], keys[len(keys)-1]}
}

func findAllOccurances(target, substr string) map[int]string {
	count := strings.Count(target, substr)
	if count > 1 {
		indicies := make(map[int]string, 0)
		target_slice := strings.Split(target, "")
		abs_ptr := 0
		for {
			new_target := strings.Join(target_slice[abs_ptr:], "")
			result := strings.Index(new_target, substr)
			if result != -1 {
				//indicies = append(indicies, abs_ptr+result)
				indicies[abs_ptr+result] = substr
				abs_ptr += result + 1
			} else {
				break
			}
		}
		return indicies
	} else if count == 1 {
		return map[int]string{strings.Index(target, substr): substr}
	} else {
		return map[int]string{}
	}
}

func summation(pairs []string) int {
	var sum int
	for _, item := range pairs {
		value, err := strconv.Atoi(item)
		if err != nil {
			pl("Failed Atoi conversion")
		} else {
			sum += value
		}
	}
	return sum
}

func main() {
	raw_text := input()
	pairs := parse(raw_text)
	sum := summation(pairs)
	pl(sum)

}
