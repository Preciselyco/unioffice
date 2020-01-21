// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"

	"github.com/Preciselyco/unioffice"
)

type EG_RunInnerContent struct {
	// Break
	Br *CT_Br
	// Text
	T *CT_Text
	// Content Part
	ContentPart *CT_Rel
	// Deleted Text
	DelText *CT_Text
	// Field Code
	InstrText *CT_Text
	// Deleted Field Code
	DelInstrText *CT_Text
	// Non Breaking Hyphen Character
	NoBreakHyphen *CT_Empty
	// Optional Hyphen Character
	SoftHyphen *CT_Empty
	// Date Block - Short Day Format
	DayShort *CT_Empty
	// Date Block - Short Month Format
	MonthShort *CT_Empty
	// Date Block - Short Year Format
	YearShort *CT_Empty
	// Date Block - Long Day Format
	DayLong *CT_Empty
	// Date Block - Long Month Format
	MonthLong *CT_Empty
	// Date Block - Long Year Format
	YearLong *CT_Empty
	// Comment Information Block
	AnnotationRef *CT_Empty
	// Footnote Reference Mark
	FootnoteRef *CT_Empty
	// Endnote Reference Mark
	EndnoteRef *CT_Empty
	// Footnote/Endnote Separator Mark
	Separator *CT_Empty
	// Continuation Separator Mark
	ContinuationSeparator *CT_Empty
	// Symbol Character
	Sym *CT_Sym
	// Page Number Block
	PgNum *CT_Empty
	// Carriage Return
	Cr *CT_Empty
	// Tab Character
	Tab *CT_Empty
	// Embedded Object
	Object *CT_Object
	// VML Object
	Pict *CT_Picture
	// Complex Field Character
	FldChar *CT_FldChar
	// Phonetic Guide
	Ruby *CT_Ruby
	// Footnote Reference
	FootnoteReference *CT_FtnEdnRef
	// Endnote Reference
	EndnoteReference *CT_FtnEdnRef
	// Comment Content Reference Mark
	CommentReference *CT_Markup
	// DrawingML Object
	Drawing *CT_Drawing
	// Absolute Position Tab Character
	Ptab *CT_PTab
	// Position of Last Calculated Page Break
	LastRenderedPageBreak *CT_Empty
}

func NewEG_RunInnerContent() *EG_RunInnerContent {
	ret := &EG_RunInnerContent{}
	return ret
}

