package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type APIError struct {
	StatusCode   int         // 应答报文的 HTTP 状态码
	Header       http.Header // 应答报文的 Header 信息
	Body         string      // 应答报文的 Body 原文
	Code         interface{} `json:"code"`
	Message      string      `json:"label"`
	ErrorMessage interface{} `json:"errorMessage"`
}

func (e *APIError) Error() string {
	var buf bytes.Buffer
	_, _ = fmt.Fprintf(&buf, "error http response:[StatusCode: %d Code: \"%s\"", e.StatusCode, e.Code)
	if e.Message != "" {
		_, _ = fmt.Fprintf(&buf, "\nMessage: %s", e.Message)
	}
	if e.ErrorMessage != nil {
		var detailBuf bytes.Buffer
		enc := json.NewEncoder(&detailBuf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(e.ErrorMessage); err == nil {
			_, _ = fmt.Fprint(&buf, "\nDetail:")
			_, _ = fmt.Fprintf(&buf, "\n%s", strings.TrimSpace(detailBuf.String()))
		}
	}
	if len(e.Header) > 0 {
		_, _ = fmt.Fprint(&buf, "\nHeader:")
		for key, value := range e.Header {
			_, _ = fmt.Fprintf(&buf, "\n - %v=%v", key, value)
		}
	}
	_, _ = fmt.Fprintf(&buf, "]")
	return buf.String()
}

func IsAPIError(err error, code string) bool {
	if ne, ok := err.(*APIError); ok {
		return ne.Code == code
	}
	return false
}
