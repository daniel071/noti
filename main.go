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
	fmt.Println("""
	noti - a simple notification library written in go.
	USAGE:
	 noti [title] [message]
	""")

}


func main(){
	firstArg := os.Args[1]
	secondArg := os.Args[2]

	if len(os.Args) > 0 {
		fmt.Println("Hello", runtime.GOOS)

		if runtime.GOOS == "linux" {
			print("Sending message...")
			app := "notify-send"
			command := exec.Command(app, firstArg, secondArg)
			command.Output()

		}

	}	else {
		help()

	}


}
