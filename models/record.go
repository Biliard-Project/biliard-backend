package models

import (
	"database/sql"
	"fmt"
)

type PatientRecord struct {
	PatientID  int      `json:"patient_id"`
	RecordID   int      `json:"record_id"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	BirthDate  JSONTime `json:"birth_date"`
	TestDate   string   `json:"test_date"`
	Keterangan string   `json:"keterangan"`
}

type Record struct {
	ID         int
	PatientID  int
	TestDate   JSONTime
	Keterangan string
}

type RecordService struct {
	DB *sql.DB
}

func (rs *RecordService) RetrievePatientRecord() (*[]PatientRecord, error) {
	patientRecords := make([]PatientRecord, 0)
	rows, err := rs.DB.Query(`
		SELECT patients.id, records.id, patients.name, patients.gender, patients.birth_date, records.test_date, records.keterangan
		FROM records
		JOIN patients ON patients.id = records.patient_id
	`)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("RetrievePatientRecord: %w", err)
	}

	for rows.Next() {
		var patientrecord PatientRecord
		err = rows.Scan(&patientrecord.PatientID, &patientrecord.RecordID, &patientrecord.Name, &patientrecord.Gender, &patientrecord.BirthDate, &patientrecord.TestDate, &patientrecord.Keterangan)
		if err != nil {
			return nil, fmt.Errorf("RetrievePatientRecord: %w", err)
		}

		patientRecords = append(patientRecords, patientrecord)
	}

	return &patientRecords, nil
}
