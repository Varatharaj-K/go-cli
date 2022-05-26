package main

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
)

// API is an API with a set of methods.
type API map[string]Method

// Method is a method of an API.
type Method struct {
	Kind         string
	Description  string
	Deprecated   string
	Experimental bool
	RPC          string
	Input        string
	Output       string
	Call         string
	CallParams   []string `yaml:"call-params"`
	Validate     []string `yaml:"validate"`
}

func GoFmt(filePath string, buf *bytes.Buffer) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, buf, parser.ParseComments)
	if err != nil {
		// If parsing fails, write out the unformatted code. Without this,
		// debugging the generator is a pain.
		_, _ = buf.WriteTo(f)
	}
	if err != nil {
		return err
	}

	err = format.Node(f, fset, file)
	if err != nil {
		return err
	}

	err = exec.Command("go", "run", "github.com/rinchsan/gosimports/cmd/gosimports", "-w", filePath).Run()
	if err != nil {
		return err
	}

	return nil
}
