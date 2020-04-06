package generate

import (
	"fmt"
	"generate-sitemap/model"
	"regexp"
)

func GenerateXml() string {
	list := model.GetJob()
	Header := `<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n"
	rowS := `<sitemap>`
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
			rowS += `<jobs>` +
				`<job>` + url[1] + `</job>` +
				`</jobs>`

		}

	}

	rowS += `</sitemap>`
	rowS += "\n"
	Closing := `</urlset>`

	xml := Header + rowS + Closing
	return xml

}
