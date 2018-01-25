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
rz
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


#多节点Fabric配置
## 1.准备环境
### 运行Fabric节点需要依赖以下工具：
#### a.Docker：用于管理Fabric镜像以及运行peer和orderer等组件
#### b.Docker-compose：用于配置Fabric容器
#### c.Fabric源码：源码提供了用于生成证书和配置channel的工具和测试代码
#### d.Go语言开发环境：源码的工具便宜依赖于Go语言
## 2.配置多节点Fabric集群
 在单节点呃e2e_cli示例中，所有节点部署在同一个docker-compose的内部网络中，通过容器的7051端口进行通信。但是在多节点的情况下，容器之间不能进行直接通讯，因此需要把容器的7051端口映射到宿主机上，通过各个宿主机的7051端口来实现节点间通信。我们在每个节点中修改 docker-compose.yaml 中的 service 定义，在不同节点只启动需要的 service。例如，在节点1中只启动peer0 的 service，在节点5中仅启动 orderer 等。
## 3.启动多节点Fabric集群
#### 在哥哥节点上配置好Fabric的启动环境后，需要依次登录到节点上通过docker-compose up的方式启动Fabric节点。由于启动环境有依赖关系，如 peer1 以 peer0 作为发现节点，因此需要先启动 peer0 再启动 peer1。
## 4.配置channel
 在Fabric中，channel代表了一个私有的广播通道，保证了消息的隔离性和私密性，它由 orderer 来管理。channel 中的成员共享该 channel 的账本，并且只有通过验证的用户才能在 channel 中进行交易，与一个 channel 相关的属性记录在该channel的初始区块中，可通过 reconfiguration 交易进行更改。channel的初始区块由 create channel 交易生成，peer 向 orderer 发送该交易时会带有的 config.tx 文件，该文件定义 channel 的相关属性。
## 发布chaincode
 chaincode 是开发人员按照特定接口编写的智能合约，通过 SDK 或者 CLI 在 Fabric 的网络上安装并且初始化后，该应用就能访问网络中的共享账本。
### chaincode的生命周期如下：
### a.Install（安装）
 chaincode 要在 Fabric 网络上运行，必须要先安装在网络中的 peer 上，安装同时注明版本号保证应用的版本控制。
### b.Instantiate（实例化）
 在 peer 上安装 chaincode 后，还需要实例化才能真正激活该 chaincode 。在实例化的过程中，chaincode 就会被编译并打包成容器镜像，然后启动运行。若 chaincode 在实例化的过程中更新了数据状态，如给某个变量赋予初始值，则该状态变化会被记录在共享账本中。每个应用只能被实例化一次，实例化可在任意一个已安装该 chaincode 的 peer 上进行。
### c.Invoke和query（调用和查询）
 chaincode 在实例化后，用户就能与它进行交互，其中 query 查询与应用相关的状态（即只读），而 invoke 则可能会改变其状态。
### d.Upgrade（升级）
 在 chaincode 添加新功能或出现 bug 需要升级时，可以通过 upgrade 交易来实现。这时需要把新的代码通过install交易安装到正在运行该 chaincode 
的 peer 上，安装时需注明比先前版本更高的版本号，接下来只需要向任意一个安装了新代码的 peer 发送 upgrade 交易就能更新 chaincode ，chaincode 在更新前的状态也会得到保留。

## 二、操作步骤
### 1、环境构建与测试
 本文中用到的宿主机环境是 Ubuntu 16.04.3 LTS，通过 Docker 容器来运行 Fabric 的节点, 版本为 v1.0 beta。因此，启动 Fabric 网络中的节点需要先安装Docker 、 Docker-compose 和 Go 语言环境，然后在网上拉取相关的 Docker镜像，再通过配置 compose 文件来启动各个节点。
### 1.1、Docker与Docker-compose安装
 与单机部署一样。
