<p align= "center" style="font-size: 2em; font-weight: bold;">
    fullfetch
</p>

<p align="center">
<img src="imgs/logo.png" alt="logo">
</p>


<p style="font-size: 1.5em; font-weight: bold;" align="center">
A completely customizable CLI system information tool made with Go
</p><br><br>

<p style="font-size: 1.1em;" align="center">
fullfetch is a neofetch and fastfetch inspired CLI application for showing system
informations and fully customizable. The customization can be modified by editing
the <code>config.json</code>.</p><br>

<p align="center">
<img src="imgs/fullfetch_ub.png" width="400px"><br>
</p>

<p style="font-size: 1.1em;" align="center">
fullfetch displays the selected properties in the config file. You can fully select the amount of properties, the order and the color for the titles. Before the output fullfetch displays a custom ASCII art, fully replaceable and customizable. In the config file you can also find pre-customized schemes.</p>

<p align="center">
<img src="imgs/fullfetch_min.png" alt="fullfetch">
</p>

<p align="center">fullfetch supports every system (in theory), and in case something is not supported open a GitHub issue.</p><br><br>

# Installation

## Windows
### Pre-Built installer
Go to the [release page](https://github.com/Buct0r/fullfetch/releases) and download the <code>fullfetch_installer_(your_arch).exe</code>.

### Portable version
Download the <code>fullfetch_portable_windows_(your_arch).zip</code> and extract it wherever you want. To make fullfetch avaible in every terminal, add the extracted folder address to your PATH variable.

### Winget

Coming soon...


## Linux
### Pre-built package
Head to the [release page](https://github.com/Buct0r/fullfetch/releases) and download the correct package for your distro and architecture. For now It's only avaible for debian based distros for both x86 and ARM, and rpm package only in the x86 version.

#### Debian based:
Download the <code>fullfetch_(your_arch).deb</code>. Head to the download folder with a terminal and type:
```bash
$ sudo dpkg -i fullfetch_(your_arch).deb
```
Reload the terminal and type
```bash
$ fullfetch
```

#### rpm based (fedora, Red Hat)
Download the <code>fullfetch-1.0-1.fc42.x86_64.rpm</code>(For now only avaible for x86). Head to the download folder with a terminal and type:
```bash
$ sudo dnf install ~/(Downloaddir)/fullfetch-1.0-1.fc42.x86_64.rpm
```
Reload the terminal and type
```bash
$ fullfetch
```

## Mac OS
Head to the [release page](https://github.com/Buct0r/fullfetch/releases) and download the <code>fullfetch_portable_macos_(your_arch).zip</code> file. Extract it wherever you want and add the binary to your path variable. For now fullfetch is only available in its portable version.

# Building

Requirements:
- Go (possibly the latest version)
- cloned repo
<br>

1) Clone the Github repo
```
git clone https://github.com/Buct0r/fullfetch.git
```
2) Head to the cloned directory and in the terminal type: 
```
go build .
```

# Customization
You can customize pretty much every aspect of fullfetch. You can choose what to display, in wich order and what color should the parameters be displayed.

Parameters:
- art    
- title (user@hostname)   
- os       
- hostname
- kernel   
- uptime 
- processes    
- cpu    
- gpu  
- memory   
- disk   
- ip      
- colors  
- credits  
- locale   
- battery 
- host

You can also change the order of the displayed parameters by modifing the <code>"order"</code> section in the <code>config.json</code> file

There are also a vast range of colors available:
- Red
- Green
- Yellow
- Blue
- Magenta
- Cyan
- Gray
- White
- Orange
- Purple
- Pink
- Brown
- Black
- LightGray
- DarkGray
<br>
And all of the light versions of the colors.

You can choose between different pre-made color schemes and also make your own

The last thing that you can customize is the ASCII art displayed before all of the parameters. The default <code>config.json</code> file includes 3 different versions of the logo in ASCII art, but you can also make your own customized ASCII art to make fullfetch more beautifull


# Future updates
## Log 1 

Pending ⚠️<br>
Next features:
- New customization options
- More parameters to display
- ARM build for rpm
- Adding fullfetch to official packages repositories (apt, dnf, AUR ecc.)
- Adding fullfetch to winget repository
- tar.gz file for every os and arch
- adding fullfetch to homebrew

# Conclusion
Thank you for your support if you decide to install fullfetch. In the next days I will work to complete the majority of the points for the first log in the future updates section. <br>
Thanks to skesko for the help during the testing phase. 

Developed by Buct0r ❤️