package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	listOfLeagueMembers := []string{"Name1", "Name2", "Name3", "Name4", "Name5", "Name6", "Name7", "Name8", "Name9", "Name10", "Name11", "Name12"}

	randomlyOrderedList, err := generateRandomOrder(listOfLeagueMembers)
	if err != nil {
		fmt.Println("ERROR", err)

		return
	}

	fmt.Println(randomlyOrderedList)

	newList := generateNewListWithRandomOrder(listOfLeagueMembers, randomlyOrderedList)
	outputList(newList)
}

func generateRandomOrder(list []string) (map[int]int, error) {
	oldPosition := make(map[int]bool)
	newPosition := make(map[int]bool)
	newOrder := make(map[int]int)

	for range list {
		oldHaveRandomNumberThatHasNotBeenTaken := false
		var oldNumber int

		for !oldHaveRandomNumberThatHasNotBeenTaken {
			num, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
			if err != nil {
				return nil, err
			}
			numInt := int(num.Int64())

			_, ok := oldPosition[numInt]
			if !ok {
				oldHaveRandomNumberThatHasNotBeenTaken = true
				oldNumber = numInt
			}
		}

		newHaveRandomNumberThatHasNotBeenTaken := false
		var newNumber int

		for !newHaveRandomNumberThatHasNotBeenTaken {
			num, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
			if err != nil {
				return nil, err
			}
			numInt := int(num.Int64())

			_, ok := newPosition[numInt]
			if !ok {
				newHaveRandomNumberThatHasNotBeenTaken = true
				newNumber = numInt
			}
		}

		oldPosition[oldNumber] = true
		newPosition[newNumber] = true
		newOrder[oldNumber] = newNumber
	}

	return newOrder, nil
}

func generateNewListWithRandomOrder(list []string, order map[int]int) []string {
	newList := make([]string, len(list))
	for oldPosition, newPosition := range order {
		newList[newPosition] = list[oldPosition]
	}

	return newList
}

func outputList(list []string) {
	for i, v := range list {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}
