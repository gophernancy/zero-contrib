package nacos

import "github.com/nacos-group/nacos-sdk-go/common/constant"

const (
	allEths  = "0.0.0.0"
	envPodIP = "POD_IP"
)

// options
type Options struct {
	ListenOn    string   `json:",optional"`
	ServiceName string   `json:",optional"`
	Prefix      string   `json:",optional"`
	Weight      float64   `json:",optional"`
	Cluster     string   `json:",optional"`
	Group       string   `json:",optional"`
	Metadata    map[string]string   `json:",optional"`

	ServerConfig []constant.ServerConfig   `json:",optional"`
	ClientConfig *constant.ClientConfig   `json:",optional"`
}

type Option func(*Options)

func NewNacosConfig(serviceName, listenOn string, sc []constant.ServerConfig, cc *constant.ClientConfig, opts ...Option) *Options {
	options := &Options{
		ServiceName:  serviceName,
		ListenOn:     listenOn,
		ServerConfig: sc,
		ClientConfig: cc,
		Weight:       1.0,
		Metadata:     make(map[string]string),
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithPrefix(prefix string) Option {
	return func(o *Options) {
		o.Prefix = prefix
	}
}

func WithWeight(weight float64) Option {
	return func(o *Options) {
		o.Weight = weight
	}
}

func WithCluster(cluster string) Option {
	return func(o *Options) {
		o.Cluster = cluster
	}
}

func WithGroup(group string) Option {
	return func(o *Options) {
		o.Group = group
	}
}

func WithMetadata(metadata map[string]string) Option {
	return func(o *Options) {
		o.Metadata = metadata
	}
}
