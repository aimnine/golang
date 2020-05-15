package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	tile     string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	// pages := getPages()
	// fmt.Println(pages)
	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(pageNumber int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(50*pageNumber)
	fmt.Println("Requesting: ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("data-jk")
		// fmt.Println(id)
		title := s.Find(".title>a").Text()
		// fmt.Println(title)
		location := s.Find(".sjcl").Text()
		fmt.Println(id, title, location)
	})
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	// fmt.Println(res, err)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// fmt.Println(doc)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html())
		// fmt.Println(s.Find("a").Length())
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

func cleanString(str string) string {

}
