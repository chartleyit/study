package main

import "fmt"

type Loot struct {
	Name       string
	Size       [2]int
	Value      int
	ValueDense int
}

func main() {
	loot := []Loot{
		{"Potion of Potionentiality", [2]int{1, 1}, 30, 0},
		{"Jeweled Dog Draught Excluder", [2]int{3, 1}, 150, 0},
		{"Spartan Shield", [2]int{2, 2}, 300, 0},
		{"Palindromic Sword o’Drows", [2]int{1, 3}, 120, 0},
		{"Unobsidian Armor", [2]int{2, 3}, 540, 0},
		{"Wardrobe of Infinite Lions", [2]int{20, 10}, 1337, 0},
	}

	// matrix to fill inventory
	inventory := [5][4]bool{}
	fmt.Println(inventory)

	for _, l := range loot {
		fmt.Println(l.Name, l.Value)
		// most value dense item
		vd := l.Value / (l.Size[0] * l.Size[1])
		fmt.Println("value density:", vd)

	}

}
