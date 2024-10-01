package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// SourceMap represents the structure of a source map file
type SourceMap struct {
	Version        int      `json:"version"`
	File           string   `json:"file"`
	Sources        []string `json:"sources"`
	SourcesContent []string `json:"sourcesContent"`
	Mappings       string   `json:"mappings"`
}

func extractJSFromSourceMap(mapData []byte) (string, error) {
	// Parse the source map file
	var sourceMap SourceMap
	err := json.Unmarshal(mapData, &sourceMap)
	if err != nil {
		return "", fmt.Errorf("error parsing source map: %v", err)
	}

	// Check if we have sourcesContent (the original JS content)
	if len(sourceMap.SourcesContent) == 0 {
		return "", fmt.Errorf("no sourcesContent found in the source map")
	}

	// Concatenate all JavaScript sources into one string
	var fullJS string
	for _, content := range sourceMap.SourcesContent {
		fullJS += content + "\n"
	}

	return fullJS, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: extract_js <input_map_file> [output_js_file]")
		return
	}

	inputFile := os.Args[1]
	outputFile := ""
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	}

	// Read the input source map file
	mapData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading source map file: %v\n", err)
		return
	}

	// Extract the JavaScript code from the source map
	jsCode, err := extractJSFromSourceMap(mapData)
	if err != nil {
		fmt.Printf("Error extracting JS: %v\n", err)
		return
	}

	// Check if an output file was provided, otherwise print to stdout
	if outputFile != "" {
		err = ioutil.WriteFile(outputFile, []byte(jsCode), 0644)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		} else {
			fmt.Printf("JavaScript code extracted and written to %s\n", outputFile)
		}
	} else {
		// Print to stdout if no output file is provided
		fmt.Println(jsCode)
	}
}
