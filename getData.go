package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// VolumeThroughput export
type VolumeThroughput struct {
	Table                  string `json:"table"`
	PxBuild                string `json:"pxbuild"`
	JobName                string `json:"jobname"`
	JenkinBuild            int    `json:"jenkinbuild"`
	Date                   string `json:"date"`
	BlockSize              string `json:"blocksize"`
	IoType                 string `json:"iotype"`
	R1LocalWrite           int    `json:"r1localwrite"`
	R1LocalRead            int    `json:"r1localread"`
	R1RemoteWrite          int    `json:"r1remotewrite"`
	R1RemoteRead           int    `json:"r1remoteread"`
	R2LocalWrite           int    `json:"r2localwrite"`
	R2LocalRead            int    `json:"r2localread"`
	R2RemoteWrite          int    `json:"r2remotewrite"`
	R2RemoteRead           int    `json:"r2remoteread"`
	R3LocalWrite           int    `json:"r3localwrite"`
	R3LocalRead            int    `json:"r3localread"`
	R2LocalWrite1ReplDown  int    `json:"r2localwrite1repldown"`
	R2LocalRead1ReplDown   int    `json:"r2localread1repldown"`
	R2RemoteWrite1ReplDown int    `json:"r2remotewrite1repldown"`
	R2RemoteRead1ReplDown  int    `json:"r2remoteread1repldown"`
	R3LocalWrite1ReplDown  int    `json:"r3localwrite1repldown"`
	R3LocalRead1ReplDown   int    `json:"r3localread1repldown"`
	R1Aggr3LocalWrite      int    `json:"r1aggr3localwrite"`
	R1Aggr3LocalRead       int    `json:"r1aggr3localread"`
	R1Aggr2LocalWrite      int    `json:"r1aggr2localwrite"`
	R1Aggr2LocalRead       int    `json:"r1aggr2localread"`
}

// Perf456 export
type Perf456 struct {
	Table        string  `json:"table"`
	Date         string  `json:"date"`
	BlockSize    string  `json:"blocksize"`
	IoType       string  `json:"iotype"`
	R1LocalWrite float32 `json:"r1localwrite"`
	R1LocalRead  float32 `json:"r1localread"`
}

// Resync export
type Resync struct {
	Table          string  `json:"table"`
	PxBuild        string  `json:"pxbuild"`
	JenkinBuild    int     `json:"jenkinbuild"`
	Date           string  `json:"date"`
	Datasize       int     `json:"datasize"`
	Throughput     int     `json:"throughput"`
	ElapsedTimeSec float32 `json:"elapsedtimesec"`
	ElapsedTime    string  `json:"elapsedtime"`
}

// Sysbench export
type Sysbench struct {
	Table              string  `json:"table"`
	PxBuild            string  `json:"pxbuild"`
	JenkinBuild        int     `json:"jenkinbuild"`
	Date               string  `json:"date"`
	Test               string  `json:"test"`
	TableSize          int     `json:"tablesize"`
	NumOfThreads       int     `json:"numofthreads"`
	TransactionsPerSec []uint8 `json:"transactionspersec"`
	Avg                int     `json:"avg"`
	Percentile         int     `json:"percentile"`
}

