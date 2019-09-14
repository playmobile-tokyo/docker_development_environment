# vscode で xdebug

## 環境

mac
virtualbox
 centOS
  network nat
docker-comporse
 nginx
 php-fpm

setting pointは余計なことせず、PCに割り振られているIPアドレスを xdebug.iniに設定する。

* virtualbox の portforwarding　は設定しない
* docker の exporse や port は設定しない
* ip アドレスは ifconfigで設定されているPCのIPアドレスを設定する
