package http

import (
	"net/http"
	"strconv"
)

func bmiHandler() http.HandlerFunc {
	type response struct {
		BMI     float64 `json:"bmi"`
		Verdict string  `json:"verdict"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		heightCmRaw := r.URL.Query().Get("height_cm")
		heightCm, err := strconv.ParseFloat(heightCmRaw, 64)
		if err != nil {
			respondError(w, nil, http.StatusBadRequest)
		}

		weightKgRaw := r.URL.Query().Get("weight_kg")
		weightKg, err := strconv.ParseFloat(weightKgRaw, 64)
		if err != nil {
			respondError(w, nil, http.StatusBadRequest)
		}

		heightM := heightCm / 100
		bmi := weightKg / (heightM * heightM)

		// Based on WHO
		var verdict string
		switch {
		case bmi > 40:
			verdict = "obese_class_3"
		case bmi >= 35:
			verdict = "obese_class_2"
		case bmi >= 30:
			verdict = "obese_class_1"
		case bmi >= 25:
			verdict = "overweight"
		case bmi >= 18.5:
			verdict = "normal"
		case bmi >= 17:
			verdict = "mild_thinness"
		case bmi >= 16:
			verdict = "moderate_thinness"
		default:
			verdict = "severe_thinness"
		}

		respondJSON(w, response{
			BMI:     bmi,
			Verdict: verdict,
		})
	}
}
