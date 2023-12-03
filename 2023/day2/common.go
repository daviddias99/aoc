package day2

import (
	"regexp"
)

var game_id_regex = regexp.MustCompile(`Game ([0-9]+):`)
var color_quantity_regex = regexp.MustCompile(`([0-9]+) ([a-z]+)`)
