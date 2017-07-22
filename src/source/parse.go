package source

import (
	"bufio"
	"fmt"
	"io"
	// "io/ioutil"
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

func DealRes(res *http.Response, keys []string) *Status {

	status := &Status{}
	status.GetHtmlCode = res.StatusCode == 200

	if !status.GetHtmlCode {
		return status
	}

	resData := &ResData{}

	resData.Head = res.Header
	resData.Cookie = res.Cookies()

	body := bufio.NewReader(res.Body)

	for {
		buf, _, state := body.ReadLine()

		if state == io.EOF {
			break
		}

		strs := string(buf)

		urls := parseURL(strs)
		infos := parseINFO(strs, keys)

		InsertINFO(infos)

		for _, url := range *urls {

			if !strings.Contains(url, "com") {
				url = res.Request.URL.Host + url
				InsertURL(url)
			}
		}
	}

	return status
}

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

func parseURL(line string) *[]string {

	urls := make([]string, 0)
	p_urls := &urls

	for _, str := range strings.Split(line, " ") {

		if strings.Contains(str, "href") {

			str = strings.NewReplacer("href=\"", "", "\"", "").Replace(str)

			if strings.ContainsAny(str, ">") {

				continue

			}

			*p_urls = append(*p_urls, str)

		} else {

			continue

		}

	}

	return p_urls
}

func parseINFO(str string, keys []string) *[]string {
	strs := make([]string, 0)
	p_strs := &strs

	str = realTag(str)

	for i := 0; i < len(keys); i++ {
		k := keys[i]

		if strings.Contains(str, k) {
			fmt.Println(str)
			*p_strs = append(*p_strs, str)

		}

	}

	return p_strs
}
