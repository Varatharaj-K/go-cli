package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var flags struct {
	Package string
	Out     string
}

func main() {
	cmd := flag.NewFlagSet("generate", flag.ExitOnError)

	args := cmd.Args()
	run(args)

}

func run(args []string) {
	api := readFile(args[0])
	tapi := convert(api)
	w := new(bytes.Buffer)
	check(Go.Execute(w, tapi))
	check(GoFmt(flags.Out, w))
}

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}

func readFile(file string) API {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	var api API
	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)
	err = dec.Decode(&api)
	check(err)

	return api
}

func check(err error) {
	if err != nil {
		fatalf("%v", err)
	}
}

func checkf(err error, format string, otherArgs ...interface{}) {
	if err != nil {
		fatalf(format+": %v", append(otherArgs, err)...)
	}
}
