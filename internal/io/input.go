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
	page := 1                 // 30 書籍 * 40 ページ
	outOfStockFlag := 1       // 在庫切れを含める
	sort := "sales"           // 売上が高い順のレスポンスを求める
	genreInformationFlag := 1 // ジャンルの情報を含める

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

type DbInfoType struct {
	RootPw   string
	User     string
	Pw       string
	Database string
	Host     string
}

func SetDbInfo() *DbInfoType {
	rootPw := os.Getenv("MYSQL_ROOT_PASSWORD")
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")

	return &DbInfoType{
		RootPw:   rootPw,
		User:     user,
		Pw:       pw,
		Database: database,
		Host:     host,
	}
}
