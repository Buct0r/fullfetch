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
		"Red":          "\033[31m",
		"Green":        "\033[32m",
		"Yellow":       "\033[33m",
		"Blue":         "\033[34m",
		"Magenta":      "\033[35m",
		"Cyan":         "\033[36m",
		"Gray":         "\033[37m",
		"White":        "\033[97m",
		"Orange":       "\033[38;5;208m",
		"Purple":       "\033[38;5;129m",
		"Pink":         "\033[38;5;200m",
		"Brown":        "\033[38;5;94m",
		"Black":        "\033[30m",
		"LightGray":    "\033[38;5;245m",
		"DarkGray":     "\033[38;5;236m",
		"LightRed":     "\033[38;5;196m",
		"LightGreen":   "\033[38;5;46m",
		"LightYellow":  "\033[38;5;226m",
		"LightBlue":    "\033[38;5;33m",
		"LightMagenta": "\033[38;5;201m",
		"LightCyan":    "\033[38;5;51m",
		"LightOrange":  "\033[38;5;214m",
		"LightPurple":  "\033[38;5;171m",
		"LightPink":    "\033[38;5;213m",
		"LightBrown":   "\033[38;5;130m",
		"LightBlack":   "\033[38;5;16m",
		"LightWhite":   "\033[38;5;255m",
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
