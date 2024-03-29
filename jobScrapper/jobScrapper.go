package jobscrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id string
	title string
	location string
	summary string
}


// Scrape Saramin by a term
func Scrape(term string){
	var baseUrl string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseUrl)
	
	for i := 0; i < totalPages; i++ {
		go getPage(i, baseUrl, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJob := <- c
		jobs = append(jobs, extractedJob...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))

}



func getPage(page int, url string, mainC chan <- []extractedJob ) {
var jobs []extractedJob
c := make(chan extractedJob)
 pageURL := url + "&ecruitPage=" + strconv.Itoa(page + 1)
 fmt.Println("Requesting", pageURL)

 res, err := http.Get(pageURL)
 checkErr(err)
 checkCode(res)

 defer res.Body.Close()

 doc, err := goquery.NewDocumentFromReader(res.Body)
 checkErr(err)

 searchCards := doc.Find(".item_recruit")
 searchCards.Each(func(i int, card *goquery.Selection){
	go extractJob(card, c)
})

for i := 0; i < searchCards.Length(); i++ {
	job := <-c
	jobs = append(jobs, job)
}
 mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan <- extractedJob){
	id, _ := card.Attr("value")
	title := CleanString(card.Find(".job_tit>a>span").Text())
	location := CleanString(card.Find(".job_condition>span>a").Text())
	summary := CleanString(card.Find(".job_sector>a").Text())

	c <- extractedJob{
		id: id,
		title: title,
		location: location,
		summary: summary,
	}
}

// Clean string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace((str)))," ")
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
		pages = s.Find("a").Length()
	})

	return pages
}

func writeJobs(jobs []extractedJob){
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func checkErr(err error){
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response ){
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status:", res.StatusCode)
	}
}