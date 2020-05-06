package main

import (
	"errors"
	"io/ioutil"
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
func parseContent(file string) (string, *FrontMatter, error) {
	fileread, err := readFile(file)
	if err != nil {
		return "", nil, err
	}

	fmStr, contentStr, err := parseFile(fileread)
	if err != nil {
		return "", nil, err
	}

	fm, err := getFrontMatter(fmStr)
	if err != nil {
		return "", nil, err
	}

	return contentStr, fm, nil
}

// ParseContent will try to parse content that has front matter and post content
func getFrontMatter(fmStr string) (*FrontMatter, error) {
	fm := &FrontMatter{}

	err := yaml.Unmarshal([]byte(fmStr), fm)
	if err != nil {
		return nil, errors.New("Can not unmarshal yaml")
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
	// TODO better reading file with buffer and error checking
	fileread, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("File is not found")
	}

	return fileread, nil
}
