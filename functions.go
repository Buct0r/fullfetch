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
	"strconv"
	"strings"
	"time"

	"github.com/distatus/battery"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func getConfigPath() string {
	exePath, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exePath)
		configPath := filepath.Join(exeDir, "config.json")
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
	} // copilot cooked this two functions, I have to admit it
	if cwd, err := os.Getwd(); err == nil {
		configPath := filepath.Join(cwd, "config.json")
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
	}
	if runtime.GOOS == "windows" {
		appdata := os.Getenv("APPDATA")
		if appdata != "" {
			configPath := filepath.Join(appdata, "fullfetch", "config.json") //Gets config folder for different OS
			if _, err := os.Stat(configPath); err == nil {
				return configPath
			}
		}
	}
	if runtime.GOOS == "linux" {
		usr, err := user.Current()
		if err == nil {
			configPath := filepath.Join(usr.HomeDir, ".config", "fullfetch", "config.json")
			if _, err := os.Stat(configPath); err == nil {
				return configPath
			}
		}
	}
	// Not found anywhere
	return ""
}

func showTitle(color string, reset string) {
	current, _ := user.Current()
	hostname, _ := os.Hostname()
	output := fmt.Sprintf("%s%s%s@%s%s%s", color, current.Username, reset, color, hostname, reset)
	fmt.Println(output, "\n"+strings.Repeat("-", len(output)))
}

