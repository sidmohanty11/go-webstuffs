package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	//connect to a database -> ref => https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=testDb user=postgres password=") //add password if any
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}

	defer conn.Close()

	fmt.Println("Connected to DB!")

	//test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping db!")
	}

	fmt.Println("Pinged DB!")

	//get rows from table
	err = getAllRows(conn)

	if err != nil {
		log.Fatalln(err.Error())
	}

	//insert a row
	query := `
		insert into users (first_name, last_name) values ($1, $2)
	`
	_, err = conn.Exec(query, "Chris", "Brown") //$1 -> Chris, $2 -> Brown
	if err != nil {
		log.Fatal(err.Error())
	}

	//get rows again, from the table
	err = getAllRows(conn)

	if err != nil {
		log.Fatalln(err.Error())
	}

	//update a row
	stmt := `
		update users set first_name = $1 where first_name = $2
	`
	_, err = conn.Exec(stmt, "Drake", "Chris") //$1 -> Drake, $2 -> Chris

	if err != nil {
		log.Fatal(err.Error())
	}

	//get rows again, from the table
	err = getAllRows(conn)

	if err != nil {
		log.Fatalln(err.Error())
	}

	//get one row by id when you add id as a col.-----------

	// query = `select id, first_name, last_name from users where id = $1`
	// var fName, lName string
	// var id int
	// row := conn.QueryRow(query, 1)
	// err = row.Scan(&id, &fName, &lName)

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// log.Println(id,fName,lName)

	//--------------------------------------------------------

	//delete a row
	query = `delete from users where first_name = $1`
	_, err = conn.Exec(query, "Drake")

	if err != nil {
		log.Fatal(err.Error())
	}

	//get rows again, from the table
	err = getAllRows(conn)

	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select first_name, last_name from users")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer rows.Close()

	var fName, lName string

	for rows.Next() {
		err := rows.Scan(&fName, &lName)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		fmt.Println("record is => ", fName, lName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("error scanning rows ", err)
	}

	fmt.Println("====================================")

	return nil
}
