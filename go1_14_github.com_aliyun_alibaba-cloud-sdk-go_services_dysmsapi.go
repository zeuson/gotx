// Code generated by 'goexports github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi'. DO NOT EDIT.

// +build go1.14,!go1.15

package main

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"reflect"
)

func init() {
	GotxSymbols["github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CreateAddSmsSignRequest":          reflect.ValueOf(dysmsapi.CreateAddSmsSignRequest),
		"CreateAddSmsSignResponse":         reflect.ValueOf(dysmsapi.CreateAddSmsSignResponse),
		"CreateAddSmsTemplateRequest":      reflect.ValueOf(dysmsapi.CreateAddSmsTemplateRequest),
		"CreateAddSmsTemplateResponse":     reflect.ValueOf(dysmsapi.CreateAddSmsTemplateResponse),
		"CreateDeleteSmsSignRequest":       reflect.ValueOf(dysmsapi.CreateDeleteSmsSignRequest),
		"CreateDeleteSmsSignResponse":      reflect.ValueOf(dysmsapi.CreateDeleteSmsSignResponse),
		"CreateDeleteSmsTemplateRequest":   reflect.ValueOf(dysmsapi.CreateDeleteSmsTemplateRequest),
		"CreateDeleteSmsTemplateResponse":  reflect.ValueOf(dysmsapi.CreateDeleteSmsTemplateResponse),
		"CreateModifySmsSignRequest":       reflect.ValueOf(dysmsapi.CreateModifySmsSignRequest),
		"CreateModifySmsSignResponse":      reflect.ValueOf(dysmsapi.CreateModifySmsSignResponse),
		"CreateModifySmsTemplateRequest":   reflect.ValueOf(dysmsapi.CreateModifySmsTemplateRequest),
		"CreateModifySmsTemplateResponse":  reflect.ValueOf(dysmsapi.CreateModifySmsTemplateResponse),
		"CreateQuerySendDetailsRequest":    reflect.ValueOf(dysmsapi.CreateQuerySendDetailsRequest),
		"CreateQuerySendDetailsResponse":   reflect.ValueOf(dysmsapi.CreateQuerySendDetailsResponse),
		"CreateQuerySmsSignRequest":        reflect.ValueOf(dysmsapi.CreateQuerySmsSignRequest),
		"CreateQuerySmsSignResponse":       reflect.ValueOf(dysmsapi.CreateQuerySmsSignResponse),
		"CreateQuerySmsTemplateRequest":    reflect.ValueOf(dysmsapi.CreateQuerySmsTemplateRequest),
		"CreateQuerySmsTemplateResponse":   reflect.ValueOf(dysmsapi.CreateQuerySmsTemplateResponse),
		"CreateSendBatchSmsRequest":        reflect.ValueOf(dysmsapi.CreateSendBatchSmsRequest),
		"CreateSendBatchSmsResponse":       reflect.ValueOf(dysmsapi.CreateSendBatchSmsResponse),
		"CreateSendSmsRequest":             reflect.ValueOf(dysmsapi.CreateSendSmsRequest),
		"CreateSendSmsResponse":            reflect.ValueOf(dysmsapi.CreateSendSmsResponse),
		"EndpointMap":                      reflect.ValueOf(&dysmsapi.EndpointMap).Elem(),
		"EndpointType":                     reflect.ValueOf(&dysmsapi.EndpointType).Elem(),
		"GetEndpointMap":                   reflect.ValueOf(dysmsapi.GetEndpointMap),
		"GetEndpointType":                  reflect.ValueOf(dysmsapi.GetEndpointType),
		"NewClient":                        reflect.ValueOf(dysmsapi.NewClient),
		"NewClientWithAccessKey":           reflect.ValueOf(dysmsapi.NewClientWithAccessKey),
		"NewClientWithEcsRamRole":          reflect.ValueOf(dysmsapi.NewClientWithEcsRamRole),
		"NewClientWithOptions":             reflect.ValueOf(dysmsapi.NewClientWithOptions),
		"NewClientWithProvider":            reflect.ValueOf(dysmsapi.NewClientWithProvider),
		"NewClientWithRamRoleArn":          reflect.ValueOf(dysmsapi.NewClientWithRamRoleArn),
		"NewClientWithRamRoleArnAndPolicy": reflect.ValueOf(dysmsapi.NewClientWithRamRoleArnAndPolicy),
		"NewClientWithRsaKeyPair":          reflect.ValueOf(dysmsapi.NewClientWithRsaKeyPair),
		"NewClientWithStsToken":            reflect.ValueOf(dysmsapi.NewClientWithStsToken),
		"SetClientProperty":                reflect.ValueOf(dysmsapi.SetClientProperty),
		"SetEndpointDataToClient":          reflect.ValueOf(dysmsapi.SetEndpointDataToClient),

		// type definitions
		"AddSmsSignRequest":         reflect.ValueOf((*dysmsapi.AddSmsSignRequest)(nil)),
		"AddSmsSignResponse":        reflect.ValueOf((*dysmsapi.AddSmsSignResponse)(nil)),
		"AddSmsSignSignFileList":    reflect.ValueOf((*dysmsapi.AddSmsSignSignFileList)(nil)),
		"AddSmsTemplateRequest":     reflect.ValueOf((*dysmsapi.AddSmsTemplateRequest)(nil)),
		"AddSmsTemplateResponse":    reflect.ValueOf((*dysmsapi.AddSmsTemplateResponse)(nil)),
		"Client":                    reflect.ValueOf((*dysmsapi.Client)(nil)),
		"DeleteSmsSignRequest":      reflect.ValueOf((*dysmsapi.DeleteSmsSignRequest)(nil)),
		"DeleteSmsSignResponse":     reflect.ValueOf((*dysmsapi.DeleteSmsSignResponse)(nil)),
		"DeleteSmsTemplateRequest":  reflect.ValueOf((*dysmsapi.DeleteSmsTemplateRequest)(nil)),
		"DeleteSmsTemplateResponse": reflect.ValueOf((*dysmsapi.DeleteSmsTemplateResponse)(nil)),
		"ModifySmsSignRequest":      reflect.ValueOf((*dysmsapi.ModifySmsSignRequest)(nil)),
		"ModifySmsSignResponse":     reflect.ValueOf((*dysmsapi.ModifySmsSignResponse)(nil)),
		"ModifySmsSignSignFileList": reflect.ValueOf((*dysmsapi.ModifySmsSignSignFileList)(nil)),
		"ModifySmsTemplateRequest":  reflect.ValueOf((*dysmsapi.ModifySmsTemplateRequest)(nil)),
		"ModifySmsTemplateResponse": reflect.ValueOf((*dysmsapi.ModifySmsTemplateResponse)(nil)),
		"QuerySendDetailsRequest":   reflect.ValueOf((*dysmsapi.QuerySendDetailsRequest)(nil)),
		"QuerySendDetailsResponse":  reflect.ValueOf((*dysmsapi.QuerySendDetailsResponse)(nil)),
		"QuerySmsSignRequest":       reflect.ValueOf((*dysmsapi.QuerySmsSignRequest)(nil)),
		"QuerySmsSignResponse":      reflect.ValueOf((*dysmsapi.QuerySmsSignResponse)(nil)),
		"QuerySmsTemplateRequest":   reflect.ValueOf((*dysmsapi.QuerySmsTemplateRequest)(nil)),
		"QuerySmsTemplateResponse":  reflect.ValueOf((*dysmsapi.QuerySmsTemplateResponse)(nil)),
		"SendBatchSmsRequest":       reflect.ValueOf((*dysmsapi.SendBatchSmsRequest)(nil)),
		"SendBatchSmsResponse":      reflect.ValueOf((*dysmsapi.SendBatchSmsResponse)(nil)),
		"SendSmsRequest":            reflect.ValueOf((*dysmsapi.SendSmsRequest)(nil)),
		"SendSmsResponse":           reflect.ValueOf((*dysmsapi.SendSmsResponse)(nil)),
		"SmsSendDetailDTO":          reflect.ValueOf((*dysmsapi.SmsSendDetailDTO)(nil)),
		"SmsSendDetailDTOs":         reflect.ValueOf((*dysmsapi.SmsSendDetailDTOs)(nil)),
	}
}