func displayArt() {

	type Config struct {
		Art  string              `json:"art"` //da modfificare
		Arts map[string][]string `json:"arts"`
	}

	scheme := `{

    "art" : "default",

    "arts": {
        "default": [
                "               @%%%%%%@@%                                                       ",
                "           %@@@@#**###%%@@@@%                                                   ",
                "         #@%#***#######%%%%%@@                                                  ",
                "         #@*#%%%%@@@@@@@@@@%%@                                                  ",
                "         #@#%%=-:-======+*@%%@  @@@@@@@@@@@@@@@@@@@@@@#  %@@@@                  ",
                "         #@#%%=:.:======+*@%%@  @-.....:@:.=@@..+#..#%   %=..%                  ",
                "         #@#%%@@@@@@@@@@@@@%%@  @-.-@@@@@:.=@@..+#..#%   %=..%                  ",
                "         #@#%%=:.:======+*@%%@  @-....-@@:.=@@..+#..#%   %=..%                  ",
                "         #@#%%=:.:======+*@%%@  @-.:###@@:.=@@..+#..#@@@@@=..%@@@@@             ",
                "         #@#%%=:.:======+*@%%@  @-.-@%%@@*:....=##.......%=.......@             ",
                "         #@#%@@@@@@@@@@@@@@%%@  @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%%%#        ",
                "         #@#%%=:.:======+*@%%@  @-.....=#.....:%......=@%....=%#..@*.-%         ",
                "         #@#%%=:.:======+*@%%@  @-.=@@@@*..@@@@@@@-.+@@%.:%@=.=#..@*.-%         ",
                "         #@#%%++=+******##@%%@  @-....*@*....:@@%@-.+@%%.:%@@@@#.....-%         ",
                "         #@#%%@@@@@@@@@@@@@%%@  @-.=%%%@*..%%%@%%@-.+@%%.:%@+:+#..%+.-%         ",
                "         #@#%%=::-======+*@%%@  @-.=@##%*.....:%%@-.+@%%+....-*#..@*.-%         ",
                "         #@#%%+=-=+++++++*@%%@  @@@@%**#@@@@@@@@%@@@@%##@@@@@@@@@@@@@@%         ",
                "         #@#%@%%@@@@@@@@@@@%%@                                                  ",
                "         #@@%#*#%%%%%%%%%%%%@@                                                  ",
                "           %@@@@@@@@@@@@@@@@%                                                   "
                                
        ],
		"defaultv2": [
                "               @%%%%%%@@%                  %s",
                "           %@@@@#**###%%@@@@%              %s",
                "         #@%#***#######%%%%%@@             %s",
                "         #@*#%%%%@@@@@@@@@@%%@             %s",
                "         #@#%%=-:-======+*@%%@             %s",
                "         #@#%%=:.:======+*@%%@             %s",
                "         #@#%%@@@@@@@@@@@@@%%@             %s",
                "         #@#%%=:.:======+*@%%@             %s",
                "         #@#%%=:.:======+*@%%@             %s",
                "         #@#%%=:.:======+*@%%@             %s",
                "         #@#%@@@@@@@@@@@@@@%%@             %s",
                "         #@#%%=:.:======+*@%%@             %s",
                "         #@#%%=:.:======+*@%%@             %s",
                "         #@#%%++=+******##@%%@             %s",
                "         #@#%%@@@@@@@@@@@@@%%@             %s",
                "         #@#%%=::-======+*@%%@             %s",
                "         #@#%%+=-=+++++++*@%%@             %s",
                "         #@#%@%%@@@@@@@@@@@%%@             %s",
                "         #@@%#*#%%%%%%%%%%%%@@             %s",
                "           %@@@@@@@@@@@@@@@@%              %s"
        ],
        "biglogo": [
        " FFFFFFFFFFFFFFFFFFFFFF                  lllllll lllllll    ffffffffffffffff                           tttt                             hhhhhhh",             
        " F::::::::::::::::::::F                  l:::::l l:::::l   f::::::::::::::::f                       ttt:::t                             h:::::h",             
        " F::::::::::::::::::::F                  l:::::l l:::::l  f::::::::::::::::::f                      t:::::t                             h:::::h ",            
        " FF::::::FFFFFFFFF::::F                  l:::::l l:::::l  f::::::fffffff:::::f                      t:::::t                             h:::::h ",            
        " F:::::F       FFFFFFuuuuuu    uuuuuu   l::::l  l::::l  f:::::f       ffffffeeeeeeeeeeee    ttttttt:::::ttttttt        cccccccccccccccch::::h hhhhh ",      
        " F:::::F             u::::u    u::::u   l::::l  l::::l  f:::::f           ee::::::::::::ee  t:::::::::::::::::t      cc:::::::::::::::ch::::hh:::::hhh" ,   
        " F::::::FFFFFFFFFF   u::::u    u::::u   l::::l  l::::l f:::::::ffffff    e::::::eeeee:::::eet:::::::::::::::::t     c:::::::::::::::::ch::::::::::::::hh " ,
        " F:::::::::::::::F   u::::u    u::::u   l::::l  l::::l f::::::::::::f   e::::::e     e:::::etttttt:::::::tttttt    c:::::::cccccc:::::ch:::::::hhh::::::h ",
        " F:::::::::::::::F   u::::u    u::::u   l::::l  l::::l f::::::::::::f   e:::::::eeeee::::::e      t:::::t          c::::::c     ccccccch::::::h   h::::::h",
        " F::::::FFFFFFFFFF   u::::u    u::::u   l::::l  l::::l f:::::::ffffff   e:::::::::::::::::e       t:::::t          c:::::c             h:::::h     h:::::h",
        " F:::::F             u::::u    u::::u   l::::l  l::::l  f:::::f         e::::::eeeeeeeeeee        t:::::t          c:::::c             h:::::h     h:::::h",
        " F:::::F             u:::::uuuu:::::u   l::::l  l::::l  f:::::f         e:::::::e                 t:::::t    ttttttc::::::c     ccccccch:::::h     h:::::h",
        " FF:::::::FF           u:::::::::::::::uul::::::ll::::::lf:::::::f        e::::::::e                t::::::tttt:::::tc:::::::cccccc:::::ch:::::h     h:::::h",
        " F::::::::FF            u:::::::::::::::ul::::::ll::::::lf:::::::f         e::::::::eeeeeeee        tt::::::::::::::t c:::::::::::::::::ch:::::h     h:::::h",
        " F::::::::FF             uu::::::::uu:::ul::::::ll::::::lf:::::::f          ee:::::::::::::e          tt:::::::::::tt  cc:::::::::::::::ch:::::h     h:::::h",
        " FFFFFFFFFFF               uuuuuuuu  uuuullllllllllllllllfffffffff            eeeeeeeeeeeeee            ttttttttttt      cccccccccccccccchhhhhhh     hhhhhhh"
        ],
        "smalllogo": [
            "FFFFFFF         lll lll  fff        tt           hh ",     
            "FF      uu   uu lll lll ff     eee  tt      cccc hh ",     
            "FFFF    uu   uu lll lll ffff ee   e tttt  cc     hhhhhh ", 
            "FF      uu   uu lll lll ff   eeeee  tt    cc     hh   hh ",
            "FF       uuuu u lll lll ff    eeeee  tttt  ccccc hh   hh"
        ]
                                       
    }

    
}`

	path := getConfigPath()
	var cfg Config
	if path == "" {
		if err := json.Unmarshal([]byte(scheme), &cfg); err != nil {
			fmt.Println("Error loading config, check your JSON syntax\nerror:", err)
		}

		selectedArt := cfg.Arts[cfg.Art]
		for _, line := range selectedArt {
			fmt.Println(line)
		}
		fmt.Print("\n")

	} else {

		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Error opening config file") //or using default config
		}

		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println("Error closing config file:", err)
			}
		}()

		if err := json.NewDecoder(file).Decode(&cfg); err != nil {
			fmt.Println("Error loading config, check your JSON syntax\nerror:", err)
		}

		selectedArt := cfg.Arts[cfg.Art]
		for _, line := range selectedArt {
			fmt.Println(line)
		}
		fmt.Print("\n")
	}
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
		hours := uptime / 3600
		minutes := (uptime % 3600) / 60
		days := hours / 24
		if days > 0 {
			if hours%24 == 0 {
				fmt.Printf("%sUptime:%s %d days %d mins\n", color, reset, days, minutes)
			} else {
				fmt.Printf("%sUptime:%s %d days %d hours %d mins\n", color, reset, days, hours%24, minutes)
			}
		} else {
			if minutes != 0 {
				fmt.Printf("%sUptime:%s %d hours %d mins\n", color, reset, hours, minutes)
			} else {
				fmt.Printf("%sUptime:%s %d hours\n", color, reset, hours)
			}
		}
	} else {
		fmt.Printf("%sUptime:%s %d mins\n", color, reset, uptime/60)
	}
}

