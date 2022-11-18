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

package net

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	gnet "github.com/fatedier/golib/net"
)

var (
	//FRPTLSHeadByte = 0x17
	FRPTLSHeadByte = 0x10
)

func CheckAndEnableTLSServerConnWithTimeout(
	c net.Conn, tlsConfig *tls.Config, tlsOnly bool, timeout time.Duration,
) (out net.Conn, isTLS bool, custom bool, err error) {
	sc, r := gnet.NewSharedConnSize(c, 2)
	buf := make([]byte, 1)
	//sc, r := gnet.NewSharedConnSize(c, 4)
	//buf := make([]byte, 3)
	var n int
	_ = c.SetReadDeadline(time.Now().Add(timeout))
	n, err = r.Read(buf)
	_ = c.SetReadDeadline(time.Time{})
	if err != nil {
		return
	}

	switch {
	case n == 1 && int(buf[0]) == FRPTLSHeadByte:
		//case n == 3 && int(buf[0]) == FRPTLSHeadByte:
		out = tls.Server(c, tlsConfig)
		isTLS = true
		custom = true
	case n == 1 && int(buf[0]) == 0x16:
		out = tls.Server(sc, tlsConfig)
		isTLS = true
	default:
		if tlsOnly {
			err = fmt.Errorf("non-TLS connection received on a TlsOnly server")
			return
		}
		out = sc
	}
	return
}
