# noti
## About
A simple cross-platform notifier library.

It uses the operating system's native notification system. For Windows, this
is toast notifications, for Linux it uses `notify-send`

Platform support:

- [x] Linux
- [x] Windows 8 and Windows 10
- [ ] MacOS 10.11 and higher (I'm still working on getting macos support)


![A screenshot of noti on the GNOME desktop](https://raw.githubusercontent.com/daniel071/noti/master/Screenshots/notiExample.png)

## Installation
### Using precompiled binaries:
**(Recommended method)**
1. Download the precompiled binary from the releases page
2. In the terminal, if you're on linux/unix, type `./noti` and if you're
on windows, type `.\noti.exe`


### Compiling it yourself:
**(Experts only)**

NOTE: This requires you to have Go installed
1. Clone this repository with `git clone https://github.com/daniel071/noti.git`
2. Run this with `go run noti.go` **OR** compile this with `go build noti.go`
which will then create an executable.

## Usage:
Using noti is very simple, all you have to do is:

`.\noti.exe title message` - If you are on Windows

`./noti title message` - If you are on Linux

Note: Replace "title" and "message" with your own text.
