package database

var connection string

func init() { // dieksekuti saat secara otomatis
	connection = "PostgrSQL"
}

func GetConnection() string {
	return connection
}
