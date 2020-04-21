package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Running PHP Code")

	filename := "test.php"
	codeSnippet := readFile(filename)
	runPHP(codeSnippet)

}

func runPHP(f string) {
	checkPhpExists()

	cmd := exec.Command("php")
	cmd.Stdin = strings.NewReader(f)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	check(err)
	fmt.Println(out.String())

}

func runDockerPhp(f string) {
	checkDockerExists()

}

func checkPhpExists() {
	path, err := exec.LookPath("php")
	if err != nil {
		fmt.Printf("didn't find 'php' executable\n")
	} else {
		fmt.Printf("'php' executable is in '%s'\n", path)
	}
}

func checkDockerExists() {
	path, err := exec.LookPath("docker")
	if err != nil {
		fmt.Printf("didn't find 'docker' executable\n")
	} else {
		fmt.Printf("'docker' executable is in '%s'\n", path)
	}
}

func readFile(f string) string {
	content, err := ioutil.ReadFile(f)
	check(err)

	t := string(content)
	return t
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
