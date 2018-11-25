// Command upload saves files to blob storage on GCP and AWS.
package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"

	"github.com/google/go-cloud/blob"
)

var (
	cloud      *string = flag.String("cloud", "gcp", "cloud to use")
	bucketName *string = flag.String("bucket", "upload-bucket", "")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	// Open a connection to the bucket.
	var (
		b   *blob.Bucket
		err error
	)
	switch *cloud {
	case "gcp":
		b, err = setupGCP(ctx, *bucketName)
	case "aws":
		// AWS is handled below in the next code sample.
		b, err = setupAWS(ctx, *bucketName)
	default:
		log.Fatalf("Failed to recognize cloud. Want gcp or aws, got: %s", *cloud)
	}
	if err != nil {
		log.Fatalf("Failed to setup bucket: %s", err)
	}

	// Prepare the file for upload.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	w, err := b.NewWriter(ctx, file, nil)
	if err != nil {
		log.Fatalf("Failed to obtain writer: %s", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to bucket: %s", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("Failed to close: %s", err)
	}
}
