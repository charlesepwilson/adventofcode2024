package day05

import (
	"fmt"
	"testing"
)

func getExample() []byte {
	return []byte("47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47")
}

func TestPart1(t *testing.T) {
	result := part1(getExample())
	if result != 143 {
		t.Errorf("Wrong answer for day %d p1: %d", DAY, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(getExample())
	if result != 123 {
		t.Errorf("Wrong answer for day %d p2: %d", DAY, result)
	}
}

func TestFullOrder(t *testing.T) {
	fmt.Println(getFullOrder([]Ordering{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	},
	))
}
