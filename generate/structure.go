package generate

import "encoding/xml"

type Sitemap struct {
	XMLName xml.Name `xml:"sitemap"`
	Jobs    []JOBS   `xml:"jobs"`
}

type JOBS struct {
	XMLName xml.Name `xml:"jobs"`
	Job     []JOB    `xml:"job"`
}

type JOB struct {
	XMLName xml.Name `xml:"job"`
	Url     string   `xml:"url"`
}
