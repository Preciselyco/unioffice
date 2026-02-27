package custom_properties

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/Preciselyco/unioffice"
)

type CT_Property struct {
	FmtidAttr      string
	PidAttr        int32
	NameAttr       *string
	LinkTargetAttr *string

	// Variant value fields
	Lpwstr   *string
	Lpstr    *string
	I4       *int32
	R8       *float64
	Bool     *bool
	Filetime *time.Time
}

func NewCT_Property() *CT_Property {
	ret := &CT_Property{}
	ret.FmtidAttr = "{D5CDD505-2E9C-101B-9397-08002B2CF9AE}"
	ret.PidAttr = 2
	return ret
}

func (m *CT_Property) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fmtid"}, Value: m.FmtidAttr})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pid"}, Value: fmt.Sprintf("%d", m.PidAttr)})
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"}, Value: *m.NameAttr})
	}
	if m.LinkTargetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "linkTarget"}, Value: *m.LinkTargetAttr})
	}
	e.EncodeToken(start)
	if m.Lpwstr != nil {
		selpwstr := xml.StartElement{Name: xml.Name{Local: "vt:lpwstr"}}
		unioffice.AddPreserveSpaceAttr(&selpwstr, *m.Lpwstr)
		e.EncodeElement(m.Lpwstr, selpwstr)
	}
	if m.Lpstr != nil {
		selpstr := xml.StartElement{Name: xml.Name{Local: "vt:lpstr"}}
		unioffice.AddPreserveSpaceAttr(&selpstr, *m.Lpstr)
		e.EncodeElement(m.Lpstr, selpstr)
	}
	if m.I4 != nil {
		sei4 := xml.StartElement{Name: xml.Name{Local: "vt:i4"}}
		e.EncodeElement(m.I4, sei4)
	}
	if m.R8 != nil {
		ser8 := xml.StartElement{Name: xml.Name{Local: "vt:r8"}}
		e.EncodeElement(m.R8, ser8)
	}
	if m.Bool != nil {
		sebool := xml.StartElement{Name: xml.Name{Local: "vt:bool"}}
		e.EncodeElement(m.Bool, sebool)
	}
	if m.Filetime != nil {
		sefiletime := xml.StartElement{Name: xml.Name{Local: "vt:filetime"}}
		e.EncodeElement(m.Filetime.Format(time.RFC3339), sefiletime)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Property) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "fmtid" {
			m.FmtidAttr = attr.Value
		}
		if attr.Name.Local == "pid" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.PidAttr = int32(parsed)
		}
		if attr.Name.Local == "name" {
			parsed := attr.Value
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "linkTarget" {
			parsed := attr.Value
			m.LinkTargetAttr = &parsed
		}
	}
lCT_Property:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "lpwstr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "lpwstr"}:
				m.Lpwstr = new(string)
				if err := d.DecodeElement(m.Lpwstr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "lpstr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "lpstr"}:
				m.Lpstr = new(string)
				if err := d.DecodeElement(m.Lpstr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i4"}:
				m.I4 = new(int32)
				if err := d.DecodeElement(m.I4, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "r8"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "r8"}:
				m.R8 = new(float64)
				if err := d.DecodeElement(m.R8, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "bool"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "bool"}:
				m.Bool = new(bool)
				if err := d.DecodeElement(m.Bool, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "filetime"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "filetime"}:
				var s string
				if err := d.DecodeElement(&s, &el); err != nil {
					return err
				}
				t, err := time.Parse(time.RFC3339, s)
				if err != nil {
					return fmt.Errorf("error parsing filetime: %s", err)
				}
				m.Filetime = &t
			default:
				unioffice.Log("skipping unsupported element on CT_Property %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Property
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Property and its children
func (m *CT_Property) Validate() error {
	return m.ValidateWithPath("CT_Property")
}

// ValidateWithPath validates the CT_Property and its children, prefixing error messages with path
func (m *CT_Property) ValidateWithPath(path string) error {
	return nil
}
