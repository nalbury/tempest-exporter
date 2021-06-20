package main

import "github.com/prometheus/client_golang/prometheus"

type MetricsMap map[string]*prometheus.GaugeVec

// Register populates and registers all metrics for the expoter
func (m MetricsMap) Register(labelsNames []string) {
	m["air_density"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "air_density",
			Help:      "Air Density",
		},
		labelNames,
	)
	m["air_temperature"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "air_temperature",
			Help:      "Air Temperature",
		},
		labelNames,
	)
	m["barometric_pressure"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "barometric_pressure",
			Help:      "Barometric Pressure",
		},
		labelNames,
	)
	m["brightness"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "brightness",
			Help:      "Brightness",
		},
		labelNames,
	)
	m["delta_t"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "delta_t",
			Help:      "Delta T",
		},
		labelNames,
	)
	m["dew_point"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "dew_point",
			Help:      "Dew Point",
		},
		labelNames,
	)
	m["feels_like"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "feels_like",
			Help:      "Feels Like",
		},
		labelNames,
	)
	m["heat_index"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "heat_index",
			Help:      "Heat Index",
		},
		labelNames,
	)
	m["lightning_strike_count"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "lightning_strike_count",
			Help:      "Lightning Strike Count",
		},
		labelNames,
	)
	m["lightning_strike_count_last_1hr"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "lightning_strike_count_last_1hr",
			Help:      "Lightning Strike Count Last 1hr",
		},
		labelNames,
	)
	m["lightning_strike_count_last_3hr"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "lightning_strike_count_last_3hr",
			Help:      "Lightning Strike Count Last 3hr",
		},
		labelNames,
	)
	m["lightning_strike_last_distance"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "lightning_strike_last_distance",
			Help:      "Lightning Strike Last Distance",
		},
		labelNames,
	)
	m["lightning_strike_last_epoch"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "lightning_strike_last_epoch",
			Help:      "Lightning Strike Last Epoch",
		},
		labelNames,
	)
	m["precip"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "precip",
			Help:      "Precip",
		},
		labelNames,
	)
	m["precip_accum_last_1hr"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "precip_accum_last_1hr",
			Help:      "Precip Accum Last 1hr",
		},
		labelNames,
	)
	m["precip_accum_local_day"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "precip_accum_local_day",
			Help:      "Precip Accum Local Day",
		},
		labelNames,
	)
	m["precip_accum_local_yesterday"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "precip_accum_local_yesterday",
			Help:      "Precip Accum Local Yesterday",
		},
		labelNames,
	)
	m["precip_accum_local_yesterday_final"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "precip_accum_local_yesterday_final",
			Help:      "Precip Accum Local Yesterday Final",
		},
		labelNames,
	)
	m["precip_analysis_type_yesterday"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "precip_analysis_type_yesterday",
			Help:      "Precip Analysis Type Yesterday",
		},
		labelNames,
	)
	m["precip_minutes_local_day"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "precip_minutes_local_day",
			Help:      "Precip Minutes Local Day",
		},
		labelNames,
	)
	m["precip_minutes_local_yesterday"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "precip_minutes_local_yesterday",
			Help:      "Precip Minutes Local Yesterday",
		},
		labelNames,
	)
	m["precip_minutes_local_yesterday_final"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "precip_minutes_local_yesterday_final",
			Help:      "Precip Minutes Local Yesterday Final",
		},
		labelNames,
	)
	m["pressure_trend"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "pressure_trend",
			Help:      "Pressure Trend",
		},
		labelNames,
	)
	m["relative_humidity"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "relative_humidity",
			Help:      "Relative Humidity",
		},
		labelNames,
	)
	m["sea_level_pressure"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "sea_level_pressure",
			Help:      "Sea Level Pressure",
		},
		labelNames,
	)
	m["solar_radiation"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "solar_radiation",
			Help:      "Solar Radiation",
		},
		labelNames,
	)
	m["station_pressure"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "station_pressure",
			Help:      "Station Pressure",
		},
		labelNames,
	)
	m["timestamp"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "timestamp",
			Help:      "Timestamp",
		},
		labelNames,
	)
	m["uv"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "uv",
			Help:      "Uv",
		},
		labelNames,
	)
	m["wet_bulb_temperature"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "wet_bulb_temperature",
			Help:      "Wet Bulb Temperature",
		},
		labelNames,
	)
	m["wind_avg"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "wind_avg",
			Help:      "Wind Avg",
		},
		labelNames,
	)
	m["wind_chill"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "wind_chill",
			Help:      "Wind Chill",
		},
		labelNames,
	)
	m["wind_direction"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "wind_direction",
			Help:      "Wind Direction",
		},
		labelNames,
	)
	m["wind_gust"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "wind_gust",
			Help:      "Wind Gust",
		},
		labelNames,
	)
	m["wind_lull"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: ss,
			Name:      "wind_lull",
			Help:      "Wind Lull",
		},
		labelNames,
	)

	// Register all metrics in our MetricsMap
	for _, met := range m {
		prometheus.MustRegister(met)
	}
}

