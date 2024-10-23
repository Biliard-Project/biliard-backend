package models

import (
	"database/sql"
	"fmt"
)

type Patient struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	BirthDate  JSONTime `json:"birth_date"`
	Keterangan string   `json:"keterangan"`
}

type PatientService struct {
	DB *sql.DB
}

func (ps *PatientService) RetrieveAllPatients() (*[]Patient, error) {
	// TODO: implement how to get all patients
	rows, err := ps.DB.Query(`select id, name, gender, birth_date, keterangan from patients;`)
	if err != nil {
		return nil, fmt.Errorf("GetAllPatients: %w", err)
	}
	defer rows.Close()

	patients := make([]Patient, 0)
	for rows.Next() {
		var patient Patient
		err = rows.Scan(&patient.ID, &patient.Name, &patient.Gender, &patient.BirthDate, &patient.Keterangan)
		if err != nil {
			return nil, fmt.Errorf("GetAllPatients: %w", err)
		}
		patients = append(patients, patient)
	}

	return &patients, nil
}

func (ps *PatientService) Create(name, gender, keterangan string, birthDate JSONTime) (*Patient, error) {
	patient := Patient{
		Name:       name,
		Gender:     gender,
		BirthDate:  birthDate,
		Keterangan: keterangan,
	}
	row := ps.DB.QueryRow(`insert into patients(name, gender, birth_date, keterangan)
		values ($1, $2, $3, $4) returning id;`, patient.Name, patient.Gender, patient.BirthDate.ConvertToYMD(), patient.Keterangan)
	err := row.Scan(&patient.ID)
	if err != nil {
		return nil, fmt.Errorf("create patient: %w", err)
	}

	return &patient, nil
}

func (ps *PatientService) RetrievePatientByID(id int) (*Patient, error) {
	patient := Patient{
		ID: id,
	}
	row := ps.DB.QueryRow(`select name, gender, birth_date, keterangan
		from patients where id = $1`, patient.ID)
	err := row.Scan(&patient.Name, &patient.Gender, &patient.BirthDate, &patient.Keterangan)
	if err != nil {
		return nil, fmt.Errorf("create patient: %w", err)
	}

	return &patient, nil
}
