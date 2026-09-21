package main

import (
	"fmt"
	"log"

	energy "github.com/cool-username-1337/lab4-variant5/pkg/powerbill"
)

func main() {
	var (
		owner   = "Иван Иванов"
		prevKwh = 1200.0
		currKwh = 1350.5
		tariff  = 5.50
		percent = 10.0
	)

	kwh, err := energy.Consumption(prevKwh, currKwh)
	if err != nil {
		log.Fatal(err)
	}

	cost, err := energy.EnergyCost(kwh, tariff)
	if err != nil {
		log.Fatal(err)
	}

	if err := energy.ApplyDiscount(&cost, percent); err != nil {
		log.Fatal(err)
	}

	report, err := energy.FormatEnergyReport(owner, kwh, cost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(report)
}
