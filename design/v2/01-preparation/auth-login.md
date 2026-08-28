# `auth login` コマンド

## 目的

現在は、usacloud を使い始めるための認証情報を `config edit` で設定する。
しかし `config edit` は zone や出力形式なども扱う汎用的な設定編集コマンドであり、
初回利用者には次のことが分かりにくい。

- usacloud の利用開始には API キーの設定が必要である。
- API キーをどこで作り、どの項目へ入力するのか。
- 設定した API キーが有効か、どのアカウントへ接続しているか。
- usacloud の利用を終了するとき、認証情報をどのように削除・失効するか。

そこで、一般設定を編集する `config` とは別に、認証の開始・確認・終了を扱う
`auth` コマンドを用意する。

```bash
usacloud auth login
usacloud auth status
usacloud auth logout
```

`auth login` の主目的は、認証方式にかかわらず「このコマンドを実行すれば usacloud を
使い始められる」という明確な入口を提供することである。API キーを手入力する段階でも、
作成場所の案内、入力内容の検証、安全な保存までを一つのフローとして扱う価値がある。

将来、サーバー側でブラウザ認証を提供できるようになった場合も、利用者は同じ
`auth login` を使い続けられる。この文書では、そのログインフローを実現可能性に
応じて 3 段階に分ける。

| 段階 | ログイン方法 | サーバー側の開発 |
|---|---|---|
| 1. 現状 | 利用者が API キーを作成し、`config` へ設定する | 不要 |
| 2. 現行 API での改善 | `auth login` が API キーの作成と設定を案内する | 不要 |
| 3. 将来 | ブラウザで認可するだけでログインできる | 必要 |

まず Phase 2 を実装し、サーバー側の準備ができた段階で、同じコマンドを Phase 3 の
フローへ切り替える。

## Phase 1: 現状

現在は、利用者が次の操作を行う。

1. さくらのクラウド コントロールパネルへログインする。
2. API キーを作成する。
3. アクセストークンとシークレットをコピーする。
4. `usacloud config edit` または `usacloud config create` へ入力する。

環境変数でも指定できる。

```bash
export SAKURACLOUD_ACCESS_TOKEN=...
export SAKURACLOUD_ACCESS_TOKEN_SECRET=...
```

### 問題

- 初回利用者には、API キーの作成場所や `config` との関係が分かりにくい。
- トークンとシークレットの組み合わせを間違えても、保存するまで分からない。
- 保存時に、API キーが有効か、どのアカウントのものかを確認しない。
- ローカルの設定を削除しても、コントロールパネル上の API キーは残る。
- 認証設定と、zone や出力形式などの一般設定が同じ `config` にまとめられている。

## Phase 2: 現行 API で実装できる改善

サーバー側を変更せず、API キーの設定を案内する `auth login` を追加する。

```bash
usacloud auth login
```

### フロー

1. 保存先の profile を選ぶ。
2. API キー作成画面の URL と作成手順を表示する。
3. 可能であればブラウザでコントロールパネルを開く。
4. 利用者が作成したアクセストークンとシークレットを入力する。
5. 保存前に API を呼び、入力した組み合わせが利用可能か確認する。
6. 認証情報を保存し、対象 profile を選択中にする。

出力例:

```text
Opening the Sakura Cloud API key page in your browser.

Enter your API token: ********
Enter your API secret: ********

✓ Authentication succeeded.
✓ Saved credentials to profile "default".
```

トークンとシークレットはコマンドライン引数では受け取らない。shell history や
プロセス一覧に残さないため、対話的に入力し、端末には表示しない。

### 保存方法

可能であれば、認証情報は OS の credential store に保存する。

| OS | 保存先 |
|---|---|
| macOS | Keychain |
| Windows | Credential Manager |
| Linux | Secret Service 対応 keyring |

profile には credential store の参照と、アカウントや認証方式など秘密でない情報を
保存する。credential store を利用できない場合は、利用者へ確認した上で既存の profile
形式へ保存する。

### Phase 2 で改善できないこと

- API キーのコピー&ペーストは必要。
- API キーは、利用者がコントロールパネルで明示的に失効するまで有効。
- `auth logout` だけでは、サーバー上の API キーを失効できない。
- ブラウザでのログイン結果を CLI が受け取ることはできない。

Phase 2 の `auth logout` はローカルの認証情報を削除し、API キーを完全に失効させる
ためのコントロールパネル URL と手順を表示する。

## Phase 3: サーバー側開発後のログイン

さくらインターネット側で Device Authorization API を提供し、ブラウザ上の認可結果を
usacloud が受け取れるようにする。

利用者から見たフローは次のようになる。

```text
$ usacloud auth login

! Copy your one-time code: ABCD-EFGH
Opening https://auth.sakura.ad.jp/device in your browser.
Waiting for authorization...

✓ Logged in to account example-account.
✓ Saved credentials to profile "default".
```

