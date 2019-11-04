package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

func main() {

	urlToVisit := flag.String("url", "https://pt.wikipedia.org/wiki/Wikip%C3%A9dia:P%C3%A1gina_principal", "Put url to start webcrawler")
	numurl := flag.Int("numurl", 100, "Number of url pages to visit")

	flag.Parse()

	var stop bool
	c := colly.NewCollector()
	visitedUrl := make(map[string]bool)
	numVisited := 0
	fo, err := os.Create("output.txt")

	if err != nil{
		panic(err)
	}

	fu, err := os.Create("urls.txt")
	if err != nil {
		panic(err)
	}

	defer fo.Close()
	defer fu.Close()

	c.OnHTML("a", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		//Evitar visitar a mesma pagina
		if visitedUrl[url] != true || !stop{
			visitedUrl[e.Attr("href")] = true
			e.Request.Visit(url)
		}
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		_, err := fo.WriteString(e.Text + "\n")
		if err != nil {
			panic(err)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		if numVisited > *numurl {
			r.Abort()
			stop = true
		}

		numVisited++
		if !stop {
			fmt.Println("Visiting", r.URL)
			fu.WriteString(r.URL.String() + "\n")
		}
	})

	c.Visit(*urlToVisit)

}