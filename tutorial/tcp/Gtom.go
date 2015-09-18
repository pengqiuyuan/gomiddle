// automatically generated, do not modify

package tcp

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Gtom struct {
	_tab flatbuffers.Table
}

func GetRootAsGtom(buf []byte, offset flatbuffers.UOffsetT) *Gtom {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Gtom{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Gtom) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Gtom) Content() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func GtomStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func GtomAddContent(builder *flatbuffers.Builder, content flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(content), 0) }
func GtomEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
