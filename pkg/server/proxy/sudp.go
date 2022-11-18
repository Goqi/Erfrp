// Copyright 2019 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	"Erfrp/pkg/config"
)

type SUDPProxy struct {
	*BaseProxy
	cfg *config.SUDPProxyConf
}

func (pxy *SUDPProxy) Run() (remoteAddr string, err error) {
	xl := pxy.xl

	listener, errRet := pxy.rc.VisitorManager.Listen(pxy.GetName(), pxy.cfg.Sk)
	if errRet != nil {
		err = errRet
		return
	}
	pxy.listeners = append(pxy.listeners, listener)
	xl.Info("sudp proxy custom listen success")

	pxy.startListenHandler(pxy, HandleUserTCPConnection)
	return
}

func (pxy *SUDPProxy) GetConf() config.ProxyConf {
	return pxy.cfg
}

func (pxy *SUDPProxy) Close() {
	pxy.BaseProxy.Close()
	pxy.rc.VisitorManager.CloseListener(pxy.GetName())
}
