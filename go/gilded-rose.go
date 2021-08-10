package main

type Item struct {
	name            string
	sellIn, quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		// Update quality basically.
		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].quality > 0 {
				// Legendary item is dealt with.
				if items[i].name != "Sulfuras, Hand of Ragnaros" {
					// Everything common.
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			// Only item with increasing quality.
			// Do nothing if already max Quality.
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				// Backstage exception: more increase depending on days left.
				if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		// End the day.
		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].sellIn = items[i].sellIn - 1
		}

		// Check if sellIn is in the past.
		if items[i].sellIn < 0 {
			if items[i].name != "Aged Brie" {
				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].quality > 0 {
						// Still do nothing with Sulfuras.
						if items[i].name != "Sulfuras, Hand of Ragnaros" {
							// Decrease a second time (twice as fast) if it is a normal item.
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					// Drop the Backstage passes Quality to zero.
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				// It is Aged Bried, always increasing Quality until max value.
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}
		}
	}

}
