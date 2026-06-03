// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document_test

import (
	"bytes"
	"testing"

	"github.com/Preciselyco/unioffice/document"
)

// TestSDTCreate verifies that AddStructuredDocumentTag + SetTag/SetAlias +
// AddParagraph produces a document body containing one block SDT with the
// expected tag, alias, and an inner paragraph with text.
func TestSDTCreate(t *testing.T) {
	doc := document.New()

	sdt := doc.AddStructuredDocumentTag()
	sdt.SetTag("clause-1")
	sdt.SetAlias("Introduction")

	p := sdt.AddParagraph()
	run := p.AddRun()
	run.AddText("Hello, world.")

	sdts := doc.StructuredDocumentTags()
	if len(sdts) != 1 {
		t.Fatalf("expected 1 SDT, got %d", len(sdts))
	}

	got := sdts[0]
	if got.Tag() != "clause-1" {
		t.Errorf("expected tag 'clause-1', got %q", got.Tag())
	}
	if got.Alias() != "Introduction" {
		t.Errorf("expected alias 'Introduction', got %q", got.Alias())
	}

	paras := got.Paragraphs()
	if len(paras) != 1 {
		t.Fatalf("expected 1 paragraph inside SDT, got %d", len(paras))
	}

	runs := paras[0].Runs()
	if len(runs) != 1 {
		t.Fatalf("expected 1 run inside SDT paragraph, got %d", len(runs))
	}
	if runs[0].Text() != "Hello, world." {
		t.Errorf("expected run text 'Hello, world.', got %q", runs[0].Text())
	}
}

// TestSDTReadBack verifies that StructuredDocumentTags() enumerates SDTs and
// Tag()/Alias()/Paragraphs() return the values that were set.
func TestSDTReadBack(t *testing.T) {
	doc := document.New()

	s1 := doc.AddStructuredDocumentTag()
	s1.SetTag("tag-a")
	s1.SetAlias("Section A")
	p1 := s1.AddParagraph()
	p1.AddRun().AddText("First")

	s2 := doc.AddStructuredDocumentTag()
	s2.SetTag("tag-b")
	s2.SetAlias("Section B")
	p2 := s2.AddParagraph()
	p2.AddRun().AddText("Second")

	sdts := doc.StructuredDocumentTags()
	if len(sdts) != 2 {
		t.Fatalf("expected 2 SDTs, got %d", len(sdts))
	}

	for i, tc := range []struct {
		tag, alias, text string
	}{
		{"tag-a", "Section A", "First"},
		{"tag-b", "Section B", "Second"},
	} {
		s := sdts[i]
		if s.Tag() != tc.tag {
			t.Errorf("SDT[%d]: expected tag %q, got %q", i, tc.tag, s.Tag())
		}
		if s.Alias() != tc.alias {
			t.Errorf("SDT[%d]: expected alias %q, got %q", i, tc.alias, s.Alias())
		}
		paras := s.Paragraphs()
		if len(paras) != 1 {
			t.Fatalf("SDT[%d]: expected 1 paragraph, got %d", i, len(paras))
		}
		runs := paras[0].Runs()
		if len(runs) != 1 {
			t.Fatalf("SDT[%d]: expected 1 run, got %d", i, len(runs))
		}
		if runs[0].Text() != tc.text {
			t.Errorf("SDT[%d]: expected text %q, got %q", i, tc.text, runs[0].Text())
		}
	}
}

// TestSDTRoundTrip builds a document with a tagged SDT, saves it to a buffer,
// reads it back, and asserts the SDT, its tag, alias, and paragraph text survive.
func TestSDTRoundTrip(t *testing.T) {
	doc := document.New()

	sdt := doc.AddStructuredDocumentTag()
	sdt.SetTag("rt-tag")
	sdt.SetAlias("Round-Trip Alias")
	p := sdt.AddParagraph()
	p.AddRun().AddText("round-trip text")

	buf := bytes.Buffer{}
	if err := doc.Save(&buf); err != nil {
		t.Fatalf("Save failed: %s", err)
	}

	doc2, err := document.ReadFromBytes(buf.Bytes())
	if err != nil {
		t.Fatalf("ReadFromBytes failed: %s", err)
	}

	sdts := doc2.StructuredDocumentTags()
	if len(sdts) != 1 {
		t.Fatalf("after round-trip: expected 1 SDT, got %d", len(sdts))
	}

	got := sdts[0]
	if got.Tag() != "rt-tag" {
		t.Errorf("after round-trip: expected tag 'rt-tag', got %q", got.Tag())
	}
	if got.Alias() != "Round-Trip Alias" {
		t.Errorf("after round-trip: expected alias 'Round-Trip Alias', got %q", got.Alias())
	}

	paras := got.Paragraphs()
	if len(paras) != 1 {
		t.Fatalf("after round-trip: expected 1 paragraph, got %d", len(paras))
	}
	runs := paras[0].Runs()
	if len(runs) != 1 {
		t.Fatalf("after round-trip: expected 1 run, got %d", len(runs))
	}
	if runs[0].Text() != "round-trip text" {
		t.Errorf("after round-trip: expected text 'round-trip text', got %q", runs[0].Text())
	}
}

// TestSDTNoLock verifies that AddStructuredDocumentTag does not set a lock
// (controls remain freely editable).
func TestSDTNoLock(t *testing.T) {
	doc := document.New()
	sdt := doc.AddStructuredDocumentTag()
	sdt.SetTag("unlocked")

	sdts := doc.StructuredDocumentTags()
	if len(sdts) != 1 {
		t.Fatalf("expected 1 SDT, got %d", len(sdts))
	}

	x := sdts[0].X()
	if x.SdtPr != nil && x.SdtPr.Lock != nil {
		t.Errorf("expected no lock on SDT, but Lock was set: %+v", x.SdtPr.Lock)
	}
}
