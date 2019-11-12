#!/bin/bash
# Copyright 2017-2019 The Usacloud Authors
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


VERSION=`git log --merges --oneline | perl -ne 'if(m/^.+Merge pull request \#[0-9]+ from .+\/bump-version-([0-9\.]+)/){print $1;exit}'`

# clone
git clone --depth=50 --branch=master https://github.com/sacloud/releases-usacloud.git releases-usacloud
cd releases-usacloud
git fetch origin

# check version
CURRENT_VERSION=`git tag -l --sort=-v:refname | perl -ne 'if(/^([0-9\.]+)$/){print $1;exit}'`
if [ "$CURRENT_VERSION" = "$VERSION" ] ; then
    echo "sacloud/releases-usacloud v$VERSION is already released."
    exit 0
fi

# build website static contents
rm -rf bin/
cp -r ../bin ./

rm -rf contrib/
cp -r ../contrib ./

rm -rf repos/
cp -r ../repos ./

cat << EOL > status.html
OK(current version: v${VERSION})
EOL

# commit and push to github.com
git config --global push.default matching
git config user.email 'sacloud.users@gmail.com'
git config user.name 'sacloud-bot'
git add .
git commit -m "v${VERSION}"
git tag "${VERSION}"

echo "Push ${VERSION} to github.com/sacloud/releases-usacloud.git"
git push --quiet -u "https://${GITHUB_TOKEN}@github.com/sacloud/releases-usacloud.git" >& /dev/null

echo "Cleanup tag ${VERSION} on github.com/sacloud/releases-usacloud.git"
git push --quiet -u "https://${GITHUB_TOKEN}@github.com/sacloud/releases-usacloud.git" :${VERSION} >& /dev/null

echo "Tagging ${VERSION} on github.com/sacloud/releases-usacloud.git"
git push --quiet -u "https://${GITHUB_TOKEN}@github.com/sacloud/releases-usacloud.git" ${VERSION} >& /dev/null
exit 0
