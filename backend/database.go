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
	Hour   uint8
	Minute uint8
	Second uint8
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
	if len(timeSlice) < 3 {
		return errors.New("too few time arguments")
	}
	hour, err := strconv.ParseUint(timeSlice[0], 10, 8)
	if err != nil {
		return fmt.Errorf("error parsing hour: %w", err)
	}
	minute, err := strconv.ParseUint(timeSlice[1], 10, 8)
	if err != nil {
		return fmt.Errorf("error parsing minute: %w", err)
	}
	second, err := strconv.ParseUint(timeSlice[2], 10, 8)
	if err != nil {
		return fmt.Errorf("error parsing seconds: %w", err)
	}

	t.Hour = uint8(hour)
	t.Minute = uint8(minute)
	t.Second = uint8(second)

	return nil
}

type Event struct {
	ID        int64     `db:"id"`
	Created   time.Time `db:"created"`
	Lecture   Lecture   `db:"-"`
	LectureID int64     `db:"lecture"`
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

func (e *Event) Create(tx *sqlx.Tx) (int64, error) {
	e.Created = time.Now()
	res, err := tx.NamedExec("INSERT INTO events (created, lecture, date, startTime, endTime, room, confirmed, comment) VALUES (:created, :lecture, :date, :startTime, :endTime, :room, :confirmed, :comment)", e)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	e.ID = id
	return id, nil
}

func (e *Event) Update(tx *sqlx.Tx) error {

	var oldCreation time.Time
	err := tx.Get(&oldCreation, "SELECT created FROM events WHERE id=?", e.ID)

	if err != nil {
		return err
	}

	if !oldCreation.Equal(e.Created) {
		return errors.New("inconsistent data update - creation time mismatch")
	}

	e.Created = time.Now()
	_, err = tx.NamedExec("UPDATE events SET created=:created, lecture=:lecture, date=:date, startTime=:startTime, endTime=:endTime, room=:room, confirmed=:confirmed, comment=:comment WHERE id=:id", e)

	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete(tx *sqlx.Tx) error {
	res, err := tx.Exec("DELETE FROM events WHERE id=?", e.ID)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if count < 1 {
			return errors.New("no rows affected")
		}
	}
	return nil
}

type Events []Event

func GetEvents(tx *sqlx.Tx, query string, args ...interface{}) (Events, error) {
	var events Events
	err := tx.Select(&events, query, args...)

	if err != nil {
		return nil, err
	}

	for index, event := range events {
		err = event.Lecture.Get(tx, "SELECT * FROM lectures WHERE id=?", event.LectureID)
		if err != nil {
			return nil, err
		}
		events[index] = event
	}

	return events, nil
}

type Course struct {
	ID      int64     `db:"id"`
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

func (c *Course) Create(tx *sqlx.Tx) (int64, error) {
	c.Created = time.Now()
	res, err := tx.NamedExec("INSERT INTO courses (created, name, year, faculty) VALUES (:created, :name, :year, :faculty)", c)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	c.ID = id
	return id, nil
}

func (c *Course) Update(tx *sqlx.Tx) error {

	var oldCreation time.Time
	err := tx.Get(&oldCreation, "SELECT created FROM courses WHERE id=?", c.ID)

	if err != nil {
		return err
	}

	if !oldCreation.Equal(c.Created) {
		return errors.New("inconsistent data update - creation time mismatch")
	}

	c.Created = time.Now()
	_, err = tx.NamedExec("UPDATE lectures SET created=:created, name=:name, year=:year, faculty=:faculty WHERE id=:id", c)

	if err != nil {
		return err
	}

	return nil
}

func (c *Course) Delete(tx *sqlx.Tx) error {
	res, err := tx.Exec("DELETE FROM courses WHERE id=?", c.ID)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if count < 1 {
			return errors.New("no rows affected")
		}
	}
	return nil
}

type Semester struct {
	ID      int64     `db:"id"`
	Created time.Time `db:"created"`
	Start   time.Time `db:"start"`
	End     time.Time `db:"end"`
	Name    string    `db:"name"`
}

func (s *Semester) Get(tx *sqlx.Tx, query string, args ...interface{}) error {
	return tx.Get(s, query, args...)
}

func (s *Semester) Create(tx *sqlx.Tx) (int64, error) {
	s.Created = time.Now()
	res, err := tx.NamedExec("INSERT INTO semesters (created, start, end, name) VALUES (:created, :start, :end, :name)", s)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	s.ID = id
	return id, nil
}

func (s *Semester) Update(tx *sqlx.Tx) error {

	var oldCreation time.Time
	err := tx.Get(&oldCreation, "SELECT created FROM semesters WHERE id=?", s.ID)

	if err != nil {
		return err
	}

	if !oldCreation.Equal(s.Created) {
		return errors.New("inconsistent data update - creation time mismatch")
	}

	s.Created = time.Now()
	_, err = tx.NamedExec("UPDATE lectures SET created=:created, start=:start, end=:end, name=:name WHERE id=:id", s)

	if err != nil {
		return err
	}

	return nil
}

func (s *Semester) Delete(tx *sqlx.Tx) error {
	res, err := tx.Exec("DELETE FROM semesters WHERE id=?", s.ID)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if count < 1 {
			return errors.New("no rows affected")
		}
	}
	return nil
}

