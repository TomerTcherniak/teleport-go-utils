package main

import (
	"bytes"
	"log"
	"os/exec"
)

// Shellcheck is function to check error
func Shellcheck(commandrun string) {
	err := Shellout(commandrun)
	if err != nil {
		log.Print("error: ", err)
		panic(err)
	}
}

// Shellout is function to run command
func Shellout(command string) error {
	const ShellToUse = "bash"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err
}
