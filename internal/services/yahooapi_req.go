package services

import (
	"fmt"
	"io/ioutil"
	"job-survey/internal/io"
	"net/http"
	"net/url"
	"strconv"
)

func SendRequest(reqParams *io.ReqParams) ([]byte, error) {
	// URL とクエリを用意する
	reqUrl := "https://job.yahooapis.jp/v1/furusato/jobinfo/"
	reqVals := url.Values{}
	reqVals.Add("appid", reqParams.AppId)
	reqVals.Add("localGovernmentCode", strconv.Itoa(reqParams.LocalGovernmentCode))
	reqVals.Add("start", strconv.Itoa(reqParams.Start))
	reqVals.Add("results", strconv.Itoa(reqParams.Results))

	// リクエストを作成する
	request, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("While creating the new Request: %w", err)
	}
	request.URL.RawQuery = reqVals.Encode()

	// リクエストを送信する
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("While connecting Yahoo API: %w", err)
	}
	defer response.Body.Close()

	// 簡易的に status error を拾う
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Status info: %s", response.Status)
	}

	// 取得結果を変数に格納する
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("While reading the response body: %w", err)
	}

	return resBody, nil
}
