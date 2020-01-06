package serviceregistration

import (
	log "github.com/hashicorp/go-hclog"
)

type State struct {
	VaultVersion                                               string
	IsInitialized, IsSealed, IsDRStandby, IsPerformanceStandby bool
}

// Factory is the factory function to create a ServiceRegistration.
type Factory func(shutdownCh <-chan struct{}, config map[string]string, logger log.Logger, state *State, redirectAddr string) (ServiceRegistration, error)

// ServiceRegistration is an interface that advertises the state of Vault to a
// service discovery network.
type ServiceRegistration interface {
	// NotifyDRStandbyStateChange is used by Core to notify that this Vault
	// instance has changed its status on whether it's the leader or is
	// a standby.
	NotifyDRStandbyStateChange(isLeader bool) error

	// NotifySealedStateChange is used by Core to notify that Vault has changed
	// its Sealed status to sealed or unsealed.
	NotifySealedStateChange(isSealed bool) error

	// NotifyPerformanceStandbyStateChange is used by Core to notify that this
	// Vault instance has changed it status to performance leader or standby.
	NotifyPerformanceStandbyStateChange(isLeader bool) error

	// NotifyInitializedStateChange is used by Core to notify that the core is
	// initialized.
	NotifyInitializedStateChange(isInitialized bool) error
}
