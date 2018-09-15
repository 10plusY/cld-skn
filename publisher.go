package cloudskine

type Publisher interface {
	Publish(Note)
}

type S3Publisher struct {
	BucketUrl string
}

func (s3p *S3Publisher) Publish(note Note) {

}
