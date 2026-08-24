// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document_test

import (
	"archive/zip"
	"bytes"
	"io"
	"testing"

	"github.com/Preciselyco/unioffice/document"
	"github.com/Preciselyco/unioffice/schema/soo/ofc/sharedTypes"
)

// zipEntry returns the raw bytes of name from a saved DOCX, or "" if absent.
func zipEntry(t *testing.T, docx []byte, name string) (string, bool) {
	t.Helper()
	zr, err := zip.NewReader(bytes.NewReader(docx), int64(len(docx)))
	if err != nil {
		t.Fatalf("zip.NewReader: %v", err)
	}
	for _, f := range zr.File {
		if f.Name != name {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			t.Fatalf("open %s: %v", name, err)
		}
		defer func() { _ = rc.Close() }()
		data, err := io.ReadAll(rc)
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		return string(data), true
	}
	return "", false
}

// TestNoFootnotesNoOrphanPart asserts a document built with no footnotes
// never emits an orphan word/footnotes.xml (or a dangling content-type
// override / relationship for one).
func TestNoFootnotesNoOrphanPart(t *testing.T) {
	doc := document.New()
	var buf bytes.Buffer
	if err := doc.Save(&buf); err != nil {
		t.Fatalf("Save: %v", err)
	}
	if _, ok := zipEntry(t, buf.Bytes(), "word/footnotes.xml"); ok {
		t.Error("word/footnotes.xml present with no footnotes added")
	}
	ct, _ := zipEntry(t, buf.Bytes(), "[Content_Types].xml")
	if bytes.Contains([]byte(ct), []byte("footnotes")) {
		t.Error("[Content_Types].xml references footnotes with none added")
	}
}

// TestEnsureFootnotesProducesValidPart is the fork-side proof this whole
// feature depends on: a document with no footnotes part on read must gain
// one on Save that Word will actually open — the content-type override and
// the document.xml.rels relationship, not just the zip entry.
func TestEnsureFootnotesProducesValidPart(t *testing.T) {
	doc := document.New()
	fns := doc.EnsureFootnotes()
	fn := fns.AddFootnote()
	body := fn.AddParagraph()
	body.AddRun().AddFootnoteRef()
	body.AddRun().AddText(" See clause 3.")

	p := doc.AddParagraph()
	ref := p.AddRun()
	ref.Properties().SetVerticalAlignment(sharedTypes.ST_VerticalAlignRunSuperscript)
	ref.AddFootnoteReference(fn)

	var buf bytes.Buffer
	if err := doc.Save(&buf); err != nil {
		t.Fatalf("Save: %v", err)
	}
	out := buf.Bytes()

	partXML, ok := zipEntry(t, out, "word/footnotes.xml")
	if !ok {
		t.Fatal("word/footnotes.xml missing after EnsureFootnotes+AddFootnote")
	}
	if !bytes.Contains([]byte(partXML), []byte(`w:id="1"`)) {
		t.Errorf("footnotes.xml has no id=1 footnote:\n%s", partXML)
	}
	if !bytes.Contains([]byte(partXML), []byte(`w:type="separator"`)) ||
		!bytes.Contains([]byte(partXML), []byte(`w:type="continuationSeparator"`)) {
		t.Errorf("footnotes.xml missing separator marks:\n%s", partXML)
	}

	ct, _ := zipEntry(t, out, "[Content_Types].xml")
	if !bytes.Contains([]byte(ct), []byte(`/word/footnotes.xml`)) {
		t.Errorf("[Content_Types].xml missing footnotes override:\n%s", ct)
	}

	rels, ok := zipEntry(t, out, "word/_rels/document.xml.rels")
	if !ok {
		t.Fatal("word/_rels/document.xml.rels missing")
	}
	if !bytes.Contains([]byte(rels), []byte("footnotes.xml")) {
		t.Errorf("document.xml.rels missing a footnotes.xml relationship:\n%s", rels)
	}

	docXML, _ := zipEntry(t, out, "word/document.xml")
	if !bytes.Contains([]byte(docXML), []byte(`w:footnoteReference w:id="1"`)) {
		t.Errorf("document.xml missing the footnote reference:\n%s", docXML)
	}
}

// TestClearFootnotesDropsThePart asserts ClearFootnotes fully removes the
// part (and its registration), not merely empties it — needed so Compose can
// reset a reused template's stale footnotes before rebuilding fresh ones.
func TestClearFootnotesDropsThePart(t *testing.T) {
	doc := document.New()
	fns := doc.EnsureFootnotes()
	fn := fns.AddFootnote()
	fn.AddParagraph().AddRun().AddText("stale")

	doc.ClearFootnotes()

	var buf bytes.Buffer
	if err := doc.Save(&buf); err != nil {
		t.Fatalf("Save: %v", err)
	}
	if _, ok := zipEntry(t, buf.Bytes(), "word/footnotes.xml"); ok {
		t.Error("word/footnotes.xml present after ClearFootnotes")
	}
}

