package main

import "fmt"

func main() {
	q0 := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tq0 := %q\n\tq0 = fmt.Sprintf(q0, q0)\n\tq := fmt.Sprintf(\"puts %%q\", q0)\n\tfmt.Println(q)\n}"
	q0 = fmt.Sprintf(q0, q0)
	q := fmt.Sprintf("puts %q", q0)
	fmt.Println(q)
}
