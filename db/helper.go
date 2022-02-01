package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oauth_provider/utils"
)

//Create creating a task in a mongo or document db
func Create[T any](object T, model string) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	result, err := client.Database(model).Collection(model).InsertOne(ctx, object)
	if err != nil {
		log.Printf("Could not create %q: %v", model, err)
		return primitive.NilObjectID, err
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func CreateMany(objects []interface{}, model string) ([]primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	result, err := client.Database(model).Collection(model).InsertMany(ctx, objects)
	if err != nil {
		log.Printf("Could not create %q: %v", model, err)
		return nil, err
	}

	oids := result.InsertedIDs
	convertedOIDs := utils.ConvertArrayToAnotherType[primitive.ObjectID](oids)
	return convertedOIDs, nil
}

//Create creating a task in a mongo or document db
func GetList[T any](model string) ([]T, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cur, err := client.Database(model).Collection(model).Find(ctx, bson.D{})

	if err != nil {
		log.Printf("Could not fetch %q: %v", model, err)
		return nil, err
	}

	res := make([]T, 0)
	if err := cur.All(ctx, &res); err != nil {
		log.Printf("Could not fetch %q: %v", model, err)
		return nil, err
	}

	return res, err
}

func Get[T any](model string, id primitive.ObjectID) (T, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	var res T
	if err := client.Database(model).Collection(model).FindOne(ctx, bson.D{{"_id", id}}).Decode(&res); err != nil {
		log.Printf("Could not fetch %q: %v", model, err)
		return res, err
	}

	return res, nil
}

func Update[T any](model string, id primitive.ObjectID, object T) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	if result := client.Database(model).Collection(model).FindOneAndUpdate(ctx, bson.D{{"_id", id}}, bson.D{{"$set", object}}); result.Err() != nil {
		log.Printf("Could not update %q: %v", model, result.Err())
		return primitive.NilObjectID, result.Err()
	}

	return id, nil
}

func FindManyByAttributes[T any](model string, attributes map[string]interface{}) ([]T, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cur, err := client.Database(model).Collection(model).Find(ctx, attributes)

	if err != nil {
		log.Printf("Could not fetch %q: %v", model, err)
		return nil, err
	}

	res := make([]T, 0)
	if err := cur.All(ctx, &res); err != nil {
		log.Printf("Could not fetch %q: %v", model, err)
		return nil, err
	}

	return res, err
}

func FindByAttributes[T any](model string, attributes map[string]interface{}) (T, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	var res T
	if err := client.Database(model).Collection(model).FindOne(ctx, attributes).Decode(&res); err != nil {
		log.Printf("Could not fetch %q: %v", model, err)
		return res, err
	}

	return res, nil
}

func Delete[T any](model string, id primitive.ObjectID) (*primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	var res T
	if err := client.Database(model).Collection(model).FindOneAndDelete(ctx, bson.D{{"_id", id}}).Decode(&res); err != nil {
		log.Printf("Could not delete %q: %v", model, err)
		return nil, err
	}

	return &id, nil
}
