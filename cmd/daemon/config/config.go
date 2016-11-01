package config

import (
	"github.com/coreos/etcd/clientv3"
	"k8s.io/client-go/1.5/rest"
	"time"
)

var config Config

func Get() *Config {
	return &config
}

func GetK8S() *rest.Config {
	return &rest.Config{
		Host: config.K8S.Host,
		TLSClientConfig: rest.TLSClientConfig{
			CAFile:   config.K8S.SSL.CA,
			KeyFile:  config.K8S.SSL.Key,
			CertFile: config.K8S.SSL.Cert,
		},
	}
}

func GetEtcd3() clientv3.Config {
	return clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
}
