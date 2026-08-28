================================================================================
eventbus-api チュートリアル
================================================================================

`usacloud eventbus-api` を使うと、CLI から EventBus のリソースを作成・確認・削除できます。
EventBus はイベント検知とジョブスケジュールを組み合わせたマネージドサービスです。
スケジュールに基づいてジョブを実行したり、イベントソースの変化を検知してジョブを実行したりできます。
実行するジョブとしては Simple Notification、SimpleMQ、AutoScale の呼び出しが利用できます。

.. warning::

   このチュートリアルでは **本番環境にリソースを作成・削除** します。
   作業を始める前に、以下のコマンドで現在の profile を確認し、
   対象プロファイルが意図したものであることを確認してください。

   .. code-block:: shell

      usacloud config current

   意図しないプロファイルやプロジェクトで作業しないよう、必ず確認を行ってください。

このサービスはいつ使うか
================================================================================

EventBus は「決まった時刻にジョブを実行したい」「特定のイベントが発生したときにジョブを実行したい」
という場面で利用します。例えば、監視アラートを検知して AutoScale を実行したり、
サーバーの起動イベントを検知してロードバランサーにサーバーを追加したりできます。

詳細は `さくらのクラウド マニュアル — EventBus <https://manual.sakura.ad.jp/cloud/appliance/eventbus/about.html>`_ を参照してください。

前提条件
================================================================================

- usacloud がインストールされていること
- 対象プロファイルが設定済みであること
- EventBus のジョブ先（Simple Notification、SimpleMQ、AutoScale など）が必要に応じて利用可能であること
- 本チュートリアルで作成するリソースの料金と影響を理解していること

このチュートリアルで作成するもの
================================================================================

本チュートリアルでは以下のリソースを作成します。
作成順序は依存関係を考慮したものです。

1. ``process-configuration`` : 実行するジョブの内容を定義します。
2. ``schedule`` : 作成した ``process-configuration`` を指定時刻に呼び出します。
3. ``trigger`` : 作成した ``process-configuration`` をイベント発生時に呼び出します。

Step 1. 実行環境の確認
================================================================================

まず、現在の profile と ``eventbus-api`` コマンドのヘルプを確認します。

.. code-block:: shell

   usacloud config current
   usacloud eventbus-api --help

プロファイルが意図したものであることを確認したら、
各 subcommand の ``create --help`` を参照し、必要なパラメータを把握します。

.. code-block:: shell

   usacloud eventbus-api process-configuration create --help
   usacloud eventbus-api schedule create --help
   usacloud eventbus-api trigger create --help

Step 2. process-configuration の作成
================================================================================

process-configuration は、EventBus が実行するジョブの内容を定義します。
以下の例では SimpleMQ にメッセージを送信するジョブを定義しています。
実際の環境では、利用する destination サービスに応じて ``Parameters`` を変更してください。

.. code-block:: shell

   PC_ID=$(usacloud eventbus-api process-configuration create \
     --name usacloud-tutorial-pc \
     --settings '{"Destination":"simplemq","Parameters":"{\"queue_name\":\"my-queue\",\"content\":\"hello\"}"}' \
     -y --quiet)

   echo "created process-configuration: ${PC_ID}"

``--quiet`` を付けると、作成されたリソースの ID のみが出力されます。

実行例:

.. code-block:: text

   <process-configuration-id>
後続の ``schedule`` と ``trigger`` はこの ID を参照します。

Step 3. schedule の作成
================================================================================

schedule は、既存の ``process-configuration`` を指定した間隔または crontab で呼び出します。
以下の例では、2030年1月1日以降、1日ごとに実行するスケジュールを作成します。
実際に試す場合は ``StartsAt`` を適切な未来の時刻に変更してください。

.. code-block:: shell

   SCHEDULE_ID=$(usacloud eventbus-api schedule create \
     --name usacloud-tutorial-schedule \
     --settings "{\"ProcessConfigurationID\":\"${PC_ID}\",\"StartsAt\":<starts-at-millis>,\"RecurringStep\":1,\"RecurringUnit\":\"day\"}" \
     -y --quiet)

   echo "created schedule: ${SCHEDULE_ID}"

``<starts-at-millis>`` は実行させたくない未来の時刻の Unix epoch ミリ秒に置き換えてください。

実行例:

.. code-block:: text

   <schedule-id>

