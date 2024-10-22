package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Biliard-Project/biliard-backend/models"
)

type Patients struct {
	PatientService *models.PatientService
}

func (pt Patients) ProcessGetPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := pt.PatientService.GetAllPatients()
	if err != nil {
		http.Error(w, "Error retrieving patients", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	patientsJson, err := json.Marshal(patients)
	if err != nil {
		fmt.Println("jsoning")
		http.Error(w, "Error retrieving patients", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(patientsJson))
}

func (pt Patients) Create(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		fmt.Println("jsoning 0")
		http.Error(w, "Error Creating patient", http.StatusBadRequest)
		return
	}

	patientDB, err := pt.PatientService.Create(patient.Name, patient.Gender, patient.BirthDate)
	if err != nil {
		fmt.Println("jsoning 1")
		fmt.Println(err)
		http.Error(w, "Error Creating patient", http.StatusInternalServerError)
		return
	}

	patientJson, err := json.Marshal(patientDB)
	if err != nil {
		fmt.Println("jsoning 2")
		http.Error(w, "Error Creating patient", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(patientJson))
}
