package models

import (
	"database/sql"
	"fmt"
)

type Patient struct {
	ID        int
	Name      string
	Gender    string
	BirthDate JSONTime
}

type PatientService struct {
	DB *sql.DB
}

func (ps *PatientService) GetAllPatients() (*[]Patient, error) {
	// TODO: implement how to get all patients
	rows, err := ps.DB.Query(`select id, name, gender, birth_date from patients;`)
	if err != nil {
		return nil, fmt.Errorf("GetAllPatients: %w", err)
	}
	defer rows.Close()

	patients := make([]Patient, 0)
	for rows.Next() {
		var patient Patient
		err = rows.Scan(&patient.Name, &patient.Name, &patient.Gender, &patient.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("GetAllPatients: %w", err)
		}
		patients = append(patients, patient)
	}

	return &patients, nil
}
