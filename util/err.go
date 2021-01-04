package util

import (
	"fmt"
	"os"
)

func PanicErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
