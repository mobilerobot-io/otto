package main

import (
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

func getCollector() (c *colly.Collector) {
	if collector == nil {
		c = colly.NewCollector(
			//colly.AllowedDomains(allsites...),
			//colly.AllowedDomains("gardenpassages.com"),
			colly.MaxDepth(*depth),
			colly.Async(true),
		)
		c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	}
	return c
}

// visit begins the walk of the given page with the collector
// provided as the first argument
func visit(url string) {

	c := getCollector()

	page := GetPage(url)
	loge := log.WithFields(log.Fields{
		"Site":  url,
		"State": page.State,
	})
	loge.Infoln("Entered visit .. ")

	// OnRequest is called just before the request is sent, we will create
	// our page
	c.OnRequest(func(r *colly.Request) {
		page.Start = time.Now()
		page.State = PageStateVisiting
		loge.Infof("  onRequest page starts now ")
	})

	// OnHTML+a[href] matches <a href="link">e.Text</a> tags
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if url == "" {
			return
		}
		loge.Infoln("  link found ", url)
		if p := purifyAndFilter(url); p != nil {
			loge.Infof("  new visit request  ", p.URL)
			e.Request.Visit(p.URL)
		}
	})

	log.Infoln("Visit URL ", page.URL)
	c.Visit(page.URL)

	t := time.Now()
	loge.Infof("       visit complete after %v ", time.Since(t))
	loge.Infoln("  waiting for threads ...")
	c.Wait()
	loge.Infof("The wait is over after %v GoodBYE!!!", time.Since(t))
}
