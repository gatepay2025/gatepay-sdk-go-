package common

// 定义协议公共头
type BaseRequest struct {
	Header  map[string]string
	Version string
}

func (r BaseRequest) GetVersion() string {
	return r.Version
}

func (r BaseRequest) GetHeaders() map[string]string {
	return r.Header
}

func (r *BaseRequest) AddHeader(key, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
}
