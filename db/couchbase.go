package db

import "github.com/couchbase/gocb/v2"

var cluster = newCluster()

type Couchbase struct {
	BucketName string
}

type Bucket interface {
	GetBucket() *gocb.Bucket
}

func (couchbase *Couchbase) GetBucket() *gocb.Bucket {
	return cluster.Bucket(couchbase.BucketName)
}

func newCluster() *gocb.Cluster {
	cluster, err := gocb.Connect(
		"couchbase",
		gocb.ClusterOptions{
			Username: "Url",
			Password: "123qwe",
		})
	if err != nil {
		panic(err)
	}

	return cluster
}