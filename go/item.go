package main

import "fmt"

// Item represents an item in the Gilded Rose inn.
type Item struct {
	name    string
	sellIn  int // sellIn denotes the number of days available to sell the item.
	quality int // quality denotes how valuable the item is.
}

// CommonItem represents an item which has no special behavior.
type CommonItem struct {
	Item
	multiplier int
}

// RareItem represents an item which has a special behavior during its daily quality update.
// The multiplier is obtained through the MultiplierFunc, allowing for configurable behaviors.
type RareItem struct {
	Item
	MultiplierFunc MultiplierFunc
}

// LegendaryItem represents an item which never has to be sold or decreases in quality.
type LegendaryItem struct {
	Item
}

// ConjuredItem represents an item which degrade in quality twice as fast as common items.
type ConjuredItem struct {
	Item
	multiplier int
}

// MultiplierFunc represents a function to compute the update mutlitplier from an item.
type MultiplierFunc func(*Item) int

// Updater is an interface for an item which changes in quality at the end of each day.
type Updater interface {
	Update()        // String performs an update on the underlying data strcuture.
	String() string // String returns a human redeable string for the underlying data strcuture.
}

// newItem returns a new item.
func newItem(name string, sellIn, quality int) Item {
	return Item{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
}

// -- CommonItem

// NewCommonItem returns a new common item. Its quality multiplier is -1.
func NewCommonItem(name string, sellIn, quality int) *CommonItem {
	return &CommonItem{
		Item:       newItem(name, sellIn, quality),
		multiplier: -1,
	}
}

// Update drops the item sell by date and quality by 1.
// If the sell by date has passed, the quality drop is doubled.
func (i *CommonItem) Update() {
	i.sellIn = i.sellIn - 1

	if i.quality < 0 || i.quality >= 50 {
		return
	}

	m := i.multiplier
	if i.sellIn < 0 {
		m = m * 2
	}

	i.quality = i.quality + 1*m
}

// String returns a human readable representation of the item.
func (i CommonItem) String() string {
	return fmt.Sprintf("%s, %d, %d", i.name, i.sellIn, i.quality)
}

// -- RareItem

// NewRareItem returns a new rare item. Its quality multiplier is provided by multiplierFunc.
func NewRareItem(name string, sellIn, quality int, multiplierFunc MultiplierFunc) *RareItem {
	return &RareItem{
		Item:           newItem(name, sellIn, quality),
		MultiplierFunc: multiplierFunc,
	}
}

// Update drops the item sell by date and updates its quality.
// MultiplierFunc is used to compute the quality change.
func (i *RareItem) Update() {
	i.sellIn = i.sellIn - 1

	if i.quality < 0 || i.quality >= 50 {
		return
	}

	m := i.MultiplierFunc(&i.Item)

	i.quality = i.quality + 1*m
}

// String returns a human readable representation of the item.
func (i RareItem) String() string {
	return fmt.Sprintf("%s, %d, %d", i.name, i.sellIn, i.quality)
}

// -- LegendaryItem

// NewRareItem returns a new legendary item. It has no quality multiplier.
func NewLegendaryItem(name string, sellIn, quality int) *LegendaryItem {
	return &LegendaryItem{
		Item: newItem(name, sellIn, quality),
	}
}

// Update does nothing to the legendary item.
func (i *LegendaryItem) Update() {}

// String returns a human readable representation of the item.
func (i LegendaryItem) String() string {
	return fmt.Sprintf("%s, %d, %d", i.name, i.sellIn, i.quality)
}

// -- ConjuredItem

// NewConjuredItem returns a new conjured item. Its quality multiplier is -2.
func NewConjuredItem(name string, sellIn, quality int) *ConjuredItem {
	return &ConjuredItem{
		Item:       newItem(name, sellIn, quality),
		multiplier: -2,
	}
}

// Update drops the item sell by date and quality by 2.
// If the sell by date has passed, the quality drop is doubled.
func (i *ConjuredItem) Update() {
	i.sellIn = i.sellIn - 1

	if i.quality < 0 || i.quality >= 50 {
		return
	}

	m := i.multiplier
	if i.sellIn < 0 {
		m = m * 2
	}

	i.quality = i.quality + 1*m
}

// String returns a human readable representation of the item.
func (i ConjuredItem) String() string {
	return fmt.Sprintf("%s, %d, %d", i.name, i.sellIn, i.quality)
}
