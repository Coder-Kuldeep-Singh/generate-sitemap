package generate

import (
	"fmt"
	"generate-sitemap/model"
	"regexp"
	"strings"
	"time"
)

func Today() string {
	a := strings.Split(time.Now().UTC().Format(time.RFC3339), "T")
	return a[0]
}

func GenerateXml() string {
	list := model.GetJob()
	Header := `<?xml version="1.0" encoding="utf-8"?>
	<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n"
	rowS := `<jobs>`
	for _, li := range list {

		str := fmt.Sprintf("%#v", li)
		regex := regexp.MustCompile(`url:"(.*)"`)
		extratingurl := regex.FindAllStringSubmatch(str, -1)
		for _, url := range extratingurl {
			// fmt.Println(url[1])
			// replaceSpace := strings.Replace(url[1], " ", "", -1)
			// v := &jobs{}
			// v.Jobs = append(v.Jobs, JOB{Url: url[1]})
			// e, err := xml.Marshal(v)
			// if err != nil {
			// 	fmt.Println(err)
			// 	return ""
			// }
			// urlslice += string(e) + "\n"
			// rowS += `<job>` +
			// 	`<url>` + url[1] + `</url>` +
			// 	`<lastmod>` + Today() + "</lastmod>" +
			// 	`<changefreq>daily</changefreq>` +
			// 	`<priority>0.6</priority>` +
			// 	`</job>`
			rowS += fmt.Sprintf(`
				<job>
					<url> %s </url>
					<lastmod> %s </lastmod>
					<changefreq>daily</changefreq>
					<priority>0.6</priority>
				</job>
				`, url[1], Today())
		}

	}

	rowS += `</jobs>`
	rowS += "\n"
	Closing := `</urlset>`

	xml := Header + rowS + Closing
	return xml

}
