#!/bin/bash
# Copyright 2017-2020 The Usacloud Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# 注:ポート番号は固定
docker run -it --rm \
  --name usacloud \
  -e SAKURACLOUD_ACCESS_TOKEN \
  -e SAKURACLOUD_ACCESS_TOKEN_SECRET \
  -e SAKURACLOUD_DEFAULT_ZONE \
  -e SAKURACLOUD_TRACE_MODE \
  usacloud:latest $@