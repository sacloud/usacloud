#!/bin/bash

VERSION=`git log --merges --oneline | perl -ne 'if(m/^.+Merge pull request \#[0-9]+ from .+\/bump-version-([0-9\.]+)/){print $1;exit}'`
SHA256_SRC=`openssl dgst -sha256 bin/usacloud_darwin-amd64.zip | awk '{print $2}'`
SHA256_BASH_COMP=`openssl dgst -sha256 contrib/completion/bash/usacloud | awk '{print $2}'`

# clone
git clone --depth=50 --branch=master https://github.com/sacloud/homebrew-usacloud.git homebrew-usacloud
cd homebrew-usacloud

# check version
CURRENT_VERSION=`git log --oneline | perl -ne 'if(/^.+ v([0-9\.]+)/){print $1;exit}'`
if [ "$CURRENT_VERSION" = "$VERSION" ] ; then
    exit 0
fi

cat << EOL > usacloud.rb
class Usacloud < Formula

  usacloud_version = "${VERSION}"
  sha256_src = "${SHA256_SRC}"
  sha256_bash_completion = "${SHA256_BASH_COMP}"

  desc "Unofficial 'sacloud' - CLI client of the SakuraCloud"
  homepage "https://github.com/sacloud/usacloud"
  url "https://github.com/sacloud/usacloud/releases/download/v#{usacloud_version}/usacloud_darwin-amd64.zip"
  sha256 sha256_src
  head "https://github.com/sacloud/usacloud.git"
  version usacloud_version

  option "without-completions", "Disable bash completions"
  resource "bash_completion" do
    url "https://usacloud.b.sakurastorage.jp/contrib/completion/bash/usacloud"
    sha256 sha256_bash_completion
  end

  def install
    bin.install "usacloud"
    if build.with? "completions"
      resource("bash_completion").stage {
        bash_completion.install "usacloud"
      }
    end

  end

  test do
    assert_match "SAKURACLOUD_ACCESS_TOKEN", shell_output("usacloud --help")
  end
end
EOL

git config --global push.default matching
git config user.email 'yamamoto.febc@gmail.com'
git config user.name 'usacloud'
git commit -am "v${VERSION}"

echo "Push ${VERSION} to github.com/sacloud/homebrew-usacloud.git"
git push --quiet -u "https://${GITHUB_TOKEN}@github.com/sacloud/homebrew-usacloud.git" >& /dev/null

echo "Cleanup tag v${VERSION} on github.com/sacloud/homebrew-usacloud.git"
git push --quiet -u "https://${GITHUB_TOKEN}@github.com/sacloud/homebrew-usacloud.git" :v${VERSION} >& /dev/null

echo "Tagging v${VERSION} on github.com/sacloud/homebrew-usacloud.git"
git tag v${VERSION} 2>&1 >/dev/null
git push --quiet -u "https://${GITHUB_TOKEN}@github.com/sacloud/homebrew-usacloud.git" v${VERSION} >& /dev/null
exit 0
