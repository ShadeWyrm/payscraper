package main

var groupURLs = []string{
	"https://www.tbs-sct.gc.ca/agreements-conventions/view-visualiser-eng.aspx?id=4#rates-ec",
	"https://www.tbs-sct.gc.ca/agreements-conventions/view-visualiser-eng.aspx?id=1#rates-cs",
	"https://www.tbs-sct.gc.ca/agreements-conventions/view-visualiser-eng.aspx?id=10#rates-fb",
	"https://www.tbs-sct.gc.ca/agreements-conventions/view-visualiser-eng.aspx?id=11#rates-fi",
	"https://www.tbs-sct.gc.ca/agreements-conventions/view-visualiser-eng.aspx?id=15#rates-as",
	"https://www.tbs-sct.gc.ca/agreements-conventions/view-visualiser-eng.aspx?id=15#rates-pm",
	"https://www.tbs-sct.gc.ca/agreements-conventions/view-visualiser-eng.aspx?id=15#rates-is",
}

func main() {

	for _, url := range groupURLs {

		g := Group{
			Identifier: url[len(url)-2:],
			URL:        url,
		}

		GetPayScales(url, &g)
	}
	// fmt.Println(groups) or do something else with them
}
