package common

import (
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"github.com/axgle/mahonia"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func SwitchContentEncoding(res *http.Response) (bodyReader io.Reader, err error) {
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		bodyReader, err = gzip.NewReader(res.Body)
	case "deflate":
		bodyReader = flate.NewReader(res.Body)
	default:
		bodyReader = res.Body
	}
	return
}

func ConvertString(src string) string {
	return ConvertToString(src, "GBK", "UTF-8")
}
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	tagCoder := mahonia.NewDecoder(tagCode)

	srcResult := srcCoder.ConvertString(src)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func UnicodeToZh(raw string) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func ZhToUnicode(sText string) string {
	textQuoted := strconv.QuoteToASCII(sText)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return textUnquoted
}

func InterfaceToString(inter interface{}) string {
	switch value := inter.(type) {
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 64)
	case int:
		return strconv.Itoa(value)
	case uint:
		return strconv.Itoa(int(value))
	case int8:
		return strconv.Itoa(int(value))
	case uint8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case uint16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case uint32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case string:
		return value
	case []byte:
		return string(value)
	default:
		newValue, _ := json.Marshal(value)
		return string(newValue)
	}
}

func JsonUnmarshal(data []byte) map[string]interface{} {
	j2 := make(map[string]interface{})
	if len(data) == 0 {
		log.Println("data is null")
		return j2
	}
	if !json.Valid(data) {
		return j2
	}
	err := json.Unmarshal(data, &j2)
	if err != nil {
		log.Println("jsonUnmarshal error : ", err.Error())
	}
	return j2
}
