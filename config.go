package main

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type config struct {
	configConnDial
	AllowService map[string]struct{} // 允许哪些服务访问本服务
	sync.RWMutex
}

var CONFIG *config

func init() {
	cc := &config{
		configConnDial: configConnDial{
			KeepAlive:        time.Second / 2,
			KeepAliveTimeout: time.Second,
			DialTimeOut:      time.Second,
			LeaseTimeOut:     1,
		},
		AllowService: map[string]struct{}{
			"*":     struct{}{},
			"DEBUG": struct{}{},
		},
	}

	dat, err := os.ReadFile("./sidecar.json")
	if err == nil {
		var c = map[string]string{}
		err = json.Unmarshal(dat, &c)
		if err == nil {
			if field, ok := c["KeepAlive"]; ok {
				i, err := strconv.Atoi(field)
				if err == nil {
					cc.KeepAlive = time.Millisecond * time.Duration(i)
				}
			}
			if field, ok := c["KeepAliveTimeout"]; ok {
				i, err := strconv.Atoi(field)
				if err == nil {
					cc.KeepAliveTimeout = time.Millisecond * time.Duration(i)
				}
			}
			if field, ok := c["DialTimeOut"]; ok {
				i, err := strconv.Atoi(field)
				if err == nil {
					cc.DialTimeOut = time.Millisecond * time.Duration(i)
				}
			}
			if field, ok := c["LeaseTimeOut"]; ok {
				i, err := strconv.Atoi(field)
				if err == nil {
					cc.LeaseTimeOut = int64(i)
				}
			}
			if field, ok := c["AllowService"]; ok {
				spans := strings.Split(field, ";")
				cc.AllowService = map[string]struct{}{}
				for _, span := range spans {
					cc.AllowService[span] = struct{}{}
				}
			}
		}
	}

	CONFIG = cc
}

func ReadConfig() string {
	var as []string
	for k, _ := range CONFIG.AllowService {
		as = append(as, k)
	}
	var c = map[string]string{
		"KeepAlive":        strconv.Itoa(int(CONFIG.KeepAlive / time.Millisecond)),
		"KeepAliveTimeout": strconv.Itoa(int(CONFIG.KeepAliveTimeout / time.Millisecond)),
		"DialTimeOut":      strconv.Itoa(int(CONFIG.DialTimeOut / time.Millisecond)),
		"LeaseTimeOut":     strconv.Itoa(int(CONFIG.LeaseTimeOut)),
		"AllowService":     strings.Join(as, ";"),
		"node_id":          strconv.Itoa(int(node_id)),
	}
	s, _ := json.Marshal(c)
	return string(s)
}

func UpdateConfig(dat []byte) error {
	CONFIG.Lock()
	defer CONFIG.Unlock()
	var c = map[string]string{}
	err := json.Unmarshal(dat, &c)
	if err != nil {
		return err
	}
	if field, ok := c["KeepAlive"]; ok {
		i, err := strconv.Atoi(field)
		if err != nil {
			return errors.New("bad KeepAlive")
		} else {
			CONFIG.KeepAlive = time.Millisecond * time.Duration(i)
		}
	}
	if field, ok := c["KeepAliveTimeout"]; ok {
		i, err := strconv.Atoi(field)
		if err != nil {
			return errors.New("bad KeepAliveTimeout")
		} else {
			CONFIG.KeepAliveTimeout = time.Millisecond * time.Duration(i)
		}
	}
	if field, ok := c["DialTimeOut"]; ok {
		i, err := strconv.Atoi(field)
		if err != nil {
			return errors.New("bad DialTimeOut")
		} else {
			CONFIG.DialTimeOut = time.Millisecond * time.Duration(i)
		}
	}
	if field, ok := c["AllowService"]; ok {
		spans := strings.Split(field, ";")
		CONFIG.AllowService = map[string]struct{}{}
		for _, span := range spans {
			CONFIG.AllowService[span] = struct{}{}
		}
	}
	return nil
}
