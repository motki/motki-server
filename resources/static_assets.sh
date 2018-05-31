#!/usr/bin/env bash
set -eo pipefail

targets="motki_ddl_zip types_zip icons_zip evesde_bz2"

motki_ddl_zip_name="motki-schema.zip"
types_zip_name="Types.zip"
icons_zip_name="Icons.zip"
evesde_bz2_name="evesde-postgres.dmp.bz2"

build_dir="$PWD/build/assets"

mkdir -p ${build_dir}

static_index_url="https://motki-static.nyc3.digitaloceanspaces.com/evesde/current"
static_index_file="$build_dir/static_index"

curl -s -L ${static_index_url} > ${static_index_file}

for kind in ${targets}; do
    name_var="${kind}_name"
    download_file="${build_dir}/$(echo "${!name_var}")"
    if [ -f ${download_file} ]; then
        echo "Skipping ${kind}, ${download_file} exists"
        continue
    fi

    download_url=$(cat ${static_index_file} | grep ${kind} | cut -d'=' -f2; exit 0)
    if [[ ${download_url} == "" ]]; then
        echo "Error downloading static assets: cannot find ${kind} in current index"
        exit 1
    fi

    echo "Downloading $download_url to $download_file ... "

    curl -s -L ${download_url} > ${download_file}
done