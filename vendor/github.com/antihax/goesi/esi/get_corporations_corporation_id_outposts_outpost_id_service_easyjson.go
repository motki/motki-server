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

func easyjson76d3c00eDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdOutpostsOutpostIdServiceList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdOutpostsOutpostIdServiceList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdOutpostsOutpostIdServiceList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdOutpostsOutpostIdService
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
func easyjson76d3c00eEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdOutpostsOutpostIdServiceList) {
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
func (v GetCorporationsCorporationIdOutpostsOutpostIdServiceList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson76d3c00eEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdOutpostsOutpostIdServiceList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson76d3c00eEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdOutpostsOutpostIdServiceList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson76d3c00eDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdOutpostsOutpostIdServiceList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson76d3c00eDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson76d3c00eDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdOutpostsOutpostIdService) {
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
		case "service_name":
			out.ServiceName = string(in.String())
		case "minimum_standing":
			out.MinimumStanding = float64(in.Float64())
		case "surcharge_per_bad_standing":
			out.SurchargePerBadStanding = float64(in.Float64())
		case "discount_per_good_standing":
			out.DiscountPerGoodStanding = float64(in.Float64())
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
func easyjson76d3c00eEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdOutpostsOutpostIdService) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ServiceName != "" {
		const prefix string = ",\"service_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ServiceName))
	}
	if in.MinimumStanding != 0 {
		const prefix string = ",\"minimum_standing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.MinimumStanding))
	}
	if in.SurchargePerBadStanding != 0 {
		const prefix string = ",\"surcharge_per_bad_standing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.SurchargePerBadStanding))
	}
	if in.DiscountPerGoodStanding != 0 {
		const prefix string = ",\"discount_per_good_standing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.DiscountPerGoodStanding))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdOutpostsOutpostIdService) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson76d3c00eEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdOutpostsOutpostIdService) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson76d3c00eEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdOutpostsOutpostIdService) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson76d3c00eDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdOutpostsOutpostIdService) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson76d3c00eDecodeGithubComAntihaxGoesiEsi1(l, v)
}
