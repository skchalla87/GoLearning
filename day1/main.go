package main

import "fmt"

// Document represents a chunk of content in a RAG pipeline.
type Document struct {
	ID int
	Content string
	ChunkSize int
	Score float64
	IsIndexed bool
}

func main() {

	var doc Document

	fmt.Println(doc.Content == "") // true
	fmt.Println(doc.Score == 0)
	fmt.Println("=== Zero Values ====")
	fmt.Printf("ID: %d\n", doc.ID)
	fmt.Printf("Content: %s\n", doc.Content)
	fmt.Printf("ChunkSize: %d\n", doc.ChunkSize)
	fmt.Printf("Score: %f\n", doc.Score)
	fmt.Printf("IsIndexed: %t\n", doc.IsIndexed)

	// short vairable declaration - type is inferred
	// := only works inside functions, not at package level
	title := "Introduction to Vector Search"
	chunkSize := 512

	fmt.Println("\n=== Short Variable Declaration ===")
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Chunk Size: %d\n", chunkSize)

	// Struct literal — unset fields get zero values
	indexedDoc := Document{
        ID:        1,
        Content:   title,
        ChunkSize: chunkSize,
        IsIndexed: true,
        // Score left out — zero value 0.0 is fine
    }
	fmt.Println("\n=== Struct Literal ===")
    fmt.Printf("Doc: %+v\n", indexedDoc)

    // Constants — compile-time, no runtime cost
    const MaxChunkSize = 1024
    fmt.Println("\n=== Constants ===")
    fmt.Printf("MaxChunkSize: %d\n", MaxChunkSize)

    // Multiple return values — preview of Day 2
    id, content := getDocumentFields(indexedDoc)
    fmt.Println("\n=== Multiple Return Values ===")
    fmt.Printf("id=%d content=%q\n", id, content)
}

// Lowercase = package-private. Uppercase = exported. That's the whole access model.
func getDocumentFields(d Document) (int, string) {
    return d.ID, d.Content
}