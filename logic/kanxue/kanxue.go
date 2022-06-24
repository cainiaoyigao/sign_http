package kanxue

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"sign_http/common"
	"strings"
)

func KanXueSign() {
	kanXueSearch()
}

func kanXueSearch() {
	httpInfo := common.HttpReqInfo{Method: "POST", Url: "https://www.cordcloud.one/user/checkin"}
	config := common.ConfigInfo{Name: "kanxue_main", Suffix: "yml"}

	httpRepInfo := common.HttpReq(httpInfo, config)
	// gzip 解压
	//body, _ := common.SwitchContentEncoding(res)
	doc, err := goquery.NewDocumentFromReader(httpRepInfo.Response.Body)
	if err != nil {
		log.Fatal(config.Name, "err: ", err)
	}
	var csrf = ""
	doc.Find("meta[name=\"csrf-token\"]").Each(func(i int, s *goquery.Selection) {
		csrf, _ = s.Attr("content")
		fmt.Printf("%d: %v\n", i, csrf)
	})

	httpReq(csrf)
}

func httpReq(csrf string) {
	postData := url.Values{}
	postData.Add("csrf_token", csrf)
	httpInfo := common.HttpReqInfo{Method: "POST", Url: "https://bbs.pediy.com/user-signin.htm", Body: strings.NewReader(postData.Encode())}
	config := common.ConfigInfo{Name: "kanxue_sign", Suffix: "yml"}
	httpRepInfo := common.HttpReq(httpInfo, config)

	httpRepInfo.OutputString(config.Name, "message", "看雪")
}
