/*
 * Apache License 2.0
 *
 * Copyright (c) 2022, Austin Zhai
 * All rights reserved.
 */
package proxy

import (
	"strings"
	"time"

	"github.com/jumboframes/conduit/pkg/log"
	nfw "github.com/jumboframes/conduit/pkg/nf_wrapper"
)

func (client *Client) setProc() error {
	err := client.initProc()
	if err != nil {
		return err
	}
	go func() {
		tick := time.NewTicker(time.Duration(client.conf.Client.Proxy.CheckTime) * time.Second)
		for {
			select {
			case <-tick.C:
				err = client.initProc()
				if err != nil {
					log.Errorf("Client::setProc | init proc err: %s", err)
				}
			case <-client.quit:
				return
			}
		}
	}()
	return nil

}

func (client *Client) initProc() error {
	infoO, infoE, err := nfw.Cmd("echo", "1", ">", "/proc/sys/net/ipv4/conf/all/route_localnet")
	if err != nil {
		log.Errorf("client::initProc | enable route local net err: %s, stdout: %s, stderr: %s",
			err, infoO, strings.TrimSuffix(string(infoE), "\n"))
		return err
	}
	return nil
}
