package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("OMGHAI!")

	d := flag.Int("days", 2, "set the number of days")
	flag.Parse()

	days := *d + 1 // Add 1 as we loop this amount of time.

	items := NewGildedRoseItems()

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("name, sellIn, quality")
		for _, item := range items {
			fmt.Printf("%s\n", item.String())
			item.Update()
		}
	}
}