func (m *EG_RunInnerContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Br != nil {
		sebr := xml.StartElement{Name: xml.Name{Local: "w:br"}}
		e.EncodeElement(m.Br, sebr)
	}
	if m.T != nil {
		set := xml.StartElement{Name: xml.Name{Local: "w:t"}}
		e.EncodeElement(m.T, set)
	}
	if m.ContentPart != nil {
		secontentPart := xml.StartElement{Name: xml.Name{Local: "w:contentPart"}}
		e.EncodeElement(m.ContentPart, secontentPart)
	}
	if m.DelText != nil {
		sedelText := xml.StartElement{Name: xml.Name{Local: "w:delText"}}
		e.EncodeElement(m.DelText, sedelText)
	}
	if m.InstrText != nil {
		seinstrText := xml.StartElement{Name: xml.Name{Local: "w:instrText"}}
		e.EncodeElement(m.InstrText, seinstrText)
	}
	if m.DelInstrText != nil {
		sedelInstrText := xml.StartElement{Name: xml.Name{Local: "w:delInstrText"}}
		e.EncodeElement(m.DelInstrText, sedelInstrText)
	}
	if m.NoBreakHyphen != nil {
		senoBreakHyphen := xml.StartElement{Name: xml.Name{Local: "w:noBreakHyphen"}}
		e.EncodeElement(m.NoBreakHyphen, senoBreakHyphen)
	}
	if m.SoftHyphen != nil {
		sesoftHyphen := xml.StartElement{Name: xml.Name{Local: "w:softHyphen"}}
		e.EncodeElement(m.SoftHyphen, sesoftHyphen)
	}
	if m.DayShort != nil {
		sedayShort := xml.StartElement{Name: xml.Name{Local: "w:dayShort"}}
		e.EncodeElement(m.DayShort, sedayShort)
	}
	if m.MonthShort != nil {
		semonthShort := xml.StartElement{Name: xml.Name{Local: "w:monthShort"}}
		e.EncodeElement(m.MonthShort, semonthShort)
	}
	if m.YearShort != nil {
		seyearShort := xml.StartElement{Name: xml.Name{Local: "w:yearShort"}}
		e.EncodeElement(m.YearShort, seyearShort)
	}
	if m.DayLong != nil {
		sedayLong := xml.StartElement{Name: xml.Name{Local: "w:dayLong"}}
		e.EncodeElement(m.DayLong, sedayLong)
	}
	if m.MonthLong != nil {
		semonthLong := xml.StartElement{Name: xml.Name{Local: "w:monthLong"}}
		e.EncodeElement(m.MonthLong, semonthLong)
	}
	if m.YearLong != nil {
		seyearLong := xml.StartElement{Name: xml.Name{Local: "w:yearLong"}}
		e.EncodeElement(m.YearLong, seyearLong)
	}
	if m.AnnotationRef != nil {
		seannotationRef := xml.StartElement{Name: xml.Name{Local: "w:annotationRef"}}
		e.EncodeElement(m.AnnotationRef, seannotationRef)
	}
	if m.FootnoteRef != nil {
		sefootnoteRef := xml.StartElement{Name: xml.Name{Local: "w:footnoteRef"}}
		e.EncodeElement(m.FootnoteRef, sefootnoteRef)
	}
	if m.EndnoteRef != nil {
		seendnoteRef := xml.StartElement{Name: xml.Name{Local: "w:endnoteRef"}}
		e.EncodeElement(m.EndnoteRef, seendnoteRef)
	}
	if m.Separator != nil {
		seseparator := xml.StartElement{Name: xml.Name{Local: "w:separator"}}
		e.EncodeElement(m.Separator, seseparator)
	}
	if m.ContinuationSeparator != nil {
		secontinuationSeparator := xml.StartElement{Name: xml.Name{Local: "w:continuationSeparator"}}
		e.EncodeElement(m.ContinuationSeparator, secontinuationSeparator)
	}
	if m.Sym != nil {
		sesym := xml.StartElement{Name: xml.Name{Local: "w:sym"}}
		e.EncodeElement(m.Sym, sesym)
	}
	if m.PgNum != nil {
		sepgNum := xml.StartElement{Name: xml.Name{Local: "w:pgNum"}}
		e.EncodeElement(m.PgNum, sepgNum)
	}
	if m.Cr != nil {
		secr := xml.StartElement{Name: xml.Name{Local: "w:cr"}}
		e.EncodeElement(m.Cr, secr)
	}
	if m.Tab != nil {
		setab := xml.StartElement{Name: xml.Name{Local: "w:tab"}}
		e.EncodeElement(m.Tab, setab)
	}
	if m.Object != nil {
		seobject := xml.StartElement{Name: xml.Name{Local: "w:object"}}
		e.EncodeElement(m.Object, seobject)
	}
	if m.Pict != nil {
		sepict := xml.StartElement{Name: xml.Name{Local: "w:pict"}}
		e.EncodeElement(m.Pict, sepict)
	}
	if m.FldChar != nil {
		sefldChar := xml.StartElement{Name: xml.Name{Local: "w:fldChar"}}
		e.EncodeElement(m.FldChar, sefldChar)
	}
	if m.Ruby != nil {
		seruby := xml.StartElement{Name: xml.Name{Local: "w:ruby"}}
		e.EncodeElement(m.Ruby, seruby)
	}
	if m.FootnoteReference != nil {
		sefootnoteReference := xml.StartElement{Name: xml.Name{Local: "w:footnoteReference"}}
		e.EncodeElement(m.FootnoteReference, sefootnoteReference)
	}
	if m.EndnoteReference != nil {
		seendnoteReference := xml.StartElement{Name: xml.Name{Local: "w:endnoteReference"}}
		e.EncodeElement(m.EndnoteReference, seendnoteReference)
	}
	if m.CommentReference != nil {
		secommentReference := xml.StartElement{Name: xml.Name{Local: "w:commentReference"}}
		e.EncodeElement(m.CommentReference, secommentReference)
	}
	if m.Drawing != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "w:drawing"}}
		e.EncodeElement(m.Drawing, sedrawing)
	}
	if m.Ptab != nil {
		septab := xml.StartElement{Name: xml.Name{Local: "w:ptab"}}
		e.EncodeElement(m.Ptab, septab)
	}
	if m.LastRenderedPageBreak != nil {
		selastRenderedPageBreak := xml.StartElement{Name: xml.Name{Local: "w:lastRenderedPageBreak"}}
		e.EncodeElement(m.LastRenderedPageBreak, selastRenderedPageBreak)
	}
	return nil
}

