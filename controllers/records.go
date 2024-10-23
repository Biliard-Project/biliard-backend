package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Biliard-Project/biliard-backend/models"
	"github.com/go-chi/chi/v5"
)

type Records struct {
	RecordService *models.RecordService
}

func (rd Records) GetAllPatientRecords(w http.ResponseWriter, r *http.Request) {
	patientRecords, err := rd.RecordService.RetrievePatientRecord()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error getting patient records", http.StatusInternalServerError)
		return
	}

	patientRecordsJson, err := json.Marshal(patientRecords)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error getting patient records", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(patientRecordsJson))
}

func (rd Records) GetRecordsByPatientID(w http.ResponseWriter, r *http.Request) {
	patientID, err := strconv.Atoi(chi.URLParam(r, "patientID"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error getting records by patient id", http.StatusBadRequest)
		return
	}
	records, err := rd.RecordService.RetrieveRecordsByPatientID(patientID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error getting records by patient id", http.StatusInternalServerError)
		return
	}

	recordsJson, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error getting records by patient id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(recordsJson))
}
