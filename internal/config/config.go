package config

import (
	"github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/SneaksAndData/esd-services-api-client-go/auth"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
)

type ServiceConfig struct {
	AuthConfig  auth.Config
	ClaimConfig claim.Config
	MLConfig    algorithm.Config
	SparkConfig spark.Config
}
