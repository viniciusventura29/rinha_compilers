package main

import (
	"encoding/json"
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

	print(interpreter(file))
}

func interpreter(file File) int {
	if file.Expression.Arguments[0].Kind == "Binary" {
		if file.Expression.Arguments[0].Op == "Add" {
			var sum = file.Expression.Arguments[0].Lhs.Value + file.Expression.Arguments[0].Rhs.Value
			return sum
		}
	}
	return 0
}
