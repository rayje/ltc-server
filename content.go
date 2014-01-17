package main

import (
	"io/ioutil"
	"errors"
)

var smallContent Content = Content{"small.out", ""}
var medContent Content = Content{"med.out", ""}
var largeContent Content = Content{"large.out", ""}
var xlargeContent Content = Content{"xlarge.out", ""}

type Content struct {
	filename string
	content string
}

func (c *Content) fileContent() (string, error) {
	result, err := ioutil.ReadFile(c.filename)
	return string(result), err
}

func (c *Content) getFileContent() (string, error) {
	var err error
	if len(c.content) == 0 {
		c.content, err = c.fileContent()
	}
	return c.content, err
}

func getContent(reqType string) (string, error) {
	switch reqType {
	case "small":
		return smallContent.getFileContent()
	case "med":
		return medContent.getFileContent()
	case "large":
		return largeContent.getFileContent()
	case "xlarge":
		return xlargeContent.getFileContent()
	}
	return "", errors.New(reqType + " is not valid")
}