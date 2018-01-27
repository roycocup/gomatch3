package debug

import "fmt"

func PrintV(o interface{}) {
	fmt.Print("\n\n")
	fmt.Print("####################\n")
	fmt.Printf("%#v", o)
	fmt.Print("####################\n")
	fmt.Print("\n")
}
