package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type FindRecordsWithCreatedAtAndTotalCountsParams struct {
	CreatedAtAfter  time.Time
	CreatedAtBefore time.Time
	TotalCountsFrom int
	TotalCountsTo   int
}

type Row struct {
	ID         string    `bson:"_id"`
	Key        string    `bson:"key"`
	CreatedAt  time.Time `bson:"createdAt"`
	TotalCount int       `bson:"totalCount"`
}

type Repo interface {
	FindRecordsWithCreatedAtAndTotalCounts(params *FindRecordsWithCreatedAtAndTotalCountsParams) ([]Row, error)
}

// MongoRepo implements repo for MongoDB
type MongoRepo struct {
	config *Config
	client *mongo.Client
	db     *mongo.Database
}

// NewMongoRepo creates instance of MongoRepo
func NewMongoRepo(config *Config, client *mongo.Client, db *mongo.Database) Repo {
	return &MongoRepo{
		config: config,
		client: client,
		db:     db,
	}
}

func (m MongoRepo) FindRecordsWithCreatedAtAndTotalCounts(params *FindRecordsWithCreatedAtAndTotalCountsParams) ([]Row, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	unwindStage := bson.M{"$unwind": "$counts"}

	groupStage := bson.M{
		"$group": bson.M{
			"_id":        "$key",
			"key":        bson.M{"$first": "$key"},
			"createdAt":  bson.M{"$first": "$createdAt"},
			"totalCount": bson.M{"$sum": "$counts"},
		},
	}

	matchStage := bson.M{
		"$match": bson.M{
			"createdAt": bson.M{
				"$gt": primitive.NewDateTimeFromTime(params.CreatedAtAfter),
				"$lt": primitive.NewDateTimeFromTime(params.CreatedAtBefore),
			},
		},
	}

	totalCountMatch := bson.M{
		"$match": bson.M{
			"totalCount": bson.M{
				"$gt": params.TotalCountsFrom,
				"$lt": params.TotalCountsTo,
			},
		},
	}

	pipeline := []bson.M{unwindStage, matchStage, groupStage, totalCountMatch}

	cursor, err := m.db.Collection("records").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var result []Row
	err = cursor.All(ctx, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
