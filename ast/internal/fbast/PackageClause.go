// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbast

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PackageClause struct {
	_tab flatbuffers.Table
}

func GetRootAsPackageClause(buf []byte, offset flatbuffers.UOffsetT) *PackageClause {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PackageClause{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PackageClause) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PackageClause) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PackageClause) BaseNode(obj *BaseNode) *BaseNode {
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

func (rcv *PackageClause) Name(obj *Identifier) *Identifier {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Identifier)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func PackageClauseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PackageClauseAddBaseNode(builder *flatbuffers.Builder, baseNode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(baseNode), 0)
}
func PackageClauseAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func PackageClauseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}