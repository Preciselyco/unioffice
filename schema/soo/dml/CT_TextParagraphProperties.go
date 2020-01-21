// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/Preciselyco/unioffice"
)

type CT_TextParagraphProperties struct {
	MarLAttr         *int32
	MarRAttr         *int32
	LvlAttr          *int32
	IndentAttr       *int32
	AlgnAttr         ST_TextAlignType
	DefTabSzAttr     *ST_Coordinate32
	RtlAttr          *bool
	EaLnBrkAttr      *bool
	FontAlgnAttr     ST_TextFontAlignType
	LatinLnBrkAttr   *bool
	HangingPunctAttr *bool
	LnSpc            *CT_TextSpacing
	SpcBef           *CT_TextSpacing
	SpcAft           *CT_TextSpacing
	BuClrTx          *CT_TextBulletColorFollowText
	BuClr            *CT_Color
	BuSzTx           *CT_TextBulletSizeFollowText
	BuSzPct          *CT_TextBulletSizePercent
	BuSzPts          *CT_TextBulletSizePoint
	BuFontTx         *CT_TextBulletTypefaceFollowText
	BuFont           *CT_TextFont
	BuNone           *CT_TextNoBullet
	BuAutoNum        *CT_TextAutonumberBullet
	BuChar           *CT_TextCharBullet
	BuBlip           *CT_TextBlipBullet
	TabLst           *CT_TextTabStopList
	DefRPr           *CT_TextCharacterProperties
	ExtLst           *CT_OfficeArtExtensionList
}

func NewCT_TextParagraphProperties() *CT_TextParagraphProperties {
	ret := &CT_TextParagraphProperties{}
	return ret
}

func (m *CT_TextParagraphProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MarLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "marL"},
			Value: fmt.Sprintf("%v", *m.MarLAttr)})
	}
	if m.MarRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "marR"},
			Value: fmt.Sprintf("%v", *m.MarRAttr)})
	}
	if m.LvlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lvl"},
			Value: fmt.Sprintf("%v", *m.LvlAttr)})
	}
	if m.IndentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "indent"},
			Value: fmt.Sprintf("%v", *m.IndentAttr)})
	}
	if m.AlgnAttr != ST_TextAlignTypeUnset {
		attr, err := m.AlgnAttr.MarshalXMLAttr(xml.Name{Local: "algn"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DefTabSzAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defTabSz"},
			Value: fmt.Sprintf("%v", *m.DefTabSzAttr)})
	}
	if m.RtlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rtl"},
			Value: fmt.Sprintf("%d", b2i(*m.RtlAttr))})
	}
	if m.EaLnBrkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "eaLnBrk"},
			Value: fmt.Sprintf("%d", b2i(*m.EaLnBrkAttr))})
	}
	if m.FontAlgnAttr != ST_TextFontAlignTypeUnset {
		attr, err := m.FontAlgnAttr.MarshalXMLAttr(xml.Name{Local: "fontAlgn"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.LatinLnBrkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "latinLnBrk"},
			Value: fmt.Sprintf("%d", b2i(*m.LatinLnBrkAttr))})
	}
	if m.HangingPunctAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hangingPunct"},
			Value: fmt.Sprintf("%d", b2i(*m.HangingPunctAttr))})
	}
	e.EncodeToken(start)
	if m.LnSpc != nil {
		selnSpc := xml.StartElement{Name: xml.Name{Local: "a:lnSpc"}}
		e.EncodeElement(m.LnSpc, selnSpc)
	}
	if m.SpcBef != nil {
		sespcBef := xml.StartElement{Name: xml.Name{Local: "a:spcBef"}}
		e.EncodeElement(m.SpcBef, sespcBef)
	}
	if m.SpcAft != nil {
		sespcAft := xml.StartElement{Name: xml.Name{Local: "a:spcAft"}}
		e.EncodeElement(m.SpcAft, sespcAft)
	}
	if m.BuClrTx != nil {
		sebuClrTx := xml.StartElement{Name: xml.Name{Local: "a:buClrTx"}}
		e.EncodeElement(m.BuClrTx, sebuClrTx)
	}
	if m.BuClr != nil {
		sebuClr := xml.StartElement{Name: xml.Name{Local: "a:buClr"}}
		e.EncodeElement(m.BuClr, sebuClr)
	}
	if m.BuSzTx != nil {
		sebuSzTx := xml.StartElement{Name: xml.Name{Local: "a:buSzTx"}}
		e.EncodeElement(m.BuSzTx, sebuSzTx)
	}
	if m.BuSzPct != nil {
		sebuSzPct := xml.StartElement{Name: xml.Name{Local: "a:buSzPct"}}
		e.EncodeElement(m.BuSzPct, sebuSzPct)
	}
	if m.BuSzPts != nil {
		sebuSzPts := xml.StartElement{Name: xml.Name{Local: "a:buSzPts"}}
		e.EncodeElement(m.BuSzPts, sebuSzPts)
	}
	if m.BuFontTx != nil {
		sebuFontTx := xml.StartElement{Name: xml.Name{Local: "a:buFontTx"}}
		e.EncodeElement(m.BuFontTx, sebuFontTx)
	}
	if m.BuFont != nil {
		sebuFont := xml.StartElement{Name: xml.Name{Local: "a:buFont"}}
		e.EncodeElement(m.BuFont, sebuFont)
	}
	if m.BuNone != nil {
		sebuNone := xml.StartElement{Name: xml.Name{Local: "a:buNone"}}
		e.EncodeElement(m.BuNone, sebuNone)
	}
	if m.BuAutoNum != nil {
		sebuAutoNum := xml.StartElement{Name: xml.Name{Local: "a:buAutoNum"}}
		e.EncodeElement(m.BuAutoNum, sebuAutoNum)
	}
	if m.BuChar != nil {
		sebuChar := xml.StartElement{Name: xml.Name{Local: "a:buChar"}}
		e.EncodeElement(m.BuChar, sebuChar)
	}
	if m.BuBlip != nil {
		sebuBlip := xml.StartElement{Name: xml.Name{Local: "a:buBlip"}}
		e.EncodeElement(m.BuBlip, sebuBlip)
	}
	if m.TabLst != nil {
		setabLst := xml.StartElement{Name: xml.Name{Local: "a:tabLst"}}
		e.EncodeElement(m.TabLst, setabLst)
	}
	if m.DefRPr != nil {
		sedefRPr := xml.StartElement{Name: xml.Name{Local: "a:defRPr"}}
		e.EncodeElement(m.DefRPr, sedefRPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextParagraphProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "marL" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.MarLAttr = &pt
			continue
		}
		if attr.Name.Local == "lvl" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.LvlAttr = &pt
			continue
		}
		if attr.Name.Local == "algn" {
			m.AlgnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "rtl" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RtlAttr = &parsed
			continue
		}
		if attr.Name.Local == "fontAlgn" {
			m.FontAlgnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "marR" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.MarRAttr = &pt
			continue
		}
		if attr.Name.Local == "latinLnBrk" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LatinLnBrkAttr = &parsed
			continue
		}
		if attr.Name.Local == "indent" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.IndentAttr = &pt
			continue
		}
		if attr.Name.Local == "eaLnBrk" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EaLnBrkAttr = &parsed
			continue
		}
		if attr.Name.Local == "hangingPunct" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HangingPunctAttr = &parsed
			continue
		}
		if attr.Name.Local == "defTabSz" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.DefTabSzAttr = &parsed
			continue
		}
	}
