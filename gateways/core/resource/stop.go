package resource

import "github.com/argoproj/argo-events/gateways"

// StopConfig stops a configuration
func (rce *ResourceConfigExecutor) StopConfig(config *gateways.ConfigContext) error {
	if config.Active == true {
		config.StopCh <- struct{}{}
	}
	return nil
}
