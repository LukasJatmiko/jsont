package jsont

import (
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

	rslt := ToJSON(a)
	if rslt != `{"b":[{"d":1},{"f":"Success"},{"h":true}]}` {
		t.Errorf("JSON Output is not match (%v)", rslt)
	}
}
