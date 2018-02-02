package main

import "time"

type UsersStruct struct {
	Id    int
	Name  string
	Email string
}

//устройства
type DevicesStruct struct {
	Id     int
	Name   string
	Userid int
}

//метрика устройства
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

//сообщение о плохих метриках
type DeviceAlertStruct struct {
	Id       int
	Deviceid int
	Message  string
}

//файл конфигурации
type MyConfig struct {
	DBuser         string
	DBname         string
	DBpassword     string
	BadMetricParam int `toml:bmp`
	Fmail          string
	Fpass          string
	Tmail          string
}
