package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rest_api/config"
	"rest_api/model"

	"github.com/gorilla/mux"
)

// AllEmployee = Select Employee API
func AllEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	var response model.Response
	var arrEmployee []model.Employee

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, city FROM employee")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Name, &employee.City)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrEmployee = append(arrEmployee, employee)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrEmployee

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// InsertEmployee = Insert Employee API
func InsertEmployee(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var employee model.Employee
	db := config.Connect()
	defer db.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &employee)

	result, err := db.Exec("INSERT INTO employee(name, city) VALUES(?, ?)", employee.Name, employee.City)

	if err != nil {
		log.Print(err)
		return
	}
	employee.Id, err = result.LastInsertId()

	response.Status = 200
	response.Message = "Insert data successfully"
	response.Data = []model.Employee{employee}
	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

// 	// id := mux.Vars(r)["id"]
// 	var response model.Response
// 	var employee model.Employee
// 	db := config.Connect()
// 	defer db.Close()
// 	body, err := io.ReadAll(r.Body)

// 	if err != nil {
// 		panic(err)
// 	}

// 	json.Unmarshal(body, &employee)

// 	result, err := db.Exec("update employee set name =?,city=? where id =?", employee.Name, employee.City)

// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	employee.Id, err = result.LastInsertId()

// 	response.Status = 200
// 	response.Message = "Insert data successfully"
// 	response.Data = []model.Employee{employee}
// 	fmt.Print("Insert data to database")

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(response)
// }

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.Connect()
	defer db.Close()
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE employee SET name = ?,city = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName := keyVal["name"]
	newCity := keyVal["city"]
	_, err = stmt.Exec(newName, newCity, params["id"])
	if err != nil {
		panic(err.Error())
	}
	// fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])
	response.Status = 200
	response.Message = "Success"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response model.Response
	params := mux.Vars(r)

	db := config.Connect()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM employee WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	response.Status = 200
	response.Message = "Success"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
