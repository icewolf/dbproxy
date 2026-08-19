name := 'dbproxy'
version := `cat VERSION.txt`

_build-setup target:
    #!/usr/bin/env bash
    set -euo pipefail
    if [[ "{{ target }}" =~ ^win.* ]]; then
      echo 'export EXT=".exe"'
      echo 'export GOOS="windows"'
      echo 'export CC="zig cc -target x86_64-windows-gnu"'
    elif [[ "{{ target }}" =~ ^lin.* ]]; then
      echo 'export EXT=""'
      echo 'export GOOS="linux"'
      echo 'export CC="zig cc -target native-native-musl"'
    else
      echo "unknown target '{{ target }}', expected win* or lin*" >&2
      exit 1
    fi
    echo 'export GOARCH="amd64"'
    echo 'export CGO_ENABLED="1"'
    echo 'export CXX="$CC"'

build target='windows':
    #!/usr/bin/env bash
    set -euo pipefail
    setup="$(just _build-setup {{ target }})"
    eval "$setup"
    echo "building for $GOOS"
    go build -o "bin/{{ name }}$EXT" ./cmd/dbproxy/

build-release target='windows':
    #!/usr/bin/env bash
    set -euo pipefail
    setup="$(just _build-setup {{ target }})"
    eval "$setup"
    echo "building release for $GOOS"
    go build -trimpath -ldflags "-s -w -X main.version={{ version }}" -o "bin/{{ name }}$EXT" ./cmd/dbproxy/
