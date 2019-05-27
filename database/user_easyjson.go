// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package database

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

func easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase(in *jlexer.Lexer, out *Wallet) {
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
		case "EncryptedSeed":
			if in.IsNull() {
				in.Skip()
				out.EncryptedSeed = nil
			} else {
				out.EncryptedSeed = in.Bytes()
			}
		case "PublicKey":
			out.PublicKey = string(in.String())
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
func easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase(out *jwriter.Writer, in Wallet) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"EncryptedSeed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.EncryptedSeed)
	}
	{
		const prefix string = ",\"PublicKey\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PublicKey))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Wallet) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Wallet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Wallet) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Wallet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase(l, v)
}
func easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase1(in *jlexer.Lexer, out *User) {
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
		case "Index":
			out.Index = int(in.Int())
		case "EncryptedSeed":
			if in.IsNull() {
				in.Skip()
				out.EncryptedSeed = nil
			} else {
				out.EncryptedSeed = in.Bytes()
			}
		case "Name":
			out.Name = string(in.String())
		case "PublicKey":
			out.PublicKey = string(in.String())
		case "City":
			out.City = string(in.String())
		case "ZipCode":
			out.ZipCode = string(in.String())
		case "Country":
			out.Country = string(in.String())
		case "RecoveryPhone":
			out.RecoveryPhone = string(in.String())
		case "Username":
			out.Username = string(in.String())
		case "Pwhash":
			out.Pwhash = string(in.String())
		case "Address":
			out.Address = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "Image":
			out.Image = string(in.String())
		case "FirstSignedUp":
			out.FirstSignedUp = string(in.String())
		case "Kyc":
			out.Kyc = bool(in.Bool())
		case "Inspector":
			out.Inspector = bool(in.Bool())
		case "Banned":
			out.Banned = bool(in.Bool())
		case "Email":
			out.Email = string(in.String())
		case "Notification":
			out.Notification = bool(in.Bool())
		case "Reputation":
			out.Reputation = float64(in.Float64())
		case "LocalAssets":
			if in.IsNull() {
				in.Skip()
				out.LocalAssets = nil
			} else {
				in.Delim('[')
				if out.LocalAssets == nil {
					if !in.IsDelim(']') {
						out.LocalAssets = make([]string, 0, 4)
					} else {
						out.LocalAssets = []string{}
					}
				} else {
					out.LocalAssets = (out.LocalAssets)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.LocalAssets = append(out.LocalAssets, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "RecoveryShares":
			if in.IsNull() {
				in.Skip()
				out.RecoveryShares = nil
			} else {
				in.Delim('[')
				if out.RecoveryShares == nil {
					if !in.IsDelim(']') {
						out.RecoveryShares = make([]string, 0, 4)
					} else {
						out.RecoveryShares = []string{}
					}
				} else {
					out.RecoveryShares = (out.RecoveryShares)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.RecoveryShares = append(out.RecoveryShares, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "PwdResetCode":
			out.PwdResetCode = string(in.String())
		case "SecondaryWallet":
			(out.SecondaryWallet).UnmarshalEasyJSON(in)
		case "EthereumWallet":
			(out.EthereumWallet).UnmarshalEasyJSON(in)
		case "PendingDocuments":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.PendingDocuments = make(map[string]string)
				} else {
					out.PendingDocuments = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 string
					v7 = string(in.String())
					(out.PendingDocuments)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		case "KYC":
			(out.KYC).UnmarshalEasyJSON(in)
		case "StarRating":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.StarRating = make(map[int]int)
				} else {
					out.StarRating = nil
				}
				for !in.IsDelim('}') {
					key := int(in.IntStr())
					in.WantColon()
					var v8 int
					v8 = int(in.Int())
					(out.StarRating)[key] = v8
					in.WantComma()
				}
				in.Delim('}')
			}
		case "GivenStarRating":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.GivenStarRating = make(map[int]int)
				} else {
					out.GivenStarRating = nil
				}
				for !in.IsDelim('}') {
					key := int(in.IntStr())
					in.WantColon()
					var v9 int
					v9 = int(in.Int())
					(out.GivenStarRating)[key] = v9
					in.WantComma()
				}
				in.Delim('}')
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
func easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Index\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Index))
	}
	{
		const prefix string = ",\"EncryptedSeed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.EncryptedSeed)
	}
	{
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"PublicKey\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PublicKey))
	}
	{
		const prefix string = ",\"City\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"ZipCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ZipCode))
	}
	{
		const prefix string = ",\"Country\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"RecoveryPhone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecoveryPhone))
	}
	{
		const prefix string = ",\"Username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"Pwhash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Pwhash))
	}
	{
		const prefix string = ",\"Address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"Description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"Image\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"FirstSignedUp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FirstSignedUp))
	}
	{
		const prefix string = ",\"Kyc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Kyc))
	}
	{
		const prefix string = ",\"Inspector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Inspector))
	}
	{
		const prefix string = ",\"Banned\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Banned))
	}
	{
		const prefix string = ",\"Email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"Notification\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Notification))
	}
	{
		const prefix string = ",\"Reputation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Reputation))
	}
	{
		const prefix string = ",\"LocalAssets\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.LocalAssets == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.LocalAssets {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"RecoveryShares\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.RecoveryShares == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.RecoveryShares {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"PwdResetCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PwdResetCode))
	}
	{
		const prefix string = ",\"SecondaryWallet\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.SecondaryWallet).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"EthereumWallet\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.EthereumWallet).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"PendingDocuments\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.PendingDocuments == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v16First := true
			for v16Name, v16Value := range in.PendingDocuments {
				if v16First {
					v16First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v16Name))
				out.RawByte(':')
				out.String(string(v16Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"KYC\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.KYC).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"StarRating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StarRating == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v17First := true
			for v17Name, v17Value := range in.StarRating {
				if v17First {
					v17First = false
				} else {
					out.RawByte(',')
				}
				out.IntStr(int(v17Name))
				out.RawByte(':')
				out.Int(int(v17Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"GivenStarRating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.GivenStarRating == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v18First := true
			for v18Name, v18Value := range in.GivenStarRating {
				if v18First {
					v18First = false
				} else {
					out.RawByte(',')
				}
				out.IntStr(int(v18Name))
				out.RawByte(':')
				out.Int(int(v18Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase1(l, v)
}
func easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase2(in *jlexer.Lexer, out *KycStruct) {
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
		case "PassportPhoto":
			out.PassportPhoto = string(in.String())
		case "IDCardPhoto":
			out.IDCardPhoto = string(in.String())
		case "DriversLicense":
			out.DriversLicense = string(in.String())
		case "PersonalPhoto":
			out.PersonalPhoto = string(in.String())
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
func easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase2(out *jwriter.Writer, in KycStruct) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"PassportPhoto\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PassportPhoto))
	}
	{
		const prefix string = ",\"IDCardPhoto\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IDCardPhoto))
	}
	{
		const prefix string = ",\"DriversLicense\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DriversLicense))
	}
	{
		const prefix string = ",\"PersonalPhoto\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PersonalPhoto))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v KycStruct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v KycStruct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *KycStruct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *KycStruct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase2(l, v)
}
func easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase3(in *jlexer.Lexer, out *EthWallet) {
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
		case "PrivateKey":
			out.PrivateKey = string(in.String())
		case "PublicKey":
			out.PublicKey = string(in.String())
		case "Address":
			out.Address = string(in.String())
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
func easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase3(out *jwriter.Writer, in EthWallet) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"PrivateKey\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PrivateKey))
	}
	{
		const prefix string = ",\"PublicKey\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PublicKey))
	}
	{
		const prefix string = ",\"Address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Address))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EthWallet) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EthWallet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComYaleOpenLabOpenxDatabase3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EthWallet) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EthWallet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComYaleOpenLabOpenxDatabase3(l, v)
}