func displayCpu(color string, reset string) {
	cpuInfo, _ := cpu.Info()
	model := strings.TrimSpace(cpuInfo[0].ModelName)
	numCores, _ := cpu.Counts(false)
	GHZ := float64(cpuInfo[0].Mhz) / 1000
	//usage, _ := cpu.Percent(time.Second, false)  <-- makes the program run too slowly
	output := fmt.Sprintf("%sCPU:%s (%v) %v @ %.2fGHz\n", color, reset, numCores, model, GHZ)
	fmt.Print(output) //usage[0]
}

func displayGpu(color string, reset string) {
	var cmd *exec.Cmd
	psCommand := "Get-CimInstance Win32_VideoController | Select-Object Name"
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell.exe", "-NoProfile", "-NonInteractive", "-Command", psCommand) //different methods to get GPU info on different OS
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
	used := out.Used / 1e6
	total := out.Total / 1e6
	percent := (used * 100) / total
	var color2 string
	if percent >= 50 && percent < 90 {
		color2 = "\033[33m"
	} else if percent >= 90 {
		color2 = "\033[31m"
	} else {
		color2 = "\033[32m"
	}
	fmt.Printf("%sMemory:%s %vMB / %vMB (%s%v%%%s)\n", color, reset, used, total, color2, int64(percent), reset)
}

func displaySwap(color string, reset string) {
	out, _ := mem.SwapMemory()
	used := out.Used / 1024 / 1024
	total := out.Total / 1024 / 1024
	percent := (used * 100) / total
	var color2 string
	if percent >= 50 && percent < 90 {
		color2 = "\033[33m"
	} else if percent >= 90 {
		color2 = "\033[31m"
	} else {
		color2 = "\033[32m"
	}
	fmt.Printf("%sSwap:%s %vMB / %vMB (%s%v%%%s)\n", color, reset, used, total, color2, int64(percent), reset)
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
		percent := (usage.Used * 100) / usage.Total
		var color2 string
		if percent >= 50 && percent < 90 {
			color2 = "\033[33m"
		} else if percent >= 90 {
			color2 = "\033[31m"
		} else {
			color2 = "\033[32m"
		}
		fmt.Printf("%sDisk:%s %s - %vGB / %vGB (%s%v%%%s)\n", color, reset, part.Mountpoint, usage.Used/1e9, usage.Total/1e9, color2, percent, reset)
	}
}