1. usacloud が一時的なコードとログイン URL を取得する。
2. URL とコードを表示し、ブラウザを開く。
3. 利用者がブラウザでログインする。
4. 利用するクラウドアカウントと権限を確認して認可する。
5. usacloud が認可完了を待ち、認証情報を受け取る。
6. 認証情報を保存し、API を利用可能にする。

ブラウザを開けない SSH 接続先などでは、別の端末で URL を開き、同じコードを入力する。
そのため、ローカル callback server は必要ない。

### サーバー側に必要なもの

実装方法の詳細は認証基盤側で設計する。この機能に必要なのは次の役割である。

- usacloud に一時的なコードとログイン URL を発行する。
- ブラウザでログイン、アカウント選択、権限確認を行う。
- 認可完了後、usacloud に期限付きの認証情報を発行する。
- logout や管理画面から、発行済みの認証情報を失効できる。
- 各サービス API が、その認証情報を受け付けられるようにする。

### 想定する要素技術

具体的な実装は認証基盤側の設計に委ねるが、標準技術として次の採用を想定する。

| 要素技術 | 用途 |
|---|---|
| OAuth 2.0 Device Authorization Grant | CLI とブラウザを一時コードで結び付ける |
| OpenID Connect | ログインしたユーザーとアカウントを確認する |
| 短命 access token | 各サービス API を呼び出す |
| refresh token | access token を再ログインなしで更新する |
| Bearer token 認証 | 各サービス API が新しい認証情報を受け付ける |
| OS credential store | CLI が refresh token などを安全に保存する |

access token の形式には JWT、サーバー側で状態を管理する opaque token などが考えられる。
この選択は既存の認証基盤や各サービス API の構成に合わせて決定し、usacloud の
コマンド仕様には露出させない。

Phase 3 では、ブラウザログインの裏側で従来の長寿命 API キーを自動作成する方式は
採用しない。短い有効期限の認証情報を発行し、必要に応じて自動更新する方式にする。

これにより、API キーのコピーが不要になり、logout による失効や、有効期限による
漏えいリスクの低減が可能になる。

## コマンド体系

```text
usacloud auth login [flags]
usacloud auth status [flags]
usacloud auth logout [flags]
```

| コマンド | 役割 |
|---|---|
| `auth login` | 認証を行い、credential を profile へ保存する |
| `auth status` | profile、アカウント、認証方式、利用可能かを表示する |
| `auth logout` | ローカルの credential を削除し、可能ならサーバー側でも失効する |

既存の `config` コマンドは、zone、出力形式、profile の一覧や切り替えなどに引き続き
利用する。

### `auth login`

```text
Usage:
  usacloud auth login [flags]

Flags:
  --no-browser      Print the URL without opening a browser
  --profile string  Profile to save credentials to (default: default)
  --with-token      Enter an existing API token and secret
```

Phase 2 では API キー入力が既定のフローになる。Phase 3 ではブラウザ認証を既定にし、
従来の API キーを使う場合だけ `--with-token` を指定する。

既存 profile を上書きする場合は確認する。zone や出力形式など、認証以外の設定は
保持する。

### `auth status`

```bash
usacloud auth status
usacloud auth status --profile staging
```

次の情報を表示する。

- profile 名
- 認証方式（API key または browser login）
- アカウント
- 認証が現在利用可能か
- 有効期限がある場合はその期限

トークンやシークレットそのものは表示しない。ネットワーク障害と、認証情報の失効や
期限切れは区別して報告する。

### `auth logout`

```bash
usacloud auth logout
usacloud auth logout --profile staging
```

Phase 2 ではローカルの認証情報だけを削除し、API キーを失効するための URL を表示する。
Phase 3 ではサーバー側の認証情報も失効してから、ローカルの情報を削除する。

profile の zone などは残す。profile 全体を削除する場合は
`usacloud config delete` を利用する。

## 実装順序

1. Phase 2 の `auth login/status/logout` を実装する。
2. 認証情報を OS の credential store へ保存できるようにする。
3. サーバー側で Device Authorization API とブラウザの認可画面を提供する。
4. 各サービス API と SDK を、新しい認証情報に対応させる。
5. `auth login` の既定を Phase 2 から Phase 3 のフローへ切り替える。

## 未決事項

- ブラウザで選択できる権限の単位。
- 認証情報の有効期限と、再ログインが必要になる期間。
- 各サービス API をどの順序で新しい認証方式へ対応させるか。
- OS の credential store を使えない環境での保存方法。

## 参考

- [GitHub CLI `gh auth login`](https://cli.github.com/manual/gh_auth_login)
- [AWS CLI IAM Identity Center authentication](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html)
- [OAuth 2.0 Device Authorization Grant](https://www.rfc-editor.org/rfc/rfc8628)
- [さくらのクラウド API キー](https://manual.sakura.ad.jp/cloud/api/apikey.html)
