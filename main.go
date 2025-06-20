package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const version = "2.1.0"

func main() {
	var showVersion bool //aggiungere flag per vedere dove si ha generato il config
	var generate bool
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.BoolVar(&showVersion, "v", false, "print version and exit")
	flag.BoolVar(&generate, "gen", false, "generetase config file")
	flag.Parse()

	if showVersion {
		fmt.Println("fullfetch version", version)
		os.Exit(0)
	}

	path := getConfigPath()
	text := `{
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

	if generate {
		if path == "" {
			configDir, _ := os.UserConfigDir()
			configPath := filepath.Join(configDir, "fullfetch", "config.json")
			if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
				panic(err)
			}
			newFile, err := os.Create(configPath)
			if err != nil {
				panic(err)
			} else {
				defer newFile.Close()
			}
			_, err = newFile.WriteString(text + "\n")

			if err != nil {
				panic(err)
			}
			fmt.Println("Config file created succesfully, you can now edit it at", configPath)
		} else {
			fmt.Println("Config file already exist, operation aborted :/")
		}
		os.Exit(0)
	}

	fmt.Print("\n \n")

	//color := "\033[38;5;208m"
	//reset := "\033[0m"

	loadConfig(text)

	fmt.Print("\n")

}
