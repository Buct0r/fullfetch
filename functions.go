package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/distatus/battery"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func getConfigPath() string {
	// Controlla nella directory locale
	if _, err := os.Stat("config.json"); err == nil {
		return "config.json"
	}
	// Windows: controlla in %APPDATA%\fullfetch\config.json
	if runtime.GOOS == "windows" {
		appdata := os.Getenv("APPDATA")
		if appdata != "" {
			configPath := filepath.Join(appdata, "fullfetch", "config.json")
			if _, err := os.Stat(configPath); err == nil {
				return configPath
			}
		}
	}
	// Linux: controlla in ~/.config/fullfetch/config.json
	if runtime.GOOS == "linux" {
		usr, err := user.Current()
		if err == nil {
			configPath := filepath.Join(usr.HomeDir, ".config", "fullfetch", "config.json")
			if _, err := os.Stat(configPath); err == nil {
				return configPath
			}
		}
	}
	// Fallback
	return "config.json"
}

func showTitle(color string, reset string) {
	current, _ := user.Current()
	hostname, _ := os.Hostname()
	output := fmt.Sprintf("%s%s%s@%s%s%s", color, current.Username, reset, color, hostname, reset)
	fmt.Println(output, "\n"+strings.Repeat("-", len(output)))
}

func displayArt() {

	type Config struct {
		Art  string              `json:"art"`
		Arts map[string][]string `json:"arts"`
	}

	file, err := os.Open(getConfigPath())
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		panic(err)
	}

	selectedArt := cfg.Arts[cfg.Art]
	for _, line := range selectedArt {
		fmt.Println(line)
	}
	fmt.Print("\n")
}

func displayOs(color string, reset string) {
	info, _ := host.Info()
	fmt.Printf("%sOS:%s %v %v %v\n", color, reset, info.Platform, info.PlatformVersion, info.KernelArch)
}

func displayHostname(color string, reset string) {
	info, _ := host.Info()
	fmt.Printf("%vHostname:%s %v\n", color, reset, info.Hostname)
}

func displayHost(color string, reset string) {
	var model string
	switch runtime.GOOS {
	case "linux":
		data, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_name")
		if err == nil {
			model = strings.TrimSpace(string(data))
		}
	case "windows":
		out, err := exec.Command("wmic", "computersystem", "get", "model").Output()
		if err == nil {
			lines := strings.Split(string(out), "\n")
			if len(lines) > 1 {
				model = strings.TrimSpace(lines[1])
			}
		}
	case "darwin":
		out, err := exec.Command("sysctl", "-n", "hw.model").Output()
		if err == nil {
			model = strings.TrimSpace(string(out))
		}
	}
	if model == "" {
		model = "Unknown"
	}
	fmt.Printf("%vHost:%s %s\n", color, reset, model)
}

func displayProcs(color string, reset string) {
	info, _ := host.Info()
	fmt.Printf("%sProcesses:%s %v\n", color, reset, info.Procs)
}

func displayKernel(color string, reset string) {
	info, _ := host.Info()
	fmt.Printf("%sKernel:%s %v\n", color, reset, info.KernelVersion)
}

func displayUptime(color string, reset string) {
	info, _ := host.Info()
	uptime := info.Uptime
	if uptime >= 3600 {
		hours := float64(uptime / 3600)
		sub := uptime / 3600
		minutes := (hours - float64(sub)) * 60
		if minutes != 0 {
			fmt.Printf("%sUptime:%s %v hours %v minutes\n", color, reset, hours, minutes)
		} else {
			fmt.Printf("%sUptime:%s %v hours\n", color, reset, hours)
		}
	} else {
		fmt.Printf("%sUptime:%s %v mins\n", color, reset, uptime/60)
	}
}

func displayCpu(color string, reset string) {
	cpuInfo, _ := cpu.Info()
	model := cpuInfo[0].ModelName
	numCores, _ := cpu.Counts(false)
	GHZ := float64(cpuInfo[0].Mhz) / 1000
	fmt.Printf("%sCPU:%s (%v) %v @ %.2fGHz\n", color, reset, numCores, model, GHZ)
}

func displayGpu(color string, reset string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	case "linux":
		cmd = exec.Command("lspci")
	case "darwin":
		cmd = exec.Command("system_profiler", "SPDisplaysDataType")
	default:
		fmt.Println("OS not supported")
		return
	}
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error retrieving GPU", err)
		return
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), "nvidia") ||
			strings.Contains(strings.ToLower(line), "vga compatible controller") ||
			strings.Contains(strings.ToLower(line), "3D controller") ||
			strings.Contains(strings.ToLower(line), "radeon") ||
			strings.Contains(strings.ToLower(line), "arc") ||
			strings.Contains(strings.ToLower(line), "graphics") {
			fmt.Printf("%sGPU:%s %s\n", color, reset, strings.TrimSpace(line))
		}
	}
}

func displayMemory(color string, reset string) {
	out, _ := mem.VirtualMemory()
	fmt.Printf("%sMemory:%s %vMB / %vMB\n", color, reset, out.Used/1e6, out.Total/1e6)
}

func displayDisk(color string, reset string) {
	quantity, err := disk.Partitions(false)
	if err != nil {
		fmt.Println("Error retrieving disk partitions", err)
		return
	}
	for _, part := range quantity {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			fmt.Println("Error retrieving disk usage", err)
			return
		}
		fmt.Printf("%sDisk:%s %s - %vGB / %vGB\n", color, reset, part.Mountpoint, usage.Used/1e9, usage.Total/1e9)
	}
}

func displayNetwork(color string, reset string) {
	info, _ := host.Info()
	ip, err := net.LookupHost(info.Hostname)
	if err != nil {
		panic(err)
	}
	for _, ip := range ip {
		if strings.Count(ip, ".") == 3 {
			fmt.Printf("%sLocal IP:%s %s\n", color, reset, ip)
		}
	}
}

func credits(color string, reset string) {
	fmt.Printf("%sDeveloped by: Buct0r :3%s\n", color, reset)
}

func printAnsiColors(color string, reset string) {
	fmt.Print("\n\n")
	for i := 0; i < 16; i++ {
		fmt.Printf("\x1b[48;5;%dm  ", i)
	}
	fmt.Print("\x1b[0m\n") // Reset
}

func displayBattery(color string, reset string) {
	batteries, err := battery.GetAll()
	if err != nil {
		return
	}
	if len(batteries) == 0 {
		return
	}
	for _, bat := range batteries {
		percent := int(bat.Current / bat.Full * 100)
		fmt.Printf("%sBattery:%s %d%% %v\n", color, reset, percent, bat.State)
	}
}

func displayLocale(color string, reset string) {
	locale := os.Getenv("LANG")
	if locale != "" {
		fmt.Printf("%sLocale:%s %s\n", color, reset, locale)
	}
}
