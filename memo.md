# wsl2
## Windows
```
1 つのコマンドで WSL を実行するために必要なすべてのものをインストールできるようになりました。 
PowerShell または Windows コマンド プロンプトを 管理者モードで 開くには、右クリックして 
[管理者として実行] を選択し、wsl --install コマンドを入力して、
アカウント作成、パス設定
コンピューターを再起動します
```
## Windows
```
Windows の [スタート] メニューにアクセスし、インストールされているディストリビューションの名前を入力して、
Linux ディストリビューションを直接開くことができます。
 例: "Ubuntu"。 
これにより、独自のコンソール ウィンドウで Ubuntu が開きます。
```

## ubuntu
- install go
``` 
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```

## Windows vscode
VSCode拡張機能(WSL)のインストール

## ubuntu
- 開いたUbuntuで「code .」で初回拡張がインストール
```
ubuntuを起動させログイン
適当な作業ディレクトリを作成、vscodeを起動
mkdir test
cd test
code .
```

## url
```
https://learn.microsoft.com/ja-jp/windows/wsl/install
https://hirofurukawa.com/how-to-install-wsl2/
https://zenn.dev/s_t_taka/articles/2d64bf0ed3c9a4
https://qiita.com/KoHey94/items/5abf60e612cfb6ec263e
https://qiita.com/SAITO_Keita/items/148f794a5b358e5cb87b
https://qiita.com/zaburo/items/27b5b819fae2bde97a3b#wsl2%E3%81%AE%E3%82%A4%E3%83%B3%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB
https://minettyo.hatenablog.com/entry/wsl_update
https://www.mitsue.co.jp/knowledge/blog/x-tech/202507/15_0843.html
https://zetamatta.hatenablog.com/
```

# docker
## wls2 ubuntu
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

## wls2 ubuntu
- gin docker build
```
sudo docker build  --tag app-hello-gin:v1.0 .
```
- gin dicker run
```
sudo docker run -p 8080:8080 app-hello-gin:v1.0 
```

# budibase
## wls2 ubuntu
- budibase docker
```
sudo docker run -d -t \
  --name=budibase \
  -p 10000:80 \
  -v /home/tttda/propjets/test_budibase/data:/data \
  --restart unless-stopped \
  budibase/budibase:latest
```

## windows
- open browser
http://localhost:10000/builder/admin

## url
https://budibase.com/blog/tutorials/employee-management-app/
