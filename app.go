package learngomodule

import "fmt"

func PrintSlc(slc []string) {
	for _, v := range slc {
		fmt.Println(v)
	}
}
