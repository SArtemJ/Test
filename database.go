package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/BurntSushi/toml"
	"fmt"
	"log"
)

var DB *sql.DB
var Myc MyConfig
var allU = make(chan []UsersStruct)
var allD = make(chan []DevicesStruct)
var allM = make(chan DevicesMetricStruct)

func init() {


	if _, err := toml.DecodeFile("myconf.toml", &Myc); err != nil {
		fmt.Println(err)
		return
	}

	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Myc.DBuser, Myc.DBpassword, Myc.DBname)
	DB, err = sql.Open("postgres", dbinfo)
	AllError(err)

	if err = DB.Ping(); err != nil {
		panic(err)
	}
}


func GetAllUsersFromDB() []UsersStruct {
	var usersSlice []UsersStruct
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
	defer rows.Close()

	for rows.Next() {
		foundUser := UsersStruct{}
		err := rows.Scan(&foundUser.Id, &foundUser.Name, &foundUser.Email)
		AllError(err)
		usersSlice = append(usersSlice, UsersStruct{Id: foundUser.Id, Name: foundUser.Name, Email: foundUser.Email})
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	log.Println(len(usersSlice))
	//log.Println(usersSlice)
	return usersSlice
}

func GetAllDevicesFromDB() []DevicesStruct {

	var devicesSlice []DevicesStruct
	rows, err := DB.Query("SELECT * FROM devices")
	AllError(err)
	defer rows.Close()

	for rows.Next() {
		oneDevice := DevicesStruct{}
		err := rows.Scan(&oneDevice.Id, &oneDevice.Name, &oneDevice.Userid)
		if err != nil {
			fmt.Println(err.Error())
			//return
		}
		devicesSlice = append(devicesSlice, DevicesStruct{Id: oneDevice.Id, Name: oneDevice.Name, Userid: oneDevice.Userid})
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
	log.Println(len(devicesSlice))
	//log.Println(devicesSlice)
	return devicesSlice
}


func TableIDs() (lastID int) {
	rows, err := DB.Query("SELECT COUNT(ID) FROM device_metrics")
	AllError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&lastID)
		if err != nil {
			fmt.Println(err.Error())
			//return
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	lastID++
	//log.Printf("LastID  - %T", lastID)
	return lastID
}

func TableIDsAlerts() (lastID int) {
	rows, err := DB.Query("SELECT COUNT(ID) FROM device_alerts")
	AllError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&lastID)
		if err != nil {
			fmt.Println(err.Error())
			//return
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	lastID++
	//log.Printf("LastID  - %T", lastID)
	return lastID
}

func GetMailToSend(p int) (email string) {
	rows, err := DB.Query("select u.email from devices d left join users u on u.id = d.user_id where d.id = $1", p)
	AllError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&email)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Println(email)
	return email

}