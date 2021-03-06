// This file is not autogenerated, but based on CT_Comment.go
// Schema: https://docs.microsoft.com/en-us/openspecs/office_standards/ms-docx/d416013d-c112-44fa-8bef-7819b8898117

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/Preciselyco/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_CommentEx struct {
	ParaIdAttr       string  // Should actually be ST_LongHexNumber according to schema
	ParaIdParentAttr *string // Should actually be ST_LongHexNumber according to schema
	DoneAttr         *sharedTypes.ST_OnOff
}

func NewCT_CommentEx() *CT_CommentEx {
	ret := &CT_CommentEx{}
	return ret
}

func (m *CT_CommentEx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w12:paraId"}, Value: m.ParaIdAttr})
	if m.ParaIdParentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w12:paraIdParent"}, Value: *m.ParaIdParentAttr})
	}
	if m.DoneAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w12:done"}, Value: fmt.Sprintf("%v", *m.DoneAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CommentEx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		switch attr.Name {
		case xml.Name{Space: "http://schemas.microsoft.com/office/word/2012/wordml", Local: "paraId"}:
			m.ParaIdAttr = attr.Value
		case xml.Name{Space: "http://schemas.microsoft.com/office/word/2012/wordml", Local: "paraIdParent"}:
			// Make a local copy
			value := attr.Value
			m.ParaIdParentAttr = &value
		case xml.Name{Space: "http://schemas.microsoft.com/office/word/2012/wordml", Local: "done"}:
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DoneAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CommentEx: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CommentEx and its children
func (m *CT_CommentEx) Validate() error {
	return m.ValidateWithPath("CT_CommentEx")
}

// ValidateWithPath validates the CT_CommentEx and its children, prefixing error messages with path
func (m *CT_CommentEx) ValidateWithPath(path string) error {
	if m.DoneAttr != nil {
		if err := m.DoneAttr.ValidateWithPath(path + "/DoneAttr"); err != nil {
			return err
		}
	}
	return nil
}
