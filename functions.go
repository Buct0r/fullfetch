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
			configPath := filepath.Join(appdata, "fullfetch", "config.json")
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
		Art  string              `json:"art"`
		Arts map[string][]string `json:"arts"`
	}

	scheme := `{
    "scheme": "all",
    
    "schemes" : {

        "all": 
        {
        "art": true,
        "title": true,
        "os" : true,
        "host": true,
        "hostname"  : true,
        "kernel" : true,
        "uptime" : true,
        "bootime": true,
        "procs" : true,
        "cpu"  : true,
        "gpu" : true,
        "memory" : true,
        "swap": true,
        "disk" : true,
        "ip"   : true,
        "colors" : true,
        "locale" : true,
        "battery" : true,
        "credits" : true
        },

         "minimal": 
         {
            "art": false,
            "title": true,
            "os" : true,
            "host": false,
            "hostname"  : true,
            "kernel" : false,
            "cpu"  : true,
            "uptime" : false,
            "bootime": false,
            "procs" : false,
            "gpu" : true,
            "memory" : true,
            "swap": true,
            "disk" : false,
            "ip"   : false,
            "colors" : false,
            "locale" : false,
            "battery" : true,
            "credits" : false
        },

        "custom": 
        {
            "art": false,
            "title": true,
            "os" : true,
            "hostname"  : true,
            "kernel" : false,
            "uptime" : false,
            "bootime": true,
            "procs" : false,
            "cpu"  : true,
            "gpu" : true,
            "memory" : true,
            "swap": true,
            "disk" : true,
            "ip"   : true,
            "colors" : false,
            "locale" : true,
            "battery" : true,
            "credits" : false
        }
    },

    "order": "default",

    "orders": {
        "default": ["art", "title", "os","host", "hostname", "kernel", "uptime", "bootime", "procs", "cpu", "gpu", "memory", "swap", "disk", "ip", "battery", "locale", "colors"],
        "custom" : ["os", "hostname", "kernel", "uptime", "cpu", "gpu", "memory", "disk", "ip", "colors"]
    },

    "colorScheme": "mono",

    "colorSchemes": {

        "default": 
        {
            "title" : "Purple",
            "os": "White",
            "host": "White",
            "hostname":   "Red",
		    "kernel":     "Red",
		    "uptime":   "Green",
            "bootime": "Green",
            "procs": "Green",
		    "cpu":  "Yellow",
		    "gpu":    "Blue",
		    "memory": "Magenta",
            "swap": "Magenta",
		    "disk":    "Cyan",
		    "ip":    "Gray",
            "battery": "Green",
            "locale": "White",
            "Reset":   "Reset"
        },

        "custom": 
        {
            "os": "Red",
            "hostname":   "Reset",
		    "kernel":     "Red",
		    "uptime":   "Green",
		    "cpu":  "Yellow",
		    "gpu":    "Blue",
		    "memory": "Magenta",
		    "disk":    "Cyan",
		    "ip":    "Gray",
            "Reset":   "Reset"
        },

        "dark": 
        {
            "os": "White",
            "Reset":   "Reset"
        },

        "mono" : 
        {
            "title" : "Orange",
            "os": "Orange",
            "host": "Orange",
            "hostname":   "Orange",
		    "kernel":     "Orange",
		    "uptime":   "Orange",
            "bootime": "Orange",
            "procs": "Orange",
		    "cpu":  "Orange",
		    "gpu":    "Orange",
		    "memory": "Orange",
            "swap": "Orange",
		    "disk":    "Orange",
		    "ip":    "Orange",
            "battery": "Orange",
            "locale": "Orange",
            "Reset":   "Reset"
        }
    },

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
			panic(err)
		}

		selectedArt := cfg.Arts[cfg.Art]
		for _, line := range selectedArt {
			fmt.Println(line)
		}
		fmt.Print("\n")

	} else {

		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&cfg); err != nil {
			panic(err)
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
		if minutes != 0 {
			fmt.Printf("%sUptime:%s %d hours %d mins\n", color, reset, hours, minutes)
		} else {
			fmt.Printf("%sUptime:%s %d hours\n", color, reset, hours)
		}
	} else {
		fmt.Printf("%sUptime:%s %d mins\n", color, reset, uptime/60)
	}
}

func displayCpu(color string, reset string) {
	cpuInfo, _ := cpu.Info()
	model := cpuInfo[0].ModelName
	numCores, _ := cpu.Counts(false)
	GHZ := float64(cpuInfo[0].Mhz) / 1000
	//usage, _ := cpu.Percent(time.Second, false)  <-- makes the program run too slowly
	fmt.Printf("%sCPU:%s (%v) %v @ %.2fGHz\n", color, reset, numCores, model, GHZ) //usage[0]
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
