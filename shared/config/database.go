package config

import (
	"fmt"
)

type Database struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Hostname  string `json:"hostname"`
	Parameters string `json:"parameters"`
	Port      int `json:"port"`
	Protocol string `json:"protocol"`

	GormSingularTable bool `json:"gorm_singular_table"`
	GormLogMode bool `json:"gorm_log_mode"`
}

func (d *Database) Dsn() string {
	return  fmt.Sprintf("%s:%s@%s(%s:%d)/%s%s",
		d.Username, d.Password, d.Protocol, d.Hostname, d.Port, d.Name, d.Parameters)
}