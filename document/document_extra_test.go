// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/Preciselyco/unioffice/document"
)

func TestDocumentX(t *testing.T) {
	doc := document.New()
	if doc.X() == nil {
		t.Error("X() must not return nil")
	}
}

func TestDocumentBodySection(t *testing.T) {
	doc := document.New()
	s := doc.BodySection()
	if s.X() == nil {
		t.Error("BodySection().X() must not return nil")
	}
}

func TestDocumentRemoveParagraph(t *testing.T) {
	doc := document.New()
	p1 := doc.AddParagraph()
	p1.AddRun().AddText("first")
	p2 := doc.AddParagraph()
	p2.AddRun().AddText("second")
	p3 := doc.AddParagraph()
	p3.AddRun().AddText("third")

	before := len(doc.Paragraphs())
	doc.RemoveParagraph(p2)
	after := len(doc.Paragraphs())
	if after != before-1 {
		t.Errorf("RemoveParagraph: got %d paragraphs, want %d", after, before-1)
	}
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			if r.Text() == "second" {
				t.Error("removed paragraph text 'second' still found")
			}
		}
	}
}

func TestDocumentRemoveParagraphNoop(t *testing.T) {
	doc := document.New()
	p := doc.AddParagraph()

	other := document.New()
	orphan := other.AddParagraph()

	before := len(doc.Paragraphs())
	doc.RemoveParagraph(orphan)
	after := len(doc.Paragraphs())
	if after != before {
		t.Error("RemoveParagraph of orphan paragraph should be a no-op")
	}
	_ = p
}

func TestDocumentSaveToFile(t *testing.T) {
	doc := document.New()
	doc.AddParagraph().AddRun().AddText("save test")

	path := filepath.Join(t.TempDir(), "out.docx")
	if err := doc.SaveToFile(path); err != nil {
		t.Fatalf("SaveToFile: %s", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("saved file not found: %s", err)
	}

	// round-trip: reopen and validate
	doc2, err := document.Open(path)
	if err != nil {
		t.Fatalf("Open after SaveToFile: %s", err)
	}
	if err := doc2.Validate(); err != nil {
		t.Errorf("Validate after SaveToFile round-trip: %s", err)
	}
}

func TestDocumentOpenTemplate(t *testing.T) {
	// Use an existing docx as a "template"
	doc := document.New()
	doc.AddParagraph().AddRun().AddText("template content")
	buf := &bytes.Buffer{}
	if err := doc.Save(buf); err != nil {
		t.Fatalf("Save: %s", err)
	}
	path := filepath.Join(t.TempDir(), "tpl.docx")
	if err := os.WriteFile(path, buf.Bytes(), 0644); err != nil {
		t.Fatalf("WriteFile: %s", err)
	}

	tpl, err := document.OpenTemplate(path)
	if err != nil {
		t.Fatalf("OpenTemplate: %s", err)
	}
	if err := tpl.Validate(); err != nil {
		t.Errorf("Validate: %s", err)
	}
}

func TestDocumentClose(t *testing.T) {
	doc := document.New()
	// Close should not panic on a freshly created document
	if err := doc.Close(); err != nil {
		t.Errorf("Close() returned error: %s", err)
	}
}

func TestDocumentStructuredDocumentTags(t *testing.T) {
	doc := document.New()
	tags := doc.StructuredDocumentTags()
	if tags == nil {
		t.Error("StructuredDocumentTags() should return a non-nil slice")
	}
}
