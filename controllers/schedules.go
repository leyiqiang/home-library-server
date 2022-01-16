package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/leyiqiang/home-library-server/models"
	"github.com/leyiqiang/home-library-server/utils"
	"net/http"
)

type SchedulePayload struct {
	models.Schedule
	Reservations []*models.Reservation `json:"reservations" bson:"reservations"`
}

// TODO check if reference work!
func (c *Controller) GetOneSchedule(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "scheduleID")
	var payload SchedulePayload
	schedule, err := c.Repo.GetScheduleByID(id)
	payload.Schedule = *schedule
	var reservations []*models.Reservation
	reservations, err = c.Repo.GetReservationsByScheduleID(schedule.ID.Hex())
	payload.Reservations = reservations
	err = utils.WriteJSON(w, http.StatusOK, schedule, "schedule")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

}

// TODO : be careful! need to delete reservations as well
func (c *Controller) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "scheduleID")

	err := c.Repo.DeleteScheduleByID(id)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
	err = utils.WriteJSON(w, http.StatusOK, "success", "message")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

}

// TODO remember to populate!
func (c *Controller) GetAllSchedules(w http.ResponseWriter, r *http.Request) {
	schedules, err := c.Repo.GetAllSchedules()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, schedules, "schedules")
}

// TODO init reservation with empty
func (c *Controller) AddSchedule(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var schedule models.Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	validate = validator.New()
	err = validate.Struct(schedule)
	if err != nil {
		utils.ErrorJSON(w, err)
		return

	}
	var oid string
	oid, err = c.Repo.AddSchedule(schedule)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, oid, "oid")

}

func (c *Controller) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := chi.URLParam(r, "scheduleID")
	var schedule models.Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	validate = validator.New()
	err = validate.Struct(schedule)
	if err != nil {
		utils.ErrorJSON(w, err)
		return

	}
	var updatedSchedule *models.Schedule
	updatedSchedule, err = c.Repo.UpdateScheduleByID(id, schedule)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, updatedSchedule, "schedule")
}
