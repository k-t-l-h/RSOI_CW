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

func easyjson2028b946DecodeRSOICWInternalModels(in *jlexer.Lexer, out *Flight) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "from":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.From).UnmarshalText(data))
			}
		case "from_city":
			out.FromCity = string(in.String())
		case "to":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.To).UnmarshalText(data))
			}
		case "to_city":
			out.ToCity = string(in.String())
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
func easyjson2028b946EncodeRSOICWInternalModels(out *jwriter.Writer, in Flight) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix)
		out.RawText((in.From).MarshalText())
	}
	if in.FromCity != "" {
		const prefix string = ",\"from_city\":"
		out.RawString(prefix)
		out.String(string(in.FromCity))
	}
	{
		const prefix string = ",\"to\":"
		out.RawString(prefix)
		out.RawText((in.To).MarshalText())
	}
	if in.ToCity != "" {
		const prefix string = ",\"to_city\":"
		out.RawString(prefix)
		out.String(string(in.ToCity))
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
	easyjson2028b946EncodeRSOICWInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Flight) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2028b946EncodeRSOICWInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Flight) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2028b946DecodeRSOICWInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Flight) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2028b946DecodeRSOICWInternalModels(l, v)
}
