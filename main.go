package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const apiURL = "https://swd.weatherflow.com/swd/rest/observations/station"

var (
	// token is our weatherflow API token
	token = os.Getenv("WEATHERFLOW_API_TOKEN")
	// labels is a map of prometheus labels to apply to the metrics retrieved
	labels     prometheus.Labels
	labelNames []string
	// METRICS //
	// all metrics should be declared here and initialized below in the "init" function
	air_temperature *prometheus.GaugeVec
)

func getTempestData(token string) (response, error) {
	var r response
	reqURL := apiURL + "/" + "51384" + "?token=" + token
	httpResp, err := http.Get(reqURL)
	if err != nil {
		return r, fmt.Errorf("error getting data from tempest station: %v", err)
	}
	defer httpResp.Body.Close()
	err = json.NewDecoder(httpResp.Body).Decode(&r)
	if err != nil {
		return r, fmt.Errorf("error parsing json into response struct: %v", err)
	}
	return r, nil
}

func init() {
	// Initialize labels
	r, err := getTempestData(token)
	if err != nil {
		log.Fatal(err)
	}
	labels = r.parseLabels()
	labelNames = []string{}
	for k := range labels {
		labelNames = append(labelNames, k)
	}
	// Initialze metrics
	// TODO see if we can hack this with some reflection?
	air_temperature = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "air_temperature",
			Help: "Temperature in Celcius",
		},
		labelNames,
	)
	prometheus.MustRegister(air_temperature)
}

// stationStatus holds our station status code
type stationStatus struct {
	Code int `json:"status_code"`
}

// observation is the typed observation data from a station
type observation struct {
	AirDensity                       float64 `json:"air_density"`
	AirTemperature                   float64 `json:"air_temperature"`
	BarometricPressure               float64 `json:"barometric_pressure"`
	Brightness                       float64 `json:"brightness"`
	DeltaT                           float64 `json:"delta_t"`
	DewPoint                         float64 `json:"dew_point"`
	FeelsLike                        float64 `json:"feels_like"`
	HeatIndex                        float64 `json:"heat_index"`
	LightningStrikeCount             float64 `json:"lightning_strike_count"`
	LightningStrikeCountLast1hr      float64 `json:"lightning_strike_count_last_1hr"`
	LightningStrikeCountLast3hr      float64 `json:"lightning_strike_count_last_3hr"`
	LightningStrikeLastDistance      float64 `json:"lightning_strike_last_distance"`
	LightningStrikeLastEpoch         float64 `json:"lightning_strike_last_epoch"`
	Precip                           float64 `json:"precip"`
	PrecipAccumLast1hr               float64 `json:"precip_accum_last_1hr"`
	PrecipAccumLocalDay              float64 `json:"precip_accum_local_day"`
	PrecipAccumLocalYesterday        float64 `json:"precip_accum_local_yesterday"`
	PrecipAccumLocalYesterdayFinal   float64 `json:"precip_accum_local_yesterday_final"`
	PrecipAnalysisTypeYesterday      float64 `json:"precip_analysis_type_yesterday"`
	PrecipMinutesLocalDay            float64 `json:"precip_minutes_local_day"`
	PrecipMinutesLocalYesterday      float64 `json:"precip_minutes_local_yesterday"`
	PrecipMinutesLocalYesterdayFinal float64 `json:"precip_minutes_local_yesterday_final"`
	PressureTrend                    string  `json:"pressure_trend"`
	RelativeHumidity                 float64 `json:"relative_humidity"`
	SeaLevelPressure                 float64 `json:"sea_level_pressure"`
	SolarRadiation                   float64 `json:"solar_radiation"`
	StationPressure                  float64 `json:"station_pressure"`
	Timestamp                        float64 `json:"timestamp"`
	Uv                               float64 `json:"uv"`
	WetBulbTemperature               float64 `json:"wet_bulb_temperature"`
	WindAvg                          float64 `json:"wind_avg"`
	WindChill                        float64 `json:"wind_chill"`
	WindDirection                    float64 `json:"wind_direction"`
	WindGust                         float64 `json:"wind_gust"`
	WindLull                         float64 `json:"wind_lull"`
}

// response is our response from the weatherflow obvservations API
type response struct {
	StationId   int           `json:"station_id"`
	StationName string        `json:"station_name"`
	PublicName  string        `json:"public_name"`
	Latitude    float64       `json:"latitude"`
	Longitude   float64       `json:"longitude"`
	Timezone    string        `json:"timezone"`
	Elevation   float64       `json:"elevation"`
	Status      stationStatus `json:"status"`
	Obs         []observation `json:"obs"`
}

// parseLabels returns a list of label values as strings matchingour "labels" var
func (r *response) parseLabels() prometheus.Labels {
	l := make(map[string]string)
	l["station_id"] = strconv.Itoa(r.StationId)
	l["station_name"] = r.StationName
	l["public_name"] = r.PublicName
	l["latitude"] = strconv.FormatFloat(r.Latitude, 'E', -1, 64)
	l["longitude"] = strconv.FormatFloat(r.Longitude, 'E', -1, 64)
	l["timezone"] = r.Timezone
	l["elevation"] = strconv.FormatFloat(r.Elevation, 'E', -1, 64)
	return l
}

func getDatas() {
	r, err := getTempestData(token)
	if err != nil {
		log.Fatal(err)
	}
	labels = r.parseLabels()
	air_temperature.With(labels).Set(r.Obs[0].AirTemperature)
	time.Sleep(time.Second * 15)
}

func main() {
	go getDatas()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:6969", nil)
}
