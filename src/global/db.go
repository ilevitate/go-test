package global

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func OpenDB() (err error) {
	if DB != nil {
		return
	}
	section := Config.Section("database")
	addr := section.Key("addr").String()
	user := section.Key("user").String()
	pwd := section.Key("pwd").String()
	name := section.Key("name").String()
	logQuery := section.Key("logQuery").MustBool()
	if addr == "" {
		return errors.New("配置文件[database]节点中的addr参数不正确")
	}
	if user == "" {
		return errors.New("配置文件[database]节点中的user参数不正确")
	}
	if pwd == "" {
		return errors.New("配置文件[database]节点中的pwd参数不正确")
	}
	if name == "" {
		return errors.New("配置文件[database]节点中的name参数不正确")
	}
	//dsn := "root:root@/test?charset=utf8mb4&parseTime=true&loc=Local"
	dsn := user + ":" + pwd + "@tcp(" + addr + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	var newLogger logger.Interface
	if logQuery == true {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,         // Disable color
			},
		)
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	//err = DB.AutoMigrate(&model.User{},&model.Article{})
	//if err != nil {
	//	return err
	//}
	return
}
