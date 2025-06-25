package core

import "regexp"

// var DefaultEndpoint = "https://openplatform.gateapi.io"
var DefaultEndpoint = "http://dev.halftrust.xyz/gfpay"

const (
	SchemeHTTP  = "http"
	SchemeHTTPS = "https"

	MethodGet    = "GET"
	MethodPut    = "PUT"
	MethodPost   = "POST"
	MethodDelete = "DELETE"
	MethodPatch  = "PATCH"
	MethodHead   = "HEAD"
)

const (
	ApplicationJSON = "application/json"
	Accept          = "Accept"       // Header 中的 Accept 字段
	ContentType     = "Content-Type" // Header 中的 ContentType 字段
	UserAgent       = "User-Agent"
)

var (
	regJSONTypeCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
)

// SDK 相关信息
const (
	Version         = "0.0.1"                        // SDK 版本
	UserAgentFormat = "GatePay-SDK-Go/%s (%s) GO/%s" // UserAgent中的信息
)
