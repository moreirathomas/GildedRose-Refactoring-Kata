package main

import "testing"

func TestCommonItem(t *testing.T) {
	item := NewCommonItem("Elixir of the Mongoose", 1, 5)

	item.Update()
	if item.sellIn != 0 {
		t.Errorf("SellIn: Expected %d but got %d", 0, item.sellIn)
	}
	if item.quality != 4 {
		t.Errorf("Quality: Expected %d but got %d", 4, item.quality)
	}

	// Once the sell by date has passed, Quality degrades twice as fast.
	item.Update()
	if item.sellIn != -1 {
		t.Errorf("SellIn: Expected %d but got %d", -1, item.sellIn)
	}
	if item.quality != 2 {
		t.Errorf("Quality: Expected %d but got %d", 2, item.quality)
	}
}

func TestRareItem(t *testing.T) {
	multiplierFunc := func(i *Item) int {
		m := 1
		if i.sellIn < 0 {
			m = m * 2
		}
		return m
	}

	item := NewRareItem("Aged Brie", 1, 0, multiplierFunc)

	item.Update()
	if item.sellIn != 0 {
		t.Errorf("SellIn: Expected %d but got %d", 0, item.sellIn)
	}
	if item.quality != 1 {
		t.Errorf("Quality: Expected %d but got %d", 1, item.quality)
	}

	// Once the sell by date has passed, Quality upgrades twice as fast.
	item.Update()
	if item.sellIn != -1 {
		t.Errorf("SellIn: Expected %d but got %d", -1, item.sellIn)
	}
	if item.quality != 3 {
		t.Errorf("Quality: Expected %d but got %d", 3, item.quality)
	}

	multiplierFunc = func(i *Item) int {
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
	item = NewRareItem("Backstage passes to a TAFKAL80ETC concert", 15, 20, multiplierFunc)

	// "Backstage passes" increases in Quality as its SellIn value approaches.
	item.Update()
	if item.sellIn != 14 {
		t.Errorf("SellIn: Expected %d but got %d", 14, item.sellIn)
	}
	if item.quality != 21 {
		t.Errorf("Quality: Expected %d but got %d", 21, item.quality)
	}

	// Quality increases by 2 when there are 10 days or less.
	item = NewRareItem("Backstage passes to a TAFKAL80ETC concert", 10, 20, multiplierFunc)
	item.Update()
	if item.sellIn != 9 {
		t.Errorf("SellIn: Expected %d but got %d", 9, item.sellIn)
	}
	if item.quality != 22 {
		t.Errorf("Quality: Expected %d but got %d", 22, item.quality)
	}

	// Quality increases by 3 when there are 5 days or less
	item = NewRareItem("Backstage passes to a TAFKAL80ETC concert", 5, 20, multiplierFunc)
	item.Update()
	if item.sellIn != 4 {
		t.Errorf("SellIn: Expected %d but got %d", 4, item.sellIn)
	}
	if item.quality != 25 {
		t.Errorf("Quality: Expected %d but got %d", 25, item.quality)
	}

	// Quality drops to 0 after the concert
	item = NewRareItem("Backstage passes to a TAFKAL80ETC concert", 0, 50, multiplierFunc)
	item.Update()
	if item.sellIn != -1 {
		t.Errorf("SellIn: Expected %d but got %d", -1, item.sellIn)
	}
	if item.quality != 0 {
		t.Errorf("Quality: Expected %d but got %d", 0, item.quality)
	}
}

func TestLegendaryItem(t *testing.T) {
	item := NewLegendaryItem("Sulfuras, Hand of Ragnaros", 0, 80)

	item.Update()
	if item.sellIn != 0 {
		t.Errorf("SellIn: Expected %d but got %d", 0, item.sellIn)
	}
	if item.quality != 80 {
		t.Errorf("Quality: Expected %d but got %d", 80, item.quality)
	}

	// Negative sellIn are allowed.
	item = NewLegendaryItem("Sulfuras, Hand of Ragnaros", -1, 80)

	item.Update()
	if item.sellIn != -1 {
		t.Errorf("SellIn: Expected %d but got %d", -1, item.sellIn)
	}
	if item.quality != 80 {
		t.Errorf("Quality: Expected %d but got %d", 80, item.quality)
	}
}

func TestConjuredItem(t *testing.T) {
	item := NewConjuredItem("Conjured Mana Cake", 1, 6)

	item.Update()
	if item.sellIn != 0 {
		t.Errorf("SellIn: Expected %d but got %d", 0, item.sellIn)
	}
	if item.quality != 4 {
		t.Errorf("Quality: Expected %d but got %d", 4, item.quality)
	}

	// Once the sell by date has passed, Quality degrades twice as fast.
	item.Update()
	if item.sellIn != -1 {
		t.Errorf("SellIn: Expected %d but got %d", -1, item.sellIn)
	}
	if item.quality != 0 {
		t.Errorf("Quality: Expected %d but got %d", 0, item.quality)
	}
}

func TestQualityLimit(t *testing.T) {
	// The Quality of an item is never negative.
	degradingItem := NewCommonItem("Elixir of the Mongoose", 5, 0)
	degradingItem.Update()
	if degradingItem.quality != 0 {
		t.Errorf("Quality: Expected %d but got %d", 0, degradingItem.quality)
	}

	// The Quality of an item is never more than 50.
	upgradingItem := NewRareItem("Elixir of the Mongoose", 5, 50, func(i *Item) int { return 1 })
	upgradingItem.Update()
	if upgradingItem.quality != 50 {
		t.Errorf("Quality: Expected %d but got %d", 50, upgradingItem.quality)
	}

	// The Quality of a legendary item is 80 and it never alters.
	// This is tested via TestLegendaryItem.
}
