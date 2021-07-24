# sushiware.net

```sh
$ node -v
v14.16.1
$ npm -v
6.14.12
```

```sh
npm outdated | sed 1d | awk '{print $1}' | xargs -I{} npm install {}@latest
```
