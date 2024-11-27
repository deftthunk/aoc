package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

var pl = fmt.Println

func input() []string {
	fh, err := os.Open("input.txt")
	if err != nil {
		pl("Failed file open")
		os.Exit(1)
	}
	defer fh.Close()

	var text []string
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func matrixInitialize[T int | string](x, y int, mType T) [][]T {
	matrix := make([][]T, x)
	for column, _ := range matrix {
		matrix[column] = make([]T, y)
	}
	return matrix
}

func matrixize(rawText []string) [][]string {
	matrix := matrixInitialize(len(rawText), len(rawText), "")
	for row, line := range rawText {
		for column, char := range line {
			matrix[row][column] = string(char)
		}
	}
	return matrix
}

func makeMask(matrix [][]string, part2 bool) [][]int {
	integers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	mask := matrixInitialize(len(matrix), len(matrix), 0)
	for row, line := range matrix {
		for column, char := range line {
			switch true {
			case char == ".":
				mask[row][column] = 0
			case slices.Contains(integers, char):
				mask[row][column] = 1
			default:
				if part2 && char == "*" {
					mask[row][column] = 3
				} else {
					mask[row][column] = 2
				}
			}
		}
	}
	return mask
}

func calcSymHash(x, y int) string {
	symString := make([]string, 0)
	//fmt.Printf("X: %d, Y: %d\n", x, y)
	for i := 0; i < y; i++ {
		symString = append(symString, inputMatrix[x][i])
	}
	hasher := sha1.New()
	io.WriteString(hasher, strings.Join(symString, ""))
	io.WriteString(hasher, strconv.Itoa(x))
	rawHash := hasher.Sum(nil)
	hash := base64.StdEncoding.EncodeToString(rawHash)
	return hash
}

func findAdjascent(x, y int, mask [][]int) ([][]int, string) {
	adjascents := make([][]int, 0)
	box := [][]int{
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
		{x + 1, y},
		{x + 1, y - 1},
		{x, y - 1},
	}

	for _, pos := range box {
		xPos, yPos := pos[0], pos[1]
		if mask[xPos][yPos] == 1 {
			adjascents = append(adjascents, pos)
		}
	}
	return adjascents, calcSymHash(x, y)
}

func symbolHunt(mask [][]int, symbolId int) map[string][][]int {
	symbolCoordinates := make(map[string][][]int, 0)
	for row, line := range mask {
		for column, val := range line {
			if val == symbolId {
				adjascents, symHash := findAdjascent(row, column, mask)
				if symbolId == 3 {
					symbolCoordinates[symHash] = append(symbolCoordinates[symHash], adjascents...)
				}
			}
		}
	}
	return symbolCoordinates
}

type Vector struct {
	start      int
	end        int
	row        int
	symbolHash string
}

func checkPos(pos, min, max int) bool {
	return min <= pos && pos < max
}

func lookAround(pos []int, mask [][]int) Vector {
	row := pos[0]
	begin, end := -1, -1
	sign := 1

	for look := 0; look < 10; look++ {
		column := pos[1] + (look * sign)
		validEnd := checkPos(column, 0, len(mask))
		if validEnd {
			maskVal := mask[row][column]
			if maskVal != 1 {
				if end == -1 {
					end = column - 1
				}
			}
		} else if end == -1 {
			end = column - 1
		}
		sign *= -1
		column = pos[1] + (look * sign)
		validBegin := checkPos(column, 0, len(mask))
		if validBegin {
			maskVal := mask[row][column]
			if maskVal != 1 {
				if begin == -1 {
					begin = column + 1
				}
			}
		} else if begin == -1 {
			begin = column + 1
		}
		sign *= -1
	}

	return Vector{
		begin,
		end,
		row,
		"",
	}
}

func getVectors(adjacents map[string][][]int, mask [][]int) map[Vector]int {
	vectors := make(map[Vector]int)
	for symHash, coordinates := range adjacents {
		for _, pos := range coordinates {
			v := lookAround(pos, mask)
			v.symbolHash = symHash
			vectors[v] = 1
		}
	}
	return vectors
}

func resolveNumbers(vectors map[Vector]int, matrix [][]string) map[string][]int {
	numbers := make(map[string][]int, 0)
	for vector, _ := range vectors {
		var strNumber []string
		for i := vector.start; i <= vector.end; i++ {
			//fmt.Printf("Number: %s, Row: %d\n", strNumber, vector.row)
			value := matrix[vector.row][i]
			strNumber = append(strNumber, value)
		}
		combined := strings.Join(strNumber, "")
		num, err := strconv.Atoi(combined)
		if err != nil {
			pl("Failed Atoi: ", err)
			os.Exit(1)
		}
		numbers[vector.symbolHash] = append(numbers[vector.symbolHash], num)
	}
	return numbers
}

func calcRatios(numbers map[string][]int) int {
	total := 0
	for _, intPairs := range numbers {
		ratio := 1
		if len(intPairs) == 2 {
			for _, val := range intPairs {
				ratio *= val
			}
			total += ratio
		}
	}
	return total
}

var inputMatrix [][]string

func main() {
	inputText := input()
	inputMatrix = matrixize(inputText)
	mask := makeMask(inputMatrix, true)
	adj := symbolHunt(mask, 3)
	vectors := getVectors(adj, mask)
	numbers := resolveNumbers(vectors, inputMatrix)
	answer := calcRatios(numbers)
	pl(answer)
}