.. note::

   ``StartsAt`` は Unix epoch ミリ秒で指定します。
   ``RecurringStep`` と ``RecurringUnit`` の代わりに ``Crontab`` を指定することもできます。

Step 4. trigger の作成
================================================================================

trigger は、イベントを検知して ``process-configuration`` を呼び出します。
以下の例では架空のイベントソースとイベントタイプを指定しています。
実際に利用する場合は、監視アラートやイベントログなどの有効な source と types を指定してください。

.. code-block:: shell

   TRIGGER_ID=$(usacloud eventbus-api trigger create \
     --name usacloud-tutorial-trigger \
     --settings "{\"Source\":\"//example/source\",\"Types\":[\"example.created\"],\"ProcessConfigurationID\":\"${PC_ID}\"}" \
     -y --quiet)

   echo "created trigger: ${TRIGGER_ID}"

実行例:

.. code-block:: text

   <trigger-id>

Step 5. 動作確認
================================================================================

作成したリソースを ``list`` または ``read`` で確認します。

.. code-block:: shell

   usacloud eventbus-api process-configuration read "${PC_ID}" -o table
   usacloud eventbus-api schedule read "${SCHEDULE_ID}" -o table
   usacloud eventbus-api trigger read "${TRIGGER_ID}" -o table

``-o table`` を付けると、見やすい表形式で出力されます。

process-configuration の読み出し例:

.. code-block:: text

   +----------------------+---------------------------+------+-------------+---------------------------------------------+
   |         ID           |           Name            | Tags | Destination |                 Parameters                  |
   +----------------------+---------------------------+------+-------------+---------------------------------------------+
   | <process-config-id>  | usacloud-tutorial-pc      | -    | simplemq    | {"queue_name":"my-queue","content":"hello"} |
   +----------------------+---------------------------+------+-------------+---------------------------------------------+

schedule の読み出し例:

.. code-block:: text

   +---------------+----------------------------+------+------------------------+---------------+---------------+---------+----------+
   |      ID       |            Name            | Tags | ProcessConfigurationID | RecurringStep | RecurringUnit | Crontab | StartsAt |
   +---------------+----------------------------+------+------------------------+---------------+---------------+---------+----------+
   | <schedule-id> | usacloud-tutorial-schedule | -    | <process-config-id>    | 1             | day           | -       | 0        |
   +---------------+----------------------------+------+------------------------+---------------+---------------+---------+----------+

trigger の読み出し例:

.. code-block:: text

   +--------------+---------------------------+------+------------------+-----------------+------------------------+
   |      ID      |           Name            | Tags |      Source      |      Types      | ProcessConfigurationID |
   +--------------+---------------------------+------+------------------+-----------------+------------------------+
   | <trigger-id> | usacloud-tutorial-trigger | -    | //example/source | example.created | <process-config-id>    |
   +--------------+---------------------------+------+------------------+-----------------+------------------------+

作成時に指定した内容が各フィールドに含まれていることを確認してください。

Step 6. クリーンアップ
================================================================================

チュートリアルで作成したリソースを削除します。
削除順は作成順の逆で行い、参照関係を考慮してください。

.. code-block:: shell

   usacloud eventbus-api trigger delete "${TRIGGER_ID}" -y
   usacloud eventbus-api schedule delete "${SCHEDULE_ID}" -y
   usacloud eventbus-api process-configuration delete "${PC_ID}" -y

.. important::

   リソースが残ったままだと課金が継続する場合があります。
   必ず ``list`` で残りがないことを確認してください。

   .. code-block:: shell

      usacloud eventbus-api process-configuration list -o table
      usacloud eventbus-api schedule list -o table
      usacloud eventbus-api trigger list -o table

まとめ
================================================================================

本チュートリアルでは ``eventbus-api`` を使って、
ジョブ定義（``process-configuration``）と、それを呼び出す ``schedule`` / ``trigger`` を作成・削除しました。
各コマンドの詳細なパラメータについては ``--help`` を、
サービスの詳細については以下の公式マニュアルを参照してください。

- `EventBus 概要 <https://manual.sakura.ad.jp/cloud/appliance/eventbus/about.html>`_
- `EventBus コントロールパネル <https://manual.sakura.ad.jp/cloud/appliance/eventbus/control_panel.html>`_
- `EventBus イベント一覧 <https://manual.sakura.ad.jp/cloud/appliance/eventbus/events.html>`_
