package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)


func main() {

	rand.Seed(time.Now().Unix())

	deviseS := make(chan DevicesStruct)
	deviseS = GetAllDevicesFromDB()

	metrics := make(chan DevicesMetricStruct)
	metrics = createMetric(deviseS)

	for n := range metrics {
		log.Println(n)
	}

}

func createMetric(d chan DevicesStruct) chan DevicesMetricStruct {
	metrics := make(chan DevicesMetricStruct)
	var mID int

	go func() {
		for v := range d {
			go func() {
				var newMetric= DevicesMetricStruct{}
				//чтобы записать новое значенеи в таблицу device_metrics получаем последний номер ID
				mID = TableIDs("device_metrics")
				log.Println("Metrics id new  - " + strconv.Itoa(mID))
				var flagMetrics = 5
				newMetric.Id = mID
				newMetric.Deviceid = v.Id
				for i := 0; i < flagMetrics; i++ {
					newMetric.Metric[i] = rand.Intn(1000)
				}
				newMetric.LocalTime = time.Now().AddDate(0, 0, -1)
				newMetric.ServerTime = time.Now()

				log.Println(newMetric)
				_, err := DB.Exec("INSERT INTO device_metrics (Id, device_Id, metric_1, metric_2, metric_3, metric_4, metric_5, local_time, server_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);",
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

				//проверка есть ли ошибки в сообщениях если есть создаем уведомление
				checkMetric(newMetric)

				metrics <- newMetric
				time.Sleep(1 * time.Second)
			}()
		}
	}()

	return metrics
}

func checkMetric(dm DevicesMetricStruct) {
	var aID = TableIDs("device_alerts")
	log.Println("Alerts id " + strconv.Itoa(aID))
	var newAlert = DeviceAlertStruct{}

	for i := 0; i < len(dm.Metric); i++ {
		if dm.Metric[i] == Myc.BadMetricParam {
			newAlert.Id = aID
			newAlert.Deviceid = dm.Deviceid
			newAlert.Message = "Bad metric param on device " + strconv.Itoa(dm.Deviceid)
			log.Println(newAlert)
			_, err := DB.Exec("INSERT INTO device_alerts (id, device_id, message) dms ($1, $2, $3);", newAlert.Id, newAlert.Deviceid, newAlert.Message)
			if err != nil {
				fmt.Println(err.Error())
				//return
			}
		}
	}
}
