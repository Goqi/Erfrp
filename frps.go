/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:2022/11/16 8:00
*/

package main

import (
	"Erfrp/cmd/frps"
	_ "Erfrp/pkg/assets/frps"
	_ "Erfrp/pkg/metrics"
	"math/rand"
	"time"

	"github.com/fatedier/golib/crypto"
)

func main() {

	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())

	frps.Execute()
}
