package main

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var Myc MyConfig

//  парсим файл конфигурации в структуру
// подключаемся к БД по данным из конфигурации
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

	if err = DB.Ping(); err != nil {
		panic(err)
	}
}

//получить все записи из таблицы устройств в канал
func GetAllDevicesFromDB() chan DevicesStruct {

	dev := make(chan DevicesStruct)
	rows, err := DB.Query("SELECT * FROM devices;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		newDevice := DevicesStruct{}
		err := rows.Scan(&newDevice.Id, &newDevice.Name, &newDevice.Userid)
		if err != nil {
			panic(err)
		}
		dev <- newDevice
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	return dev

}

//проверка что запись новой строки в таблицу  будет уникальна
//инкремент
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