func getVolumeThroughput(tableName string) []VolumeThroughput {
	entries := make([]VolumeThroughput, 0, 0)
	var date, pxBuild, jobName, blockSize, ioType string
	var jenkinBuild, r1LocalWrite, r1LocalRead, r1RemoteWrite, r1RemoteRead, r2LocalWrite, r2LocalRead, r2RemoteWrite, r2RemoteRead, r3LocalWrite, r3LocalRead, r2LocalWrite1ReplDown, r2LocalRead1ReplDown, r2RemoteWrite1ReplDown, r2RemoteRead1ReplDown, r3LocalWrite1ReplDown, r3LocalRead1ReplDown, r1Aggr3LocalWrite, r1Aggr3LocalRead, r1Aggr2LocalWrite, r1Aggr2LocalRead int

	fmt.Println("Connecting to MySQL Database - VolumeThroughput")

	db, err := sql.Open("mysql", "root:password@tcp(70.0.0.41)/portworx")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var query string
	query = fmt.Sprintf("select * from %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		en := new(VolumeThroughput)
		err := rows.Scan(&pxBuild, &jobName, &jenkinBuild, &date, &blockSize, &ioType, &r1LocalWrite, &r1LocalRead, &r1RemoteWrite, &r1RemoteRead, &r2LocalWrite, &r2LocalRead, &r2RemoteWrite, &r2RemoteRead, &r3LocalWrite, &r3LocalRead, &r2LocalWrite1ReplDown, &r2LocalRead1ReplDown, &r2RemoteWrite1ReplDown, &r2RemoteRead1ReplDown, &r3LocalWrite1ReplDown, &r3LocalRead1ReplDown, &r1Aggr3LocalWrite, &r1Aggr3LocalRead, &r1Aggr2LocalWrite, &r1Aggr2LocalRead)
		if err != nil {
			log.Fatal(err)
		}
		en.Table = tableName
		en.PxBuild = pxBuild
		en.JobName = jobName
		en.JenkinBuild = jenkinBuild
		en.Date = date
		en.BlockSize = blockSize
		en.IoType = ioType
		en.R1LocalWrite = r1LocalWrite
		en.R1LocalRead = r1LocalRead
		en.R1RemoteWrite = r1RemoteWrite
		en.R1RemoteRead = r1RemoteRead
		en.R2LocalWrite = r2LocalWrite
		en.R2LocalRead = r2LocalRead
		en.R2RemoteWrite = r2RemoteWrite
		en.R2RemoteRead = r2RemoteRead
		en.R3LocalWrite = r3LocalWrite
		en.R3LocalRead = r3LocalRead
		en.R2LocalWrite1ReplDown = r2LocalWrite1ReplDown
		en.R2LocalRead1ReplDown = r2LocalRead1ReplDown
		en.R2RemoteWrite1ReplDown = r2RemoteWrite1ReplDown
		en.R2RemoteRead1ReplDown = r2RemoteRead1ReplDown
		en.R3LocalWrite1ReplDown = r3LocalWrite1ReplDown
		en.R3LocalRead1ReplDown = r2LocalRead1ReplDown
		en.R1Aggr3LocalWrite = r1Aggr2LocalWrite
		en.R1Aggr3LocalRead = r1Aggr3LocalRead
		en.R1Aggr2LocalWrite = r1Aggr2LocalWrite
		en.R1Aggr2LocalRead = r1Aggr2LocalRead

		entries = append(entries, *en)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
		entries[i], entries[j] = entries[j], entries[i]
	}
	return entries
}

func getPerf456(tableName string) []Perf456 {
	entries := make([]Perf456, 0, 0)
	var date, blockSize, ioType string
	var r1LocalWrite, r1LocalRead float32

	fmt.Println("Connecting to MySQL Database - Perf456")

	db, err := sql.Open("mysql", "root:password@tcp(70.0.0.41)/portworx")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var query string
	query = fmt.Sprintf("select * from %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		en := new(Perf456)
		err := rows.Scan(&date, &blockSize, &ioType, &r1LocalWrite, &r1LocalRead)
		if err != nil {
			log.Fatal(err)
		}
		en.Table = tableName
		en.Date = date
		en.BlockSize = blockSize
		en.IoType = ioType
		en.R1LocalWrite = r1LocalWrite
		en.R1LocalRead = r1LocalRead

		entries = append(entries, *en)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
		entries[i], entries[j] = entries[j], entries[i]
	}
	return entries
}

func getResync(tableName string) []Resync {
	entries := make([]Resync, 0, 0)
	var date, pxBuild, elapsedTime string
	var jenkinBuild, datasize, throughput int
	var elapsedTimeSec float32

	fmt.Println("Connecting to MySQL Database - Resync")

	db, err := sql.Open("mysql", "root:password@tcp(70.0.0.41)/portworx")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var query string
	query = fmt.Sprintf("select * from %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		en := new(Resync)
		err := rows.Scan(&pxBuild, &jenkinBuild, &date, &datasize, &throughput, &elapsedTimeSec, &elapsedTime)
		if err != nil {
			log.Fatal(err)
		}
		en.Table = tableName
		en.Date = date
		en.PxBuild = pxBuild
		en.Throughput = throughput
		en.Datasize = datasize
		en.ElapsedTime = elapsedTime
		en.ElapsedTimeSec = elapsedTimeSec
		en.JenkinBuild = jenkinBuild
		entries = append(entries, *en)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
		entries[i], entries[j] = entries[j], entries[i]
	}
	return entries
}

func getSysbench(tableName string) []Sysbench {
	entries := make([]Sysbench, 0, 0)
	var date, pxBuild, test string
	var jenkinBuild, tableSize, numOfThreads, avg, percentile int
	var transactionPerSec []uint8

	fmt.Println("Connecting to MySQL Database - Sysbench")

	db, err := sql.Open("mysql", "root:password@tcp(70.0.0.41)/portworx")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var query string
	query = fmt.Sprintf("select * from %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		en := new(Sysbench)
		err := rows.Scan(&pxBuild, &jenkinBuild, &date, &test, &tableSize, &numOfThreads, &transactionPerSec, &avg, &percentile)
		if err != nil {
			log.Fatal(err)
		}
		en.Table = tableName
		en.Avg = avg
		en.Date = date
		en.JenkinBuild = jenkinBuild
		en.NumOfThreads = numOfThreads
		en.Percentile = percentile
		en.PxBuild = pxBuild
		en.TableSize = tableSize
		en.Test = test
		en.TransactionsPerSec = transactionPerSec

		entries = append(entries, *en)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
		entries[i], entries[j] = entries[j], entries[i]
	}
	return entries
}