func displayNetwork(color string, reset string) {
	info, _ := host.Info()
	ip, err := net.LookupHost(info.Hostname)
	if err != nil {
		fmt.Println("Error retrieving IP\nerror:", err)
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
	for i := 0; i < 8; i++ {
		fmt.Printf("\x1b[48;5;%dm  ", i)
	}
	fmt.Print("\x1b[0m\n")
	for i := 8; i < 16; i++ {
		fmt.Printf("\x1b[48;5;%dm  ", i)
	}
	fmt.Print("\x1b[0m\n")
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
		var color2 string
		if percent <= 50 && percent > 20 {
			color2 = "\033[33m"
		} else if percent <= 20 {
			color2 = "\033[31m"
		} else {
			color2 = "\033[32m"
		}
		fmt.Printf("%sBattery:%s %s%d%%%s %v\n", color, reset, color2, percent, reset, bat.State)
	}
}

func displayLocale(color string, reset string) {
	locale := os.Getenv("LANG")
	if locale != "" {
		fmt.Printf("%sLocale:%s %s\n", color, reset, locale)
	}
}

func displayBootTime(color string, reset string) {
	info, _ := host.Info()
	time := time.Unix(int64(info.BootTime), 0)
	fmt.Printf("%sBoot time:%s %v\n", color, reset, time)
}

func displayPackages(color string, reset string) {
	var managers []string
	var totalCount int

	switch runtime.GOOS {
	case "linux":
		if _, err := exec.LookPath("dpkg"); err == nil {
			out, err := exec.Command("sh", "-c", "dpkg-query -W -f='${Status}' | grep -c 'install ok installed'").Output()
			if err == nil {
				count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
				if count > 0 {
					managers = append(managers, fmt.Sprintf("%d (dpkg)", count))
					totalCount += count
				}
			}
		}
		if _, err := exec.LookPath("rpm"); err == nil {
			out, err := exec.Command("sh", "-c", "rpm -qa | wc -l").Output()
			if err == nil {
				count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
				if count > 0 {
					managers = append(managers, fmt.Sprintf("%d (rpm)", count))
					totalCount += count
				}
			}
		}
		if _, err := exec.LookPath("pacman"); err == nil {
			out, err := exec.Command("sh", "-c", "pacman -Qq | wc -l").Output()
			if err == nil {
				count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
				if count > 0 {
					managers = append(managers, fmt.Sprintf("%d (pacman)", count))
					totalCount += count
				}
			}
		}
		if _, err := exec.LookPath("apk"); err == nil {
			out, err := exec.Command("sh", "-c", "apk info | wc -l").Output()
			if err == nil {
				count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
				if count > 0 {
					managers = append(managers, fmt.Sprintf("%d (apk)", count))
					totalCount += count
				}
			}
		}
		if _, err := exec.LookPath("flatpak"); err == nil {
			out, err := exec.Command("sh", "-c", "flatpak list --app | wc -l").Output()
			if err == nil {
				count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
				if count > 0 {
					managers = append(managers, fmt.Sprintf("%d (flatpak)", count))
					totalCount += count
				}
			}
		}
		if _, err := exec.LookPath("snap"); err == nil {
			out, err := exec.Command("sh", "-c", "snap list --all | wc -l").Output()
			if err == nil {
				count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
				if count > 1 {
					count--
					managers = append(managers, fmt.Sprintf("%d (snap)", count))
					totalCount += count
				}
			}
		}

	case "darwin":
		if _, err := exec.LookPath("brew"); err == nil {
			out, err := exec.Command("sh", "-c", "brew list --formula | wc -l").Output()
			if err == nil {
				count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
				if count > 0 {
					managers = append(managers, fmt.Sprintf("%d (brew)", count))
					totalCount += count
				}
			}
			out, err = exec.Command("sh", "-c", "brew list --cask | wc -l").Output()
			if err == nil {
				count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
				if count > 0 {
					managers = append(managers, fmt.Sprintf("%d (cask)", count))
					totalCount += count
				}
			}
		}
		if _, err := exec.LookPath("port"); err == nil {
			out, err := exec.Command("sh", "-c", "port installed | wc -l").Output()
			if err == nil {
				count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
				if count > 0 {
					managers = append(managers, fmt.Sprintf("%d (macports)", count))
					totalCount += count
				}
			}
		}

	case "windows":
		if _, err := exec.LookPath("winget"); err == nil {
			out, err := exec.Command("winget", "list").Output()
			if err == nil {
				lines := strings.Split(string(out), "\n")
				count := 0
				for _, line := range lines {
					trimmed := strings.TrimSpace(line)
					if trimmed != "" && !strings.HasPrefix(trimmed, "Name") && !strings.HasPrefix(trimmed, "-") && !strings.HasPrefix(trimmed, "Arp") {
						count++
					}
				}
				if count > 0 {
					managers = append(managers, fmt.Sprintf("%d (winget)", count))
					totalCount += count
				}
			}
		}
		if _, err := exec.LookPath("choco"); err == nil {
			out, err := exec.Command("choco", "list", "-l", "--limit-output").Output()
			if err == nil {
				lines := strings.Split(string(out), "\n")
				if len(lines) > 0 {
					count := len(lines)
					managers = append(managers, fmt.Sprintf("%d (choco)", count))
					totalCount += count
				}
			}
		}
	}

	if len(managers) == 0 {
		return
	}

	if len(managers) == 1 {
		fmt.Printf("%sPackages:%s %s\n", color, reset, managers[0])
	} else {
		fmt.Printf("%sPackages:%s %d (%s)\n", color, reset, totalCount, strings.Join(managers, ", "))
	}
}