func (m *EG_RunInnerContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_RunInnerContent:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "br"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "br"}:
				m.Br = NewCT_Br()
				if err := d.DecodeElement(m.Br, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "t"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "t"}:
				m.T = NewCT_Text()
				if err := d.DecodeElement(m.T, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "contentPart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "contentPart"}:
				m.ContentPart = NewCT_Rel()
				if err := d.DecodeElement(m.ContentPart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "delText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "delText"}:
				m.DelText = NewCT_Text()
				if err := d.DecodeElement(m.DelText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "instrText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "instrText"}:
				m.InstrText = NewCT_Text()
				if err := d.DecodeElement(m.InstrText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "delInstrText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "delInstrText"}:
				m.DelInstrText = NewCT_Text()
				if err := d.DecodeElement(m.DelInstrText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noBreakHyphen"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "noBreakHyphen"}:
				m.NoBreakHyphen = NewCT_Empty()
				if err := d.DecodeElement(m.NoBreakHyphen, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "softHyphen"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "softHyphen"}:
				m.SoftHyphen = NewCT_Empty()
				if err := d.DecodeElement(m.SoftHyphen, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "dayShort"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "dayShort"}:
				m.DayShort = NewCT_Empty()
				if err := d.DecodeElement(m.DayShort, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "monthShort"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "monthShort"}:
				m.MonthShort = NewCT_Empty()
				if err := d.DecodeElement(m.MonthShort, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "yearShort"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "yearShort"}:
				m.YearShort = NewCT_Empty()
				if err := d.DecodeElement(m.YearShort, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "dayLong"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "dayLong"}:
				m.DayLong = NewCT_Empty()
				if err := d.DecodeElement(m.DayLong, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "monthLong"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "monthLong"}:
				m.MonthLong = NewCT_Empty()
				if err := d.DecodeElement(m.MonthLong, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "yearLong"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "yearLong"}:
				m.YearLong = NewCT_Empty()
				if err := d.DecodeElement(m.YearLong, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "annotationRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "annotationRef"}:
				m.AnnotationRef = NewCT_Empty()
				if err := d.DecodeElement(m.AnnotationRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "footnoteRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "footnoteRef"}:
				m.FootnoteRef = NewCT_Empty()
				if err := d.DecodeElement(m.FootnoteRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "endnoteRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "endnoteRef"}:
				m.EndnoteRef = NewCT_Empty()
				if err := d.DecodeElement(m.EndnoteRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "separator"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "separator"}:
				m.Separator = NewCT_Empty()
				if err := d.DecodeElement(m.Separator, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "continuationSeparator"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "continuationSeparator"}:
				m.ContinuationSeparator = NewCT_Empty()
				if err := d.DecodeElement(m.ContinuationSeparator, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sym"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "sym"}:
				m.Sym = NewCT_Sym()
				if err := d.DecodeElement(m.Sym, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pgNum"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pgNum"}:
				m.PgNum = NewCT_Empty()
				if err := d.DecodeElement(m.PgNum, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "cr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "cr"}:
				m.Cr = NewCT_Empty()
				if err := d.DecodeElement(m.Cr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tab"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "tab"}:
				m.Tab = NewCT_Empty()
				if err := d.DecodeElement(m.Tab, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "object"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "object"}:
				m.Object = NewCT_Object()
				if err := d.DecodeElement(m.Object, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pict"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pict"}:
				m.Pict = NewCT_Picture()
				if err := d.DecodeElement(m.Pict, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "fldChar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "fldChar"}:
				m.FldChar = NewCT_FldChar()
				if err := d.DecodeElement(m.FldChar, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ruby"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "ruby"}:
				m.Ruby = NewCT_Ruby()
				if err := d.DecodeElement(m.Ruby, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "footnoteReference"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "footnoteReference"}:
				m.FootnoteReference = NewCT_FtnEdnRef()
				if err := d.DecodeElement(m.FootnoteReference, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "endnoteReference"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "endnoteReference"}:
				m.EndnoteReference = NewCT_FtnEdnRef()
				if err := d.DecodeElement(m.EndnoteReference, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentReference"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "commentReference"}:
				m.CommentReference = NewCT_Markup()
				if err := d.DecodeElement(m.CommentReference, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "drawing"}:
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ptab"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "ptab"}:
				m.Ptab = NewCT_PTab()
				if err := d.DecodeElement(m.Ptab, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "lastRenderedPageBreak"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "lastRenderedPageBreak"}:
				m.LastRenderedPageBreak = NewCT_Empty()
				if err := d.DecodeElement(m.LastRenderedPageBreak, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_RunInnerContent %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_RunInnerContent
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_RunInnerContent and its children
func (m *EG_RunInnerContent) Validate() error {
	return m.ValidateWithPath("EG_RunInnerContent")
}

// ValidateWithPath validates the EG_RunInnerContent and its children, prefixing error messages with path
func (m *EG_RunInnerContent) ValidateWithPath(path string) error {
	if m.Br != nil {
		if err := m.Br.ValidateWithPath(path + "/Br"); err != nil {
			return err
		}
	}
	if m.T != nil {
		if err := m.T.ValidateWithPath(path + "/T"); err != nil {
			return err
		}
	}
	if m.ContentPart != nil {
		if err := m.ContentPart.ValidateWithPath(path + "/ContentPart"); err != nil {
			return err
		}
	}
	if m.DelText != nil {
		if err := m.DelText.ValidateWithPath(path + "/DelText"); err != nil {
			return err
		}
	}
	if m.InstrText != nil {
		if err := m.InstrText.ValidateWithPath(path + "/InstrText"); err != nil {
			return err
		}
	}
	if m.DelInstrText != nil {
		if err := m.DelInstrText.ValidateWithPath(path + "/DelInstrText"); err != nil {
			return err
		}
	}
	if m.NoBreakHyphen != nil {
		if err := m.NoBreakHyphen.ValidateWithPath(path + "/NoBreakHyphen"); err != nil {
			return err
		}
	}
	if m.SoftHyphen != nil {
		if err := m.SoftHyphen.ValidateWithPath(path + "/SoftHyphen"); err != nil {
			return err
		}
	}
	if m.DayShort != nil {
		if err := m.DayShort.ValidateWithPath(path + "/DayShort"); err != nil {
			return err
		}
	}
	if m.MonthShort != nil {
		if err := m.MonthShort.ValidateWithPath(path + "/MonthShort"); err != nil {
			return err
		}
	}
	if m.YearShort != nil {
		if err := m.YearShort.ValidateWithPath(path + "/YearShort"); err != nil {
			return err
		}
	}
	if m.DayLong != nil {
		if err := m.DayLong.ValidateWithPath(path + "/DayLong"); err != nil {
			return err
		}
	}
	if m.MonthLong != nil {
		if err := m.MonthLong.ValidateWithPath(path + "/MonthLong"); err != nil {
			return err
		}
	}
	if m.YearLong != nil {
		if err := m.YearLong.ValidateWithPath(path + "/YearLong"); err != nil {
			return err
		}
	}
	if m.AnnotationRef != nil {
		if err := m.AnnotationRef.ValidateWithPath(path + "/AnnotationRef"); err != nil {
			return err
		}
	}
	if m.FootnoteRef != nil {
		if err := m.FootnoteRef.ValidateWithPath(path + "/FootnoteRef"); err != nil {
			return err
		}
	}
	if m.EndnoteRef != nil {
		if err := m.EndnoteRef.ValidateWithPath(path + "/EndnoteRef"); err != nil {
			return err
		}
	}
	if m.Separator != nil {
		if err := m.Separator.ValidateWithPath(path + "/Separator"); err != nil {
			return err
		}
	}
	if m.ContinuationSeparator != nil {
		if err := m.ContinuationSeparator.ValidateWithPath(path + "/ContinuationSeparator"); err != nil {
			return err
		}
	}
	if m.Sym != nil {
		if err := m.Sym.ValidateWithPath(path + "/Sym"); err != nil {
			return err
		}
	}
	if m.PgNum != nil {
		if err := m.PgNum.ValidateWithPath(path + "/PgNum"); err != nil {
			return err
		}
	}
	if m.Cr != nil {
		if err := m.Cr.ValidateWithPath(path + "/Cr"); err != nil {
			return err
		}
	}
	if m.Tab != nil {
		if err := m.Tab.ValidateWithPath(path + "/Tab"); err != nil {
			return err
		}
	}
	if m.Object != nil {
		if err := m.Object.ValidateWithPath(path + "/Object"); err != nil {
			return err
		}
	}
	if m.Pict != nil {
		if err := m.Pict.ValidateWithPath(path + "/Pict"); err != nil {
			return err
		}
	}
	if m.FldChar != nil {
		if err := m.FldChar.ValidateWithPath(path + "/FldChar"); err != nil {
			return err
		}
	}
	if m.Ruby != nil {
		if err := m.Ruby.ValidateWithPath(path + "/Ruby"); err != nil {
			return err
		}
	}
	if m.FootnoteReference != nil {
		if err := m.FootnoteReference.ValidateWithPath(path + "/FootnoteReference"); err != nil {
			return err
		}
	}
	if m.EndnoteReference != nil {
		if err := m.EndnoteReference.ValidateWithPath(path + "/EndnoteReference"); err != nil {
			return err
		}
	}
	if m.CommentReference != nil {
		if err := m.CommentReference.ValidateWithPath(path + "/CommentReference"); err != nil {
			return err
		}
	}
	if m.Drawing != nil {
		if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
			return err
		}
	}
	if m.Ptab != nil {
		if err := m.Ptab.ValidateWithPath(path + "/Ptab"); err != nil {
			return err
		}
	}
	if m.LastRenderedPageBreak != nil {
		if err := m.LastRenderedPageBreak.ValidateWithPath(path + "/LastRenderedPageBreak"); err != nil {
			return err
		}
	}
	return nil
}
