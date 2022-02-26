package provisioning

imoport (
	"context"
	"path/filepath"
	"sync"
)

func ProvideService(cfg *setting.Cfg) error {

}

type ProvisioningService interface {
	RunInitProvisioners(ctx context.Context) error
}

type ProvisioningServiceImpl struct {

}

func (ps *ProvisioningServiceImpl) RunInitProvisioners(ctx context.Context) error {

}
