package main

import "time"

type UsersStruct struct {
	Id    int
	Name  string
	Email string
}

type DevicesStruct struct {
	Id     int
	Name   string
	Userid int
}

type DevicesMetricStruct struct {
	Id       int
	Deviceid int
	//metric1    int
	//metric2    int
	//metric3    int
	//metric4    int
	//metric5    int
	Metric     [5]int
	LocalTime  time.Time
	ServerTime time.Time
}

type DeviceAlertStruct struct {
	Id       int
	Deviceid int
	Message  string
}

type MyConfig struct {
	DBuser string
	DBname string
	DBpassword string
	BadMetricParam int `toml:bmp`
	Fmail string
	Fpass string
	Tmail string
}

func AllError(err error) {
	if err != nil {
		panic(err)
	}
}

