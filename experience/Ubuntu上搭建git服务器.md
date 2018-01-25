有了github为什么还要自己搭建git服务器呢，因为GitHub是开源的，但是很多时候我们又想把东西从不同的地方同步，又不想让别人看到，这个时候我们就可以自己来搭建一个git服务器了。因为windows对git的兼容性不好，所以建议放在ubuntu系统上。
下面就是自己安装git服务器的步骤：
# 一、安装git服务器所需软件
## 安装git-core
####   git-core是git版本控制核心软件，安装命令如下：
   sudo apt-get install git-core

####   若安装提示失败，可能是因为系统软件库的检索文件太旧了，需要先更新一下，更新命令如下：
   sudo apt-get update

## 安装openssh-server和openssh-client
#### openssh和openssh-client用于git通过ssh协议来在服务器与客户端之间传输文件，安装命令如下：
sudo apt-get install openssh-server openssh-client

## 安装python
#### 由于安装gitosis需要用到Python的一些工具，所以需要先安装setup，安装命令如下：
sudo apt-get install python-setuptools

## 初始化服务器的git用户信息
#### 安装gitosis之前需要初始化服务器的git用户信息，初始化命令如下：
git config --global user.name "[username]"
git config --flobal user.email "[email]"

##获取gitosis版本文件
#### 这个就类似于我们下载安装包，命令如下：
git clone https://github.com/res0nat0r/gitosis.git

## 安装gitosis
### 进入文件目录
cd gitosis/

### 安装gitosis
#### 这里需要使用python命令安装目录下的setup.py的python脚本进行安装，命令如下：
sudo python setup.py install

#### 到这里整个安装步骤就完成了，下面是对git进行一些基本配置。

# 二、创建git管理员账户、配置git

## 创建管理员账户
#### 创建一个账户(gitmanager)作为git服务器的管理员，可以管理其他用户的项目权限。命令为：
sudo useradd -m gitmanager
sudo passwd gitmanager

## 创建仓库存储点
#### 在/home目录下创建一个项目仓库存储点，命令为：
sudo mkdir /home/gitrepository

## 权限设置
#### 设置只有git用户拥有所有权限，其他用户没有任何权限，命令为：
sudo chown gitmanager /home/gitrepository/
sudo chmod 700 /home/gitrepository/

## 创建链接映射
#### 由于gitosis默认状态下会将仓库放在用户的repositorise目录下，例如gitmanager用户的仓库地址默认在/home/gitmanager/repositories/目录下，这里我们需要创建一个链接映射。让他指向我们前面创建的专门用于存放项目的仓库目录/home/gitrepository。
命令如下：
sudo ln -s /home/gitrepository /home/gitmanager/repository

## 生成公钥
#### 在管理机器（你主要使用的电脑）上生成一个ssh公钥，命令如下：
ssh-krygen -t rsa

## 拷贝公钥文件到服务器
#### 将公钥文件拷贝到服务器上，命令如下：
 scp /home/[username]/.ssh/id_rsa.pub gitmanager@192.168.0.23:/home/gitmanager/id_rse.pub

## 初始化gitosis
#### 注意：初始化之前需要切换至gitmanager用户
su gitmanager
gitosis-init < home/gitmanager/id_rsa.pub(传到服务器的地址)

# 三、在服务器上创建项目仓库

## 创建仓库
#### 使用gitmanager账户在服务器上创建一个目录（mytestproject.git）并初始化成git项目仓库。命令如下：
su gitmanager
cd /home/gitrepository
mkdir mytestproject.git
git init --bare
exit

## SSh验证
#### 使用初始化Gitosis公钥的拥有者身份SSH进服务器，命令如下：
ssh gitmanager@192.168.0.23

## 可令Gitosis的控制仓库到本地
#### 命令为：
git clone gitmanager@192.168.0.23:/home/gitmanager/repositories/gitosis-admin.git

## gitosis-admin目录结构
### gitosis.conf
#### 用来设置用户、仓库和权限的控制文件
### keydir
#### 保存所有具有访问权限用户公钥的地方，每人一个

## 多人协助开发同一个版本
#### 将他们每个人的公钥文件添加到keydir文件夹然后push到服务器。文件的命名将决定在gitosis.conf配置文件中的称呼。

## 添加协同成员示例
### 1.为John,Josie和Jessica添加公钥：
$ cp /tmp/id_rsa.john.pub keydir/john.pub
$ cp /tmp/id_rsa.josie.pub keydir/josie.pub
$ cp /tmp/id_rsa.jessica.pub keydir/jessica.pub

### 2.把他们都加进 ‘mobile’ 团队,让他们对iphone_project具有读写权限:
[group mobile]
writable = iphone_project
members = scott john josie jessica

### 权限控制
#### Gitosis 也具有简单的访问控制功能。如果想让 John 只有读权限,可以这样做:
[group mobile]
writable = iphone_project
members = scott josie jessica
[group mobile_ro]
readonly = iphone_project
members = john

现在 John 可以克隆和获取更新,但 Gitosis 不会允许他向项目推送任何内容。

# 四、常见问题
## 运行ssh git@192.168.0.23出错
#### 重启电脑

## ERROR:gitosis.serve.main:Repository read access denied
### 原因
gitosis.conf中的members与keydir中的用户名不一致，如gitosis中的members = foo@bar，但keydir中的公密名却叫foo.pub
### 解决方法
使keydir的名称与gitosis中members所指的名称一致。 改为members = foo 或 公密名称改为foo@bar.pub

## clone时报does not appear to be a git repository
### 原因
####clone时不能用绝对路径，只能写相对于gitmanager用户home的相对路径。
我用的路径是：git clone gitmanager@192.168.0.6:gitosis-admin.git
### 解决方案
#### 将路径改为：相对gitmanager用户的路径：
git clone gitmanager@192.168.0.23:/home/gitmanager/repositories/gitosis-admin.git

# 五、参考资料
gitosis使用笔记 by YSHY
ubuntu完美搭建git服务器 by wxie的Linux人生

本文摘自：https://www.jianshu.com/p/d03efd263fe2
