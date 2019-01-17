package main

import (
	"encoding/json"
	"io"

	log "github.com/sirupsen/logrus"
)

func WriteJSON(wr io.Writer, obj interface{}, data []byte) (err error) {
	var jbytes []byte
	jbytes, err = json.Marshal(data)
	check(err)

	nbytes, err := wr.Write(jbytes)
	check(err)

	log.Infoln("saved json ", string(nbytes))
	return
}
