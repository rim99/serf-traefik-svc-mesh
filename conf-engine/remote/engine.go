package main

import (
	"encoding/json"
	"flag"
	"maps"
	"os"
	"slices"
	"strings"
	"text/template"
)

type ServiceInfo struct {
	Svc       string
	Endpoints []string
}

type Member struct {
	Name    string            `json:"name"`
	Address string            `json:"addr"`
	Status  string            `json:"status"`
	Tags    map[string]string `json:"tags"`
}

type SerfOutput struct {
	Members []Member `json:"members"`
}

func main() {
	var err error

	// Input from cmdline args
	services := flag.String("services", "", "Services to connected with, comma separated")
	jsonFileInput := flag.String("json-file", "", "The JSON file of serf members output")
	templatePath := flag.String("template", "", "template file path")
	outputPath := flag.String("output-dir", "", "for rendered file")

	// validate flags
	flag.Parse()
	if *templatePath == "" {
		panic("-template must be provided")
	}

	// read template
	tmplContent, err := os.ReadFile(*templatePath)
	if err != nil {
		panic("Failed to read template file: " + err.Error())
	}

	// Parse serf members
	input := os.Stdin
	if *jsonFileInput != "" {
		input, err = os.Open(*jsonFileInput)
		if err != nil {
			panic("Failed to open serf members json file: " + err.Error())
		}
	}
	defer input.Close()
	var serfOutput SerfOutput
	decoder := json.NewDecoder(input)
	if err := decoder.Decode(&serfOutput); err != nil {
		panic("Failed to decode JSON input: " + err.Error())
	}

	// Turn "VM -> svc" pairs into "svc -> VM" pairs
	requiredRemoteServices := strings.Split(*services, ",")
	var serviceInfoList = make(map[string]ServiceInfo)
	for _, member := range serfOutput.Members {
		hostedServices := strings.Split(member.Tags["services"], ",")
		for _, svc := range hostedServices {
			if slices.Contains(requiredRemoteServices, svc) {
				keys := slices.Sorted(maps.Keys(serviceInfoList))
				if !slices.Contains(keys, svc) {
					serviceInfoList[svc] = ServiceInfo{
						Svc:       svc,
						Endpoints: []string{},
					}
				}
				info := serviceInfoList[svc]
				info.Endpoints = append(info.Endpoints, strings.Split(member.Address, ":")[0])
				serviceInfoList[svc] = info
			}
		}
	}

	// Render templates
	for svc, info := range serviceInfoList {
		tmpl := template.Must(template.New("tpl").Parse(string(tmplContent)))
		// render template
		output := os.Stdout // Dry-run
		if *outputPath != "" {
			output, err = os.Create(*outputPath + "/remote-" + svc + ".yaml")
			if err != nil {
				panic("Failed to create output file: " + err.Error())
			}
		}
		err = tmpl.Execute(output, info)
		if err != nil {
			panic("Failed to render template: " + err.Error())
		}
	}
}
