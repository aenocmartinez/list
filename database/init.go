package database

import (
	"sync"
)

var lock = &sync.Mutex{}
var instance *ConnectDB
var instanceSetting *settings

var user string = dataSource().UserDB()
var pass string = dataSource().PassDB()
var host string = dataSource().HostDB()
var port string = dataSource().PortDB()
var database string = dataSource().NameDB()

func dataSource() *settings {
	if instanceSetting == nil {
		if instanceSetting == nil {
			instanceSetting = &settings{}
		}
	}
	return instanceSetting
}

func getDatabase() interface{} {
	var source interface{}
	if dataSource().DriverDB() == "mysql" {
		source = NewMySQL()
	}
	return source
}

func Instance() *ConnectDB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &ConnectDB{
				source: getDatabase(),
			}
		}
	}
	return instance
}
