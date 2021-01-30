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

// pushes a new element to the stack
func (s *Stack) Push(element byte) {
	node := Node{data: element, next: nil}
	if s.head != nil {
		temp := s.head
		node.next = temp
	}
	s.head = &node
	s.Length++
}

// removes the last added element from the stack
func (s *Stack) Pop() {
	if s.head != nil {
		temp := s.head.next
		s.head = temp
		s.Length--
	}
}

// receives the byte slice that and write its content to a file
// with the filename given
func writeToRevisionFile(b []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(b)

	return nil
}

// return the working dir concatened to the file passed
// in the command line
func getFullPath(fileName string) (string, error) {
	//from where i am calling the program
	workingDir, err := os.Getwd()
	if err != nil {
		return "", errors.New("Invalid working dir")
	}

	return workingDir + "/" + fileName, nil
}

// returns the file name given in the command line
func getFileName() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("You need to pass the name of one file")
	}
	fileName := os.Args[1]

	return fileName, nil
}

func main() {
	fileName, err := getFileName()
	if err != nil {
		fmt.Printf("Error getting filename %v\n", err)
		return
	}

	filePath, err := getFullPath(fileName)
	if err != nil {
		fmt.Printf("Error getting working dir %v\n", err)
		return
	}

	var stack Stack
	lines, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading markdown file")
		return
	}

	//TODO modularize
	// all bold words
	var b []byte
	if err == nil {
		// a line of bold chars
		var bcurrent []byte
		for _, char := range lines {
			if char == '*' {
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

		writeToRevisionFile(b, "revision")
	}
}
