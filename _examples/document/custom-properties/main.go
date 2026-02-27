package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Preciselyco/unioffice/document"
)

func main() {
	doc := document.New()

	// Add some content
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText("Document with custom properties. Check File > Properties > Custom in Word/LibreOffice.")

	// Set custom properties of various types
	cp := doc.CustomProperties
	cp.SetPropertyAsString("Department", "Engineering")
	cp.SetPropertyAsString("Project", "UniOffice")
	cp.SetPropertyAsInt("Revision", 3)
	cp.SetPropertyAsBool("Approved", true)
	cp.SetPropertyAsFloat64("Score", 98.6)
	cp.SetPropertyAsDate("ReviewDate", time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC))

	// Read them back
	fmt.Println("Custom Properties:")
	for _, p := range cp.Properties() {
		fmt.Printf("  %s = %v\n", p.Name(), p.Value())
	}

	if err := doc.SaveToFile("custom-properties.docx"); err != nil {
		log.Fatalf("error saving document: %s", err)
	}
	fmt.Println("\nSaved custom-properties.docx")
}
