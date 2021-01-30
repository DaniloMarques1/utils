package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := Stack{}
	stack.Push('a')
	stack.Push('b')
	if stack.Length != 2 || stack.head.data != 'b' {
		t.Fatalf("Failed to push to stack\n")
	}
}

func TestStackPop(t *testing.T) {
	stack := Stack{}
	stack.Push('a')
	stack.Push('b')
	stack.Push('c')
	stack.Pop()
	if stack.Length != 2 || stack.head.data != 'b' {
		t.Fatalf("Failed to pop from stack\n")
	}
}

func TestGetFullPath(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting working dir %v\n", err)
	}

	expect := workingDir + "/filename"
	result, err := getFullPath("filename")
	if err != nil {
		t.Errorf("Error getting full path %v\n", err)
	}

	if result != expect {
		t.Errorf("Returned wrong path to file\n")
	}
}

func TestWriteToRevisionFile(t *testing.T) {
	content := "Danilo Marques\nEh top"
	filename := "test"
	if err := writeToRevisionFile([]byte(content), filename); err != nil {
		t.Errorf("Error writing to file %v\n", err)
	}
	defer os.Remove(filename)

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Error reading file %v\n", err)
	}
	if string(fileContent) != content {
		t.Errorf("Wrong file content\n")
	}
}
