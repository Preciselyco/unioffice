package document

import (
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// Markup is used for anchoring a comment in a document.
type Markup struct {
	d *Document
	x *wml.CT_Markup
}

// X returns the inner wrapped XML type.
func (m Markup) X() *wml.CT_Markup {
	return m.x
}

// SetID sets the comment ID.
func (m Markup) SetID(id int64) Markup {
	m.x.IdAttr = id
	return m
}
