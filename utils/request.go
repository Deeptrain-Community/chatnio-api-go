package utils

import (
	"github.com/valyala/fasthttp"
)

type Headers map[string]string

func Get(uri string, headers Headers) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(uri)
	req.Header.SetMethod("GET")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func GetForm[T any](uri string, headers Headers) (*T, error) {
	res, err := Get(uri, headers)
	if err != nil {
		return nil, err
	}

	return UnmarshalByteForm[T](res)
}

func Post(uri string, headers Headers, body string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(uri)
	req.Header.SetMethod("POST")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.SetBodyString(body)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func PostForm[T any](uri string, headers Headers, body interface{}) (*T, error) {
	data := MarshalForm(body)
	res, err := Post(uri, headers, data)
	if err != nil {
		return nil, err
	}

	return UnmarshalByteForm[T](res)
}
