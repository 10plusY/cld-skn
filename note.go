package cloudskine

import (
	"fmt"
	"regexp"
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
	tags := make([]string, len(htags))
	copy(tags, htags)
	for i := range btags {
		for _, tag := range tags {
			if btags[i] != tag {
				tags = append(tags, btags[i])
			}
		}
	}
	return tags
}

func (n Note) toDict(tagged, separate bool) map[string]interface{} {
	dict := make(map[string]interface{})
	dict["header"] = n.Header
	dict["body"] = n.Body

	if n.hasNamespace() == true {
		dict["namespace"] = n.Namespace
	}

	if tagged == true {
		if separate == true {
			dict["headertags"], dict["bodytags"] = n.parseTags()
		} else {
			dict["tags"] = n.parseAllTags()
		}
	}

	return dict
}

func (n Note) toRecord(dict map[string]interface{}) []string {
	rec := make([]string, len(dict))
	for key, val := range dict {
		if n.annotate == false {
			if _, ok := val.(string); ok == true {
				rec = append(rec, val.(string))
			} else {
				rec = append(rec, strings.Join(val.([]string), ""))
			}	
		} else {
			if _, ok := val.(string); ok == true {
				rec = append(rec, fmt.Sprintf("%s%s", key, val.(string)))
			} else {
				rec = append(rec, fmt.Sprintf("%s%s", key, strings.Join(val.([]string), "")))
			}
		}
	}
	return rec
}

func (n Note) ToRecord() []string {
	return n.toRecord(n.toDict(false, false))
}

func (n Note) ToTaggedRecord(separate bool) []string {
	return n.toRecord(n.toDict(true, separate))
} 
