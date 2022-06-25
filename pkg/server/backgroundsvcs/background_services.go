package backgroundsvcs

import (
	"github.com/taidevops/grafana/pkg/registry"
)

func ProvideBackgroundServiceRegistry() *BackgroundServiceRegistry {

}

// BackgroundServiceRegistry provides background services.
type BackgroundServiceRegistry struct {
	Services []registry.BackgroundService
}

func NewBackgroundServiceRegistry(services ...registry.BackgroundService) *BackgroundServiceRegistry {
	return &BackgroundServiceRegistry{services}
}

func (r *BackgroundServiceRegistry) GetServices() []registry.BackgroundService {
	return r.Services
}
