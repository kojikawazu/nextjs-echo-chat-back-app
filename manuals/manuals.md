# GoLang/Echoの環境構築

## 環境構築

### 1. GoLangのインストール

```bash
# インストール
brew install go

# バージョン確認
go -version
```

### 2. Goプロジェクトの初期化

```bash
# プロジェクトの初期化
go mod init echo-chat-app

# モジュールの確認
go mod tidy
```

### 3. Echoのインストール

```bash
# インストール
go get -u github.com/labstack/echo/v4

# モジュールの確認
go mod tidy
```



