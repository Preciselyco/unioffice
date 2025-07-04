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

// Paragraph is a paragraph within a document.
type Paragraph struct {
	d *Document
	x *wml.CT_P
}

// X returns the inner wrapped XML type.
func (p Paragraph) X() *wml.CT_P {
	return p.x
}

func (p Paragraph) ensurePPr() {
	if p.x.PPr == nil {
		p.x.PPr = wml.NewCT_PPr()
	}
}

// RemoveRun removes a child run from a paragraph.
func (p Paragraph) RemoveRun(r Run) {
	for _, c := range p.x.EG_PContent {
		for i, rc := range c.EG_ContentRunContent {
			if rc.R == r.x {
				copy(c.EG_ContentRunContent[i:], c.EG_ContentRunContent[i+1:])
				c.EG_ContentRunContent = c.EG_ContentRunContent[0 : len(c.EG_ContentRunContent)-1]
			}
			if rc.Sdt != nil && rc.Sdt.SdtContent != nil {
				for i, rc2 := range rc.Sdt.SdtContent.EG_ContentRunContent {
					if rc2.R == r.x {
						copy(rc.Sdt.SdtContent.EG_ContentRunContent[i:], rc.Sdt.SdtContent.EG_ContentRunContent[i+1:])
						rc.Sdt.SdtContent.EG_ContentRunContent = rc.Sdt.SdtContent.EG_ContentRunContent[0 : len(rc.Sdt.SdtContent.EG_ContentRunContent)-1]
					}
				}
			}
		}
	}
}

// Properties returns the paragraph properties.
func (p Paragraph) Properties() ParagraphProperties {
	p.ensurePPr()
	return ParagraphProperties{p.d, p.x.PPr}
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (p Paragraph) Style() string {
	if p.x.PPr != nil && p.x.PPr.PStyle != nil {
		return p.x.PPr.PStyle.ValAttr
	}
	return ""
}

// SetStyle sets the style of a paragraph and is identical to setting it on the
// paragraph's Properties()
func (p Paragraph) SetStyle(s string) {
	p.ensurePPr()
	if s == "" {
		p.x.PPr.PStyle = nil
	} else {
		p.x.PPr.PStyle = wml.NewCT_String()
		p.x.PPr.PStyle.ValAttr = s
	}
}

// AddRun adds a run to a paragraph.
func (p Paragraph) AddRun() Run {
	pc := wml.NewEG_PContent()
	p.x.EG_PContent = append(p.x.EG_PContent, pc)

	rc := wml.NewEG_ContentRunContent()
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, rc)
	r := wml.NewCT_R()
	rc.R = r
	return Run{p.d, r}
}

// AddRunTrackChange adds an insertion change to a paragraph.
func (p Paragraph) AddRunTrackChange(mode RunTrackChangeMode) RunTrackChange {
	pc := wml.NewEG_PContent()
	crc := wml.NewEG_ContentRunContent()
	rle := wml.NewEG_RunLevelElts()
	rtc := wml.NewCT_RunTrackChange()
	switch mode {
	case RunTrackChangeInsert:
		rle.Ins = rtc
	case RunTrackChangeDelete:
		rle.Del = rtc
	}
	crc.EG_RunLevelElts = append(crc.EG_RunLevelElts, rle)
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, crc)
	p.x.EG_PContent = append(p.x.EG_PContent, pc)

	return RunTrackChange{p.d, rtc}
}

// AddCommentRangeStart adds a start marker for commented text to a paragraph.
func (p Paragraph) AddCommentRangeStart() MarkupRange {
	pc := wml.NewEG_PContent()
	crc := wml.NewEG_ContentRunContent()
	rle := wml.NewEG_RunLevelElts()
	rme := wml.NewEG_RangeMarkupElements()
	mr := wml.NewCT_MarkupRange()
	rme.CommentRangeStart = mr
	rle.EG_RangeMarkupElements = append(rle.EG_RangeMarkupElements, rme)
	crc.EG_RunLevelElts = append(crc.EG_RunLevelElts, rle)
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, crc)
	p.x.EG_PContent = append(p.x.EG_PContent, pc)

	return MarkupRange{p.d, mr}
}

