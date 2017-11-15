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

func easyjson29139c1DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PutFleetsFleetIdNewSettingsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PutFleetsFleetIdNewSettingsList, 0, 2)
			} else {
				*out = PutFleetsFleetIdNewSettingsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PutFleetsFleetIdNewSettings
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
func easyjson29139c1EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PutFleetsFleetIdNewSettingsList) {
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
func (v PutFleetsFleetIdNewSettingsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29139c1EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PutFleetsFleetIdNewSettingsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29139c1EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PutFleetsFleetIdNewSettingsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29139c1DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PutFleetsFleetIdNewSettingsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29139c1DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson29139c1DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PutFleetsFleetIdNewSettings) {
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
		case "motd":
			out.Motd = string(in.String())
		case "is_free_move":
			out.IsFreeMove = bool(in.Bool())
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
func easyjson29139c1EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PutFleetsFleetIdNewSettings) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Motd != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"motd\":")
		out.String(string(in.Motd))
	}
	if in.IsFreeMove {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_free_move\":")
		out.Bool(bool(in.IsFreeMove))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PutFleetsFleetIdNewSettings) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29139c1EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PutFleetsFleetIdNewSettings) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29139c1EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PutFleetsFleetIdNewSettings) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29139c1DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PutFleetsFleetIdNewSettings) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29139c1DecodeGithubComAntihaxGoesiEsi1(l, v)
}
