package main

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hackattic/utils"
	"io"
	"os"
	"os/exec"

	l "github.com/charmbracelet/log"
	_ "github.com/lib/pq"
)

const (
	chalName = "backup_restore"
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "21BCY10011"
	dbname   = "hackattic"
)

var log *l.Logger

type chalBody struct {
	B64Str string `json:"dump"`
}

type solBody struct {
	Ssns []string `json:"alive_ssns"`
}

type chalSol struct {
	alive_ssns []int
}

func init() {
	log = utils.GetLogger(chalName)
}

func main() {
	b := chalBody{}
	s := solBody{}
	resp := utils.GetChal(chalName)
	log.Debug(string(resp))
	json.Unmarshal(resp, &b)
	log.Debug(b)
	byteArray, err := base64.StdEncoding.DecodeString(b.B64Str)
	if err != nil {
		log.Errorf("err: %v\n", err)
	}
	log.Debugf("rawDecodedTxt: %v\n", byteArray)
	log.Infof("len(byteArray):%d", len(byteArray))
	reader, err := gzip.NewReader(bytes.NewReader(byteArray))
	if err != nil {
		log.Errorf("err: %v\n", err)
	}
	defer reader.Close()

	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		log.Errorf("err: %v\n", err)
	}
	dbRecovery := string(decompressedData)

	file, err := os.OpenFile("dump.sql", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Errorf("err creating file: %v\n", err)
	}
	defer file.Close()
	_, err = file.WriteString(dbRecovery)
	if err != nil {
		log.Errorf("err writing to recovery text to file : %v\n", err)
	}

	log.Info("dump file created")

	// Form the db connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	log.Info("Established a connection to the db successfully")

	// // Recover the db
	// _, err = db.Exec(dbRecovery)
	// if err != nil {
	// 	log.Fatalf("Err executing query %s", err)
	// }

	sPort := fmt.Sprintf("%d", port) // stringified port number
	_, err = db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	if err != nil {
		log.Fatalf("Err deleting everything from db %v", err)
	}
	cmd := exec.Command("psql", "-h", host, "-p", sPort, "-U", user, "-W", password, "-d", dbname, "-f", "dump.sql")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error restoring the db %v", err)
	}
	// query the database
	query := "SELECT * FROM criminal_records;"
	rows, err := db.Query(query)
	if err != nil {
		log.Errorf("Err executing db query %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var felony string
		var ssn string
		var home_add string
		var city string
		var status string
		var entry string

		if err := rows.Scan(&id, &name, &felony, &ssn, &home_add, &entry, &city, &status); err != nil {
			log.Fatal(err)
		}
		log.Debugf("%d %s %s %s %s %s %s", id, name, felony, ssn, home_add, city, status)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	solQuery := "SELECT ssn FROM criminal_records WHERE status ='alive';"
	solRows, err := db.Query(solQuery)
	if err != nil {
		log.Errorf("Err executing db query %s", err)
	}
	defer solRows.Close()
	for solRows.Next() {
		var ssn string
		if err := solRows.Scan(&ssn); err != nil {
			log.Fatal(err)
		}
		log.Debugf("ssn : %v", ssn)
		s.Ssns = append(s.Ssns, ssn)
	}
	jsonData, err := json.Marshal(s)
	if err != nil {
		log.Errorf("Err marshalling %v", err)
	}
	status := utils.SendSol(chalName, jsonData)
	log.Infof("Status: %v\n", status)
}