当安装完成后，可通过 docker version 命令来查看 docker 的版本信息并确认安装成功。
接下来下载并安装 docker-compose 。安装完毕后可以通过 docker-compose version 来确认安装是否成功。
### 1.2、Go安装
与单机部署一样
### 1.3、下载Fabric源码
### 1.4、Docker镜像下载
进入到 fabric/examples/e2e_cli 目录下，运行 ./download-dockerimages.sh 来下载必要镜像。镜像下载完成后，就可以通过 docker save 命令把镜像打包成压缩文件，传送到各个VM。当VM接收到压缩文件后，可以通过 docker load 来解压和导入镜像。如果有私有的容器镜像仓库registry，如 Harbor 等，也可以把镜像推送到私有registry，再从各个机器中拉取。
通过以下命令来保存所有 tag 含有 beta 标识的镜像到名字为 images 的压缩文件中：

docker save $(docker images | grep beta | awk {‘print $1’} ) -o images

生成 images 文件后，就可以通过scp把它拷贝到还没有镜像的其他节点中，例如，地址为 10.112.122.6 的节点需要安装以上镜像，可以通过以下命令把images 远程拷贝到 10.112.122.6 的home目录下：

scp images root@10.112.122.6:~

然后在 10.112.112.6 这台主机的home目录上运行：

docker load -i images

等待一段时间后, 通过 docker images 命令就能查看到相关镜像的信息。

### 1.5、运行测试
进入到 fabric/example/e2e_cli 文件夹，文件结构如下：
。。
network_setup.sh 是一键测试脚本，该脚本启动了6个 docker 容器，其中有4个容器运行 peer 节点和1个容器运行 orderer 节点，它们组成一个Fabric集群。另外, 还有一个 cli 容器用于执行创建 channel 、加入 channel 、安装和执行chaincode 等操作。测试用的 chaincode 定义了两个变量，在实例化的时候给这两个变量赋予了初值，通过invoke操作可以使两个变量的值发生变化。
通过以下命令执行测试：

bash network_setup.sh up

接下来会有许多的调试信息，具体可参考 e2e_cli 目录下的 script/script.sh 文件，当终端出现以下信息时说明测试通过，所有部件工作正常：

===All GOOD, End-2-End execution completed ===

至此，环境配置工作完毕，通过 docker ps -a 命令可以查看各容器的状态。 chaincode 会在独立的容器中运行，因此会出现3个以 dev 开头的容器，它们与各自的 peer 对应，记录了 peer 对 chaincode 的操作。

### 2、创建Fabric多节点集群
### 2.1、前期准备
我们将重现 Fabric 自带的 e2e_cli 示例中的集群，不同的是要把容器分配到不同的虚拟机上，彼此之间通过网络来进行通信，网络搭建完成后则进行相关的 channel 和 chaincode 操作。

先准备5台虚拟机（VM），所有虚拟机均按照上述环境构建与测试步骤配置，当然也可安装一个虚拟机模板，然后克隆出其他虚拟机。其中4台虚拟机运行 peer 节点，另外一台运行 orderer 节点，为其他的四个节点提供order服务。
### 2.2、生成证书和config.tx
在任意VM上运行 fabric/examples/e2e_cli 目录下的 generateArtifacts.sh 脚本，可生成两个目录，它们分别为 channel-artifacts/ 和 crypto-config/，两个目录的结构分别如下:

   -channel-artifacts 
       -channel.tx
      -genesis.block
       -Org1MSPanchors.tx
       -Org2MSPanchors.tx

上述目录里的文件用于 orderer 创建 channel , 它们根据 configtx.yaml 的配置生成。

  -crypto-config
   -ordererOrganizations
   -peerOrganizations 

上述目录里面有 orderer 和 peer 的证书、私钥和以及用于通信加密的tls证书等文件，它通过 configtx.yaml 配置文件生成。
### 2.3、多节点Fabric的配置
以下各VM的工作目录为： 

$GOPATH/src/github.com/hyperledger/fabric/examples/e2e_cli

