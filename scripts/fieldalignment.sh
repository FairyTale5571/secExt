#!/usr/bin/env bash
## Check if gogroup is installed
if ! tool_loc="$(type -p fieldalignment)" || [[ -z ${tool_loc} ]]; then
      echo "fieldalignment is not installed. installing...."
      go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
fi

fieldalignment -fix ./...
