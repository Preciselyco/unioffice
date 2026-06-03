// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "github.com/Preciselyco/unioffice/schema/soo/wml"

// StructuredDocumentTag are a tagged bit of content in a document.
type StructuredDocumentTag struct {
	d *Document
	x *wml.CT_SdtBlock
}

// X returns the inner wrapped XML type.
func (s StructuredDocumentTag) X() *wml.CT_SdtBlock {
	return s.x
}

// ensurePr ensures that the SdtPr (properties) struct is initialised.
func (s StructuredDocumentTag) ensurePr() {
	if s.x.SdtPr == nil {
		s.x.SdtPr = wml.NewCT_SdtPr()
	}
}

// SetTag sets the machine-readable tag on the structured document tag.
func (s StructuredDocumentTag) SetTag(tag string) {
	s.ensurePr()
	if s.x.SdtPr.Tag == nil {
		s.x.SdtPr.Tag = wml.NewCT_String()
	}
	s.x.SdtPr.Tag.ValAttr = tag
}

// Tag returns the machine-readable tag of the structured document tag,
// or an empty string if none is set.
func (s StructuredDocumentTag) Tag() string {
	if s.x.SdtPr == nil || s.x.SdtPr.Tag == nil {
		return ""
	}
	return s.x.SdtPr.Tag.ValAttr
}

// SetAlias sets the human-readable title (alias) of the structured document tag.
func (s StructuredDocumentTag) SetAlias(alias string) {
	s.ensurePr()
	if s.x.SdtPr.Alias == nil {
		s.x.SdtPr.Alias = wml.NewCT_String()
	}
	s.x.SdtPr.Alias.ValAttr = alias
}

// Alias returns the human-readable title of the structured document tag,
// or an empty string if none is set.
func (s StructuredDocumentTag) Alias() string {
	if s.x.SdtPr == nil || s.x.SdtPr.Alias == nil {
		return ""
	}
	return s.x.SdtPr.Alias.ValAttr
}

// AddParagraph adds a new paragraph inside the structured document tag's content.
func (s StructuredDocumentTag) AddParagraph() Paragraph {
	if s.x.SdtContent == nil {
		s.x.SdtContent = wml.NewCT_SdtContentBlock()
	}
	p := wml.NewCT_P()
	s.x.SdtContent.P = append(s.x.SdtContent.P, p)
	return Paragraph{s.d, p}
}

// Paragraphs returns the paragraphs within a structured document tag.
func (s StructuredDocumentTag) Paragraphs() []Paragraph {
	if s.x.SdtContent == nil {
		return nil
	}
	ret := []Paragraph{}
	for _, p := range s.x.SdtContent.P {
		ret = append(ret, Paragraph{s.d, p})
	}
	return ret
}
