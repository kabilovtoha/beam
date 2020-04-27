package main

import (
	"beam.mw/service-balancer/services"
	"fmt"
	"log"
	"runtime"
)

func NewWalletServices () (* services.Services, error)  {
	log.Printf("initializing wallet services, CPU count %v, service count %v", runtime.NumCPU(), config.WalletServiceCnt)

	var cliOptions []string
	if len(config.AllowedOrigin) > 0 {
		cliOptions = append(cliOptions, "--allowed_origin")
		cliOptions = append(cliOptions, fmt.Sprintf(`"%s"`, config.AllowedOrigin))
	}

	cfg := services.Config {
		BeamNodeAddress:  config.BeamNodeAddress,
		ServiceExePath:   config.WalletServicePath,
		StartTimeout:     config.ServiceLaunchTimeout,
		HeartbeatTimeout: config.ServiceHeartbeatTimeout,
		AliveTimeout:     config.ServiceAliveTimeout,
		FirstPort:        config.WalletServiceFirstPort,
		LastPort:         config.WalletServiceLastPort,
		Debug:            config.Debug,
		NoisyLogs:        config.NoisyLogs,
		CliOptions:       cliOptions,
	}

	return services.NewServices(&cfg, config.WalletServiceCnt, "service")
}

func NewBbsServices () (* services.Services, error) {
	var svcsCnt = 1
	log.Printf("initializing wallet services, CPU count %v, service count %v", runtime.NumCPU(), svcsCnt)
	cfg := services.Config{
		BeamNodeAddress:  config.BeamNodeAddress,
		ServiceExePath:   config.BbsMonitorPath,
		StartTimeout:     config.ServiceLaunchTimeout,
		HeartbeatTimeout: config.ServiceHeartbeatTimeout,
		AliveTimeout:     config.ServiceAliveTimeout,
		FirstPort:        config.BbsMonitorFirstPort,
		LastPort:         config.BbsMonitorLastPort,
		Debug:            config.Debug,
		NoisyLogs:        config.NoisyLogs,
	}
	return services.NewServices(&cfg, svcsCnt, "bbs")
}
