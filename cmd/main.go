package main

import (
	"book-survey/internal/db"
	"book-survey/internal/io"
	"book-survey/internal/preproc"
	"book-survey/internal/services"
	"errors"
	"fmt"
	"log"
)

func run() error {
	// リクエストのパラメータを取得する
	reqParams := io.SetReqParams()

	// Rakuten API に通信する
	resBody, err := services.SendRequest(reqParams)
	if err != nil {
		return err
	}

	// JSON をデコードする
	decResBody, err := services.DecordResponse(resBody)
	if err != nil {
		return err
	}

	// 発売日の表記揺れを修正する
	preproc.FmtSalesDate(decResBody)
	fmt.Printf("%v\n", decResBody)

	// DB との接続を開始する
	dbInfo := io.SetDbInfo()
	database, err := db.Initialize(dbInfo)
	if err != nil {
		return err
	}

	// 書籍情報を DB に挿入する
	if err := db.InsertItems(database, decResBody); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(errors.Unwrap(err))
	}
}
