package main

import(
  "fmt"
  "reflect"
  "time"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "net/http"
  "log"
)

//uin, name, surname, middlename, bloodtype, phone, email, emergencyContact, address, maritalStatus, birthdate, regDate
//newUser := UserObject{uin:"010101500200",name:"Test", surname:"Testov", middlename:"Testovich", bloodtype: "O(-)", phone:"+77771112233",email:"email@example.com",emergencyContact:"+77472223344",address:"Kabanbay 53",maritalStatus:"single",birthdate:"2000-01-01",regDate:time.Now().Format("2017-09-07")}

func insertPatientRequest() bool{
  newUser := UserObject{
    name: r.URL.Query().Get("name"),
    surname: r.URL.Query().Get("surname"),
    middlename: r.URL.Query().Get("middlename"),
    bloodtype: r.URL.Query().Get("bloodtype"),
    phone: r.URL.Query().Get("phone"),
    email: r.URL.Query().Get("email"),
    emergencyContact: r.URL.Query().Get("emergency_contact"),
    address: r.URL.Query().Get("address"),
    maritalStatus: r.URL.Query().Get("marital_status"),
    birthdate: r.URL.Query().Get("birthdate"),
    regDate: time.Now().Format("2017-09-07"),
  }
  if newUser.isValidUser(){
    return insertPatient(newUser)
  }
  return false
}

func (s string) isEmpty() bool{
  return len(s) == 0
}
