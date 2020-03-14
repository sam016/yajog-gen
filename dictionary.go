package gogen

import "github.com/francoispqt/gojay"

// PlainDictionary stores foo values
type PlainDictionary struct {
	data map[string]interface{}
}

// UnMarDictionary stores foo values which implements
// the MarshalerJSONArray, MarshalerJSONObject, UnmarshalerJSONArray, UnmarshalerJSONObject
type UnMarDictionary struct {
	data map[string]interface{}
}

func (d *UnMarDictionary) UnmarshalJSONObject(*gojay.Decoder, string) error { return nil }   // UnmarshalerJSONObject
func (d *UnMarDictionary) NKeys() int                                       { return 0 }     // UnmarshalerJSONObject
func (d *UnMarDictionary) UnmarshalJSONArray(*gojay.Decoder) error          { return nil }   // UnmarshalerJSONArray
func (d *UnMarDictionary) MarshalJSONObject(enc *gojay.Encoder)             {}               // MarshalerJSONObject
func (d *UnMarDictionary) IsNil() bool                                      { return false } // MarshalerJSONObject
func (d *UnMarDictionary) MarshalJSONArray(enc *gojay.Encoder)              {}               // MarshalerJSONArray
//func (d *UnMarDictionary) IsNil() bool                                    { return false } // MarshalerJSONArray

// PlainClassroom using PlainDictionary under the hood
type PlainClassroom struct {
	ID       int
	Students PlainDictionary
}

// UnMarClassroom using UnMarDictionary under the hood
type UnMarClassroom struct {
	ID       int
	Students UnMarDictionary
}
