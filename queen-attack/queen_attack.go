package queenattack

import (
	"strconv"
	"errors"
)

var posIntMap = map[byte]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
}

type position struct {
	x, y int
}

func strPositionToInteger(p string) (*position, error) {

	if len(p) != 2 {
		return nil, errors.New("Invalid argument")
	}

	pos := &position{}
	p1, ok := posIntMap[p[0]]
	if !ok {
		return nil, errors.New("Invalid argument")
	}
	pos.x = p1

	p2, err := strconv.Atoi(p[1:])
	if err != nil {
		return nil, err
	}
	if p2 > 8 || p2 < 1 {
		return nil, errors.New("Invalid position")
	}
	pos.y = p2
	return pos, nil
}

func CanQueenAttack(w, b string) (bool, error) {

	posW, errW := strPositionToInteger(w)
	if errW != nil {
		return false, errW
	}

	posB, errB := strPositionToInteger(b)
	if errB != nil {
		return false, errB
	}

	if posW.x == posB.x && posW.y == posB.y {
		return false, errors.New("Two queen can't be at same position")
	}

	if posW.x == posB.x || posW.y == posB.y {
		return true, nil
	}

	if posW.x-posB.x == posW.y-posB.y || posW.x-posB.x == posB.y-posW.y {
		return true, nil
	}

	return false, nil
}
