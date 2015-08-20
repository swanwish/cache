package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/swanwish/go-helper/logs"
)

func GetUrlContent(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		logs.Errorf("Failed to get content from url %s, the error is %v", url, err)
		return nil, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Errorf("Failed to read content from response, the error is %v\n", err)
		return nil, err
	}
	return contents, nil
}

func PostUrlContent(url string, content []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logs.Errorf("Failed to post content, the error is %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	logs.Debugf("The content is %s", string(body))
	return body, nil
}

func PostForm(url string, data url.Values, headers http.Header) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
	if err != nil {
		logs.Errorf("Failed to get request, the error is %v", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	if headers != nil {
		for key, value := range headers {
			req.Header[key] = value
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		logs.Errorf("Failed to post content, the error is %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	logs.Debugf("The content is %s", string(body))
	return body, nil
}
