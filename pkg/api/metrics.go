package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	challengeUpdatesTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "acmedns_challenge_updates_total",
			Help: "Total ACME challenge TXT record update attempts.",
		},
		[]string{"result"},
	)

	registrationsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "acmedns_registrations_total",
			Help: "Total ACME DNS account registration attempts.",
		},
		[]string{"result"},
	)

	registeredAccounts = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "acmedns_registered_accounts",
			Help: "Current number of registered ACME DNS accounts.",
		},
	)
)

const (
	metricLabelSuccess = "success"
	metricLabelFailure = "failure"
)
