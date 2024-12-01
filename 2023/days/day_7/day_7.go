package day7

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

/*
Five of a kind, where all five cards have the same label: AAAAA
Four of a kind, where four cards have the same label and one card has a different label: AA8AA
Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
High card, where all cards' labels are distinct: 23456
*/

var cardToValue = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardToValueP2 = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 0, // joker is the weakest card
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type hand struct {
	cards []int
	bid   int
}

func ResolvePartOne(r io.Reader) int {
	return calculateTotalBid(parseHands(r, cardToValue), false)
}

func ResolvePartTwo(r io.Reader) int {
	return calculateTotalBid(parseHands(r, cardToValueP2), true)
}

func calculateTotalBid(hands []hand, withJoker bool) int {
	combinations := make(map[int][]hand, 5)
	matches := make(map[int]int, 5)
	for _, h := range hands {
		clear(matches)
		for _, c := range h.cards {
			matches[c] += 1
		}

		if withJoker {
			jokerMatches, ok := matches[0]
			if ok {
				// if there is not only jokers - then joker becomes different card
				if len(matches) != 1 {
					delete(matches, 0)
					for k, v := range matches {
						matches[k] = v + jokerMatches
					}
				}
			}
		}

		maxMatch := matches[0]
		for _, match := range matches {
			if match > maxMatch {
				maxMatch = match
			}
		}

		switch maxMatch {
		case 5:
			maxMatch += 2
			combinations[maxMatch] = append(combinations[maxMatch], h)
		case 4:
			maxMatch += 2
			combinations[maxMatch] = append(combinations[maxMatch], h)
		case 3:
			maxMatch += 1
			if len(matches) == 2 {
				maxMatch += 1
				combinations[maxMatch] = append(combinations[maxMatch], h)
				break
			}

			combinations[maxMatch] = append(combinations[maxMatch], h)
		case 2:
			if len(matches) == 3 {
				maxMatch += 1
				combinations[maxMatch] = append(combinations[maxMatch], h)
				break
			}

			combinations[maxMatch] = append(combinations[maxMatch], h)
		case 1:
			combinations[maxMatch] = append(combinations[maxMatch], h)
		}
	}

	totalCCnt := 0
	for cCnt, combination := range combinations {
		if cCnt > totalCCnt {
			totalCCnt = cCnt
		}

		combinations[cCnt] = sortHands(combination)
	}

	rang := 1
	result := 0
	for i := 1; i <= totalCCnt; i++ {
		hs, ok := combinations[i]
		if !ok {
			continue
		}

		for _, h := range hs {
			result += h.bid * rang
			rang += 1
		}

	}

	return result
}

func sortHands(h []hand) []hand {
	sort.Slice(h, func(i, j int) bool {
		hi := h[i].cards
		hj := h[j].cards
		for idx, c := range hi {
			if c < hj[idx] {
				return true
			}

			if c > hj[idx] {
				return false
			}
		}

		return false
	})

	return h
}

func parseHands(r io.Reader, mapper map[rune]int) []hand {
	s := bufio.NewScanner(r)
	var hands []hand

	for s.Scan() {
		st := strings.Split(s.Text(), " ")
		cards := make([]int, 5)
		for idx, c := range st[0] {
			cards[idx] = mapper[c]
		}

		bid, _ := strconv.Atoi(st[1])
		hands = append(hands, hand{
			cards: cards,
			bid:   bid,
		})
	}

	return hands
}
