package main

import (
	"backproj/api"
	// "fmt"
    // "log"
    // "net/http"
	// "encoding/json"
    // "github.com/gorilla/mux"
	// "io/ioutil"
)

// type Article struct {
//     Username string `json:"User"`
//     Title string `json:"Title"`
//     Content string `json:"content"`
// }

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	}
// func createNewTopic(w http.ResponseWriter, r *http.Request) {
// 	// enableCors(&w)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
//     // get the body of our POST request
//     // return the string response containing the request body    
//     // reqBody, _ := ioutil.ReadAll(r.Body)
// 	// fmt.Print("caca")
//     // fmt.Print( string(reqBody))
// 	// var article Article 
//     // json.Unmarshal(reqBody, &article)
//     // // update our global Articles array to include
//     // // our new Article
	
//     // js, _ := json.Marshal(article)
// 	// w.Write(js)
// }





func main() {
	apiHandler.InitApi()
	apiHandler.HandleRequests()
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mash:toumatim@cluster0.7cqu3.mongodb.net/collegeproj?retryWrites=true&w=majority"))
	// //client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://cluster0.7cqu3.mongodb.net/myFirstDatabase?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority"))
	// //mongodb+srv://cluster0.7cqu3.mongodb.net/myFirstDatabase?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority


	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
    // 	log.Fatal(err)

	// }
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)
}