// TestFootnoteIDsAvoidSeparatorSlots asserts AddFootnote allocates ids
// starting at 1, never colliding with the reserved separator ids (-1, 0).
func TestFootnoteIDsAvoidSeparatorSlots(t *testing.T) {
	doc := document.New()
	fns := doc.EnsureFootnotes()
	first := fns.AddFootnote()
	second := fns.AddFootnote()
	if first.ID() != 1 {
		t.Errorf("first footnote id = %d, want 1", first.ID())
	}
	if second.ID() != 2 {
		t.Errorf("second footnote id = %d, want 2", second.ID())
	}
}

// TestEndnotesSymmetric is a light smoke test that Endnotes follows the same
// shape as Footnotes.
func TestEndnotesSymmetric(t *testing.T) {
	doc := document.New()
	ens := doc.EnsureEndnotes()
	en := ens.AddEndnote()
	if en.ID() != 1 {
		t.Errorf("first endnote id = %d, want 1", en.ID())
	}
	en.AddParagraph().AddRun().AddText("end note body")

	var buf bytes.Buffer
	if err := doc.Save(&buf); err != nil {
		t.Fatalf("Save: %v", err)
	}
	if _, ok := zipEntry(t, buf.Bytes(), "word/endnotes.xml"); !ok {
		t.Error("word/endnotes.xml missing after EnsureEndnotes+AddEndnote")
	}
}

// TestFootnotesAddFootnoteWithoutEnsureLazilyCreatesPart asserts that calling
// AddFootnote() via Document.Footnotes() — which wraps nil when the document
// has no footnotes part yet, rather than via EnsureFootnotes() — does not
// panic, and instead lazily creates and registers the part.
func TestFootnotesAddFootnoteWithoutEnsureLazilyCreatesPart(t *testing.T) {
	doc := document.New()
	fn := doc.Footnotes().AddFootnote()
	if fn.ID() != 1 {
		t.Errorf("footnote id = %d, want 1", fn.ID())
	}
	if !doc.Footnotes().NonEmpty() {
		t.Error("document footnotes should be non-empty after lazy AddFootnote")
	}

	var buf bytes.Buffer
	if err := doc.Save(&buf); err != nil {
		t.Fatalf("Save: %v", err)
	}
	if _, ok := zipEntry(t, buf.Bytes(), "word/footnotes.xml"); !ok {
		t.Error("word/footnotes.xml missing after lazy AddFootnote")
	}
}

// TestEndnotesAddEndnoteWithoutEnsureLazilyCreatesPart is
// TestFootnotesAddFootnoteWithoutEnsureLazilyCreatesPart's endnote twin.
func TestEndnotesAddEndnoteWithoutEnsureLazilyCreatesPart(t *testing.T) {
	doc := document.New()
	en := doc.Endnotes().AddEndnote()
	if en.ID() != 1 {
		t.Errorf("endnote id = %d, want 1", en.ID())
	}
	if !doc.Endnotes().NonEmpty() {
		t.Error("document endnotes should be non-empty after lazy AddEndnote")
	}

	var buf bytes.Buffer
	if err := doc.Save(&buf); err != nil {
		t.Fatalf("Save: %v", err)
	}
	if _, ok := zipEntry(t, buf.Bytes(), "word/endnotes.xml"); !ok {
		t.Error("word/endnotes.xml missing after lazy AddEndnote")
	}
}

// TestAddEndnoteReferenceWritesEndnoteReference is AddFootnoteReference's
// endnote-side round-trip check (see TestEnsureFootnotesProducesValidPart) —
// exercises Run.AddEndnoteReference, which had no direct test coverage.
func TestAddEndnoteReferenceWritesEndnoteReference(t *testing.T) {
	doc := document.New()
	ens := doc.EnsureEndnotes()
	en := ens.AddEndnote()
	en.AddParagraph().AddRun().AddText("end note body")

	p := doc.AddParagraph()
	p.AddRun().AddEndnoteReference(en)

	var buf bytes.Buffer
	if err := doc.Save(&buf); err != nil {
		t.Fatalf("Save: %v", err)
	}

	docXML, _ := zipEntry(t, buf.Bytes(), "word/document.xml")
	if !bytes.Contains([]byte(docXML), []byte(`w:endnoteReference w:id="1"`)) {
		t.Errorf("document.xml missing the endnote reference:\n%s", docXML)
	}
}
