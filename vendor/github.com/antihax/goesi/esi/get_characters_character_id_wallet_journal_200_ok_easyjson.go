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

func easyjson46989833DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdWalletJournal200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdWalletJournal200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdWalletJournal200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdWalletJournal200Ok
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
func easyjson46989833EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdWalletJournal200OkList) {
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
func (v GetCharactersCharacterIdWalletJournal200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson46989833EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdWalletJournal200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson46989833EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletJournal200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson46989833DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletJournal200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson46989833DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson46989833DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdWalletJournal200Ok) {
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
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		case "ref_id":
			out.RefId = int64(in.Int64())
		case "ref_type":
			out.RefType = string(in.String())
		case "first_party_id":
			out.FirstPartyId = int32(in.Int32())
		case "first_party_type":
			out.FirstPartyType = string(in.String())
		case "second_party_id":
			out.SecondPartyId = int32(in.Int32())
		case "second_party_type":
			out.SecondPartyType = string(in.String())
		case "amount":
			out.Amount = float32(in.Float32())
		case "balance":
			out.Balance = float32(in.Float32())
		case "reason":
			out.Reason = string(in.String())
		case "tax_reciever_id":
			out.TaxRecieverId = int32(in.Int32())
		case "tax":
			out.Tax = float32(in.Float32())
		case "extra_info":
			easyjson46989833DecodeGithubComAntihaxGoesiEsi2(in, &out.ExtraInfo)
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
func easyjson46989833EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdWalletJournal200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"date\":")
		out.Raw((in.Date).MarshalJSON())
	}
	if in.RefId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ref_id\":")
		out.Int64(int64(in.RefId))
	}
	if in.RefType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ref_type\":")
		out.String(string(in.RefType))
	}
	if in.FirstPartyId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"first_party_id\":")
		out.Int32(int32(in.FirstPartyId))
	}
	if in.FirstPartyType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"first_party_type\":")
		out.String(string(in.FirstPartyType))
	}
	if in.SecondPartyId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"second_party_id\":")
		out.Int32(int32(in.SecondPartyId))
	}
	if in.SecondPartyType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"second_party_type\":")
		out.String(string(in.SecondPartyType))
	}
	if in.Amount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"amount\":")
		out.Float32(float32(in.Amount))
	}
	if in.Balance != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"balance\":")
		out.Float32(float32(in.Balance))
	}
	if in.Reason != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"reason\":")
		out.String(string(in.Reason))
	}
	if in.TaxRecieverId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"tax_reciever_id\":")
		out.Int32(int32(in.TaxRecieverId))
	}
	if in.Tax != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"tax\":")
		out.Float32(float32(in.Tax))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"extra_info\":")
		easyjson46989833EncodeGithubComAntihaxGoesiEsi2(out, in.ExtraInfo)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdWalletJournal200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson46989833EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdWalletJournal200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson46989833EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletJournal200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson46989833DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletJournal200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson46989833DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson46989833DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdWalletJournalExtraInfo) {
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
func easyjson46989833EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdWalletJournalExtraInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LocationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"location_id\":")
		out.Int64(int64(in.LocationId))
	}
	if in.TransactionId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"transaction_id\":")
		out.Int64(int64(in.TransactionId))
	}
	if in.NpcName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"npc_name\":")
		out.String(string(in.NpcName))
	}
	if in.NpcId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"npc_id\":")
		out.Int32(int32(in.NpcId))
	}
	if in.DestroyedShipTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"destroyed_ship_type_id\":")
		out.Int32(int32(in.DestroyedShipTypeId))
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
	if in.AllianceId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance_id\":")
		out.Int32(int32(in.AllianceId))
	}
	if in.JobId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"job_id\":")
		out.Int32(int32(in.JobId))
	}
	if in.ContractId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"contract_id\":")
		out.Int32(int32(in.ContractId))
	}
	if in.SystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"system_id\":")
		out.Int32(int32(in.SystemId))
	}
	if in.PlanetId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"planet_id\":")
		out.Int32(int32(in.PlanetId))
	}
	out.RawByte('}')
}
