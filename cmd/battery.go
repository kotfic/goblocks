/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"encoding/json"
	"github.com/distatus/battery"
	"math"
	"os"

	"github.com/spf13/cobra"
)

// TODO enum structs for Markup and Align fields
// TODO Move this out of the battery.go and into a more central location where other subcommands can use it
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

// batteryCmd represents the battery command
var batteryCmd = &cobra.Command{
	Use:   "battery",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// TODO This function is going to be mostly the same across all
	// subcommands, is there a way to move it (along with I3BarInput out to
	// a better location?)
	Run: func(cmd *cobra.Command, args []string) {

		b, err := json.Marshal(BatteryInfo())
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)
	},
}


// TODO What is this doing?
func init() {
	rootCmd.AddCommand(batteryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// batteryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// batteryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
