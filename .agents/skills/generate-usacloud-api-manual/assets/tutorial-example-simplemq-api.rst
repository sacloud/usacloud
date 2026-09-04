================================================================================
simplemq-api チュートリアル
================================================================================

``usacloud simplemq-api`` を使うと、ソフトウェアコンポーネント間でデータを非同期に
受け渡すためのキューを作成し、メッセージを送受信できます。
例えば、時間のかかるメール送信処理をキューへ登録して別の処理系に任せることで、
リクエスト元へ先に応答したい場合に利用できます。

.. warning::

   このチュートリアルでは **本番環境にリソースを作成・削除** します。
   作業を始める前に、以下のコマンドで現在の profile を確認し、
   対象プロファイルが意図したものであることを確認してください。

   .. code-block:: shell

      usacloud config current

   意図しないプロファイルやプロジェクトで作業しないよう、必ず確認を行ってください。

   キュー管理 API を実行する profile には「作成・削除」のアクセスレベルが必要です。

このサービスはいつ使うか
================================================================================

SimpleMQ は、送信側と受信側を同期させずにソフトウェア間でデータを受け渡したい場合に
利用するマネージド型のメッセージキューサービスです。
配信方式は Pull 型、配信保証は at least once です。

前提条件
================================================================================

- usacloud と jq がインストールされていること
- 対象 profile が設定済みで、キュー管理に必要な「作成・削除」のアクセスレベルを持つこと
- API キーを保存するローカル一時ファイルを安全に扱えること
- SimpleMQ はグローバルリソースであり、全ネットワークゾーンで共通であることを理解していること

外部の依存サービスや既存リソースは使用しません。
メッセージ API の認証にはキューごとに発行される API キーを使用します。

このチュートリアルで作成するもの
================================================================================

以下の順序で作成・確認します。

1. ``queue`` : メッセージを保存するキュー
2. キュー認証用 API キー : メッセージ API の呼び出しに使用する認証情報
3. ``message`` : 動作確認用の payload

クリーンアップでは、処理済みメッセージ、キューの順に削除し、API キーの一時ファイルも
削除します。

Step 1. 実行環境の確認
================================================================================

まず、現在の profile と ``simplemq-api`` コマンドのヘルプを確認します。

.. code-block:: shell

   usacloud config current
   usacloud simplemq-api --help

プロファイルが意図したものであることを確認したら、
各 subcommand のヘルプを参照し、必要なパラメータを把握します。

.. code-block:: shell

   usacloud simplemq-api queue create --help
   usacloud simplemq-api queue rotate-api-key --help
   usacloud simplemq-api message send --help
   usacloud simplemq-api message receive --help
   usacloud simplemq-api message delete --help

Step 2. キューと API キーの作成
================================================================================

キュー名は同一プロジェクト内で一意である必要があり、作成後は変更できません。
この例では実行日時を付けて一意な名前を作ります。

.. code-block:: shell

   QUEUE_NAME="tutorial-simplemq-$(date +%Y%m%d%H%M%S)"
   QUEUE_ID=$(usacloud simplemq-api queue create \
     --name "${QUEUE_NAME}" \
     --description "usacloud SimpleMQ API tutorial" \
     -y --quiet)

作成したキューを表形式で確認します。

.. code-block:: shell

   usacloud simplemq-api queue read "${QUEUE_ID}" \
     -o table

出力例:

.. code-block:: text

   +------------+---------------------------+---------------------------------+--------------------------+---------------+------+--------------+
   |     ID     |           Name            |           Description           | VisibilityTimeoutSeconds | ExpireSeconds | Tags | Availability |
   +------------+---------------------------+---------------------------------+--------------------------+---------------+------+--------------+
   | <queue-id> | tutorial-simplemq-...     | usacloud SimpleMQ API tutorial  | 30                       | 345600        | []   | available    |
   +------------+---------------------------+---------------------------------+--------------------------+---------------+------+--------------+

メッセージ API 用の API キーを発行します。API キーの再発行は古い API キーを無効に
するため、既存キューで実行する場合は利用中のクライアントへの影響を確認してください。

.. code-block:: shell

   umask 077
   API_KEY_FILE=$(mktemp "${TMPDIR:-/tmp}/simplemq-api-key.XXXXXX")
   RECEIVE_FILE=$(mktemp "${TMPDIR:-/tmp}/simplemq-message.XXXXXX")
   trap 'rm -f "${API_KEY_FILE}" "${RECEIVE_FILE}"' EXIT

   usacloud simplemq-api queue rotate-api-key "${QUEUE_ID}" \
     -y --format '{{.APIKey}}' >"${API_KEY_FILE}"
   chmod 600 "${API_KEY_FILE}"

