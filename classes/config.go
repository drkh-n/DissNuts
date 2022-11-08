package main

import(
  "fmt"
)

type DataBaseCred struct{
  ip string
  dbName string
  dbUserName string
  dbPassword string
  dbPort int32
}

func(std *DataBaseCred) defaultValues(){
  ip := "127.0.0.1"
  dbName := "swe"
  dbUserName := "root"
  dbPassword := "root"
  dbPort := 3306
}
