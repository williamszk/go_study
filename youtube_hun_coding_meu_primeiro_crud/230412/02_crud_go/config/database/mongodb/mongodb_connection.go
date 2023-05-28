package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {
	// ctx := context.Background()
	mongodb_url := os.Getenv(MONGODB_URL)
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_url),
	)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	mongodb_user_db := os.Getenv(MONGODB_USER_DB)

	return client.Database(mongodb_user_db), nil
}

// func InitConnection() {
// 	ctx := context.Background()
// 	// mongodb_url :=
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err = client.Ping(ctx, nil); err != nil {
// 		panic(err)
// 	}

// 	logger.Info("Able to connect to MongoDB")
// }
