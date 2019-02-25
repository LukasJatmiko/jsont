package jsont

import (
	"regexp"
	"testing"
)

func TestToJSON(t *testing.T) {
	a := make(map[string]interface{})
	c := make(map[string]interface{})
	c["d"] = 1
	e := make(map[string]interface{})
	e["f"] = "Success"
	g := make(map[string]interface{})
	g["h"] = true
	var ar []map[string]interface{}
	ar = append(ar, c)
	ar = append(ar, e)
	ar = append(ar, g)
	a["b"] = ar
	var j int64
	j = 100
	a["i"] = j
	a["k"] = 2.00
	a["l"] = nil

	rslt := ToJSON(a)

	// {"b":[{"d":1},{"f":"Success"},{"h":true}],"i":100,"k":2.00,"l":""}
	rgx := regexp.MustCompile(`(?i)\"b\"\:\[\{\"d\"\:1\}\,\{\"f\"\:\"Success\"}\,\{\"h\"\:true\}\]`)
	if !rgx.Match([]byte(rslt)) {
		t.Errorf("JSON Output is not match (%v)", `"b":[{"d":1},{"f":"Success"},{"h":true}]`)
	}
	rgx = regexp.MustCompile(`(?i)\"i\"\:100`)
	if !rgx.Match([]byte(rslt)) {
		t.Errorf("JSON Output is not match (%v)", `"i":100`)
	}
	rgx = regexp.MustCompile(`(?i)\"k\"\:2\.00`)
	if !rgx.Match([]byte(rslt)) {
		t.Errorf("JSON Output is not match (%v)", `"k":2.00`)
	}
	rgx = regexp.MustCompile(`(?i)\"l\"\:\"\"`)
	if !rgx.Match([]byte(rslt)) {
		t.Errorf("JSON Output is not match (%v)", `"l":""`)
	}
}
