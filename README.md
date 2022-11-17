# Erfrp-[frp](https://github.com/fatedier/frp)二开-免杀与隐藏

本项目是frp的二开项目。frp是fatedier开发的一款优秀的快速反向代理工具，可以将NAT或防火墙后面的本地服务器暴露在互联网上。但原程序对攻击队而言并不优雅，希望本项目可以为攻击队贡献完美的FRP二开项目！

目前程序和代码未发布，敬请期待！

本项目创建于2022年4月18日，最近的更新时间为2022年11月16日。

- [01-项目结构修改](https://github.com/Goqi/Erfrp#01-%E9%A1%B9%E7%9B%AE%E7%BB%93%E6%9E%84%E4%BF%AE%E6%94%B9)
- [02-项目功能修改](https://github.com/Goqi/Erfrp#02-%E9%A1%B9%E7%9B%AE%E5%8A%9F%E8%83%BD%E4%BF%AE%E6%94%B9)
- [03-静态特征修改](https://github.com/Goqi/Erfrp#03-%E9%9D%99%E6%80%81%E7%89%B9%E5%BE%81%E4%BF%AE%E6%94%B9)
- [04-流量特征修改](https://github.com/Goqi/Erfrp#04-%E6%B5%81%E9%87%8F%E7%89%B9%E5%BE%81%E4%BF%AE%E6%94%B9)
- [05-敏感信息隐藏](https://github.com/Goqi/Erfrp#05-%E6%95%8F%E6%84%9F%E4%BF%A1%E6%81%AF%E9%9A%90%E8%97%8F)
- [06-参考项目资源](https://github.com/Goqi/Erfrp#06-%E5%8F%82%E8%80%83%E9%A1%B9%E7%9B%AE%E8%B5%84%E6%BA%90)

## 01-项目结构修改

```
│  frpc.go
│  frps.go
├─cmd
│  ├─frpc
│  │      http.go
│  │      https.go
│  │      reload.go
│  │      root.go
│  │      status.go
│  │      stcp.go
│  │      sudp.go
│  │      tcp.go
│  │      tcpmux.go
│  │      udp.go
│  │      verify.go
│  │      xtcp.go
│  │      
│  └─frps
│          root.go
│          verify.go
│          
├─pkg
│  ├─assets
│  │  │  assets.go
│  │  │  
│  │  ├─frpc
│  │  │  │  embed.go
│  │  │  └─static
│  │  │          
│  │  └─frps
│  │      │  embed.go
│  │      └─static
│  │              
│  ├─auth
│  │      auth.go
│  │      oidc.go
│  │      token.go
│  │      
│  ├─client
│  │  │  admin.go
│  │  │  admin_api.go
│  │  │  control.go
│  │  │  service.go
│  │  │  visitor.go
│  │  │  visitor_manager.go
│  │  │  
│  │  ├─event
│  │  │      event.go
│  │  │      
│  │  ├─health
│  │  │      health.go
│  │  │      
│  │  └─proxy
│  │          proxy.go
│  │          proxy_manager.go
│  │          proxy_wrapper.go
│  │      
│  ├─config
│  │      client.go
│  │      client_test.go
│  │      DefaultiniBytefrpc.go
│  │      DefaultiniBytefrps.go
│  │      parse.go
│  │      proxy.go
│  │      proxy_test.go
│  │      README.md
│  │      server.go
│  │      server_test.go
│  │      types.go
│  │      types_test.go
│  │      utils.go
│  │      value.go
│  │      visitor.go
│  │      visitor_test.go
│  │      
│  ├─consts
│  │      consts.go
│  │      
│  ├─crypto
│  │      aes.go
│  │      aes1.go
│  │      aes2.go
│  │      des.go
│  │      md5.go
│  │      rsa.go
│  │      rsa_private_key.pem
│  │      rsa_public_key.pem
│  │      sha1.go
│  │      var.go
│  │      
│  ├─errors
│  │      errors.go
│  │      
│  ├─metrics
│  │  │  metrics.go
│  │  │  
│  │  ├─aggregate
│  │  │      server.go
│  │  │      
│  │  ├─mem
│  │  │      server.go
│  │  │      types.go
│  │  │      
│  │  └─prometheus
│  │          server.go
│  │          
│  ├─msg
│  │      ctl.go
│  │      msg.go
│  │      
│  ├─nathole
│  │      nathole.go
│  │      
│  ├─plugin
│  │  ├─client
│  │  │      http2https.go
│  │  │      https2http.go
│  │  │      https2https.go
│  │  │      http_proxy.go
│  │  │      plugin.go
│  │  │      socks5.go
│  │  │      static_file.go
│  │  │      unix_domain_socket.go
│  │  │      
│  │  └─server
│  │          http.go
│  │          manager.go
│  │          plugin.go
│  │          tracer.go
│  │          types.go
│  │          
│  ├─proto
│  │  └─udp
│  │          udp.go
│  │          udp_test.go
│  │          
│  ├─server
│  │  │  control.go
│  │  │  dashboard.go
│  │  │  dashboard_api.go
│  │  │  service.go
│  │  │  
│  │  ├─controller
│  │  │      resource.go
│  │  │      
│  │  ├─group
│  │  │      group.go
│  │  │      http.go
│  │  │      tcp.go
│  │  │      tcpmux.go
│  │  │      
│  │  ├─metrics
│  │  │      metrics.go
│  │  │      
│  │  ├─ports
│  │  │      ports.go
│  │  │      
│  │  ├─proxy
│  │  │      http.go
│  │  │      https.go
│  │  │      proxy.go
│  │  │      stcp.go
│  │  │      sudp.go
│  │  │      tcp.go
│  │  │      tcpmux.go
│  │  │      udp.go
│  │  │      xtcp.go
│  │  │      
│  │  └─visitor
│  │          visitor.go
│  │          
│  ├─transport
│  │      tls.go
│  │      
│  ├─util
│  │  ├─limit
│  │  │      reader.go
│  │  │      writer.go
│  │  │      
│  │  ├─log
│  │  │      log.go
│  │  │      
│  │  ├─metric
│  │  │      counter.go
│  │  │      counter_test.go
│  │  │      date_counter.go
│  │  │      date_counter_test.go
│  │  │      metrics.go
│  │  │      
│  │  ├─net
│  │  │      conn.go
│  │  │      dial.go
│  │  │      http.go
│  │  │      kcp.go
│  │  │      listener.go
│  │  │      tls.go
│  │  │      udp.go
│  │  │      websocket.go
│  │  │      
│  │  ├─tcpmux
│  │  │      httpconnect.go
│  │  │      
│  │  ├─util
│  │  │      http.go
│  │  │      util.go
│  │  │      util_test.go
│  │  │      
│  │  ├─version
│  │  │      version.go
│  │  │      version_test.go
│  │  │      
│  │  ├─vhost
│  │  │      http.go
│  │  │      https.go
│  │  │      https_test.go
│  │  │      resource.go
│  │  │      router.go
│  │  │      vhost.go
│  │  │      
│  │  └─xlog
│  │          ctx.go
│  │          xlog.go
│  │          
│  └─web
│      ├─frpc
│      │  └─src            
│      └─frps
│          └─src
```

## 02-项目功能修改

- [x] 程序运行判断是否存在frpc.ini或frps.ini文件，不存在则自动创建。
- [ ] 加入命令执行模块
- [ ] 全部的参数都从ini文件获取？or 全部的参数都写到go文件中？

## 03-静态特征修改

- [ ] 去除日志打印相关内容

## 04-流量特征修改

- [x] 0x17特征修改

## 05-敏感信息隐藏

- [ ] 服务端IP地址加密
- [x] [程序运行后删除配置文件](https://github.com/Goqi/Erfrp/blob/main/pkg/cmd/frpc/root.go)：例子：frpc.exe --delini
- [x] [远程加载配置文件](https://github.com/Goqi/Erfrp/blob/main/pkg/config/value.go)：例子：frpc.exe -c http://127.0.0.1/frpc.ini

## 06-参考项目资源

- https://github.com/atsud0/frp-modify
- https://github.com/OrangeWatermelon/frp_cmd
- https://github.com/baibaicloud/frp
- https://github.com/seventeenman/Forest
- https://github.com/XiaoMingXD/SakuraFrp
- https://github.com/uknowsec/frpModify
- https://github.com/NS-Sp4ce/Frp_modify
- https://github.com/g0h3aler/frp-owner
- https://www.nctry.com/2358.html
- https://www.anquanke.com/post/id/231424
- https://www.anquanke.com/post/id/231685
- https://cn-sec.com/archives/1369447.html
- https://blog.csdn.net/weixin_38989540/article/details/105924886
- https://blog.csdn.net/Workers77/article/details/124802570
