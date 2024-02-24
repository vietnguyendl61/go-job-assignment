package repo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pricing-service/model"
	"time"
)

const (
	generalQueryTimeout = 600 * time.Second
)

type PriceRepo struct {
	mongodb *mongo.Client
}

func NewPriceRepo(mongodb *mongo.Client) PriceRepo {
	return PriceRepo{mongodb: mongodb}
}

func (r PriceRepo) GetPriceByJobId(ctx context.Context, jobId string) (*model.Price, error) {
	var price *model.Price
	filter := bson.D{{"job_id", jobId}}
	err := r.mongodb.Database("price").Collection("price").FindOne(ctx, filter).Decode(&price)
	if err != nil {
		return nil, err
	}

	return price, nil
}

func (r PriceRepo) CreatePriceMongo(ctx context.Context, price *model.Price) error {
	price.ID = uuid.New().String()
	price.CreatedAt = time.Now().String()
	price.UpdatedAt = time.Now().String()

	dataInsert, err := toDoc(price)

	_, err = r.mongodb.Database("price").Collection("price").InsertOne(ctx, dataInsert)
	if err != nil {
		return err
	}

	return nil
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
