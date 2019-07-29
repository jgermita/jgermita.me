package model

import (
	"database/sql"
	"log"
	"math"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var BaseDir = "."
var fileName = ""

type Database struct {
	Path string
	db   *sql.DB
}

// Open the database connection
func (database *Database) OpenDatabase(filename string) {
	fileName = "./" + filename

}

// Close database connection
func (database *Database) Close() {
	database.db.Close()
}

func (database *Database) GetRobot(nameQuery string) Robot {
	db, err := sql.Open("sqlite3",
		fileName)
	if err != nil {
		log.Fatal(err)
	}
	var thisBot Robot

	var (
		Id          int
		Name        string
		WeightClass string
		Status      string
		Wins        string
		Losses      string
		Events      string
		logo        string
		img         string
		media       string
		weapon      string
	)

	// Get robot data
	rows, err := db.Query("select * from robots where name = ?", nameQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Id, &Name, &WeightClass, &weapon, &Status, &logo, &img, &media)
		if err != nil {
			log.Fatal(err)
		}
		thisBot.Id = Id
		thisBot.Name = strings.Title(Name)
		thisBot.WeightClass = WeightClass
		thisBot.Status = Status

		winCount, err := db.Query("select sum(case when myWin=\"true\" then 1 else 0 end) from fights where myRobot = ?", thisBot.Name)
		winCount.Next()
		err = winCount.Scan(&Wins)
		thisBot.Wins, err = strconv.Atoi(Wins)
		defer winCount.Close()

		lossCount, err := db.Query("select sum(case when myWin=\"false\" then 1 else 0 end) from fights where myRobot = ?", thisBot.Name)
		lossCount.Next()
		err = lossCount.Scan(&Losses)
		thisBot.Losses, err = strconv.Atoi(Losses)
		defer lossCount.Close()

		thisBot.Events = ""

		allEvents, err := db.Query("select distinct event from fights where myRobot = ?", thisBot.Name)
		defer allEvents.Close()
		for allEvents.Next() {
			var thisEvent string
			err := allEvents.Scan(&thisEvent)
			if err != nil {
				log.Fatal(err)
			}
			thisBot.Events = thisBot.Events + thisEvent + ", "

		}

		thisBot.Logo = logo
		thisBot.Img = img
		thisBot.Media = strings.Split(media, ",")

		if len(thisBot.Media) == 1 && thisBot.Media[0] == "" {
			thisBot.Media = nil
		}
		thisBot.Weapon = weapon
		thisBot.Winrate = math.Floor(float64(thisBot.Wins)/(float64(thisBot.Wins)+float64(thisBot.Losses))*10000) / 100

		if thisBot.Wins+thisBot.Losses == 0 {
			thisBot.Winrate = 0
		}

		log.Println(Id, Name, WeightClass, Status, Wins, Losses, Events)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// get fights
	rows, err = db.Query("select * from fights where myRobot = ?", thisBot.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		FightId       int
		Event         string
		MyRobot       string
		OpponentRobot string
		MyWin         bool
		Video         string
	)
	for rows.Next() {
		err := rows.Scan(&FightId, &Event, &MyRobot, &OpponentRobot, &MyWin, &Video)
		if err != nil {
			log.Fatal(err)
		}
		thisFight := new(Fight)

		thisFight.id = FightId
		thisFight.Event = Event
		thisFight.MyRobot = MyRobot
		thisFight.OpponentRobot = OpponentRobot
		thisFight.MyWin = MyWin
		thisFight.Video = Video

		thisFight.VideoExists = Video != ""

		thisBot.Fights = append(thisBot.Fights, *thisFight)

		log.Println(FightId, MyRobot, OpponentRobot)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return thisBot
}

func (database *Database) GetAllRobots() []Robot {
	db, err := sql.Open("sqlite3",
		fileName)
	if err != nil {
		log.Fatal(err)
	}
	var allRobots []Robot

	var (
		Id          int
		Name        string
		WeightClass string
		Status      string
		Wins        string
		Losses      string
		Events      string
		logo        string
		img         string
		media       string
		weapon      string
	)

	// Get robot data
	rows, err := db.Query("select * from robots")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Id, &Name, &WeightClass, &weapon, &Status, &logo, &img, &media)
		if err != nil {
			log.Fatal(err)
		}
		thisBot := new(Robot)

		thisBot.Id = Id
		thisBot.Name = strings.Title(Name)
		thisBot.WeightClass = WeightClass
		thisBot.Status = Status

		winCount, err := db.Query("select sum(case when myWin=\"true\" then 1 else 0 end) from fights where myRobot = ?", thisBot.Name)
		winCount.Next()
		err = winCount.Scan(&Wins)
		thisBot.Wins, err = strconv.Atoi(Wins)
		defer winCount.Close()

		lossCount, err := db.Query("select sum(case when myWin=\"false\" then 1 else 0 end) from fights where myRobot = ?", thisBot.Name)
		lossCount.Next()
		err = lossCount.Scan(&Losses)
		thisBot.Losses, err = strconv.Atoi(Losses)
		defer lossCount.Close()

		thisBot.Events = ""

		allEvents, err := db.Query("select distinct event from fights where myRobot = ?", thisBot.Name)
		defer allEvents.Close()
		for allEvents.Next() {
			var thisEvent string
			err := allEvents.Scan(&thisEvent)
			if err != nil {
				log.Fatal(err)
			}
			thisBot.Events = thisBot.Events + thisEvent + ", "

		}

		thisBot.Logo = logo
		thisBot.Img = img
		thisBot.Media = strings.Split(media, ",")
		thisBot.Weapon = weapon
		thisBot.Winrate = math.Floor(float64(thisBot.Wins)/(float64(thisBot.Wins)+float64(thisBot.Losses))*10000) / 100

		if thisBot.Wins+thisBot.Losses == 0 {
			thisBot.Winrate = 0
		}

		allRobots = append(allRobots, *thisBot)

		log.Println(Id, Name, WeightClass, Status, Wins, Losses, Events)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return allRobots
}
