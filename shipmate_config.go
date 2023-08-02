package go_shipmate

import (
	"encoding/base64"
	"errors"
	"os"
)

type ShipmateConfig struct {
}

func (c *ShipmateConfig) GetAccessId() (string, error) {
	accessId := os.Getenv("SHIPMATE_ACCESS_ID")

	if len(accessId) == 0 {
		return "", errors.New("the `SHIPMATE_ACCESS_ID` environment variable is not set")
	}

	return accessId, nil
}

func (c *ShipmateConfig) GetAccessKey() (string, error) {
	accessKey := os.Getenv("SHIPMATE_ACCESS_KEY")

	if len(accessKey) == 0 {
		return "", errors.New("the `SHIPMATE_ACCESS_KEY` environment variable is not set")
	}

	decodedAccessKey, _ := base64.StdEncoding.DecodeString(accessKey)

	return string(decodedAccessKey), nil
}

func (c *ShipmateConfig) GetEnvironmentId() (string, error) {
	environmentId := os.Getenv("SHIPMATE_ENVIRONMENT_ID")

	if len(environmentId) == 0 {
		return "", errors.New("the `SHIPMATE_ENVIRONMENT_ID` environment variable is not set")
	}

	return environmentId, nil
}

func (c *ShipmateConfig) GetRegionId() (string, error) {
	regionId := os.Getenv("SHIPMATE_REGION_ID")

	if len(regionId) == 0 {
		return "", errors.New("the `SHIPMATE_REGION_ID` environment variable is not set")
	}

	return regionId, nil
}
