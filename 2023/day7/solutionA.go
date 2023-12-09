package day7

import (
	"aoc/2023/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Type int

const (
	HIGH_CARD Type = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

var card_value = map[string]int {
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

func classify_play(play []rune) Type{

	card_count := make([]int, 13)

	for i := 0; i < 5; i++ {
		card_count[card_value[string(play[i])]] += 1
	}

	slices.Sort(card_count)

	if card_count[12] == 5 {
		return FIVE_OF_A_KIND
	}

	if card_count[12] == 4 {
		return FOUR_OF_A_KIND
	}

	if card_count[12] == 3 && card_count[11] == 2 {
		return FULL_HOUSE
	}

	if card_count[12] == 3 {
		return THREE_OF_A_KIND
	}

	if card_count[12] == 2 && card_count[11] == 2 {
		return TWO_PAIR
	}

	if card_count[12] == 2{
		return ONE_PAIR
	}

	return HIGH_CARD
}

func is_value_less(a string, b string) bool {
	cards_a := []rune(a)
	cards_b := []rune(b)

	type_a := classify_play(cards_a)
	type_b := classify_play(cards_b)

	if type_a  == type_b {
		for i := 0; i < len(cards_a); i++ {
			if cards_a[i] == cards_b[i] {
				continue
			}

			return card_value[string(cards_a[i])] < card_value[string(cards_b[i])]
		}
	}

	return type_a < type_b
}


func SolveA() {

	fmt.Println("Day 7 Problem A")

	lines := utils.ReadLines("day7/input/a.input")

	sort.Slice(lines, func(i, j int) bool {
		return is_value_less(lines[i], lines[j])
	})

	total := 0

	for i, v := range lines {
		hands := strings.Split(v, " ")
		bid, _ := strconv.Atoi(hands[1])

		total += (i+1) * bid
	}

	fmt.Println(total)
}
