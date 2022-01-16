package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, userName, password, database, authDb string) (db *mongo.Database, err error) {
	var mongoDbUrl string
	var isAuth bool
	if userName == "" || password == "" {
		mongoDbUrl = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongoDbUrl = fmt.Sprintf("mongodb://%s:%s@%s:%s", userName, password, host, password)
	}

	clientOptions := options.Client().ApplyURI(mongoDbUrl)
	if isAuth {
		if authDb == "" {
			authDb = database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDb,
			Username:   userName,
			Password:   password,
		})
	}

	connect, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)
	}

	err = connect.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err)
	}

	return connect.Database(database), nil
}
