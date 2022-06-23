package cc

import (
	"fmt"
	"log"
	"sign_http/common"
)

func CCSign() {

	httpInfo := common.HttpReqInfo{Method: "POST", Url: "https://www.cordcloud.one/user/checkin"}
	config := common.ConfigInfo{Name: "cc", Suffix: "yml"}

	httpRepInfo := common.HttpReq(httpInfo, config)
	ret, boo := httpRepInfo.BodyMap["ret"]
	if boo && "1" != common.InterfaceToString(ret) {
		fmt.Printf("CC签到 re: %#v\n", ret)
	}
	val, boo := httpRepInfo.BodyMap["msg"]
	if boo {
		str, err := common.UnicodeToZh(common.InterfaceToString(val))
		if err != nil {
			log.Println(config.Name, "err: ", err)
		}
		fmt.Printf("CC签到: %#v\n", string(str))
	}
}
