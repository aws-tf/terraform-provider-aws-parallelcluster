#!/bin/bash

set -ex

# On Mac OS, the default implementation of sed is BSD sed, but this script requires GNU sed.
if [ "$(uname)" == "Darwin" ]; then
  command -v gsed >/dev/null 2>&1 || { echo >&2 "[ERROR] Mac OS detected: please install GNU sed with 'brew install gnu-sed'"; exit 1; }
  PATH="/usr/local/opt/gnu-sed/libexec/gnubin:$PATH"
fi

_error_exit() {
   echo "$1"
   exit 1
}

_help() {
    local -- _cmd
    _cmd=$(basename "$0")

    cat <<EOF

  Usage: ${_cmd} [OPTION]...

  Bump ParallelCluster Terraform version.

  --version <version>                                               ParallelCluster Terraform version
  -h, --help                                                        Print this help message
EOF
}

main() {
    # parse input options
    while [ $# -gt 0 ] ; do
        case "$1" in
            --version)                            _version="$2"; shift;;
            --version=*)                          _version="${1#*=}";;
            -h|--help|help)                       _help; exit 0;;
            *)                                    _help; _error_exit "[error] Unrecognized option '$1'";;
        esac
        shift
    done

    # verify required parameters
    if [ -z "${_version}" ]; then
        _error_exit "--version parameter not specified"
        _help;
    else
        NEW_VERSION=$_version
        CURRENT_VERSION=$(gsed -ne "s/^VERSION = \(.*\)/\1/p" GNUmakefile)

        sed -i "s/VERSION = $CURRENT_VERSION/VERSION = $NEW_VERSION/g" GNUmakefile
        sed -i "s/Version: $CURRENT_VERSION/Version: $NEW_VERSION/g" THIRD-PARTY-LICENSES.txt
        sed -i "s/version = \"$CURRENT_VERSION\"/version = \"$NEW_VERSION\"/g" */*/*/*/provider.tf
        sed -i "s/version = \"$CURRENT_VERSION\"/version = \"$NEW_VERSION\"/g" */*/provider.tf
    fi
}

main "$@"

# vim:syntax=sh
