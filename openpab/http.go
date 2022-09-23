package openpab

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type HttpRequest struct {
	URL     string
	StartAt time.Time
	Raw     *http.Request
}

type HttpResponse struct {
	request *HttpRequest
	raw     *http.Response

	body       []byte
	size       int64
	receivedAt time.Time
	traceId    string
}

// Request return HTTP request
func (r *HttpResponse) Request() *HttpRequest {
	return r.request
}

// Body method returns HTTP response as []byte array for the executed request.
func (r *HttpResponse) Body() []byte {
	if r.raw == nil {
		return []byte{}
	}
	return r.body
}

// Status method returns the HTTP status string for the executed request.
//
//	Example: 200 OK
func (r *HttpResponse) Status() string {
	if r.raw == nil {
		return ""
	}
	return r.raw.Status
}

// StatusCode method returns the HTTP status code for the executed request.
//
//	Example: 200
func (r *HttpResponse) StatusCode() int {
	if r.raw == nil {
		return 0
	}
	return r.raw.StatusCode
}

// Proto method returns the HTTP response protocol used for the request.
func (r *HttpResponse) Proto() string {
	if r.raw == nil {
		return ""
	}
	return r.raw.Proto
}

// Header method returns the response headers
func (r *HttpResponse) Header() http.Header {
	if r.raw == nil {
		return http.Header{}
	}
	return r.raw.Header
}

// Cookies method to access all the response cookies
func (r *HttpResponse) Cookies() []*http.Cookie {
	if r.raw == nil {
		return make([]*http.Cookie, 0)
	}
	return r.raw.Cookies()
}

// String method returns the body of the server response as String.
func (r *HttpResponse) String() string {
	if r.body == nil {
		return ""
	}
	return strings.TrimSpace(string(r.body))
}

// Latency method returns the time of HTTP response time that from request we sent and received a request.
//
// See `Response.receivedAt` to know when client received response and see `Response.Request.StartAt` to know
// when client sent a request.
func (r *HttpResponse) Latency() time.Duration {
	return r.receivedAt.Sub(r.request.StartAt)
}

// ReceivedAt method returns when response got received from server for the request.
func (r *HttpResponse) ReceivedAt() time.Time {
	return r.receivedAt
}

// Size method returns the HTTP response size in bytes. Ya, you can relay on HTTP `Content-Length` header,
// however it won't be good for chucked transfer/compressed response. Since Resty calculates response size
// at the client end. You will get actual size of the http response.
func (r *HttpResponse) Size() int64 {
	return r.size
}

// RawBody method exposes the HTTP raw response body. Use this method in-conjunction with `SetDoNotParseResponse`
// option otherwise you get an error as `read err: http: read on closed response body`.
//
// Do not forget to close the body, otherwise you might get into connection leaks, no connection reuse.
// Basically you have taken over the control of response parsing from `Resty`.
func (r *HttpResponse) RawBody() io.ReadCloser {
	if r.raw == nil {
		return nil
	}
	return r.raw.Body
}

// IsSuccess method returns true if HTTP status `code >= 200 and <= 299` otherwise false.
func (r *HttpResponse) IsSuccess() bool {
	return r.StatusCode() > 199 && r.StatusCode() < 300
}

// IsError method returns true if HTTP status `code >= 400` otherwise false.
func (r *HttpResponse) IsError() bool {
	return r.StatusCode() > 399
}

// TraceId method returns this request response trace id.
func (r *HttpResponse) TraceId() string {
	return r.traceId
}

func (c *Client) post(ctx context.Context, url string, headers map[string]string, body string) (*HttpResponse, error) {
	return c.execute(ctx, http.MethodPost, url, headers, body)
}

func (c *Client) execute(ctx context.Context, method, url string, headers map[string]string, body string) (*HttpResponse, error) {
	traceId := c.getTraceId(ctx)

	c.log.Debugf(`{"req.proto":{"method":"%s","url":"%s"},"traceId":"%s"}`, method, url, traceId)

	req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "aop-sdk-go")
	for k, v := range headers {
		req.Header[k] = []string{v}
	}
	if method != http.MethodDelete {
		req.Header.Add("Accept", "text/plain")
	}
	if method != http.MethodDelete && method != http.MethodGet {
		req.Header.Add("Content-Type", "application/json")
	}

	markReq := &HttpRequest{
		URL:     url,
		StartAt: time.Now(),
		Raw:     req,
	}

	c.log.Infof(`{"req":{"method":"%s","url":"%s","header":"%v","body":%s},"traceId":"%s"}`, method, url, req.Header, body, traceId)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	defer resp.Body.Close()
	receivedAt := time.Now()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	markResp := &HttpResponse{
		request:    markReq,
		raw:        resp,
		body:       b,
		size:       int64(len(b)),
		receivedAt: receivedAt,
		traceId:    traceId,
	}
	c.log.Infof(`{"rsp":{"code":%d,"header":"%v","body":%s},"traceId":"%s","latency":%v}`,
		markResp.StatusCode(), markResp.Header(), markResp.String(), traceId, markResp.Latency())
	return markResp, nil
}

func (c *Client) getFile(ctx context.Context, url string, headers map[string]string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	rsp, err := c.httpClient.Do(req)
	if err != nil {
		if rsp != nil {
			rsp.Body.Close()
		}
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		_, _ = io.ReadAll(rsp.Body)
		rsp.Body.Close()
		return nil, errors.New("文件下载异常")
	}
	return rsp.Body, nil
}