可在任意VM上运行以下命令，生成构建 Fabric 网络所需的成员证书等必要材料：

bash generateArtifacts.sh

该命令只需在某个VM上运行一次，其他VM上就不需要运行。

在运行该命令的VM中会生成 channel-artifacts 和 crypto-config 目录，需要把它们拷贝到其他VM的 e2e_cli 目录下，如果在VM中已经存在该目录，则先把目录删除。当每个VM中都有统一的 channel-artifacts 和 crypto-config 目录后接下来就开始配置 compose 文件。
#### VM1的配置
##### 1.修改/etc/hosts的映射关系
因为容器内部通过域名的方式访问 orderer , 因此需要通过修改 /etc/hosts 把orderer 的域名和 ip 地址对应起来，在文件中添加:

10.112.122.69   orderer.example.com
##### 2.修改docker-compose-cli.yaml
在默认的情况下，docker-compose-cli.yaml会启动6个service（容器），它们分别为 peer0.org1.example.com、 peer1.org1.example.com、 peer0.org2.example.com、 peer1.org2.example.com、 orderer.example.com 和 cli，因为每台机器只运行与之对应的一个节点，因此需要注释掉无需启动的 service。

###### (1) 除 peer0.org1.example.com 和 cli service 外，其他 service 全部注释。
###### (2) 在 cli 的 volumes 中加入映射关系：
-./peer/:/opt/gopath/src/github.com/hyperledger/fabric/peer/
-/etc/hosts:/etc/hosts
  
###### (3) 注释 cli 中的 depends_on 和 command :
   depends_on:
      #- orderer.example.com
      - peer0.org1.example.com
      #- peer1.org1.example.com
      #- peer0.org2.example.com
      #- peer1.org2.example.com

      #command: /bin/bash -c './scripts/script.sh ${CHANNEL_NAME}; sleep $TIMEOUT'
 
之前我们把容器中的工作目录挂载到宿主机的 e2e_cli/peer 目录下, 是因为在执行 create channel 的过程中，orderer 会返回一个 mychannel.block 作为 peer 加入 channel 的依据，其他的 peer 要加入到相同的 channel 中必须先获取该 mychannel.block 文件。因此，通过挂载目录从宿主机就能方便获得 mychannel.block ，并且把它传输到其他的 VM 上。

挂载 /etc/hosts 的目的是把主机中 orderer.exmaple.com 与 IP 地址10.112.122.69 的映射关系带入容器中,目的是让 cli 能通过域名访问 orderer  。在实际环境中，建议通过配置 DNS 而不是修改 /etc/hosts 文件（下同）。
##### 3.修改base/peer-base.yaml，添加volumes:
volumes:
-/etc/hosts:/etc/hosts
 
这样 peer 容器能通过域名访问orderer了。

#### VM2配置：
##### 1.修改/etc/hosts的映射关系
peer1.org1.example.com 使用了 peer0.org1.example.com 作为它的初始化节点，因此需要在主机中还需要加入 VM1 的 ip 地址。
10.112.122.69   orderer.example.com
10.112.122.144  peer0.org1.example.com
##### 2. 修改docker-compose-cli.yaml
###### (1) 类似VM1，除 peer1.org1.example.com 和 cli service 外，其他 service 全部注释。

###### (2) 在 cli 的 volumes 中加入映射关系：
-./peer/:/opt/gopath/src/github.com/hyperledger/fabric/peer/
-/etc/hosts:/etc/hosts
 
###### (3) 注释cli中的 depends_on 和 command:
depends_on:
 #- orderer.example.com
 #- peer0.org1.example.com
 - peer1.org1.example.com
 #- peer0.org2.example.com
 #- peer1.org2.example.com
          
#command:/bin/bash -c './scripts/script.sh ${CHANNEL_NAME}; sleep $TIMEOUT'
     
