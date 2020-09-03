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
	//
	//

	var localtest = false

	if localtest {
		fileName = "./app.db"
	} else {

		fileUrl := "https://github.com/jgermita/jgermita.me/raw/master/app.db"

		if err := DownloadFile("app_remote.db", fileUrl); err != nil {
			panic(err)
		}
		fileName = "./app_remote.db"
	}

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
		versions    string
	)

	// Get robot data
	rows, err := db.Query("select * from robots order by status asc, weight desc")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&Id, &Name, &desc, &WeightClass, &weapon, &Status, &logo, &img, &media, &rgb, &versions)
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

	// Get robot data
	rows, err := db.Query("select name, robot, finish, date, report, video from events order by date DESC, robot ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// var thisEvent string
		var thisEvent Event

		var eName string
		var eRobot string
		var eFinish string
		var eDate string
		var eReport string
		var eVideo string

		err := rows.Scan(&eName, &eRobot, &eFinish, &eDate, &eReport, &eVideo)
		if err != nil {
			log.Fatal(err)
		}
		thisEvent.Name = eName
		thisEvent.Finish = eFinish
		thisEvent.Robot = eRobot
		thisEvent.Date = eDate
		thisEvent.Report = eReport
		thisEvent.Video = eVideo

		thisEvent.VideoExists = eVideo != ""

		allEvents = append(allEvents, thisEvent)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return allEvents
}

func (database *Database) GetHighlights() []Event {
	db, err := sql.Open("sqlite3",
		fileName)
	if err != nil {
		log.Fatal(err)
	}
	var allEvents []Event

	// Get robot data
	rows, err := db.Query("select distinct video from events where video IS NOT '' order by date DESC, robot ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// var thisEvent string
		var thisEvent Event

		var eVideo string

		err := rows.Scan(&eVideo)
		if err != nil {
			log.Fatal(err)
		}
		thisEvent.Video = eVideo

		thisEvent.VideoExists = eVideo != ""

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
		versions    string
	)

	// Get robot data
	rows, err := db.Query("select * from robots where name = ?", nameQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&Id, &Name, &desc, &WeightClass, &weapon, &Status, &logo, &img, &media, &rgb, &versions)
		if err != nil {
			log.Fatal(err)
		}
		thisBot.Id = Id
		thisBot.Name = strings.Title(Name)
		thisBot.Name2 = strings.ToLower(Name)
		thisBot.WeightClass = WeightClass
		thisBot.Status = Status
		thisBot.Desc = desc
		thisBot.Rgb = rgb

		v_temp := strings.Split(versions, ",")
		for i := 0; i < len(v_temp)/2; i++ {
			j := len(v_temp) - i - 1
			v_temp[i], v_temp[j] = v_temp[j], v_temp[i]
		}

		thisBot.Versions = v_temp
		thisBot.CurrentVersion = thisBot.Versions[0]

		Wins = ""
		Losses = ""
		thisBot.Wins = 0
		winCount, err := db.Query("select sum(case when myWin=\"true\" then 1 else 0 end) from fights where myRobot = ? AND (id < 99990000)", thisBot.Name)
		if err != nil {
			thisBot.Wins = 0
		} else {
			winCount.Next()
			err = winCount.Scan(&Wins)
			thisBot.Wins, err = strconv.Atoi(Wins)

		}

		lossCount, err := db.Query("select sum(case when myWin=\"false\" then 1 else 0 end) from fights where myRobot = ? AND (id < 99990000)", thisBot.Name)

		thisBot.Losses = 0
		if err != nil {
			thisBot.Losses = 0
		} else {
			lossCount.Next()
			err = lossCount.Scan(&Losses)
			thisBot.Losses, err = strconv.Atoi(Losses)

		}

		allEvents, err := db.Query("select name, finish, date, report, video, location, website, bracket from events where robot = ? ORDER BY id DESC", thisBot.Name)
		defer allEvents.Close()
		for allEvents.Next() {
			// var thisEvent string
			var thisEvent Event

			var eName string
			var eFinish string
			var eDate string
			var eReport string
			var eVideo string
			var eWebsite string
			var eLocation string
			var eBracket string

			err := allEvents.Scan(&eName, &eFinish, &eDate, &eReport, &eVideo, &eLocation, &eWebsite, &eBracket)
			if err != nil {
				log.Fatal(err)
			}
			thisEvent.Name = eName
			thisEvent.Finish = eFinish
			thisEvent.Date = eDate
			thisEvent.Report = eReport
			thisEvent.Video = eVideo
			thisEvent.Robot = thisBot.Name
			thisEvent.Website = eWebsite
			thisEvent.Location = eLocation
			thisEvent.Bracket = eBracket

			thisEvent.VideoExists = eVideo != ""

			var loss = 0
			var win = 0
			println(thisBot.Name, thisEvent.Name)
			stmt, err := db.Prepare("select * from fights where myRobot = ? AND event = ?")
			eFights, err := stmt.Query(thisBot.Name, thisEvent.Name)
			for eFights.Next() {

				var (
					FightId       int
					Event         string
					MyRobot       string
					OpponentRobot string
					MyWin         bool
					Video         string
					Recap         string
					Type          string
				)
				err := eFights.Scan(&FightId, &Event, &MyRobot, &OpponentRobot, &MyWin, &Video, &Recap, &Type)
				if err != nil {
					log.Fatal(err)
				}

				var eFight Fight

				eFight.Id = FightId
				eFight.Event = Event
				eFight.MyRobot = MyRobot
				eFight.OpponentRobot = OpponentRobot
				eFight.MyWin = MyWin
				eFight.Video = Video
				eFight.Recap = Recap
				eFight.Type = Type
				eFight.IsRumble = Type == "Rumble"

				eFight.VideoExists = Video != ""

				if !eFight.IsRumble {

					if MyWin {
						win++
					} else {
						loss++
					}
				}

				thisEvent.Fights = append(thisEvent.Fights, eFight)
			}
			thisEvent.WinCount = win
			thisEvent.LossCount = loss
			win = 0
			loss = 0

			thisBot.Events = append(thisBot.Events, thisEvent)

		}

		thisBot.Logo = logo
		thisBot.Img = img
		thisBot.Media = media

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
		Recap         string
		Type          string
	)
	for rows.Next() {
		err := rows.Scan(&FightId, &Event, &MyRobot, &OpponentRobot, &MyWin, &Video, &Recap, &Type)
		if err != nil {
			log.Fatal(err)
		}
		thisFight := new(Fight)

		thisFight.Id = FightId
		thisFight.Event = Event
		thisFight.MyRobot = MyRobot
		thisFight.OpponentRobot = OpponentRobot
		thisFight.MyWin = MyWin
		thisFight.Video = Video
		thisFight.Type = Type

		thisFight.VideoExists = Video != ""

		thisBot.Fights = append(thisBot.Fights, *thisFight)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return thisBot
}

