package main

import (
	"os"
	"testing"
)

func TestPush(t *testing.T) {
	stack := Stack{}
	stack.Push('a')
	stack.Push('b')
	if stack.Length != 2 || stack.head.data != 'b' {
		t.Fatalf("Failed to push to stack\n")
	}
}

func TestPop(t *testing.T) {
	stack := Stack{}
	stack.Push('a')
	stack.Push('b')
	stack.Push('c')
	stack.Pop()
	if stack.Length != 2 || stack.head.data != 'b' {
		t.Fatalf("Failed to pop from stack\n")
	}
}

func TestWriteToRevisionFile(t *testing.T) {
	s := "Danilo eh top"
	file, err := os.Create("test_revision")
	if err != nil {
		t.Fatal("Failed to write to revision file")
	}
	defer file.Close()
	defer os.Remove("test_revision")
	_, err = file.WriteString(s)
	if err != nil {
		t.Fatal("Failed to write to revision file")
	}
}
