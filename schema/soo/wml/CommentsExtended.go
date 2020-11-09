// Added by Precisely
// This file is based on Comments.go
// Schema: https://docs.microsoft.com/en-us/openspecs/office_standards/ms-docx/d416013d-c112-44fa-8bef-7819b8898117

package wml

import (
	"encoding/xml"

	"github.com/Preciselyco/unioffice"
)

type CommentsExtended struct {
	CT_CommentsEx
}

func NewCommentsExtended() *CommentsExtended {
	ret := &CommentsExtended{}
	ret.CT_CommentsEx = *NewCT_CommentsEx()
	return ret
}

func (m *CommentsExtended) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w12"}, Value: "http://schemas.microsoft.com/office/word/2012/wordml"})
	start.Name.Local = "w12:commentsEx"
	return m.CT_CommentsEx.MarshalXML(e, start)
}

func (m *CommentsExtended) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_CommentsEx = *NewCT_CommentsEx()
lCommentsExtended:
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
				unioffice.Log("skipping unsupported element on CommentsExtended %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCommentsExtended
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CommentsExtended and its children
func (m *CommentsExtended) Validate() error {
	return m.ValidateWithPath("CommentsExtended")
}

// ValidateWithPath validates the CommentsExtended and its children, prefixing error messages with path
func (m *CommentsExtended) ValidateWithPath(path string) error {
	if err := m.CT_CommentsEx.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}

