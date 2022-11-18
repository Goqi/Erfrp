/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:2022/11/17 16:36
*/

package dscrypto

// 对服务器IP进行隐藏需要修改此处的AESKey和AESencryptCode。
// 同时需要对frpc.ini中的server_addr进行修改，修改成AESencryptCode。
// server_addr支持正常的ip和加密之后的ip，2种形式。
var (
	VpsIP          = "192.168.1.22"
	AESKey         = "9d9d14b5f6650726afe17e1af4052632" //Erfrp
	AESencryptCode = "J6X+PfMnVldSaM1tpjaNKw=="
	//AESencryptCode = "2HrQDAPV5JgjckfYkO9u4g=="
)
