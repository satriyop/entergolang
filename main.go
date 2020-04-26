package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Session preparation")

	fmt.Println("Running Code in Docker")
	startServer()
	// runDockerPhp()
}

// Session
// var globalSessions *Manager

func runDockerPhp() string {
	var res string
	isExist, err := checkDockerExists()
	check(err)
	if isExist {
		res = startDocker()
	}
	return res
}

func startDocker() string {
	fmt.Println("going to run code in docker, please wait...")

	cmd := exec.Command("bash", "-c", "docker run -w /app --rm --volumes-from phpcontainer php:7.4-alpine php test.php")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	check(err)
	r := out.String()

	return r
}

func checkDockerExists() (bool, error) {
	path, err := exec.LookPath("docker")
	if err != nil {
		fmt.Printf("didn't find 'docker' executable\n")
		return false, err
	} else {
		fmt.Printf("'docker' executable is in '%s'\n", path)
		return true, nil
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
