package main

import (
	"fmt"
)

func main() {
	chr := byte(96)
	q := `package main

import (
	"fmt"
)

func main() {
	chr := byte(96)
	q := %c%v%c
	fmt.Printf(q, chr, q, chr)
}
`
	fmt.Printf(q, chr, q, chr)
}
