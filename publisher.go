package cloudskine

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"

    "os"
    "fmt"
    "bytes"
    "net/http"
)

type Publisher interface {
    Publish() error
}

type S3Publisher struct {
    BucketUrl string
    Session *session.Session
    Logger Logger
}

func getPutObject(file *os.File, size int64) (interface{}, error) {
    buffer := make([]byte, size)
    file.Read(buffer)
    return s3.New(s3p.Session).PutObject(&s3.PutObjectInput{
       Bucket: aws.String(s3p.BucketUrl),
       Key: aws.String(file.Name()),
       ACL: aws.String("private"),
       Body: bytes.NewReader(buffer),
       ContentLength: aws.Int64(size),
       ContentType: aws.String(http.DetectContentType(buffer)),
       ContentDisposition: aws.String("attachment"),
       ServerSideEncryption: aws.String("AWS256"),
   })

}

func (s3p *S3Publisher) Publish(file *os.File) error {
    var info os.FileInfo
    var size int64
    var err error

    defer file.Close()

    if info, err = file.Stat(); err != nil {
        s3p.Logger.Log((fmt.Sprintf("Error: %s", err)))
    } else {
        size = info.Size()
    }

    buffer := make([]byte, size)
    file.Read(buffer)

    _, err = getPutObject()
    return err




}
