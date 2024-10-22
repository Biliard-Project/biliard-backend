package models

import "database/sql"

type Record struct {
	ID         int
	PatientID  int
	TestDate   string
	Keterangan string
}

type RecordService struct {
	DB *sql.DB
}

func (rs *RecordService) GetAllRecordService() (*[]Record, error) {
	// TODO: implement how to get all records
	return nil, nil
}
