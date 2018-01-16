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

func easyjsonEbe19d94DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdRolesHistory200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdRolesHistory200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdRolesHistory200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdRolesHistory200Ok
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
func easyjsonEbe19d94EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdRolesHistory200OkList) {
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
func (v GetCorporationsCorporationIdRolesHistory200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbe19d94EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdRolesHistory200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbe19d94EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdRolesHistory200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbe19d94DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdRolesHistory200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbe19d94DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonEbe19d94DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdRolesHistory200Ok) {
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
		case "changed_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ChangedAt).UnmarshalJSON(data))
			}
		case "issuer_id":
			out.IssuerId = int32(in.Int32())
		case "role_type":
			out.RoleType = string(in.String())
		case "old_roles":
			if in.IsNull() {
				in.Skip()
				out.OldRoles = nil
			} else {
				in.Delim('[')
				if out.OldRoles == nil {
					if !in.IsDelim(']') {
						out.OldRoles = make([]string, 0, 4)
					} else {
						out.OldRoles = []string{}
					}
				} else {
					out.OldRoles = (out.OldRoles)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.OldRoles = append(out.OldRoles, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "new_roles":
			if in.IsNull() {
				in.Skip()
				out.NewRoles = nil
			} else {
				in.Delim('[')
				if out.NewRoles == nil {
					if !in.IsDelim(']') {
						out.NewRoles = make([]string, 0, 4)
					} else {
						out.NewRoles = []string{}
					}
				} else {
					out.NewRoles = (out.NewRoles)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.NewRoles = append(out.NewRoles, v5)
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
func easyjsonEbe19d94EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdRolesHistory200Ok) {
	out.RawByte('{')
	first := true
	_ = first
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
	if true {
		const prefix string = ",\"changed_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ChangedAt).MarshalJSON())
	}
	if in.IssuerId != 0 {
		const prefix string = ",\"issuer_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.IssuerId))
	}
	if in.RoleType != "" {
		const prefix string = ",\"role_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RoleType))
	}
	if len(in.OldRoles) != 0 {
		const prefix string = ",\"old_roles\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.OldRoles {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.NewRoles) != 0 {
		const prefix string = ",\"new_roles\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.NewRoles {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdRolesHistory200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbe19d94EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdRolesHistory200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbe19d94EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdRolesHistory200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbe19d94DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdRolesHistory200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbe19d94DecodeGithubComAntihaxGoesiEsi1(l, v)
}