###### (4) 修改cli中的环境变量
CORE_PEER_ADDRESS=peer1.org1.example.com:7051
 
##### 3. 修改base/peer-base.yaml，同VM1的修改。
 
#### VM3配置：

##### 1. 修改 /etc/hosts 的映射关系
 10.112.122.69      orderer.example.com
 
##### 2. 修改docker-compose-cli.yaml
###### (1) VM3 上运行 peer2 节点，因此除 peer0.org2.example.com 和 cli service 外,其他 service 全部注释。

###### (2) 在cli的 volumes 中加入映射关系：
- ./peer/:/opt/gopath/src/github.com/hyperledger/fabric/peer/
-/etc/hosts:/etc/hosts
 
###### (3) 注释cli中的 depends_on 和 command :
depends_on:
 #- orderer.example.com
 #- peer0.org1.example.com
 #- peer1.org1.example.com
 - peer0.org2.example.com
 #- peer1.org2.example.com

#command:/bin/bash -c './scripts/script.sh ${CHANNEL_NAME}; sleep $TIMEOUT'
         
###### (4) 修改cli中的环境变量
CORE_PEER_LOCALMSPID="Org2MSP"
CORE_PEER_ADDRESS=peer0.org2.example.com:7051
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
 
##### 3. 修改base/peer-base.yaml，同VM1。
 
#### VM4配置：

##### 1. 修改/etc/hosts的映射关系
peer1.org2.example.com 使用了 peer0.org2.example.com 作为它的初始化节点，因此需要在映射关系中加入 VM3 的 ip 地址
10.112.122.69       orderer.example.com
10.112.122.12       peer0.org2.example.com
 
##### 2. 修改docker-compose-cli.yaml
###### (1) VM4运行peer3，因此除peer1.org2.example.com和cliservice 外,其他service全部注释

###### (2) 在cli的volumes中加入映射关系：
-./peer/:/opt/gopath/src/github.com/hyperledger/fabric/peer/
-/etc/hosts:/etc/hosts
 
###### (3) 修改cli中的 depends_on 和 command:
depends_on:
  - peer1.org2.example.com
#command:/bin/bash -c './scripts/script.sh ${CHANNEL_NAME}; sleep $TIMEOUT'

###### (4) 修改cli中的环境变量
CORE_PEER_LOCALMSPID="Org2MSP"
CORE_PEER_ADDRESS=peer1.org2.example.com:7051
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
 
##### 3. 修改base/peer-base.yaml，同VM1。
 
#### VM5配置如下：
##### 1. 修改docker-compose-cli.yaml
除orderer外的其他 service 全部注释，即只启动 orderer 。
 
### 2.4 启动多节点Fabric集群

#### 1.启动orderer
进入到 VM5 的 fabric/examples/e2e_cli 目录下，运行
docker-compose -f docker-compose-cli.yaml up -d

此时终端会出现大量记录，当出现Beginning to service requests时，orderer启动完成。有了 orderer 之后，就可以通过它来管理 channel 。             
 
#### 2.启动 org1的第一个节点 peer0 ，即 peer0.org1.example.com

进入到 VM1 的 fabric/examples/e2e_cli 目录下，运行
docker-compose -f docker-compose-cli.yaml up -d

此时通过docker ps -a 命令可以看到成功启动了 peer0.org1.example.com 和 cli 两个容器。

接下来实现创建 channel 、加入 channel 和安装 chanicode 。首先进入到cli容器内：
docker exec -it cli bash
 
cli 与 orderer 之间的通讯使用 tls 加密，设置环境变量 ORDERER_CA 以作建立握手的凭证：
$ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/cacerts/ca.example.com-cert.pem

注：以下所有涉及到 ORDERER_CA 环境变量的命令都需预先给该变量赋值。

进入到 cli 容器后会自动跳转到 /opt/gopath/src/github.com/hyperledger/fabric/peer 目录，即工作目录，通过compose文件的配置，该目录映射为宿主机的 /e2e_cli/peer 。

