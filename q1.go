package main

import (
	"fmt"
)

func main() {
	q := "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tq := %q\n\tfmt.Printf(q, q)\n}\n"
	fmt.Printf(q, q)
}
