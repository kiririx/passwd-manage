package conf

import (
	"database/sql"
	"fmt"
	"github.com/kiririx/krutils/logx"
	"github.com/kiririx/passwd-manage/env"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Sqlx *gorm.DB

func init() {
	initDataBase()
}

func initDataBase() *sql.DB {
	var err error
	db := os.Getenv(env.DB)
	logx.INFO("current db is: " + db)
	var dial gorm.Dialector
	if db == "" || db == "sqlite" {
		dial = sqlite.Open("./data/sensitive.db")
	} else if db == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Local&timeout=5s&charset=utf8mb4&collation=utf8mb4_unicode_ci&interpolateParams=true&parseTime=true&loc=Local",
			os.Getenv("db_username"), os.Getenv("db_password"), os.Getenv("db_host"), os.Getenv("db_port"), os.Getenv("db_database"))
		dial = mysql.Open(dsn)
	} else {
		panic("db not validate")
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	Sqlx, err = gorm.Open(dial, &gorm.Config{
		SkipDefaultTransaction: false, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 设置为true时，表名为复数形式 User的表名应该是user
			TablePrefix:   "t_",  // 表名前缀 User的表名应该是t_user
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 设置成为逻辑外键(在物理数据库上没有外键，仅体现在代码上)
		Logger:                                   newLogger,
	})
	if err != nil {
		log.Fatalln("数据库连接错误:", err)
	}
	pool, err := Sqlx.DB()
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(10)
	pool.SetConnMaxLifetime(time.Minute)
	return pool
}
