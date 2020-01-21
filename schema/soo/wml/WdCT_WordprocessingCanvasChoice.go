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
	"fmt"

	"github.com/Preciselyco/unioffice"
	"github.com/Preciselyco/unioffice/schema/soo/dml/picture"
)

type WdCT_WordprocessingCanvasChoice struct {
	Wsp          []*WdWsp
	Pic          []*picture.Pic
	ContentPart  []*WdCT_WordprocessingContentPart
	Wgp          []*WdWgp
	GraphicFrame []*WdCT_GraphicFrame
}

func NewWdCT_WordprocessingCanvasChoice() *WdCT_WordprocessingCanvasChoice {
	ret := &WdCT_WordprocessingCanvasChoice{}
	return ret
}

func (m *WdCT_WordprocessingCanvasChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Wsp != nil {
		sewsp := xml.StartElement{Name: xml.Name{Local: "wp:wsp"}}
		for _, c := range m.Wsp {
			e.EncodeElement(c, sewsp)
		}
	}
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "pic:pic"}}
		for _, c := range m.Pic {
			e.EncodeElement(c, sepic)
		}
	}
	if m.ContentPart != nil {
		secontentPart := xml.StartElement{Name: xml.Name{Local: "wp:contentPart"}}
		for _, c := range m.ContentPart {
			e.EncodeElement(c, secontentPart)
		}
	}
	if m.Wgp != nil {
		sewgp := xml.StartElement{Name: xml.Name{Local: "wp:wgp"}}
		for _, c := range m.Wgp {
			e.EncodeElement(c, sewgp)
		}
	}
	if m.GraphicFrame != nil {
		segraphicFrame := xml.StartElement{Name: xml.Name{Local: "wp:graphicFrame"}}
		for _, c := range m.GraphicFrame {
			e.EncodeElement(c, segraphicFrame)
		}
	}
	return nil
}

func (m *WdCT_WordprocessingCanvasChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdCT_WordprocessingCanvasChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wsp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wsp"}:
				tmp := NewWdWsp()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Wsp = append(m.Wsp, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "pic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/picture", Local: "pic"}:
				tmp := picture.NewPic()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Pic = append(m.Pic, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "contentPart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "contentPart"}:
				tmp := NewWdCT_WordprocessingContentPart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ContentPart = append(m.ContentPart, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wgp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wgp"}:
				tmp := NewWdWgp()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Wgp = append(m.Wgp, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "graphicFrame"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "graphicFrame"}:
				tmp := NewWdCT_GraphicFrame()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GraphicFrame = append(m.GraphicFrame, tmp)
			default:
				unioffice.Log("skipping unsupported element on WdCT_WordprocessingCanvasChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WordprocessingCanvasChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WordprocessingCanvasChoice and its children
func (m *WdCT_WordprocessingCanvasChoice) Validate() error {
	return m.ValidateWithPath("WdCT_WordprocessingCanvasChoice")
}

// ValidateWithPath validates the WdCT_WordprocessingCanvasChoice and its children, prefixing error messages with path
func (m *WdCT_WordprocessingCanvasChoice) ValidateWithPath(path string) error {
	for i, v := range m.Wsp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Wsp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Pic {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Pic[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ContentPart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ContentPart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Wgp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Wgp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GraphicFrame {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GraphicFrame[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
