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

func easyjsonFc8244e9DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdWalletJournalExtraInfoList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdWalletJournalExtraInfoList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdWalletJournalExtraInfoList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdWalletJournalExtraInfo
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
func easyjsonFc8244e9EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdWalletJournalExtraInfoList) {
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
func (v GetCharactersCharacterIdWalletJournalExtraInfoList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFc8244e9EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdWalletJournalExtraInfoList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFc8244e9EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletJournalExtraInfoList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFc8244e9DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletJournalExtraInfoList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFc8244e9DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonFc8244e9DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdWalletJournalExtraInfo) {
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
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "transaction_id":
			out.TransactionId = int64(in.Int64())
		case "npc_name":
			out.NpcName = string(in.String())
		case "npc_id":
			out.NpcId = int32(in.Int32())
		case "destroyed_ship_type_id":
			out.DestroyedShipTypeId = int32(in.Int32())
		case "character_id":
			out.CharacterId = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "job_id":
			out.JobId = int32(in.Int32())
		case "contract_id":
			out.ContractId = int32(in.Int32())
		case "system_id":
			out.SystemId = int32(in.Int32())
		case "planet_id":
			out.PlanetId = int32(in.Int32())
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
func easyjsonFc8244e9EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdWalletJournalExtraInfo) {
	out.RawByte('{')
	first := true
	_ = first
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
	if in.TransactionId != 0 {
		const prefix string = ",\"transaction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.TransactionId))
	}
	if in.NpcName != "" {
		const prefix string = ",\"npc_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.NpcName))
	}
	if in.NpcId != 0 {
		const prefix string = ",\"npc_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.NpcId))
	}
	if in.DestroyedShipTypeId != 0 {
		const prefix string = ",\"destroyed_ship_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.DestroyedShipTypeId))
	}
	if in.CharacterId != 0 {
		const prefix string = ",\"character_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CharacterId))
	}
	if in.CorporationId != 0 {
		const prefix string = ",\"corporation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CorporationId))
	}
	if in.AllianceId != 0 {
		const prefix string = ",\"alliance_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.AllianceId))
	}
	if in.JobId != 0 {
		const prefix string = ",\"job_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.JobId))
	}
	if in.ContractId != 0 {
		const prefix string = ",\"contract_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ContractId))
	}
	if in.SystemId != 0 {
		const prefix string = ",\"system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SystemId))
	}
	if in.PlanetId != 0 {
		const prefix string = ",\"planet_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.PlanetId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdWalletJournalExtraInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFc8244e9EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdWalletJournalExtraInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFc8244e9EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletJournalExtraInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFc8244e9DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletJournalExtraInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFc8244e9DecodeGithubComAntihaxGoesiEsi1(l, v)
}
