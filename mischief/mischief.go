package mischief

var Counter uint8 = 0

var Ratio uint8 = 16

func TrickOrTreat() bool {

	Counter++
	if Counter/Ratio == 0 {
		return true
	}
	return false
}
