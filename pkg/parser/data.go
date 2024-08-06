package parser

func PostgresInterfaceToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case []byte:
		str := string(v)

		if str == "{NULL}" {
			return ""
		}
		return str
	case string:

		return v
	default:
		return ""
	}
}
