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

func easyjsonD831633DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdWalletTransactions200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdWalletTransactions200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdWalletTransactions200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdWalletTransactions200Ok
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
func easyjsonD831633EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdWalletTransactions200OkList) {
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
func (v GetCharactersCharacterIdWalletTransactions200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD831633EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdWalletTransactions200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD831633EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletTransactions200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD831633DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletTransactions200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD831633DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonD831633DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdWalletTransactions200Ok) {
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
		case "client_id":
			out.ClientId = int32(in.Int32())
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		case "is_buy":
			out.IsBuy = bool(in.Bool())
		case "is_personal":
			out.IsPersonal = bool(in.Bool())
		case "journal_ref_id":
			out.JournalRefId = int64(in.Int64())
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "quantity":
			out.Quantity = int32(in.Int32())
		case "transaction_id":
			out.TransactionId = int64(in.Int64())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "unit_price":
			out.UnitPrice = int32(in.Int32())
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
func easyjsonD831633EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdWalletTransactions200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ClientId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"client_id\":")
		out.Int32(int32(in.ClientId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"date\":")
		out.Raw((in.Date).MarshalJSON())
	}
	if in.IsBuy {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_buy\":")
		out.Bool(bool(in.IsBuy))
	}
	if in.IsPersonal {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_personal\":")
		out.Bool(bool(in.IsPersonal))
	}
	if in.JournalRefId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"journal_ref_id\":")
		out.Int64(int64(in.JournalRefId))
	}
	if in.LocationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"location_id\":")
		out.Int64(int64(in.LocationId))
	}
	if in.Quantity != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quantity\":")
		out.Int32(int32(in.Quantity))
	}
	if in.TransactionId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"transaction_id\":")
		out.Int64(int64(in.TransactionId))
	}
	if in.TypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type_id\":")
		out.Int32(int32(in.TypeId))
	}
	if in.UnitPrice != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"unit_price\":")
		out.Int32(int32(in.UnitPrice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdWalletTransactions200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD831633EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdWalletTransactions200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD831633EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletTransactions200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD831633DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletTransactions200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD831633DecodeGithubComAntihaxGoesiEsi1(l, v)
}
