package learngomodule

import "fmt"

func print(slc []string) {
	for _, v := range slc {
		fmt.Println(v)
	}
}
