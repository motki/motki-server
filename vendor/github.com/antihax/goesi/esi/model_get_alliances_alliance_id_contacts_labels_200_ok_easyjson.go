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

func easyjson7c9005b1DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetAlliancesAllianceIdContactsLabels200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetAlliancesAllianceIdContactsLabels200OkList, 0, 2)
			} else {
				*out = GetAlliancesAllianceIdContactsLabels200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetAlliancesAllianceIdContactsLabels200Ok
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
func easyjson7c9005b1EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetAlliancesAllianceIdContactsLabels200OkList) {
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
func (v GetAlliancesAllianceIdContactsLabels200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7c9005b1EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdContactsLabels200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7c9005b1EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdContactsLabels200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7c9005b1DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdContactsLabels200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7c9005b1DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson7c9005b1DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetAlliancesAllianceIdContactsLabels200Ok) {
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
		case "label_id":
			out.LabelId = int64(in.Int64())
		case "label_name":
			out.LabelName = string(in.String())
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
func easyjson7c9005b1EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetAlliancesAllianceIdContactsLabels200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LabelId != 0 {
		const prefix string = ",\"label_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LabelId))
	}
	if in.LabelName != "" {
		const prefix string = ",\"label_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LabelName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetAlliancesAllianceIdContactsLabels200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7c9005b1EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdContactsLabels200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7c9005b1EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdContactsLabels200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7c9005b1DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdContactsLabels200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7c9005b1DecodeGithubComAntihaxGoesiEsi1(l, v)
}
