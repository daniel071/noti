// hello world!
package main

import "fmt"
import "runtime"
import "os/exec"
import "os"


func print(text string) {
	fmt.Println(text)
}


func help() {
	fmt.Println(
	"----- noti version 0.1.0 ----\n" +
	"noti - a simple notification library written in go.\n" +
	"USAGE:\n" +
	" noti [title] [message]",
	)

}


func main(){
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		secondArg := os.Args[2]

		if runtime.GOOS == "linux" {
			app := "notify-send"
			command := exec.Command(app, firstArg, secondArg)
			command.Output()

		} else if runtime.GOOS == "windows" {
			notification := toast.Notification{
	        AppID: "Example App",
	        Title: firstArg,
	        Message: secondArg,
			notification.Push()


		} else if runtime.GOOS == "darwin" {
			print("MacOS is not currently supported!")

		} else {
			fmt.Println("Your current operating system,", runtime.GOOS, ", is not supported.")


		}

	}	else {
		help()

	}


}