// AddCommentRangeEnd adds an end marker for commented text to a paragraph.
func (p Paragraph) AddCommentRangeEnd() MarkupRange {
	pc := wml.NewEG_PContent()
	crc := wml.NewEG_ContentRunContent()
	rle := wml.NewEG_RunLevelElts()
	rme := wml.NewEG_RangeMarkupElements()
	mr := wml.NewCT_MarkupRange()
	rme.CommentRangeEnd = mr
	rle.EG_RangeMarkupElements = append(rle.EG_RangeMarkupElements, rme)
	crc.EG_RunLevelElts = append(crc.EG_RunLevelElts, rle)
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, crc)
	p.x.EG_PContent = append(p.x.EG_PContent, pc)

	return MarkupRange{p.d, mr}
}

// SetParagraphID sets the paragraph ID (8 hexadecimal digits).
func (p Paragraph) SetParagraphID(id string) {
	p.x.ParaIdAttr = &id
}

// Runs returns all of the runs in a paragraph.
func (p Paragraph) Runs() []Run {
	ret := []Run{}
	for _, c := range p.x.EG_PContent {
		for _, rc := range c.EG_ContentRunContent {
			if rc.R != nil {
				ret = append(ret, Run{p.d, rc.R})
			}
			if rc.Sdt != nil && rc.Sdt.SdtContent != nil {
				for _, rc2 := range rc.Sdt.SdtContent.EG_ContentRunContent {
					if rc2.R != nil {
						ret = append(ret, Run{p.d, rc2.R})
					}
				}
			}
		}
	}
	return ret
}

// InsertRunAfter inserts a run in the paragraph after the relative run.
func (p Paragraph) InsertRunAfter(relativeTo Run) Run {
	return p.insertRun(relativeTo, false)
}

// InsertRunBefore inserts a run in the paragraph before the relative run.
func (p Paragraph) InsertRunBefore(relativeTo Run) Run {
	return p.insertRun(relativeTo, true)
}

func (p Paragraph) insertRun(relativeTo Run, before bool) Run {
	for _, c := range p.x.EG_PContent {
		for i, rc := range c.EG_ContentRunContent {
			if rc.R == relativeTo.X() {
				r := wml.NewCT_R()
				c.EG_ContentRunContent = append(c.EG_ContentRunContent, nil)
				if before {
					copy(c.EG_ContentRunContent[i+1:], c.EG_ContentRunContent[i:])
					c.EG_ContentRunContent[i] = wml.NewEG_ContentRunContent()
					c.EG_ContentRunContent[i].R = r
				} else {
					copy(c.EG_ContentRunContent[i+2:], c.EG_ContentRunContent[i+1:])
					c.EG_ContentRunContent[i+1] = wml.NewEG_ContentRunContent()
					c.EG_ContentRunContent[i+1].R = r
				}
				return Run{p.d, r}

			}
			if rc.Sdt != nil && rc.Sdt.SdtContent != nil {
				for _, rc2 := range rc.Sdt.SdtContent.EG_ContentRunContent {
					if rc2.R == relativeTo.X() {
						r := wml.NewCT_R()
						rc.Sdt.SdtContent.EG_ContentRunContent = append(rc.Sdt.SdtContent.EG_ContentRunContent, nil)
						if before {
							copy(rc.Sdt.SdtContent.EG_ContentRunContent[i+1:], rc.Sdt.SdtContent.EG_ContentRunContent[i:])
							rc.Sdt.SdtContent.EG_ContentRunContent[i] = wml.NewEG_ContentRunContent()
							rc.Sdt.SdtContent.EG_ContentRunContent[i].R = r
						} else {
							copy(rc.Sdt.SdtContent.EG_ContentRunContent[i+2:], rc.Sdt.SdtContent.EG_ContentRunContent[i+1:])
							rc.Sdt.SdtContent.EG_ContentRunContent[i+1] = wml.NewEG_ContentRunContent()
							rc.Sdt.SdtContent.EG_ContentRunContent[i+1].R = r
						}
						return Run{p.d, r}
					}
				}
			}
		}
	}
	return p.AddRun()
}

// AddHyperLink adds a new hyperlink to a parapgraph.
func (p Paragraph) AddHyperLink() HyperLink {
	pc := wml.NewEG_PContent()
	p.x.EG_PContent = append(p.x.EG_PContent, pc)

	pc.Hyperlink = wml.NewCT_Hyperlink()
	return HyperLink{p.d, pc.Hyperlink}
}

