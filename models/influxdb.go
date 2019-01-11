package models

import (
	"encoding/json"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

//influxdb数据库连接
func ConnInfluxdb() client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://127.0.0.1:8086",
		Username: "admin",
		Password: "admin",
	})
	if err != nil {
		log.Fatal("influxdb数据库连接错误： ", err)
	}
	return cli
}

//influxdb查询
func QueryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "test",
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return nil, response.Error()
		}
		res = response.Results
	} else {
		return nil, err
	}
	return res, nil
}

//influxdb写入
func WritesPoints(cli client.Client, field map[string]string) error {
	//t := time.Now()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s",
	})
	if err != nil {
		return err
	}
	for key, value := range field {
		m := make(map[string]interface{})
		json.Unmarshal([]byte(value), &m)
		for value1, value2 := range m {
			tags := map[string]string{"key": key, "value1": value1}
			fields := map[string]interface{}{
				"value": value2,
			}
			pt, err := client.NewPoint(
				"map",
				tags,
				fields,
				time.Now(),
			)
			if err != nil {
				return err
			}
			bp.AddPoint(pt)
		}
	}
	if err := cli.Write(bp); err != nil {
		return err
	}
	//elapsed := time.Since(t)
	//fmt.Println(elapsed)
	return nil
}