func displayMB(color string, reset string) {

	switch runtime.GOOS {
	case "linux":
		data, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_vendor")
		if err != nil {
			fmt.Println("Error retrieving motherboard vendor\nerror:", err)
			return
		}
		data2, err2 := os.ReadFile("/sys/devices/virtual/dmi/id/board_name")
		if err2 != nil {
			fmt.Println("Error retrieving motherboard name\nerror:", err2)
			return
		}
		output := fmt.Sprintf("%sMotherboard:%s %s %s\n", color, reset, strings.TrimSpace(string(data)), strings.TrimSpace(string(data2)))
		fmt.Print(output)
	case "darwin":
		out, err := exec.Command("sysctl", "-n", "hw.model").Output()
		if err == nil {
			fmt.Printf("%sMotherboard:%s %s\n", color, reset, strings.TrimSpace(string(out)))
		}
	case "windows":
		manufacturer, err := exec.Command("powershell.exe", "-NoProfile", "-NonInteractive", "-Command", "Get-CimInstance Win32_BaseBoard | Select-Object -ExpandProperty  Manufacturer").Output()
		if err != nil {
			manufacturer = []byte("Unknown")
		}
		product, err := exec.Command("powershell.exe", "-NoProfile", "-NonInteractive", "-Command", "Get-CimInstance Win32_BaseBoard | Select-Object -ExpandProperty  Product").Output()
		if err != nil {
			product = []byte("Unknown")
		}
		output := fmt.Sprintf("%sMotherboard:%s %s %s\n", color, reset, strings.TrimSpace(string(manufacturer)), strings.TrimSpace(string(product)))

		fmt.Print(output)

	}

}

func displayRamModel(color string, reset string) {
	switch runtime.GOOS {
	case "windows":
		out, err := exec.Command("powershell.exe", "-NoProfile", "-NonInteractive", "-Command",
			"Get-CimInstance Win32_PhysicalMemory | ForEach-Object { \"$($_.Manufacturer) $($_.PartNumber) $($_.Speed)MHz $($_.Capacity / 1GB)GB\" }").Output()
		if err != nil {
			fmt.Println("Error retrieving RAM info")
			return
		}
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		for i, line := range lines {
			if strings.TrimSpace(line) != "" && i > 0 {
				fmt.Printf("%sRAM specs:%s %s\n", color, reset, strings.TrimSpace(line))
			}
		}

	case "linux":
		out, err := exec.Command("sh", "-c", "dmidecode -t memory 2>/dev/null | grep 'Part Number\\|Manufacturer\\|Speed'").Output()
		if err != nil {
			return
		}
		fmt.Printf("%sRAM specs:%s %s\n", color, reset, strings.TrimSpace(string(out)))

	case "darwin":
		out, err := exec.Command("system_profiler", "SPMemoryDataType").Output()
		if err != nil {
			return
		}
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Type:") || strings.Contains(line, "Size:") || strings.Contains(line, "Speed:") {
				fmt.Printf("%sRAM specs:%s %s\n", color, reset, strings.TrimSpace(line))
			}
		}
	}
}
