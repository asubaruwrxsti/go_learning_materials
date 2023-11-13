package main
import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"os"
)

// declare a global variable to hold the connection pool.
	// this is the db handle that we'll use to execute queries.

// use mysql.Config to collect connection properties
	// this is a struct that implements the DSN method

// use sql.Open initialize the db variable
	// passing the return value of cfg.FormatDSN()

// check for errors
	// sql.Open doesn't actually create a connection
		// sql.Open only validates the arguments
		// sql.Open might fail

// call log.Fatal if there is an error

// db.Ping() actually creates a connection

// check for errors
	// db.Ping() might fail

var db *sql.DB

type Car struct {
	ID int
	Name string
}

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "carrentone",
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

    fmt.Printf("> Get by Car name \n")
    // Get car by name audi1 using getCars(name)
    cars, err := getCars("audi1")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("cars: %+v\n", cars)

    fmt.Printf("> Get by Car id \n")
    // Get car by id 31 using getCar(id)
    c, err := getCarById(31)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Car found: %v\n", c)

    fmt.Printf("> Inserting Car \n")
    carId, err := insertCar(Car {
        Name: "TestGOLANG",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID of the added car: %v\n", carId)
}

func getCars(name string) ([]Car, error) {
    var cars []Car

    rows, err := db.Query("SELECT id, name FROM cars WHERE name = ?", name)
    if err != nil {
        return nil, err
    }
    defer rows.Close() // any resources it holds will be released when the function exits

    for rows.Next() { // iterate over the rows
        var c Car
        // Scan takes a list of pointers to Go values, 
        // where the column values will be written. 
        // Here, you pass pointers to fields in the alb variable, 
        // created using the & operator. 
        
        // Scan writes through the pointers to update the struct fields.

        if err := rows.Scan(&c.ID, &c.Name); err != nil {
            return nil, err
        }
        cars = append(cars, c)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return cars, nil
}

// use db.QueryRow() to execute select statement
    // returns an sql.Row
    // sql.Row doesnt return an error, returns a query error

// use row.Scan() to copy column values into struct fields

// check error from scan
    // the special error sql.ErrNoRows indicates that the qeury returned no results

func getCarById(id int) (Car, error) {
    var c Car
    row := db.QueryRow("SELECT id, name FROM cars WHERE id = ?", id)
    if err := row.Scan(&c.ID, &c.Name); err != nil {
        if err == sql.ErrNoRows {
            return c, fmt.Errorf("Cars by id %d: no such car", id)
        }
        return c, fmt.Errorf("carsById %d: %v", id, err)
    }
    return c, nil
}

func insertCar(c Car) (int64, error) {
    result, err := db.Exec("INSERT INTO cars (name) VALUES (?)", c.Name)
    if err != nil {
        return 0, fmt.Errorf("insertCar: %v", err)
    }

    // Retrieve the ID of the inserted database row using Result.LastInsertId.
    id, err := result.LastInsertId()

    if err != nil {
        return 0, fmt.Errorf("insertCar: %v", err)
    }
    return id, nil
}