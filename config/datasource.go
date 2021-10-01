package config

type Datasource struct {
	DriverName string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
	Charset    string
	Loc        string
}
