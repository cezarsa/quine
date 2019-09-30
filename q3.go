package main

func main() {
	chr := byte(96)
	q := `package main

func main() {
	chr := byte(96)
	q := AHOY QUINE!!!
	println(q[:51] + string(chr) + q + string(chr) + q[64:])
}`
	println(q[:51] + string(chr) + q + string(chr) + q[64:])
}
