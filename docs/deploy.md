## 合约编译

```shell

Docker 映像运行编译器可执行文件，因此您可以将所有编译器参数传递给它。例如，下面的命令会拉取solc镜像的稳定版本（如果您还没有），并在新容器中运行它，并传递--help参数。

docker run ethereum/solc:stable --help
您还可以在标记中指定发布构建版本，例如 0.5.4 版本。

docker run ethereum/solc:0.5.4 --help
要使用 Docker 镜像在主机上编译 Solidity 文件，请挂载本地文件夹用于输入和输出，并指定要编译的合约。例如。

docker run -v /local/path:/sources ethereum/solc:stable -o /sources/output --abi --bin /sources/Contract.sol
您还可以使用标准 JSON 接口（在使用带有工具的编译器时推荐）。使用此接口时，只要 JSON 输入是自包含的（即它不引用必须 由导入回调加载的任何外部文件），就不需要挂载任何目录。

docker run ethereum/solc:stable --standard-json < input.json > output.json

```   

## 合约部署
