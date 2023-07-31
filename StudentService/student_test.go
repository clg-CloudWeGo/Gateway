package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	demo "github.com/Raccoon-njuse/rpcsvr/kitex_gen/demo"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const (
	queryURLFmt = "http://127.0.0.1:8888/query?id="
	registerURL = "http://127.0.0.1:8888/gateway/StudentService/Register"
)

var httpCli = &http.Client{Timeout: 3 * time.Second}

func BenchmarkStudentService(b *testing.B) {
	for i := 1; i < b.N; i++ {
		fmt.Println(i)
		newStu := genStudent(i)
		resp, err := register(newStu)
		Assert(b, err == nil, err)
		Assert(b, resp.Success, resp.Message)
	}
}

func register(stu *demo.Student) (rResp *demo.RegisterResp, err error) {
	reqBody, err := json.Marshal(stu)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: err=%v", err)
	}
	var resp *http.Response
	req, err := http.NewRequest(http.MethodPost, registerURL, bytes.NewBuffer(reqBody))
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

	if err = json.Unmarshal(body, &rResp); err != nil {
		return
	}
	return
}

func query(id int) (student demo.Student, err error) {
	var resp *http.Response
	resp, err = httpCli.Get(fmt.Sprint(queryURLFmt, id))
	defer resp.Body.Close()
	if err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &student); err != nil {
		return
	}
	return
}

func genStudent(id int) *demo.Student {
	return &demo.Student{
		Id:   int32(id),
		Name: fmt.Sprintf("student-%d", id),
		College: &demo.College{
			Name:    "",
			Address: "",
		},
		Email: []string{fmt.Sprintf("student-%d@nju.com", id)},
	}
}

// Assert asserts cond is true, otherwise fails the test.
func Assert(t testingTB, cond bool, val ...interface{}) {
	t.Helper()
	if !cond {
		if len(val) > 0 {
			val = append([]interface{}{"assertion failed:"}, val...)
			t.Fatal(val...)
		} else {
			t.Fatal("assertion failed")
		}
	}
}

// testingTB is a subset of common methods between *testing.T and *testing.B.
type testingTB interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
}
