package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) format() string {
	fs := fmt.Sprintf("Bill breakdown: %v\n", b.name)
	total := 0.0

	for n, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", n+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total+b.tip)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
}
