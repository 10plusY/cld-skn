package cloudskine

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

type Publisher struct {
    Publish()
}

type S3Publisher struct {
    BucketUrl string
}

func (s3p *S3Publisher) Publish() {
    
}
