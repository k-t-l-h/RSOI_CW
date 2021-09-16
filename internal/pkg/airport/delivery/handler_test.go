package delivery

import (
	"RSOI_CW/internal/pkg/airport"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAirport(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mock := airport.NewMockIRepo(ctl)
	handler := NewAirportHandler(mock)

	mock.EXPECT().SelectAirport(gomock.Any()).Return().Times(0)
	r := httptest.NewRequest("POST", "/persons", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.GetAirport(w, r)
}

func TestGetAirports(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mock := airport.NewMockIRepo(ctl)
	handler := NewAirportHandler(mock)

	mock.EXPECT().SelectAirports().Return().Times(0)
	r := httptest.NewRequest("POST", "/persons", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.GetAirports(w, r)
}
