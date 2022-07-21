package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type DistanceMatrixResponse struct {
	// OriginAddresses contains an array of addresses as returned by the API from
	// your original request.
	OriginAddresses []string `json:"origin_addresses"`
	// DestinationAddresses contains an array of addresses as returned by the API
	// from your original request.
	DestinationAddresses []string `json:"destination_addresses"`
	// Rows contains an array of elements.
	Rows []DistanceMatrixElementsRow `json:"rows"`
}

type DistanceMatrixElementsRow struct {
	Elements []*DistanceMatrixElement `json:"elements"`
}

type DistanceMatrixElement struct {
	Status string `json:"status"`
	// Duration is the length of time it takes to travel this route.
	Duration time.Duration `json:"duration"`
	// DurationInTraffic is the length of time it takes to travel this route
	// considering traffic.
	DurationInTraffic time.Duration `json:"duration_in_traffic"`
	// Distance is the total distance of this route.
	Distance Distance `json:"distance"`
}

type Distance struct {
	// HumanReadable is the human friendly distance. This is rounded and in an
	// appropriate unit for the request. The units can be overriden with a request
	// parameter.
	HumanReadable string `json:"text"`
	// Meters is the numeric distance, always in meters. This is intended to be used
	// only in algorithmic situations, e.g. sorting results by some user specified
	// metric.
	Meters int `json:"value"`
}

func DistanceMatrix(latSource, longSource, latDest, longDest string) string {
	if latSource == "" || longSource == "" {
		return "Distance Not Found! Please Allow Access to Your Location"
	} else {
		var apiKey = os.Getenv("API_KEY")
		url := "https://maps.googleapis.com/maps/api/distancematrix/json?origins=" + latSource + "," + longSource + "&destinations=" + latDest + "," + longDest + "&units=metric&key=" + apiKey
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return "Error"
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return "Error"
		}
		defer res.Body.Close()

		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var responseObject DistanceMatrixResponse
		json.Unmarshal(responseData, &responseObject)

		var distance string
		for i := 0; i < len(responseObject.Rows); i++ {
			distance = responseObject.Rows[i].Elements[i].Distance.HumanReadable
		}
		return distance
	}
}
