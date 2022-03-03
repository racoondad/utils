package utils

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// CreateDB 创建数据库
func CreateDB(dbType, username, password, host, port, dbName string) error {
	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8mb4&parseTime=True&loc=Local",
			username,
			password,
			host,
			port,
		)
		conn, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("create database %s;", dbName))
		return err
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable TimeZone=Asia/Shanghai",
			host,
			username,
			password,
			port,
		)
		conn, err := sql.Open("postgres", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("create database %s;", dbName))
		return err
	case "sqlserver":
		dsn := fmt.Sprintf("server=%s;port%s;database=master;user id=%s;password=%s", host, port, username, password)
		conn, err := sql.Open("mssql", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("create database %s;", dbName))
		return err
	default:
		return nil
	}
}

// CreateDB 创建数据库
func DropDatabase(dbType, username, password, host, port, dbName string) error {
	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=uutf8mb4&parseTime=True&loc=Local",
			username,
			password,
			host,
			port,
		)
		conn, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("drop database %s;", dbName))
		return err
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable TimeZone=Asia/Shanghai",
			host,
			username,
			password,
			port,
		)
		conn, err := sql.Open("postgres", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("drop database %s;", dbName))
		return err
	case "sqlserver":
		dsn := fmt.Sprintf("server=%s;port%s;database=master;user id=%s;password=%s", host, port, username, password)
		conn, err := sql.Open("mssql", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("drop database %s;", dbName))
		return err
	default:
		return nil
	}
}

// GormSqlServer 初始化SqlServer数据库
func GormSqlServer(username, password, host, dbName, config, tablePrefix, port string, maxIdleConns, maxOpenConns int) (*gorm.DB, error) {
	dsn := "sqlserver://" + username + ":" + password + "@" + host + ":" + port + "?database=" + dbName + config
	if db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	}); err != nil {
		return nil, errors.New("连接数据库失败")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
		return db, nil
	}
}

// GormPostgreSql 初始化PostgreSql数据库
func GormPostgreSql(username, password, host, dbName, config, tablePrefix, port string, maxIdleConns, maxOpenConns int) (*gorm.DB, error) {
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbName + " port=" + port + " " + config
	postgresConfig := postgres.Config{
		DSN:                  dsn,   // DSN data source name
		PreferSimpleProtocol: false, // 禁用隐式 prepared statement
	}
	if db, err := gorm.Open(postgres.New(postgresConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	}); err != nil {
		return nil, errors.New("连接数据库失败")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
		return db, nil
	}
}

// GormMysql 初始化PostgreSql数据库
func GormMysql(username, password, host, dbName, config, tablePrefix, port string, maxIdleConns, maxOpenConns int) (*gorm.DB, error) {

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?" + config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	}); err != nil {

		return nil, errors.New("连接数据库失败")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
		return db, nil
	}
}
