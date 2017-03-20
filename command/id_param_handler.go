package command

import (
	"strconv"
)

func toSakuraID(id string) (int64, bool) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, false
	}
	return i, true
}

func isSakuraID(id string) bool {
	_, err := strconv.ParseInt(id, 10, 64)
	return err == nil
}
