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

	//получим все устройства в канал и передадим его на обработку функции CreateMetric
	deviseS := make(chan DevicesStruct)
	deviseS = GetAllDevicesFromDB()

	metrics := make(chan DevicesMetricStruct)
	metrics = createMetric(deviseS)

	for n := range metrics {
		log.Println(n)
	}

}

//Все утсройства из БД полученные,
//параллельно каждые 5 минут генерируют произвольную метрику
func createMetric(d chan DevicesStruct) chan DevicesMetricStruct {
	metrics := make(chan DevicesMetricStruct)
	var mID int

	go func() {
		for v := range d {
			go func() {
				var newMetric= DevicesMetricStruct{}
				//каждая новая метрика каждого устройства записыввается в таблицу
				//но чтобы не произошло нарушения целостности ключа
				//то ищем последний id в таблице метрик и записываем метрику на id+1
				//так как это функция то возможно ее тоже нужно бросить в рутину
				//но тогда рутину создания метрики и вычисления id-шника придейтся сливать в один канал
				mID = TableIDs("device_metrics")

				log.Println("Metrics id new  - " + strconv.Itoa(mID))

				//метрики рандомно генерирурются
				//это массив из 5 значений
				var flagMetrics = 5
				newMetric.Id = mID
				newMetric.Deviceid = v.Id
				for i := 0; i < flagMetrics; i++ {
					newMetric.Metric[i] = rand.Intn(1000)
				}
				newMetric.LocalTime = time.Now().AddDate(0, 0, -1)
				newMetric.ServerTime = time.Now()

				log.Println(newMetric)

				//каждая новая мтерика записывается в БД
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

				//если значение метрики равно значению из файа myconf.toml 755
				//то это ошибка метрики и нужно создать сообщение и записать в БД
				checkMetric(newMetric)

				metrics <- newMetric
				time.Sleep(1 * time.Second)
			}()
		}
	}()

	return metrics
}

func checkMetric(dm DevicesMetricStruct) {
	//так же дял каждого сообщения об ошике вычисляем ID+1
	var aID = TableIDs("device_alerts")
	log.Println("Alerts id " + strconv.Itoa(aID))
	var newAlert = DeviceAlertStruct{}

	//проверяем все значения всех 5 метрик
	//если одно ошибочно то пишес сообщени в таблицу
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
