package source

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type ResData struct {
	Head   http.Header
	Body   string
	Cookie []*http.Cookie
}

type Status struct {
	GetHtmlCode bool
	Url         string
}

func DealRes(res *http.Response) *Status {

	status := &Status{}
	status.GetHtmlCode = res.StatusCode == 200

	if !status.GetHtmlCode {
		return status
	}

	resData := &ResData{}

	resData.Head = res.Header

	body, _ := ioutil.ReadAll(res.Body)
	resData.Body = string(body)
	resData.Cookie = res.Cookies()

	fmt.Println(resData.Body, 11111)

	/*
		resData.Head = res.Header
		resData.Body = res.Body
		resData.Cookie = res.Cookies()
	*/

	return status
}

/*
	@html [string]

	returns [string]
*/
func realTag(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return src
}