func (m MetricsMap) SetAll(o observation, labels prometheus.Labels) {
	metrics["air_density"].With(labels).Set(o.AirDensity)
	metrics["air_temperature"].With(labels).Set(o.AirTemperature)
	metrics["barometric_pressure"].With(labels).Set(o.BarometricPressure)
	metrics["brightness"].With(labels).Set(o.Brightness)
	metrics["delta_t"].With(labels).Set(o.DeltaT)
	metrics["dew_point"].With(labels).Set(o.DewPoint)
	metrics["feels_like"].With(labels).Set(o.FeelsLike)
	metrics["heat_index"].With(labels).Set(o.HeatIndex)
	metrics["lightning_strike_count"].With(labels).Set(o.LightningStrikeCount)
	metrics["lightning_strike_count_last_1hr"].With(labels).Set(o.LightningStrikeCountLast1hr)
	metrics["lightning_strike_count_last_3hr"].With(labels).Set(o.LightningStrikeCountLast3hr)
	metrics["lightning_strike_last_distance"].With(labels).Set(o.LightningStrikeLastDistance)
	metrics["lightning_strike_last_epoch"].With(labels).Set(o.LightningStrikeLastEpoch)
	metrics["precip"].With(labels).Set(o.Precip)
	metrics["precip_accum_last_1hr"].With(labels).Set(o.PrecipAccumLast1hr)
	metrics["precip_accum_local_day"].With(labels).Set(o.PrecipAccumLocalDay)
	metrics["precip_accum_local_yesterday"].With(labels).Set(o.PrecipAccumLocalYesterday)
	metrics["precip_accum_local_yesterday_final"].With(labels).Set(o.PrecipAccumLocalYesterdayFinal)
	metrics["precip_analysis_type_yesterday"].With(labels).Set(o.PrecipAnalysisTypeYesterday)
	metrics["precip_minutes_local_day"].With(labels).Set(o.PrecipMinutesLocalDay)
	metrics["precip_minutes_local_yesterday"].With(labels).Set(o.PrecipMinutesLocalYesterday)
	metrics["precip_minutes_local_yesterday_final"].With(labels).Set(o.PrecipMinutesLocalYesterdayFinal)
	// TODO convert this to a numeric data point
	//metrics["pressure_trend"].With(labels).Set(o.PressureTrend)
	metrics["relative_humidity"].With(labels).Set(o.RelativeHumidity)
	metrics["sea_level_pressure"].With(labels).Set(o.SeaLevelPressure)
	metrics["solar_radiation"].With(labels).Set(o.SolarRadiation)
	metrics["station_pressure"].With(labels).Set(o.StationPressure)
	metrics["timestamp"].With(labels).Set(o.Timestamp)
	metrics["uv"].With(labels).Set(o.Uv)
	metrics["wet_bulb_temperature"].With(labels).Set(o.WetBulbTemperature)
	metrics["wind_avg"].With(labels).Set(o.WindAvg)
	metrics["wind_chill"].With(labels).Set(o.WindChill)
	metrics["wind_direction"].With(labels).Set(o.WindDirection)
	metrics["wind_gust"].With(labels).Set(o.WindGust)
	metrics["wind_lull"].With(labels).Set(o.WindLull)
}
