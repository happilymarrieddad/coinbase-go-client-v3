package coinbasegoclientv3

// AssignStrIfSetToMap maps are pointers so this works
func AssignStrIfSetToMap(key string, val string, params map[string]string) {
	if len(key) > 0 {
		params[key] = val
	}
}

func StrArrToStr(arr []string) (str string) {
	for idx, s := range arr {
		if idx > 0 {
			str += ","
		}
		str += s
	}

	return str
}