在工作目录下输入以下命令，创建名为 mychannel 的 channel ：
peer channel create -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/channel.tx --tls --cafile $ORDERER_CA

channel 创建成功后，会在当前目录下生成mychannel.block文件。每个peer 在向 orderer 发送 join channel 交易的时候，需要提供这个文件才能加入到 mychannel 中，因此运行在其他 VM 上的 peer 需要得到 mychannel.block 文件来加入到 mychannel 中。由于之前的文件映射关系， mychannel.block 文件可在宿主机的 e2e_cli/peer 目录下获取，这时可以通过宿主机把 mychannel.block 拷贝到 VM2, VM3, VM4的 e2e_cli/peer 目录下。

把 peer0.org1.example.com 加入到 mychannel 中：
peer channel join -b mychannel.block

更新 mychannel 中 org1 的 anchor peer 的信息：
peer channel update -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile $ORDERER_CA            

安装 chaincode 示例 chaincode_example02 到 peer0.org1.example.com 中：
peer chaincode install -nmycc -v 1.0 -p \
github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02

这时候 chaincode 代码已经安装到了 peer0 节点上，但并未实例化运行。接下来先配置好其他节点。
 
#### 3. 启动 org1 的第二个节点 peer1，即 peer1.org1.example.com

进入到VM2的 fabric/examples/e2e_cli 目录下，运行
docker-compose -f docker-compose-cli.yaml up -d

进入到 cli 容器内部：
docker exec -it cli bash

由于前面已经把 mychannel.block 拷贝到了 VM2 的 e2e_cli/peer 目录下，因此 mychannel.block 可通过容器内的 /opt/gopath/src/github.com/hyperledger/fabric/peer 目录访问。

把 peer1.org1.example.com 加入到 mychannel 中：
peer channel join -b mychannel.block
             
安装 chaincode_example02 到 peer1.org1.example.com 中：
peer chaincode install -nmycc -v 1.0 –p \
github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02
 
#### 4. 启动 org2 的第一个节点 peer2，即 peer0.org2.example.com

进入到 VM3 的 fabric/examples/e2e_cli 目录下，运行
docker-compose-f docker-compose-cli.yaml up -d
 
进入到cli容器内部：
docker exec -it cli bash
 
把peer0.org2.example.com加入到mychannel中：
peer channel join -b mychannel.block

更新 mychannel 中 org2 的 anchor peer 的信息：
peer channel update -oorderer.example.com:7050 -c mychannel -f ./channel-artifacts/Org2MSPanchors.tx --tls --cafile $ORDERER_CA   
 
安装 chaincode_example02 到 peer0.org2.example.com 中：
peer chaincode install -nmycc -v 1.0 -p \
github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02
 
#### 5. 启动org2的第二个节点 peer3 ，即启动 peer1.org2.example.com

进入到 VM4 的 fabric/examples/e2e_cli 目录下，运行
docker-compose-f docker-compose-cli.yaml up -d

首先进入到cli容器内部：
docker exec -it cli bash
 
把 peer1.org2.example.com 加入到 mychannel 中：
peer channel join -b mychannel.block
             
安装 chaincode_example02 到 peer1.org2.example.com 中：

peer chaincode install -nmycc -v 1.0 -p \
github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02

### 2.5 运行chaincode

通过前面的步骤，整个 多节点Fabric 网络已经运行起来了，每个peer都加入到了标识为 mychannel 的 channel 中，并且都安装了一个简单的 chaincode (该 chaincode 在安装时被标识为 mycc ) 。下面步骤运行和维护 chaincode。

#### 1. 实例化chaincode

chaincode 的实例化可在任意 peer 上进行，并且 chaincode 只能被实例化一次，下面以在 peer0.org2.example.com 上实例化 chaincode 为例。

首先登录VM3并进入到cli容器内部运行：
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile $ORDERER_CA -C mychannel-nmycc -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "OR     ('Org1MSP.member','Org2MSP.member')"

