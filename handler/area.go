package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/anaabdi/location-mogo/repository"
)

type Area struct {
	areaRepo repository.Area
}

func NewAreaHandler(areaRepo repository.Area) *Area {
	return &Area{
		areaRepo: areaRepo,
	}
}

func (area *Area) GetByLocation() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		lng, _ := strconv.ParseFloat(query.Get("longitude"), 64)
		lat, _ := strconv.ParseFloat(query.Get("latitude"), 64)

		areaObj, err := area.areaRepo.GetByLocation(lng, lat)
		if err != nil {
			// TODO: structurized error response
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := json.Marshal(areaObj)
		if err != nil {
			// TODO: structurized error response
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})

}
