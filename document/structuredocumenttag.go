// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"bytes"

	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// StructuredDocumentTag is a tagged bit of content in a document. It wraps
// either a block-level content control (CT_SdtBlock) or a run-level / inline
// content control (CT_SdtRun); exactly one of x / run is set. The writer side
// (AddStructuredDocumentTag and the setters below) only ever produces block
// controls. Run-level controls are surfaced by the reader so documents that
// another editor (e.g. LibreOffice) has down-converted from block to inline
// controls can still be enumerated by tag.
type StructuredDocumentTag struct {
	d   *Document
	x   *wml.CT_SdtBlock // set for block-level controls
	run *wml.CT_SdtRun   // set for run-level (inline) controls
}

// X returns the inner wrapped block-level XML type, or nil for a run-level
// control.
func (s StructuredDocumentTag) X() *wml.CT_SdtBlock {
	return s.x
}

// pr returns the structured-document-tag properties for whichever control
// kind is wrapped, or nil when neither carries properties.
func (s StructuredDocumentTag) pr() *wml.CT_SdtPr {
	switch {
	case s.x != nil:
		return s.x.SdtPr
	case s.run != nil:
		return s.run.SdtPr
	}
	return nil
}

// ensurePr ensures the SdtPr (properties) struct is initialised on whichever
// control kind is wrapped and returns it. Returns nil only when neither a
// block nor a run control is set.
func (s StructuredDocumentTag) ensurePr() *wml.CT_SdtPr {
	switch {
	case s.x != nil:
		if s.x.SdtPr == nil {
			s.x.SdtPr = wml.NewCT_SdtPr()
		}
		return s.x.SdtPr
	case s.run != nil:
		if s.run.SdtPr == nil {
			s.run.SdtPr = wml.NewCT_SdtPr()
		}
		return s.run.SdtPr
	}
	return nil
}

// SetTag sets the machine-readable tag on the structured document tag. Works
// for both block-level and run-level controls.
func (s StructuredDocumentTag) SetTag(tag string) {
	pr := s.ensurePr()
	if pr == nil {
		return
	}
	if pr.Tag == nil {
		pr.Tag = wml.NewCT_String()
	}
	pr.Tag.ValAttr = tag
}

// Tag returns the machine-readable tag of the structured document tag,
// or an empty string if none is set.
func (s StructuredDocumentTag) Tag() string {
	pr := s.pr()
	if pr == nil || pr.Tag == nil {
		return ""
	}
	return pr.Tag.ValAttr
}

// SetAlias sets the human-readable title (alias) of the structured document
// tag. Works for both block-level and run-level controls.
func (s StructuredDocumentTag) SetAlias(alias string) {
	pr := s.ensurePr()
	if pr == nil {
		return
	}
	if pr.Alias == nil {
		pr.Alias = wml.NewCT_String()
	}
	pr.Alias.ValAttr = alias
}

// Alias returns the human-readable title of the structured document tag,
// or an empty string if none is set.
func (s StructuredDocumentTag) Alias() string {
	pr := s.pr()
	if pr == nil || pr.Alias == nil {
		return ""
	}
	return pr.Alias.ValAttr
}

// AddParagraph adds a new paragraph inside the structured document tag's
// content. This is a block-level operation: a run-level (inline) control holds
// runs rather than paragraphs, so calling AddParagraph on one is unsupported
// and returns a zero Paragraph rather than panicking.
func (s StructuredDocumentTag) AddParagraph() Paragraph {
	if s.x == nil {
		return Paragraph{}
	}
	if s.x.SdtContent == nil {
		s.x.SdtContent = wml.NewCT_SdtContentBlock()
	}
	p := wml.NewCT_P()
	s.x.SdtContent.P = append(s.x.SdtContent.P, p)
	return Paragraph{s.d, p}
}

// Paragraphs returns the paragraphs within a structured document tag. A
// run-level (inline) control has no paragraphs of its own and returns nil.
func (s StructuredDocumentTag) Paragraphs() []Paragraph {
	if s.x == nil || s.x.SdtContent == nil {
		return nil
	}
	ret := []Paragraph{}
	for _, p := range s.x.SdtContent.P {
		ret = append(ret, Paragraph{s.d, p})
	}
	return ret
}

// Text returns the concatenated plain text of the control's content, working
// for both block-level and run-level controls. Paragraphs within a block
// control are joined with newlines. Deleted (w:del) content is excluded;
// tracked insertions (w:ins) are included as current content.
func (s StructuredDocumentTag) Text() string {
	buf := bytes.Buffer{}
	if s.x != nil && s.x.SdtContent != nil {
		for i, p := range s.x.SdtContent.P {
			if i > 0 {
				buf.WriteByte('\n')
			}
			writeParagraphText(&buf, p)
		}
	}
	if s.run != nil && s.run.SdtContent != nil {
		for _, crc := range s.run.SdtContent.EG_ContentRunContent {
			writeRunContentText(&buf, crc)
		}
	}
	return buf.String()
}

// writeParagraphText appends the plain text of a paragraph's run content.
func writeParagraphText(buf *bytes.Buffer, p *wml.CT_P) {
	for _, pc := range p.EG_PContent {
		for _, crc := range pc.EG_ContentRunContent {
			writeRunContentText(buf, crc)
		}
	}
}

// writeRunContentText appends the text of one run-content element: the direct
// run plus any runs inside tracked insertions. Deletions are skipped.
func writeRunContentText(buf *bytes.Buffer, crc *wml.EG_ContentRunContent) {
	if crc.R != nil {
		writeRunText(buf, crc.R)
	}
	for _, rle := range crc.EG_RunLevelElts {
		if rle.Ins != nil {
			for _, irc := range rle.Ins.EG_ContentRunContent {
				if irc.R != nil {
					writeRunText(buf, irc.R)
				}
			}
		}
		// rle.Del (w:del) is deliberately skipped — deleted content is not the
		// current state of the control.
	}
}

// writeRunText appends a run's visible text (w:t and tabs), mirroring Run.Text.
func writeRunText(buf *bytes.Buffer, r *wml.CT_R) {
	for _, ic := range r.EG_RunInnerContent {
		if ic.T != nil {
			buf.WriteString(ic.T.Content)
		}
		if ic.Tab != nil {
			buf.WriteByte('\t')
		}
	}
}
