package models 

import (
    "strconv"
    "database/sql"
    _"github.com/mattn/go-sqlite3")

var DB *sql.DB 

func ConnectDatabase() error {

    db, err := sql.Open("sqlite3", "trips.db" ) 

    if err != nil {
        return err} 

    DB = db
    return nil}
        
type Event struct {
    Eid  int "json: 'id'"
    Etitle  string "json: 'title'"
    Elocation  string "json: 'location'"
    Edescription  string "json: 'description'" 
    Edate  string "json: 'date'" 
    Etime   string "json: 'time'"
    Ecost  int "json: 'cost'"
    MaxPeople *string "json: 'max_people' " }

func GetEvents(count int) ([]Event, error){
 
    rows, err := DB.Query("SELECT * from events LIMIT" + strconv.Itoa(count))


    if err != nil { 
        return nil, err}

    defer rows.Close() 

    events := make([]Event, 0) 

    for rows.Next() {
        singleEvent := Event{}
    err = rows.Scan(&singleEvent.Eid, &singleEvent.Etitle,  &singleEvent.Elocation, &singleEvent.Edescription, &singleEvent.Edate, &singleEvent.Etime, &singleEvent.Ecost, 

    &singleEvent.MaxPeople)

    if err != nil {
			return nil, err
		}
   
        events = append(events, singleEvent)
}

err = rows.Err()

	if err != nil {
		return nil, err
	}

	return events, err
}
//
// func AddEvent(newEvent Event) (bool, error){
//     tx, err := DB.Begin()
//     if err!= nil{
//         return false, err}
//
//         stmt, err := tx.Prepare("INSERT INTO events (id, title, location, description, date, time, cost, max_people) VALUES (?. ?, ?, ?, ?, ?, ?, ?)")
//
//         if err != nil{
//             return false, err}
//
//         defer stmt.Close()
//
//     _, err = stmt.Exec(&newEvent.Eid, &newEvent.Etitle, newEvent.Elcoation, &newEvent.Edescription, &newEvent.Edate, E&newEvent.time, &newEvent.Ecost, &newEvent.MaxPeople)        
//
//     if err != nil{
//         return false, err}
//
//     tx.Commit()
//
//     return true nil
// }

