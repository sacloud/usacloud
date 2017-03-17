package command

type MonitorValues []MonitorValue
type MonitorValue map[string]string

func (v MonitorValues) Len() int {
	return len(v)
}

func (v MonitorValues) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v MonitorValues) Less(i, j int) bool {
	_, hasIndex := v[i]["Index"]
	if hasIndex {
		if v[i]["UnixTime"] == v[j]["UnixTime"] {
			return v[i]["Index"] < v[j]["Index"]
		}
	}
	return v[i]["UnixTime"] < v[j]["UnixTime"]
}
