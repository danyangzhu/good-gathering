package service

import (
	"fmt"
	"log"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

func WriteInfluxdb(c client.Client, table string, tags map[string]string, fields map[string]interface{}) {
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database: "good_gathering",
	})

	pt, err := client.NewPoint(table, tags, fields, time.Now())
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(pt)

	err = c.Write(bp)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
