package helper

import "fmt"

func PanicIfError(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf(msg, err))
	}
}
