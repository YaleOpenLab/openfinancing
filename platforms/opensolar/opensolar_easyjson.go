// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package opensolar

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

func easyjson4a457b9dDecodeGithubComYaleOpenLabOpenxPlatformsOpensolar(in *jlexer.Lexer, out *statusResponse) {
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
		case "Code":
			out.Code = int(in.Int())
		case "Status":
			out.Status = string(in.String())
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
func easyjson4a457b9dEncodeGithubComYaleOpenLabOpenxPlatformsOpensolar(out *jwriter.Writer, in statusResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"Status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v statusResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4a457b9dEncodeGithubComYaleOpenLabOpenxPlatformsOpensolar(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v statusResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4a457b9dEncodeGithubComYaleOpenLabOpenxPlatformsOpensolar(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *statusResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4a457b9dDecodeGithubComYaleOpenLabOpenxPlatformsOpensolar(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *statusResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4a457b9dDecodeGithubComYaleOpenLabOpenxPlatformsOpensolar(l, v)
}
func easyjson4a457b9dDecodeGithubComYaleOpenLabOpenxPlatformsOpensolar1(in *jlexer.Lexer, out *SolarProjectArray) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(SolarProjectArray, 0, 1)
			} else {
				*out = SolarProjectArray{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Project
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
func easyjson4a457b9dEncodeGithubComYaleOpenLabOpenxPlatformsOpensolar1(out *jwriter.Writer, in SolarProjectArray) {
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
func (v SolarProjectArray) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4a457b9dEncodeGithubComYaleOpenLabOpenxPlatformsOpensolar1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SolarProjectArray) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4a457b9dEncodeGithubComYaleOpenLabOpenxPlatformsOpensolar1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SolarProjectArray) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4a457b9dDecodeGithubComYaleOpenLabOpenxPlatformsOpensolar1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SolarProjectArray) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4a457b9dDecodeGithubComYaleOpenLabOpenxPlatformsOpensolar1(l, v)
}
func easyjson4a457b9dDecodeGithubComYaleOpenLabOpenxPlatformsOpensolar2(in *jlexer.Lexer, out *Project) {
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
		case "TotalValue":
			out.TotalValue = float64(in.Float64())
		case "MoneyRaised":
			out.MoneyRaised = float64(in.Float64())
		case "Years":
			out.Years = int(in.Int())
		case "InterestRate":
			out.InterestRate = float64(in.Float64())
		case "BalLeft":
			out.BalLeft = float64(in.Float64())
		case "Votes":
			out.Votes = int(in.Int())
		case "InvestorAssetCode":
			out.InvestorAssetCode = string(in.String())
		case "DebtAssetCode":
			out.DebtAssetCode = string(in.String())
		case "PaybackAssetCode":
			out.PaybackAssetCode = string(in.String())
		case "SeedAssetCode":
			out.SeedAssetCode = string(in.String())
		case "DateInitiated":
			out.DateInitiated = string(in.String())
		case "DateFunded":
			out.DateFunded = string(in.String())
		case "DateLastPaid":
			out.DateLastPaid = int64(in.Int64())
		case "Location":
			out.Location = string(in.String())
		case "PanelSize":
			out.PanelSize = string(in.String())
		case "Inverter":
			out.Inverter = string(in.String())
		case "ChargeRegulator":
			out.ChargeRegulator = string(in.String())
		case "ControlPanel":
			out.ControlPanel = string(in.String())
		case "CommBox":
			out.CommBox = string(in.String())
		case "ACTransfer":
			out.ACTransfer = string(in.String())
		case "SolarCombiner":
			out.SolarCombiner = string(in.String())
		case "Batteries":
			out.Batteries = string(in.String())
		case "IoTHub":
			out.IoTHub = string(in.String())
		case "Metadata":
			out.Metadata = string(in.String())
		case "Originator":
			(out.Originator).UnmarshalEasyJSON(in)
		case "OriginatorFee":
			out.OriginatorFee = float64(in.Float64())
		case "Developer":
			(out.Developer).UnmarshalEasyJSON(in)
		case "Guarantor":
			(out.Guarantor).UnmarshalEasyJSON(in)
		case "Contractor":
			(out.Contractor).UnmarshalEasyJSON(in)
		case "ContractorFee":
			out.ContractorFee = float64(in.Float64())
		case "SecondaryContractor":
			(out.SecondaryContractor).UnmarshalEasyJSON(in)
		case "SecondaryContractorFee":
			out.SecondaryContractorFee = float64(in.Float64())
		case "TertiaryContractor":
			(out.TertiaryContractor).UnmarshalEasyJSON(in)
		case "TertiaryContractorFee":
			out.TertiaryContractorFee = float64(in.Float64())
		case "DeveloperFee":
			out.DeveloperFee = float64(in.Float64())
		case "RecipientIndex":
			out.RecipientIndex = int(in.Int())
		case "InvestorIndices":
			if in.IsNull() {
				in.Skip()
				out.InvestorIndices = nil
			} else {
				in.Delim('[')
				if out.InvestorIndices == nil {
					if !in.IsDelim(']') {
						out.InvestorIndices = make([]int, 0, 8)
					} else {
						out.InvestorIndices = []int{}
					}
				} else {
					out.InvestorIndices = (out.InvestorIndices)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int
					v4 = int(in.Int())
					out.InvestorIndices = append(out.InvestorIndices, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "SeedInvestorIndices":
			if in.IsNull() {
				in.Skip()
				out.SeedInvestorIndices = nil
			} else {
				in.Delim('[')
				if out.SeedInvestorIndices == nil {
					if !in.IsDelim(']') {
						out.SeedInvestorIndices = make([]int, 0, 8)
					} else {
						out.SeedInvestorIndices = []int{}
					}
				} else {
					out.SeedInvestorIndices = (out.SeedInvestorIndices)[:0]
				}
				for !in.IsDelim(']') {
					var v5 int
					v5 = int(in.Int())
					out.SeedInvestorIndices = append(out.SeedInvestorIndices, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Stage":
			out.Stage = float64(in.Float64())
		case "AuctionType":
			out.AuctionType = string(in.String())
		case "OriginatorMoUHash":
			out.OriginatorMoUHash = string(in.String())
		case "ContractorContractHash":
			out.ContractorContractHash = string(in.String())
		case "InvPlatformContractHash":
			out.InvPlatformContractHash = string(in.String())
		case "RecPlatformContractHash":
			out.RecPlatformContractHash = string(in.String())
		case "SpecSheetHash":
			out.SpecSheetHash = string(in.String())
		case "Reputation":
			out.Reputation = float64(in.Float64())
		case "Lock":
			out.Lock = bool(in.Bool())
		case "LockPwd":
			out.LockPwd = string(in.String())
		case "InvestmentType":
			out.InvestmentType = string(in.String())
		case "PaybackPeriod":
			out.PaybackPeriod = int(in.Int())
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
func easyjson4a457b9dEncodeGithubComYaleOpenLabOpenxPlatformsOpensolar2(out *jwriter.Writer, in Project) {
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
		const prefix string = ",\"TotalValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.TotalValue))
	}
	{
		const prefix string = ",\"MoneyRaised\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.MoneyRaised))
	}
	{
		const prefix string = ",\"Years\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Years))
	}
	{
		const prefix string = ",\"InterestRate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.InterestRate))
	}
	{
		const prefix string = ",\"BalLeft\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.BalLeft))
	}
	{
		const prefix string = ",\"Votes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Votes))
	}
	{
		const prefix string = ",\"InvestorAssetCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InvestorAssetCode))
	}
	{
		const prefix string = ",\"DebtAssetCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DebtAssetCode))
	}
	{
		const prefix string = ",\"PaybackAssetCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PaybackAssetCode))
	}
	{
		const prefix string = ",\"SeedAssetCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SeedAssetCode))
	}
	{
		const prefix string = ",\"DateInitiated\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DateInitiated))
	}
	{
		const prefix string = ",\"DateFunded\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DateFunded))
	}
	{
		const prefix string = ",\"DateLastPaid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DateLastPaid))
	}
	{
		const prefix string = ",\"Location\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Location))
	}
	{
		const prefix string = ",\"PanelSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PanelSize))
	}
	{
		const prefix string = ",\"Inverter\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Inverter))
	}
	{
		const prefix string = ",\"ChargeRegulator\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ChargeRegulator))
	}
	{
		const prefix string = ",\"ControlPanel\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ControlPanel))
	}
	{
		const prefix string = ",\"CommBox\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CommBox))
	}
	{
		const prefix string = ",\"ACTransfer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ACTransfer))
	}
	{
		const prefix string = ",\"SolarCombiner\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SolarCombiner))
	}
	{
		const prefix string = ",\"Batteries\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Batteries))
	}
	{
		const prefix string = ",\"IoTHub\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IoTHub))
	}
	{
		const prefix string = ",\"Metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Metadata))
	}
	{
		const prefix string = ",\"Originator\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Originator).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"OriginatorFee\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.OriginatorFee))
	}
	{
		const prefix string = ",\"Developer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Developer).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Guarantor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Guarantor).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Contractor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Contractor).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"ContractorFee\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.ContractorFee))
	}
	{
		const prefix string = ",\"SecondaryContractor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.SecondaryContractor).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"SecondaryContractorFee\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.SecondaryContractorFee))
	}
	{
		const prefix string = ",\"TertiaryContractor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.TertiaryContractor).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"TertiaryContractorFee\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.TertiaryContractorFee))
	}
	{
		const prefix string = ",\"DeveloperFee\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.DeveloperFee))
	}
	{
		const prefix string = ",\"RecipientIndex\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.RecipientIndex))
	}
	{
		const prefix string = ",\"InvestorIndices\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.InvestorIndices == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.InvestorIndices {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v7))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"SeedInvestorIndices\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.SeedInvestorIndices == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.SeedInvestorIndices {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Stage\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Stage))
	}
	{
		const prefix string = ",\"AuctionType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AuctionType))
	}
	{
		const prefix string = ",\"OriginatorMoUHash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OriginatorMoUHash))
	}
	{
		const prefix string = ",\"ContractorContractHash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContractorContractHash))
	}
	{
		const prefix string = ",\"InvPlatformContractHash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InvPlatformContractHash))
	}
	{
		const prefix string = ",\"RecPlatformContractHash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecPlatformContractHash))
	}
	{
		const prefix string = ",\"SpecSheetHash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SpecSheetHash))
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
		const prefix string = ",\"Lock\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Lock))
	}
	{
		const prefix string = ",\"LockPwd\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LockPwd))
	}
	{
		const prefix string = ",\"InvestmentType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InvestmentType))
	}
	{
		const prefix string = ",\"PaybackPeriod\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.PaybackPeriod))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Project) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4a457b9dEncodeGithubComYaleOpenLabOpenxPlatformsOpensolar2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Project) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4a457b9dEncodeGithubComYaleOpenLabOpenxPlatformsOpensolar2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Project) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4a457b9dDecodeGithubComYaleOpenLabOpenxPlatformsOpensolar2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Project) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4a457b9dDecodeGithubComYaleOpenLabOpenxPlatformsOpensolar2(l, v)
}
