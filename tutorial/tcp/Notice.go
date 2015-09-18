// automatically generated, do not modify

package tcp

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Notice struct {
	_tab flatbuffers.Table
}

func GetRootAsNotice(buf []byte, offset flatbuffers.UOffsetT) *Notice {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Notice{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Notice) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Notice) Content() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func NoticeStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func NoticeAddContent(builder *flatbuffers.Builder, content flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(content), 0) }
func NoticeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
