package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/3scale/3scale-istio-adapter/pkg/threescale"
	"github.com/3scale/3scale-istio-adapter/pkg/threescale/metrics"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
	istiolog "istio.io/pkg/log"
)

var version string

func init() {
	viper.SetEnvPrefix("threescale")
	viper.BindEnv("log_level")
	viper.BindEnv("log_json")
	viper.BindEnv("log_grpc")
	viper.BindEnv("listen_addr")
	viper.BindEnv("report_metrics")
	viper.BindEnv("metrics_port")

	viper.BindEnv("cache_ttl_seconds")
	viper.BindEnv("cache_refresh_seconds")
	viper.BindEnv("cache_entries_max")

	viper.BindEnv("client_timeout_seconds")
	viper.BindEnv("allow_insecure_conn")

	viper.BindEnv("grpc_conn_max_seconds")

	options := istiolog.DefaultOptions()

	if viper.IsSet("log_level") {
		loglevel := viper.GetString("log_level")
		parsedLogLevel, err := stringToLogLevel(loglevel)

		if err != nil {
			fmt.Printf("THREESCALE_LOG_LEVEL is not valid, expected: debug,info,warn,error or none. Got: %v\n", loglevel)
			os.Exit(1)
		}

		options.SetOutputLevel(istiolog.DefaultScopeName, parsedLogLevel)
	}

	if viper.IsSet("log_json") {
		options.JSONEncoding = viper.GetBool("log_json")
	}

	if !viper.IsSet("log_grpc") || !viper.GetBool("log_grpc") {
		options.LogGrpc = false
		grpclog.SetLogger(log.New(ioutil.Discard, "", 0))
	}

	istiolog.Configure(options)
	istiolog.Infof("Logging started")

}

func stringToLogLevel(loglevel string) (istiolog.Level, error) {

	stringToLevel := map[string]istiolog.Level{
		"debug": istiolog.DebugLevel,
		"info":  istiolog.InfoLevel,
		"warn":  istiolog.WarnLevel,
		"error": istiolog.ErrorLevel,
		"none":  istiolog.NoneLevel,
	}

	if val, ok := stringToLevel[strings.ToLower(loglevel)]; ok {
		return val, nil
	}

	return istiolog.InfoLevel, errors.New("invalid log_level")
}

func parseMetricsConfig() *metrics.Reporter {
	if !viper.IsSet("report_metrics") || !viper.GetBool("report_metrics") {
		return nil
	}

	var port int
	if viper.IsSet("metrics_port") {
		port = viper.GetInt("metrics_port")
	} else {
		port = 8080
	}

	return metrics.NewMetricsReporter(true, port)
}

func parseClientConfig() *http.Client {
	c := &http.Client{
		// Setting some sensible default here for http timeouts
		Timeout: time.Duration(time.Second * 10),
	}

	if viper.IsSet("client_timeout_seconds") {
		c.Timeout = time.Duration(viper.GetInt("client_timeout_seconds")) * time.Second
	}

	if viper.IsSet("allow_insecure_conn") {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: viper.GetBool("allow_insecure_conn")},
		}
		c.Transport = tr
	}

	return c
}

func cacheConfigBuilder() *threescale.ProxyConfigCache {
	cacheTTL := threescale.DefaultCacheTTL
	cacheRefreshInterval := threescale.DefaultCacheRefreshBuffer
	cacheEntriesMax := threescale.DefaultCacheLimit
	cacheUpdateRetries := threescale.DefaultCacheUpdateRetries

	if viper.IsSet("cache_ttl_seconds") {
		cacheTTL = time.Duration(viper.GetInt("cache_ttl_seconds")) * time.Second
	}

	if viper.IsSet("cache_refresh_seconds") {
		cacheRefreshInterval = time.Duration(viper.GetInt("cache_refresh_seconds")) * time.Second
	}

	if viper.IsSet("cache_entries_max") {
		cacheEntriesMax = viper.GetInt("cache_entries_max")
	}

	if viper.IsSet("cache_refresh_retries") {
		cacheUpdateRetries = viper.GetInt("cache_refresh_retries")
	}

	return threescale.NewProxyConfigCache(cacheTTL, cacheRefreshInterval, cacheUpdateRetries, cacheEntriesMax)
}

func main() {
	var addr string

	if viper.IsSet("listen_addr") {
		addr = viper.GetString("listen_addr")
	} else {
		addr = "0"
	}

	proxyCache := cacheConfigBuilder()

	grpcKeepAliveFor := time.Minute
	if viper.IsSet("grpc_conn_max_seconds") {
		grpcKeepAliveFor = time.Second * time.Duration(viper.GetInt("grpc_conn_max_seconds"))
	}
	adapterConfig := threescale.NewAdapterConfig(proxyCache, parseMetricsConfig(), grpcKeepAliveFor)

	s, err := threescale.NewThreescale(addr, parseClientConfig(), adapterConfig)

	if err != nil {
		istiolog.Errorf("Unable to start sever: %v", err)
		os.Exit(1)
	}

	if version == "" {
		version = "undefined"
	}
	istiolog.Infof("Started server version %s", version)

	proxyCache.StartRefreshWorker()
	if err != nil {
		istiolog.Errorf("error while starting cache refresh worker %s", err.Error())
	}

	shutdown := make(chan error, 1)
	go func() {
		s.Run(shutdown)
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM)

	go func() {
		_ = <-sigC
		fmt.Println("SIGTERM received. Attempting graceful shutdown")
		err := s.Close()
		if err != nil {
			fmt.Println("Error calling graceful shutdown")
			os.Exit(1)
		}
		return
	}()

	_ = <-shutdown
}
