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
	Keterangan string   `json:"keterangan"`
	TestDate   JSONTime `json:"test_date"`
	Bilirubin  float64  `json:"bilirubin"`
	Oxygen     float64  `json:"oxygen"`
	HeartRate  float64  `json:"heart_rate"`
}

type Record struct {
	ID        int      `json:"id"`
	PatientID int      `json:"patient_id"`
	TestDate  JSONTime `json:"test_date"`
	Bilirubin float64  `json:"bilirubin"`
	Oxygen    float64  `json:"oxygen"`
	HeartRate float64  `json:"heart_rate"`
}

type RecordService struct {
	DB *sql.DB
}

func (rs *RecordService) RetrievePatientRecord() (*[]PatientRecord, error) {
	patientRecords := make([]PatientRecord, 0)
	rows, err := rs.DB.Query(`
		SELECT patients.id, records.id, patients.name, patients.gender, patients.birth_date, patients.keterangan, records.test_date, records.bilirubin, records.oxygen, records.heart_rate 
		FROM records
		JOIN patients ON patients.id = records.patient_id;
	`)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("RetrievePatientRecord: %w", err)
	}

	for rows.Next() {
		var patientrecord PatientRecord
		err = rows.Scan(&patientrecord.PatientID, &patientrecord.RecordID, &patientrecord.Name, &patientrecord.Gender, &patientrecord.BirthDate, &patientrecord.Keterangan, &patientrecord.TestDate, &patientrecord.Bilirubin, &patientrecord.Oxygen, &patientrecord.HeartRate)
		if err != nil {
			return nil, fmt.Errorf("RetrievePatientRecord: %w", err)
		}

		patientRecords = append(patientRecords, patientrecord)
	}

	return &patientRecords, nil
}

func (rs *RecordService) RetrieveRecordsByPatientID(patientID int) (*[]Record, error) {
	records := make([]Record, 0)
	rows, err := rs.DB.Query(`
		SELECT records.id, records.patient_id, records.test_date, records.bilirubin, records.oxygen, records.heart_rate
		FROM records
		where records.patient_id = $1;
	`, patientID)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("retrieverecordbyid: %w", err)
	}

	for rows.Next() {
		var record Record
		err = rows.Scan(&record.ID, &record.PatientID, &record.TestDate, &record.Bilirubin, &record.Oxygen, &record.HeartRate)
		if err != nil {
			return nil, fmt.Errorf("retrieverecordbyid: %w", err)
		}
		records = append(records, record)
	}

	return &records, nil
}

func (rs *RecordService) InsertNewRecord(patientID int, testDate JSONTime, bilirubin, oxygen, heart_rate float64) (*Record, error) {
	row := rs.DB.QueryRow(`
		INSERT INTO records (patient_id, test_date, bilirubin, oxygen, heart_rate)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
		`, patientID, testDate.ConvertToYMD(), bilirubin, oxygen, heart_rate)

	record := Record{
		PatientID: patientID,
		TestDate:  testDate,
		Bilirubin: bilirubin,
		Oxygen:    oxygen,
		HeartRate: heart_rate,
	}
	err := row.Scan(&record.ID)
	if err != nil {
		return nil, fmt.Errorf("insert new record: %w", err)
	}
	return &record, nil
}
