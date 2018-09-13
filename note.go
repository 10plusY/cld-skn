package cloudskine

import (
	"fmt"
	"regexp"
	"io"
	"strings"
)

type Note struct {
	Namespace string
	Header    string
	Body      string
	tagChar   rune
	annotate  bool
}

func (n Note) hasNamespace() bool {
	return len(strings.TrimSpace(n.Namespace)) != 0
}

func (n Note) compileTagRegex() (*regexp.Regexp, error) {
	r, err := regexp.Compile(fmt.Sprintf("(?:^|/s)(?:%s)([a-zA-Z/d]+)", string(n.tagChar)))
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (n Note) isTagged() bool {
	r, err := n.compileTagRegex()
	if err != nil {
		return false
	}

	return r.MatchString(n.Header) || r.MatchString(n.Body) 
}

func (n Note) parseTags() ([]string, []string) {
	r, err := n.compileTagRegex()
	if err != nil {
		return []string{}, []string{}
	}

	return r.FindAllString(n.Header, -1), r.FindAllString(n.Body, -1)
}

func (n Note) parseAllTags() []string {
	htags, btags := n.parseTags()
	return append(htags, btags...)
}

func (n Note) recordVals() []interface{} {
	if n.hasNamespace() == true {
		return []interface{}{n.Namespace, n.Header, n.Body}
	}
	return []interface{}{n.Header, n.Body}
}

func (n Note) annotationList() []string {
	if n.hasNamespace() == true {
		return []string{"NAMESPACE", "HEADER", "BODY"}
	}
	return []string{"HEADER", "BODY"}
}

func (n Note) toRecord() string {
	var tks []string
	vals := n.recordVals()

	if n.annotate == false {
		for _ = range vals {
			tks = append(tks, "%s")
		}
	} else {
		tks = n.annotationList()	
		for i, _ := range vals {
			shift := 2 * i + 1
			tks = append(tks[:shift], append([]string{"%s"}, tks[shift:]...)...)
		}
	}

	return fmt.Sprintf(strings.Join(tks, ","), vals...)
}

func (n Note) toTaggedRecord(separate bool) string {
	var rstr string
	rec := n.toRecord()

	if n.annotate == false {
		if separate == true {
			rstr = "%s,%s,%s"
		} else {
			rstr = "%s,%s"
		}
	} else {
		if separate == true {
			rstr = "%s,HEADERTAGS:%s,BODYTAGS:%s"
		} else {
			rstr = "%s,TAGS:%s"
		}
	}

	if separate == true {
		htags, btags := n.parseTags()
		return fmt.Sprintf(rstr, rec, htags, btags)
	}

	return fmt.Sprintf(rstr, rec, n.parseAllTags())
}

func (n Note) bufferNote() io.Reader {
	return strings.NewReader(n.toRecord())
}
