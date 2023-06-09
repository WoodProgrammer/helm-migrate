#!/usr/bin/env bash

set \
  -o errexit \
  -o nounset \
  -o pipefail

if [ -n "${HELM_MIGRATE_PLUGIN_NO_INSTALL_HOOK:-}" ]; then
    echo "Development mode: not downloading versioned release."
    exit 0
fi

validate_checksum() {
    if ! grep -q ${1} ${2}; then
        echo "Invalid checksum" > /dev/stderr
        exit 1
    fi
    echo "Checksum is valid."
}

on_exit() {
    exit_code=$?
    if [ ${exit_code} -ne 0 ]; then
        echo "helm-migrate install hook failed. Please remove the plugin using 'helm plugin remove migrate' and install again." > /dev/stderr
    fi
    exit ${exit_code}
}
trap on_exit EXIT

version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Downloading and installing helm-migrate ${version} ..."

arch=""
case "$(uname -m)" in
  x86_64|amd64) arch="amd64" ;;
  aarch64|arm64) arch="arm64" ;;
  *)
    echo "Arch '$(uname -m)' not supported!" >&2
    exit 1
    ;;
esac

os=""
case "$(uname)" in
  Darwin) os="darwin" ;;
  Linux) os="linux" ;;
  CYGWIN*|MINGW*|MSYS_NT*) os="windows" ;;
  *)
    echo "OS '$(uname)' not supported!" >&2
    exit 1
    ;;
esac

binary_url="https://github.com/WoodProgrammer/helm-migrate/releases/download/${version}/helm-migrate_${version}_${os}_${arch}.tar.gz"
checksum_url="https://github.com/WoodProgrammer/helm-migrate/releases/download/${version}/helm-migrate_${version}_checksums.txt"

mkdir -p "bin"
mkdir -p "releases/${version}"
binary_filename="releases/${version}.tar.gz"
checksums_filename="releases/${version}_checksums.txt"

# Download binary and checksums files.
(
    if command -v curl >/dev/null 2>&1; then
        curl -sSL "${binary_url}" -o "${binary_filename}"
        curl -sSL "${checksum_url}" -o "${checksums_filename}"
    elif command -v wget >/dev/null 2>&1; then
        wget -q "${binary_url}" -O "${binary_filename}"
        wget -q "${checksum_url}" -O "${checksums_filename}"
    else
      echo "ERROR: no curl or wget found to download files." > /dev/stderr
    fi
)

# Verify checksum.
(
    if command -v sha256sum >/dev/null 2>&1; then
        checksum=$(sha256sum ${binary_filename} | awk '{ print $1 }')
        validate_checksum ${checksum} ${checksums_filename}
    elif command -v openssl >/dev/null 2>&1; then
        checksum=$(openssl dgst -sha256 ${binary_filename} | awk '{ print $2 }')
        validate_checksum ${checksum} ${checksums_filename}
    else
        echo "WARNING: no tool found to verify checksum" > /dev/stderr
    fi
)

# Unpack the binary.
tar xzf "${binary_filename}" -C "releases/${version}"
mv "releases/${version}/bin/helm-migrate" "bin/helm-migrate"
exit 0