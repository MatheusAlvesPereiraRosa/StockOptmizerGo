package mongodb

import (
	"context"

	"gear-priority-api/internal/domain"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoGearRepository struct {
	collection *mongo.Collection
}

func NewGearRepository(client *Client) *MongoGearRepository {
	return &MongoGearRepository{
		collection: client.DB.Collection("gears"),
	}
}

func (r *MongoGearRepository) Create(ctx context.Context, gear *domain.Gear) error {
	if gear.ID == uuid.Nil {
		gear.ID = uuid.New()
	}

	_, err := r.collection.InsertOne(ctx, gear)

	if err != nil {
		return err
	}

	return nil
}

func (r *MongoGearRepository) FindAll(ctx context.Context) ([]domain.Gear, error) {

	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var gears []domain.Gear

	if err := cursor.All(ctx, &gears); err != nil {
		return nil, err
	}

	return gears, nil
}

func (r *MongoGearRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Gear, error) {

	var gear domain.Gear

	err := r.collection.
		FindOne(ctx, bson.M{
			"_id": id,
		}).
		Decode(&gear)

	if err != nil {
		return nil, err
	}

	return &gear, nil
}

func (r *MongoGearRepository) Update(ctx context.Context, gear *domain.Gear) error {

	_, err := r.collection.UpdateByID(
		ctx,
		gear.ID,
		bson.M{
			"$set": gear,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *MongoGearRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	_, err := r.collection.DeleteOne(
		ctx,
		bson.M{
			"_id": id,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

// var _ domain.MongoGearRepository = (*MongoGearRepository)(nil)
