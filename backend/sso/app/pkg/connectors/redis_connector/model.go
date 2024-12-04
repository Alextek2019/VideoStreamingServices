package redis_connector

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"time"

	Client "github.com/redis/go-redis/v9"
)

type Redis struct {
	Host               string `yaml:"Host" validate:"required"`
	Port               string `yaml:"Port" validate:"required"`
	MinIdleConns       int    `yaml:"MinIdleConns" validate:"required"`
	PoolSize           int    `yaml:"PoolSize" validate:"required"`
	PoolTimeout        int    `yaml:"PoolTimeout" validate:"required"`
	Password           string `yaml:"Password" validate:"required"`
	UseCertificates    bool   `yaml:"UseCertificates"`
	InsecureSkipVerify bool   `yaml:"InsecureSkipVerify"`
	CertificatesPaths  struct {
		Cert string `yaml:"Cert"`
		Key  string `yaml:"Key"`
		Ca   string `yaml:"Ca"`
	} `yaml:"CertificatesPaths"`
	DB int `yaml:"DB"`
}

func (c *Redis) BuildOptions() (*Client.Options, error) {
	opts := &Client.Options{}
	if c.UseCertificates {
		certs := make([]tls.Certificate, 0, 0)
		if c.CertificatesPaths.Cert != "" && c.CertificatesPaths.Key != "" {
			cert, err := tls.LoadX509KeyPair(c.CertificatesPaths.Cert, c.CertificatesPaths.Key)
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"certPath: %v, keyPath: %v",
					c.CertificatesPaths.Cert,
					c.CertificatesPaths.Key,
				)
			}
			certs = append(certs, cert)
		}
		caCert, err := os.ReadFile(c.CertificatesPaths.Ca)
		if err != nil {
			return nil, errors.Wrapf(err, "ca load path: %v", c.CertificatesPaths.Ca)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		opts = &Client.Options{
			Addr:         fmt.Sprintf("%s:%s", c.Host, c.Port),
			MinIdleConns: c.MinIdleConns,
			PoolSize:     c.PoolSize,
			PoolTimeout:  time.Duration(c.PoolTimeout) * time.Second,
			Password:     c.Password,
			DB:           c.DB,
			TLSConfig: &tls.Config{
				InsecureSkipVerify: c.InsecureSkipVerify,
				Certificates:       certs,
				RootCAs:            caCertPool,
			},
		}
	} else {
		opts = &Client.Options{
			Addr:         fmt.Sprintf("%s:%s", c.Host, c.Port),
			MinIdleConns: c.MinIdleConns,
			PoolSize:     c.PoolSize,
			PoolTimeout:  time.Duration(c.PoolTimeout) * time.Second,
			Password:     c.Password,
			DB:           c.DB,
		}
	}

	return opts, nil
}
