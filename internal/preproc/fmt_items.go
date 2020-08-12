package preproc

import (
	"book-survey/internal/services"
	"regexp"
	"strings"
)

func FmtSalesDate(decResBody *services.DecResBodyType) {
	r := regexp.MustCompile(`(\d{4})年(\d{2})月(\d{2})*`)

	for i, item := range decResBody.Items {
		matchDate := r.FindAllStringSubmatch(item.SalesDate, -1)

		// 発売日の表記揺れを修正
		if matchDate[0][3] == "" {
			matchDate[0][3] = "00"
		}

		// [YYYY, MM, DD]
		ymdArr := []string{matchDate[0][1], matchDate[0][2], matchDate[0][3]}
		// YYYY-MM-DD
		fmtDate := strings.Join(ymdArr, "-")

		// update the struct pointer
		decResBody.Items[i].SalesDate = fmtDate
	}
}

// BooksGenreID : 000000000.../11111111.../222222222/... を整形する
