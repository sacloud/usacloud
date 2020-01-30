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


travis encrypt GPG_PASSPHRASE=$GPG_PASSPHRASE -a
travis encrypt GITHUB_TOKEN=$GITHUB_TOKEN -a
travis encrypt SACLOUD_OJS_ACCESS_KEY_ID=$SACLOUD_OJS_ACCESS_KEY_ID -a
travis encrypt SACLOUD_OJS_SECRET_ACCESS_KEY=$SACLOUD_OJS_SECRET_ACCESS_KEY -a
travis encrypt SAKURACLOUD_ACCESS_TOKEN=$SAKURACLOUD_ACCESS_TOKEN -a
travis encrypt SAKURACLOUD_ACCESS_TOKEN_SECRET=$SAKURACLOUD_ACCESS_TOKEN_SECRET -a