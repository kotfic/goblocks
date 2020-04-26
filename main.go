package main

import (
	"fmt"
	"math"
	"github.com/distatus/battery"
)

type BatteryInfo struct {
	State   battery.State
	Percent float64
}

func calc_bat_info(b *battery.Battery) (s string) {

	// bi.Percent = (b.Current / b.Full) * 100.00
	symbol := "⚡"
	if b.State == battery.Discharging {
		symbol = "🔋"
	}
	s = fmt.Sprintf("%s %.0f%%", symbol, math.Round((b.Current / b.Full) * 100.00))
	return s
}

func main() {
	battery, err := battery.Get(0)
	if err != nil {
		fmt.Println("Could not get battery info!")
		return
	}
	fmt.Print(calc_bat_info(battery))

}
