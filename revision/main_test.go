package main

import (
	"io/ioutil"
	"os"
	"strings"
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
	content := []string{"Danilo Marques", "Eh top"}
	//content := "Danilo Marques\nEh top"
	filename := "test"
	if err := writeToRevisionFile(content, filename); err != nil {
		t.Errorf("Error writing to file %v\n", err)
	}
	defer os.Remove(filename)

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Error reading file %v\n", err)
	}
	contentJoined := strings.Join(content, "\n")
	if string(fileContent) != contentJoined {
		t.Errorf("Wrong file content\n")
	}
}

func TestGetBold(t *testing.T) {
	content := `
	Ticking away the moments that make up a dull day
	**Fritter and waste** the hours in an offhand way.
	Kicking around on a piece of ground in your home town
	Waiting for someone or **something to** show you the way.

	Tired of lying in the sunshine staying **home** to watch the rain.
	You are young and life is long and there is time to kill today.
	And then one day you find ten years have got behind you.
	No one told you when to run, you missed the starting gun.

	So you run and you run to catch up with the sun but it's sinking
	Racing around to come up behind you again.
	The sun is the same in a relative way but you're older,
	Shorter of breath and one day closer to death.

	Every year is **getting shorter** never seem to find the time.
	Plans that either come to naught or half a page of scribbled lines
	Hanging on in quiet desperation is the English way
	**The time is gone** the song is over,
	Thought I'd something more to say.
	`
	boldWords := getBold([]byte(content))
	if len(boldWords) != 5 {
		t.Errorf("Wrong length of bold words\n")
	}

	if boldWords[0] != "Fritter and waste" {
		t.Errorf("Wrong bold word. Expect 'Fritter and waste' got '%v'\n", boldWords[0])
	}

	if boldWords[1] != "something to" {
		t.Errorf("Wrong bold word. Expect 'something to' got '%v'\n", boldWords[1])
	}

	if boldWords[2] != "home" {
		t.Errorf("Wrong bold word. Expect 'home' got '%v'\n", boldWords[1])
	}

	if boldWords[3] != "getting shorter" {
		t.Errorf("Wrong bold word. Expect 'getting shorter' got '%v'\n", boldWords[2])
	}

	if boldWords[4] != "The time is gone" {
		t.Errorf("Wrong bold word. Expect 'The time is gone' got '%v'\n", boldWords[2])
	}

	content = `
		**this will** be a math problem to you son
		3 * 2 is equal to what?

		what about 7 * 2? do you have a clue

		You better **have it**
	`
	boldWords = getBold([]byte(content))
	if len(boldWords) != 2 {
		t.Errorf("Wrong length of bold words\n")
	}

	if boldWords[0] != "this will" {
		t.Errorf("Wrong bold word. Expect 'this will' got '%v'\n", boldWords[0])
	}

	if boldWords[1] != "have it" {
		t.Errorf("Wrong bold word. Expect 'have it' got '%v'\n", boldWords[1])
	}
}
