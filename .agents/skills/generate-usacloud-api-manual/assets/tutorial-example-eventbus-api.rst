================================================================================
eventbus-api チュートリアル
================================================================================

``usacloud eventbus-api`` を使うと、スケジュールやクラウド上のイベントを契機として
ジョブを実行できます。本チュートリアルでは、EventBus から SimpleMQ へ実際に
メッセージを送り、次の2つの経路が動作することを確認します。

- schedule の開始時刻を迎えると SimpleMQ にメッセージが届く
- スイッチを作成すると trigger がイベントを検知し、SimpleMQ にメッセージが届く

.. warning::

   このチュートリアルでは **本番環境に EventBus、SimpleMQ、スイッチのリソースを
   作成・削除** します。作業を始める前に、以下のコマンドで現在の profile を確認し、
   対象プロファイルが意図したものであることを確認してください。

   .. code-block:: shell

      usacloud config current

   対象 profile でこれらのリソースを作成・削除することと、料金や影響を確認してから
   続行してください。

このサービスはいつ使うか
================================================================================

EventBus は、イベント検知サービスとジョブスケジュールサービスを統合した
マネージドサービスです。「一定間隔で処理したい」「クラウド上のリソース変更を
検知して処理したい」という場合に、スケジュールやイベントソースをトリガーとして
ジョブを自動実行できます。

スケジュールの周期は最短1分で、ジョブの実行はベストエフォートです。
本チュートリアルでは、結果を SimpleMQ のメッセージとして受信することで、
設定の作成だけでなくジョブが実行されたことまで確認します。

詳細は `EventBusの基本情報 <https://manual.sakura.ad.jp/cloud/appliance/eventbus/about.html>`_
を参照してください。

前提条件
================================================================================

- usacloud と jq がインストールされていること
- 対象 profile が設定済みで、リソースの作成・削除に必要な権限を持つこと
- SimpleMQ の API キーと受信結果を保存するローカル一時ファイルを安全に扱えること
- スイッチを作成するゾーンとして ``is1c`` を利用できること

SimpleMQ のメッセージ API にはキューごとの API キーを使用します。
API キーはチュートリアルやコマンドラインへ直接記載せず、権限を制限した一時ファイルに
保存します。

このチュートリアルで作成するもの
================================================================================

以下の順序で作成・確認します。

1. ``queue`` : EventBus からのメッセージを受信する SimpleMQ キュー
2. 2つの ``process-configuration`` : schedule と trigger が送るメッセージの実行設定
3. ``schedule`` : 直近の開始時刻から1分ごとに実行設定を呼び出すスケジュール
4. ``trigger`` : スイッチ作成イベントで実行設定を呼び出すトリガー
5. ``switch`` : trigger の動作確認用に作成するスイッチ

Step 1. 実行環境の確認
================================================================================

現在の profile と、使用するコマンドのヘルプを確認します。

.. code-block:: shell

   usacloud config current
   usacloud eventbus-api --help
   usacloud simplemq-api --help
   usacloud iaas switch create --help

EventBus の各リソースと、SimpleMQ のメッセージ操作も確認します。

.. code-block:: shell

   usacloud eventbus-api process-configuration create --help
   usacloud eventbus-api process-configuration update-secret --help
   usacloud eventbus-api schedule create --help
   usacloud eventbus-api trigger create --help
   usacloud simplemq-api message receive --help

以降のコマンドは同じシェルで実行します。リソース名が重複しないように実行日時を付け、
スイッチの作成先を指定します。

.. code-block:: shell

   SUFFIX=$(date +%Y%m%d%H%M%S)
   QUEUE_NAME="tutorial-eventbus-${SUFFIX}"
   ZONE=is1c

Step 2. SimpleMQ キューと実行設定の作成
================================================================================

SimpleMQ キューを作成し、メッセージ API 用の API キーを発行します。
API キーを再発行すると古い API キーは無効になるため、このチュートリアル専用の
キューを使用します。

