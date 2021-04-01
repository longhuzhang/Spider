package paraseURL

import (
	"fmt"
	"strconv"
	"strings"
)

func ParaseUnicode(source string) string {
	sUnicodev := strings.Split(source, `\u`)
	aim := ""
	for _, v := range sUnicodev {
		if v == ""{
			continue
		}
		fmt.Println("v",v)
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		aim += fmt.Sprintf("%c", temp)
	}
	return aim
}
