package kdisc

import (
	"time"

	"google.golang.org/grpc"

	"github.com/1nterdigital/aka-im-discover/pkg/common/config"
	"github.com/1nterdigital/aka-im-tools/discovery"
	"github.com/1nterdigital/aka-im-tools/discovery/etcd"
	"github.com/1nterdigital/aka-im-tools/discovery/kubernetes"
	"github.com/1nterdigital/aka-im-tools/errs"
)

const (
	ETCDCONST          = "etcd"
	KUBERNETESCONST    = "kubernetes"
	DefaultMessageSize = 20
	KB                 = 1024
	MB                 = KB * KB
	DefaultTimeout     = 10 * time.Second
)

// NewDiscoveryRegister creates a new service discovery and registry client based on the provided environment type.
func NewDiscoveryRegister(disc *config.Discovery, _ string, watchNames []string) (discovery.SvcDiscoveryRegistry, error) {
	switch disc.Enable {
	case KUBERNETESCONST:
		return kubernetes.NewKubernetesConnManager(disc.Kubernetes.Namespace,
			disc.Etcd.Address, // TODO: find what is this for
			grpc.WithDefaultCallOptions(
				grpc.MaxCallSendMsgSize(MB*DefaultMessageSize),
			),
		)
	case ETCDCONST:
		return etcd.NewSvcDiscoveryRegistry(
			disc.Etcd.RootDirectory,
			disc.Etcd.Address,
			watchNames,
			etcd.WithDialTimeout(DefaultTimeout),
			etcd.WithMaxCallSendMsgSize(DefaultMessageSize*MB),
			etcd.WithUsernameAndPassword(disc.Etcd.Username, disc.Etcd.Password))
	default:
		return nil, errs.New("unsupported discovery type", "type", disc.Enable).Wrap()
	}
}
