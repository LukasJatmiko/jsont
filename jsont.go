package jsont

import "strconv"

//ToJSON : Create JSON object from map string interface
func ToJSON(obj map[string]interface{}) string {
	var jsonStr = "{"
	ctr := 1
	for k, v := range obj {
		if ctr != 1 {
			jsonStr = jsonStr + `,`
		}
		switch v.(type) {
		case int:
			{
				jsonStr = jsonStr + `"` + k + `":` + strconv.Itoa(v.(int))
			}
		case bool:
			{
				jsonStr = jsonStr + `"` + k + `":` + strconv.FormatBool(v.(bool))
			}
		case string:
			{
				jsonStr = jsonStr + `"` + k + `":"` + v.(string) + `"`
			}
		case map[string]interface{}:
			{
				jsonStr = jsonStr + `"` + k + `":` + ToJSON(v.(map[string]interface{}))
			}
		case []map[string]interface{}:
			{
				sv := ""
				for _, mv := range v.([]map[string]interface{}) {
					sv += `,` + ToJSON(mv)
				}
				if len(sv) > 0 {
					sv = sv[1:]
				}
				jsonStr = jsonStr + `"` + k + `":[` + sv + `]`
			}
		case nil:
			{
				jsonStr = jsonStr + `"` + k + `":""`
			}
		}
		ctr++
	}
	jsonStr = jsonStr + "}"
	return jsonStr
}