Step 3. メッセージの送受信
================================================================================

``Hello`` という payload を送信し、返されたメッセージ ID を保存します。
payload を常に Base64 エンコードする必要はありません。

.. code-block:: shell

   CONTENT='Hello'
   MESSAGE_ID=$(usacloud simplemq-api message send \
     --queue-name "${QUEUE_NAME}" \
     --api-key-file "${API_KEY_FILE}" \
     --content "${CONTENT}" \
     -y --format '{{.ID}}')

キューからメッセージを受信し、結果の JSON を表示します。

.. code-block:: shell

   usacloud simplemq-api message receive \
     --queue-name "${QUEUE_NAME}" \
     --api-key-file "${API_KEY_FILE}" \
     -y -o json >"${RECEIVE_FILE}"

   cat "${RECEIVE_FILE}"

出力例:

.. code-block:: json

   [
       {
           "AcquiredAt": 1788134400000,
           "Content": "Hello",
           "CreatedAt": 1788134400000,
           "ExpiresAt": 1788479999000,
           "ID": "00000000-0000-0000-0000-000000000000",
           "UpdatedAt": 1788134400000,
           "VisibilityTimeoutAt": 1788134430000
       }
   ]

JSON の ``Content`` から payload がそのまま返っていることを確認できます。
送信時と同じ ID と payload であることも確認します。

.. code-block:: shell

   RECEIVED_ID=$(jq -r '.[0].ID' "${RECEIVE_FILE}")
   RECEIVED_CONTENT=$(jq -r '.[0].Content' "${RECEIVE_FILE}")
   test "${RECEIVED_ID}" = "${MESSAGE_ID}"
   test "${RECEIVED_CONTENT}" = "${CONTENT}"

受信すると、そのメッセージは可視性タイムアウトまで他の受信リクエストへは配信されません。
処理が完了したらメッセージを削除し、キュー内のメッセージ数が 0 になったことを確認します。

.. code-block:: shell

   usacloud simplemq-api message delete \
     --queue-name "${QUEUE_NAME}" \
     --api-key-file "${API_KEY_FILE}" \
     --message-id "${MESSAGE_ID}" \
     -y

   usacloud simplemq-api queue count-messages "${QUEUE_ID}" \
     -o table

出力例:

.. code-block:: text

   +--------------+
   | MessageCount |
   +--------------+
   | 0            |
   +--------------+

.. note::

   メッセージ処理が可視性タイムアウトを超える場合は、
   ``usacloud simplemq-api message extend-timeout`` でタイムアウトを延長できます。

Step 4. クリーンアップ
================================================================================

作成したキューを削除し、API キーと受信結果の一時ファイルも削除します。

.. code-block:: shell

   usacloud simplemq-api queue delete "${QUEUE_ID}" \
     -y

   rm -f "${API_KEY_FILE}" "${RECEIVE_FILE}"
   trap - EXIT

キュー一覧を表形式で表示し、作成したキューが残っていないことを確認します。

.. code-block:: shell

   usacloud simplemq-api queue list -o table

.. important::

   途中で手順が失敗した場合も、``QUEUE_ID`` を使ってキューを削除してください。
   API キーの一時ファイルも必ず削除してください。

まとめ
================================================================================

本チュートリアルでは ``simplemq-api`` を使い、キューの作成、API キーの発行、
payload の送受信、処理済みメッセージの削除、キューの削除までを確認しました。
実際の受信処理では at least once 配信を前提とし、処理完了後の削除や、必要に応じた
可視性タイムアウトの延長を実装してください。

詳細は以下の公式資料とコマンドヘルプを参照してください。

- `シンプルMQの基本情報 <https://manual.sakura.ad.jp/cloud/appliance/simplemq/about.html>`_
- `コントロールパネルでの操作 <https://manual.sakura.ad.jp/cloud/appliance/simplemq/control_panel.html>`_
- `API利用の基本手順 <https://manual.sakura.ad.jp/cloud/appliance/simplemq/api.html>`_
- `SimpleMQ APIドキュメント <https://manual.sakura.ad.jp/api/cloud/simplemq/>`_
- ``usacloud simplemq-api queue --help``
- ``usacloud simplemq-api message --help``
