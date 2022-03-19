package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arcology-network/frontend-svc/backend"
	"github.com/arcology-network/frontend-svc/handler"
	"github.com/arcology-network/frontend-svc/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port            int    `yaml:"port"`
		AccessTokenHash string `yaml:"accessTokenHash"`
	} `yaml:"server"`
	ZooKeeper struct {
		Servers []string `yaml:"servers"`
	} `yaml:"zk"`
}

func main() {
	configPath := "./config.yml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		panic(err)
	}

	handler.Init(cfg.Server.AccessTokenHash)
	backend.InitParams(cfg.ZooKeeper.Servers)

	duration := promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "process_duration",
		Help: "HTTP request process duration.",
	}, []string{"function"})

	// http.Handle("/metrics", promhttp.Handler())
	// go http.ListenAndServe(":19001", nil)

	router := httprouter.New()
	router.GET("/latestheight", middleware.PrometheusMiddleware(duration.WithLabelValues("GetLatestHeight"), handler.GetLatestHeight))
	router.GET("/nonce/:address", middleware.PrometheusMiddleware(duration.WithLabelValues("GetNonce"), handler.GetNonce))
	router.GET("/balances/:address", middleware.PrometheusMiddleware(duration.WithLabelValues("GetBalance"), handler.GetBalance))
	router.GET("/blocks/:height", middleware.PrometheusMiddleware(duration.WithLabelValues("GetBlock"), handler.GetBlock))
	router.GET("/receipts/:hashes", middleware.PrometheusMiddleware(duration.WithLabelValues("GetReceipts"), handler.GetReceipts))
	router.GET("/containers/:address/:id/:key", middleware.PrometheusMiddleware(duration.WithLabelValues("GetContainer"), handler.GetContainer))
	router.POST("/txs", middleware.PrometheusMiddleware(duration.WithLabelValues("SendTransactions"), handler.SendTransactions))
	router.POST("/config", middleware.PrometheusMiddleware(duration.WithLabelValues("UpdateConfig"), handler.UpdateConfig))
	router.POST("/connect", handler.Connect)
	router.GET("/connect", handler.Connect)

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), router)
}
