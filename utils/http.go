package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"

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
