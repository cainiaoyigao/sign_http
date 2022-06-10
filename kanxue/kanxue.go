package kanxue

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sign_http/common"
	"strings"
)

func KanXueSign(){
	kanXueSearch()
}

func kanXueSearch() {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://www.cordcloud.one/user/checkin",nil)
	var cookie = "__jsluid_s=0ccb11284a16418997cfef6c9906565a; __utmz=181774708.1645498217.49.4.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __jsluid_h=83bd0301e5dcda51f2c1716068b36cb3; _ga=GA1.2.1244538903.1646025979; bbs_sid=453fb97a93d507591a7c59665bc8b2ed; __utma=181774708.279083200.1639708281.1646878307.1654681017.68; __utmc=181774708; __utmt=1; bbs_token=_2BB4kbIzfxu2H_2FHyN60csncuXGTuwtaSsYiqOXGu_2FB10af6DuJ1CrBzLDqXuZt3xlMV5prTtU_2B1mdNioUqoAXg3aMQYF5OlcFBBuz9S5I69_2FER6iYb8bxCzpZNoYnm2O_2B_2F7kKYA_3D_3D; __utmb=181774708.12.9.1654681071204"
	req.Header.Add("Cookie",cookie)
	req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Encoding","gzip, deflate, br")
	req.Header.Add("Accept-Language","zh-CN,zh-TW;q=0.9,zh;q=0.8,en;q=0.7")
	req.Header.Add("Cache-Control","max-age=0")
	req.Header.Add("Connection","keep-alive")
	req.Header.Add("Host","bbs.pediy.com")
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36")
	req.Header.Add("Sec-Fetch-User","1")
	req.Header.Add("Sec-Fetch-Site","none")
	req.Header.Add("Sec-Fetch-Mode","navigate")
	req.Header.Add("Sec-Fetch-Dest","document")
	req.Header.Add("sec-ch-ua-platform","Windows")
	req.Header.Add("sec-ch-ua","\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"102\", \"Google Chrome\";v=\"102\"")
	req.Header.Add("sec-ch-ua-mobile","?0")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// gzip 解压
	body, _ := common.SwitchContentEncoding(res)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	var csrf = "";
	doc.Find("meta[name=\"csrf-token\"]").Each(func(i int, s *goquery.Selection) {
		csrf,_ = s.Attr("content")
		fmt.Printf("%d: %v\n", i,csrf)
	})

	httpReq(csrf)
}

func httpReq(csrf string){
	client := &http.Client{}

	postData := url.Values{}
	postData.Add("csrf_token", csrf)
	req, _ := http.NewRequest("POST", "https://bbs.pediy.com/user-signin.htm",strings.NewReader(postData.Encode()))
	var cookie = "__jsluid_s=0ccb11284a16418997cfef6c9906565a; __utmz=181774708.1645498217.49.4.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __jsluid_h=83bd0301e5dcda51f2c1716068b36cb3; _ga=GA1.2.1244538903.1646025979; bbs_sid=453fb97a93d507591a7c59665bc8b2ed; __utma=181774708.279083200.1639708281.1646878307.1654681017.68; __utmc=181774708; __utmt=1; bbs_token=_2BB4kbIzfxu2H_2FHyN60csncuXGTuwtaSsYiqOXGu_2FB10af6DuJ1CrBzLDqXuZt3xlMV5prTtU_2B1mdNioUqoAXg3aMQYF5OlcFBBuz9S5I69_2FER6iYb8bxCzpZNoYnm2O_2B_2F7kKYA_3D_3D; __utmb=181774708.12.9.1654681071204"
	req.Header.Add("Cookie",cookie)
	req.Header.Add("Accept","text/plain, */*; q=0.01")
	req.Header.Add("Accept-Encoding","gzip, deflate, br")
	req.Header.Add("Accept-Language","zh-CN,zh-TW;q=0.9,zh;q=0.8,en;q=0.7")
	req.Header.Add("Connection","keep-alive")
	req.Header.Add("Content-Length","43")
	req.Header.Add("Host","bbs.pediy.com")
	req.Header.Add("Content-Type","application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Origin","https://bbs.pediy.com")
	req.Header.Add("Referer","https://bbs.pediy.com/")
	req.Header.Add("Sec-Fetch-Site","same-origin")
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36")
	req.Header.Add("X-Requested-With","XMLHttpRequest")
	req.Header.Add("Sec-Fetch-Site","same-origin")
	req.Header.Add("Sec-Fetch-Mode","cors")
	req.Header.Add("Sec-Fetch-Dest","empty")
	req.Header.Add("sec-ch-ua-platform","Windows")
	req.Header.Add("sec-ch-ua","\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"102\", \"Google Chrome\";v=\"102\"")
	req.Header.Add("sec-ch-ua-mobile","?0")
	resp, err := client.Do(req)
	if err != nil {
		panic(nil)
	}
	// gzip 解压
	body, _ := common.SwitchContentEncoding(resp)

	bo, _ := ioutil.ReadAll(body)
	defer resp.Body.Close()

	val, boo:= common.JsonUnmarshal(bo)["message"]
	if boo{
		str, err := common.UnicodeToZh(val.(string))
		if err != nil{
			log.Println(err)
		}
		fmt.Printf("看雪签到: %#v\n", string(str))
	}
}