.. code-block:: shell

   usacloud simplemq-api queue create \
     --name "${QUEUE_NAME}" \
     --description "usacloud EventBus API tutorial" \
     -y

   umask 077
   API_KEY_FILE=$(mktemp "${TMPDIR:-/tmp}/eventbus-api-key.XXXXXX")
   SECRET_FILE=$(mktemp "${TMPDIR:-/tmp}/eventbus-secret.XXXXXX")
   SCHEDULE_RESULT=$(mktemp "${TMPDIR:-/tmp}/eventbus-schedule.XXXXXX")
   TRIGGER_RESULT=$(mktemp "${TMPDIR:-/tmp}/eventbus-trigger.XXXXXX")
   trap 'rm -f "${API_KEY_FILE}" "${SECRET_FILE}" "${SCHEDULE_RESULT}" "${TRIGGER_RESULT}"' EXIT

   usacloud simplemq-api queue rotate-api-key "${QUEUE_NAME}" \
     -y --format '{{.APIKey}}' >"${API_KEY_FILE}"
   chmod 600 "${API_KEY_FILE}"

キューを表形式で確認します。

.. code-block:: shell

   usacloud simplemq-api queue read "${QUEUE_NAME}" -o table

schedule と trigger のどちらが実行されたかを区別できるように、異なるメッセージを送る
2つの実行設定を作成します。SimpleMQ のメッセージ本文は Base64 で指定します。

.. code-block:: shell

   SCHEDULE_CONTENT=$(printf 'eventbus schedule fired' | base64 | tr -d '\n')
   TRIGGER_CONTENT=$(printf 'eventbus switch trigger fired' | base64 | tr -d '\n')

   jq -nc \
     --arg queue_name "${QUEUE_NAME}" \
     --arg content "${SCHEDULE_CONTENT}" \
     '{Destination:"simplemq",Parameters:({queue_name:$queue_name,content:$content}|tojson)}' \
     >schedule-process.json

   jq -nc \
     --arg queue_name "${QUEUE_NAME}" \
     --arg content "${TRIGGER_CONTENT}" \
     '{Destination:"simplemq",Parameters:({queue_name:$queue_name,content:$content}|tojson)}' \
     >trigger-process.json

   SCHEDULE_PC_ID=$(usacloud eventbus-api process-configuration create \
     --name "tutorial-eventbus-schedule-${SUFFIX}" \
     --settings schedule-process.json \
     -y --quiet)

   TRIGGER_PC_ID=$(usacloud eventbus-api process-configuration create \
     --name "tutorial-eventbus-trigger-${SUFFIX}" \
     --settings trigger-process.json \
     -y --quiet)

EventBus が SimpleMQ へ送信するときに使う API キーを、両方の実行設定へ登録します。
``update-secret`` は ``Secret`` で包む前の JSON を受け取ります。

.. code-block:: shell

   jq -n --arg api_key "$(cat "${API_KEY_FILE}")" \
     '{APIKey:$api_key}' >"${SECRET_FILE}"

   usacloud eventbus-api process-configuration update-secret \
     "${SCHEDULE_PC_ID}" --secret-file "${SECRET_FILE}" -y
   usacloud eventbus-api process-configuration update-secret \
     "${TRIGGER_PC_ID}" --secret-file "${SECRET_FILE}" -y

   usacloud eventbus-api process-configuration read \
     "${SCHEDULE_PC_ID}" -o table
   usacloud eventbus-api process-configuration read \
     "${TRIGGER_PC_ID}" -o table

Step 3. schedule からメッセージが届くことを確認
================================================================================

現在時刻から2分後の分境界を開始時刻とし、最短周期である1分ごとに実行する schedule を
作成します。``StartsAt`` は Unix epoch ミリ秒で指定します。

