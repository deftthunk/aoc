package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var pl = fmt.Println

type Card struct {
	number   int
	winners  []int
	possesed []int
	points   int
}

func input() []string {
	fh, err := os.Open("../input.txt")
	if err != nil {
		pl("Failed to open file: ", err)
		os.Exit(1)
	}
	defer fh.Close()

	text := make([]string, 0)
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func splitOffCardNumber(line string) (int, string) {
	result := strings.Split(line, ":")
	re := regexp.MustCompile(`\d+$`)
	test := re.FindString(result[0])

	cardNumber, err := strconv.Atoi(test)
	if err != nil {
		pl("Atoi CardNumber Fail: ", err)
		os.Exit(1)
	}
	return cardNumber, result[1]
}

func convertSlice(target []string) []int {
	newSlice := make([]int, 0)
	for _, val := range target {
		if val == "" {
			continue
		}
		number, err := strconv.Atoi(val)
		if err != nil {
			pl("Atoi Fail: ", err)
			os.Exit(1)
		}
		newSlice = append(newSlice, number)
	}
	return newSlice
}

func splitWinnerPossessed(line string) ([]int, []int) {
	result := strings.Split(line, "|")
	winnersRaw, possessedRaw := strings.TrimSpace(result[0]), strings.TrimSpace(result[1])
	winners := strings.Split(winnersRaw, " ")
	possessed := strings.Split(possessedRaw, " ")
	return convertSlice(winners), convertSlice(possessed)
}

func (card *Card) getCardPointCandidates() []int {
	pointCandidates := make([]int, 0)
	for _, possess := range card.possesed {
		if slices.Contains(card.winners, possess) {
			pointCandidates = append(pointCandidates, possess)
		}
	}
	return pointCandidates
}

func (card *Card) calcCardPoints() {
	for range card.getCardPointCandidates() {
		if card.points == 0 {
			card.points = 1
		} else {
			card.points *= 2
		}
	}
}

func createCards(input []string) []Card {
	cards := make([]Card, 0)
	for _, line := range input {
		cardNumber, lineRemainder := splitOffCardNumber(line)
		winners, possessed := splitWinnerPossessed(lineRemainder)
		card := Card{
			cardNumber,
			winners,
			possessed,
			0,
		}
		card.calcCardPoints()
		cards = append(cards, card)
	}
	return cards
}

func total(cards []Card) int {
	total := 0
	for _, card := range cards {
		total += card.points
	}
	return total
}

func main() {
	rawInput := input()
	cards := createCards(rawInput)
	total := total(cards)
	pl(total)
}
