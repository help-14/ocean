package db

func AddLog(data struct {
	job          string
	path         string
	success      bool
	errorMessage string
}) error {
	checkDb()
	insertSQL := `INSERT INTO logs(job, path, success, error) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		AddSystemLog(serviceTag, err)
		return err
	}
	_, err = statement.Exec(data.job, data.path, data.success, data.errorMessage)
	if err != nil {
		AddSystemLog(serviceTag, err)
		return err
	}
	return nil
}
