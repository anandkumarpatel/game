package engine

import (
	"testing"
)

type GetWinnerObj struct {
	p1     Move
	p2     Move
	winner int
}

var tests = []GetWinnerObj{
	{Move{1, 1}, Move{1, 1}, 0},
	{Move{3, 3}, Move{1, 1}, 1},
	{Move{1, 1}, Move{3, 3}, 2},
}

func TestGetWinner(t *testing.T) {
	for _, testObj := range tests {
		winner, err := getWinner(testObj.p1, testObj.p2)
		if err != nil {
			t.Error(err)
		}
		if winner != testObj.winner {
			t.Error("expected", testObj.winner, "got", winner)
		}
	}
}
