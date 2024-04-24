#!/bin/bash

# モックを生成するディレクトリを指定
MOCK_DIR="./infrastructure/repository/mocks"

# プロジェクトのルートディレクトリ
PROJECT_ROOT="$(pwd)"

# リポジトリ層のモックを生成
for file in $(find $PROJECT_ROOT/infrastructure/repository -name '*_repository_interface.go'); do
    destination=$(basename $file | sed -e 's _interface.go  g' -e 's $ _mock.go g')

    package_name=$(dirname $file | sed 's .*infrastructure repository/  g' | tr '/' '_')
    interface_name=$(basename token_repository_interface.go | sed -e 's _repository_interface.go  g' -e 's \b\(.\) \u\1 ' -e 's $ Repository g')

    simple_file=$(echo $file | sed -e 's .*/server  g')
    echo -e "created at \033[0;35m$MOCK_DIR/$destination \033[0mfrom \033[0;35m$simple_file\033[0m"
    mkdir -p $MOCK_DIR

    mockgen -source=$file -destination=$MOCK_DIR/$destination -package=mocks $package_name $interface_name
done

echo -e "created at \033[0;35m$MOCK_DIR/$destination \033[0mfrom \033[0;35m$simple_file\033[0m"
# db層のモックを生成
mockgen -source infrastructure/db/store.go -destination infrastructure/db/mocks/store_mock.go
