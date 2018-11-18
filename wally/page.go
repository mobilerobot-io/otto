package main

import (
	"time"

	"github.com/PuerkitoBio/purell"
)

const (
	PageStateNone = iota
	PageStateReady
	PageStateVisiting
	PageStateVisited
	PageStateErrored
	PageStateBlocked
	PageStateNotFound
	PageStateRedirected
)

//                            Page
// ====================================================================

// Page is used to keep tabs on things
type Page struct {

	// Page identification
	URL string `json:"url"`

	// Document Structure - todo ...
	Title string         `json:"title"` // Document structure
	Links map[string]int `json:"links"` // Links contained in this page

	// Visitation stats ~ start and finish crawl times for last visit
	Created time.Time `json:"created"`
	Start   time.Time `json:"start"`
	Finish  time.Time `json:"finish"`

	// Page State
	State int

	LastError string `json:"LastError"` // "" if no error

	Accessed int `json:-` // How many times has this page/link been looked up
}

func (p *Page) String() (str string) {
	str = p.URL + "links (" + string(len(p.Links)) + ") ..."
	return str
}

// NewPage returns a new page to the given url
func NewPage(url string) *Page {
	p := &Page{
		URL:     url,
		Links:   make(map[string]int),
		Created: time.Now(),
		State:   PageStateNone,
	}
	pages.Store(url, p)
	return p
}

// GetPage will return an existing page or create a new one representing
// the URL parameter
func GetPage(url string) (p *Page) {
	if i, ok := pages.Load(url); !ok {
		p = NewPage(url)
	} else {
		p = i.(*Page)
	}
	return p
}

func purifyURL(urlstr string) (purl string) {
	var err error
	flags := purell.FlagsUsuallySafeGreedy
	if purl, err = purell.NormalizeURLString(urlstr, flags); err != nil {
		panic(err)
	}
	return purl
}

//  Filter URL
// ====================================================================

// filterURL will determine if a given URL will be placed on the
// process queue
func filterURL(url string) (page *Page) {
	if p, ok := pages.Load(url); !ok {
		page = NewPage(url)
	} else {
		page = p.(*Page)
	}
	page.Accessed++
	if page.State != PageStateReady {
		return nil
	}
	return page
}

// purifyAndFilter the url
func purifyAndFilter(url string) (page *Page) {
	if purl := purifyURL(url); purl != "" {
		if page = filterURL(purl); page != nil {
			return page
		}
	}
	return nil
}
