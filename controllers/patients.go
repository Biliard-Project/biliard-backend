package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Biliard-Project/biliard-backend/models"
	"github.com/go-chi/chi/v5"
)

type Patients struct {
	PatientService *models.PatientService
}

func (pt Patients) ProcessGetPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := pt.PatientService.RetrieveAllPatients()
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

	patientDB, err := pt.PatientService.Create(patient.Name, patient.Gender, patient.Keterangan, patient.BirthDate)
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

func (pt Patients) ProcessGetPatientByID(w http.ResponseWriter, r *http.Request) {
	patientID, err := strconv.Atoi(chi.URLParam(r, "patientID"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error getting patient by patient id", http.StatusBadRequest)
		return
	}

	patientDB, err := pt.PatientService.RetrievePatientByID(patientID)
	if err != nil {
		fmt.Println("jsoning 1")
		fmt.Println(err)
		http.Error(w, "error getting patient by patient id", http.StatusInternalServerError)
		return
	}

	patientJson, err := json.Marshal(patientDB)
	if err != nil {
		fmt.Println("jsoning 2")
		http.Error(w, "error getting patient by patient id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(patientJson))
}

func (pt Patients) DeletePatientByID(w http.ResponseWriter, r *http.Request) {
	patientID, err := strconv.Atoi(chi.URLParam(r, "patientID"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error deleting patient", http.StatusBadRequest)
		return
	}

	err = pt.PatientService.Delete(patientID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error deleting patient", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message":"success"}`)
}

func (pt Patients) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "updating patient: error while decoding json", http.StatusBadRequest)
		return
	}

	err = pt.PatientService.UpdatePatient(patient)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "updating patient: error while updating patient inside sql", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message":"success"}`)
}
