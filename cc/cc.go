package cc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sign_http/common"
)

func CCSign() {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://www.cordcloud.one/user/checkin", nil)
	var cookie = ""
	req.Header.Add("authority", "www.cordcloud.one")
	req.Header.Add("method", "POST")
	req.Header.Add("path", "/user/checkin")
	req.Header.Add("scheme", "https")
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", "")
	req.Header.Add("Accept-Language", "zh-CN,zh-TW;q=0.9,zh;q=0.8,en;q=0.7")
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Origin", "https://www.cordcloud.one")
	req.Header.Add("Referer", "https://www.cordcloud.one/user")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("sec-ch-ua-platform", "Windows")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"102\", \"Google Chrome\";v=\"102\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: \v", err)
	}

	bo, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	val, boo := common.JsonUnmarshal(bo)["msg"]
	if boo {

		str, err := common.UnicodeToZh(val.(string))
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("CC签到: %#v\n", string(str))
	}
}
