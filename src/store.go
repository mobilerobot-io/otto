package otto

import (
	"encoding/json"
	"io"

	log "github.com/sirupsen/logrus"
)

func SaveJSON(wr io.Writer, obj interface{}, data []byte) (err error) {
	var jbytes []byte
	jbytes, err = json.Marshal(data)
	fatalIfError(err)

	nbytes, err := wr.Write(jbytes)
	fatalIfError(err)

	log.Infoln("saved json ", string(nbytes))
	return
}
