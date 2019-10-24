// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbast

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Identifier struct {
	_tab flatbuffers.Table
}

func GetRootAsIdentifier(buf []byte, offset flatbuffers.UOffsetT) *Identifier {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Identifier{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Identifier) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Identifier) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Identifier) BaseNode(obj *BaseNode) *BaseNode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BaseNode)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Identifier) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func IdentifierStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func IdentifierAddBaseNode(builder *flatbuffers.Builder, baseNode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(baseNode), 0)
}
func IdentifierAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func IdentifierEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}