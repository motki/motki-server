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

func easyjsonAc6a9211DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdAssets200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdAssets200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdAssets200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdAssets200Ok
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
func easyjsonAc6a9211EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdAssets200OkList) {
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
func (v GetCorporationsCorporationIdAssets200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAc6a9211EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdAssets200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAc6a9211EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdAssets200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAc6a9211DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdAssets200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAc6a9211DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonAc6a9211DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdAssets200Ok) {
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
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "quantity":
			out.Quantity = int32(in.Int32())
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "location_type":
			out.LocationType = string(in.String())
		case "item_id":
			out.ItemId = int64(in.Int64())
		case "location_flag":
			out.LocationFlag = string(in.String())
		case "is_singleton":
			out.IsSingleton = bool(in.Bool())
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
func easyjsonAc6a9211EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdAssets200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.TypeId != 0 {
		const prefix string = ",\"type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TypeId))
	}
	if in.Quantity != 0 {
		const prefix string = ",\"quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Quantity))
	}
	if in.LocationId != 0 {
		const prefix string = ",\"location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LocationId))
	}
	if in.LocationType != "" {
		const prefix string = ",\"location_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LocationType))
	}
	if in.ItemId != 0 {
		const prefix string = ",\"item_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ItemId))
	}
	if in.LocationFlag != "" {
		const prefix string = ",\"location_flag\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LocationFlag))
	}
	if in.IsSingleton {
		const prefix string = ",\"is_singleton\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsSingleton))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdAssets200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAc6a9211EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdAssets200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAc6a9211EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdAssets200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAc6a9211DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdAssets200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAc6a9211DecodeGithubComAntihaxGoesiEsi1(l, v)
}
