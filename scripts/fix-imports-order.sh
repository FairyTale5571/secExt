#!/usr/bin/env bash
## Check if gogroup is installed
if ! tool_loc="$(type -p gogroup)" || [[ -z ${tool_loc} ]]; then
      echo "gogroup is not installed. installing...."
      go get -u -v github.com/Bubblyworld/gogroup/...
fi

gogroup -order std,other,prefix=fox_telegram --rewrite $(find . -type f -name "*.go" | grep -v /vendor/ |grep -v /.git/)