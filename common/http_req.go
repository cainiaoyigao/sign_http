package common

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpReq(httpReq HttpReqInfo, info ConfigInfo) HttpRepInfo {
	client := &http.Client{}

	req, _ := http.NewRequest(httpReq.Method, httpReq.Url, httpReq.Body)

	headMap := ReadConfig(info.Name, info.Suffix)

	for key, value := range headMap {
		req.Header.Add(key, InterfaceToString(value))
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: \v", err)
	}

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("url %#v req fail code:%#v \n", httpReq.Url, InterfaceToString(resp.StatusCode)))
	}

	bo, err := ioutil.ReadAll(resp.Body)
	//defer resp.Body.Close()
	return HttpRepInfo{
		Response: resp,
		BodyMap:  JsonUnmarshal(bo),
	}
}

type HttpReqInfo struct {
	Method string
	Url    string
	Body   io.Reader
}

type HttpRepInfo struct {
	Response *http.Response
	BodyMap  map[string]interface{}
}
