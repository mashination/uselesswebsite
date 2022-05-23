package apiHandler
import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
    "github.com/gorilla/mux"
	"io/ioutil"
	"time"

)


type Topic struct {
    Usr string 
    Title string 
    Content string }

type TopicR struct {
	Id string
	Usr string 
	Title string 
	Content string }

type TopicG struct {
	Id string
	Usr string 
	Title string 
	Content string 
	Replies []ReplyR
}



type Reply struct {
	Usr string
	TopicId string
	Content string
}
type ReplyR struct {
	Id string
	TopicId string
	Usr string
	Content string
}
type LaunchR struct {
	Id int
	Message string
}

type Launch struct {
	ValidAuth bool `json:"valid_auth"`
	Count     int  `json:"count"`
	Limit     int  `json:"limit"`
	Total     int  `json:"total"`
	LastPage  int  `json:"last_page"`
	Result    []struct {
		ID       int    `json:"id"`
		CosparID string `json:"cospar_id"`
		SortDate string `json:"sort_date"`
		Name     string `json:"name"`
		Provider struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"provider"`
		Vehicle struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			CompanyID int    `json:"company_id"`
			Slug      string `json:"slug"`
		} `json:"vehicle"`
		Pad struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Location struct {
				ID        int    `json:"id"`
				Name      string `json:"name"`
				State     string `json:"state"`
				Statename string `json:"statename"`
				Country   string `json:"country"`
				Slug      string `json:"slug"`
			} `json:"location"`
		} `json:"pad"`
		Missions []struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			Description interface{} `json:"description"`
		} `json:"missions"`
		MissionDescription interface{} `json:"mission_description"`
		LaunchDescription  string      `json:"launch_description"`
		WinOpen            string      `json:"win_open"`
		T0                 interface{} `json:"t0"`
		WinClose           interface{} `json:"win_close"`
		EstDate            struct {
			Month   interface{} `json:"month"`
			Day     interface{} `json:"day"`
			Year    interface{} `json:"year"`
			Quarter interface{} `json:"quarter"`
		} `json:"est_date"`
		DateStr string `json:"date_str"`
		Tags    []struct {
			ID   int    `json:"id"`
			Text string `json:"text"`
		} `json:"tags"`
		Slug             string        `json:"slug"`
		WeatherSummary   string        `json:"weather_summary"`
		WeatherTemp      float64       `json:"weather_temp"`
		WeatherCondition string        `json:"weather_condition"`
		WeatherWindMph   float64       `json:"weather_wind_mph"`
		WeatherIcon      string        `json:"weather_icon"`
		WeatherUpdated   time.Time     `json:"weather_updated"`
		Quicktext        string        `json:"quicktext"`
		Media            []interface{} `json:"media"`
		Result           int           `json:"result"`
		Suborbital       bool          `json:"suborbital"`
		Modified         time.Time     `json:"modified"`
	} `json:"result"`
}

var  mstore *MongoStore 
func InitApi() {
	mstore = Initdb()
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Content-Type", "application/json")
	}
func createNewTopic(w http.ResponseWriter, r *http.Request) {
	
	enableCors(&w)
    reqBody, _ := ioutil.ReadAll(r.Body)
    //fmt.Print(string(reqBody))
	fmt.Print("caca")
	var topic Topic
    json.Unmarshal(reqBody, &topic)
	fmt.Print(topic)
	
	res := AddTopic(mstore, topic)
	

	
	w.Header().Set("Content-Type", "application/json")
    // js, _ := json.Marshal(res)
	// w.Write(js)
	json.NewEncoder(w).Encode(res)
}
func createNewReply(w http.ResponseWriter, r *http.Request) {
	
	enableCors(&w)
    reqBody, _ := ioutil.ReadAll(r.Body)
    //fmt.Print(string(reqBody))
	fmt.Print("boudin")
	var reply Reply
    json.Unmarshal(reqBody, &reply)
	fmt.Print(reply)
	res := AddReply(mstore, reply)
	w.Header().Set("Content-Type", "application/json")
    // js, _ := json.Marshal(res)
	// w.Write(js)
	json.NewEncoder(w).Encode(res)
}
func getSingleTopic(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	vars := mux.Vars(r)
    key := vars["id"]
	res := GetTopic(mstore, key)
	json.NewEncoder(w).Encode(res)

}
func getTopics(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	res := GetTopics(mstore)
	json.NewEncoder(w).Encode(res)

}
func getRocketLaunches(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	var reply Launch
	resp, err := http.Get("https://fdo.rocketlaunch.live/json/launches/next/5")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }
	json.Unmarshal(body, &reply)
	fmt.Print(reply.Result)
	var res []LaunchR
	for _,s := range reply.Result {
		var launch LaunchR
		launch.Id = s.ID
		launch.Message = s.LaunchDescription
		res = append( res,launch)
	}
	json.NewEncoder(w).Encode(res)

	
	


}

func HandleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/topic", createNewTopic).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/reply", createNewReply).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/topics", getTopics)
	myRouter.HandleFunc("/topic/{id}", getSingleTopic)
	myRouter.HandleFunc("/launches", getRocketLaunches)

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}