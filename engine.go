package engine

import (
	"errors"
)

const (
	TIE     = 0
	P1      = 1
	P2      = 2
	INVALID = -1
)

type Move struct {
	Type  int // 1, 2, 3 (3 beats 2 beats 1 beats 3)
	Level int // 1, 2, 3
}

// (a + b*3) - 4 = index
// index of fight map to use
var advantageArray = [9]int{
	0: TIE, // "1-1"
	4: TIE, // "2-2"
	8: TIE, // "3-3"
	6: P1,  // "1-3"
	1: P1,  // "2-1"
	5: P1,  // "3-2"
	3: P2,  // "1-2"
	7: P2,  // "2-3"
	2: P2}  // "3-1"

var tieFightArray = [9]int{
	0: TIE, // "1-1"
	3: P2,  // "1-2"
	6: P2,  // "1-3"
	1: P1,  // "2-1"
	4: TIE, // "2-2"
	7: P2,  // "2-3"
	2: P1,  // "3-1"
	5: P1,  // "3-2"
	8: TIE} // "3-3"

// used for P1
var p1FightArray = [9]int{
	0: P1,  // "1-1"
	3: TIE, // "1-2"
	6: P2,  // "1-3"
	1: P1,  // "2-1"
	4: P1,  // "2-2"
	7: TIE, // "2-3"
	2: P1,  // "3-1"
	5: P1,  // "3-2"
	8: P1}  // "3-3"

var p2FightArray = [9]int{
	0: P2,  // "1-1"
	1: TIE, // "2-1"
	2: P1,  // "3-1"
	3: P2,  // "1-2"
	4: P2,  // "2-2"
	5: TIE, // "3-2"
	6: P2,  // "1-3"
	7: P2,  // "2-3"
	8: P2}  // "3-3"

/**
 * p1 = type of Move for player 1
 * p2 = type of Move for player 2
 * return: index to arrays
 */
func getIndex(p1, p2 int) int {
	return (p1 + p2*3) - 4
}

/**
 * calculates winner, winner wil be nil if tie
 * return (winner, err)
 */
func getWinner(p1, p2 Move) (int, error) {
	// first see who has advantage and fight index
	advantageIndex := getIndex(p1.Type, p2.Type)
	fightIndex := getIndex(p1.Type, p2.Type)

	switch advantage := advantageArray[advantageIndex]; advantage {
	case TIE:
		return tieFightArray[fightIndex], nil
	case P1:
		return p1FightArray[fightIndex], nil
	case P2:
		return p2FightArray[fightIndex], nil
	default:
		return INVALID, errors.New("invalid advantage. arguments might be invalid")
	}
}