type Lecture struct {
	ID      int64     `db:"id"`
	Created time.Time `db:"created"`
	Name    string    `db:"name"`

	LecturerID int64    `db:"lecturer"`
	Lecturer   Lecturer `db:"-"`

	SemesterID int64    `db:"semester"`
	Semester   Semester `db:"-"`

	CourseID int64  `db:"course"`
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

func (l *Lecture) Create(tx *sqlx.Tx) (int64, error) {
	l.Created = time.Now()
	res, err := tx.NamedExec("INSERT INTO lectures (created, name, lecturer, semester, course) VALUES (:created, :name, :lecturer, :semester, :course)", l)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	l.ID = id
	return id, nil
}

func (l *Lecture) Update(tx *sqlx.Tx) error {

	var oldCreation time.Time
	err := tx.Get(&oldCreation, "SELECT created FROM lecturers WHERE id=?", l.ID)

	if err != nil {
		return err
	}

	if !oldCreation.Equal(l.Created) {
		return errors.New("inconsistent data update - creation time mismatch")
	}

	l.Created = time.Now()
	_, err = tx.NamedExec("UPDATE lectures SET created=:created, name=:name, lecturer=:lecturer, semester=:semester, course=:course WHERE id=:id", l)

	if err != nil {
		return err
	}

	return nil
}

func (l *Lecture) Delete(tx *sqlx.Tx) error {
	res, err := tx.Exec("DELETE FROM lectures WHERE id=?", l.ID)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if count < 1 {
			return errors.New("no rows affected")
		}
	}
	return nil
}

type Lecturer struct {
	ID        int64         `db:"id"`
	Created   time.Time     `db:"created"`
	Surname   string        `db:"surname"`
	GivenName string        `db:"givenName"`
	Email     string        `db:"email"`
	SAMLUID   sql.NullInt64 `db:"samlUID"`
}

func (l *Lecturer) Get(tx *sqlx.Tx, query string, args ...interface{}) error {
	return tx.Get(l, query, args...)
}

func (l *Lecturer) Create(tx *sqlx.Tx) (int64, error) {
	l.Created = time.Now()
	res, err := tx.NamedExec("INSERT INTO lecturers (surname, givenName, email, samlUID) VALUES (:surname, :givenName, :email, :samlUID)", l)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	l.ID = id
	return id, nil
}

func (l *Lecturer) Update(tx *sqlx.Tx) error {

	var oldCreation time.Time
	err := tx.Get(&oldCreation, "SELECT created FROM lecturers WHERE id=?", l.ID)

	if err != nil {
		return err
	}

	if !oldCreation.Equal(l.Created) {
		return errors.New("inconsistent data update - creation time mismatch")
	}

	l.Created = time.Now()
	_, err = tx.NamedExec("UPDATE lecturers SET created=:created, surname=:surname, givenName=:givenName, email=:email, samlUID=:samlUID WHERE id=:id", l)

	if err != nil {
		return err
	}

	return nil
}

