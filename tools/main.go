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

	args := cmd.String("file", "", "File Name")
	if len(os.Args) < 2 {
		fmt.Println("expected 'generate' subcommand")
		os.Exit(1)
	}

	Run(cmd, args)

}

func Run(cmd *flag.FlagSet, args *string) {
	cmd.Parse(os.Args[2:])
	api := readFile(args)
	tapi := convert(api)
	w := new(bytes.Buffer)
	check(Go.Execute(w, tapi))
	check(GoFmt(flags.Out, w))
}

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}

func readFile(file *string) API {
	fil := *file
	fmt.Print("File " + fil)
	f, err := os.Open(fil)
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
