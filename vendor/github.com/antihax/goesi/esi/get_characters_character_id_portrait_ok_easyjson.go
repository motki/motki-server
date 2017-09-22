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

func easyjsonFa80e40cDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPortraitOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPortraitOkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdPortraitOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPortraitOk
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
func easyjsonFa80e40cEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPortraitOkList) {
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
func (v GetCharactersCharacterIdPortraitOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa80e40cEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPortraitOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa80e40cEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPortraitOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa80e40cDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPortraitOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa80e40cDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonFa80e40cDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPortraitOk) {
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
		case "px128x128":
			out.Px128x128 = string(in.String())
		case "px256x256":
			out.Px256x256 = string(in.String())
		case "px512x512":
			out.Px512x512 = string(in.String())
		case "px64x64":
			out.Px64x64 = string(in.String())
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
func easyjsonFa80e40cEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPortraitOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Px128x128 != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"px128x128\":")
		out.String(string(in.Px128x128))
	}
	if in.Px256x256 != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"px256x256\":")
		out.String(string(in.Px256x256))
	}
	if in.Px512x512 != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"px512x512\":")
		out.String(string(in.Px512x512))
	}
	if in.Px64x64 != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"px64x64\":")
		out.String(string(in.Px64x64))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPortraitOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa80e40cEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPortraitOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa80e40cEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPortraitOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa80e40cDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPortraitOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa80e40cDecodeGithubComAntihaxGoesiEsi1(l, v)
}
