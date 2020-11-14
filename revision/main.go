package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Node struct {
	data byte
	next *Node
}

type Stack struct {
	head   *Node
	Length int
}

func (s *Stack) Push(element byte) {
	node := Node{data: element, next: nil}
	if s.head != nil {
		temp := s.head
		node.next = temp
	}
	s.head = &node
	s.Length++
}

func (s *Stack) Pop() {
	if s.head != nil {
		temp := s.head.next
		s.head = temp
		s.Length--
	}
}

// receives the byte array that will write to the revision file
func WriteToRevisionFile(b []byte) {
	file, err := os.Create("revision")
	if err == nil {
		defer file.Close()
		file.Write(b)
	}
}

func GetFilePath() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("You need to pass a file")
	}

	fileName := os.Args[1]
	workingDir, err := os.Getwd()

	if err != nil {
		return "", errors.New("Invalid working dir")
	}

	return workingDir + "/" + fileName, nil
}

func main() {
	filePath, err := GetFilePath()
	if err != nil {
		fmt.Println(err)
		return
	}

	stack := Stack{}
	lines, err := ioutil.ReadFile(filePath)

	// all bold words
	var b []byte
	if err == nil {
		// a line of bold chars
		var bcurrent []byte
		for _, char := range lines {
			if string(char) == "*" {
				stack.Push(char)
			} else {
				if stack.Length == 2 {
					bcurrent = append(bcurrent, char)
				} else if stack.Length == 4 {
					stack = Stack{nil, 0}
					bcurrent = append(bcurrent, '\n')
					b = append(b, bcurrent...)
					bcurrent = make([]byte, 0)
				} else {
					// if '*' was used to do something other than bolding a word
					stack.Pop()
				}
			}
		}

		WriteToRevisionFile(b)
	}
}
