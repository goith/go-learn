## 依赖管理
```bash

# 安装dep
# 首先设置好GOPATH，保证$GOPATH/bin在环境变量PATH中，然后执行如下命令：
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh 

# 然后执行如下命令确认安装：
dep -h

# 首次使用，需要在项目代码根目录下执行
dep init

# 下载依赖包
dep ensure -add github.com/foo/bar github.com/baz/quux

#  也可以执行如下命令，下载全部依赖包，但会对已下载了且修改过的包进行还原
dep ensure

# 更新所有依赖项的锁定版本
dep ensure -update

# 检查导入、Gopkg.toml和Gopkg.lock是否同步
dep check

# 查看依赖状态
dep status 

```
