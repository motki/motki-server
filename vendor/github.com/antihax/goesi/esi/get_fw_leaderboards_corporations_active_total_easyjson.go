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

func easyjson81d4fc9bDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsActiveTotalList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsCorporationsActiveTotalList, 0, 8)
			} else {
				*out = GetFwLeaderboardsCorporationsActiveTotalList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsCorporationsActiveTotal
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
func easyjson81d4fc9bEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsCorporationsActiveTotalList) {
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
func (v GetFwLeaderboardsCorporationsActiveTotalList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson81d4fc9bEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCorporationsActiveTotalList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson81d4fc9bEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsActiveTotalList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson81d4fc9bDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsActiveTotalList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson81d4fc9bDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson81d4fc9bDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsActiveTotal) {
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
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "amount":
			out.Amount = int32(in.Int32())
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
func easyjson81d4fc9bEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsCorporationsActiveTotal) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CorporationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"corporation_id\":")
		out.Int32(int32(in.CorporationId))
	}
	if in.Amount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"amount\":")
		out.Int32(int32(in.Amount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsCorporationsActiveTotal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson81d4fc9bEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCorporationsActiveTotal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson81d4fc9bEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsActiveTotal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson81d4fc9bDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsActiveTotal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson81d4fc9bDecodeGithubComAntihaxGoesiEsi1(l, v)
}
