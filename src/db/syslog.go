package db

import "log"

func AddSystemLog(tag string, err error) error {
	checkDb()
	insertSQL := `INSERT INTO syslogs(tag, error) VALUES (?, ?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = statement.Exec(tag, err.Error())
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
