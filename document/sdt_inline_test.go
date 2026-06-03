// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document_test

import (
	"testing"

	"github.com/Preciselyco/unioffice/document"
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// TestStructuredDocumentTags_findsInlineRunLevelControls exercises a real
// exported document whose content controls are RUN-LEVEL (inline, plain-text)
// SDTs nested inside paragraphs — the shape another editor (LibreOffice)
// produces when it down-converts our block controls. The document contains 8
// such controls (8 <w:sdtContent><w:r> with <w:text/> markers in
// word/document.xml). StructuredDocumentTags must enumerate them by tag and
// expose their text.
func TestStructuredDocumentTags_findsInlineRunLevelControls(t *testing.T) {
	doc, err := document.Open("testdata/inline-sdt.docx")
	if err != nil {
		t.Fatalf("open fixture: %s", err)
	}

	sdts := doc.StructuredDocumentTags()
	if len(sdts) != 8 {
		t.Fatalf("expected 8 content controls, got %d — run-level (inline) SDTs nested inside paragraphs were not enumerated", len(sdts))
	}

	for i, s := range sdts {
		if s.Tag() == "" {
			t.Errorf("control[%d]: empty tag", i)
		}
		if s.Text() == "" {
			t.Errorf("control[%d] (tag %q): empty Text()", i, s.Tag())
		}
	}
}

// TestStructuredDocumentTags_synthRunLevel builds a run-level (inline) content
// control by hand and asserts the enumerator finds it and reads its tag/text,
// so coverage of the run-level path does not depend solely on the binary
// fixture. The writer API only produces block controls, so the inline control
// is assembled directly from the wml types.
func TestStructuredDocumentTags_synthRunLevel(t *testing.T) {
	doc := document.New()
	p := doc.AddParagraph()

	run := wml.NewCT_SdtRun()
	run.SdtPr = wml.NewCT_SdtPr()
	run.SdtPr.Tag = wml.NewCT_String()
	run.SdtPr.Tag.ValAttr = "pbi_RUN"
	run.SdtPr.Alias = wml.NewCT_String()
	run.SdtPr.Alias.ValAttr = "Inline Clause"
	run.SdtContent = wml.NewCT_SdtContentRun()

	r := wml.NewCT_R()
	ic := wml.NewEG_RunInnerContent()
	ic.T = wml.NewCT_Text()
	ic.T.Content = "Inline heading text"
	r.EG_RunInnerContent = append(r.EG_RunInnerContent, ic)

	inner := wml.NewEG_ContentRunContent()
	inner.R = r
	run.SdtContent.EG_ContentRunContent = append(run.SdtContent.EG_ContentRunContent, inner)

	// Attach the run-level control to the paragraph's content.
	pc := wml.NewEG_PContent()
	crc := wml.NewEG_ContentRunContent()
	crc.Sdt = run
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, crc)
	p.X().EG_PContent = append(p.X().EG_PContent, pc)

	sdts := doc.StructuredDocumentTags()
	if len(sdts) != 1 {
		t.Fatalf("expected 1 run-level control, got %d", len(sdts))
	}
	got := sdts[0]
	if got.Tag() != "pbi_RUN" {
		t.Errorf("tag: got %q, want %q", got.Tag(), "pbi_RUN")
	}
	if got.Alias() != "Inline Clause" {
		t.Errorf("alias: got %q, want %q", got.Alias(), "Inline Clause")
	}
	if got.Text() != "Inline heading text" {
		t.Errorf("text: got %q, want %q", got.Text(), "Inline heading text")
	}
	if got.X() != nil {
		t.Errorf("run-level control should have nil X() (no block wrapper)")
	}
	if got.Paragraphs() != nil {
		t.Errorf("run-level control should have no paragraphs")
	}
}

// runLevelControl builds a paragraph containing a single run-level (inline)
// content control tagged with tag, and returns the document.
func runLevelControlDoc(tag string) *document.Document {
	doc := document.New()
	p := doc.AddParagraph()
	run := wml.NewCT_SdtRun()
	run.SdtPr = wml.NewCT_SdtPr()
	run.SdtPr.Tag = wml.NewCT_String()
	run.SdtPr.Tag.ValAttr = tag
	run.SdtContent = wml.NewCT_SdtContentRun()
	pc := wml.NewEG_PContent()
	crc := wml.NewEG_ContentRunContent()
	crc.Sdt = run
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, crc)
	p.X().EG_PContent = append(p.X().EG_PContent, pc)
	return doc
}

// TestStructuredDocumentTag_writerMethodsSafeOnRunLevel checks that the writer
// methods do not panic when invoked on a run-level control returned by the
// reader (where the block pointer is nil): SetTag/SetAlias mutate the run
// control, and AddParagraph is a safe no-op returning a zero Paragraph.
func TestStructuredDocumentTag_writerMethodsSafeOnRunLevel(t *testing.T) {
	doc := runLevelControlDoc("run-orig")
	sdts := doc.StructuredDocumentTags()
	if len(sdts) != 1 {
		t.Fatalf("expected 1 run-level control, got %d", len(sdts))
	}
	s := sdts[0]

	s.SetTag("run-new") // must not panic on a run-level control
	s.SetAlias("Run Alias")
	if s.Tag() != "run-new" {
		t.Errorf("SetTag on run-level: Tag() = %q, want run-new", s.Tag())
	}
	if s.Alias() != "Run Alias" {
		t.Errorf("SetAlias on run-level: Alias() = %q, want Run Alias", s.Alias())
	}

	if got := s.AddParagraph(); got.X() != nil { // must not panic; zero Paragraph
		t.Errorf("AddParagraph on run-level should return a zero Paragraph")
	}
}

// TestStructuredDocumentTags_recursesNestedAndTables verifies enumeration
// reaches controls nested inside a block SDT's content: a nested block SDT and
// a run-level control inside one of the SDT's content paragraphs.
func TestStructuredDocumentTags_recursesNestedAndTables(t *testing.T) {
	doc := document.New()
	outer := doc.AddStructuredDocumentTag()
	outer.SetTag("outer")
	outer.AddParagraph().AddRun().AddText("outer body")

	// Nest a block SDT inside the outer control's content.
	nested := wml.NewCT_SdtBlock()
	nested.SdtPr = wml.NewCT_SdtPr()
	nested.SdtPr.Tag = wml.NewCT_String()
	nested.SdtPr.Tag.ValAttr = "nested"
	outer.X().SdtContent.Sdt = nested

	// Nest a run-level control inside one of the outer control's paragraphs.
	innerPara := outer.X().SdtContent.P[0]
	run := wml.NewCT_SdtRun()
	run.SdtPr = wml.NewCT_SdtPr()
	run.SdtPr.Tag = wml.NewCT_String()
	run.SdtPr.Tag.ValAttr = "inner-run"
	pc := wml.NewEG_PContent()
	crc := wml.NewEG_ContentRunContent()
	crc.Sdt = run
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, crc)
	innerPara.EG_PContent = append(innerPara.EG_PContent, pc)

	tags := map[string]bool{}
	for _, s := range doc.StructuredDocumentTags() {
		tags[s.Tag()] = true
	}
	for _, want := range []string{"outer", "nested", "inner-run"} {
		if !tags[want] {
			t.Errorf("StructuredDocumentTags missing nested control %q; found %v", want, tags)
		}
	}
}
