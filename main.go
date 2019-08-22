package main
import (
	"database/sql"
	   "fmt"
	   _ "github.com/mattn/go-sqlite3"
	   "net/http" 
		"log"
		"strconv"
		"strings"
		"encoding/json"
		"io/ioutil"
		
)

var db *sql.DB = init_db()

type user struct {
	Id int
	First_name string
    Last_name string
	Email string
	Gender string
	Age int
}

func home(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprint(w, "Welcome to the HomePage!") 
}

func showAll(w http.ResponseWriter, r *http.Request) {
	
	rows, err := db.Query("select id,first_name,last_name,email,gender,age from users")
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

	addUser := []user{}  
    for rows.Next() {
		rows.Scan(&id,&first_name,&last_name,&email,&gender,&age)
		
		data := user {
			Id : id,
			First_name : first_name,
			Last_name : last_name,
			Email : email,
			Gender : gender,
			Age : age,
		} 
		addUser = append(addUser,data)
		// fmt.Println(addUser)
		// fmt.Fprint(w,"\n"+strconv.Itoa(id)+" "+first_name+" "+last_name)
	}
	
	// fmt.Println(addUser)
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
	json.NewEncoder(w).Encode(addUser)
}

func showByid(w http.ResponseWriter, r *http.Request) {
	find_id := strings.Split(r.URL.Path[1:],"/")[1]

	rows, err := db.Query("select first_name,last_name,email,gender,age from users where id = " + find_id)
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
		rows.Scan(&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
		fmt.Println(first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func showByfirst_name(w http.ResponseWriter, r *http.Request) {
	find_first_name := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(find_first_name)
	rows, err := db.Query("select first_name,last_name,email,gender,age  from users where first_name = " + "'" + find_first_name + "'")
	var (
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
		rows.Scan(&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func showBylast_name(w http.ResponseWriter, r *http.Request) {
	find_last_name := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(find_last_name)
	rows, err := db.Query("select first_name,last_name,email,gender,age  from users where last_name = " + "'" + find_last_name + "'")
	var (
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
		rows.Scan(&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func showByemail(w http.ResponseWriter, r *http.Request) {
	find_email := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(find_email)
	rows, err := db.Query("select first_name,last_name,email,gender,age  from users where email = " + "'" + find_email + "'")
	var (
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
		rows.Scan(&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func showBygender(w http.ResponseWriter, r *http.Request) {
	find_gender := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(find_gender)
	rows, err := db.Query("select first_name,last_name,email,gender,age  from users where gender = " + "'" + find_gender + "'")
	var (
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
		rows.Scan(&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
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

	rows, err := db.Query("select first_name,last_name,email,gender,age  from users where age >= " + find_age_min + " and age <= " + find_age_max)
	
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
		rows.Scan(&first_name,&last_name,&email,&gender,&age)
		fmt.Fprint(w,first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age)+"\n")
		fmt.Println(first_name+" "+last_name+" "+email+" "+gender+" "+strconv.Itoa(age))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	rBody , _ := ioutil.ReadAll(r.Body)
	
	str := string(rBody)
	data := strings.Split(str,"&")
	first_name := strings.Split(data[0],"=")[1]
	last_name := strings.Split(data[1],"=")[1]
	email := strings.Split(data[2],"=")[1]
	gender := strings.Split(data[3],"=")[1]
	age,_ := strconv.Atoi(strings.Split(data[4],"=")[1])
	fmt.Println(first_name,last_name,email,gender,age)
	
	result, _ := db.Exec(
		"INSERT INTO users ('first_name','last_name','email','gender','age') VALUES (?,?,?,?,?)",
		first_name,last_name,email,gender,age)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	// w.Write([]byte("Received a POST request\n"))
}

func update(w http.ResponseWriter, r *http.Request) {
	rBody , _ := ioutil.ReadAll(r.Body)
	
	str := string(rBody)
	data := strings.Split(str,"&")
	first_name := strings.Split(data[0],"=")[1]
	last_name := strings.Split(data[1],"=")[1]
	email := strings.Split(data[2],"=")[1]
	gender := strings.Split(data[3],"=")[1]
	age,_ := strconv.Atoi(strings.Split(data[4],"=")[1])
	// fmt.Println(first_name,last_name,email,gender,age)
	
	result, _ := db.Exec(
		"INSERT INTO users ('first_name','last_name','email','gender','age') VALUES (?,?,?,?,?)",
		first_name,last_name,email,gender,age)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	// w.Write([]byte("Received a POST request\n"))
}

func delete(w http.ResponseWriter, r *http.Request) {
	del_id := strings.Split(r.URL.Path[1:],"/")[1]
	fmt.Println(del_id)
	fmt.Println("DEL")
	result, _ := db.Exec(
		"DELETE FROM users WHERE id=?",
		del_id)
	fmt.Println(result.RowsAffected())
	// w.Write([]byte("Received a POST request\n"))
}

func handleRequest() { 
	http.HandleFunc("/", home) 
	http.HandleFunc("/showAll", showAll)
	http.HandleFunc("/showByid/", showByid)
	http.HandleFunc("/showByfirst_name/", showByfirst_name)
	http.HandleFunc("/showBylast_name/", showBylast_name)
	http.HandleFunc("/showByemail/", showByemail)
	http.HandleFunc("/showBygender/", showBygender)
	http.HandleFunc("/showByage/", showByage)
	http.HandleFunc("/create", create)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete/", delete)
    http.ListenAndServe(":8000", nil) 
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

