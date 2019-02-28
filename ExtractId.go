package go_sportradar_football

import (
	"errors"
	"strconv"
	"strings"
)

func ExtractId(str string) (int, error) {
	split := strings.Split(str, ":")
	if len(split) == 0 {
		return 0, errors.New("Split failed")
	}

	id, err := strconv.Atoi(split[len(split)-1])
	if err != nil {
		return 0, err
	}

	return id, nil
}
