package db

import (
	"fmt"
	"time"
)

type DbConfig struct {
	Driver          string
	Host            string
	Port            string
	User            string
	Password        string
	DatabaseName    string
	SSLMode         string
	Timezone        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func (d *DbConfig) GetDSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s", d.Driver, d.User, d.Password, d.Host, d.Port, d.DatabaseName)
}

// Builder Object for DbConfig
type DbConfigBuilder struct {
	driver          string
	host            string
	port            string
	user            string
	password        string
	databaseName    string
	sslMode         string
	timeZone        string
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration
}

// Constructor for DbConfigBuilder
func NewDbConfigBuilder() *DbConfigBuilder {
	o := new(DbConfigBuilder)
	return o
}

// Build Method which creates DbConfig
func (b *DbConfigBuilder) Build() *DbConfig {
	o := new(DbConfig)
	o.Driver = b.driver
	o.Host = b.host
	o.Port = b.port
	o.User = b.user
	o.Password = b.password
	o.DatabaseName = b.databaseName
	o.MaxIdleConns = b.maxIdleConns
	o.ConnMaxLifetime = b.connMaxLifetime
	o.MaxOpenConns = b.maxOpenConns
	o.SSLMode = b.sslMode
	o.Timezone = b.timeZone
	return o
}

// Setter method for the field driver of type string in the object DbConfigBuilder
func (d *DbConfigBuilder) SetDriver(driver string) {
	d.driver = driver
}

// Setter method for the field host of type string in the object DbConfigBuilder
func (d *DbConfigBuilder) SetHost(host string) {
	d.host = host
}

// Setter method for the field port of type string in the object DbConfigBuilder
func (d *DbConfigBuilder) SetPort(port string) {
	d.port = port
}

// Setter method for the field user of type string in the object DbConfigBuilder
func (d *DbConfigBuilder) SetUser(user string) {
	d.user = user
}

// Setter method for the field password of type string in the object DbConfigBuilder
func (d *DbConfigBuilder) SetPassword(password string) {
	d.password = password
}

// Setter method for the field databaseName of type string in the object DbConfigBuilder
func (d *DbConfigBuilder) SetDatabaseName(databaseName string) {
	d.databaseName = databaseName
}

// Setter method for the field sslMode of type string in the object DbConfigBuilder
func (d *DbConfigBuilder) SetSslMode(sslMode string) {
	d.sslMode = sslMode
}

// Setter method for the field timeZone of type string in the object DbConfigBuilder
func (d *DbConfigBuilder) SetTimeZone(timeZone string) {
	d.timeZone = timeZone
}

// Setter method for the field maxIdleConss of type int in the object DbConfigBuilder
func (d *DbConfigBuilder) SetMaxIdleConss(maxIdleConns int) {
	d.maxIdleConns = maxIdleConns
}

// Setter method for the field maxOpenConns of type int in the object DbConfigBuilder
func (d *DbConfigBuilder) SetMaxOpenConns(maxOpenConns int) {
	d.maxOpenConns = maxOpenConns
}

// Setter method for the field connMaxLifetime of type time.Duration in the object DbConfigBuilder
func (d *DbConfigBuilder) SetConnMaxLifetime(connMaxLifetime time.Duration) {
	d.connMaxLifetime = connMaxLifetime
}
