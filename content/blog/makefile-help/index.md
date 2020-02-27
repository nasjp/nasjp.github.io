---
title: Makefile Help
date: 2020-02-27T16:50:19+09:00
draft: false
tags: ["makefile"]
---

突然ですが、`Makefile` に下記のような行を見たことないですか?

```bash
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
```

これがほんとに便利なんです。

コードの詳細は説明しませんが、コマンド名の通り`Makefile` に ヘルプを追加します。

例えば、、、

上記の行と、`hello`コマンドと`bye`コマンドが書かれた`Makefile`を用意します。

```bash
$ cat Makefile
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: hello
hello: ## this command echo hello
	echo hello

.PHONY: bye
bye: ## this command echo bye
	echo bye
```

`make` を実行してみます。

```bash
$ make
hello                          this command echo hello
bye                            this command echo bye
```

こんなかんじで　コマンドのヘルプを出せるようになりました。

肝は `Makefile` の下記の行です。

```bash
hello: ## this command echo hello
```

`##` 以降をコマンドのヘルプとして表示します。

逆にこれを書かなければヘルプは表示されません。(だから`help`コマンドのヘルプは表示されない)

ぜひ使ってみてください。

以上です。
