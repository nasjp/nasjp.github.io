---
title: Easay to login Mysql datbase on Docker container
date: 2020-02-26T15:42:52+09:00
draft: false
tags: ["docker", "mysql"]
---

すでに立ち上がっている docker container に `docker exec` で入って、データベースを確認したいときありますよね

でも container id は毎回変わてしまいます。

なので `docker ps` で container id を調べて、`docker exec` して、`mysql -u root` を毎回しなければいけません。

だるい。

下記で解決です。

```sh
docker exec -it $(docker ps | grep "container-name_1" | awk '{print $1;}') mysql -u root
```

`grep "container-name_1"` を任意の container name に変更してください。

記事にするほどでもねえ。。。

---

ちなみに

`mysql -u root` を変更すれば任意のコマンドを実行できます。
