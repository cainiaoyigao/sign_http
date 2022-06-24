package cc

import (
	"sign_http/common"
)

func CCSign() {

	httpInfo := common.HttpReqInfo{Method: "POST", Url: "https://www.cordcloud.one/user/checkin"}
	config := common.ConfigInfo{Name: "cc", Suffix: "yml"}
	httpRepInfo := common.HttpReq(httpInfo, config)
	httpRepInfo.OutputString(config.Name, "msg", "cc")
}
