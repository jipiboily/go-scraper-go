package main
import (
    "fmt"
    "log"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    var doc *goquery.Document
    var e error

    if doc, e = goquery.NewDocument("http://www.cotecour.ca/index.php/programmation/programmation-2013-14"); e != nil {
        log.Fatal(e)
    }

    // Find the review items (the type of the Selection would be *goquery.Selection)
    doc.Find(".event_detail_container").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the band and title
        name := s.Find("h1.event-list-title a").Text()
        desc := s.Find(".ohanah-event-short-description").Text()
        date := s.Find("h2").Text()
        fmt.Printf("Spectacle %d: %s (%s)\n\t\t%s\n\n", i + 1, name, date, desc)
    })
}
