package cloudskine

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"bytes"
	"fmt"
	"net/http"
	"os"
)

type Publisher interface {
	Prepare(*os.File, int64) (interface{}, error)
	Publish(*os.File) error
}

type S3Publisher struct {
	BucketUrl string
	Session   *session.Session
	Logger    Logger
}

func (s3p *S3Publisher) Prepare(file *os.File, size int64) (interface{}, error) {
	buffer := make([]byte, size)
	file.Read(buffer)
	return s3.New(s3p.Session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s3p.BucketUrl),
		Key:                  aws.String(file.Name()),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AWS256"),
	})

}

func (s3p *S3Publisher) Publish(file *os.File) error {
	defer file.Close()

	if info, err := file.Stat(); err != nil {
		s3p.Logger.Log((fmt.Sprintf("Error: %s", err)), true)
		return err
	}

	_, err := getPutObject(file, info.Size())
	return err
}
