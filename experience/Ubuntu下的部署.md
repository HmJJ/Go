#单节点流程
##在虚拟机下装好一个Ubuntu 16.04 系统
安装完Ubuntu之后，需要保证apt source是国内的，如果是国外的话会很慢很慢的：
sudo vim /etc/apt/sources.list     *如果没有vim的可以直接用vi或是安装一个:sudo apt-get stall vim*
打开这个apt源列表之后，如果看到里面有http://us.xxxx之类的（nova等），那么就是国外的，需要转换，在该文件下，先按esc退出插入模式，然后输入：
:%s/us./cn./g
就可以把所有的us.改为cn.了。然后输入:wq保存退出。
接着更新一下源
sudo apt-get update
然后安装ssh，这样接下来就可以用putty或者SecureCRT之类的客户端远程连接Ubuntu了
sudo apt-get stall ssh

##Go的安装
Ubuntu的apt-get虽然提供了Go的安装（其实我并不知道），但是版本比较旧，最好还是去官网下一个最新的go，命令为：（没有用）
wget https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.9.linux-amd64.tar.gz
注意不要使用apt的方式安装go，不然就给你apt的低版本的了。

直接把下载好的go压缩文件传到Ubuntu里，再解压
sudo apt-get install lrzsz
sudo tar -C /usr/local -xzf go1.9.linux-amd64.tar.gz

接下来编辑当前用户的环境变量
vim ~/.profile

添加一下内容，记得要按i进入插入模式哦
export PATH=$PATH:/usr/local/go/bin 
export GOROOT=/usr/local/go 
export GOPATH=$HOME/go 
export PATH=$PATH:$HOME/go/bin
编辑保存并退出vim后，还需要把这些环境载入一下
source ~/.profile

我们把go的目录GOPATH设置为当前用户的文件夹下，所以记得创建go文件夹
cd ~
mkdir go
校验：$GOPATH    或者    go -version

##Docker安装
我们可以使用阿里提供的镜像，安装也十分方便。通过一下命令来安装docker
### step 1: 安装必要的一些系统工具
sudo apt-get update
sudo apt-get -y install apt-transport-https ca-certificates curl software-properties-common
### step 2: 安装GPG证书
curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
### Step 3: 写入软件源信息
sudo add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
### Step 4: 更新并安装Docker-CE
sudo apt-get -y update
sudo apt-get -y install docker-ce

### 安装指定版本的Docker-CE:
### Step 1: 查找Docker-CE的版本:
### apt-cache madison docker-ce
###   docker-ce | 17.03.1~ce-0~ubuntu-xenial | http://mirrors.aliyun.com/docker-ce/linux/ubuntu xenial/stable amd64 Packages
###   docker-ce | 17.03.0~ce-0~ubuntu-xenial | http://mirrors.aliyun.com/docker-ce/linux/ubuntu xenial/stable amd64 Packages
### Step 2: 安装指定版本的Docker-CE: (VERSION例如上面的17.03.1~ce-0~ubuntu-xenial)
### sudo apt-get -y install docker-ce=[VERSION]
注意：不同的版本所用到的添加方法也是不一样的，官方文档如下：
https://cr.console.aliyun.com/#/accelerator


##Docker-Compose的安装
Docker-compose是支持通过模板脚本批量创建Docker容器的一个组件，在安装Docker-Compose之前，需要安装Python-pip,运行脚本：
sudo apt-get install python-pip

然后就是安装docker-compose，我们可以从官网（https://github.com/docker/compose/releases）下载也可以从国内的进行DaoClound下载，为了速度快接下来从DaoClound安装Docker-compose，运行脚本：
curl -L https://get.daocloud.io/docker/compose/releases/download/1.12.0/docker-compose-`uname -s`-`uname -m` > ~/docker-compose
sudo mv ~/docker-compose /usr/local/bin/docker-compose 
chmod +x /usr/local/bin/docker-compose

##Fabric源码下载
我们可以使用Git命令下载源码，首先需要建立对应的目录，然后进入该目录，Git下载源码：
mkdir -p ~/go/src/github.com/hyperledger 
cd ~/go/src/github.com/hyperledger 
git clone https://github.com/hyperledger/fabric.git
由于Fabric一直在更新，所以我们并不需要最新最新的源码，需要切换到v1.0.0版本的源码即可：
cd ~/go/src/github.com/hyperledger/fabric
git checkout v1.0.0

##Fabric Docker镜像的下载
这个跟docker hub镜像下载有些类似，官方提供了批量下载的脚本：
cd ~/go/src/github.com/hyperledger/fabric/examples/e2e_cli/
source download-dockerimages.sh -c x86_64-1.0.0 -f x86_64-1.0.0

下载完毕后，我们运行以下命令检查下载的镜像列表：
docker images
然后就能看到我们下载的东西了

##启动Fabric网络并完成ChainCode的测试
我们仍然停留在e2e_cli文件夹，这里提供了启动、关闭Fabric网络的自动化脚本。我们要启动Fabric网络，并自动运行Example02 ChainCode的测试，执行一个命令：
./network_setup.sh up
这个东西做了以下操作：
###编译生成Fabric公私钥、证书的程序，程序在目录：fabric/release/linux-amd64/bin
###基于config.yaml生成创世区块和通道相关信息，并保存在channel-artifacts文件夹。
###基于crypto-config.yaml生成公私钥和证书信息，并保存在crypto-config文件夹中。
###基于docker-compose-cli.yaml启动1Orderer+4Peer+1CLI的Fabric容器。
###在CLI启动的时候，会运行script/script.sh文件，这个脚本文件包含了创建Channel，安装Example02,运行Example02等功能。
最后运行完毕，我们可以看到一个别致的界面
如果您看到这个界面，说明我们整个Fabric网络已经通了

##手动测试一下Fabric网络
我们仍然是以现在安装好的Example02为例，在官方例子中，channel名字是mychannel，链码的名字是mycc，我们首先进入CLI，我们中心打开一个命令行窗口，输入：
docker exec -it cli bash
运行一下命令可以查询a账户的余额：
peer chaincode query -C mychannel -n mycc -c '{"Args":{"query","a"}}'
可以看到余额是90
然后，我们试一试把a账户的余额再转20元给b庄户，运行命令：
peer chaincode invoke -o orderer.example.com:7050  --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C mychannel -n mycc -c '{"Args":["invoke","a","b","20"]}'
运行结果。。。
转账完毕之后，查询一下a账户的余额，没问题的话，应该是只剩下70了。
peer chaincode query -C mychannel -n mycc -c '{"Args":["query","a"]}'

最后我们要关闭Fabric网络，首先需要运行exit命令退出cli容器。关闭Favric的命令与启动类似，命令为：
cd ~/go/src/github.com/hyperledger/fabric/examples/e2e_cli./network_setup.sh down

整个单节点流程就是这样。


#多节点流程
##