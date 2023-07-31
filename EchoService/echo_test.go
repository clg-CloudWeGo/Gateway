package main

import (
	"bytes"
	"encoding/json"
	"github.com/clg-CloudWeGo/Echo/kitex_gen/api"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
)

const echoURL = "http://127.0.0.1:8888/gateway/EchoService/Echo"

var httpCli = &http.Client{Timeout: 3 * time.Second}

func BenchmarkEchoService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "hello raccoon" + strconv.Itoa(i)
		echoReq := getEchoMessage(s)
		reqBody, err := json.Marshal(echoReq)
		if err != nil {
			panic("marshal error")
		}
		var resp *http.Response
		req, err := http.NewRequest(http.MethodPost, echoURL, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err = httpCli.Do(req)
		defer resp.Body.Close()
		if err != nil {
			return
		}
		var body []byte
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			return
		}
		var echoResp *api.Response
		if err := json.Unmarshal(body, &echoResp); err != nil {
			panic("unmarshal error")
		}
		Assert(echoReq.Message, echoResp.Message)
	}
}

func Assert(a string, b string) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func getEchoMessage(s string) *api.Request {
	return &api.Request{Message: s}
}
