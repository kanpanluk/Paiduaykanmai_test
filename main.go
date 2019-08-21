package main
import (
	"database/sql"
	   "fmt"
	   _ "github.com/mattn/go-sqlite3"
	   "net/http" 
		"log"
		"strconv"
		"strings"
)

var db *sql.DB = init_db()

func home(w http.ResponseWriter, r *http.Request) { // (1)
    fmt.Fprint(w, "Welcome to the HomePage!") // (2)
}

func showAll(w http.ResponseWriter, r *http.Request) {
	
	rows, err := db.Query("select id,first_name,last_name from users where id = 1")
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

func showByid(w http.ResponseWriter, r *http.Request) {
	find_id := strings.Split(r.URL.Path[1:],"/")[1]

	rows, err := db.Query("select * from users where id = " + find_id)
	var id int
	var first_name string
	var last_name string
	var email string
	var gender string
	var age int
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
		rows.Scan(&id,&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
		fmt.Println(strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func showByfirst_name(w http.ResponseWriter, r *http.Request) {
	find_first_name := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(find_first_name)
	rows, err := db.Query("select * from users where first_name = " + "'" + find_first_name + "'")
	var (id int
	first_name string
	last_name string
	email string
	gender string
	age int)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
		rows.Scan(&id,&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func showBylast_name(w http.ResponseWriter, r *http.Request) {
	find_last_name := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(find_last_name)
	rows, err := db.Query("select * from users where last_name = " + "'" + find_last_name + "'")
	var (id int
	first_name string
	last_name string
	email string
	gender string
	age int)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
		rows.Scan(&id,&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func showByemail(w http.ResponseWriter, r *http.Request) {
	find_email := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(find_email)
	rows, err := db.Query("select * from users where email = " + "'" + find_email + "'")
	var (id int
	first_name string
	last_name string
	email string
	gender string
	age int)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
		rows.Scan(&id,&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func showBygender(w http.ResponseWriter, r *http.Request) {
	find_gender := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(find_gender)
	rows, err := db.Query("select * from users where gender = " + "'" + find_gender + "'")
	var (id int
	first_name string
	last_name string
	email string
	gender string
	age int)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
		rows.Scan(&id,&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func showByage(w http.ResponseWriter, r *http.Request) {
	find_age := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(find_age)

	find_age_min := strings.Split(find_age,"-")[:1][0]
	var find_age_max string

	if find_age_max = find_age_min ; len(strings.Split(find_age,"-")[1:]) > 0 {
		find_age_max = strings.Split(find_age,"-")[1:][0]
	}

	rows, err := db.Query("select * from users where age >= " + find_age_min + " and age <= " + find_age_max)
	var id int
	var first_name string
	var last_name string
	var email string
	var gender string
	var age int
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
		rows.Scan(&id,&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(strconv.Itoa(id)+" "+first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func handleRequest() { // (3)
	http.HandleFunc("/", home) // (4)home 
	http.HandleFunc("/showAll", showAll)
	http.HandleFunc("/showByid/", showByid)
	http.HandleFunc("/showByfirst_name/", showByfirst_name)
	http.HandleFunc("/showBylast_name/", showBylast_name)
	http.HandleFunc("/showByemail/", showByemail)
	http.HandleFunc("/showBygender/", showBygender)
	http.HandleFunc("/showByage/", showByage)
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

