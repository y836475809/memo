# virtualbox
- 共有フォルダ
- https://tanoike.com/virtualbox-shared-folder-linux

# virtualbox docker
## ubuntu
https://docs.docker.com/engine/install/ubuntu/

- remove old
```
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done
```

- Add Docker's official GPG key
```
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
```

- Add the repository to Apt sources
```
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```
- install docker
```
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

# go
- install go
``` 
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```
- https://qiita.com/TakanoriVega/items/6d7210147c289b45298a
- https://qiita.com/enoenoeno11/items/c24516ebe42606e601b7
- sample2
  - https://qiita.com/hichihara/items/176d4c15bd240d7f2b0d
```
$ gcc -c hello.c
$ ar rusv libhello.a hello.o
```
- test
  - https://zenn.dev/geniee/articles/207821e953a900
  - https://qiita.com/theoden9014/items/ac8763381758148e8ce5
  - https://qiita.com/268iop/items/dc6e73138ee97dff2f73
  - https://gin-gonic.com/ja/docs/testing/
  - https://qiita.com/torat3158/items/7b38bc17c2d1129119f9
  - https://dev.to/truongpx396/golang-integration-test-with-gin-gorm-testify-postgresql-1e8m

# ubuntuでbudibase(docker)
- budibase docker
```
sudo docker run -d -t \
  --name=budibase \
  -p 10000:80 \
  -v /home/tttda/propjets/test_budibase/data:/data \
  --restart unless-stopped \
  budibase/budibase:latest
```
- hostでurlひらく
http://localhost:10000/builder/admin

# sample3 sql server(docker)
- [【SQL Server 2019】 Developer と Express の違い](https://qiita.com/Sanada-code/items/fb19e98f9201bf1d8e1f)
  - Express: 機能制限あり、本番環境もふくめて自由に使用可能
  - Developer：全機能使用可能、テスト用としてのみライセンスを認める
- [Docker で SQL Server を実行する(Expressエディション)](https://www.curict.com/item/99/99ff31e.html)
- [クイック スタート:Docker を使用して SQL Server Linux コンテナー イメージを実行する](https://learn.microsoft.com/ja-jp/sql/linux/quickstart-install-connect-docker?view=sql-server-ver17&tabs=cli&pivots=cs1-bash)
- [SQL Server 2022をDockerコンテナで動かす](https://qiita.com/charon/items/6a7cae83b0d2aea6258e)
- [Docker で Microsoft SQL Server を実行するときのメモ](https://zenn.dev/shimiyu/scraps/e4b93ef1c47a08)
  - docker-compose.yamlでMSSQL_PID=ExpressでExpress Editionを指定
- [DockerでSQLServerを起動してデータベースをリストアするまでの記録](https://twinbird-htn.hatenablog.com/entry/2025/01/11/043000)
- https://learn.microsoft.com/ja-jp/sql/linux/quickstart-install-connect-docker?view=sql-server-ver17&tabs=cli&pivots=cs1-bash#remove-your-container
- https://github.com/microsoft/mssql-docker/issues/13#issuecomment-641904197
- https://blog.hn-pgtech.com/2023-02-06/
- https://docs.docker.com/engine/storage/drivers/overlayfs-driver/#configure-docker-with-the-overlay-or-overlay2-storage-driver

- virtualboxのubuntuにdockerのsql server
  - 共有フォルダだとエラーでdocker起動できない(ホームの適当なフォルダにしたら起動できた)
  - vbのポートフォワーディング設定(https://qiita.com/nacanaca/items/34b8a7c4d01dab993e4a) 
- windows(ホスト)にMicrosoft SQL Server Management Studioインストール, vbのdockerのsql serverに接続
  - 「ログイン」のサーバ名はlocalhost,1433
  - 「追加の接続パラメータ」にTrustServerCertificate=Trueを追加
  - https://zenn.dev/nagiyu/articles/68f005eda86218
  - https://ippun-blog.com/dockeronsqlserverkaisetsu/

# sample4 minio(docker)
- virtualboxのubuntuにdockerのminio
- image
  - https://hub.docker.com/r/minio/minio
  - sudo docker pull minio/minio:RELEASE.2025-07-23T15-54-02Z
- vbのポートフォワーディング設定(https://qiita.com/nacanaca/items/34b8a7c4d01dab993e4a) 
- compose
  - https://blog.adglobe.co.jp/entry/2024/11/15/100000
  - https://qiita.com/h_tyokinuhata/items/c2b76b39b99728d8246a
  - webui http://127.0.0.1:9001

# sample5 ローカルモジュール参照(go.work), http server test sample
- [Go言語でパッケージの修飾名が重複した場合の対応方法](https://taknb2nch.hatenablog.com/entry/20140211/1392046399)
- [Gin を使って複数のサービスを稼働させる](https://gin-gonic.com/ja/docs/examples/run-multiple-service/)
- [HTTPサーバーは安全に終了](https://zenn.dev/tksx1227/articles/5ab5b3c99336c3)
- [ローカルPCでGoでserverを立ち上げたときにファイアウォールを出なくさせる](https://ludwig125.hatenablog.com/entry/2020/06/10/064119)
- https://zenn.dev/ikawaha/articles/20220701-a053ec54b77435
- https://zenn.dev/jy8752/scraps/21b9d548cc9b09
- https://go.dev/doc/tutorial/workspaces
- https://qiita.com/h-masano/items/bf7bdfba7c4e8f6b28c4

# sample6 http get post test 
- https://github.com/h2non/gock

# sample7 cgo
- https://qiita.com/endoh/items/4410309f5114cafc265e
- https://qiita.com/saturday06/items/84535c61a3328c02032c
- https://ericchiang.github.io/post/cgo/
- https://hnakamur.github.io/blog/2019/12/29/cgo-and-unsafe/
