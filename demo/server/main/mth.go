package main

import "github.com/fumeboy/pome/demo/server/main/guestbook"

type serverT struct{}
type methodT struct {}
var serverIns guestbook.GuestBookServiceServer = &serverT{}
