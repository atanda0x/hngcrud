package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"

	"github.com/atanda0x/mirrorFinder/chapter4/gorestMetro/dbutils"
)

var DB *sql.DB

// TrainResources is the model for holding rail info
type TrainResources struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource hold info about location
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource link both trains and station
type ScheduleResource struct {
	ID        int
	TrainID   int
	StationID int
	Arrival   time.Time
}

// Register adds paths and routes to a new service instance
func (t *TrainResources) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{tain_id}").To(t.removeTrain))
	container.Add(ws)
}

func (t TrainResources) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("SELECT ID, DRIVER_NAME, OPERATING_STATUS FROM train WHERE id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		response.WriteEntity(t)
	}
}

func (t TrainResources) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResources
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)

	// Error handling is obvious here. so omitting
	statement, _ := DB.Prepare("INSERT INTO train (DRIVER_NAME, OPERATING_STATUS) VALUES (?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err != nil {
		newId, _ := result.LastInsertId()
		b.ID = int(newId)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (t TrainResources) removeTrain(requst *restful.Request, response *restful.Response) {
	id := requst.PathParameter("train-id")
	statement, _ := DB.Prepare("DELETE FROM train WHERE id=?")
	_, err := statement.Exec(id)
	if err == nil {
		response.WriteHeader(http.StatusOK)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func main() {
	// Connect to Database
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	// Create tables
	dbutils.Initilize(db)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResources{}
	t.Register(wsContainer)
	log.Printf("start listening on localhost:9000")
	server := &http.Server{Addr: ":9000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
