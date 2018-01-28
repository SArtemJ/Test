package main

import (
	"fmt"
	"time"
	"math/rand"
)


//генерируем метрики
/*
передаем слайс всех девайсов
прозодим в цикле
создаем новою метрику
генерируем случайные номера для метрик от 1...5
записываем в БД - device_metric
проверяем значения метрик, если == значению в файле myconf
создаем запись в БД _ alert
вызываем отправку email владельцу девайса

 */
func StartWrite(dev []DevicesStruct) DevicesMetricStruct {


		var setMetrics= DevicesMetricStruct{}

		for i := 1; i < len(dev)+1; i++ {
			var mID = TableIDs()
			setMetrics.Id = mID
			setMetrics.Deviceid = dev[i].Id
			setMetrics.LocalTime = time.Now().AddDate(0, 0, -1)
			setMetrics.ServerTime = time.Now()
			var flagMetrics = 5
			for i := 0; i < flagMetrics; i++ {
				setMetrics.Metric[i] = rand.Intn(1000)
			}

			_, err := DB.Exec("INSERT INTO device_metrics (Id, device_Id, metric_1, metric_2, metric_3, metric_4, metric_5, local_time, server_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
				setMetrics.Id,
				setMetrics.Deviceid,
				setMetrics.Metric[0],
				setMetrics.Metric[1],
				setMetrics.Metric[2],
				setMetrics.Metric[3],
				setMetrics.Metric[4],
				setMetrics.LocalTime,
				setMetrics.ServerTime)
			if err != nil {
				fmt.Println(err.Error())
				//return
			}

			for i:=0; i<len(setMetrics.Metric); i++ {
				if setMetrics.Metric[i] == Myc.BadMetricParam {
					var alert= DeviceAlertStruct{}
					alert.Id = TableIDsAlerts()
					alert.Deviceid = setMetrics.Deviceid
					alert.Message = "Bad Metric Param"
					_, err := DB.Exec("INSERT INTO device_alerts (id, device_id, message) values ($1, $2, $3)", alert.Id, alert.Deviceid, alert.Message)
					//Myc.Tmail =	GetMailToSend(alert.Deviceid)
					Send(Myc.Fmail, Myc.Fpass, GetMailToSend(alert.Deviceid), "Error metrics on your device")
					//Send()
					if err != nil {
						fmt.Println(err.Error())
						//return
					}
				}
			}

		}

	return setMetrics
}

func main() {

	rand.Seed(time.Now().Unix())

	//получаем всех юзеров БД
	go func() {
		allU <- GetAllUsersFromDB()
	}()

	//получаем все девайсы
	go func() {
		allD <- GetAllDevicesFromDB()
	}()

	allDevSlice := <-allD

	//создаем метрики
	go func() {
			allM <- StartWrite(allDevSlice)

	}()

	select {
		case <- allM:
	default:
		time.Sleep(5*time.Second)
	}


}