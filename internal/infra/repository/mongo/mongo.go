package mongo

import (
	"context"
	"log"

	"github.com/megalypse/golang-functional-clean/internal/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DatabaseName string = "user_service"
const CtxOperationKey string = "operation_id"

var mongoSessions map[string]mongo.Session
var client *mongo.Client

func init() {
	mongoSessions = make(map[string]mongo.Session)
	clientOptions := options.Client().ApplyURI("mockuri")

	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCollection(session mongo.Session, collectionName string) *mongo.Collection {
	return session.Client().Database(DatabaseName).Collection(collectionName)
}

func GetSession(ctx context.Context) mongo.Session {
	requestId := ctx.Value(common.MakeCtxKey(CtxOperationKey)).(string)

	session, ok := mongoSessions[requestId]
	if !ok {
		var err error
		session, err = client.StartSession()
		if err != nil {
			log.Fatal(err)
		}

		mongoSessions[requestId] = session
	}

	return session
}

func CloseSession(ctx context.Context) {
	requestId := ctx.Value(common.MakeCtxKey(CtxOperationKey)).(string)

	session, ok := mongoSessions[requestId]

	if ok {
		session.EndSession(ctx)
		mongoSessions[requestId] = nil
	}
}
