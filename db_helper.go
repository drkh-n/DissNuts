package main

import(
  "fmt"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)

func connectDB() *sql.DB{
  dbCredentials := new(DataBaseCred)
  dbCredentials.defaultValues()

  fmt.Println(ip, dbname, username, password, port)
  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbCredentials.dbUserName,dbCredentials.dbPassword, dbCredentials.ip, dbCredentials.dbPort, dbCredentials.dbName))
  if err != nil{
    panic(err)
  }else{
    fmt.Println("Подключено к БД")
    return db
    //fmt.Println(reflect.TypeOf(db))
  }
}

func insertPatient(newUser UserObject) bool{
  db := connectDB()
  insert, err := db.Query(fmt.Sprintf("INSERT IGNORE INTO `users` (`uin`, `name`, `surname`, `middlename`, `blood_group`, `phone`, `email`, `address`, `emergency_contact`, `marital_status`, `birthdate`, `date`) VALUES('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');",newUser.uin, newUser.name, newUser.surname, newUser.middlename, newUser.bloodtype, newUser.phone, newUser.email, newUser.emergencyContact, newUser.address, newUser.maritalStatus, newUser.birthdate, newUser.regDate))
  if err != nil{
    panic(err)
    return false
  }else{
    fmt.Println("Insert query successful")
  }
  return true
  defer insert.Close()
  defer db.Close()
}
