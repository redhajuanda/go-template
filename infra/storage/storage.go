package storage

import (
	"context"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
)

// Client is a storage client
type Client struct {
	client *storage.Client
}

var storageURL = "https://storage.googleapis.com"

// New Creates a new storage client
func New() (Client, error) {
	c, err := storage.NewClient(context.Background())
	if err != nil {
		return Client{}, err
	}
	return Client{c}, nil
}

// Upload uploads file to bucket and returns url to uploaded file
func (c Client) Upload(bucket string, object string, file io.Reader) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	wc := c.client.Bucket(bucket).Object(object).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return "", errors.Wrap(err, "error io.Copy")
	}
	if err := wc.Close(); err != nil {
		return "", errors.Wrap(err, "writer.Close")
	}

	return storageURL + "/" + bucket + "/" + object, nil
}
