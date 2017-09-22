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

func easyjsonDb6fb9f8DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetIndustrySystems200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetIndustrySystems200OkList, 0, 2)
			} else {
				*out = GetIndustrySystems200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetIndustrySystems200Ok
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
func easyjsonDb6fb9f8EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetIndustrySystems200OkList) {
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
func (v GetIndustrySystems200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDb6fb9f8EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetIndustrySystems200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDb6fb9f8EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetIndustrySystems200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDb6fb9f8DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetIndustrySystems200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDb6fb9f8DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonDb6fb9f8DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetIndustrySystems200Ok) {
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
		case "cost_indices":
			if in.IsNull() {
				in.Skip()
				out.CostIndices = nil
			} else {
				in.Delim('[')
				if out.CostIndices == nil {
					if !in.IsDelim(']') {
						out.CostIndices = make([]GetIndustrySystemsCostIndice, 0, 2)
					} else {
						out.CostIndices = []GetIndustrySystemsCostIndice{}
					}
				} else {
					out.CostIndices = (out.CostIndices)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetIndustrySystemsCostIndice
					easyjsonDb6fb9f8DecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.CostIndices = append(out.CostIndices, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "solar_system_id":
			out.SolarSystemId = int32(in.Int32())
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
func easyjsonDb6fb9f8EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetIndustrySystems200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.CostIndices) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cost_indices\":")
		if in.CostIndices == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.CostIndices {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjsonDb6fb9f8EncodeGithubComAntihaxGoesiEsi2(out, v6)
			}
			out.RawByte(']')
		}
	}
	if in.SolarSystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"solar_system_id\":")
		out.Int32(int32(in.SolarSystemId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetIndustrySystems200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDb6fb9f8EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetIndustrySystems200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDb6fb9f8EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetIndustrySystems200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDb6fb9f8DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetIndustrySystems200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDb6fb9f8DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjsonDb6fb9f8DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetIndustrySystemsCostIndice) {
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
		case "activity":
			out.Activity = string(in.String())
		case "cost_index":
			out.CostIndex = float32(in.Float32())
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
func easyjsonDb6fb9f8EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetIndustrySystemsCostIndice) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Activity != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"activity\":")
		out.String(string(in.Activity))
	}
	if in.CostIndex != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cost_index\":")
		out.Float32(float32(in.CostIndex))
	}
	out.RawByte('}')
}
