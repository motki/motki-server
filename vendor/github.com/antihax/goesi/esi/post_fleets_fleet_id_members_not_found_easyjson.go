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

func easyjsonE5dab29eDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PostFleetsFleetIdMembersNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostFleetsFleetIdMembersNotFoundList, 0, 4)
			} else {
				*out = PostFleetsFleetIdMembersNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostFleetsFleetIdMembersNotFound
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
func easyjsonE5dab29eEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PostFleetsFleetIdMembersNotFoundList) {
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
func (v PostFleetsFleetIdMembersNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE5dab29eEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFleetsFleetIdMembersNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE5dab29eEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFleetsFleetIdMembersNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE5dab29eDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFleetsFleetIdMembersNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE5dab29eDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonE5dab29eDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PostFleetsFleetIdMembersNotFound) {
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
		case "error":
			out.Error_ = string(in.String())
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
func easyjsonE5dab29eEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PostFleetsFleetIdMembersNotFound) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostFleetsFleetIdMembersNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE5dab29eEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFleetsFleetIdMembersNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE5dab29eEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFleetsFleetIdMembersNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE5dab29eDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFleetsFleetIdMembersNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE5dab29eDecodeGithubComAntihaxGoesiEsi1(l, v)
}
