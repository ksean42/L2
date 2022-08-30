package main

import (
	"bufio"
	"fmt"
	p "github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func parse(input string) {
	commands := strings.Split(input, "|")

	for _, v := range commands {
		tmp := strings.Trim(v, " ")
		args := strings.Split(tmp, " ")
		switch args[0] {
		case "echo":
			echo(args[1:])
		case "pwd":
			pwd()
		case "cd":
			cd(args[1])
		case "ps":
			ps()
		case "kill":
			kill(args[1])
		default:
			execute(args)
		}
	}
}

func execute(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

}

func cd(dest string) {
	if dest == ".." {
		curDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		dest = filepath.Dir(curDir)
	}

	dir, err := os.Stat(dest)
	if err != nil {
		fmt.Println("cd: not exists", dest)
		return
	}
	if !dir.IsDir() {
		fmt.Println("cd: not a directory", dest)
		return
	}
	if err := os.Chdir(dest); err != nil {
		fmt.Println(err)
		return
	}
}
func ps() {
	processes, err := p.Processes()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%10s %10s %10s\n", "PID", "Parent PID", "Name")
	for _, v := range processes {
		var p p.Process
		p = v
		fmt.Printf("%10d  %10d %10s\n", p.Pid(), p.PPid(), p.Executable())
	}
}

func kill(arg string) {
	pid, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Wrong pid")
		return
	}
	err = syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
}
func echo(args []string) {

	fmt.Println(strings.Join(args, " "))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print(">> ")
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	go gracefulShutdown(exit)

	for sc.Scan() {
		input := sc.Text()
		if input == "quit" {
			return
		}
		parse(input)
		fmt.Print(">> ")
	}
}

func gracefulShutdown(exit chan os.Signal) {
	<-exit
	fmt.Println("\nBye!")
	os.Exit(0)
}
