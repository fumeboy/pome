package model

import (
	"context"
)

type Msg struct {
	Email     string
	Content   string
	Timestamp int64
}

var tempDatabase = &tempDatabaseT{}
type tempDatabaseT struct {
	msgList []*Msg
}

func (l *tempDatabaseT) add(ctx context.Context, msg *Msg) (err error) {
	l.msgList = append(l.msgList, msg)
	return
}

func (l *tempDatabaseT) get(ctx context.Context, offset, limit uint32) (result []*Msg, err error) {
	if offset < 0 || limit <= 0 {
		return
	}
	if offset >= uint32(len(l.msgList)) {
		return
	}
	end := offset+limit
	if end > uint32(len(l.msgList)) {
		end = uint32(len(l.msgList))
	}
	result = l.msgList[offset : end]
	return
}

func AddMsg(ctx context.Context, msg *Msg) error {
	return tempDatabase.add(ctx, msg)
}

func GetMsg(ctx context.Context, offset, limit uint32) (result []*Msg, err error) {
	return tempDatabase.get(ctx, offset, limit)
}
