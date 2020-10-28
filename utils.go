package main

import (
	"bufio"
	"github.com/kr/pty"
	"io"
	"log"
	"os"
	"os/exec"
)

func checkBinaryExist(command string) (bool, string) {
	path, err := exec.LookPath(command)

	return err == nil, path
}

func runCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), err
}

func runCommandTty(directory string, cmdName string, arg ...string) {
	cmd := exec.Command(cmdName, arg...)
	cmd.Dir = directory
	tty, err := pty.Start(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	defer tty.Close()

	go func() {
		scanner := bufio.NewScanner(tty)
		for scanner.Scan() {
			log.Println("[" + cmdName + "] " + scanner.Text())
		}
	}()
	go func() {
		io.Copy(tty, os.Stdin)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}

func copyFile(source string, dest string) {
	from, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}
