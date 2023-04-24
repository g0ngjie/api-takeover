# API Takeover

- [x] 修改 **http(s)** 响应数据
- [x] 注入 **vConsole, eruda, mdebug**
- [ ] 操作面板
- [ ] 数据抓取、缓存

## 配置

首次执行会在当前目录下生成 **config**文件夹、**config.yaml** 配置文件和 **install_ca.crt** 证书

### install_ca 证书

用来对 **HTTPS** 目标进行操作的时候，需要安装到系统根证书下，windows 叫 **受信任的根证书颁发机构**

### config.yaml

用来配置需要注入调试工具的 **domains**

### config 文件夹

- 第一层存放需要修改响应的域名文件夹
- 第二层存放以**path**命名的接口文件，并以**txt**结尾
  - **txt** 文件保存需要修改的结果
