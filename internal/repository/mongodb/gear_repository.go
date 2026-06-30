package mongodb

import (
	"context"

	"gear-priority-api/internal/dto"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoGearRepository struct {
	collection *mongo.Collection
}

func NewGearRepository(client *Client) *MongoGearRepository {
	return &MongoGearRepository{
		collection: client.DB.Collection("gears"),
	}
}

func (r *MongoGearRepository) Create(
	ctx context.Context,
	gear *dto.Gear,
) error {
	if gear.ID == uuid.Nil {
		gear.ID = uuid.New()
	}

	_, err := r.collection.InsertOne(ctx, gear)
	return err
}

func (r *MongoGearRepository) FindAll(
	ctx context.Context,
) ([]dto.Gear, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var gears []dto.Gear

	if err := cursor.All(ctx, &gears); err != nil {
		return nil, err
	}

	return gears, nil
}

func (r *MongoGearRepository) FindByID(
	ctx context.Context,
	id string,
) (*dto.Gear, error) {
	var gear dto.Gear

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

func (r *MongoGearRepository) FindByCategory(
	ctx context.Context,
	category string,
) ([]dto.Gear, error) {
	cursor, err := r.collection.Find(
		ctx,
		bson.M{"category": category},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var gears []dto.Gear

	if err := cursor.All(ctx, &gears); err != nil {
		return nil, err
	}

	return gears, nil
}

// findPaginated is private because it is only an internal MongoDB helper.
// It supports both general and filtered pagination.
func (r *MongoGearRepository) findPaginated(
	ctx context.Context,
	filter bson.M,
	page, limit int,
) ([]dto.Gear, int64, error) {
	skip := int64((page - 1) * limit)

	totalItems, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	cursor, err := r.collection.Find(
		ctx,
		filter,
		options.Find().
			SetSkip(skip).
			SetLimit(int64(limit)).
			SetSort(bson.D{
				{Key: "name", Value: 1},
			}),
	)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var gears []dto.Gear

	if err := cursor.All(ctx, &gears); err != nil {
		return nil, 0, err
	}

	return gears, totalItems, nil
}

func (r *MongoGearRepository) FindPaginated(
	ctx context.Context,
	page, limit int,
) ([]dto.Gear, int64, error) {
	return r.findPaginated(
		ctx,
		bson.M{},
		page,
		limit,
	)
}

func (r *MongoGearRepository) FindByCategoryPaginated(
	ctx context.Context,
	category string,
	page, limit int,
) ([]dto.Gear, int64, error) {
	return r.findPaginated(
		ctx,
		bson.M{"category": category},
		page,
		limit,
	)
}

func (r *MongoGearRepository) Update(
	ctx context.Context,
	gear *dto.Gear,
) error {
	_, err := r.collection.UpdateByID(
		ctx,
		gear.ID,
		bson.M{
			"$set": gear,
		},
	)

	return err
}

func (r *MongoGearRepository) Delete(
	ctx context.Context,
	id string,
) error {
	_, err := r.collection.DeleteOne(
		ctx,
		bson.M{
			"_id": id,
		},
	)

	return err
}

func (r *MongoGearRepository) DeleteAll(
	ctx context.Context,
) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{})
	return err
}
