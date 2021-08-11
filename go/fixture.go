package main

// NewGildedRoseItems returns a collection of new items available at the Gilded Rose inn.
func NewGildedRoseItems() []Updater {
	agedBrieMultiplierFunc := func(i *Item) int {
		m := 1
		if i.sellIn < 0 {
			m = m * 2
		}
		return m
	}

	backstagePassesMultiplierFunc := func(i *Item) int {
		switch {
		case i.sellIn < 0:
			return -i.quality
		case i.sellIn <= 5:
			return 5
		case i.sellIn <= 10:
			return 2
		default:
			return 1
		}
	}

	return []Updater{
		NewCommonItem("+5 Dexterity Vest", 10, 20),
		NewRareItem("Aged Brie", 2, 0, agedBrieMultiplierFunc),
		NewCommonItem("Elixir of the Mongoose", 5, 7),
		NewLegendaryItem("Sulfuras, Hand of Ragnaros", 0, 80),
		NewLegendaryItem("Sulfuras, Hand of Ragnaros", -1, 80),
		NewRareItem("Backstage passes to a TAFKAL80ETC concert", 15, 20, backstagePassesMultiplierFunc),
		NewRareItem("Backstage passes to a TAFKAL80ETC concert", 10, 49, backstagePassesMultiplierFunc),
		NewRareItem("Backstage passes to a TAFKAL80ETC concert", 5, 49, backstagePassesMultiplierFunc),
		NewConjuredItem("Conjured Mana Cake", 3, 6),
	}
}
