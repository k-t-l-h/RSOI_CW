package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/flights"
	"RSOI_CW/internal/pkg/middleware"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"net/http"
)

type FlightHandler struct {
	repo flights.IRepo
}

func NewFlightHandler(repo flights.IRepo) *FlightHandler {
	return &FlightHandler{repo: repo}
}

//список всех рейсов
//GET /flights
func (h *FlightHandler) AllFlights(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8887")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

	if r.Method == http.MethodOptions {
		return
	}
	readFlights, status := h.repo.ReadFlights()
	middleware.Response(w, status, readFlights)
}

//добавить рейс
//POST /flights
func (h *FlightHandler) AddFlight(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8887")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

	if r.Method == http.MethodOptions {
		return
	}
	flight := models.Flight{}
	err := easyjson.UnmarshalFromReader(
		r.Body,
		&flight,
	)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
		return
	}

	flight.ID = uuid.New()
	status := h.repo.CreateFlight(flight)
	middleware.Response(w, status, nil)

}

//изменить рейс
//PATCH /flights
func (h *FlightHandler) UpdateFlight(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8887")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE")

	if r.Method == http.MethodOptions {
		return
	}

	vars := mux.Vars(r)
	uuids, _ := vars["UUID"]
	id, err := uuid.Parse(uuids)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
		return
	}
	flight := models.Flight{}
	err = easyjson.UnmarshalFromReader(
		r.Body,
		&flight,
	)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
		return
	}
	status := h.repo.UpdateFlight(id, flight)
	middleware.Response(w, status, nil)

}
