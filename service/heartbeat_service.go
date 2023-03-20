package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"storage/conf"
	"storage/consts"
	"storage/data"
	"storage/model"
	"time"
)

type HeartbeatService struct {
	mongoClient *mongo.Client
}

func NewHeartbeatService() (*HeartbeatService, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getURI()))
	if err != nil {
		return nil, err
	}

	return &HeartbeatService{mongoClient: client}, nil
}

func (hs *HeartbeatService) InitReport() error {
	return insertStorage(hs.mongoClient)
}

func (hs *HeartbeatService) StartReport() {
	go func() {
		for true {
			go func() {
				err := hs.updateStatus()
				if err != nil {
					log.Println("UpdateStatus failed: ", err)
				}
			}()
			time.Sleep(conf.HeartbeatInternal * time.Second)
		}
	}()
}

func (hs *HeartbeatService) updateStatus() error {
	return updateStorage(hs.mongoClient)
}

func insertStorage(mongoClient *mongo.Client) error {
	storage := &model.Storage{
		StorageAddress: conf.IP + ":" + conf.PORT,
		GroupID:        conf.GROUP,
		Capacity:       conf.Capacity,
		AvailableCap:   conf.Capacity,
		UpdateTime:     time.Now().Unix(),
	}
	_, err := mongoClient.Database(conf.DATABASE).Collection(consts.Storages).InsertOne(context.TODO(), storage)

	return err
}

func updateStorage(mongoClient *mongo.Client) error {
	filter := bson.M{"storageAddress": conf.IP + ":" + conf.PORT}
	data.AvailableCapLock.Lock()
	update := bson.M{"$set": bson.M{"availableCap": data.AvailableCap, "updateTime": time.Now().Unix()}}
	data.AvailableCapLock.Unlock()
	_, err := mongoClient.Database(conf.DATABASE).Collection(consts.Storages).UpdateOne(context.TODO(), filter, update)

	return err
}

/*
func getURI() string {
	protocol := conf.PROTOCOL
	username := conf.USERNAME
	password := conf.PASSWORD
	address := conf.ADDRESS
	authentication := conf.AUTHENTICATION

	return fmt.Sprintf("%s://%s:%s@%s/%s", protocol, username, password, address, authentication)
}
*/

func getURI() string {
	return "mongodb://localhost:27017"
}
