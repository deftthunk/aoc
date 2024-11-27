package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

type Game struct {
	name  int
	red   int
	green int
	blue  int
	total int
}

func getInput() []string {
	file, err := os.Open("games.txt")
	if err != nil {
		pl(err)
		os.Exit(1)
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func carveName(line string) (name int, remainder string) {
	pieces := strings.Split(line, ":")
	if len(pieces) != 2 {
		pl("Found wrong number of pieces!")
		os.Exit(1)
	}
	nameNum := strings.Split(pieces[0], " ")
	name, err := strconv.Atoi(nameNum[1])
	if err != nil {
		pl("Failed name Atoi: ", err)
		os.Exit(1)
	}
	return name, pieces[1]
}

func (g *Game) calcTotal() {
	g.total = g.red + g.blue + g.green
}

func (g *Game) getColorCounts(round string) {
	pairs := strings.Split(round, ",")
	for _, p := range pairs {
		p = strings.TrimSpace(p)
		items := strings.Split(p, " ")
		count, err := strconv.Atoi(items[0])
		if err != nil {
			pl("Failed Atoi: ", err)
			os.Exit(1)
		}
		switch items[1] {
		case "red":
			if g.red < count {
				g.red = count
			}
		case "blue":
			if g.blue < count {
				g.blue = count
			}
		case "green":
			if g.green < count {
				g.green = count
			}
		default:
			pl("Unknown failure in switch")
			os.Exit(1)
		}
	}
}

func genGames(input []string) []Game {
	games := make([]Game, 0)
	for _, line := range input {
		name, remainder := carveName(line)
		rounds := strings.Split(remainder, ";")
		game := Game{name, 0, 0, 0, 0}
		for _, round := range rounds {
			game.getColorCounts(round)
			game.calcTotal()
		}
		games = append(games, game)
	}
	return games
}

func getConstraints() Game {
	g := Game{0, 12, 13, 14, 0}
	g.calcTotal()
	return g
}

func (g *Game) enoughRed(constraint int) bool {
	return g.red <= constraint
}

func (g *Game) enoughBlue(constraint int) bool {
	return g.blue <= constraint
}

func (g *Game) enoughGreen(constraint int) bool {
	return g.green <= constraint
}

func (g *Game) enoughCubes(constraint int) bool {
	return g.total <= constraint
}

func q1(games []Game) int {
	/*which games would be possible with only the cubes given in the constraints?*/
	constraint := getConstraints()
	pl("Constraint: ", constraint)
	candidates := make([]Game, 0)
	for _, game := range games {
		if game.enoughBlue(constraint.blue) &&
			game.enoughGreen(constraint.green) &&
			game.enoughRed(constraint.red) &&
			game.enoughCubes(constraint.total) {
			candidates = append(candidates, game)
		}
	}
	pl("Candidates: ", candidates)
	sum := 0
	for _, game := range candidates {
		sum += game.name
	}
	return sum
}

func (g *Game) getCube() int {
	return g.red * g.green * g.blue
}

func q2(games []Game) {
	/* what is the min number of cubes-colors required for each game. "cube" those numbers and sum them all up */
	result := 0
	for _, game := range games {
		cube := game.getCube()
		result += cube
	}
	pl("Answer 2: ", result)
}

func main() {
	inputText := getInput()
	games := genGames(inputText)
	answer := q1(games)
	pl(answer)
	q2(games)
}
