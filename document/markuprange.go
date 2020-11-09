// Added by Precisely

package document

import (
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

type MarkupRange struct {
	d *Document
	x *wml.CT_MarkupRange
}

// X returns the inner wrapped XML type.
func (mr MarkupRange) X() *wml.CT_MarkupRange {
	return mr.x
}

// SetID sets the ID.
func (mr MarkupRange) SetID(id int64) MarkupRange {
	mr.x.IdAttr = id
	return mr
}
