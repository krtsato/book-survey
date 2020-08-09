package io

import "os"

type ReqParams struct {
	AppId               string
	YJobId              string
	LocalGovernmentCode int
	Fields              string
	Start               int
	Results             int
}

func SetReqParams() *ReqParams {
	appId := os.Getenv("APP_ID")
	yJobId := ""
	localGovernmentCode := 130001
	fields := ""
	start := 1
	results := 1

	return &ReqParams{
		AppId:               appId,
		YJobId:              yJobId,
		LocalGovernmentCode: localGovernmentCode,
		Fields:              fields,
		Start:               start,
		Results:             results,
	}
}
