---
title: "SSH Port Forwarding via bastion"
date: 2020-01-26T13:39:17+09:00
draft: false
tags: ["mysql", "aws"]
---

おそらく、データベースへの接続は

踏み台サーバーを経由して接続していると思いますが、

少し面倒くさいですよね。

私の職場では

```bash
$ ssh ec2-user@99.99.99.999 -i $HOME/.ssh/dev

Last login: Mon Jan 10 10:00:00 2020 from ***.***.***.***

       __|  __|_  )
       _|  (     /   Amazon Linux AMI
      ___|\___|___|


https://aws.amazon.com/amazon-linux-ami/$version-release-notes/
[ec2-user@ip-99-9-9-999 ~]$ mysql -udev  -hdev.ap-northeast-1.rds.amazonaws.com -pdev

Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 2 to server version: 9.9.99-community-nt

Type 'help;' or '\h' for help. Type '\c' to clear the buffer.

mysql>
```

一発でログインしたいものです。

ポートフォワーディングで解決します。

```bash
ssh -f -N -L 3306:dev.ap-northeast-1.rds.amazonaws.com:3306 -i $HOME/.ssh/dev ec2-user@99.99.99.999
mysql -udev -h 127.0.0.1 -P 3306 -pdev
```

私はローカルで[mycli](https://github.com/dbcli/mycli)を使用しています。
これを、ポートフォワーディングと組み合わせています。

```bash
ssh -f -N -L 3306:dev.ap-northeast-1.rds.amazonaws.com:3306 -i $HOME/.ssh/dev ec2-user@99.99.99.999
mycli -udev -h 127.0.0.1 -P 3306 -pdev
```

これらのステップを行う前にポートが空けておきましょう
```bash
kill $(ps aux | grep "ssh -f" | awk '{print $2;}')
ssh -f -N -L 3306:dev.ap-northeast-1.rds.amazonaws.com:3306 -i $HOME/.ssh/dev ec2-user@99.99.99.999
mycli -udev -h 127.0.0.1 -P 3306 -pdev
```

これを関数にまとめて、`.zshrc`に書いておけば最高の踏み台ライフの完成です。

私は`fish shell`を使用していますので下記のような`fish function`を使っています。
```bash
function ssh_dev
  echo "Login to dev mysql server via bastion..."
  kill (ps aux | grep "ssh -f" | awk '{print $2;}') >/dev/null 2>&1
  ssh -f -N -L 4306:dev.ap-northeast-1.rds.amazonaws.com:3306 -i $HOME/.ssh/dev ec2-user@99.99.99.999
  mycli -u dev -h 127.0.0.1 -P 4306  -p dev
end
```

最高すぎる。。。

ここまで書きましたが、もっと良い記事があった。。。

https://qiita.com/lighttiger2505/items/ea33291639a8656d50b4

以上です。
