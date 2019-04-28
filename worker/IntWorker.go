package worker

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/global"
	"github.com/Deansquirrel/goMonitorV5/object"
	"github.com/Deansquirrel/goMonitorV5/repository/task"
	"github.com/Deansquirrel/goMonitorV5/repository/taskHis"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolMSSql"
	"reflect"
	"strconv"
	"strings"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

type intWorker struct {
	config *object.IntTaskConfig
}

func (w *intWorker) GetMsg() (string, object.ITaskHis) {
	comm := common{}
	if w.config == nil {
		msg := comm.getMsg("", "config is nil")
		return w.formatMsg(msg), nil
	}
	hisData := object.IntHisData{
		FId:       strings.ToUpper(goToolCommon.Guid()),
		FConfigId: w.config.FId,
		FOprTime:  time.Now(),
	}
	num, err := w.getCheckNum()
	if err != nil {
		errMsg := fmt.Sprintf("get check num error: %s", err.Error())
		msg := comm.getMsg(w.config.FMsgTitle, errMsg)
		hisData.FContent = msg
		return w.formatMsg(msg), &hisData
	}
	hisData.FNum = num
	var msg string
	if num <= w.config.FCheckMin || num >= w.config.FCheckMax {
		msg = comm.getMsg(w.config.FMsgTitle,
			strings.Replace(w.config.FMsgContent, "title", strconv.Itoa(num), -1))
		dMsg := w.getDMsg()
		if dMsg != "" {
			msg = msg + "\n" + dMsg
		}
		if msg == "" {
			msg = dMsg
		} else {
			if dMsg != "" {
				msg = msg + "\n" + "--------------------" + "\n" + dMsg
			}
		}
		msg = w.formatMsg(msg)
		hisData.FContent = msg
		return msg, &hisData
	}
	return "", &hisData
}

func (w *intWorker) SaveData(data object.ITaskHis) error {
	if data == nil {
		return nil
	}
	hisRep, err := taskHis.NewHisRepository(object.Int)
	if err != nil {
		return err
	}
	return hisRep.SetHis(data)
}

func (w *intWorker) DelHisData() error {
	hisRep, err := taskHis.NewHisRepository(object.Int)
	if err != nil {
		return err
	}
	d := time.Duration(1000 * 1000 * 1000 * 60 * 60 * 24 * global.SysConfig.TaskConfig.KeepDays)
	return hisRep.ClearHis(d)
}

func (w *intWorker) CheckAction() error {
	//TODO
	return nil
}

func (w *intWorker) formatMsg(msg string) string {
	if msg != "" {
		msg = goToolCommon.GetDateTimeStr(time.Now()) + "\n" + msg
	}
	return msg
}

//获取待检测值
func (w *intWorker) getCheckNum() (int, error) {
	rows, err := w.getRowsBySQL(w.config.FSearch)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = rows.Close()
	}()
	list := make([]int, 0)
	var num int
	for rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			log.Error(err.Error())
			break
		} else {
			list = append(list, num)
		}
	}
	if err != nil {
		return 0, err
	}
	if len(list) != 1 {
		errMsg := fmt.Sprintf("SQL返回数量异常，exp:1,act:%d", len(list))
		log.Error(errMsg)
		return 0, errors.New(errMsg)
	}
	return list[0], nil
}

//查询数据
func (w *intWorker) getRowsBySQL(sql string) (*sql.Rows, error) {
	conn, err := goToolMSSql.GetConn(w.getDBConfig())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	rows, err := conn.Query(sql)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return rows, nil
}

//获取DB配置
func (w *intWorker) getDBConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: w.config.FServer,
		Port:   w.config.FPort,
		DbName: w.config.FDbName,
		User:   w.config.FDbUser,
		Pwd:    w.config.FDbPwd,
	}
}

func (w *intWorker) getDMsg() string {
	rep, err := task.NewRepository(object.IntD)
	if err != nil {
		log.Error(fmt.Sprintf(err.Error()))
		return err.Error()
	}
	dConfig, err := rep.GetConfig(w.config.FId)
	if err != nil {
		errMsg := fmt.Sprintf("get intd config error,id: %s,error: %s", w.config.FId, err.Error())
		log.Error(errMsg)
		return errMsg
	}
	//无明细配置
	if dConfig == nil {
		return ""
	}
	switch reflect.TypeOf(dConfig).String() {
	case "*object.IntDTaskConfig":
		c, ok := dConfig.(*object.IntDTaskConfig)
		if ok {
			return w.getSingleDMsg(c.FMsgSearch)
		} else {
			return fmt.Sprintf("convert intd config error")
		}
	default:
		return fmt.Sprintf("type error,expected %s,get %s",
			reflect.TypeOf(object.IntDTaskConfig{}).String(),
			reflect.TypeOf(dConfig).String())
	}
}

func (w *intWorker) getSingleDMsg(search string) string {
	if search == "" {
		return ""
	}
	rows, err := w.getRowsBySQL(search)
	if err != nil {
		return fmt.Sprintf("查询明细内容时遇到错误：%s，查询语句为：%s", err.Error(), search)
	}
	defer func() {
		_ = rows.Close()
	}()
	titleList, err := rows.Columns()
	if err != nil {
		return fmt.Sprintf("获取明细内容表头时遇到错误：%s，查询语句为：%s", err.Error(), search)
	}
	counter := len(titleList)
	values := make([]interface{}, counter)
	valuePointers := make([]interface{}, counter)
	for i := 0; i < counter; i++ {
		valuePointers[i] = &values[i]
	}

	var result string
	for rows.Next() {
		err = rows.Scan(valuePointers...)
		if err != nil {
			return fmt.Sprintf("读取明细内容时遇到错误：%s，查询语句为：%s", err.Error(), search)
		}
		if result != "" {
			result = result + "\n" + "--------------------"
		}
		for i := 0; i < counter; i++ {
			if result != "" {
				result = result + "\n"
			}
			var v string
			if values[i] == nil {
				v = "Null"
			} else {
				v = goToolCommon.ConvertToString(values[i])
			}
			result = result + titleList[i] + " - " + v
		}
	}
	if rows.Err() != nil {
		return fmt.Sprintf("读取明细内容时遇到错误：%s，查询语句为：%s", err.Error(), search)
	}
	return result
}
