package http

import (
	"fmt"
	"net"
	"net/http"

	"github.com/oschwald/geoip2-golang"
)

func geoipHandler(cityDB *geoip2.Reader, ispDB *geoip2.Reader) http.HandlerFunc {
	type response struct {
		ASN          string `json:"asn"`
		City         string `json:"city"`
		Country      string `json:"country"`
		ISP          string `json:"isp"`
		Organization string `json:"organization"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ip := net.ParseIP(r.PathValue("ip"))
		if ip == nil {
			respondError(w, nil, http.StatusBadRequest)
			return
		}

		city, err := cityDB.City(ip)
		if err != nil {
			respondError(w, err, http.StatusInternalServerError)
			return
		}

		isp, err := ispDB.ISP(ip)
		if err != nil {
			respondError(w, err, http.StatusInternalServerError)
			return
		}

		respondJSON(w, response{
			ASN:          fmt.Sprint(isp.AutonomousSystemNumber),
			City:         city.City.Names["en"],
			Country:      city.Country.Names["en"],
			ISP:          isp.ISP,
			Organization: isp.AutonomousSystemOrganization,
		})
	}
}
