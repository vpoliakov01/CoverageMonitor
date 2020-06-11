// author: github.com/everestmz
// from: https://gist.github.com/everestmz/26181027b3edd9d1c880a40e11bc8564

package parser

import (
	"bytes"
	"strconv"
	"strings"
)

type ProjectCoverage struct {
	Files []FileCoverage `json:"files"`
}

type FileCoverage struct {
	FileName       string          `json:"file_name"`
	CoverageBlocks []CoverageBlock `json:"coverage_blocks"`
}

type CoverageBlock struct {
	StartLine int `json:"start_line"`
	EndLine   int `json:"end_line"`
}

// ParseCover consumes the contents of a coverage.out file produced
// by running `go test -coverprofile coverage.out ./...` on a package
// It produces a list of FileCoverage structs that format the coverage
// information in a structure more suitable for JSON
// Marshalling its output to JSON looks like:
// [
// 	{
// 		"file_name": "github.com/google/uuid/marshal.go",
// 		"coverage_blocks": [
// 			{
// 				"start_line": 19,
// 				"end_line": 21
// 			},
// 			...
// 		]
// 	},
// 	...
// ]
func ParseCover(coverFile []byte) ProjectCoverage {
	var files []FileCoverage
	// Go Cover File returns code blocks categorized by file, so we can
	// just have one current file
	currentFile := FileCoverage{}
	for _, line := range bytes.Split(coverFile, []byte("\n")) {
		if bytes.Contains(line, []byte("mode: ")) || len(line) == 0 {
			continue
		}
		// splits into lines formatted like:
		// <package_name>/<file_name>.go:<line_start>.<col_start>,<line_end>.<col_end> <numstatements> <count>
		// We want package/file name, line_start and line_end
		firstPart := string(bytes.Split(line, []byte(" "))[0])
		colonSplit := strings.Split(firstPart, ":")
		fileName := colonSplit[0]
		if !strings.HasSuffix(fileName, ".go") {
			// This isn't a well-formatted line
			continue
		}

		lineNumbersSplit := strings.Split(colonSplit[1], ",")
		startLine, _ := strconv.Atoi(strings.Split(lineNumbersSplit[0], ".")[0])
		endLine, _ := strconv.Atoi(strings.Split(lineNumbersSplit[1], ".")[0])

		if fileName != currentFile.FileName {
			if len(currentFile.CoverageBlocks) > 0 {
				files = append(files, currentFile)
			}
			currentFile = FileCoverage{
				FileName: fileName,
				CoverageBlocks: []CoverageBlock{{
					StartLine: startLine,
					EndLine:   endLine,
				}},
			}
		} else {
			currentFile.CoverageBlocks = append(currentFile.CoverageBlocks, CoverageBlock{
				StartLine: startLine,
				EndLine:   endLine,
			})
		}
	}

	return ProjectCoverage{
		Files: append(files, currentFile),
	}
}

// If you're not using Go, you can turn this function into a CLI by uncommenting the
// below main file, adding "package main" to the top of this file, and compiling with,
// go build -o cover2json ./parse_cover.go
// It can be used by running cover2json <path_to_cover.out>, and will output JSON to
// stdout

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Printf("command usage: %s <coverage profile file>\n", os.Args[0])
// 		os.Exit(1)
// 	}

// 	coverFile, err := ioutil.ReadFile(os.Args[1])
// 	if err != nil {
// 		fmt.Printf("could not read file: %s", err)
// 		os.Exit(1)
// 	}

// 	structured := ParseCover(coverFile)
// 	outBs, err := json.Marshal(&structured)
// 	if err != nil {
// 		fmt.Printf("could not marshal struct into json: %s", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(string(outBs))
// }
