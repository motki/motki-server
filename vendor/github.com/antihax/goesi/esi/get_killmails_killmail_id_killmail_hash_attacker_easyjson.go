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

func easyjson3080782cDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashAttackerList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetKillmailsKillmailIdKillmailHashAttackerList, 0, 1)
			} else {
				*out = GetKillmailsKillmailIdKillmailHashAttackerList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetKillmailsKillmailIdKillmailHashAttacker
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
func easyjson3080782cEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashAttackerList) {
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
func (v GetKillmailsKillmailIdKillmailHashAttackerList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3080782cEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashAttackerList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3080782cEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashAttackerList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3080782cDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashAttackerList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3080782cDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson3080782cDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashAttacker) {
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
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "character_id":
			out.CharacterId = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "damage_done":
			out.DamageDone = int32(in.Int32())
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "final_blow":
			out.FinalBlow = bool(in.Bool())
		case "security_status":
			out.SecurityStatus = float32(in.Float32())
		case "ship_type_id":
			out.ShipTypeId = int32(in.Int32())
		case "weapon_type_id":
			out.WeaponTypeId = int32(in.Int32())
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
func easyjson3080782cEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashAttacker) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllianceId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance_id\":")
		out.Int32(int32(in.AllianceId))
	}
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
	if in.DamageDone != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"damage_done\":")
		out.Int32(int32(in.DamageDone))
	}
	if in.FactionId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"faction_id\":")
		out.Int32(int32(in.FactionId))
	}
	if in.FinalBlow {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"final_blow\":")
		out.Bool(bool(in.FinalBlow))
	}
	if in.SecurityStatus != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"security_status\":")
		out.Float32(float32(in.SecurityStatus))
	}
	if in.ShipTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ship_type_id\":")
		out.Int32(int32(in.ShipTypeId))
	}
	if in.WeaponTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"weapon_type_id\":")
		out.Int32(int32(in.WeaponTypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashAttacker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3080782cEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashAttacker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3080782cEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashAttacker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3080782cDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashAttacker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3080782cDecodeGithubComAntihaxGoesiEsi1(l, v)
}
