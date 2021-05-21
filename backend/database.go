package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Event struct {
	ID        int       `db:"id"`
	Created   time.Time `db:"created"`
	Lecture   Lecture   `db:"lecture"`
	Date      time.Time `db:"date"`
	StartTime time.Time `db:"startTime"`
	EndTime   time.Time `db:"endTime"`
	Room      string    `db:"room"`
	Confirmed bool      `db:"confirmed"`
	Comment   string    `db:"comment"`
}

type Course struct {
	ID      int       `db:"id"`
	Created time.Time `db:"created"`
	Name    string    `db:"name"`
	Year    int       `db:"year"`
	Faculty string    `db:"faculty"`
}

type Semester struct {
	ID      int       `db:"id"`
	Created time.Time `db:"created"`
	Start   time.Time `db:"start"`
	End     time.Time `db:"end"`
	Name    string    `db:"name"`
}

type Lecture struct {
	ID       int       `db:"id"`
	Created  time.Time `db:"created"`
	Name     string    `db:"name"`
	Lecturer Lecturer  `db:"lecturer"`
	Semester Semester  `db:"semester"`
	Course   Course    `db:"course"`
}

type Lecturer struct {
	ID        int       `db:"id"`
	Created   time.Time `db:"created"`
	Surname   string    `db:"surname"`
	GivenName string    `db:"givenName"`
	Email     string    `db:"email"`
	SAMLUID   int       `db:"samlUID"`
}

type UserRights struct {
	SAMLUID   int    `db:"samlUID"`
	Email     string `db:"email"`
	IsAdmin   bool   `db:"isAdmin"`
	AddCourse bool   `db:"addCourse"`
	EditEvent bool   `db:"editEvent"`
	AddEvent  bool   `db:"addEvent"`
}

type UnavailableTime struct {
	ID        int       `db:"id"`
	Created   time.Time `db:"created"`
	StartTime time.Time `db:"startTime"`
	EndTime   time.Time `db:"endTime"`
	Lecturer  Lecturer  `db:"lecturer"`
}

type PreferredWeekday struct {
	ID        int      `db:"id"`
	Lecturer  Lecturer `db:"lecturer"`
	Monday    bool     `db:"monday"`
	Tuesday   bool     `db:"tuesday"`
	Wednesday bool     `db:"wednesday"`
	Thursday  bool     `db:"thursday"`
	Friday    bool     `db:"friday"`
	Saturday  bool     `db:"saturday"`
}

type PreferredDay struct {
	ID       int       `db:"id"`
	Created  time.Time `db:"created"`
	Lecturer Lecturer  `db:"lecturer"`
	Day      time.Time `db:"day"`
}

func openDB() {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("mysql", "user:password@/lectureplan")
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	err = tx.Commit()
	log.Println(err)
}
