package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type File struct {
	Name       string `json:"name"`
	Expression Expression
}

type Expression struct {
	Kind   string `json:"kind"`
	Callee struct {
		Kind string `json:"kind"`
		Text string `json:"text"`
	} `json:"callee"`
	Arguments Arguments
}

type Arguments []struct {
	Kind string `json:"kind"`
	Lhs  struct {
		Kind  string `json:"kind"`
		Value int
	} `json:"lhs"`
	Op  string `json:"op"`
	Rhs struct {
		Kind  string `json:"kind"`
		Value int
	}
	Value string `json:"value"`
}

func main() {
	var file File
	jsonFile, err := os.Open("source.rinha.json")

	if err != nil {
		println("Erro ao ler o json")
	}

	defer jsonFile.Close()

	byteJson, err := io.ReadAll(jsonFile)

	json.Unmarshal(byteJson, &file)

	interpreter(file)
}

func interpreter(file File) int {

	switch file.Expression.Callee.Text {
	case "Print":
		fmt.Print(Operations(file.Expression.Arguments))
	}

	return 0
}

func Operations(arguments Arguments) interface{} {
	switch arguments[0].Kind {
	case "Binary":
		if arguments[0].Op == "Add" {
			if arguments[0].Lhs.Kind == "Str" || arguments[0].Rhs.Kind == "Str" {
				newString := fmt.Sprintf("%v %v", arguments[0].Lhs.Value, arguments[0].Rhs.Value)
				return newString
			} else {
				var sum = arguments[0].Lhs.Value + arguments[0].Rhs.Value
				return sum
			}
		}
	case "Str":
		return arguments[0].Value
	default:
		panic(fmt.Sprintf("Unknown node kind: <%s>", arguments[0].Kind))
	}
	return 0
}
