package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "postgres"
	password = "root"
)

type Author struct {
	ID int
	Name string
}
type Book struct {
	ID int
	Name string
	Page int
	Author_Name string
	Author_ID int
}
type Result struct {
	Author string
	Books string
	Pages string
}

func main() {
	con := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Print("Successfully connected!\n\n")

	/*info, err := db.Query("select * from author;")
	if err != nil {
		panic(err)
	}
	for info.Next() {
		a := Author{}
		err = info.Scan(&a.ID, &a.Name)
		if err != nil {
			panic(err)
		}
		fmt.Println(a.ID, a.Name)
	}

	_, err = db.Exec("update author set name = 'F.Benjamin' where id = 7 and name like 'F%'")
	if err != nil {
		panic(err)
	} else {
		fmt.Print("\nTable updated!\n\n")
	}

	info, err = db.Query("select * from author;")
	if err != nil {
		panic(err)
	}
	for info.Next() {
		a := Author{}
		err = info.Scan(&a.ID, &a.Name)
		if err != nil {
			panic(err)
		}
		fmt.Println(a.ID, a.Name)
	}*/


	query := "select a.name author, array_agg(b.name) books, array_agg(b.page) pages from author a left join book b on a.id = b.author_id group by a.name"
	info, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for info.Next() {
		r := Result{}
		err = info.Scan(&r.Author, &r.Books, &r.Pages)
		if err != nil {
			panic(err)
		}
		fmt.Println(r.Author, r.Books, r.Pages)
	}
}