func (database *Database) GetFight(id string) Fight {
	db, err := sql.Open("sqlite3",
		fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Get robot data
	rows, err := db.Query("select * from fights where id = ?", id)

	var thisFight Fight
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
		Recap         string
		Type          string
	)
	for rows.Next() {
		err := rows.Scan(&FightId, &Event, &MyRobot, &OpponentRobot, &MyWin, &Video, &Recap, &Type)
		if err != nil {
			log.Fatal(err)
		}

		thisFight.Id = FightId
		thisFight.Event = Event
		thisFight.MyRobot = MyRobot
		thisFight.OpponentRobot = OpponentRobot
		thisFight.MyWin = MyWin
		thisFight.Video = Video
		thisFight.Recap = Recap
		thisFight.Type = Type

		thisFight.IsRumble = FightId > 99990000
		println(Video)

		thisFight.VideoExists = Video != ""

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return thisFight
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
		versions    string
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

		err := rows.Scan(&Id, &Name, &desc, &WeightClass, &weapon, &Status, &logo, &img, &media, &rgb, &versions)
		if err != nil {
			log.Fatal(err)
		}
		thisBot.Id = Id
		thisBot.Name = strings.Title(Name)

		thisBot.Name2 = strings.ToLower(Name)
		thisBot.WeightClass = WeightClass
		thisBot.Status = Status
		thisBot.Desc = desc
		thisBot.Rgb = rgb

		Wins = ""
		Losses = ""
		thisBot.Wins = 0
		winCount, err := db.Query("select sum(case when myWin=\"true\" then 1 else 0 end) from fights where myRobot = ? AND (id < 99990000)", thisBot.Name)
		if err != nil {
			thisBot.Wins = 0
		} else {
			winCount.Next()
			err = winCount.Scan(&Wins)
			thisBot.Wins, err = strconv.Atoi(Wins)

		}

		lossCount, err := db.Query("select sum(case when myWin=\"false\" then 1 else 0 end) from fights where myRobot = ? AND (id < 99990000)", thisBot.Name)

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

func (database *Database) GetFirstRobots() []FirstRobot {
	db, err := sql.Open("sqlite3",
		fileName)
	if err != nil {
		log.Fatal(err)
	}

	var firstRobots []FirstRobot

	// Get robot data
	rows, err := db.Query("select * from firstrobots order by year desc")

	var thisRobot FirstRobot
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		thisRobot.Video = ""
		thisRobot.Award = ""
		thisRobot.Tba = ""
		err := rows.Scan(&thisRobot.Id, &thisRobot.Team, &thisRobot.Year, &thisRobot.Name, &thisRobot.Img, &thisRobot.Video, &thisRobot.Award, &thisRobot.Role)
		if err != nil {
			log.Fatal(err)
		}

		thisRobot.Videos = strings.Split(thisRobot.Video, ",")
		thisRobot.Awards = strings.Split(strings.Trim(thisRobot.Award, "\n"), "\n")

		if strings.Contains(thisRobot.Team, "FRC") && (thisRobot.Name != "Nimbus") {
			thisRobot.Tba = "https://www.thebluealliance.com/team/" + strings.Trim(thisRobot.Team, "FRC") + "/" + thisRobot.Year
		}

		thisRobot.Name = strings.ToUpper(thisRobot.Name)
		thisRobot.Team = strings.Trim(strings.Trim(thisRobot.Team, "FTC"), "FRC")

		firstRobots = append(firstRobots, thisRobot)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return firstRobots
}
