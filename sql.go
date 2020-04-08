/*
 * Copyright (c) 2020. Aberic - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gnomon

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"time"
)

// SQL sql 连接对象
type SQLCommon struct {
	DB     *gorm.DB // 数据库任务入口
	DBUrl  string   // dbURL 数据库 URL
	DBUser string   // dbUser 数据库用户名
	DBPass string   // dbPass 数据库用户密码
	DBName string   // dbName 数据库名称
	// LogMode set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
	LogModeEnable   bool
	MaxIdleConnects int
	MaxOpenConnects int
	scheduled       *time.Timer   // 定时任务
	stop            chan struct{} // 释放当前角色chan
}

func (s *SQLCommon) DisConnect() {
	if nil != s.scheduled {
		s.stop <- struct{}{}
	}
}

// Connect 链接数据库服务
//
// dbURL 数据库 URL
//
// dbUser 数据库用户名
//
// dbPass 数据库用户密码
//
// dbName 数据库名称
//
// logModeEnable set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
//
// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
//
// SetMaxOpenConns sets the maximum number of open connections to the database.
func (s *SQLCommon) Connect(dbURL, dbUser, dbPass, dbName string, logModeEnable bool, maxIdleConns, maxOpenConns int) error {
	if nil == s.DB {
		if String().IsEmpty(dbURL) || String().IsEmpty(dbUser) || String().IsEmpty(dbPass) || String().IsEmpty(dbName) {
			return errors.New("db connect params can not be nil")
		}
		s.DBUrl = dbURL
		s.DBUser = dbUser
		s.DBPass = dbPass
		s.DBName = dbName
		s.LogModeEnable = logModeEnable
		s.MaxIdleConnects = maxIdleConns
		s.MaxOpenConnects = maxOpenConns
		s.scheduled = time.NewTimer(time.Millisecond * time.Duration(10))
		s.stop = make(chan struct{}, 1)
		Log().Info("init DB Manager")
		dbValue := strings.Join([]string{s.DBUser, ":", s.DBPass, "@tcp(", s.DBUrl, ")/", s.DBName,
			"?charset=utf8&parseTime=True&loc=Local"}, "")
		Log().Debug("dbValue = " + dbValue)
		var err error
		s.DB, err = gorm.Open("mysql", dbValue)
		if err != nil {
			Log().Error("failed to connect database, err = " + err.Error())
			return err
		}
		s.DB.LogMode(logModeEnable)
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		s.DB.DB().SetMaxIdleConns(maxIdleConns)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		s.DB.DB().SetMaxOpenConns(maxOpenConns)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		s.DB.DB().SetConnMaxLifetime(time.Hour)
		go s.dbKeepAlive(s.DB)
	}
	return nil
}

func (s *SQLCommon) reConnect() error {
	return s.Connect(s.DBUrl, s.DBUser, s.DBPass, s.DBName, s.LogModeEnable, s.MaxIdleConnects, s.MaxOpenConnects)
}

// Exec 执行自定义 SQL
func (s *SQLCommon) Exec(f func(s *SQLCommon)) error {
	if nil == s.DB {
		if err := s.reConnect(); nil == err {
			f(s)
		} else {
			return err
		}
	}
	f(s)
	return nil
}

// ExecSQL 执行自定义 SQL 语句，该方法是对 func Exec(f func(db *gorm.DB)) error 的实现
//
// dest 期望通过该过程赋值的对象
//
// sql 即将执行的 SQL 语句，可以包含 "?" 来做通配符
//
// values 上述 SQL 语句中 "?" 通配符所表达的值
//
// eg：在 db_user 表中根据用户编号和年龄查询用户基本信息，如下所示：
//
// ExecSQL(&user, "select id,name,age from db_user where id=? and age=?", 1, 18)
func (s *SQLCommon) ExecSQL(dest interface{}, sql string, values ...interface{}) {
	s.DB.Raw(Format(sql), values).Scan(dest)
}

// Format SQL 格式化
func Format(elem ...string) string {
	return strings.Join(elem, " ")
}

func (s *SQLCommon) dbKeepAlive(db *gorm.DB) {
	s.scheduled.Reset(time.Second * time.Duration(10))
	for {
		select {
		case <-s.scheduled.C:
			err := db.DB().Ping()
			if nil != err {
				_ = s.Exec(func(sql *SQLCommon) {})
			}
			s.scheduled.Reset(time.Second * time.Duration(10))
		case <-s.stop:
			s.scheduled = nil
			return
		}
	}
}
