package document

import (
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// CommentsExtended contains relationship information for the comments attached to a document.
type CommentsExtended struct {
	x *wml.CommentsExtended
}

// NewCommentsExtended constructs a new empty CommentsExtended.
func NewCommentsExtended() CommentsExtended {
	return CommentsExtended{
		x: wml.NewCommentsExtended(),
	}
}

// X returns the inner wrapped XML type.
func (ce CommentsExtended) X() *wml.CommentsExtended {
	return ce.x
}

// Clear clears the contents.
func (ce CommentsExtended) Clear() {
	ce.x = wml.NewCommentsExtended()
}

// AddCommentExtended adds a new empty CommentExtended.
func (ce CommentsExtended) AddCommentExtended() CommentExtended {
	cex := wml.NewCT_CommentEx()
	ce.x.CommentEx = append(ce.x.CommentEx, cex)
	return CommentExtended{x: cex}
}

// CommentsExtended returns the contents.
func (ce CommentsExtended) CommentsExtended() []CommentExtended {
	ret := []CommentExtended{}
	for _, cex := range ce.x.CommentEx {
		ret = append(ret, CommentExtended{x: cex})
	}
	return ret
}

// NonEmpty returns true if the contents is not empty.
func (ce CommentsExtended) NonEmpty() bool {
	return len(ce.x.CommentEx) > 0
}
