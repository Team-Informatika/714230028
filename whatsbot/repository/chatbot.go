package repository

import (
	"context"
	"time"
	"whatsbot/config"
	"whatsbot/model"

	"go.mongodb.org/mongo-driver/bson"
)

func GetKeywordsFromDB() ([]model.Keyword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.KeywordCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var keywords []model.Keyword
	for cursor.Next(ctx) {
		var keyword model.Keyword
		if err := cursor.Decode(&keyword); err != nil {
			return nil, err
		}
		keywords = append(keywords, keyword)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return keywords, nil
}
