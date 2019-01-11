package models

import (
	"fmt"
	"log"
	"time"
)

type MapResults struct {
	MapTime string
	Value   interface{}
}

func AddInfluxdbData() error {
	field, err := GetRedisData()
	if err != nil {
		return err
	}
	conn := ConnInfluxdb()
	err = WritesPoints(conn, field)
	if err != nil {
		return err
	}
	return nil
}

func GetInfluxdbData(t1 string, t2 string, tags map[string]string) ([]*MapResults, error) {
	conn := ConnInfluxdb()
	fmt.Println(tags)
	cmd1 := " "
	for key, value := range tags {
		switch key {
		case "key":
			if value == "" {
				break
			}
			cmd1 = cmd1 + " and " + "\"" + "key" + "\"" + " = " + "'" + value + "'" + " "
		case "value1":
			if value == "" {
				break
			}
			cmd1 = cmd1 + " and " + "\"" + "value1" + "\"" + " = " + "'" + value + "'" + " "
		default:

		}
	}
	cmd := fmt.Sprintf("select * from %s where time >= %s and time < %s"+cmd1+"tz('Asia/Shanghai')",
		"map", "'"+t1+"'", "'"+t2+"'")
	fmt.Println(cmd)
	res, err := QueryDB(conn, cmd)
	if err != nil {
		return nil, err
	}

	if len(res[0].Series) == 0 {
		return nil, nil
	}
	results := make([]*MapResults, 0)
	for _, row := range res[0].Series[0].Values {
		_, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		m := new(MapResults)
		m.MapTime = row[0].(string)
		m.Value = row[2]
		//m.Key = row[1].(string)

		/*x := row[2].(string)
		n := make(map[string]interface{})
		json.Unmarshal([]byte(x), &n)
		m.Value = n*/
		/*
			switch row[2].(type) {
			case json.Number:
				str := string(row[2].(json.Number))
				m.Value, err = strconv.Atoi(str)
				if err != nil {
					log.Fatal(err)
				}
				//fmt.Println("执行了我")
			case string:
				m.Value, err = strconv.Atoi(row[2].(string))
				if err != nil {
					log.Fatal(err)
				}
			case :

			}*/
		results = append(results, m)
	}
	return results, nil
}
