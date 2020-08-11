package db

import (
	"book-survey/internal/services"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertGenres(database *sql.DB, decResBody *services.DecResBodyType) error {
	const genreSql = "insert into genres(genre_code, name) values (?, ?)"

	for i := range decResBody.GenreInformation[0].Children {
		// genres の挿入
		res, err := database.Exec(genreSql, decResBody.GenreInformation[0].Children[i].BooksGenreID, decResBody.GenreInformation[0].Children[i].BooksGenreName)
		if err != nil {
			return fmt.Errorf("While inserting data to genres table: %w", err)
		}

		// genreId 確認のため返却
		genreId, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("While checking id after inserting genres: %w", err)
		}
		fmt.Printf("Successful insertion: genre.id = %d\n", genreId)
	}

	return nil
}

func InsertItems(database *sql.DB, decResBoby *services.DecResBodyType) error {
	const (
		reviewsSql = "insert into reviews(review_cnt, review_avg) values (?, ?)"
		itemsSql   = "insert into items(genre_id, review_id, title, price, sales_date, url) values (?, ?, ?, ?, ?, ?)"
	)

	for i := range decResBoby.Items {
		// reviews の挿入
		res, err := database.Exec(reviewsSql, decResBoby.Items[i].ReviewCount, decResBoby.Items[i].ReviewAverage)
		if err != nil {
			return fmt.Errorf("While inserting data to reviews table: %w", err)
		}

		// reviewId 利用のため返却
		reviewId, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("While checking id after inserting reviews: %w", err)
		}

		// items の挿入
		// 注 : genre_id 値の取得方法は要検討
		//			ひとまず 0 を代入している
		//			API から取得する値は e.g. "000111222333/444555666777/..."
		//			正規化して genre_id を外部キーとしたが, ここから genre_id を取得するには
		//			ヘビーな前処理と大量のトランザクションが必要になる
		res, err = database.Exec(itemsSql, 0, reviewId, decResBoby.Items[i].Title, decResBoby.Items[i].ItemPrice, decResBoby.Items[i].SalesDate, decResBoby.Items[i].ItemURL)
		if err != nil {
			return fmt.Errorf("While inserting data to items table: %w", err)
		}

		// itemId 確認のため返却
		itemId, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("While checking id after inserting items: %w", err)
		}
		fmt.Printf("Successful insertion: items.id = %d\n", itemId)
	}

	return nil
}
