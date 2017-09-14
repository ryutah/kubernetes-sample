# CloudSQL接続手順
1. GKEクラスタを作成する  
  ```console
  $ gcloud config configurations activate sandbox-hara

  $ gcloud container clusters create sample-gke \
      --network sample-gke \
      --scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring.write","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append"
  ```

  \# 必要なスコープはまだ未整理(一応この設定なら問題なし)

2. クレデンシャルを設定
  ```console
  $ gcloud container clusters get-credentials sample-gke
  ```

3. CloudSQLにアクセス可能なサービスアカウントを作成する

4. CloudSQLのユーザを作成する
  ```console
  $ gcloud beta sql users create gke-app cloudsqlproxy~% --instance=test-tryumphdb --password=gke-app
  ```

5. 接続名を取得する
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

6. インスタンスレベルアクセスを設定する
  \# [インスタンスのアクセス制御について](https://cloud.google.com/sql/docs/mysql/instance-access-control?hl=ja)
  ```console
  # インスタンスレベルのアクセスのためのシークレットを作成
  $ kubectl create secret generic cloudsql-instance-credentials \
        --from-file=credentials.json=${CLOUD_SQL_SERVICE_ACCOUNT_KEY_FILE}
  ```

7. データベースレベルアクセスを設定する
  \# [インスタンスのアクセス制御について](https://cloud.google.com/sql/docs/mysql/instance-access-control?hl=ja)
  ```console
  # データベースアクセスに必要なシークレットを作成
  $ kubectl create secret generic cloudsql-db-credentials \
        --from-literal=username=gke-app \
        --from-literal=password=gke-app
  ```

8. ポッド設定ファイルの設定
  `k8s/deployment.yml` 参照
