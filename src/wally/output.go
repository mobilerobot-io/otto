package main

import (
	"fmt"
	"io"
)

// PrintSiteInfo will write out site info to the writer.
func PrintPages(w io.Writer) {

	fmt.Fprintf(w, "Pages ...\n")
	pages.Range(func(key, value interface{}) bool {
		if key == nil || value == nil {
			return true
		}
		name := key.(string)
		page := value.(*Page)

		fmt.Fprintf(w, "\t%s state %4d\n", name, page.State)
		return true
	})
}
