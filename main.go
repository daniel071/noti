// hello world!
package main

import "fmt"
import "runtime"
import "os/exec"
import "os"


func print(text string){
	fmt.Println(text)
}

func main(){
	firstArg := os.Args[1]
	secondArg := os.Args[2]

	fmt.Println("Hello", runtime.GOOS)
	if runtime.GOOS == "linux" {
		print("Sending message...")
		app := "notify-send"
		command := exec.Command(app, firstArg, secondArg)
		command.Output()


	}


}
