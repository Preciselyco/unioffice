// Added by Precisely

package document

import (
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

type Markup struct {
	d *Document
	x *wml.CT_Markup
}

// X returns the inner wrapped XML type.
func (m Markup) X() *wml.CT_Markup {
	return m.x
}

// SetID sets the ID.
func (m Markup) SetID(id int64) Markup {
	m.x.IdAttr = id
	return m
}
