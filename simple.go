package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func getAllPossibleNumbersWithoutZeroUntilXNotInY(x int, y []int) []int {
	var retVal []int

	for i := 1; i <= x; i++ {
		if notInSlice(i, y) {
			retVal = append(retVal, i)
		}
	}

	return retVal
}

func notInSlice(x int, y []int) bool {
	for _, i := range y {
		if i == x {
			return false
		}
	}

	return true
}

func getNotPickedRandom(previousPicks []int, upperBound int) []int {
	var unpickedNumbers = getAllPossibleNumbersWithoutZeroUntilXNotInY(upperBound, previousPicks)

	var indexOfNextPick = rand.Intn(len(unpickedNumbers))
	var nextPick = unpickedNumbers[indexOfNextPick]

	return append(previousPicks, nextPick)
}

func drawLotteryForUpperBound(howManyBalls int, upperBound int) []int {
	var picks []int
	for i := 0; i < howManyBalls; i++ {
		picks = getNotPickedRandom(picks, upperBound)
	}

	return picks
}

func returnSortedInts(ints []int) []int {
	var sortedInts = ints
	sort.Ints(sortedInts)
	return sortedInts
}

func drawLotteryForUpperBoundSorted(howManyBalls int, upperBound int) []int {
	return returnSortedInts(drawLotteryForUpperBound(howManyBalls, upperBound))
}

func untilMatch(guess []int, upperBound int) int {

	var sortedGuess = returnSortedInts(guess)
	var howManyBalls = len(guess)

	var roundsPlayed = 1
	for sortedDraw := drawLotteryForUpperBoundSorted(howManyBalls, upperBound); matchInts(sortedDraw, sortedGuess) == false; sortedDraw = drawLotteryForUpperBoundSorted(howManyBalls, upperBound) {
		roundsPlayed++
		/*
			if (roundsPlayed % 100) == 0 {
				fmt.Println(roundsPlayed)
			}
		*/
	}

	return roundsPlayed
}

func matchInts(ints1 []int, ints2 []int) bool {
	if len(ints1) != len(ints2) {
		fmt.Println("lens not matching!!!")
		return false
	}

	for i := range ints1 {
		if ints1[i] != ints2[i] {
			return false
		}
	}

	return true
}

func averageOverXWinCycles(x int, guess []int, upperBound int) int {
	var numOfTries []int

	for i := 0; i < x; i++ {
		numOfTries = append(numOfTries, untilMatch(guess, upperBound))
		fmt.Println("Try:", i, " of ", x)
	}

	return averageOverSlice(numOfTries)
}

func averageOverSlice(tries []int) int {

	var sumOfTries = 0
	for _, numOfTries := range tries {
		sumOfTries += numOfTries
	}

	return sumOfTries / len(tries)
}

func main() {

	var guess = []int{1, 2, 3, 4, 5, 6}
	var upperBound = 49

	fmt.Println("Rounds played until match:", averageOverXWinCycles(100, guess, upperBound))

}
