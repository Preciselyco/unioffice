package document

import (
	"time"
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// RunTrackChange is used for marking text as inserted or deleted.
type RunTrackChange struct {
	d *Document
	x *wml.CT_RunTrackChange
}

// RunTrackChangeMode is used when inserting a RunTrackChange.
type RunTrackChangeMode int
const (
	RunTrackChangeInsert RunTrackChangeMode = iota
	RunTrackChangeDelete
)

// X returns the inner wrapped XML type.
func (rtc RunTrackChange) X() *wml.CT_RunTrackChange {
	return rtc.x
}

// AddRun adds a run of inserted or deleted text.
func (rtc RunTrackChange) AddRun() Run {
	crc := wml.NewEG_ContentRunContent()
	crc.R = wml.NewCT_R()
	rtc.x.EG_ContentRunContent = append(rtc.x.EG_ContentRunContent, crc)
	return Run{rtc.d, crc.R}
}

// SetAuthor sets the author's name.
func (rtc RunTrackChange) SetAuthor(author string) RunTrackChange {
	rtc.x.AuthorAttr = author
	return rtc
}

// SetDate sets the date.
func (rtc RunTrackChange) SetDate(date time.Time) RunTrackChange {
	rtc.x.DateAttr = &date
	return rtc
}

// SetID sets the annotation ID.
func (rtc RunTrackChange) SetID(id int64) RunTrackChange {
	rtc.x.IdAttr = id
	return rtc
}
