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
