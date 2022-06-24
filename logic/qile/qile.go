package qile

import (
	"log"
	"sign_http/common"
)

func QiLeSign() {
	httpInfo := common.HttpReqInfo{Method: "GET", Url: "https://keylol.com/"}
	config := common.ConfigInfo{Name: "qile", Suffix: "yml"}
	httpRepInfo := common.HttpReq(httpInfo, config)
	httpRepInfo.OutputString(config.Name, "", "其乐")

	if httpRepInfo.Response.StatusCode == 200 {
		log.Println("其乐访问成功")
	}
}
