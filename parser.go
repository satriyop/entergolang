package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	// Header is front matter file header
	Header = "---\n"
	// Footer is front matter and content separator
	Footer = "\n---"
)

// ParseContent is to parse file that contain front matter and content
func parseContent(file string) (string, *FrontMatter) {
	fileread, err := readFile(file)
	check(err)

	fmStr, contentStr, err := parseFile(fileread)
	check(err)

	fm, err := getFrontMatter(fmStr)
	check(err)

	return contentStr, fm
}

// ParseContent will try to parse content that has front matter and post content
func getFrontMatter(fmStr string) (*FrontMatter, error) {
	fm := &FrontMatter{}

	err := yaml.Unmarshal([]byte(fmStr), fm)
	if err != nil {
		return nil, err
	}

	return fm, nil
}

func parseFile(fileread []byte) (string, string, error) {
	// TO DO replace with buffer
	post := string(fileread)

	fmStart := strings.HasPrefix(post, Header)
	if !fmStart {
		return "", "", errors.New("No Front Matter Found")
	}

	// trim --- at start of file
	post = strings.Trim(post, Header)
	// split between front matter and content
	doc := strings.Split(post, Footer)

	fmStr := doc[0]
	contentStr := doc[1]
	return fmStr, contentStr, nil
}

func readFile(filename string) (fileread []byte, err error) {
	fileread, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return fileread, nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
