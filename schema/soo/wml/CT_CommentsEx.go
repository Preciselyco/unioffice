// Added by Precisely
// This file is based on CT_Comments.go
// Schema: https://docs.microsoft.com/en-us/openspecs/office_standards/ms-docx/d416013d-c112-44fa-8bef-7819b8898117

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/Preciselyco/unioffice"
)

type CT_CommentsEx struct {
	// Comment Content
	CommentEx []*CT_CommentEx
}

func NewCT_CommentsEx() *CT_CommentsEx {
	ret := &CT_CommentsEx{}
	return ret
}

func (m *CT_CommentsEx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CommentEx != nil {
		commentStart := xml.StartElement{Name: xml.Name{Local: "w12:commentEx"}}
		for _, c := range m.CommentEx {
			e.EncodeElement(c, commentStart)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CommentsEx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CommentsEx:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.microsoft.com/office/word/2012/wordml", Local: "commentEx"}:
				tmp := NewCT_CommentEx()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CommentEx = append(m.CommentEx, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_CommentsEx %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CommentsEx
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CommentsEx and its children
func (m *CT_CommentsEx) Validate() error {
	return m.ValidateWithPath("CT_CommentsEx")
}

// ValidateWithPath validates the CT_CommentsEx and its children, prefixing error messages with path
func (m *CT_CommentsEx) ValidateWithPath(path string) error {
	for i, v := range m.CommentEx {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CommentEx[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
