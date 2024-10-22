package models

import (
	"database/sql"
	"fmt"
)

type Patient struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Gender    string   `json:"gender"`
	BirthDate JSONTime `json:"birth_date"`
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
		err = rows.Scan(&patient.ID, &patient.Name, &patient.Gender, &patient.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("GetAllPatients: %w", err)
		}
		patients = append(patients, patient)
	}

	return &patients, nil
}

func (ps *PatientService) Create(name, gender string, birthDate JSONTime) (*Patient, error) {
	patient := Patient{
		Name:      name,
		Gender:    gender,
		BirthDate: birthDate,
	}
	row := ps.DB.QueryRow(`insert into patients(name, gender, birth_date)
		values ($1, $2, $3) returning id;`, patient.Name, patient.Gender, patient.BirthDate.ConvertToYMD())
	err := row.Scan(&patient.ID)
	if err != nil {
		return nil, fmt.Errorf("create patient: %w", err)
	}

	return &patient, nil
}
