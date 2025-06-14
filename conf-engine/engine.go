package main

import (
	"flag"
	"os"
	"strings"
	"text/template"
)

type ServiceInfo struct {
	Svc       string
	Endpoints []string
}

func main() {
	var err error

	// Input from cmdline args
	svc := flag.String("svc", "", "The service name to process with")
	endpoints  := flag.String("endpoints", "", "The endpoints to process with, separated by comma")
	templatePath := flag.String("template", "", "template file path")
	outputPath := flag.String("output-dir", "", "for rendered file")
	outputFileName := flag.String("output-file", "", "for rendered file")

	// validate flags
	flag.Parse()
	if *templatePath == "" {
		panic("-template must be provided")
	}
	if *endpoints == "" {
		panic("-endpoints must be provided")
	}
	if *svc == "" {
		panic("-svc must be provided")
	}

	// read template
	tmplContent, err := os.ReadFile(*templatePath)
	if err != nil {
		panic("Failed to read template file: " + err.Error())
	}
	tmpl := template.Must(template.New("tpl").Parse(string(tmplContent)))

	// render template
	service := ServiceInfo{
		Svc:       *svc,
		Endpoints: strings.Split(*endpoints, ","),
	}
	output := os.Stdout
	if *outputPath != "" {
		output, err = os.Create(*outputPath + "/" + *outputFileName)
		if err != nil {
			panic("Failed to create output file: " + err.Error())
		}
	}
	err = tmpl.Execute(output, service)
	if err != nil {
		panic("Failed to render template: " + err.Error())
	}
}
