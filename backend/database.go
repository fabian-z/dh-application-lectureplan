package main

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Time struct {
	Hour   int8
	Minute int8
	Second int8
}

// Value - Implementation of valuer for database/sql
func (t *Time) Value() (driver.Value, error) {
	return fmt.Sprintf("%d:%d:%d", t.Hour, t.Minute, t.Second), nil
}

// Scan - Implement the database/sql scanner interface
func (t *Time) Scan(value interface{}) error {
	// if value is nil, false
	if value == nil {
		// set the value of the pointer yne to YesNoEnum(false)
		*t = Time{0, 0, 0}
		return nil
	}
	res, err := driver.String.ConvertValue(value)

	if err != nil {
		return err
	}

	timeString := ""

	switch resTyped := res.(type) {
	case string:
		timeString = resTyped
	case []byte:
		timeString = string(resTyped)
	default:
		return errors.New("invalid time type from driver")
	}

	timeSlice := strings.Split(timeString, ":")
	hour, err := strconv.ParseInt(timeSlice[0], 10, 8)
	if err != nil {
		return fmt.Errorf("error parsing hour: %w", err)
	}
	minute, err := strconv.ParseInt(timeSlice[1], 10, 8)
	if err != nil {
		return fmt.Errorf("error parsing minute: %w", err)
	}
	second, err := strconv.ParseInt(timeSlice[2], 10, 8)
	if err != nil {
		return fmt.Errorf("error parsing seconds: %w", err)
	}

	t.Hour = int8(hour)
	t.Minute = int8(minute)
	t.Second = int8(second)

	return nil
}

type Event struct {
	ID        int       `db:"id"`
	Created   time.Time `db:"created"`
	Lecture   Lecture   `db:"-"`
	LectureID int       `db:"lecture"`
	Date      time.Time `db:"date"`
	StartTime Time      `db:"startTime"`
	EndTime   Time      `db:"endTime"`
	Room      string    `db:"room"`
	Confirmed bool      `db:"confirmed"`
	Comment   string    `db:"comment"`
}

func (e *Event) Get(tx *sqlx.Tx, query string, args ...interface{}) error {

	err := tx.Get(e, query, args...)
	if err != nil {
		return err
	}

	return e.Lecture.Get(tx, "SELECT * FROM lectures WHERE id=?", e.LectureID)
}

type Course struct {
	ID      int       `db:"id"`
	Created time.Time `db:"created"`
	Name    string    `db:"name"`
	Year    int       `db:"year"`
	Faculty string    `db:"faculty"`
}

func (c *Course) Get(tx *sqlx.Tx, query string, args ...interface{}) error {
	err := tx.Get(c, query, args...)
	if err != nil {
		return err
	}
	return nil
}

type Semester struct {
	ID      int       `db:"id"`
	Created time.Time `db:"created"`
	Start   time.Time `db:"start"`
	End     time.Time `db:"end"`
	Name    string    `db:"name"`
}

func (s *Semester) Get(tx *sqlx.Tx, query string, args ...interface{}) error {
	return tx.Get(s, query, args...)
}

type Lecture struct {
	ID      int       `db:"id"`
	Created time.Time `db:"created"`
	Name    string    `db:"name"`

	LecturerID int      `db:"lecturer"`
	Lecturer   Lecturer `db:"-"`

	SemesterID int      `db:"semester"`
	Semester   Semester `db:"-"`

	CourseID int    `db:"course"`
	Course   Course `db:"-"`
}

func (l *Lecture) Get(tx *sqlx.Tx, query string, args ...interface{}) error {

	err := tx.Get(l, query, args...)
	if err != nil {
		return err
	}

	err = l.Lecturer.Get(tx, "SELECT * FROM lecturers WHERE id=?", l.LecturerID)
	if err != nil {
		return err
	}

	err = l.Semester.Get(tx, "SELECT * FROM semesters WHERE id=?", l.SemesterID)
	if err != nil {
		return err
	}

	err = l.Course.Get(tx, "SELECT * FROM courses WHERE id=?", l.CourseID)
	if err != nil {
		return err
	}

	return nil
}

type Lecturer struct {
	ID        int           `db:"id"`
	Created   time.Time     `db:"created"`
	Surname   string        `db:"surname"`
	GivenName string        `db:"givenName"`
	Email     string        `db:"email"`
	SAMLUID   sql.NullInt64 `db:"samlUID"`
}

func (l *Lecturer) Get(tx *sqlx.Tx, query string, args ...interface{}) error {
	return tx.Get(l, query, args...)
}

type UserRights struct {
	SAMLUID   int    `db:"samlUID"`
	Email     string `db:"email"`
	IsAdmin   bool   `db:"isAdmin"`
	AddCourse bool   `db:"addCourse"`
	EditEvent bool   `db:"editEvent"`
	AddEvent  bool   `db:"addEvent"`
}

func (ur *UserRights) Get(tx *sqlx.Tx, query string, args ...interface{}) error {
	return tx.Get(ur, query, args...)
}

type UnavailableTime struct {
	ID        int       `db:"id"`
	Created   time.Time `db:"created"`
	StartTime time.Time `db:"startTime"`
	EndTime   time.Time `db:"endTime"`

	Lecturer   Lecturer `db:"-"`
	LecturerID int      `db:"lecturer"`
}

func (ut *UnavailableTime) Get(tx *sqlx.Tx, query string, args ...interface{}) error {

	err := tx.Get(ut, query, args...)
	if err != nil {
		return err
	}

	return ut.Lecturer.Get(tx, "SELECT * FROM lecturers WHERE id=?")
}

type PreferredWeekday struct {
	ID         int      `db:"id"`
	Lecturer   Lecturer `db:"-"`
	LecturerID int      `db:"lecturer"`
	Monday     bool     `db:"monday"`
	Tuesday    bool     `db:"tuesday"`
	Wednesday  bool     `db:"wednesday"`
	Thursday   bool     `db:"thursday"`
	Friday     bool     `db:"friday"`
	Saturday   bool     `db:"saturday"`
}

func (pw *PreferredWeekday) Get(tx *sqlx.Tx, query string, args ...interface{}) error {

	err := tx.Get(pw, query, args...)
	if err != nil {
		return err
	}

	return pw.Lecturer.Get(tx, "SELECT * FROM lecturers WHERE id=?")
}

type PreferredDay struct {
	ID       int       `db:"id"`
	Created  time.Time `db:"created"`
	Lecturer Lecturer  `db:"lecturer"`
	Day      time.Time `db:"day"`
}

func (pd *PreferredDay) Get(tx *sqlx.Tx, query string, args ...interface{}) error {

	err := tx.Get(pd, query, args...)
	if err != nil {
		return err
	}

	return pd.Lecturer.Get(tx, "SELECT * FROM lecturers WHERE id=?")
}

func openDB() {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("mysql", "mysql:mysql@/lectureplan?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	tx := db.MustBegin()

	var event Event
	err = event.Get(tx, "SELECT * FROM events WHERE id=?", 2)
	if err != nil {
		log.Println(err)
		tx.Rollback()
	}

	err = tx.Commit()

	log.Println(event)
	log.Println(err)
}
