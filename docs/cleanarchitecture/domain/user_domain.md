### ユースケース
- ユーザー新規登録
- 単一ユーザー検索
- 全ユーザー検索
- 監査ログ出力

### ビジネスルール
- Username
   - 1文字以上20文字以下

   - 英数字及びアンダースコアのみ
      日本語対応はエンコード処理の実装が必要そうなので後回し
- Email
   - メールアドレス形式
   　 RFC 5322に準拠

### ドメインモデル
- User
   - ID
   - Username
   - Email
   - CreatedAt
   - UpdatedAt