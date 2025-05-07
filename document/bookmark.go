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

// Bookmark is a bookmarked location within a document that can be referenced
// with a hyperlink.
type Bookmark struct {
	x  *wml.CT_Bookmark
	p  *Paragraph
	pc *wml.EG_PContent
}

// X returns the inner wrapped XML type.
func (b Bookmark) X() *wml.CT_Bookmark {
	return b.x
}

// SetName sets the name of the bookmark. This is the name that is used to
// reference the bookmark from hyperlinks.
func (b Bookmark) SetName(name string) {
	b.x.NameAttr = name
}

// Name returns the name of the bookmark whcih is the document unique ID that
// identifies the bookmark.
func (b Bookmark) Name() string {
	return b.x.NameAttr
}

// AddRun adds a run of inserted or deleted text.
func (b Bookmark) AddRun() Run {
	for i, c := range b.p.x.EG_PContent {
		if c == b.pc {
			r := wml.NewCT_R()
			for len(b.p.x.EG_PContent) <= i+2 {
				b.p.x.EG_PContent = append(b.p.x.EG_PContent, nil) // ensure space for new content
			}
			copy(b.p.x.EG_PContent[i+2:], b.p.x.EG_PContent[i+1:]) // shift existing field end
			b.p.x.EG_PContent[i+1] = wml.NewEG_PContent()
			b.p.x.EG_PContent[i+1].EG_ContentRunContent = append(b.p.x.EG_PContent[i+1].EG_ContentRunContent, nil)
			b.p.x.EG_PContent[i+1].EG_ContentRunContent[0] = wml.NewEG_ContentRunContent()
			b.p.x.EG_PContent[i+1].EG_ContentRunContent[0].R = r
			return Run{b.p.d, r}
		}
	}
	return Run{}
}
