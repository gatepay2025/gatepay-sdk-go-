package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gate/gatepay-sdk-go/core/signature"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"time"
)

type httpResponse struct {
	Status string `json:"status"`
	//Code         string `json:"code"`
	Code interface{} `json:"code"`

	Label        string `json:"label,omitempty"`
	ErrorMessage string `json:"errorMessage"`
	//Data         interface{} `json:"data"`
	Data json.RawMessage `json:"data"`
}

type APIResult struct {
	Request  *http.Request
	Response *http.Response
}

type Client struct {
	httpClient *http.Client
	Credential *Credential
	Config     *Config
}

func NewClient(config *Config, credential *Credential) (client *Client, err error) {
	client = &Client{}
	if config == nil {
		config = NewConfig()
	}

	if credential == nil {
		return nil, fmt.Errorf("credential is required")
	}

	client.WithConfig(config).WithCredential(credential)
	if client.httpClient == nil {
		client.httpClient = &http.Client{
			Timeout: config.Timeout,
		}
	}

	return
}

func (cli *Client) WithCredential(cred *Credential) *Client {
	cli.Credential = cred
	return cli
}

func (c *Client) WithConfig(config *Config) *Client {
	c.Config = config
	return c
}

func (cli *Client) Request(ctx context.Context,
	method, requestPath string,
	headerParams http.Header,
	queryParams url.Values,
	postBody interface{},
	contentType string) (result *APIResult, err error) {

	varURL, err := url.Parse(requestPath)
	if err != nil {
		return nil, err
	}

	//获取path里的查询参数
	query := varURL.Query()
	//获取传入查询参数
	for k, values := range queryParams {
		for _, v := range values {
			query.Add(k, v)
		}
	}

	if len(query) > 0 {
		varURL.RawQuery = query.Encode()
	}

	if postBody == nil {
		//不带BODY
		return cli.doRequest(ctx, method, varURL.String(), headerParams, contentType, nil, "")
	}

	if contentType == "" {
		contentType = ApplicationJSON
	}

	var body *bytes.Buffer
	body, err = setBody(postBody, contentType)
	if err != nil {
		return nil, err
	}

	return cli.doRequest(ctx, method, varURL.String(), headerParams, contentType, body, body.String())
}

func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	switch b := body.(type) {
	case string:
		_, err = bodyBuf.WriteString(b)
	case *string:
		_, err = bodyBuf.WriteString(*b)
	case []byte:
		_, err = bodyBuf.Write(b)
	case **os.File:
		_, err = bodyBuf.ReadFrom(*b)
	case io.Reader:
		_, err = bodyBuf.ReadFrom(b)
	default:
		if regJSONTypeCheck.MatchString(contentType) {
			err = json.NewEncoder(bodyBuf).Encode(body)
		}
	}
	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

func (cli *Client) doRequest(
	ctx context.Context,
	method string,
	requestURL string,
	header http.Header,
	contentType string,
	reqBody io.Reader,
	signBody string,
) (*APIResult, error) {

	var (
		err     error
		request *http.Request
	)
	if request, err = http.NewRequestWithContext(ctx, method, requestURL, reqBody); err != nil {
		return nil, err
	}

	//for key, values := range header {
	//	for _, v := range values {
	//		request.Header.Add(key, v)
	//	}
	//}
	request.Header = header.Clone()
	request.Header.Set(Accept, "*/*")
	request.Header.Set(ContentType, contentType)
	ua := fmt.Sprintf(UserAgentFormat, Version, runtime.GOOS, runtime.Version())
	request.Header.Set(UserAgent, ua)
	//签名
	t := time.Now().UnixMilli()
	ts := strconv.FormatInt(t, 10)
	nonce := signature.GenerateNonce(9)
	sing := signature.VerifySignature(ts, nonce, signBody, cli.Credential.SecretKey)
	request.Header.Set(signature.HeaderGatePayTimestamp, ts)
	request.Header.Set(signature.HeaderGatePayNonce, nonce)
	request.Header.Set(signature.HeaderGatePaySignature, sing)

	result, err := cli.doHTTP(request)
	if err != nil {
		return result, err
	}

	if err = CheckResponse(result.Response); err != nil {
		return result, err
	}
	return result, nil
}

func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	slurp, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid response, read body error: %w", err)
	}
	_ = resp.Body.Close()

	resp.Body = io.NopCloser(bytes.NewBuffer(slurp))
	apiError := &APIError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(slurp),
	}
	// 忽略 JSON 解析错误，均返回 apiError
	_ = json.Unmarshal(slurp, apiError)
	return apiError
}

func (cli *Client) doHTTP(req *http.Request) (result *APIResult, err error) {
	result = &APIResult{
		Request: req,
	}

	result.Response, err = cli.httpClient.Do(req)
	return result, err
}

func UnMarshalResponse(httpResp *http.Response, resp interface{}) error {
	body, err := io.ReadAll(httpResp.Body)
	_ = httpResp.Body.Close()

	if err != nil {
		return err
	}

	httpResp.Body = io.NopCloser(bytes.NewBuffer(body))
	httpResponse := &httpResponse{}
	err = json.Unmarshal(body, httpResponse)
	if err != nil {
		return err
	}

	//判断状态码
	if httpResponse.Status != "SUCCESS" {
		slurp, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return fmt.Errorf("invalid response, read body error: %w", err)
		}
		_ = httpResp.Body.Close()
		httpResp.Body = io.NopCloser(bytes.NewBuffer(slurp))
		apiError := &APIError{
			StatusCode: httpResp.StatusCode,
			Header:     httpResp.Header,
			Body:       string(slurp),
		}

		_ = json.Unmarshal(slurp, apiError)
		return apiError
	}
	err = json.Unmarshal(httpResponse.Data, resp)
	if err != nil {
		return err
	}
	return nil
}
