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
	"strconv"

	"github.com/Preciselyco/unioffice"
	"github.com/Preciselyco/unioffice/schema/soo/dml"
)

type WdCT_Inline struct {
	DistTAttr         *uint32
	DistBAttr         *uint32
	DistLAttr         *uint32
	DistRAttr         *uint32
	Extent            *dml.CT_PositiveSize2D
	EffectExtent      *WdCT_EffectExtent
	DocPr             *dml.CT_NonVisualDrawingProps
	CNvGraphicFramePr *dml.CT_NonVisualGraphicFrameProperties
	Graphic           *dml.Graphic
}

func NewWdCT_Inline() *WdCT_Inline {
	ret := &WdCT_Inline{}
	ret.Extent = dml.NewCT_PositiveSize2D()
	ret.DocPr = dml.NewCT_NonVisualDrawingProps()
	ret.Graphic = dml.NewGraphic()
	return ret
}

func (m *WdCT_Inline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DistTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"},
			Value: fmt.Sprintf("%v", *m.DistTAttr)})
	}
	if m.DistBAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"},
			Value: fmt.Sprintf("%v", *m.DistBAttr)})
	}
	if m.DistLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"},
			Value: fmt.Sprintf("%v", *m.DistLAttr)})
	}
	if m.DistRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"},
			Value: fmt.Sprintf("%v", *m.DistRAttr)})
	}
	e.EncodeToken(start)
	seextent := xml.StartElement{Name: xml.Name{Local: "wp:extent"}}
	e.EncodeElement(m.Extent, seextent)
	if m.EffectExtent != nil {
		seeffectExtent := xml.StartElement{Name: xml.Name{Local: "wp:effectExtent"}}
		e.EncodeElement(m.EffectExtent, seeffectExtent)
	}
	sedocPr := xml.StartElement{Name: xml.Name{Local: "wp:docPr"}}
	e.EncodeElement(m.DocPr, sedocPr)
	if m.CNvGraphicFramePr != nil {
		secNvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "wp:cNvGraphicFramePr"}}
		e.EncodeElement(m.CNvGraphicFramePr, secNvGraphicFramePr)
	}
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_Inline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Extent = dml.NewCT_PositiveSize2D()
	m.DocPr = dml.NewCT_NonVisualDrawingProps()
	m.Graphic = dml.NewGraphic()
	for _, attr := range start.Attr {
		if attr.Name.Local == "distT" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistTAttr = &pt
			continue
		}
		if attr.Name.Local == "distB" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistBAttr = &pt
			continue
		}
		if attr.Name.Local == "distR" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistRAttr = &pt
			continue
		}
		if attr.Name.Local == "distL" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistLAttr = &pt
			continue
		}
	}
lWdCT_Inline:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "extent"}:
				if err := d.DecodeElement(m.Extent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "effectExtent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "effectExtent"}:
				m.EffectExtent = NewWdCT_EffectExtent()
				if err := d.DecodeElement(m.EffectExtent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "docPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "docPr"}:
				if err := d.DecodeElement(m.DocPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvGraphicFramePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "cNvGraphicFramePr"}:
				m.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()
				if err := d.DecodeElement(m.CNvGraphicFramePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "graphic"}:
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on WdCT_Inline %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_Inline
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_Inline and its children
func (m *WdCT_Inline) Validate() error {
	return m.ValidateWithPath("WdCT_Inline")
}

// ValidateWithPath validates the WdCT_Inline and its children, prefixing error messages with path
func (m *WdCT_Inline) ValidateWithPath(path string) error {
	if err := m.Extent.ValidateWithPath(path + "/Extent"); err != nil {
		return err
	}
	if m.EffectExtent != nil {
		if err := m.EffectExtent.ValidateWithPath(path + "/EffectExtent"); err != nil {
			return err
		}
	}
	if err := m.DocPr.ValidateWithPath(path + "/DocPr"); err != nil {
		return err
	}
	if m.CNvGraphicFramePr != nil {
		if err := m.CNvGraphicFramePr.ValidateWithPath(path + "/CNvGraphicFramePr"); err != nil {
			return err
		}
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	return nil
}
