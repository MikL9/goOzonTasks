package main

import (
	"fmt"
)

type Player struct {
	PlayerId int
	Cards    []string
}

func getCards() []string {
	cards := make([]string, 15)
	cards[2] = "2"
	cards[3] = "3"
	cards[4] = "4"
	cards[5] = "5"
	cards[6] = "6"
	cards[7] = "7"
	cards[8] = "8"
	cards[9] = "9"
	cards[10] = "T"
	cards[11] = "J"
	cards[12] = "Q"
	cards[13] = "K"
	cards[14] = "A"

	return cards
}

func getTypes() []string {
	cardTypes := make([]string, 4)
	cardTypes[0] = "S"
	cardTypes[1] = "C"
	cardTypes[2] = "D"
	cardTypes[3] = "H"
	return cardTypes
}

func createKoloda() []string {
	var koloda []string
	cards := getCards()
	cardTypes := getTypes()

	for _, card := range cards {
		for _, cardType := range cardTypes {
			if card != "" {
				karta := card + cardType
				koloda = append(koloda, karta)
			}
		}
	}
	return koloda
}

func sliceCards(koloda *[]string, cards string) {
	for i, c := range *koloda {
		if cards == c {
			*koloda = append((*koloda)[:i], (*koloda)[i+1:]...)
			return
		}
	}
}

func main() {
	swappedMap := make(map[string]int)
	for i, card := range getCards() {
		if card != "" {
			swappedMap[card] = i
		}
	}

	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		koloda := createKoloda()
		var n int
		fmt.Scan(&n)
		var solved []string

		//GET ALL PLAYERS
		var allPlayersData []Player
		for j := 1; j <= n; j++ {
			var playerData Player
			var oppCards []string
			for c := 1; c <= 2; c++ {
				var oppCard string
				fmt.Scan(&oppCard)
				oppCards = append(oppCards, oppCard)
				sliceCards(&koloda, oppCard)
			}
			playerData.Cards = oppCards
			playerData.PlayerId = j
			allPlayersData = append(allPlayersData, playerData)
		}
		//GET WINNING CARDS
		for _, card := range koloda {
			if getWinning(swappedMap, card, allPlayersData) {
				solved = append(solved, card)
			}
		}
		fmt.Println(len(solved))
		for _, card := range solved {
			fmt.Println(card)
		}
	}
}

func getWinning(cardsMap map[string]int, card string, data []Player) bool {
	bestSet := 0
	elderCard := 0
	bestPlayer := 0
	for _, player := range data {
		player.Cards = append(player.Cards, card)
		cards := player.Cards

		//SET3
		if string(cards[0][0]) == string(cards[1][0]) && string(cards[0][0]) == string(cards[2][0]) {
			e, _ := cardsMap[string(cards[0][0])]
			if bestSet != 3 {
				elderCard = e
				bestPlayer = player.PlayerId
			} else if e > elderCard {
				elderCard = e
				bestPlayer = player.PlayerId
			}
			bestSet = 3
		}

		//SET2
		if bestSet != 3 {
			var e int
			setSetTwo := false
			if string(cards[0][0]) == string(cards[1][0]) || string(cards[0][0]) == string(cards[2][0]) {
				e, _ = cardsMap[string(cards[0][0])]
				setSetTwo = true
			} else if string(cards[1][0]) == string(cards[2][0]) {
				e, _ = cardsMap[string(cards[1][0])]
				setSetTwo = true
			}
			if setSetTwo {
				if bestSet != 2 {
					elderCard = e
					bestPlayer = player.PlayerId
				} else if e > elderCard {
					elderCard = e
					bestPlayer = player.PlayerId
				}
				bestSet = 2
			}
		}

		//SET1
		if bestSet != 3 && bestSet != 2 {
			if string(cards[0][0]) != string(cards[1][0]) ||
				string(cards[0][0]) != string(cards[2][0]) ||
				string(cards[1][0]) != string(cards[2][0]) {
				e1, _ := cardsMap[string(cards[0][0])]
				e2, _ := cardsMap[string(cards[1][0])]
				e3, _ := cardsMap[string(cards[2][0])]
				var e int
				if e1 >= e2 && e1 >= e3 {
					e = e1
				} else if e2 >= e1 && e2 >= e3 {
					e = e2
				} else {
					e = e3
				}
				if e > elderCard {
					elderCard = e
					bestPlayer = player.PlayerId
				}
				bestSet = 1
			}
		}
	}
	if bestPlayer == 1 {
		return true
	}
	return false
}
