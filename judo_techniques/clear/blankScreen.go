package clear

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func() //makes a map for storing clear func expressions / functions

func init() {
	clear = make(map[string]func()) // initialize the clear variable, making it accessible to the rest of our code
	clear["linux"] = func() {
		cmd := exec.Command("clear") //this command clears the terminal on Unix Os
		cmd.Stdout = os.Stdout       //gives the cmd variable access to Stdout on the os
		cmd.Run()                    //executes the command
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls") //clears the terminal in Windows Os
		cmd.Stdout = os.Stdout     //gives cmd variable access to Stdout on the os
		cmd.Run()
	}

}

//CallClear is a function to clear the terminal screen, needs to be exported to use
func CallClear() {
	value, ok := clear[runtime.GOOS] //checks os, linux, windows, etc
	if ok {                          //if we defined a clear function for that platform
		value() //execute it
	} else { //if platform unsupported
		fmt.Println("Cannot clear screen between selections with this Operating System.")
	}
}

func main() {
	fmt.Println("Refreshing screen!")
	time.Sleep(1 * time.Second)
	CallClear()

}
