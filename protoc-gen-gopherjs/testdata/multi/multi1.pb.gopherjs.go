// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: multi/multi1.proto

/*
Package multitest is a generated protocol buffer package.

It is generated from these files:
	multi/multi1.proto
	multi/multi2.proto
	multi/multi3.proto

It has these top-level messages:
	Multi1
	Multi2
	Multi3
*/
package multitest

import js "github.com/gopherjs/gopherjs/js"
import jspb "github.com/johanbrandhorst/protobuf/jspb"

type Multi1 struct {
	*js.Object
}

// GetMulti2 gets the Multi2 of the Multi1.
func (m *Multi1) GetMulti2() *Multi2 {
	return &Multi2{Object: m.Call("getMulti2")}
}

// SetMulti2 sets the Multi2 of the Multi1.
func (m *Multi1) SetMulti2(v *Multi2) {
	m.Call("setMulti2", v.Call("toArray"))
}

// HasMulti2 indicates whether the Multi2 of the Multi1 is set.
func (m *Multi1) HasMulti2() bool {
	return m.Call("hasMulti2").Bool()
}

// ClearMulti2 clears the Multi2 of the Multi1.
func (m *Multi1) ClearMulti2() {
	m.Call("clearMulti2")
}

// GetColor gets the Color of the Multi1.
func (m *Multi1) GetColor() Multi2_Color {
	return Multi2_Color(m.Call("getColor").Int())
}

// SetColor sets the Color of the Multi1.
func (m *Multi1) SetColor(v Multi2_Color) {
	m.Call("setColor", v)
}

// GetHatType gets the HatType of the Multi1.
func (m *Multi1) GetHatType() Multi3_HatType {
	return Multi3_HatType(m.Call("getHatType").Int())
}

// SetHatType sets the HatType of the Multi1.
func (m *Multi1) SetHatType(v Multi3_HatType) {
	m.Call("setHatType", v)
}

// New creates a new Multi1.
func (m *Multi1) New(multi2 *Multi2, color Multi2_Color, hatType Multi3_HatType) *Multi1 {
	m = &Multi1{
		Object: js.Global.Get("proto").Get("multitest").Get("Multi1").New([]interface{}{
			multi2.Call("toArray"),
			color,
			hatType,
		}),
	}

	return m
}

// Serialize marshals Multi1 to a slice of bytes.
func (m *Multi1) Serialize() ([]byte, error) {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a Multi1 from a slice of bytes.
func (m *Multi1) Deserialize(rawBytes []byte) (*Multi1, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("multitest").Get("Multi1"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Multi1{
		Object: obj,
	}, nil
}
