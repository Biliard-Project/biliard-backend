package models

import (
	"database/sql"
	"fmt"
)

type PatientScanService struct {
	DB *sql.DB
}

func (pss *PatientScanService) Update(patientID int) error {
	_, err := pss.DB.Exec(`
    UPDATE patient_scans
    SET patient_id = $1
    WHERE id = 1;
  `, patientID)
	if err != nil {
		return fmt.Errorf("update patient_scan: %w", err)
	}

	return nil
}

func (pss *PatientScanService) Get() (*Patient, error) {
	var patient Patient
	row := pss.DB.QueryRow(`
		SELECT patients.id, patients.name, patients.gender, patients.birth_date, patients.keterangan
		FROM patient_scans
		JOIN patients ON patients.id = patient_scans.patient_id
		WHERE  patient_scans.id = 1;
	`)
	err := row.Scan(&patient.ID, &patient.Name, &patient.Gender, &patient.BirthDate, &patient.Keterangan)
	if err != nil {
		return nil, fmt.Errorf("get current patient for scan: %w", err)
	}

	return &patient, err
}
