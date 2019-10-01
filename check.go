package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var code = 0
	matches, err := filepath.Glob("./*.go")
	if err != nil {
		panic(err)
	}
	if os.Getenv("RECURSE") == "0" {
		goto out
	}
	for _, f := range matches {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			panic(err)
		}
		cmd := exec.Command("go", "run", f)
		cmd.Env = append(os.Environ(), "RECURSE=0")
		out, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(data, out) {
			code = 1
			goto out
		}
	}
out:
	q := "package main\n\nimport (\n\t\"bytes\"\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"os\"\n\t\"os/exec\"\n\t\"path/filepath\"\n)\n\nfunc main() {\n\tvar code = 0\n\tmatches, err := filepath.Glob(\"./*.go\")\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\tif os.Getenv(\"RECURSE\") == \"0\" {\n\t\tgoto out\n\t}\n\tfor _, f := range matches {\n\t\tdata, err := ioutil.ReadFile(f)\n\t\tif err != nil {\n\t\t\tpanic(err)\n\t\t}\n\t\tcmd := exec.Command(\"go\", \"run\", f)\n\t\tcmd.Env = append(os.Environ(), \"RECURSE=0\")\n\t\tout, err := cmd.CombinedOutput()\n\t\tif err != nil {\n\t\t\tpanic(err)\n\t\t}\n\t\tif !bytes.Equal(data, out) {\n\t\t\tcode = 1\n\t\t\tgoto out\n\t\t}\n\t}\nout:\n\tq := %q\n\tfmt.Printf(q, q)\n\tos.Exit(code)\n}\n"
	fmt.Printf(q, q)
	os.Exit(code)
}
