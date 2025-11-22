# How to: 在 linux 系统中安装 go



1. 确定将要安装的版本.

    目前, 最新的版本为:

          Version: 1.25.4
             File: go1.25.4.linux-amd64.tar.gz
              URL: https://go.dev/dl/go1.25.4.linux-amd64.tar.gz
        SHA256SUM: 9fa5ffeda4170de60f67f3aa0f824e426421ba724c21e133c1e35d6159ca1bec

2. 下载新版本.

        # wget https://go.dev/dl/go1.25.4.linux-amd64.tar.gz

3. 移除以前安装的版本(如果有).

        # sudo rm -rf /usr/local/go 

4. 把新版本解压到指定位置.

        # sudo tar -C /usr/local -xzf go1.25.4.linux-amd64.tar.gz

5. 配置环境变量.

    直接在命令行终端输入(临时生效); 或写入配置文件(长期生效):

        export PATH=$PATH:/usr/local/go/bin

    注意: 各个发行版的环境变量配置文件有所不同, 应当视具体版本而定.
    以下几个文件路径供参考:

        /etc/env
        /etc/profile

        ~/.bash
        ~/.bash

6. 配置 goproxy.

        export GOPROXY=https://mirrors.aliyun.com/goproxy/

7. 最后,测试安装结果.

        # go version

        # go env