lCT_TextParagraphProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lnSpc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lnSpc"}:
				m.LnSpc = NewCT_TextSpacing()
				if err := d.DecodeElement(m.LnSpc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "spcBef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "spcBef"}:
				m.SpcBef = NewCT_TextSpacing()
				if err := d.DecodeElement(m.SpcBef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "spcAft"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "spcAft"}:
				m.SpcAft = NewCT_TextSpacing()
				if err := d.DecodeElement(m.SpcAft, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buClrTx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buClrTx"}:
				m.BuClrTx = NewCT_TextBulletColorFollowText()
				if err := d.DecodeElement(m.BuClrTx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buClr"}:
				m.BuClr = NewCT_Color()
				if err := d.DecodeElement(m.BuClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buSzTx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buSzTx"}:
				m.BuSzTx = NewCT_TextBulletSizeFollowText()
				if err := d.DecodeElement(m.BuSzTx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buSzPct"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buSzPct"}:
				m.BuSzPct = NewCT_TextBulletSizePercent()
				if err := d.DecodeElement(m.BuSzPct, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buSzPts"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buSzPts"}:
				m.BuSzPts = NewCT_TextBulletSizePoint()
				if err := d.DecodeElement(m.BuSzPts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buFontTx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buFontTx"}:
				m.BuFontTx = NewCT_TextBulletTypefaceFollowText()
				if err := d.DecodeElement(m.BuFontTx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buFont"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buFont"}:
				m.BuFont = NewCT_TextFont()
				if err := d.DecodeElement(m.BuFont, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buNone"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buNone"}:
				m.BuNone = NewCT_TextNoBullet()
				if err := d.DecodeElement(m.BuNone, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buAutoNum"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buAutoNum"}:
				m.BuAutoNum = NewCT_TextAutonumberBullet()
				if err := d.DecodeElement(m.BuAutoNum, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buChar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buChar"}:
				m.BuChar = NewCT_TextCharBullet()
				if err := d.DecodeElement(m.BuChar, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buBlip"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buBlip"}:
				m.BuBlip = NewCT_TextBlipBullet()
				if err := d.DecodeElement(m.BuBlip, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tabLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tabLst"}:
				m.TabLst = NewCT_TextTabStopList()
				if err := d.DecodeElement(m.TabLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "defRPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "defRPr"}:
				m.DefRPr = NewCT_TextCharacterProperties()
				if err := d.DecodeElement(m.DefRPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TextParagraphProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextParagraphProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextParagraphProperties and its children
func (m *CT_TextParagraphProperties) Validate() error {
	return m.ValidateWithPath("CT_TextParagraphProperties")
}

// ValidateWithPath validates the CT_TextParagraphProperties and its children, prefixing error messages with path
func (m *CT_TextParagraphProperties) ValidateWithPath(path string) error {
	if m.MarLAttr != nil {
		if *m.MarLAttr < 0 {
			return fmt.Errorf("%s/m.MarLAttr must be >= 0 (have %v)", path, *m.MarLAttr)
		}
		if *m.MarLAttr > 51206400 {
			return fmt.Errorf("%s/m.MarLAttr must be <= 51206400 (have %v)", path, *m.MarLAttr)
		}
	}
	if m.MarRAttr != nil {
		if *m.MarRAttr < 0 {
			return fmt.Errorf("%s/m.MarRAttr must be >= 0 (have %v)", path, *m.MarRAttr)
		}
		if *m.MarRAttr > 51206400 {
			return fmt.Errorf("%s/m.MarRAttr must be <= 51206400 (have %v)", path, *m.MarRAttr)
		}
	}
	if m.LvlAttr != nil {
		if *m.LvlAttr < 0 {
			return fmt.Errorf("%s/m.LvlAttr must be >= 0 (have %v)", path, *m.LvlAttr)
		}
		if *m.LvlAttr > 8 {
			return fmt.Errorf("%s/m.LvlAttr must be <= 8 (have %v)", path, *m.LvlAttr)
		}
	}
	if m.IndentAttr != nil {
		if *m.IndentAttr < -51206400 {
			return fmt.Errorf("%s/m.IndentAttr must be >= -51206400 (have %v)", path, *m.IndentAttr)
		}
		if *m.IndentAttr > 51206400 {
			return fmt.Errorf("%s/m.IndentAttr must be <= 51206400 (have %v)", path, *m.IndentAttr)
		}
	}
	if err := m.AlgnAttr.ValidateWithPath(path + "/AlgnAttr"); err != nil {
		return err
	}
	if m.DefTabSzAttr != nil {
		if err := m.DefTabSzAttr.ValidateWithPath(path + "/DefTabSzAttr"); err != nil {
			return err
		}
	}
	if err := m.FontAlgnAttr.ValidateWithPath(path + "/FontAlgnAttr"); err != nil {
		return err
	}
	if m.LnSpc != nil {
		if err := m.LnSpc.ValidateWithPath(path + "/LnSpc"); err != nil {
			return err
		}
	}
	if m.SpcBef != nil {
		if err := m.SpcBef.ValidateWithPath(path + "/SpcBef"); err != nil {
			return err
		}
	}
	if m.SpcAft != nil {
		if err := m.SpcAft.ValidateWithPath(path + "/SpcAft"); err != nil {
			return err
		}
	}
	if m.BuClrTx != nil {
		if err := m.BuClrTx.ValidateWithPath(path + "/BuClrTx"); err != nil {
			return err
		}
	}
	if m.BuClr != nil {
		if err := m.BuClr.ValidateWithPath(path + "/BuClr"); err != nil {
			return err
		}
	}
	if m.BuSzTx != nil {
		if err := m.BuSzTx.ValidateWithPath(path + "/BuSzTx"); err != nil {
			return err
		}
	}
	if m.BuSzPct != nil {
		if err := m.BuSzPct.ValidateWithPath(path + "/BuSzPct"); err != nil {
			return err
		}
	}
	if m.BuSzPts != nil {
		if err := m.BuSzPts.ValidateWithPath(path + "/BuSzPts"); err != nil {
			return err
		}
	}
	if m.BuFontTx != nil {
		if err := m.BuFontTx.ValidateWithPath(path + "/BuFontTx"); err != nil {
			return err
		}
	}
	if m.BuFont != nil {
		if err := m.BuFont.ValidateWithPath(path + "/BuFont"); err != nil {
			return err
		}
	}
	if m.BuNone != nil {
		if err := m.BuNone.ValidateWithPath(path + "/BuNone"); err != nil {
			return err
		}
	}
	if m.BuAutoNum != nil {
		if err := m.BuAutoNum.ValidateWithPath(path + "/BuAutoNum"); err != nil {
			return err
		}
	}
	if m.BuChar != nil {
		if err := m.BuChar.ValidateWithPath(path + "/BuChar"); err != nil {
			return err
		}
	}
	if m.BuBlip != nil {
		if err := m.BuBlip.ValidateWithPath(path + "/BuBlip"); err != nil {
			return err
		}
	}
	if m.TabLst != nil {
		if err := m.TabLst.ValidateWithPath(path + "/TabLst"); err != nil {
			return err
		}
	}
	if m.DefRPr != nil {
		if err := m.DefRPr.ValidateWithPath(path + "/DefRPr"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
