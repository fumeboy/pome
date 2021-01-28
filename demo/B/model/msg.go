package model

type Msg struct {
	Email     string
	Content   string
	Timestamp int64
}

var fakeDatabase = &fakeDatabaseTyp{}
type fakeDatabaseTyp struct {
	msgList []*Msg
}

func (l *fakeDatabaseTyp) add(msg *Msg) (err error) {
	l.msgList = append(l.msgList, msg)
	return
}

func (l *fakeDatabaseTyp) get(offset, limit uint32) (result []*Msg, err error) {
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

func AddMsg(msg *Msg) error {
	return fakeDatabase.add(msg)
}

func GetMsg(offset, limit uint32) (result []*Msg, err error) {
	return fakeDatabase.get(offset, limit)
}
