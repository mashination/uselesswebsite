package apiHandler

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
    "context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
type MongoStore struct {
	db *mongo.Database
	client *mongo.Client
	// ctx *context.Context
}

type TResult struct {
	id string `json:"id"`

}
type topicRR struct {
	Id primitive.ObjectID `json:"_id"`
	Usr string 
	Title string 
	Content string
}

type replyRR struct {
	Id primitive.ObjectID
	TopicId primitive.ObjectID
	Usr string
	Content string
}


func  Initdb()  *MongoStore {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mash:toumatim@cluster0.7cqu3.mongodb.net/collegeproj?retryWrites=true&w=majority"))
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://cluster0.7cqu3.mongodb.net/myFirstDatabase?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority"))
	//mongodb+srv://cluster0.7cqu3.mongodb.net/myFirstDatabase?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority


	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	mstore := new(MongoStore)
	mstore.db = client.Database("collegeproj")
	mstore.client = client
	// mstore.ctx = &ctx


	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
    	log.Fatal(err)

	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)


	return mstore
}

func AddTopic(mstore *MongoStore , top Topic) TResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	topicscoll := mstore.db.Collection("topics")
	fmt.Print(top)
	topicRes, err := topicscoll.InsertOne(ctx, bson.D{
			{"usr", top.Usr},
			{"title", top.Title},
			{"content", top.Content},
		})
	if err != nil {
		log.Fatal(err)
	}
	oid, ok := topicRes.InsertedID.(primitive.ObjectID)
	if !ok  {
		log.Fatal(err)
	}

	res := new(TResult)
	res.id = oid.Hex()
	return *res

}
func AddReply(mstore *MongoStore , rep Reply) TResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	topicscoll := mstore.db.Collection("replies")
	fmt.Print(rep)
	repid, _ :=primitive.ObjectIDFromHex(rep.TopicId)
	replyRes, err := topicscoll.InsertOne(ctx, bson.D{
			{"topicid", repid},
			{"usr", rep.Usr},
			{"content", rep.Content},
		})
	if err != nil {
		log.Fatal(err)
	}
	oid, ok := replyRes.InsertedID.(primitive.ObjectID)
	if !ok  {
		log.Fatal(err)
	}

	res := new(TResult)
	res.id = oid.Hex()
	return *res

}

func convTopic(top topicRR) TopicR{
	
	res := new(TopicR)
	res.Id = top.Id.Hex()
	res.Content = top.Content
	res.Title = top.Title
	res.Usr = top.Usr
	return *res
}
func convReply(top replyRR) ReplyR{
	
	res := new(ReplyR)
	res.Id = top.Id.Hex()
	res.Content = top.Content
	res.TopicId = top.TopicId.Hex()
	res.Usr = top.Usr
	return *res
}

func GetTopic(mstore *MongoStore , id string) TopicG {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	topicscoll := mstore.db.Collection("topics")
	objID, _ := primitive.ObjectIDFromHex(id)
	var resb topicRR
	if err := topicscoll.FindOne(ctx, bson.M{"_id": objID}).Decode(&resb); err!= nil {
		log.Fatal(err)
	}
	
	res := convTopic(resb)
	res2 := new(TopicG)
	res2.Usr = res.Usr
	res2.Title = res.Title
	res2.Content = res.Content
	res2.Id = id
	res2.Replies = GetReplies(mstore,id)
	return *res2

}
type HexId struct {

    ID primitive.ObjectID `bson:"_id"`
}
func GetTopics(mstore *MongoStore ) []TopicR {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var topics []TopicR
	topicscoll := mstore.db.Collection("topics")
	cursor, err := topicscoll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var topic TopicR
		var test bson.M
		var hexid HexId
		// if err = cursor.Decode(&topic); err != nil {
		// 	log.Fatal(err)
		// 
		
		if err = cursor.Decode(&test); err != nil {
			log.Fatal(err)
		}
		err = cursor.Decode(&hexid)
		bsonBytes, _ := bson.Marshal(test)
		bson.Unmarshal(bsonBytes, &topic)
		topic.Id = hexid.ID.Hex()
		fmt.Print(hexid)
		fmt.Print("hexid")

		fmt.Print(hexid.ID.Hex())


		topics = append([]TopicR{topic},  topics...)
		
	}
	fmt.Print(topics)
	return topics
}



func GetReplies(mstore *MongoStore, topid string) []ReplyR{
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var replies []ReplyR
	repliescoll := mstore.db.Collection("replies")
	objID, _ := primitive.ObjectIDFromHex(topid)
	cursor, err := repliescoll.Find(ctx, bson.M{"topicid" : objID})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var reply ReplyR
		var test bson.M
		var hexid HexId
		// if err = cursor.Decode(&topic); err != nil {
		// 	log.Fatal(err)
		// 
		
		if err = cursor.Decode(&test); err != nil {
			log.Fatal(err)
		}
		err = cursor.Decode(&hexid)
		bsonBytes, _ := bson.Marshal(test)
		bson.Unmarshal(bsonBytes, &reply)
		reply.Id = hexid.ID.Hex()
		fmt.Print(hexid)
		fmt.Print("hexid")

		fmt.Print(hexid.ID.Hex())


		replies = append(replies,  reply)
		
	}
	fmt.Print(replies)
	return replies
	// var replies []ReplyR
	// repliescoll := mstore.db.Collection("replies")
	// objID, _ := primitive.ObjectIDFromHex(topid)
	// cursor, err := repliescoll.Find(*mstore.ctx, bson.M{"topicid" : objID})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(*mstore.ctx)
	// for cursor.Next(*mstore.ctx) {
	// 	var reply replyRR
	// 	var hexid HexId
	// 	if err = cursor.Decode(&reply); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	replies = append(replies, convReply( reply))
		
	// }
	// return replies

}



