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

func easyjsonEbea04a9DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashVictimList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetKillmailsKillmailIdKillmailHashVictimList, 0, 1)
			} else {
				*out = GetKillmailsKillmailIdKillmailHashVictimList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetKillmailsKillmailIdKillmailHashVictim
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
func easyjsonEbea04a9EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashVictimList) {
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
func (v GetKillmailsKillmailIdKillmailHashVictimList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbea04a9EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashVictimList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbea04a9EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashVictimList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbea04a9DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashVictimList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbea04a9DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonEbea04a9DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashVictim) {
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
		case "character_id":
			out.CharacterId = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "damage_taken":
			out.DamageTaken = int32(in.Int32())
		case "ship_type_id":
			out.ShipTypeId = int32(in.Int32())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]GetKillmailsKillmailIdKillmailHashItem1, 0, 1)
					} else {
						out.Items = []GetKillmailsKillmailIdKillmailHashItem1{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetKillmailsKillmailIdKillmailHashItem1
					easyjsonEbea04a9DecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "position":
			(out.Position).UnmarshalEasyJSON(in)
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
func easyjsonEbea04a9EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashVictim) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CharacterId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"character_id\":")
		out.Int32(int32(in.CharacterId))
	}
	if in.CorporationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"corporation_id\":")
		out.Int32(int32(in.CorporationId))
	}
	if in.AllianceId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance_id\":")
		out.Int32(int32(in.AllianceId))
	}
	if in.FactionId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"faction_id\":")
		out.Int32(int32(in.FactionId))
	}
	if in.DamageTaken != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"damage_taken\":")
		out.Int32(int32(in.DamageTaken))
	}
	if in.ShipTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ship_type_id\":")
		out.Int32(int32(in.ShipTypeId))
	}
	if len(in.Items) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"items\":")
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Items {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjsonEbea04a9EncodeGithubComAntihaxGoesiEsi2(out, v6)
			}
			out.RawByte(']')
		}
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"position\":")
		(in.Position).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashVictim) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbea04a9EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashVictim) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbea04a9EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashVictim) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbea04a9DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashVictim) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbea04a9DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjsonEbea04a9DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashItem1) {
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
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]GetKillmailsKillmailIdKillmailHashItem, 0, 2)
					} else {
						out.Items = []GetKillmailsKillmailIdKillmailHashItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v7 GetKillmailsKillmailIdKillmailHashItem
					(v7).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v7)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonEbea04a9EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashItem1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ItemTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"item_type_id\":")
		out.Int32(int32(in.ItemTypeId))
	}
	if in.QuantityDestroyed != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quantity_destroyed\":")
		out.Int64(int64(in.QuantityDestroyed))
	}
	if in.QuantityDropped != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quantity_dropped\":")
		out.Int64(int64(in.QuantityDropped))
	}
	if in.Singleton != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"singleton\":")
		out.Int32(int32(in.Singleton))
	}
	if in.Flag != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"flag\":")
		out.Int32(int32(in.Flag))
	}
	if len(in.Items) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"items\":")
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Items {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
