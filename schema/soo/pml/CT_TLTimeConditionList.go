// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/Preciselyco/unioffice"
)

type CT_TLTimeConditionList struct {
	// Condition
	Cond []*CT_TLTimeCondition
}

func NewCT_TLTimeConditionList() *CT_TLTimeConditionList {
	ret := &CT_TLTimeConditionList{}
	return ret
}

func (m *CT_TLTimeConditionList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	second := xml.StartElement{Name: xml.Name{Local: "p:cond"}}
	for _, c := range m.Cond {
		e.EncodeElement(c, second)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTimeConditionList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLTimeConditionList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cond"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cond"}:
				tmp := NewCT_TLTimeCondition()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cond = append(m.Cond, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_TLTimeConditionList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeConditionList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTimeConditionList and its children
func (m *CT_TLTimeConditionList) Validate() error {
	return m.ValidateWithPath("CT_TLTimeConditionList")
}

// ValidateWithPath validates the CT_TLTimeConditionList and its children, prefixing error messages with path
func (m *CT_TLTimeConditionList) ValidateWithPath(path string) error {
	for i, v := range m.Cond {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cond[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
