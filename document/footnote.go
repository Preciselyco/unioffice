// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// Footnote is the body of one footnote or endnote (CT_FtnEdn is shared by
// both parts in the OOXML schema, so this type serves as both).
type Footnote struct {
	d *Document
	x *wml.CT_FtnEdn
}

// X returns the inner wrapped XML type.
func (f Footnote) X() *wml.CT_FtnEdn {
	return f.x
}

// ID returns the note's id — the value a run's w:footnoteReference /
// w:endnoteReference must reference to point at this note.
func (f Footnote) ID() int64 {
	return f.x.IdAttr
}

// AddParagraph adds a paragraph to the footnote/endnote body. Mirrors
// Document.AddParagraph — CT_FtnEdn nests EG_BlockLevelElts exactly like
// CT_Body does.
func (f Footnote) AddParagraph() Paragraph {
	elts := wml.NewEG_BlockLevelElts()
	f.x.EG_BlockLevelElts = append(f.x.EG_BlockLevelElts, elts)
	c := wml.NewEG_ContentBlockContent()
	elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, c)
	p := wml.NewCT_P()
	c.P = append(c.P, p)
	return Paragraph{f.d, p}
}
