package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func run(cmd_in ...string) {
	//fmt.Printf("Running %v from run\n", os.Args[2:])
	fmt.Printf("Running %v from run\n", cmd_in)
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	//cmd := exec.Command("/proc/self/exe", append([]string{"child"}, cmd_in[0:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Cloneflags only available in Linux
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	must(cmd.Run())
}

func child(cmd_in ...string) {
	fmt.Printf("Running %v from child\n", cmd_in)
	//cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd := exec.Command(cmd_in[0], cmd_in[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("my-container")))
	must(cmd.Run())

}

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Unknown command")
	}
}