这时候会构建一个新的容器来运行chaincode，通过docker ps -a 命令可以看到新容器：

dev-peer0.org2.example.com-mycc-1.0
 
上述实例化中，我们对两个变量‘a’和‘b’分别赋予初值100和200，通过 channel 它们的值被同步到了其他peer的账本上，即使其他peer还没有构建运行 chaincode 的容器。
 
#### 2.    执行 chaincode 的 query 交易

由于 chaincode 已经被 peer0.org2.example.com 实例化了，因此其他 peer 不需要再次实例化它了，但是 chaincode 的状态（world state）却是已经记录在各个peer的账本上的。

 接下来我们在peer0.org1.example.com上查看chaincode的状态，登录到VM1上并进入cli容器内部执行：
peer chaincode query -C mychannel -nmycc -c '{"Args":["query","a"]}'

上面的命令查看 mycc 中变量 a 的值，由于在 peer 跟 chaincode 发生互动之前还不存在运行 chaincode 的容器，因此第一次交互的时候需要先构建运行 chaincode 的容器，等待一段时间后返回结果：100 。

此时通过 docker ps -a 命令能看到新容器：

dev-peer0.org1.example.com-mycc-1.0

该值与实例化时的赋值一致，说明 peer0.org1 和 peer0.org2 两个 peer 可以相互通信。
 
#### 3. 执行chaincode的invoke交易

接下来，我们执行一个 invoke 交易，使得变量 a 向变量 b 转帐 20，得到最终值为["a":"80","b":"220"]。

登录到VM2并进入到cli容器中中通过以下命令查询mycc的状态：
peer chaincode query -C mychannel -n mycc -c '{"Args":["query","a"]}'

稍作等待后返回结果为100，下面执行 invoke 交易，改变 a 的值为 80 ：
peer chaincode invoke -oorderer.example.com:7050  --tls --cafile $ORDERER_CA -C mychannel -n mycc -c '{"Args":["invoke","a","b","20"]}'
 
#### 4. 再次执行 chaincode 的 query 交易

在peer1.org1.example.com 上重复以上查看 chaincode 的步骤，得到返回结果为 80 ，说明测试通过，至此，Fabric网络构建完毕，各个部件工作正常。

### 2.6  更新chaincode

通过 channel upgrade 命令可以使得 chaincode 更新到最新的版本，而低版本 chaincode 将不能再使用。

登录到VM1的 cli 容器中再次安装 chaincode_example02 ，但赋予它更高的版本号 2.0：
peer chaincode install -n mycc -v 2.0 -p \
github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02

在 VM1 的 cli 容器升级 chaincode ,添加两个变量 ‘c’和‘d’:

peer chaincode upgrade -o  orderer.example.com:7050 --tls --cafile $ORDERER_CA -n mycc -v 2.0 -c '{"Args":["init","c", "10","d","20"]}'

等待一段时间后，可以通过docker ps -a 来查看新容器构建的容器,该容器的名称为：
dev-peer0.org1.example.com-mycc-2.0

通过以下命令查询c的变量：
peer chaincode -n mycc -C mychannel -v 2.0 -c '{"Args":["query","c"]}'
返回结果为10

再次查询a的变量：
 peer chaincode -n mycc -C mychannel -v 2.0 -c'{"Args":["query","a"]}'

返回结果为80，说明更新 chaincode 成功。

这时候对账本的修改会通过 orderer 同步到其他 peer 上，但是在其他 peer 上将无法查看或更改 chaincode 的状态，因为它们还在使用旧版的 chaincode ，所以其他 pee r要想正常访问还需再次安装 chaincode ，并且设置相同的版本号 ( chaincode 代码没发生改变，只是安装时版本号更新为 2.0 )，命令如下：

peerchaincode install -n mycc -v 2.0 –p \
github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02

##全文下载：http://8btc.com/doc-view-1376.html