.. code-block:: shell

   NOW=$(date +%s)
   START_SECONDS=$(( (NOW / 60 + 2) * 60 ))
   STARTS_AT=$(( START_SECONDS * 1000 ))

   jq -nc \
     --arg process_configuration_id "${SCHEDULE_PC_ID}" \
     --argjson starts_at "${STARTS_AT}" \
     '{
       ProcessConfigurationID:$process_configuration_id,
       StartsAt:$starts_at,
       RecurringStep:1,
       RecurringUnit:"min"
     }' >schedule-settings.json

   SCHEDULE_ID=$(usacloud eventbus-api schedule create \
     --name "tutorial-eventbus-${SUFFIX}" \
     --settings schedule-settings.json \
     -y --quiet)

   usacloud eventbus-api schedule read "${SCHEDULE_ID}" -o table

SimpleMQ のメッセージ数を5秒ごとに確認し、メッセージが届くまで最大5分待ちます。
EventBus のジョブ実行はベストエフォートであり、開始時刻と受信時刻が厳密に一致するとは
限りません。

.. code-block:: shell

   for i in $(seq 1 60); do
     MESSAGE_COUNT=$(usacloud simplemq-api queue count-messages \
       "${QUEUE_NAME}" --format '{{.MessageCount}}')
     if test "${MESSAGE_COUNT}" -gt 0; then
       break
     fi
     sleep 5
   done

   test "${MESSAGE_COUNT}" -gt 0
   usacloud simplemq-api message receive \
     --queue-name "${QUEUE_NAME}" \
     --api-key-file "${API_KEY_FILE}" \
     -y -o json >"${SCHEDULE_RESULT}"

   test "$(jq -r '.[0].Content' "${SCHEDULE_RESULT}")" = "${SCHEDULE_CONTENT}"
   jq -r '.[0].Content' "${SCHEDULE_RESULT}" | base64 --decode
   echo

出力例:

.. code-block:: text

   eventbus schedule fired

受信したメッセージを削除します。また、次の trigger の確認中に schedule が再実行されない
よう、ここで schedule も削除します。

.. code-block:: shell

   SCHEDULE_MESSAGE_ID=$(jq -r '.[0].ID' "${SCHEDULE_RESULT}")
   usacloud simplemq-api message delete \
     --queue-name "${QUEUE_NAME}" \
     --api-key-file "${API_KEY_FILE}" \
     --message-id "${SCHEDULE_MESSAGE_ID}" \
     -y

   usacloud eventbus-api schedule delete "${SCHEDULE_ID}" -y

Step 4. スイッチ作成イベントでメッセージが届くことを確認
================================================================================

イベントログをソースとし、スイッチ作成イベントを検知する trigger を作成します。
このイベントタイプでは ``zone`` 条件を指定せず、タイプだけで絞り込みます。

.. code-block:: shell

   jq -nc \
     --arg process_configuration_id "${TRIGGER_PC_ID}" \
     '{
       Source:"//eventbus.sakura.ad.jp/eventlog",
       Types:["jp.ad.sakura.eventbus.eventlog.IaaS.request.Switch.normal.created"],
       ProcessConfigurationID:$process_configuration_id
     }' >trigger-settings.json

   TRIGGER_ID=$(usacloud eventbus-api trigger create \
     --name "tutorial-eventbus-${SUFFIX}" \
     --settings trigger-settings.json \
     -y --quiet)

   usacloud eventbus-api trigger read "${TRIGGER_ID}" -o table

trigger の反映を60秒待ってから、対象ゾーンにスイッチを作成してイベントを発生させます。

.. code-block:: shell

   sleep 60

   SWITCH_ID=$(usacloud iaas switch create \
     --zone "${ZONE}" \
     --name "tutorial-eventbus-${SUFFIX}" \
     -y --quiet)

SimpleMQ のメッセージ数を5秒ごとに確認し、trigger のメッセージが届くまで最大5分待ちます。

