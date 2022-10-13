# vodka
> vodka 是個手腳架(鷹架)，對於現在Golang沒有比較好定義應用程式專案的目錄結構的問題．vodka 作為一個定義應用程式專案的目錄結構，已將每一個層級定義出來．並在每個層級融入OOP化，把**專案可控、初階工程師、進階工程師，混合作業作為最高指導原則所開發的**．
vodka是遵循[SOLID](https://en.wikipedia.org/wiki/SOLID)設計規範．

## 目錄結構
```
├── LICENSE
├── Makefile
├── README.md
├── air.example
├── api
├── build
│   └── circleci
│       └── config.yml
├── cmd
├── config
│   ├── config.go.example
│   └── debug_config.go
├── deploy
├── docs
├── go.mod
├── go.sum
├── http
├── internal
│   ├── entity
│   │   ├── access_log
│   │   │   └── access_log_entity.go
│   │   ├── db
│   │   │   ├── access_logs
│   │   │   │   └── access_logs.go
│   │   │   └── members_phone
│   │   │       └── members_phone.go
│   │   └── member_phone
│   │       ├── member_phone_entity.go
│   │       ├── member_phone_suite_test.go
│   │       └── member_phone_test.go
│   ├── interactor
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── helpers
│   │   │   └── amazon
│   │   │       └── s3
│   │   │           └── s3.go
│   │   ├── manager
│   │   ├── models
│   │   │   ├── access_logs
│   │   │   │   └── access_logs.go
│   │   │   ├── page
│   │   │   │   └── page.go
│   │   │   ├── section
│   │   │   │   └── section.go
│   │   │   └── special
│   │   │       └── backend.go
│   │   ├── service
│   │   │   └── access_log
│   │   │       └── access_log_service.go
│   │   └── util
│   │       ├── aes.go
│   │       ├── base.go
│   │       ├── code
│   │       │   └── status.go
│   │       ├── connect
│   │       │   ├── mysql.go
│   │       │   └── postgres.go
│   │       ├── jwe.go
│   │       ├── log
│   │       │   └── log.go
│   │       ├── page.go
│   │       ├── time.go
│   │       ├── util.go
│   │       └── uuid.go
│   ├── presenter
│   └── router
│       ├── middleware
│       └── router.go
├── main.go
├── migrations
│   ├── 20220906071052_create_access_logs.down.sql
│   ├── 20220906071052_create_access_logs.up.sql
│   ├── 20220906072034_create_members_phone.down.sql
│   └── 20220906072034_create_members_phone.up.sql
├── scripts
│   ├── environment.example
│   └── migration.example
└── tools
    ├── autoMigrate
    │   └── main.go
    ├── log
    │   └── run.go
    └── testData
        └── main.go
```
### LICENSE
* 授權檔案 MIT License

### Makefile
* 詳細記錄了,所有能夠使用的命令集

### air.example
* 熱加載設定檔

### /api
* API DOC 置放

### /build
* CI/CD 設定檔

### /cmd
* 本專案的主要應用程式

### /config
* 設定檔

### /deploy
* 被編譯過後的檔案

### /docs
* golang doc

### go.mod
* go mod檔

### /http
* restful api 測試文件

### /internal
* 私有應用程式和函式庫的程式碼,是你不希望其他人在其應用程式或函式庫中匯入的程式碼.

### /internal/entity
* 對應資料庫的CRUD

### /internal/entity/db
* 對應資料表結構檔

### /internal/interactor
* 可共用或不可共用的函式庫的程式碼

### /internal/interactor/constants
* 置放常數資料夾

### /internal/interactor/helpers
* 置放一些共用的manager

### /internal/interactor/manager
* 置放交互調度程式

### /internal/interactor/models
* 置放共用結構檔

### /internal/interactor/models/page
* 置放有關於分頁的結構檔

### /internal/interactor/models/section
* 置放有關於時間的結構檔

### /internal/interactor/models/special
* 置放有關於後端共用結構檔

### /internal/interactor/service
* 置放有關於所有的服務的程式碼

### /internal/interactor/util
* 置放一些小工具,實用程序

### /internal/interactor/util/code
* 置放錯誤代碼或是回應代碼

### /internal/interactor/util/connect
* 置放對應資料庫的連線

### /internal/presenter
* 對應前端第一接觸的地方,API文件註解的地方,驗證輸入的地方

### /internal/router
* 置放路由設定

### main.go
* main func

### /migrations
* 放置資料庫的SQL檔案

### /scripts
* 置放一些可以執行的SH檔

### /tools
* 可以置放一些小工具

# 安裝
* 以下命令可以建議一些執行的sh檔跟設定檔
```shell!
make setup
```

# 執行
* 測試環境
```shell!
make runTesting
```

* 正式環境
```shell!
make runProduction
```

# 資料庫遷移
> 使用[golang-migrate](https://github.com/golang-migrate/migrate)做資料庫遷移及做資料表版控
* 測試環境
```shell!
make migrationTesting
```

* 正式環境
```shell!
make migrationProduction
```

# 更新日誌檔
* 可以產生change log
```shell!
make changeLog
```

# License
Vodka is released under the MIT license.