// AddBookmark adds a bookmark to a document that can then be used from a hyperlink. Name is a document
// unique name that identifies the bookmark so it can be referenced from hyperlinks.
func (p Paragraph) AddBookmark(id int64, name string) Bookmark {
	pc := wml.NewEG_PContent()
	rc := wml.NewEG_ContentRunContent()
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, rc)

	pcStart := wml.NewEG_PContent()
	rcStart := wml.NewEG_ContentRunContent()
	pcStart.EG_ContentRunContent = append(pcStart.EG_ContentRunContent, rcStart)

	relt := wml.NewEG_RunLevelElts()
	rcStart.EG_RunLevelElts = append(rcStart.EG_RunLevelElts, relt)

	markEl := wml.NewEG_RangeMarkupElements()
	bmStart := wml.NewCT_Bookmark()
	bmStart.IdAttr = id
	markEl.BookmarkStart = bmStart
	relt.EG_RangeMarkupElements = append(relt.EG_RangeMarkupElements, markEl)

	pcEnd := wml.NewEG_PContent()
	rcEnd := wml.NewEG_ContentRunContent()
	pcEnd.EG_ContentRunContent = append(pcEnd.EG_ContentRunContent, rcEnd)

	relt = wml.NewEG_RunLevelElts()
	rcEnd.EG_RunLevelElts = append(rcEnd.EG_RunLevelElts, relt)
	markEl = wml.NewEG_RangeMarkupElements()

	bmEnd := wml.NewCT_MarkupRange()
	bmEnd.IdAttr = id
	markEl.BookmarkEnd = bmEnd

	relt.EG_RangeMarkupElements = append(relt.EG_RangeMarkupElements, markEl)
	p.x.EG_PContent = append(p.x.EG_PContent, pcStart, pcEnd)

	bm := Bookmark{x: bmStart, p: &p, pc: pcStart}
	bm.SetName(name)
	return bm
}

// SetNumberingLevel sets the numbering level of a paragraph.  If used, then the
// NumberingDefinition must also be set via SetNumberingDefinition or
// SetNumberingDefinitionByID.
func (p Paragraph) SetNumberingLevel(listLevel int) {
	p.ensurePPr()
	if p.x.PPr.NumPr == nil {
		p.x.PPr.NumPr = wml.NewCT_NumPr()
	}
	lvl := wml.NewCT_DecimalNumber()
	lvl.ValAttr = int64(listLevel)
	p.x.PPr.NumPr.Ilvl = lvl
}

// SetNumberingDefinition sets the numbering definition ID via a NumberingDefinition
// defined in numbering.xml
func (p Paragraph) SetNumberingDefinition(nd NumberingDefinition) {
	p.ensurePPr()
	if p.x.PPr.NumPr == nil {
		p.x.PPr.NumPr = wml.NewCT_NumPr()
	}
	lvl := wml.NewCT_DecimalNumber()

	numID := int64(-1)
	for _, n := range p.d.Numbering.x.Num {
		if n.AbstractNumId != nil && n.AbstractNumId.ValAttr == nd.AbstractNumberID() {
			numID = n.NumIdAttr
		}
	}
	if numID == -1 {
		num := wml.NewCT_Num()
		p.d.Numbering.x.Num = append(p.d.Numbering.x.Num, num)
		num.NumIdAttr = int64(len(p.d.Numbering.x.Num))
		num.AbstractNumId = wml.NewCT_DecimalNumber()
		num.AbstractNumId.ValAttr = nd.AbstractNumberID()
		numID = num.NumIdAttr
	}

	lvl.ValAttr = numID
	p.x.PPr.NumPr.NumId = lvl
}

// SetNumberingDefinitionByID sets the numbering definition ID directly, which must
// match an ID defined in numbering.xml
func (p Paragraph) SetNumberingDefinitionByID(abstractNumberID int64) {
	p.ensurePPr()
	if p.x.PPr.NumPr == nil {
		p.x.PPr.NumPr = wml.NewCT_NumPr()
	}
	lvl := wml.NewCT_DecimalNumber()
	lvl.ValAttr = int64(abstractNumberID)
	p.x.PPr.NumPr.NumId = lvl
}
