// Added by Precisely

package document

import (
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// Comments are the comments contained in comments.xml.
type Comments struct {
	d *Document
	x *wml.Comments
}

// NewComments constructs a new empty Comments.
func NewComments(doc *Document) Comments {
	return Comments{
		d: doc,
		x: wml.NewComments(),
	}
}

// X returns the inner wrapped XML type.
func (comments Comments) X() *wml.Comments {
	return comments.x
}

// Clear clears the comments.
func (comments Comments) Clear() {
	comments.x = wml.NewComments()
}

// AddComment adds a new empty comment.
func (comments Comments) AddComment() Comment {
	cx := wml.NewCT_Comment()
	comments.x.Comment = append(comments.x.Comment, cx)
	return Comment{d: comments.d, x: cx}
}

// Comments returns all comments.
func (comments Comments) Comments() []Comment {
	ret := []Comment{}
	for _, cx := range comments.x.Comment {
		ret = append(ret, Comment{d: comments.d, x: cx})
	}
	return ret
}

// NonEmpty returns true if at least one comment exists.
func (comments Comments) NonEmpty() bool {
	return len(comments.x.Comment) > 0
}
