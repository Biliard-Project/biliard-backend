package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Biliard-Project/biliard-backend/models"
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
