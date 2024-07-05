package main

import (
	"fmt"
	"testing"
)

func TestGeneratePlayer(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		want := 3
		got := len(GeneratePlayers(3, 6))
		if got != want {
			t.Errorf("I want %d but I got %d", want, got)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		want := "P2"
		got := GeneratePlayers(3, 6)[1].Name
		if got != want {
			t.Errorf("I want %s but I got %s", want, got)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		want := 6
		got := len(GeneratePlayers(3, 6)[1].Dice)
		if got != want {
			t.Errorf("I want %d but I got %d", want, got)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		want := "P2"
		got := GeneratePlayers(3, 6)[0].Next.Name
		if got != want {
			t.Errorf("I want %s but I got %s", want, got)
		}
	})

	t.Run("Case 5", func(t *testing.T) {
		want := "P3"
		got := GeneratePlayers(3, 6)[1].Next.Name
		if got != want {
			t.Errorf("I want %s but I got %s", want, got)
		}
	})

	t.Run("Case 6", func(t *testing.T) {
		want := "P1"
		got := GeneratePlayers(3, 6)[2].Next.Name
		if got != want {
			t.Errorf("I want %s but I got %s", want, got)
		}
	})

	t.Run("Case 7", func(t *testing.T) {
		want := "P3"
		got := GeneratePlayers(3, 6)[0].Prev.Name
		if got != want {
			t.Errorf("I want %s but I got %s", want, got)
		}
	})

	t.Run("Case 8", func(t *testing.T) {
		want := "P1"
		got := GeneratePlayers(3, 6)[1].Prev.Name
		if got != want {
			t.Errorf("I want %s but I got %s", want, got)
		}
	})

	t.Run("Case 9", func(t *testing.T) {
		want := "P2"
		got := GeneratePlayers(3, 6)[2].Prev.Name
		if got != want {
			t.Errorf("I want %s but I got %s", want, got)
		}
	})

	t.Run("Case 10", func(t *testing.T) {
		players := GeneratePlayers(3, 6)
		for _, v := range players {
			fmt.Println(v.Name)
			fmt.Println(v.Score)
			fmt.Println(v.Dice)
			fmt.Println(v.Next.Name)
			fmt.Println(v.Prev.Name)
			fmt.Println()
		}
	})
}

func TestRoll(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		players := GeneratePlayers(3, 6)
		Roll(players)
		if players[0].Dice[5] == 0 {
			t.Errorf("I want > 0 but I got %d", 0)
		}
	})
}

func TestEvaluate(t *testing.T) {

	t.Run("Sub test 1", func(t *testing.T) {

		players := GeneratePlayers(3, 4)
		players[0].Dice = []int{3, 6, 1, 3}
		players[1].Dice = []int{2, 4, 5, 5}
		players[2].Dice = []int{1, 2, 5, 6}
		Evaluate(players)

		want := 1
		got := players[0].Score
		if got != want {
			t.Errorf("Case 1: I want %v but I got %v", want, got)
		}

		want = 0
		got = players[1].Score
		if got != want {
			t.Errorf("Case 2: I want %v but I got %v", want, got)
		}

		want = 1
		got = players[2].Score
		if got != want {
			t.Errorf("Case 3: I want %v but I got %v", want, got)
		}

		want = 1
		got = players[0].Dice[2]
		if got != want {
			t.Errorf("Case 4: I want %v but I got %v", want, got)
		}

		want = 1
		got = players[1].Dice[4]
		if got != want {
			t.Errorf("Case 5: I want %v but I got %v", want, got)
		}
	})

	t.Run("Sub test 2", func(t *testing.T) {
		players := GeneratePlayers(3, 4)
		players[0].Dice = []int{6, 1}
		players[0].Score = 2
		players[1].Dice = []int{2, 5, 6, 4, 6}
		players[1].Score = 0
		players[2].Dice = []int{1}
		players[2].Score = 2
		Evaluate(players)

		want := 3
		got := players[0].Score
		if got != want {
			t.Errorf("Case 1: I want %v but I got %v", want, got)
		}

		want = 2
		got = players[1].Score
		if got != want {
			t.Errorf("Case 2: I want %v but I got %v", want, got)
		}

		want = 2
		got = players[2].Score
		if got != want {
			t.Errorf("Case 3: I want %v but I got %v", want, got)
		}

		want = 0
		got = len(players[2].Dice)
		if got != want {
			t.Errorf("Case 4: I want %v but I got %v", want, got)
		}

		want = 2
		got = players[1].Dice[0]
		if got != want {
			t.Errorf("Case 5: I want %v but I got %v", want, got)
		}

	})

	t.Run("Sub test 3", func(t *testing.T) {
		players := GeneratePlayers(3, 4)
		players[0].Dice = []int{1}
		players[0].Score = 3
		players[1].Dice = []int{3, 4, 5, 5}
		players[1].Score = 2
		players[2].Dice = []int{}
		players[2].Score = 2
		Evaluate(players)

		want := 0
		got := len(players[0].Dice)
		if got != want {
			t.Errorf("Case 1: I want %v but I got %v", want, got)
		}

		want = 0
		got = len(players[2].Dice)
		if got != want {
			t.Errorf("Case 2: I want %v but I got %v", want, got)
		}

		want = 3
		got = players[0].Score
		if got != want {
			t.Errorf("Case 3: I want %v but I got %v", want, got)
		}

		want = 2
		got = players[2].Score
		if got != want {
			t.Errorf("Case 4: I want %v but I got %v", want, got)
		}

	})

}

func TestAddScore(t *testing.T) {
	t.Run("Sub test 1", func(t *testing.T) {
		players := GeneratePlayers(3, 4)
		players[0].Dice = []int{6, 1}
		players[0].Score = 2
		players[1].Dice = []int{2, 5, 6, 4, 6}
		players[1].Score = 0
		players[2].Dice = []int{1}
		players[2].Score = 2
		Evaluate(players)

		score := map[string]int{}
		players = AddScore(players, score)

		want := 2
		got := len(players)
		if got != want {
			t.Errorf("Case 1: I want %v but I got %v", want, got)
		}
	})

	t.Run("Sub test 2", func(t *testing.T) {
		players := GeneratePlayers(3, 4)
		players[0].Dice = []int{1}
		players[0].Score = 3
		players[1].Dice = []int{3, 4, 5, 5}
		players[1].Score = 2
		players[2].Dice = []int{}
		players[2].Score = 2
		Evaluate(players)

		score := map[string]int{}
		players = AddScore(players, score)

		want := 1
		got := len(players)
		if got != want {
			t.Errorf("Case 1: I want %v but I got %v", want, got)
		}
	})
}

func TestGetWinner(t *testing.T) {
	score := map[string]int{
		"P1": 3,
		"P2": 2,
		"P3": 2,
	}

	want := "P1"
	got := GetWinner(score)
	if got != want {
		t.Errorf("I want %s but I got %s", want, got)
	}
}
