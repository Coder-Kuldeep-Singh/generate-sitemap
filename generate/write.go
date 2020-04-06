package generate

import (
	"fmt"
	"os"
)

func WriteIntoFile() {
	xml := GenerateXml()
	// everything ok now, write to file.
	filename := "sitemap-https-jora.xml"
	f, err := os.Create("./sitemaps/" + filename)
	if err != nil {
		fmt.Println("Error to create file:", err)
		return
	}
	l, err := f.WriteString(xml)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "xml written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
