桌面版
1、首先对系统进行更新
   sudo apt-get update
   sudo apt-get dist-upgrade  选Y
   上面的命令会下载最新的内核和软件包

2、重启系统完成安装
   sudo reboot

3、执行如下命令打开更新管理器
   sudo update-manager -d
   他会自动查找最新可用版本
   点击"Upgrade"->"ok"->"Upgrade"->"start upgrade"
   安装完成之后重启系统

服务器版
Ubuntu 14.04/15.10 Server升级到Ubuntu 16.04 Server
1、如果你的系统中没有安装update-manager-core软件包，安装他：
   sudo apt-get install update-manager-core

2、编辑文件 _/etc/update-manager/release-upgrades_:
   sudo vim /etc/update-manager/release-upgrades
   设置Prompt=normal
   Normal：检查新版本，如果有多个新版本可以升级，系统试图升级离当前使用         版本最近的。
   LTS：检查长期支持新版本，如果当前版本不为LTS，不要使用它。
   如果你的系统是Ubuntu 14.04 Server，设置Prompt=lts。
3、在升级系统前，先更新一下系统
   sudo apt-get update && sudo apt-get dist-upgrade
4、重启系统
   sudo reboot

如果你使用的是ssh登陆服务器升级，建议使用screen，防止SSH连接断开
1、安装screen：
   sudo apt-get install screen
2、启动screen
   screen
3、升级Ubuntu
   sudo do-release-upgrade -d
   根据提示一路Y、Y、Y、Y。。。
   等待升级完成。