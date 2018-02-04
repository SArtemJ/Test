package main

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"math/rand"
	"strconv"
	"time"
	"log"

)

var DB *sql.DB
var Myc MyConfig

// подключаемся к БД по данным из конфигурации toml
func init() {

	if _, err := toml.DecodeFile("myconf.toml", &Myc); err != nil {
		fmt.Println(err)
		return
	}

	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Myc.DBuser, Myc.DBpassword, Myc.DBname)
	DB, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	//defer DB.Close() //базу не закрываем она нам еще понадобится



}

//получить все записи из таблицы устройств в канал
func GetAllDevicesFromDB() [] DevicesStruct {

	//out := make(chan DevicesStruct, 10000)
	var dSlice []DevicesStruct
	rows, err := DB.Query("SELECT * FROM devices")
	if err != nil {
		//panic(err)
	}
	defer rows.Close()

	for rows.Next() {
			var newDevice DevicesStruct
			err := rows.Scan(&newDevice.Id, &newDevice.Name, &newDevice.Userid)
			if err != nil {
				//panic(err)
			}
			log.Println(newDevice)
			dSlice = append(dSlice, newDevice)
		}

	return dSlice
}

func CreateMetric(d []DevicesStruct)  {
	time.Sleep(5 * time.Second)
		var newMetric DevicesMetricStruct
		for _, v := range d {

			newMetric.Id = TableIDs("device_metrics")
			//log.Println(newMetric.Id)
			newMetric.Deviceid = v.Id
			for i := 0; i < len(newMetric.Metric); i++ {
				newMetric.Metric[i] = rand.Intn(50)
			}
			newMetric.LocalTime = time.Now().AddDate(0, 0, -1)
			newMetric.ServerTime = time.Now()

			log.Println(newMetric)
			var stringQ= "INSERT INTO device_metrics (Id, device_Id, metric_1, metric_2, metric_3, metric_4, metric_5, local_time, server_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
			_, err := DB.Exec(stringQ,
				newMetric.Id,
				newMetric.Deviceid,
				newMetric.Metric[0],
				newMetric.Metric[1],
				newMetric.Metric[2],
				newMetric.Metric[3],
				newMetric.Metric[4],
				newMetric.LocalTime,
				newMetric.ServerTime)
			if err != nil {
				fmt.Println(err.Error())
				//return
			}

			checkMetrics(newMetric)
		}

}

func checkMetrics(r DevicesMetricStruct) {

	var newAlert DeviceAlertStruct
			for i := 0; i < len(r.Metric); i++ {
				if r.Metric[i] == Myc.BadMetricParam {
					newAlert.Id = TableIDs("device_alerts")
					newAlert.Deviceid = r.Deviceid
					newAlert.Message = "Bad metric param on device " + strconv.Itoa(r.Deviceid)
					//getValues(newAlert.Id)
					setValues(newAlert.Deviceid, newAlert.Message)
					getValues(newAlert.Deviceid)
					_, err := DB.Exec("INSERT INTO device_alerts (id, device_id, message) VALUES ($1, $2, $3)", newAlert.Id, newAlert.Deviceid, newAlert.Message)
					if err != nil {
						fmt.Println(err.Error())
						//return
					}
				}
			}
}

func TableIDs(nameT string) (lastID int) {
	stringQ := "SELECT COUNT(ID) FROM " + nameT + ";"
	rows, err := DB.Query(stringQ)
	if err != nil {
		panic(err)
	}
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
