// automatically generated, do not modify

package tcp

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Mtog struct {
	_tab flatbuffers.Table
}

func GetRootAsMtog(buf []byte, offset flatbuffers.UOffsetT) *Mtog {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Mtog{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Mtog) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Mtog) ServerZoneId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Mtog) ServerId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Mtog) GameId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Mtog) PageNumber() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Mtog) PageSize() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Mtog) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Mtog) ItemId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MtogStart(builder *flatbuffers.Builder) { builder.StartObject(7) }
func MtogAddServerZoneId(builder *flatbuffers.Builder, serverZoneId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(serverZoneId), 0) }
func MtogAddServerId(builder *flatbuffers.Builder, serverId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(serverId), 0) }
func MtogAddGameId(builder *flatbuffers.Builder, gameId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(gameId), 0) }
func MtogAddPageNumber(builder *flatbuffers.Builder, pageNumber flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(pageNumber), 0) }
func MtogAddPageSize(builder *flatbuffers.Builder, pageSize flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(pageSize), 0) }
func MtogAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(id), 0) }
func MtogAddItemId(builder *flatbuffers.Builder, itemId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(itemId), 0) }
func MtogEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
