package cloudskine

import (
	"fmt"
	"regexp"
)

type Note struct {
	Header  string
	Body    string
	TagChar string
}

func (n Note) compileTagRegex() (*regexp.Regexp, error) {
	r, err := regexp.Compile(fmt.Sprintf("(?:^|/s)(?:%s)([a-zA-Z/d]+)", n.TagChar))
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (n Note) 