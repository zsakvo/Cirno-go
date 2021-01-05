package util

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func unicodeToZh(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func PanicErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func ErrTip(str string) string {
	var valid = regexp.MustCompile(`"tip":"(\S*)"`)
	var unicode = valid.FindAllStringSubmatch(str, -1)[0][1]
	v, _ := unicodeToZh([]byte(unicode))
	return "err: " + string(v)
}
