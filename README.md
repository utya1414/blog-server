- dirty version set
```
migrate -path db/migrations -database "postgresql://root:password@localhost:5434/blog?sslmode=disable" force 
1
```

- swagger editor shortcut
    Shift + Alt + P

- スキーマの変更
```
make migrateup
queryを変更
sqlc generate
ユニットテストを変更
```
#### ディレクトリ構造
- infrastructure
    - db/migrations golang-migrateのマイグレーションファイル
    - db/query sqlcのクエリファイル
    - repository/ 永続化の責務

### 思想
#### クリーンアーキテクチャ
クリーンアーキテクチャの採用により以下の性質をソフトウェアに持たせます。
- フレームワーク⾮依存：アーキテクチャは、機能満載のソフトウェアのライブラリに依
存していない。これにより、システムをフレームワークの制約で縛るのではなく、フ
レームワークをツールとして使⽤できる。
- テスト可能：ビジネスルールは、UI、データベース、ウェブサーバー、その他外部の要
素がなくてもテストできる。
- UI ⾮依存：UI は、システムのほかの部分を変更することなく、簡単に変更できる。た
とえば、ビジネスルールを変更することなく、ウェブ UI をコンソール UI に置き換える
ことができる。
- データベース⾮依存：Oracle や SQL Server を Mongo、BigTable、CouchDB などに
置き換えることができる。ビジネスルールはデータベースに束縛されていない。
- 外部エージェント⾮依存：ビジネスルールは、外界のインターフェイスについて何も知
らない。

また、以下の効果も期待できます。
- 関心ごとの分離
- 保守性とテスタビリティの向上

#### 責任
- repository
    データ永続化
- domain
    ビジネスルールの表現（バリデーション）
    - 必要なビジネスルールはapiのユースケースに応じて変化する。
        

#### スキーマ駆動開発
スキーマ駆動開発の採用により、以下の効果を期待できます。
- 

#### sqlc
#### golang-migrate
#### ULID
