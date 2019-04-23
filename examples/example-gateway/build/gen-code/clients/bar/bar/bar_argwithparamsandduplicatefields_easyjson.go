// Code generated by zanzibar
// @generated
// Checksum : w+vBwewhkvFL9NlSicST5Q==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
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

func easyjsonAc132a5fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields(in *jlexer.Lexer, out *Bar_ArgWithParamsAndDuplicateFields_Result) {
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
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(BarResponse)
				}
				(*out.Success).UnmarshalEasyJSON(in)
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
func easyjsonAc132a5fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields(out *jwriter.Writer, in Bar_ArgWithParamsAndDuplicateFields_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Success).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithParamsAndDuplicateFields_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAc132a5fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithParamsAndDuplicateFields_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAc132a5fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithParamsAndDuplicateFields_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAc132a5fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithParamsAndDuplicateFields_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAc132a5fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields(l, v)
}
func easyjsonAc132a5fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields1(in *jlexer.Lexer, out *Bar_ArgWithParamsAndDuplicateFields_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var RequestSet bool
	var EntityUUIDSet bool
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
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(RequestWithDuplicateType)
				}
				easyjsonAc132a5fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in, out.Request)
			}
			RequestSet = true
		case "entityUUID":
			out.EntityUUID = string(in.String())
			EntityUUIDSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !RequestSet {
		in.AddError(fmt.Errorf("key 'request' is required"))
	}
	if !EntityUUIDSet {
		in.AddError(fmt.Errorf("key 'entityUUID' is required"))
	}
}
func easyjsonAc132a5fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields1(out *jwriter.Writer, in Bar_ArgWithParamsAndDuplicateFields_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"request\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Request == nil {
			out.RawString("null")
		} else {
			easyjsonAc132a5fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out, *in.Request)
		}
	}
	{
		const prefix string = ",\"entityUUID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EntityUUID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithParamsAndDuplicateFields_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAc132a5fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithParamsAndDuplicateFields_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAc132a5fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithParamsAndDuplicateFields_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAc132a5fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithParamsAndDuplicateFields_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAc132a5fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithParamsAndDuplicateFields1(l, v)
}
func easyjsonAc132a5fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in *jlexer.Lexer, out *RequestWithDuplicateType) {
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
		case "request1":
			if in.IsNull() {
				in.Skip()
				out.Request1 = nil
			} else {
				if out.Request1 == nil {
					out.Request1 = new(BarRequest)
				}
				(*out.Request1).UnmarshalEasyJSON(in)
			}
		case "request2":
			if in.IsNull() {
				in.Skip()
				out.Request2 = nil
			} else {
				if out.Request2 == nil {
					out.Request2 = new(BarRequest)
				}
				(*out.Request2).UnmarshalEasyJSON(in)
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
func easyjsonAc132a5fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out *jwriter.Writer, in RequestWithDuplicateType) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Request1 != nil {
		const prefix string = ",\"request1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Request1).MarshalEasyJSON(out)
	}
	if in.Request2 != nil {
		const prefix string = ",\"request2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Request2).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
