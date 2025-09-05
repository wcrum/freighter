#!/bin/bash

# Usage:
#   - curl -sfL... | ENV_VAR=... bash
#   - ENV_VAR=... ./install.sh
#
# Install Usage:
#   Install Latest Release
#     - curl -sfL https://get.freighter.dev | bash
#     - ./install.sh
#
#   Install Specific Release
#     - curl -sfL https://get.freighter.dev | FREIGHTER_VERSION=1.0.0 bash
#     - FREIGHTER_VERSION=1.0.0 ./install.sh
#
#   Set Install Directory
#     - curl -sfL https://get.freighter.dev | FREIGHTER_INSTALL_DIR=/usr/local/bin bash
#     - FREIGHTER_INSTALL_DIR=/usr/local/bin ./install.sh
#
#   Set Freighter Directory
#     - curl -sfL https://get.freighter.dev | FREIGHTER_DIR=$HOME/.freighter bash
#     - FREIGHTER_DIR=$HOME/.freighter ./install.sh
#
# Debug Usage:
#   - curl -sfL https://get.freighter.dev | FREIGHTER_DEBUG=true bash
#   - FREIGHTER_DEBUG=true ./install.sh
#
# Uninstall Usage:
#   - curl -sfL https://get.freighter.dev | FREIGHTER_UNINSTALL=true bash
#   - FREIGHTER_UNINSTALL=true ./install.sh
#
# Documentation:
#   - https://freighter.dev
#   - https://github.com/freighter-dev/freighter

# set functions for logging
function verbose {
    echo "$1"
}

function info {
    echo && echo "[INFO] Freighter: $1"
}

function warn {
    echo && echo "[WARN] Freighter: $1"
}

function fatal {
    echo && echo "[ERROR] Freighter: $1"
    exit 1
}

# debug freighter from argument or environment variable
if [ "${FREIGHTER_DEBUG}" = "true" ]; then
    set -x
fi

# start freighter preflight checks
info "Starting Preflight Checks..."

# check for required packages and dependencies
for cmd in echo curl grep sed rm mkdir awk openssl tar install source; do
    if ! command -v "$cmd" &> /dev/null; then
        fatal "$cmd is required to install Freighter"
    fi
done

# set install directory from argument or environment variable
FREIGHTER_INSTALL_DIR=${FREIGHTER_INSTALL_DIR:-/usr/local/bin}

# ensure install directory exists and/or create it
if [ ! -d "${FREIGHTER_INSTALL_DIR}" ]; then
    mkdir -p "${FREIGHTER_INSTALL_DIR}" || fatal "Failed to Create Install Directory: ${FREIGHTER_INSTALL_DIR}"
fi

# ensure install directory is writable (by user or root privileges)
if [ ! -w "${FREIGHTER_INSTALL_DIR}" ]; then
    if [ "$(id -u)" -ne 0 ]; then
        fatal "Root privileges are required to install Freighter to Directory: ${FREIGHTER_INSTALL_DIR}"
    fi
fi

# uninstall freighter from argument or environment variable
if [ "${FREIGHTER_UNINSTALL}" = "true" ]; then
    # remove the freighter binary
    rm -rf "${FREIGHTER_INSTALL_DIR}/freighter" || fatal "Failed to Remove Freighter from ${FREIGHTER_INSTALL_DIR}"

    # remove the freighter directory
    rm -rf "${FREIGHTER_DIR}" || fatal "Failed to Remove Freighter Directory: ${FREIGHTER_DIR}"

    info "Successfully Uninstalled Freighter" && echo
    exit 0
fi

# set version environment variable
if [ -z "${FREIGHTER_VERSION}" ]; then
    # attempt to retrieve the latest version from GitHub
    FREIGHTER_VERSION=$(curl -sI https://github.com/freighter-dev/freighter/releases/latest | grep -i location | sed -e 's#.*tag/v##' -e 's/^[[:space:]]*//g' -e 's/[[:space:]]*$//g')

    # exit if the version could not be detected
    if [ -z "${FREIGHTER_VERSION}" ]; then
        fatal "FREIGHTER_VERSION is unable to be detected and/or retrieved from GitHub. Please set: FREIGHTER_VERSION"
    fi
fi

# detect the operating system
PLATFORM=$(uname -s | tr '[:upper:]' '[:lower:]')
case $PLATFORM in
    linux)
        PLATFORM="linux"
        ;;
    darwin)
        PLATFORM="darwin"
        ;;
    *)
        fatal "Unsupported Platform: ${PLATFORM}"
        ;;
esac

# detect the architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64 | x86-32 | x64 | x32 | amd64)
        ARCH="amd64"
        ;;
    aarch64 | arm64)
        ARCH="arm64"
        ;;
    *)
        fatal "Unsupported Architecture: ${ARCH}"
        ;;
esac

# set freighter directory from argument or environment variable
FREIGHTER_DIR=${FREIGHTER_DIR:-$HOME/.freighter}

# start freighter installation
info "Starting Installation..."

# display the version, platform, and architecture
verbose "- Version: v${FREIGHTER_VERSION}"
verbose "- Platform: ${PLATFORM}"
verbose "- Architecture: ${ARCH}"
verbose "- Install Directory: ${FREIGHTER_INSTALL_DIR}"
verbose "- Freighter Directory: ${FREIGHTER_DIR}"

