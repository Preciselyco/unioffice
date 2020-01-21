// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/Preciselyco/unioffice"
)

type CT_Query struct {
	// MDX Query String
	MdxAttr string
	// Tuples
	Tpls *CT_Tuples
}

func NewCT_Query() *CT_Query {
	ret := &CT_Query{}
	return ret
}

func (m *CT_Query) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mdx"},
		Value: fmt.Sprintf("%v", m.MdxAttr)})
	e.EncodeToken(start)
	if m.Tpls != nil {
		setpls := xml.StartElement{Name: xml.Name{Local: "ma:tpls"}}
		e.EncodeElement(m.Tpls, setpls)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Query) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "mdx" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MdxAttr = parsed
			continue
		}
	}
lCT_Query:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tpls"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "tpls"}:
				m.Tpls = NewCT_Tuples()
				if err := d.DecodeElement(m.Tpls, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Query %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Query
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Query and its children
func (m *CT_Query) Validate() error {
	return m.ValidateWithPath("CT_Query")
}

// ValidateWithPath validates the CT_Query and its children, prefixing error messages with path
func (m *CT_Query) ValidateWithPath(path string) error {
	if m.Tpls != nil {
		if err := m.Tpls.ValidateWithPath(path + "/Tpls"); err != nil {
			return err
		}
	}
	return nil
}
