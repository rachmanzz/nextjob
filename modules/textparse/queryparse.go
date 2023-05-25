package textparse

import (
	"regexp"
	"strings"
)

func QueryParse(q string) (*[]string, *[]map[string]string) {
	regex := regexp.MustCompile(`(--[a-zA-Z\-]+)=([^=]*\w)(?:\s|$)`)
	var positions []string = []string{}
	var params []map[string]string

	argsGroup := regex.FindAllString(q, -1)
	indexGroup := regex.FindAllStringIndex(q, -1)
	if len(indexGroup) >= 1 {
		argStart := indexGroup[0][0]
		if argStart >= 1 {
			positionStr := q[0:argStart]
			positionArr := strings.Split(positionStr, ";")
			positions = append(positions, positionArr...)
		}
	} else {
		positions = append(positions, q)
	}

	if len(argsGroup) >= 1 {
		for _, arg := range argsGroup {
			keyVal := strings.Split(arg, "=")

			if len(keyVal) == 2 {
				argVal := strings.TrimSpace(keyVal[1])

				if argVal != "" {
					switch keyVal[0] {
					case "--loc":
						locations := strings.Split(argVal, ";")
						for _, location := range locations {
							params = append(params, map[string]string{
								"location[]": strings.TrimSpace(location),
							})
						}
					case "--type":
						params = append(params, map[string]string{"type": argVal})
					case "--currency":
						params = append(params, map[string]string{"currency": argVal})
					case "--salary":
						salaryArr := strings.Split(argVal, "-")
						if len(salaryArr) >= 1 {
							params = append(params, map[string]string{"salaryFrom": salaryArr[0]})

							if len(salaryArr) == 2 {
								params = append(params, map[string]string{"salaryTo": salaryArr[1]})
							}
						}
					}
				}
			}
		}
	}

	return &positions, &params
}
