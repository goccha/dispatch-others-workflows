package debug

import "fmt"

func Print(name, value string) {
	fmt.Printf("::debug::%s=%s\n", name, value)
}
