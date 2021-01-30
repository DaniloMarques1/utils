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

// returns a pointer to empty stack
func NewStack() *Stack {
	return &Stack{
		nil,
		0,
	}
}

// receives a string slice where each index must be
// written to the filename file
func writeToRevisionFile(strSlice []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	for idx, str := range strSlice {
		if idx == len(strSlice) - 1{
			// do not add new line to last line
			file.WriteString(str)
		} else {
			file.WriteString(str + "\n")
		}
	}

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

// returns a slice of strings where each index
// is a word that was marked as bold
func getBold(lines []byte) []string {
	var boldWords []string
	stack := NewStack()
	var currentWord []byte
	for _, char := range lines {
		if char == '*' {
			if stack.Length == 3 {
				boldWords = append(boldWords, string(currentWord))
				stack = NewStack()
				currentWord = []byte{}
			} else {
				stack.Push(char)
			}
		} else {
			if stack.Length == 2 {
				currentWord = append(currentWord, char)
			} else if stack.Length == 1 {
				// if the previous was a * but the current is not
				// means we are trying to bold the word
				stack = NewStack()
			}
		}
	}

	return boldWords
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
	lines, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading markdown file")
		return
	}

	boldWords := getBold(lines)
	writeToRevisionFile(boldWords, "revision")
}
