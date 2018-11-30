package jsont

import "strconv"

//ToJSON : Create JSON object from map string interface
func ToJSON(obj map[string]interface{}) string {
	var jsonStr = "{"
	ctr := 1
	for k, v := range obj {
		if ctr != 1 {
			switch v.(type) {
			case int:
				{
					jsonStr = jsonStr + `,"` + k + `":` + strconv.Itoa(v.(int))
				}
			case string:
				{
					jsonStr = jsonStr + `,"` + k + `":"` + v.(string) + `"`
				}
			case map[string]interface{}:
				{
					jsonStr = jsonStr + `,"` + k + `":` + ToJSON(v.(map[string]interface{}))
				}
			}
		} else {
			switch v.(type) {
			case int:
				{
					jsonStr = jsonStr + `"` + k + `":` + strconv.Itoa(v.(int)) + `,`
				}
			case string:
				{
					jsonStr = jsonStr + `"` + k + `":"` + v.(string) + `",`
				}
			case map[string]interface{}:
				{
					jsonStr = jsonStr + `"` + k + `":` + ToJSON(v.(map[string]interface{})) + `,`
				}
			}
		}
		ctr++
	}
	jsonStr = jsonStr + "}"
	return jsonStr
}
