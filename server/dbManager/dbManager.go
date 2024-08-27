package dbManager

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func parseDateTime(originalString, parseFormat string) string {
	parsedTime, err := time.Parse(time.RFC3339, originalString)
	if err != nil {
		fmt.Println("Error parsing date-time:", err)
		return ""
	}
	fmt.Println("Parsed Time:", parsedTime)
	result := parsedTime.Format(parseFormat)
	fmt.Println("Date only:", result)
	return result

}

// Connection represents a hypothetical connection structure
type Connection struct {
	// connection fields (e.g., address, port, etc.)
}

// global variable to  hold the single instance of Connection
// var instance *Connection
var db *sql.DB

var err error

// to ensure thread-safe access to the connection initialization
var once sync.Once

func DbConnection() (*sql.DB, error) {
	once.Do(func() {
		fmt.Println("Initializing connection...")
		connStr := "user=brangapp dbname=postgres sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		err = db.Ping()
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
			db = nil
		} else {
			fmt.Println("Database connection initialized successfully")
		}
	})
	return db, err
}

func GetConnection() (*sql.DB, error) {
	// var x *sql.DB
	// var y error
	var d, e = DbConnection()
	fmt.Println("fetching instance GetConnection", d, e)
	if d == nil {
		return nil, fmt.Errorf("db connection is not initialized")

	}
	return d, e
}

func CheckMeetingRoomAvailability(empid int, bldno, floor uint32, mroom string, booking_date, start_time, end_time string) (bool, error) {
	// Get the database connection
	db1, err := GetConnection()
	if err != nil {
		return false, fmt.Errorf("failed to get database connection: %v", err)
	}
	//defer db1.Close()

	query := `SELECT id, empid, bldno, floor, mroom, booking_date, start_time, end_time 
		  FROM reserveMeetingRoom 
		  WHERE bldno=$1 AND floor=$2 AND mroom=$3 AND booking_date=$4`
	rows, err := db1.Query(query, bldno, floor, mroom, booking_date)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var (
		r_id           int
		r_empno        int
		r_bldno        int
		r_floor        int
		r_mroom        string
		r_booking_date string
		r_start_time   string
		r_end_time     string
	)

	for rows.Next() {

		if err := rows.Scan(&r_id, &r_empno, &r_bldno, &r_floor, &r_mroom, &r_booking_date, &r_start_time, &r_end_time); err != nil {
			return false, fmt.Errorf("failed to scan row: %v", err)
		}
		fmt.Println(r_bldno, r_floor, r_mroom, r_booking_date, "--", r_start_time, "--", r_end_time)

		r_bdate := parseDateTime(r_booking_date, "2006-01-02")
		fmt.Println(r_bdate, "-->check date")

		r_st := parseDateTime(r_start_time, "15:04:05")
		fmt.Println(r_st, "-->check stime")

		r_et := parseDateTime(r_end_time, "15:04:05")
		fmt.Println(r_et, "-->check etime")
		if uint32(r_bldno) == bldno && uint32(r_floor) == floor && r_mroom == mroom && r_bdate == booking_date && r_st == start_time && r_et == end_time {
			return false, fmt.Errorf("matching record present: %v", err)

		}

		if (start_time >= r_start_time && start_time < r_end_time) ||
			(end_time > r_start_time && end_time <= r_end_time) ||
			(start_time <= r_start_time && end_time >= r_end_time) {
			fmt.Println("Meeting room is already booked for the requested time.")
			return false, nil
		}
	}

	// If no conflicts were found, the room is available
	fmt.Println("Meeting room is available for the requested time.")

	query1 := `INSERT INTO reservemeetingroom (empid,bldno,floor,mroom,booking_date,start_time,end_time)
        VALUES ($1, $2, $3,$4, $5,$6,$7) RETURNING empid`

	// Execute the query and get the inserted record's ID
	err = db.QueryRow(query1, empid, bldno, floor, mroom, booking_date, start_time, end_time).Scan(&empid)
	if err != nil {
		log.Fatal("Failed to insert the record:", err)

		return false, fmt.Errorf("failed to insert recrod to table  reservemeetingroom: %v", err)
	}

	fmt.Printf("Record inserted successfully with empid: %d\n", empid)

	return true, nil
}

func CheckCredentials(username, password string) (int, error) {
	db, err = GetConnection()
	var employeeNo int

	// Query to get the password hash for the provided username
	query := "SELECT empid FROM employee WHERE empname=$1 and password=$2"
	err := db.QueryRow(query, username, password).Scan(&employeeNo)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return 0, err
	case nil:
		fmt.Println("Employee No:", employeeNo)

		return employeeNo, nil
	default:
		panic(err)

	}
}
func CloseDb() error {
	return db.Close()
}
