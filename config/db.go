package config

const (
	Mysqldb = "root:hellodemo749895452@tcp(127.0.0.1:3306)/ranking?charset=utf8mb4&parseTime=True&loc=Local"
)

// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