.. code-block:: shell

   for i in $(seq 1 60); do
     MESSAGE_COUNT=$(usacloud simplemq-api queue count-messages \
       "${QUEUE_NAME}" --format '{{.MessageCount}}')
     if test "${MESSAGE_COUNT}" -gt 0; then
       break
     fi
     sleep 5
   done

   test "${MESSAGE_COUNT}" -gt 0
   usacloud simplemq-api message receive \
     --queue-name "${QUEUE_NAME}" \
     --api-key-file "${API_KEY_FILE}" \
     -y -o json >"${TRIGGER_RESULT}"

   test "$(jq -r '.[0].Content' "${TRIGGER_RESULT}")" = "${TRIGGER_CONTENT}"
   jq -r '.[0].Content' "${TRIGGER_RESULT}" | base64 --decode
   echo

出力例:

.. code-block:: text

   eventbus switch trigger fired

これで、スイッチ作成イベントを EventBus が検知し、実行設定を通じて SimpleMQ に
メッセージを送ったことを確認できました。受信したメッセージを削除します。

.. code-block:: shell

   TRIGGER_MESSAGE_ID=$(jq -r '.[0].ID' "${TRIGGER_RESULT}")
   usacloud simplemq-api message delete \
     --queue-name "${QUEUE_NAME}" \
     --api-key-file "${API_KEY_FILE}" \
     --message-id "${TRIGGER_MESSAGE_ID}" \
     -y

Step 5. クリーンアップ
================================================================================

trigger を先に削除してから動作確認用のスイッチを削除します。その後、参照されなくなった
実行設定と SimpleMQ キューを削除します。schedule は Step 3 で削除済みです。

.. code-block:: shell

   usacloud eventbus-api trigger delete "${TRIGGER_ID}" -y
   usacloud iaas switch delete "${SWITCH_ID}" --zone "${ZONE}" -y
   usacloud eventbus-api process-configuration delete \
     "${TRIGGER_PC_ID}" -y
   usacloud eventbus-api process-configuration delete \
     "${SCHEDULE_PC_ID}" -y
   usacloud simplemq-api queue delete "${QUEUE_NAME}" -y

   rm -f schedule-process.json trigger-process.json \
     schedule-settings.json trigger-settings.json
   rm -f "${API_KEY_FILE}" "${SECRET_FILE}" \
     "${SCHEDULE_RESULT}" "${TRIGGER_RESULT}"
   trap - EXIT

作成したリソースが残っていないことを確認します。

.. code-block:: shell

   usacloud eventbus-api process-configuration list -o table
   usacloud eventbus-api schedule list -o table
   usacloud eventbus-api trigger list -o table
   usacloud simplemq-api queue list -o table

.. important::

   途中で手順が失敗した場合も、作成済みの ID を使って trigger、schedule、スイッチ、
   実行設定、SimpleMQ キューを削除してください。API キーを保存した一時ファイルも
   必ず削除してください。

まとめ
================================================================================

本チュートリアルでは ``eventbus-api`` を使い、schedule の時刻到来とスイッチ作成イベントの
それぞれから SimpleMQ にメッセージが届くことを確認しました。EventBus の設定を
``read`` できるだけでなく、ジョブが実際に実行されたことを受信メッセージで検証できます。

詳細は以下の公式資料とコマンドヘルプを参照してください。

- `EventBusの基本情報 <https://manual.sakura.ad.jp/cloud/appliance/eventbus/about.html>`_
- `EventBus コントロールパネル <https://manual.sakura.ad.jp/cloud/appliance/eventbus/control_panel.html>`_
- `EventBus イベント一覧 <https://manual.sakura.ad.jp/cloud/appliance/eventbus/events.html>`_
- `ネットワーク関連のタイプ <https://manual.sakura.ad.jp/cloud/appliance/eventbus/events_network.html>`_
- `シンプルMQの基本情報 <https://manual.sakura.ad.jp/cloud/appliance/simplemq/about.html>`_
- `SimpleMQ API利用の基本手順 <https://manual.sakura.ad.jp/cloud/appliance/simplemq/api.html>`_
- ``usacloud eventbus-api --help``
- ``usacloud simplemq-api --help``
