package main

import (
	"log"
	"time"
)

func main() {

	var app _App

	app.Name = "Memory-Guard"
	app.Cfg.Init()

	db := _DB{
		User: "services",
		Pass: "services",
		Addr: "dbfs:5432",
		Name: "services",
	}

	for err := db.Connect(); err != nil; {

		log.Printf("Retrying connection to DB.. %s", err.Error())
		time.Sleep(5000 * time.Millisecond)
	}

	log.Printf("Succesfully connected to DB Server '%s'", db.Addr)
	defer db.Close()

	log.Printf("appCFG - %v", app.Cfg)

	for err := db.LoadCfg(&app, false); err != nil; {
		log.Printf("Retrying loading config from DB.. %s", err.Error())
		time.Sleep(5000 * time.Millisecond)
	}

	log.Printf("Succesfully loaded config from DB..")

	go app.AddCron(func() { db.LoadCfg(&app, true) }, 10)

	select {}
}
