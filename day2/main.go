package main

import (
	"errors"
	"fmt"
	"os"
)

type Document struct {
	ID        int
	Context   string
	ChunkSize int
}

// loadDocument reads a file and returns a Document.
// The (Document, error) signature is the core Go error pattern.
// In Java this would throw an IOException. In Go, error is just a return value.
func loadDocument(path string) (Document, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return Document{}, fmt.Errorf("loadDocument: %w", err)
	}

	if len(data) == 0 {
		return Document{}, errors.New("loadDocument: file is empty")
	}

	return Document{
		ID: 1,
		Context: string(data),
		ChunkSize: len(data),
	}, nil
}

func main() {
	// case 1: file done not exist
	doc, err := loadDocument("nonexistent.txt")

	if err != nil {
		fmt.Printf("Error reading the file: %v", err)
	} else {
		fmt.Printf("Document loaded successfully: %v", doc.Context)
	}

	// valid case, file exists happy path
	// create a temp file with some content
	_ = os.WriteFile("sample.txt", []byte("vector search introduction"), 0644)

	doc, err = loadDocument("sample.txt")
	if err != nil {
		fmt.Printf("Error loading document: %v", err)
		return
	} else {
		fmt.Printf("Loaded doc: ID=%d ChuckSize=%d Content=%q\n", doc.ID, doc.ChunkSize, doc.Context)
	}
}