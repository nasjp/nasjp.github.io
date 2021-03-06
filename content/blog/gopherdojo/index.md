---
title: Join GopherDojo
draft: false
date: 2019-09-08T16:49:03+09:00
---

[【リモート可】Gopher道場#7](https://gopherdojo.connpass.com/event/142892/)  
従前より参加し、Goのレベルを底上げしたいと考えておりました。  
Gopher道場に参加するためには、課題を提出し、選考を通過する必要があります。  
喜びとともに、*何に気をつけたのか* をつらつらと書き連ねます。

# 何に気をつけたのか

今回申し込みをする以前、[【リモート可】Go同新卒研修vol.1 by Gopher道場](https://mercari.connpass.com/event/135590/) に参加させていただきました。  
こちらとGopher道場のSlackのワーススペースは同じなんですが、Grobalチャンネルに主催の [@tenntenn](https://twitter.com/tenntenn)さんがGopher道場の告知と合わせて下記のような投稿をしておりました

> ↑もし知人の方にオススメする際には、以下の点を見るとお伝えください。  
> ・妥協なく品質を追求してるか  
> ・人が読むことを意識しているか  
> 特に品質を担保するにはどうすれば良いかという視点が漏れてる方が多数見受けられます。

さらにあわせて、 下記の記事を別の参加者の方が投稿をしてらっしゃいました。

> [34人のGoの課題をレビューして感じたこと](https://docs.google.com/presentation/d/1TAwxT9mRmiEjQOZurz-TbXc8SQf0mlEJDS8tL49Q-M4/edit#slide=id.g4e721da061_5_5)

ドキッとしました。  
Go同新卒研修 の課題の提出の際、わたしは`main.go`をドカンと送りつけていたのです。  


> ・妥協なく品質を追求してるか  
> ・人が読むことを意識しているか

これらを意識し、具体的には下記を提出ソースコードに盛り込みました。

1. ソースコードを分割する  
2. コメントを書く  
3. テストを書く  
4. READMEを書く  
5. タスクランナーとしてMakefileを使う
6. Go Modulesを導入する

###### 1.ソースコードを分割する

課題は Go Playground のリンクの共有によって提出します。  
提出フォームにも記載されておりましたが、下記の[@tenntenn](https://twitter.com/tenntenn)さんの記事を参考にしました。

[The Go Playground上で簡単に複数のファイルをシェアする #golang](https://qiita.com/tenntenn/items/4c2d33f795aa6e23e188)

##### 2. コメントを書く

当たり前ですね。  
[Effective Go のCommentaryの項目](https://golang.org/doc/effective_go.html#commentary) が非常に参考になります。

##### 3. テストを書く  

下記の記事がまとまっているので度々参考にさせていただいております。  
[Goのtestを理解する in 2018 #go](https://budougumi0617.github.io/2018/08/19/go-testing2018/)

##### 4. REDAMEを書く

ビルドとテストの仕方を書きました。

##### 5. タスクランナーとしてMakefileを使う

Goで書かれた有名なツールのMakefileを覗いて回るのが趣味になってます。  
下記からリポジトリを徘徊し、参考にさせていただいております。  
[Awesome Go](https://github.com/avelino/awesome-go)  
Makefile の ディグに夢中になって開発できなかったなんてのは本末が転倒していますのでほどほどに...

##### 6. Go Modulesを導入する

[Go & Versioning](https://research.swtch.com/vgo)  
[Modules](https://github.com/golang/go/wiki/Modules)  
Go Modules をあまり理解せずに使っていると GO Playground での共有で詰みます。  
課題の根幹に触れるような気がするのであまり言及しませんが、  
Go Playground 上に共有 -> 共有リンクから$GOPATH外に落としてくる -> go build  
みたいなことしてみると「おや?」ということになる可能性が高いので試してみることをおすすめします。

---

このような課題提出だけでなく、わたし以外(3ヶ月後のわたしも含む)が触れるソースコードは常に意識していきたい所存です。  
明日は Gopher道場 2回目です。台風が顔を出していますがどうなることやら...
