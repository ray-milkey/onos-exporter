// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"flag"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/onosproject/onos-exporter/pkg/export"
)

const (
	endpoint_address = ":9861"
	endpoint_path    = "/metrics"
	exporter_mode    = "prometheus"
	e2tEndpoint      = "onos-e2t:5150"
)

var log = logging.GetLogger("main")

func main() {
	address := flag.String("address", endpoint_address, "Exporter endpoint address:port or just :port")
	path := flag.String("path", endpoint_path, "Exporter endpoint path be used to export kpis")
	mode := flag.String("mode", exporter_mode, "Exporter mode (e.g., prometheus, ...)")
	caPath := flag.String("caPath", "", "path to CA certificate")
	keyPath := flag.String("keyPath", "", "path to client private key")
	certPath := flag.String("certPath", "", "path to client certificate")
	e2tEndpoint := flag.String("e2tEndpoint", e2tEndpoint, "E2T service endpoint")

	flag.Parse()

	log.Info("Starting onos-exporter")

	cfg := export.Config{
		Address:     *address,
		Path:        *path,
		Mode:        *mode,
		CAPath:      *caPath,
		KeyPath:     *keyPath,
		CertPath:    *certPath,
		E2tEndpoint: *e2tEndpoint,
	}

	exporter := export.NewExporter(cfg)

	if err := exporter.Run(); err != nil {
		log.Fatal("ONOS exporter error %s", err)
	}
}
