package mongodb

import (
	"context"
	"hex-microservices/shortener"
	"time"

	errs "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, err

}

func NewMongoRepository(mongoURL string, mongoDB string, mongoTimeout int) (*MongoRepository, error) {

	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errs.Wrap(err, "repository.NewMongoRepository")
	}

	repo := &MongoRepository{
		client:   client,
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	return repo, nil
}

func (r *MongoRepository) Find(code string) (*shortener.Redirect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	var redirect *shortener.Redirect
	redirect = &shortener.Redirect{}

	collection := r.client.Database(r.database).Collection("redirects")

	filter := bson.M{"code":code}

	err := collection.FindOne(ctx, filter).Decode(redirect)



}

func (r *MongoRepository) Store(*shortener.Redirect) error {


}