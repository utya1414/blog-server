#!/bin/bash

# モックを生成するディレクトリを指定
MOCK_DIR="./infrastructure/repository/mocks"

# プロジェクトのルートディレクトリ
PROJECT_ROOT="$(pwd)"

# リポジトリ層のモックを生成
for file in $(find $PROJECT_ROOT/domain -name '*_repository_interface.go'); do
    destination=$(basename $file | sed -e 's/_interface.go//' -e 's/$/_mock.go/')

    package_name=$(dirname $file | sed 's/.*domain\///' | tr '/' '_')
    interface_name=$(basename token_repository_interface.go | sed -e 's/_repository_interface.go//' -e 's/\b\(.\)/\u\1/' -e 's/$/Repository/')

    mkdir -p $MOCK_DIR
    mockgen -source=$file -destination=$MOCK_DIR/$destination -package=mocks $package_name $interface_name
done

# db層のモックを生成
mockgen -source infrastructure/db/store.go -destination infrastructure/db/mocks/store_mock.go
