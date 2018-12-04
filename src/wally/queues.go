package main

import log "github.com/sirupsen/logrus"

func addUrlQ(url string) {
	l := log.WithFields(log.Fields{
		"news": url,
	})
	l.Infof("Adding site to URL Q")

	//urlQ <- url
	visit(url)
	l.Infoln("      queing complete ")
}

func addVisitQ(url string) {
	l := log.WithFields(log.Fields{
		"visit": url,
	})
	l.Infof("Add to Visit Q")
	visitQ <- url
	l.Infoln("      queing complete ")
}

// watch channels and make cool stuff happen
func watchUrlQ(urlQ chan string) {
	loop := 0
	for {
		loop = loop + 1
		log.Infof("watching channels loop %d", loop)

		select {
		case url := <-urlQ:
			l := log.WithFields(log.Fields{"URL": url})

			purl := purifyURL(url)
			if purl == "" {
				l.Errorf("failed purification")
				continue
			}

			l.Infoln("Incoming URL")
			var page *Page
			if page = filterURL(purl); page == nil {
				l.Infoln("  skipping page ...")
				continue
			}
			addVisitQ(page.URL)
		}
	}
}

func watchVisitQ(visitQ chan string) {
	for {
		select {
		case url := <-visitQ:
			log.Infoln("Got something from the visitQ ")
			go visit(url)
			log.Infoln("     visiting finished ")
		}
	}
}
