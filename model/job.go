package model

import (
	"generate-sitemap/config"
	"log"
)

type Jora struct {
	url string
}

func GetJob() []Jora {
	db := config.DBConfig()
	rows, err := db.Query(`SELECT site_url
                    FROM
                      site`)

	if err != nil {
		log.Println("Query failed.....")
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()
	jora := []Jora{}
	for rows.Next() {
		var pageurl string
		rows.Scan(&pageurl)
		jora = append(jora, Jora{
			url: pageurl,
		},
		)

	}
	return jora
}
