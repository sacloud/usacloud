# usacloud v1でのコマンド実装方法について

## 全体の流れ

- `pkg/cmd/commands`配下にリソース名のパッケージを作成(なるべくlibsacloud serviceのパッケージ名と合わせる)
- 作成したパッケージ内にリソース定義(`resource.go`)を作成
- 作成したパッケージ内にコマンド定義(例:`list.go`)を作成
- `pkg/cmd/resources.go`に作成したリソース定義を追加する
- コード生成する - `make` or `make gen`

## 各手順について

### `pkg/cmd/commands`配下にリソース名のパッケージを作成

追加したいリソースが`archive`の場合、`pkg/cmd/commands/archive`ディレクトリを作成する。

Note: 名前はlibsacloud serviceのものに合わせる。

#### 基本的な命名規則

- リソース名を小文字にしたもの
- 単語の区切りはなし。ハイフンやアンダーバーを用いない
- 以下2つはGoの予約語とぶつからないように特殊なルールを適用
  - NIC(インターフェース)リソース => `iface`
  - スイッチ リソース => `swytch`

### 作成したパッケージ内にリソース定義(`resource.go`)を作成

以下のように実装。詳細は`core.Resource`や既存リソースの実装を参照。

```go
var Resource = &core.Resource{
	Name:        "archive",
	ServiceType: reflect.TypeOf(&archive.Service{}),
	Category:    core.ResourceCategoryStorage,
	CommandCategories: []core.Category{
		{
			Key:         "basic",
			DisplayName: "Basic Commands",
			Order:       10,
		},
		{
			Key:         "operation",
			DisplayName: "Archive Operation Commands",
			Order:       20,
		},
		{
			Key:         "other",
			DisplayName: "Other Commands",
			Order:       1000,
		},
	},
}

var defaultColumnDefs = []output.ColumnDef{
	{Name: "Zone"},
	{Name: "ID"},
	{Name: "Name"},
	{Name: "Scope", Template: "{{ scope_to_key .Scope }}"},
}
```

### 作成したパッケージ内にコマンド定義(例:`list.go`)を作成

作成の際はv0の`pkg/define`内の定義とlibsacloud serviceのリクエストの定義を参照しながら必要なパラメータをstruct+フィールドタグで定義していく

```go
var listCommand = &core.Command{
	Name:               "list",
	Aliases:            []string{"ls", "find", "select"},
	Category:           "basic",
	Order:              10,
	ServiceFuncAltName: "Find",
	NoProgress:         true,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newListParameter()
	},
}

type listParameter struct {
	cflag.ZoneParameter `cli:",squash" mapconv:",squash"`

	Names               []string `cli:",category=filter" validate:"omitempty"`
	Tags                []string `cli:",category=filter" validate:"omitempty"`
	OSType              string   `cli:",category=filter,options=os_type" mapconv:",omitempty,filters=os_type_to_value" validate:"omitempty,os_type"`
	Scope               string   `cli:",category=filter,options=scope" mapconv:",omitempty,filters=scope_to_value" validate:"omitempty,scope"`
	cflag.FindParameter `cli:",squash" mapconv:",squash"`

	cflag.OutputParameter `cli:",squash" mapconv:"-"`
}

func newListParameter() *listParameter {
	return &listParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	Resource.AddCommand(listCommand)
}
```

この際にタグでのバリデーションやフィルタが必要であれば`pkg/vdev/definitions.go`に値定義(この例だと`scope`)を定義する。

```go
// definitions usacloudで使う名称(key)/値(value)のペア
var definitions = map[string][]*definition{
	/* ... */
	"scope": {
		{key: types.Scopes.User.String(), value: types.Scopes.User},
		{key: types.Scopes.Shared.String(), value: types.Scopes.Shared},
	},
}
```

`ConflictsWith`などのタグでのバリデーションでカバーできないものはここでカスタムバリデーションを実装する。
例: `archive`や`disk`の`create`サブコマンド、`download`コマンドなど

### `pkg/cmd/resources.go`に作成したリソース定義を追加する

`pkg/cmd`の`Resources`に作成したリソースを追記する。

```go
var Resources = core.Resources{
	archive.Resource, // 追記
	/* ... */
}
```

### コード生成する - `make` or `make gen`

Note: 生成したコードが原因でビルドできなくなったり`make`が失敗する場合は`make clean-all`で生成したコードを一旦全削除するとうまくいくことがある。