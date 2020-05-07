package main

import (
  "fmt"
  "log"
  "net/http"
  "strings"

  "github.com/PuerkitoBio/goquery"
)

// ExampleScrape comment
func ExampleScrape() {
  // Request the HTML page.
  res, err := http.Get("https://www.naver.com/")
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

//   fmt.Println("qwer")
  // Find the review items
  doc.Find(".thumb_box").Each(func(i int, s *goquery.Selection) {
	  s_html, _ := s.Html()
	//   fmt.Println(s_html)
	// fmt.Println(strings.Count(s_html, "a"))
	fmt.Println(strings.Trim(s_html, "\n"))
    // For each item found, get the band and title
    // img_src, _ := s.Find("a img").Attr("src")
	// fmt.Println(img_src)
  })

  
}

func main() {
  ExampleScrape()
}