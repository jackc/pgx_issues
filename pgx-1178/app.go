package main

import (
	"time"
)

type _App struct {
	Name string `json:"name"`
	Cfg  _Cfg
}

type _Cfg struct {
	CheckInterval int64  `json:"CheckInterval"`
	MaxPercent    uint64 `json:"MaxPercentt"`
	MaxRetr       uint   `json:"MaxRetr"`
}

func (c *_Cfg) Init() {
	*c = _Cfg{
		CheckInterval: 15,
		MaxPercent:    80,
		MaxRetr:       4,
	}
}

func (app *_App) AddCron(f func(), v time.Duration) {

	for {
		time.Sleep(v * time.Second)
		f()
	}
}
