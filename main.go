package main

import (
	"encoding/json"
	"fmt"
	"github.com/distatus/battery"
	"math"
	"os"
)



// TODO enum structs for Markup and Align fields
type I3BarInput struct {
	FullText            string `json:"full_text,omitempty"`
	ShortText           string `json:"short_text,omitempty"`
	Color               string `json:"color,omitempty"`
	Background          string `json:"background,omitempty"`
	Border              string `json:"border,omitempty"`
	BorderTop           int    `json:"border_top,omitempty"`
	BorderRight         int    `json:"border_right,omitempty"`
	BorderBottom        int    `json:"border_bottom,omitempty"`
	BorderLeft          int    `json:"border_left,omitempty"`
	MinWidth            int    `json:"min_width,omitempty"`
	Align               string `json:"align,omitempty"`
	Urgent              bool   `json:"urgent,omitempty"`
	Name                string `json:"name,omitempty"`
	Instance            string `json:"instance,omitempty"`
	Separator           bool   `json:"separator,omitempty"`
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"`
	Markup              string `json:"markup,omitempty"`
}

func BatteryInfo() (i3bi I3BarInput) {

	b, err := battery.Get(0)
	if err != nil {
		fmt.Println("Could not get battery info!")
		return
	}

	// bi.Percent = (b.Current / b.Full) * 100.00
	symbol := "⚡"
	if b.State == battery.Discharging {
		symbol = "🔋"
	}
	i3bi.FullText = fmt.Sprintf("%s %.0f%%", symbol, math.Round((b.Current/b.Full)*100.00))

	return i3bi
}

func main() {
	b, err := json.Marshal(BatteryInfo())
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}
