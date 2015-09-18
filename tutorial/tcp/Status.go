// automatically generated, do not modify

package tcp

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Status struct {
	_tab flatbuffers.Table
}

func GetRootAsStatus(buf []byte, offset flatbuffers.UOffsetT) *Status {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Status{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Status) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Status) ServerZoneId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Status) PlatForm(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j * 4))
	}
	return nil
}

func (rcv *Status) PlatFormLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Status) ServerId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Status) GameId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Status) Ip() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Status) Port() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Status) Status() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func StatusStart(builder *flatbuffers.Builder) { builder.StartObject(7) }
func StatusAddServerZoneId(builder *flatbuffers.Builder, serverZoneId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(serverZoneId), 0) }
func StatusAddPlatForm(builder *flatbuffers.Builder, platForm flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(platForm), 0) }
func StatusStartPlatFormVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func StatusAddServerId(builder *flatbuffers.Builder, serverId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(serverId), 0) }
func StatusAddGameId(builder *flatbuffers.Builder, gameId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(gameId), 0) }
func StatusAddIp(builder *flatbuffers.Builder, ip flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(ip), 0) }
func StatusAddPort(builder *flatbuffers.Builder, port uint16) { builder.PrependUint16Slot(5, port, 0) }
func StatusAddStatus(builder *flatbuffers.Builder, status int32) { builder.PrependInt32Slot(6, status, 0) }
func StatusEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
