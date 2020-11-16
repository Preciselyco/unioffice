package document

import (
	"github.com/Preciselyco/unioffice/schema/soo/wml"
	"github.com/Preciselyco/unioffice/schema/soo/ofc/sharedTypes"
)

// CommentExtended contains relationship information for a Comment.
type CommentExtended struct {
	x *wml.CT_CommentEx
}

// NewCommentExtended constructs a new empty CommentExtended.
func NewCommentExtended(doc *Document) CommentExtended {
	return CommentExtended{
		x: wml.NewCT_CommentEx(),
	}
}

// X returns the inner wrapped XML type.
func (ce CommentExtended) X() *wml.CT_CommentEx {
	return ce.x
}

// SetParagraphID sets the ID of the last paragraph in the Comment.
func (ce CommentExtended) SetParagraphID(id string) CommentExtended {
	ce.x.ParaIdAttr = id
	return ce
}

// SetParentParagraphID sets the ID of the last paragraph in the parent Comment.
func (ce CommentExtended) SetParentParagraphID(id string) CommentExtended {
	ce.x.ParaIdParentAttr = &id
	return ce
}

// SetDone sets the 'done' attribute.
func (ce CommentExtended) SetDone(done bool) CommentExtended {
	ce.x.DoneAttr = &sharedTypes.ST_OnOff{}
	ce.x.DoneAttr.Bool = &done
	return ce
}