func (l *Lecturer) Delete(tx *sqlx.Tx) error {
	res, err := tx.Exec("DELETE FROM lecturers WHERE id=?", l.ID)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if count < 1 {
			return errors.New("no rows affected")
		}
	}
	return nil
}

type UserRights struct {
	ID        int64     `db:"id"`
	Created   time.Time `db:"created"`
	SAMLUID   int64     `db:"samlUID"`
	Email     string    `db:"email"`
	IsAdmin   bool      `db:"isAdmin"`
	AddCourse bool      `db:"addCourse"`
	EditEvent bool      `db:"editEvent"`
	AddEvent  bool      `db:"addEvent"`
}

func (ur *UserRights) Get(tx *sqlx.Tx, query string, args ...interface{}) error {
	return tx.Get(ur, query, args...)
}

func (ur *UserRights) Create(tx *sqlx.Tx) (int64, error) {
	ur.Created = time.Now()
	res, err := tx.NamedExec("INSERT INTO userRights (samlUID, email, isAdmin, addCourse, editEvent, addEvent) VALUES (:samlUID, :email, :isAdmin, :addCourse, :editEvent, :addEvent)", ur)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	ur.ID = id
	return id, nil
}

func (ur *UserRights) Update(tx *sqlx.Tx) error {

	var oldCreation time.Time
	err := tx.Get(&oldCreation, "SELECT created FROM userRights WHERE id=?", ur.ID)

	if err != nil {
		return err
	}

	if !oldCreation.Equal(ur.Created) {
		return errors.New("inconsistent data update - creation time mismatch")
	}

	ur.Created = time.Now()
	_, err = tx.NamedExec("UPDATE userRights SET created=:created, samlUID=:samlUID, email=:email, isAdmin=:isAdmin, addCourse=:addCourse, editEvent=:editEvent, addEvent=:addEvent WHERE id=:id", ur)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRights) Delete(tx *sqlx.Tx) error {
	res, err := tx.Exec("DELETE FROM userRights WHERE id=?", ur.ID)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if count < 1 {
			return errors.New("no rows affected")
		}
	}
	return nil
}

type UnavailableTime struct {
	ID        int64     `db:"id"`
	Created   time.Time `db:"created"`
	StartTime time.Time `db:"startTime"`
	EndTime   time.Time `db:"endTime"`

	Lecturer   Lecturer `db:"-"`
	LecturerID int64    `db:"lecturer"`
}

func (ut *UnavailableTime) Get(tx *sqlx.Tx, query string, args ...interface{}) error {

	err := tx.Get(ut, query, args...)
	if err != nil {
		return err
	}

	return ut.Lecturer.Get(tx, "SELECT * FROM unavailableTimes WHERE id=?")
}

func (ut *UnavailableTime) Create(tx *sqlx.Tx) (int64, error) {
	ut.Created = time.Now()
	res, err := tx.NamedExec("INSERT INTO unavailableTimes (created, startTime, endTime, lecturer) VALUES (:created, :startTime, :endTime, :lecturer)", ut)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	ut.ID = id
	return id, nil
}

func (ut *UnavailableTime) Update(tx *sqlx.Tx) error {

	var oldCreation time.Time
	err := tx.Get(&oldCreation, "SELECT created FROM unavailableTimes WHERE id=?", ut.ID)

	if err != nil {
		return err
	}

	if !oldCreation.Equal(ut.Created) {
		return errors.New("inconsistent data update - creation time mismatch")
	}

	ut.Created = time.Now()
	_, err = tx.NamedExec("UPDATE userRights SET created=:created, startTime=:startTime, endTime=:endTime, lecturer=:lecturer) WHERE id=:id", ut)

	if err != nil {
		return err
	}

	return nil
}

func (ut *UnavailableTime) Delete(tx *sqlx.Tx) error {
	res, err := tx.Exec("DELETE FROM unavailableTimes WHERE id=?", ut.ID)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if count < 1 {
			return errors.New("no rows affected")
		}
	}
	return nil
}

