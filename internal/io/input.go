package io

import (
	"os"
	"strings"
)

type ReqParams struct {
	ApplicationId        string
	Elements             string // レスポンスが含む情報
	FormatVersion        int    // 常に version 2 を指定
	Page                 int    // 30 アイテム * ページ数
	OutOfStockFlag       int    // 1 で在庫切れも含める
	Sort                 string // 売上高順でリクエスト
	GenreInformationFlag int    // ジャンル情報を含める
}

func SetReqParams() *ReqParams {
	applicationId := os.Getenv("APP_ID")
	elementSlc := []string{
		"count", "page", "pageCount", "first", "last", "hits",
		"title", "booksGenreId", "booksGenreName", "salesDate",
		"itemPrice", "itemUrl", "reviewCount", "reviewAverage",
	}
	elements := strings.Join(elementSlc, ",")
	formatVersion := 2
	page := 1
	outOfStockFlag := 1
	sort := "sales"
	genreInformationFlag := 1

	return &ReqParams{
		ApplicationId:        applicationId,
		Elements:             elements,
		FormatVersion:        formatVersion,
		Page:                 page,
		OutOfStockFlag:       outOfStockFlag,
		Sort:                 sort,
		GenreInformationFlag: genreInformationFlag,
	}
}
