package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Node struct {
	Name  string
	Score int
	Dice  []int
	Out   []int
	Next  *Node
	Prev  *Node
}

func main() {

	dice := 4
	player := 3
	score := map[string]int{}
	players := GeneratePlayers(player, dice)

	for len(players) > 1 {
		Roll(players)
		Evaluate(players)
		players = AddScore(players, score)
	}

	fmt.Println("The winner is: ", GetWinner(score))
}

func GeneratePlayers(numberOfPlayer int, numberOfDice int) []Node {
	players := make([]Node, numberOfPlayer)

	for i := 0; i < len(players); i++ {
		players[i].SetName(i)
		players[i].SetDice(numberOfDice)
		players[i].SetNext(i, players)
		players[i].SetPrev(i, players)
	}

	return players
}

func Roll(players []Node) {
	for i := 0; i < len(players); i++ {
		players[i].Roll()
	}
}

func Evaluate(players []Node) {
	for i, player := range players {

		players[i].Out = []int{}
		newPlayerDice := []int{}

		for _, dice := range player.Dice {

			if dice == 6 {
				players[i].Score += 1
				continue
			}
			if dice == 1 {
				players[i].Out = append(players[i].Out, 1)
				continue
			}

			newPlayerDice = append(newPlayerDice, dice)
		}
		players[i].Dice = newPlayerDice

		if i == len(players)-1 {
			players[i].Dice = append(players[i].Dice, players[i].Prev.Out...)
			players[i].Next.Dice = append(players[i].Next.Dice, players[i].Out...)
		} else if i != 0 {
			players[i].Dice = append(players[i].Dice, players[i].Prev.Out...)
		}
	}
}

func AddScore(players []Node, score map[string]int) []Node {
	var newPlayers []Node
	for i := 0; i < len(players); i++ {
		if len(players[i].Dice) == 0 {
			score[players[i].Name] = players[i].Score
			continue
		}

		newPlayers = append(newPlayers, players[i])
	}

	if len(newPlayers) == 1 {
		score[newPlayers[0].Name] = newPlayers[0].Score
	}

	return newPlayers
}

func GetWinner(score map[string]int) string {

	winner := ""
	point := 0

	for player, scr := range score {
		if scr > point {
			point = scr
			winner = player
		}
	}

	return winner
}

func (p *Node) SetName(i int) {
	p.Name = "P" + strconv.Itoa(i+1)
}

func (p *Node) SetDice(numberOfDice int) {
	for i := 0; i < numberOfDice; i++ {
		p.Dice = make([]int, numberOfDice)
	}
}

func (p *Node) SetNext(i int, players []Node) {
	if i == len(players)-1 {
		p.Next = &players[0]
		return
	}
	p.Next = &players[i+1]
}

func (p *Node) SetPrev(i int, players []Node) {
	if i == 0 {
		p.Prev = &players[len(players)-1]
		return
	}
	p.Prev = &players[i-1]
}

func (p *Node) Roll() {
	for i := 0; i < len(p.Dice); i++ {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		p.Dice[i] = random.Intn(6) + 1
	}
}
