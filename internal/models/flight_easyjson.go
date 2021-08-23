// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson2028b946DecodeRsoiKpKTLHInternalModels(in *jlexer.Lexer, out *Flight) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "from":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.From).UnmarshalText(data))
			}
		case "to":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.To).UnmarshalText(data))
			}
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2028b946EncodeRsoiKpKTLHInternalModels(out *jwriter.Writer, in Flight) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix[1:])
		out.RawText((in.From).MarshalText())
	}
	{
		const prefix string = ",\"to\":"
		out.RawString(prefix)
		out.RawText((in.To).MarshalText())
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Flight) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2028b946EncodeRsoiKpKTLHInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Flight) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2028b946EncodeRsoiKpKTLHInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Flight) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2028b946DecodeRsoiKpKTLHInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Flight) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2028b946DecodeRsoiKpKTLHInternalModels(l, v)
}
