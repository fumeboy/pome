package main

import (
	"context"
	"fmt"

	// "errors"
	"io"
	"net"

	"syscall"

	"github.com/sirupsen/logrus"
)

// from pkg/net/parse.go
// Convert i to decimal string.
func itod(i uint) string {
	if i == 0 {
		return "0"
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; i > 0; i /= 10 {
		bp--
		b[bp] = byte(i%10) + '0'
	}

	return string(b[bp:])
}

const SO_ORIGINAL_DST = 80

func getOriginalDestination(leftConn net.Conn) (net.Conn, string, uint16) {
	tcpConn := leftConn.(*net.TCPConn)
	tcpConnFile, err := tcpConn.File()
	if err != nil {
		panic(err)
	} else {
		tcpConn.Close()
	}
	addr, err := syscall.GetsockoptIPv6Mreq(int(tcpConnFile.Fd()), syscall.IPPROTO_IP, SO_ORIGINAL_DST)
	if err != nil {
		panic(err)
	}
	// file => connection
	leftConn, err = net.FileConn(tcpConnFile)
	if err != nil {
		panic(err)
	}
	dst := itod(uint(addr.Multiaddr[4])) + "." +
		itod(uint(addr.Multiaddr[5])) + "." +
		itod(uint(addr.Multiaddr[6])) + "." +
		itod(uint(addr.Multiaddr[7]))
	dport := uint16(addr.Multiaddr[2])<<8 + uint16(addr.Multiaddr[3])
	return leftConn, dst, dport
}

func handlerInTCP(nodeActiveContext context.Context, ln net.Listener) (err error) {
	for {
		select {
		case <-nodeActiveContext.Done():
			return
		default:
		}
		var log_f = logrus.Fields{"direction": "OUT", "proxy": "tcp"}
		conn, err := ln.Accept()
		if err != nil {
			log_f["error"] = err
			logger.WithFields(log_f).Error()
			continue
		}

		conn, _, port := getOriginalDestination(conn)
		log_f["port"] = port
		logger.WithFields(log_f).Info()

		remote_tcp, err := net.Dial("tcp", P.local.addr+fmt.Sprintf(":%d", port)) //连接目标服务器
		if err != nil {
			log_f["error"] = err
			logger.WithFields(log_f).Error()
			conn.Close()
			continue
		}
		go io.Copy(remote_tcp, conn)
		go io.Copy(conn, remote_tcp)
	}
}

func handlerOutTCP(nodeActiveContext context.Context, ln net.Listener) (err error) {
	for {
		select {
		case <-nodeActiveContext.Done():
			return
		default:
		}
		conn, err := ln.Accept()
		var log_f = logrus.Fields{"direction": "OUT", "proxy": "tcp"}
		if err != nil {
			log_f["error"] = err
			logger.WithFields(log_f).Error()
			continue
		}
		conn, ip, port := getOriginalDestination(conn)
		log_f["fake_ip"] = ip
		log_f["port"] = port
		logger.WithFields(log_f).Info()
		endpoint := P.discoverer.directByFakeIP(ip)
		if endpoint == nil {
			log_f["error"] = "endpoint not found"
			logger.WithFields(log_f).Error()
			conn.Close()
			continue
		}
		log_f["real_ip"] = endpoint.addr
		remote_tcp, err := net.Dial("tcp", endpoint.addr+fmt.Sprintf(":%d", port)) //连接目标服务器
		if err != nil {
			log_f["error"] = err
			logger.WithFields(log_f).Error()
			conn.Close()
			continue
		}
		go func(remote_tcp, conn net.Conn) {
			go io.Copy(remote_tcp, conn)
			io.Copy(conn, remote_tcp)
			conn.Close()
			remote_tcp.Close()
		}(remote_tcp, conn)
		// 20 -18
		//25-19
		//26-20
		// 19-21
		//21-22
	}
}
