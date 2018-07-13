package define

func formatBoolFunc(key string) func(map[string]string) string {
	return func(values map[string]string) string {
		v, ok := values[key]
		if !ok {
			return "false"
		}
		if v == "true" || v == "True" {
			return "true"
		}
		return "false"
	}
}
