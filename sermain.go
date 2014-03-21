package main

import (
	"github.com/glycerine/go-capnproto"
	"github.com/glycerine/serlist"
)

func main() {
	seg := capn.NewBuffer(nil)
	list := serlist.NewMPGList(seg, 2)
	m1 := serlist.NewMPG(seg)
	//m2 := serlist.NewMPG(seg)

	pl := capn.PointerList(list)
	pl.Set(0, capn.Object(m1))
	//pl.Set(0, capn.Object(m2))

}
