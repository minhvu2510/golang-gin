package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)
func EncodeParam(s string) string {
	return url.QueryEscape(s)
}
//func Notify(message string) {
func Notify() {
	address := EncodeParam("<a href=\"http://www.example.com/\">inline URL</a>5 <strong>ERROR(S)</strong><b>bold <i>italic bold <s>italic bold strikethrough</s> \n <u>underline italic bold</u></i> bold</b>")

	fmt.Sprintf("I live in %v", address)


	url_base := "https://api.telegram.org/bot2032472682:AAEo2UkQ6k3N5TERvR3f3FIQMxnWC705qQY/sendMessage?chat_id=-1001607522369&text=%s&parse_mode=html"

	url_pr := fmt.Sprintf(url_base, address)
	fmt.Println(url_pr)
	method := "GET"

	//url_i := url.URL{}
	//url_proxy, _ := url_i.Parse("10.60.117.113:8080")

	proxyStr := "http://10.60.117.113:8080"
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Println(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	timeout := time.Duration(7 * time.Second)
	client := &http.Client {
		Timeout: timeout,
		Transport: transport,
	}
	req, err := http.NewRequest(method, url_pr, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	log.Println(string(body))
}