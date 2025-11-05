package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Scheme  string                       `json:"scheme"`
	Schemes map[string]map[string]bool   `json:"schemes"`
	Order   string                       `json:"order"`
	Orders  map[string][]string          `json:"orders"`
	Color   string                       `json:"colorScheme"`
	Colors  map[string]map[string]string `json:"colorSchemes"`
	Version string                       `json:"version"`
	Arts    map[string][]string          `json:"arts"`
	Art     string                       `json:"art"`
}

func loadConfig(scheme string) {
	path := getConfigPath()
	var cfg Config
	if path == "" {
		if err := json.Unmarshal([]byte(scheme), &cfg); err != nil {
			fmt.Println("Error loading config, check your JSON syntax\nerror:", err)
		}
	} else {
		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Error opening config file") //or using default config
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&cfg); err != nil {
			fmt.Println("Error loading config, check your JSON syntax\nerror:", err)
		}
	}

	selectedScheme := cfg.Schemes[cfg.Scheme]

	actions := map[string]func(string, string){
		"art":      func(color string, reset string) { displayArt() },
		"title":    func(color string, reset string) { showTitle(color, reset) },
		"os":       func(color string, reset string) { displayOs(color, reset) },
		"hostname": func(color string, reset string) { displayHostname(color, reset) },
		"kernel":   func(color string, reset string) { displayKernel(color, reset) },
		"uptime":   func(color string, reset string) { displayUptime(color, reset) },
		"bootime":  func(color string, reset string) { displayBootTime(color, reset) },
		"procs":    func(color string, reset string) { displayProcs(color, reset) },
		"cpu":      func(color string, reset string) { displayCpu(color, reset) },
		"gpu":      func(color string, reset string) { displayGpu(color, reset) },
		"memory":   func(color string, reset string) { displayMemory(color, reset) },
		"swap":     func(color string, reset string) { displaySwap(color, reset) },
		"disk":     func(color string, reset string) { displayDisk(color, reset) },
		"ip":       func(color string, reset string) { displayNetwork(color, reset) },
		"colors":   printAnsiColors,
		"credits":  func(color string, reset string) { credits(color, reset) },
		"locale":   func(color string, reset string) { displayLocale(color, reset) },
		"battery":  func(color string, reset string) { displayBattery(color, reset) },
		"host":     func(color string, reset string) { displayHost(color, reset) },
	}

	colors := map[string]string{
		"Reset":        "\033[0m",
		"Red":          "\033[31m",               // #FF0000
		"Green":        "\033[32m",               // #00FF00
		"Yellow":       "\033[33m",               // #FFFF00
		"Blue":         "\033[34m",               // #0000FF
		"Magenta":      "\033[35m",               // #FF00FF
		"Cyan":         "\033[36m",               // #00FFFF
		"Gray":         "\033[37m",               // #BFBFBF
		"White":        "\033[97m",               // #FFFFFF
		"Orange":       "\033[38;5;208m",         // #FF8700 (approx)
		"Purple":       "\033[38;5;129m",         // #875FD7 (approx)
		"Pink":         "\033[38;5;200m",         // #FF87D7 (approx)
		"Brown":        "\033[38;5;94m",          // #5F3F00 (approx)
		"Black":        "\033[30m",               // #000000
		"LightGray":    "\033[38;5;245m",         // #A8A8A8 (approx)
		"DarkGray":     "\033[38;5;236m",         // #585858 (approx)
		"LightRed":     "\033[38;5;196m",         // #FF0000 (bright)
		"LightGreen":   "\033[38;5;46m",          // #00FF00 (bright)
		"LightYellow":  "\033[38;5;226m",         // #FFFF00 (bright)
		"LightBlue":    "\033[38;5;33m",          // #5F87FF (approx)
		"LightMagenta": "\033[38;5;201m",         // #FF87D7 (approx)
		"LightCyan":    "\033[38;5;51m",          // #00AFD7 (approx)
		"LightOrange":  "\033[38;5;214m",         // #FFAF5F (approx)
		"LightPurple":  "\033[38;5;171m",         // #D787FF (approx)
		"LightPink":    "\033[38;5;213m",         // #FFAFD7 (approx)
		"LightBrown":   "\033[38;5;130m",         // #AF8700 (approx)
		"BrightPink":   "\033[38;2;241;49;105m",  // #F13169
		"Rose":         "\033[38;2;242;121;182m", // #F279B6
		"Coral":        "\033[38;2;255;111;97m",  // #FF6F61
		"Teal":         "\033[38;2;0;150;136m",   // #009688
		"Mint":         "\033[38;2;152;255;159m", // #98FF9F
		"Lime":         "\033[38;2;181;255;0m",   // #B5FF00
		"Sunset":       "\033[38;2;255;165;84m",  // #FFA554
		"Lavender":     "\033[38;2;180;167;255m", // #B4A7FF
		"Slate":        "\033[38;2;83;100;120m",  // #536478
		"Mustard":      "\033[38;2;211;176;76m",  // #D3B04C
		"DeepBlue":     "\033[38;2;37;64;143m",   // #25408F
		"Amber":        "\033[38;2;255;191;0m",   // #FFBF00
		"Crimson":      "\033[38;2;220;20;60m",   // #DC143C
		"Olive":        "\033[38;2;128;128;0m",   // #808000
		"Turquoise":    "\033[38;2;64;224;208m",  // #40E0D0
		"HotPink":      "\033[38;2;255;105;180m", // #FF69B4
	}

	selectedOrder := cfg.Orders[cfg.Order]

	selectedColors := cfg.Colors[cfg.Color]

	/*
		selectedArt := cfg.Arts[cfg.Art]

		version := cfg.Version



		for i, line := range selectedArt {
			fmt.Println(line, )

	*/

	//var stats []string

	for _, key := range selectedOrder {
		if value, ok := selectedScheme[key]; ok && value {
			if fn, ok := actions[key]; ok {
				colorname := selectedColors[key] //og
				colorCode := colors[colorname]
				reset := colors["Reset"]
				fn(colorCode, reset)
			}
		}
	}

}
