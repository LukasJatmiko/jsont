package jsont

import "strconv"

//ToJSON : Create JSON object from map string interface
func ToJSON(obj interface{}) string {
	var jsonStr string
	switch obj.(type) {
	case map[string]interface{}:
		{
			jsonStr = "{"
			ctr := 1
			for k, v := range obj.(map[string]interface{}) {
				if ctr != 1 {
					jsonStr = jsonStr + `,`
				}
				switch v.(type) {
				case int:
					{
						jsonStr = jsonStr + `"` + k + `":` + strconv.Itoa(v.(int))
					}
				case int64:
					{
						jsonStr = jsonStr + `"` + k + `":` + strconv.FormatInt(v.(int64), 10)
					}
				case float64:
					{
						jsonStr = jsonStr + `"` + k + `":` + strconv.FormatFloat(v.(float64), 'f', 2, 64)
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
						jsonStr = jsonStr + `"` + k + `":` + ToJSON(v.([]map[string]interface{}))
					}
				case nil:
					{
						jsonStr = jsonStr + `"` + k + `":""`
					}
				}
				ctr++
			}
			jsonStr = jsonStr + "}"
		}
	case []map[string]interface{}:
		{
			jsonStr = "["
			var sv string
			for _, sObj := range obj.([]map[string]interface{}) {
				sv += `,` + ToJSON(sObj)
			}
			if len(sv) > 0 {
				sv = sv[1:]
			}
			jsonStr += sv + "]"
		}
	}
	return jsonStr
}
