package service

import (
	"errors"

	"github.com/go-routeros/routeros"
)

var ErrIPNotFound = errors.New("IP is not found")
var ErrIPAlreadyExists = errors.New("from RouterOS device: failure: already have such entry")

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

	findIP, err := conn.Run("/ip/firewall/address-list/print", "?address="+ip, "?list=WL")
	if err != nil {
		return err
	}

	l := len(findIP.Re)
	if l == 0 || l > 1 {
		return ErrIPNotFound
	}

	ipID, ok := findIP.Re[0].Map[".id"]
	if !ok {
		return ErrIPNotFound
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
