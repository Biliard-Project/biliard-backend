package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Biliard-Project/biliard-backend/models"
	"github.com/go-chi/chi/v5"
)

type PatientScans struct {
	PatientScansService *models.PatientScanService
}

func (pts PatientScans) Set(w http.ResponseWriter, r *http.Request) {
	patientID, err := strconv.Atoi(chi.URLParam(r, "patientID"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error setting patient to scan, make sure the patient id is correct", http.StatusBadRequest)
		return
	}

	err = pts.PatientScansService.Update(patientID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error setting patient to scan, make sure patientID actually exists", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message":"success"}`)
}

func (pts PatientScans) Get(w http.ResponseWriter, r *http.Request) {
	patient, err := pts.PatientScansService.Get()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error getting current patient to scan", http.StatusInternalServerError)
		return
	}

	patientJson, err := json.Marshal(patient)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error getting current patient to scan", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(patientJson))
}
