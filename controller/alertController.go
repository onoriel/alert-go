package controller

import (
	"encoding/json"

	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	. "alerts/models"
	"alerts/repository"

	"github.com/gorilla/mux"
)

type AlertController struct {
	repository repository.AlertsRepository
}

func NewAlertController(rep repository.AlertsRepository) *AlertController {
	return &AlertController{rep}
}

// GET list of alerts
func (ac *AlertController) AllAlertsEndPoint(w http.ResponseWriter, r *http.Request) {
	alerts, err := ac.repository.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, alerts)
}

// GET an alert by its ID
func (ac *AlertController) FindAlertEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	alert, err := ac.repository.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Alert ID")
		return
	}
	respondWithJson(w, http.StatusOK, alert)
}

// POST a new alert
func (ac *AlertController) CreateAlertEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var alert Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	alert.ID = primitive.NewObjectID()
	if err := ac.repository.Insert(alert); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, alert)
}

// PUT update an existing alert
func (ac *AlertController) UpdateAlertEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var alert Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := ac.repository.Update(alert); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing alert
func (ac *AlertController) DeleteAlertEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var alert Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := ac.repository.Delete(alert); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
