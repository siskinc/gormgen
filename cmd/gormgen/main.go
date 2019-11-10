package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/siskinc/gormgen"
)

type config struct {
	structs []string
	client  string
}

var cnf config

func parseFlags() {
	var structs, client string
	flag.StringVar(&structs, "structs", "", "[Required] The name of schema structs to generate structs for, comma seperated")
	//flag.StringVar(&output, "output", "", "[Required] The name of the output file")
	flag.StringVar(&client, "client", "", "[Required] The name of *grom.DB object")
	flag.Parse()

	if structs == "" || client == "" {
		flag.Usage()
		os.Exit(1)
	}

	cnf = config{
		structs: strings.Split(structs, ","),
		client:  client,
	}
}

func main() {
	parseFlags()
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	parser := gormgen.NewParser()
	parser.ParseDir(wd)

	pkgName := parser.PkgName
	output := fmt.Sprintf("%s_gormgen.go", pkgName)

	gen := gormgen.NewGenerator(output)
	if err := gen.Init(parser, cnf.structs, cnf.client); err != nil {
		log.Fatalf("Error Initializing Generator: %v", err.Error())
	}
	if err := gen.Generate(); err != nil {
		log.Fatalf("Error Generating file: %v", err.Error())
	}
	if err := gen.Imports(); err != nil {
		log.Fatalf("Error adding imports to output file: %v", err.Error())
	}
	if err := gen.Format(); err != nil {
		log.Fatalf("Error Formating output file: %v", err.Error())
	}
	if err := gen.Flush(); err != nil {
		log.Fatalf("Error writing output file: %v", err.Error())
	}
}
