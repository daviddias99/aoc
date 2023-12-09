package day7

import (
	"aoc/2023/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)


var card_value_joker = map[string]int {
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

func classify_play_joker(play []rune) Type{

	card_count := make([]int, 14)
	for i := 0; i < 5; i++ {
		card_count[card_value_joker[string(play[i])]] += 1
	}

	card_count_sorted := make([]int, 14)
	copy(card_count_sorted, card_count)

	slices.Sort(card_count_sorted)

	if card_count_sorted[13] == 5 {
		return FIVE_OF_A_KIND
	}

	if card_count_sorted[13] == 4 {

		if card_count[0] == 1 {
			return FIVE_OF_A_KIND
		}

		if card_count[0] == 4 {
			return FIVE_OF_A_KIND
		}

		return FOUR_OF_A_KIND
	}

	if card_count_sorted[13] == 3 && card_count_sorted[12] == 2 {

		if card_count[0] == 2 {
			return FIVE_OF_A_KIND
		}

		if card_count[0] == 3 {
			return FIVE_OF_A_KIND
		}

		return FULL_HOUSE
	}

	if card_count_sorted[13] == 3 {

		if card_count[0] == 1 {
			return FOUR_OF_A_KIND
		}

		if card_count[0] == 3 {
			return FOUR_OF_A_KIND
		}

		return THREE_OF_A_KIND
	}

	if card_count_sorted[13] == 2 && card_count_sorted[12] == 2 {

		if card_count[0] == 1 {
			return FULL_HOUSE
		}

		if card_count[0] == 2 {
			return FOUR_OF_A_KIND
		}

		return TWO_PAIR
	}

	if card_count_sorted[13] == 2{

		if card_count[0] == 1 {
			return THREE_OF_A_KIND
		}

		if card_count[0] == 2 {
			return THREE_OF_A_KIND
		}

		return ONE_PAIR
	}

	if card_count[0] == 1 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func is_value_less_joker(a string, b string) bool {
	cards_a := []rune(a)
	cards_b := []rune(b)

	type_a := classify_play_joker(cards_a)
	type_b := classify_play_joker(cards_b)

	if type_a  == type_b {
		for i := 0; i < len(cards_a); i++ {
			if cards_a[i] == cards_b[i] {
				continue
			}

			return card_value_joker[string(cards_a[i])] < card_value_joker[string(cards_b[i])]
		}
	}

	return type_a < type_b
}


func SolveB() {

	fmt.Println("Day 7 Problem B")

	lines := utils.ReadLines("day7/input/b.input")

	sort.Slice(lines, func(i, j int) bool {
		return is_value_less_joker(lines[i], lines[j])
	})

	total := 0

	for i, v := range lines {
		hands := strings.Split(v, " ")
		bid, _ := strconv.Atoi(hands[1])

		total += (i+1) * bid
	}

	fmt.Println(total)
}
