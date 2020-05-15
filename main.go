// hello world!
package main

import "fmt"
import "runtime"
import "os/exec"


func print(text string){
	fmt.Println(text)
}

func main(){
	fmt.Println("Hello", runtime.GOOS)
	if runtime.GOOS == "linux" {
		print("Sending message...")
		app := "notify-send"
		arg1 := "'Test titel'"
		arg2 := "'test message'"
		exec.Command(app, arg1, arg2)

	}


}
