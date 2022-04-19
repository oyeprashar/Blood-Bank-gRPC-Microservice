package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
    SERVICE_NAME = "BloodBankSystemService"
    DATABASE = "database"
)

var Metrics MetricWrapper

var (
	FindPassword_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "NF_BBSS_FINDPASSWORD",
		Help: "Sumary metrics for FindPassword",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},[]string{"nservice","nmethod","ncode"})
)

var (
	FindPassword_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_BBSS_FINDPASSWORD_Error",
		Help: "Sumary metrics for FindPassword_Error",
	},[]string{"nservice","ntype","nerror"})
)

var (
	AddUser_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "NF_BBSS_ADDUSER",
		Help: "Sumary metrics for AddUser",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},[]string{"nservice","nmethod","ncode"})
)

var (
	AddUser_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_BBSS_ADDUSER_Error",
		Help: "Sumary metrics for AddUser_Error",
	},[]string{"nservice","ntype","nerror"})
)

var (
	BulkAddUser_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "NF_BBSS_BULKADDUSER",
		Help: "Sumary metrics for BulkAddUser",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},[]string{"nservice","nmethod","ncode"})
)

var (
	BulkAddUser_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_BBSS_BULKADDUSER_Error",
		Help: "Sumary metrics for BulkAddUser_Error",
	},[]string{"nservice","ntype","nerror"})
)

var (
	FindBlood_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "NF_BBSS_FINDBLOOD",
		Help: "Sumary metrics for FindBlood",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},[]string{"nservice","nmethod","ncode"})
)

var (
	FindBlood_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_BBSS_FINDBLOOD_Error",
		Help: "Sumary metrics for FindBlood_Error",
	},[]string{"nservice","ntype","nerror"})
)

var (
	AddBlood_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "NF_BBSS_ADDBLOOD",
		Help: "Sumary metrics for AddBlood",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},[]string{"nservice","nmethod","ncode"})
)

var (
	AddBlood_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_BBSS_ADDBLOOD_Error",
		Help: "Sumary metrics for AddBlood_Error",
	},[]string{"nservice","ntype","nerror"})
)

var (
	BulkAddBlood_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "NF_BBSS_BULKADDBLOOD",
		Help: "Sumary metrics for BulkAddBlood",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},[]string{"nservice","nmethod","ncode"})
)

var (
	BulkAddBlood_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_BBSS_BULKADDBLOOD_Error",
		Help: "Sumary metrics for BulkAddBlood_Error",
	},[]string{"nservice","ntype","nerror"})
)



func init() {
	Metrics = &Helper{
    	SERVICE_NAME: SERVICE_NAME,
    	DATABASE: DATABASE,
    }
	    prometheus.MustRegister(FindPassword_Metrics)
    prometheus.MustRegister(FindPassword_Error_Metrics)
    prometheus.MustRegister(AddUser_Metrics)
    prometheus.MustRegister(AddUser_Error_Metrics)
    prometheus.MustRegister(BulkAddUser_Metrics)
    prometheus.MustRegister(BulkAddUser_Error_Metrics)
    prometheus.MustRegister(FindBlood_Metrics)
    prometheus.MustRegister(FindBlood_Error_Metrics)
    prometheus.MustRegister(AddBlood_Metrics)
    prometheus.MustRegister(AddBlood_Error_Metrics)
    prometheus.MustRegister(BulkAddBlood_Metrics)
    prometheus.MustRegister(BulkAddBlood_Error_Metrics)

}
