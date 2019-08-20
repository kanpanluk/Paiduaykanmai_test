package main
import (
	"database/sql"
	   "fmt"
	   _ "github.com/mattn/go-sqlite3"
	   "net/http" 
		"log"
		"strconv"
)

var db *sql.DB = init_db()

func home(w http.ResponseWriter, r *http.Request) { // (1)
    fmt.Fprint(w, "Welcome to the HomePage!") // (2)
}

func showAll(w http.ResponseWriter, r *http.Request) {
	
	rows, err := db.Query("select id,first_name,last_name from users")
	var id int
	var first_name string
	var last_name string
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
		rows.Scan(&id,&first_name,&last_name)
		fmt.Fprint(w,"\n"+strconv.Itoa(id)+" "+first_name+" "+last_name)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func handleRequest() { // (3)
	http.HandleFunc("/", home) // (4)home 
	http.HandleFunc("/showAll", showAll)
    http.ListenAndServe(":8080", nil) // (5)
}

func init_db() *sql.DB {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db
}


func main() {
	fmt.Println("Create File Main.go Success")	
	// fmt.Printf("%T\n", db)
  	handleRequest()
}

