package services

import (
	"book-survey/internal/io"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func SendRequest(reqParams *io.ReqParams) ([]byte, error) {
	reqUrl := "https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404"

	// クエリを準備する
	reqVals := url.Values{}
	reqVals.Add("applicationId", reqParams.ApplicationId)
	reqVals.Add("elements", reqParams.Elements)
	reqVals.Add("formatVersion", strconv.Itoa(reqParams.FormatVersion))
	reqVals.Add("page", strconv.Itoa(reqParams.Page))
	reqVals.Add("outOfStockFlag", strconv.Itoa(reqParams.OutOfStockFlag))
	reqVals.Add("sort", reqParams.Sort)
	reqVals.Add("genreInformationFlag", strconv.Itoa(reqParams.GenreInformationFlag))

	// リクエストを作成する
	request, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("While creating the new Request: %w", err)
	}
	request.URL.RawQuery = reqVals.Encode()
	fmt.Println(request.URL.RawQuery)

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
