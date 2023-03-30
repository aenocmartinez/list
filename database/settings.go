package database

import "os"

type settings struct {
}

func (s *settings) DriverDB() string {
	return os.Getenv("DB_DRIVER")
}

func (s *settings) PortDB() string {
	return os.Getenv("DB_PORT")
}

func (s *settings) UserDB() string {
	return os.Getenv("DB_USER")
}

func (s *settings) PassDB() string {
	return os.Getenv("DB_PASS")
}

func (s *settings) NameDB() string {
	return os.Getenv("DB_NAME")
}

func (s *settings) HostDB() string {
	return os.Getenv("DB_HOST")
}
