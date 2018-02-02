package main

import "time"

//владельцы девайсов
type UsersStruct struct {
	Id    int
	Name  string
	Email string
}

//девайс
type DevicesStruct struct {
	Id     int
	Name   string
	Userid int
}

//метрика
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

//сообщение о метриках
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
