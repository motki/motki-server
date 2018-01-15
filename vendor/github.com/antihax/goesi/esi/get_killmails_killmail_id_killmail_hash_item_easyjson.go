// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson8e5c6700DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashItemList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetKillmailsKillmailIdKillmailHashItemList, 0, 2)
			} else {
				*out = GetKillmailsKillmailIdKillmailHashItemList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetKillmailsKillmailIdKillmailHashItem
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8e5c6700EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashItemList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashItemList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8e5c6700EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashItemList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8e5c6700EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashItemList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8e5c6700DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashItemList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8e5c6700DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson8e5c6700DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashItem) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "item_type_id":
			out.ItemTypeId = int32(in.Int32())
		case "quantity_destroyed":
			out.QuantityDestroyed = int64(in.Int64())
		case "quantity_dropped":
			out.QuantityDropped = int64(in.Int64())
		case "singleton":
			out.Singleton = int32(in.Int32())
		case "flag":
			out.Flag = int32(in.Int32())
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
func easyjson8e5c6700EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ItemTypeId != 0 {
		const prefix string = ",\"item_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ItemTypeId))
	}
	if in.QuantityDestroyed != 0 {
		const prefix string = ",\"quantity_destroyed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.QuantityDestroyed))
	}
	if in.QuantityDropped != 0 {
		const prefix string = ",\"quantity_dropped\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.QuantityDropped))
	}
	if in.Singleton != 0 {
		const prefix string = ",\"singleton\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Singleton))
	}
	if in.Flag != 0 {
		const prefix string = ",\"flag\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Flag))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8e5c6700EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8e5c6700EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8e5c6700DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8e5c6700DecodeGithubComAntihaxGoesiEsi1(l, v)
}
