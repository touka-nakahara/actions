package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewEqual(t *testing.T) {
	if errors.New("abc") == errors.New("abc") {
		t.Errorf(`New("abc") == New("abc")`)
	}
}

func TestNewEqual2(t *testing.T) {
	if errors.New("abc") == errors.New("abc") {
		t.Errorf(`New("abc") == New("abc")`)
	}
}

func TestNewEqual3(t *testing.T) {
	if errors.New("abc") == errors.New("abc") {
		t.Errorf(`New("abc") == New("abc")`)
	}
}

func main() {
	fmt.Println("GitHubActions")
	fmt.Print("Heelo GitHub Aciotns\n")
	fmt.Print("Heelo GitHub Aciotns\n")
}
