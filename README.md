## Golang-Gin-Sample
このリポジトリは、Golang + Gin + Air + MySQLで以下の機能を実装したAPIです。

- 登録 / ログイン / ログアウト
- ログイン中のユーザの情報更新 & パスワード再設定
- User（ユーザー）, Role（ロール）, Permission（権限）, Product（商品）, Order（注文）、OrderItem（注文詳細）のCRUD（一部除く）
- MiddlewareによるJWT認証・権限チェック, CORS
- 画像アップロード
- Order・OrderItemのデータをCSVへ出力
- ページネーション

etc...

Dockerは用いていないので、ローカル環境は別途ご準備ください。
