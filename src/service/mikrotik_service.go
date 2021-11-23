package service

import (
	"errors"
	"fmt"
	"github.com/go-routeros/routeros"
)

type Mikrotik struct {
	routerAddr     string
	routerLogin    string
	routerPassword string
}

func NewMikrotik(routerAddr, routerLogin, routerPassword string) (*Mikrotik, error) {
	router := &Mikrotik{
		routerAddr:     routerAddr,
		routerLogin:    routerLogin,
		routerPassword: routerPassword,
	}

	conn, err := router.dial()
	if err != nil {
		return nil, err
	}
	conn.Close()

	return router, nil
}

func (rc *Mikrotik) AddIP(ip string, comment string) error {
	conn, err := rc.dial()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Run("/ip/firewall/address-list/add", "=list=WL", "=address="+ip, "=comment=\""+comment+"\"")
	if err != nil {
		return err
	}

	return nil
}

func (rc *Mikrotik) RemoveIP(ip string) error {
	conn, err := rc.dial()
	if err != nil {
		return err
	}
	defer conn.Close()

	findIP, err := conn.Run("/ip/firewall/address-list/print", fmt.Sprintf("address=%s", ip), ".proplist=.id")
	if err != nil {
		return err
	}

	if len(findIP.Re) <= 0 {
		return errors.New("IP is not found")
	}

	ipID, ok := findIP.Re[0].Map[".id"]
	if !ok {
		return errors.New("IP is not found")
	}

	_, err = conn.Run("/ip/firewall/address-list/remove", "=.id="+ipID)
	if err != nil {
		return err
	}

	return nil

}

func (rc *Mikrotik) dial() (*routeros.Client, error) {
	return routeros.Dial(rc.routerAddr, rc.routerLogin, rc.routerPassword)
}
