package main

import (
	"example.com/m/v2/item"
	"fmt"
)

func main() {
	qtd := 75
	itemAnalyzed := item.MachadaoDeAco
	report := itemAnalyzed.CalcItemsReport(qtd)

	fmt.Printf("%s, qtd(%d), exp_total(%d)\n", itemAnalyzed.Name, qtd, qtd*itemAnalyzed.Exp)

	fmt.Println("\nPara Coletar")
	for _, value := range report {
		if value.IsCollectable() {
			fmt.Printf("%s: %d\n", value.Name, value.Quantity)
		}
	}

	fmt.Println("\nPara Craftar")
	for _, value := range report {
		if !value.IsCollectable() {
			fmt.Printf("%s: %d\n", value.Name, value.Quantity)
		}
	}
}