type PreferredWeekday struct {
	ID         int64     `db:"id"`
	Created    time.Time `db:"created"`
	Lecturer   Lecturer  `db:"-"`
	LecturerID int64     `db:"lecturer"`
	Monday     bool      `db:"monday"`
	Tuesday    bool      `db:"tuesday"`
	Wednesday  bool      `db:"wednesday"`
	Thursday   bool      `db:"thursday"`
	Friday     bool      `db:"friday"`
	Saturday   bool      `db:"saturday"`
}

func (pw *PreferredWeekday) Get(tx *sqlx.Tx, query string, args ...interface{}) error {
	err := tx.Get(pw, query, args...)
	if err != nil {
		return err
	}

	return pw.Lecturer.Get(tx, "SELECT * FROM preferredWeekdays WHERE id=?", pw.LecturerID)
}

func (pw *PreferredWeekday) Create(tx *sqlx.Tx) (int64, error) {
	pw.Created = time.Now()
	res, err := tx.NamedExec("INSERT INTO preferredWeekdays (created, lecturer, monday, tuesday, wednesday, thursday, friday, saturday) VALUES (:created, :lecturer, :monday, :tuesday, :wednesday, :thursday, :friday, :saturday)", pw)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	pw.ID = id
	return id, nil
}

func (pw *PreferredWeekday) Update(tx *sqlx.Tx) error {

	var oldCreation time.Time
	err := tx.Get(&oldCreation, "SELECT created FROM preferredWeekdays WHERE id=?", pw.ID)

	if err != nil {
		return err
	}

	if !oldCreation.Equal(pw.Created) {
		return errors.New("inconsistent data update - creation time mismatch")
	}

	pw.Created = time.Now()
	_, err = tx.NamedExec("UPDATE preferredWeekdays SET created=:created, lecturer=:lecturer, monday=:monday, tuesday=:tuesday, wednesday=:wednesday, thursday=:thursday, friday=:friday, saturday=:saturday) WHERE id=:id", pw)

	if err != nil {
		return err
	}

	return nil
}

func (pw *PreferredWeekday) Delete(tx *sqlx.Tx) error {
	res, err := tx.Exec("DELETE FROM preferredWeekdays WHERE id=?", pw.ID)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if count < 1 {
			return errors.New("no rows affected")
		}
	}
	return nil
}

type PreferredDay struct {
	ID         int64     `db:"id"`
	Created    time.Time `db:"created"`
	Lecturer   Lecturer  `db:"-"`
	LecturerID int64     `db:"lecturer"`
	Day        time.Time `db:"day"`
}

func (pd *PreferredDay) Get(tx *sqlx.Tx, query string, args ...interface{}) error {

	err := tx.Get(pd, query, args...)
	if err != nil {
		return err
	}

	return pd.Lecturer.Get(tx, "SELECT * FROM lecturers WHERE id=?", pd.LecturerID)
}

func (pd *PreferredDay) Create(tx *sqlx.Tx) (int64, error) {
	pd.Created = time.Now()
	res, err := tx.NamedExec("INSERT INTO preferredDays (created, lecturer, day) VALUES (:created, :lecturer, :day)", pd)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	pd.ID = id
	return id, nil
}

func (pd *PreferredDay) Update(tx *sqlx.Tx) error {

	var oldCreation time.Time
	err := tx.Get(&oldCreation, "SELECT created FROM preferredDays WHERE id=?", pd.ID)

	if err != nil {
		return err
	}

	if !oldCreation.Equal(pd.Created) {
		return errors.New("inconsistent data update - creation time mismatch")
	}

	pd.Created = time.Now()
	_, err = tx.NamedExec("UPDATE preferredDays SET created=:created, lecturer=:lecturer, day=:day) WHERE id=:id", pd)

	if err != nil {
		return err
	}

	return nil
}

func (pd *PreferredDay) Delete(tx *sqlx.Tx) error {
	res, err := tx.Exec("DELETE FROM preferredDays WHERE id=?", pd.ID)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if count < 1 {
			return errors.New("no rows affected")
		}
	}
	return nil
}

func openDB() *sqlx.DB {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("mysql", "mysql:mysql@/lectureplan?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db

}