# ensure freighter directory exists and/or create it
if [ ! -d "${FREIGHTER_DIR}" ]; then
    mkdir -p "${FREIGHTER_DIR}" || fatal "Failed to Create Freighter Directory: ${FREIGHTER_DIR}"
fi

# ensure freighter directory is writable (by user or root privileges)
chmod -R 777 "${FREIGHTER_DIR}" || fatal "Failed to Update Permissions of Freighter Directory: ${FREIGHTER_DIR}"

# change to freighter directory
cd "${FREIGHTER_DIR}" || fatal "Failed to Change Directory to Freighter Directory: ${FREIGHTER_DIR}"

# start freighter artifacts download
info "Starting Download..."

# download the checksum file
if ! curl -sfOL "https://github.com/freighter-dev/freighter/releases/download/v${FREIGHTER_VERSION}/FREIGHTER_${FREIGHTER_VERSION}_checksums.txt"; then
    fatal "Failed to Download: FREIGHTER_${FREIGHTER_VERSION}_checksums.txt"
fi

# download the archive file
if ! curl -sfOL "https://github.com/freighter-dev/freighter/releases/download/v${FREIGHTER_VERSION}/FREIGHTER_${FREIGHTER_VERSION}_${PLATFORM}_${ARCH}.tar.gz"; then
    fatal "Failed to Download: FREIGHTER_${FREIGHTER_VERSION}_${PLATFORM}_${ARCH}.tar.gz"
fi

# start freighter checksum verification
info "Starting Checksum Verification..."

# verify the Freighter checksum
EXPECTED_CHECKSUM=$(awk -v FREIGHTER_VERSION="${FREIGHTER_VERSION}" -v PLATFORM="${PLATFORM}" -v ARCH="${ARCH}" '$2 == "FREIGHTER_"FREIGHTER_VERSION"_"PLATFORM"_"ARCH".tar.gz" {print $1}' "FREIGHTER_${FREIGHTER_VERSION}_checksums.txt")
DETERMINED_CHECKSUM=$(openssl dgst -sha256 "FREIGHTER_${FREIGHTER_VERSION}_${PLATFORM}_${ARCH}.tar.gz" | awk '{print $2}')

if [ -z "${EXPECTED_CHECKSUM}" ]; then
    fatal "Failed to Locate Checksum: FREIGHTER_${FREIGHTER_VERSION}_${PLATFORM}_${ARCH}.tar.gz"
elif [ "${DETERMINED_CHECKSUM}" = "${EXPECTED_CHECKSUM}" ]; then
    verbose "- Expected Checksum: ${EXPECTED_CHECKSUM}"
    verbose "- Determined Checksum: ${DETERMINED_CHECKSUM}"
    verbose "- Successfully Verified Checksum: FREIGHTER_${FREIGHTER_VERSION}_${PLATFORM}_${ARCH}.tar.gz"
else
    verbose "- Expected: ${EXPECTED_CHECKSUM}"
    verbose "- Determined: ${DETERMINED_CHECKSUM}"
    fatal "Failed Checksum Verification: FREIGHTER_${FREIGHTER_VERSION}_${PLATFORM}_${ARCH}.tar.gz"
fi

# uncompress the freighter archive
tar -xzf "FREIGHTER_${FREIGHTER_VERSION}_${PLATFORM}_${ARCH}.tar.gz" || fatal "Failed to Extract: FREIGHTER_${FREIGHTER_VERSION}_${PLATFORM}_${ARCH}.tar.gz"

# install the freighter binary
install -m 755 freighter "${FREIGHTER_INSTALL_DIR}" || fatal "Failed to Install Freighter: ${FREIGHTER_INSTALL_DIR}"

# add freighter to the path
if [[ ":$PATH:" != *":${FREIGHTER_INSTALL_DIR}:"* ]]; then
    if [ -f "$HOME/.bashrc" ]; then
        echo "export PATH=\$PATH:${FREIGHTER_INSTALL_DIR}" >> "$HOME/.bashrc"
        source "$HOME/.bashrc"
    elif [ -f "$HOME/.bash_profile" ]; then
        echo "export PATH=\$PATH:${FREIGHTER_INSTALL_DIR}" >> "$HOME/.bash_profile"
        source "$HOME/.bash_profile"
    elif [ -f "$HOME/.zshrc" ]; then
        echo "export PATH=\$PATH:${FREIGHTER_INSTALL_DIR}" >> "$HOME/.zshrc"
        source "$HOME/.zshrc"
    elif [ -f "$HOME/.profile" ]; then
        echo "export PATH=\$PATH:${FREIGHTER_INSTALL_DIR}" >> "$HOME/.profile"
        source "$HOME/.profile"
    else
        warn "Failed to add ${FREIGHTER_INSTALL_DIR} to PATH: Unsupported Shell"
    fi
fi

# display success message
info "Successfully Installed Freighter at ${FREIGHTER_INSTALL_DIR}/freighter"

# display availability message
info "Freighter v${FREIGHTER_VERSION} is now available for use!"

# display freighter docs message
verbose "- Documentation: https://freighter.dev" && echo
