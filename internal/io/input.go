package io

import "os"

type ReqParams struct {
	AppId string
}

func SetReqParams() *ReqParams {
	appId := os.Getenv("APP_ID")

	return &ReqParams{
		AppId: appId,
	}
}
