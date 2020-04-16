package model

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
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
	//	fileName = "./" + filename
	//fileName = "https://github.com/jgermita/jgermita.me/raw/master/app.db"

	fileUrl := "https://github.com/jgermita/jgermita.me/raw/master/app.db"

	if err := DownloadFile("app_remote.db", fileUrl); err != nil {
		panic(err)
	}
	fileName = "./app_remote.db"

}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Close database connection
func (database *Database) Close() {
	database.db.Close()
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
		logo        string
		img         string
		media       string
		weapon      string
		desc        string
		rgb         string
	)

	// Get robot data
	rows, err := db.Query("select * from robots order by status, id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&Id, &Name, &desc, &WeightClass, &weapon, &Status, &logo, &img, &media, &rgb)
		if err != nil {
			log.Fatal(err)
		}

		thisBot := database.GetRobot(Name)
		allRobots = append(allRobots, thisBot)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return allRobots
}

func (database *Database) GetAllEvents() []Event {
	db, err := sql.Open("sqlite3",
		fileName)
	if err != nil {
		log.Fatal(err)
	}
	var allEvents []Event

	var (
		Name   string
		Bot    string
		Finish string
		Date   string
	)

	// Get robot data
	rows, err := db.Query("select name, robot, finish, date from events order by date DESC, robot ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&Name, &Bot, &Finish, &Date)
		if err != nil {
			log.Fatal(err)
		}

		var thisEvent Event
		thisEvent.Name = Name
		thisEvent.Robot = Bot
		thisEvent.Finish = Finish
		thisEvent.Date = Date

		allEvents = append(allEvents, thisEvent)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return allEvents
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
		logo        string
		img         string
		media       string
		weapon      string
		desc        string
		rgb         string
	)

	// Get robot data
	rows, err := db.Query("select * from robots where name = ?", nameQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&Id, &Name, &desc, &WeightClass, &weapon, &Status, &logo, &img, &media, &rgb)
		if err != nil {
			log.Fatal(err)
		}
		thisBot.Id = Id
		thisBot.Name = strings.Title(Name)
		thisBot.WeightClass = WeightClass
		thisBot.Status = Status
		thisBot.Desc = desc
		thisBot.Rgb = rgb

		Wins = ""
		Losses = ""
		thisBot.Wins = 0
		winCount, err := db.Query("select sum(case when myWin=\"true\" then 1 else 0 end) from fights where myRobot = ?", thisBot.Name)
		if err != nil {
			thisBot.Wins = 0
		} else {
			winCount.Next()
			err = winCount.Scan(&Wins)
			thisBot.Wins, err = strconv.Atoi(Wins)

		}

		lossCount, err := db.Query("select sum(case when myWin=\"false\" then 1 else 0 end) from fights where myRobot = ?", thisBot.Name)

		thisBot.Losses = 0
		if err != nil {
			thisBot.Losses = 0
		} else {
			lossCount.Next()
			err = lossCount.Scan(&Losses)
			thisBot.Losses, err = strconv.Atoi(Losses)

		}

		allEvents, err := db.Query("select name, finish, date, report from events where robot = ?", thisBot.Name)
		defer allEvents.Close()

		for allEvents.Next() {
			// var thisEvent string
			var thisEvent Event

			var eName string
			var eFinish string
			var eDate string
			var eReport string

			err := allEvents.Scan(&eName, &eFinish, &eDate, &eReport)
			if err != nil {
				log.Fatal(err)
			}
			thisEvent.Name = eName
			thisEvent.Finish = eFinish
			thisEvent.Date = eDate
			thisEvent.Report = eReport
			thisBot.Events = append(thisBot.Events, thisEvent)

		}

		thisBot.Logo = logo
		thisBot.Img = img
		thisBot.Media = strings.Split(media, ",")

		if len(thisBot.Media) == 1 && thisBot.Media[0] == "" {
			thisBot.Media = nil
		}
		thisBot.Weapon = weapon
		thisBot.Winrate = math.Floor(float64(thisBot.Wins) / (float64(thisBot.Wins) + float64(thisBot.Losses)) * 100)

		if thisBot.Wins+thisBot.Losses == 0 {
			thisBot.Winrate = 0
		}

		lossCount.Close()
		winCount.Close()

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

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return thisBot
}

func (database *Database) GetRecord() string {
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
		logo        string
		img         string
		media       string
		weapon      string
		desc        string
		rgb         string
		totalWins   int
		totalLosses int
		answer      string
	)

	totalWins = 0
	totalLosses = 0

	// Get robot data
	rows, err := db.Query("select * from robots")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&Id, &Name, &desc, &WeightClass, &weapon, &Status, &logo, &img, &media, &rgb)
		if err != nil {
			log.Fatal(err)
		}
		thisBot.Id = Id
		thisBot.Name = strings.Title(Name)
		thisBot.WeightClass = WeightClass
		thisBot.Status = Status
		thisBot.Desc = desc
		thisBot.Rgb = rgb

		Wins = ""
		Losses = ""
		thisBot.Wins = 0
		winCount, err := db.Query("select sum(case when myWin=\"true\" then 1 else 0 end) from fights where myRobot = ?", thisBot.Name)
		if err != nil {
			thisBot.Wins = 0
		} else {
			winCount.Next()
			err = winCount.Scan(&Wins)
			thisBot.Wins, err = strconv.Atoi(Wins)

		}

		lossCount, err := db.Query("select sum(case when myWin=\"false\" then 1 else 0 end) from fights where myRobot = ?", thisBot.Name)

		thisBot.Losses = 0
		if err != nil {
			thisBot.Losses = 0
		} else {
			lossCount.Next()
			err = lossCount.Scan(&Losses)
			thisBot.Losses, err = strconv.Atoi(Losses)

		}

		totalWins += thisBot.Wins
		totalLosses += thisBot.Losses

		lossCount.Close()
		winCount.Close()

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	var winrate = math.Floor(float64(totalWins) / (float64(totalWins) + float64(totalLosses)) * 100)
	answer = "Overall record is " + strconv.Itoa(totalWins) + " wins and " + strconv.Itoa(totalLosses) + " losses for a winrate of " + fmt.Sprintf("%2.f", winrate) + "%"

	return answer
}
