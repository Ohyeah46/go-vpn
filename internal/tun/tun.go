// internal/tun/tun.go
package tun

import (
	"github.com/songgao/water"
	"log"
)

func CreateTUN(name string) (*water.Interface, error) {
	cfg := water.Config{
		DeviceType: water.TUN,
	}
	cfg.Name = name
	ifce, err := water.New(cfg)
	if err != nil {
		return nil, err
	}
	log.Println("Created TUN interface:", ifce.Name())
	return ifce, nil
}
