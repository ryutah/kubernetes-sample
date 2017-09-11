# CloudSQL接続手順
1. CloudSQLにアクセス可能なサービスアカウントを作成する

2. CloudSQLのユーザを作成する
  ```console
  $ gcloud beta sql users create gke-app cloudsqlproxy~% --instance=test-tryumphdb --password=gke-app
  ```

3. 接続名を取得する
  ```console
  $ gcloud beta sql instances describe test-tryumphdb

    backendType: SECOND_GEN
    connectionName: tryumph-test01-project:asia-east1:test-tryumphdb # これ
    databaseVersion: MYSQL_5_7
    etag: '"AX0XCS8cmlmWUsraAUQ8dEhn1gI/MTU"'
    gceZone: asia-east1-a
    instanceType: CLOUD_SQL_INSTANCE
    ipAddresses:
    - ipAddress: 35.194.136.188
      type: PRIMARY
    ...
  ```

4. インスタンスレベルアクセスを設定する
  \# [インスタンスのアクセス制御について](https://cloud.google.com/sql/docs/mysql/instance-access-control?hl=ja)
  ```console
  # インスタンスレベルのアクセスのためのシークレットを作成
  $ kubectl create secret generic cloudsql-instance-credentials \
                       --from-file=credentials.json=${CLOUD_SQL_SERVICE_ACCOUNT_KEY_FILE}
  ```

5. データベースレベルアクセスを設定する
  \# [インスタンスのアクセス制御について](https://cloud.google.com/sql/docs/mysql/instance-access-control?hl=ja)
  ```console
  # データベースアクセスに必要なシークレットを作成
  $ kubectl create secret generic cloudsql-db-credentials --from-literal=username=${CLOUD_SQL_USER_NAME}
  ```

6. ポッド設定ファイルの設定
  `k8s/deployment.yml` 参照
