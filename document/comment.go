// Added by Precisely

package document

import (
	"time"

	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// Comment is a comment contained in comments.xml.
type Comment struct {
	d *Document
	x *wml.CT_Comment
}

// NewComment constructs a new empty Comment.
func NewComment(doc *Document) Comment {
	return Comment{
		d: doc,
		x: wml.NewCT_Comment(),
	}
}

// X returns the inner wrapped XML type.
func (comment Comment) X() *wml.CT_Comment {
	return comment.x
}

// AddParagraph adds a paragraph of text to the comment.
func (comment Comment) AddParagraph() Paragraph {
	ble := wml.NewEG_BlockLevelElts()
	cbc := wml.NewEG_ContentBlockContent()
	p := wml.NewCT_P()
	cbc.P = append(cbc.P, p)
	ble.EG_ContentBlockContent = append(ble.EG_ContentBlockContent, cbc)
	comment.x.EG_BlockLevelElts = append(comment.x.EG_BlockLevelElts, ble)
	return Paragraph{comment.d, p}
}

// SetAuthor sets the author's name.
func (comment Comment) SetAuthor(author string) Comment {
	comment.x.AuthorAttr = author
	return comment
}

// SetDate sets the creation date.
func (comment Comment) SetDate(date time.Time) Comment {
	comment.x.DateAttr = &date
	return comment
}

// SetID sets the comment ID.
func (comment Comment) SetID(id int64) Comment {
	comment.x.IdAttr = id
	return comment
}

// SetInitials sets the author's initials.
func (comment Comment) SetInitials(initials string) Comment {
	comment.x.InitialsAttr = &initials
	return comment
}
