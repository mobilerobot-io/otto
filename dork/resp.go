package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Response contains an io.Writer and a format to tell
// it how to respond with queries
type Response struct {
	Format string
	io.Writer
}

func getResponse(output string, format string) (resp *Response) {

	resp = &Response{
		Format: format,
		Writer: os.Stdout,
	}
	return resp
}

func (resp *Response) Respond(obj interface{}) {
	switch config.Format {
	case "json":
		resp.JSON(obj)
	case "dot":
		//resp.DOT(obj)
	default:
		resp.Text(obj)
	}
}

func (r *Response) Text(obj interface{}) {
	fmt.Fprintf(r.Writer, "%v", obj)
}

func (r *Response) JSON(obj interface{}) {
	jbuf, err := json.Marshal(obj)
	panicIfError(err)
	_, err = r.Write(jbuf)
	panicIfError(err)
}

/*
func (r *Response) DOT(obj interface{}) {
	dbuf, err := dot.Marshal(network.Undirect, "digital-ocean", "", "  ")
	panicIfError(err)
	_, err = r.Write(dbuf)
	panicIfError(err)
}
*/
