# CacheflySSLUpDate

CacheflySSLUpDate 是一个使用 Go 语言编写的工具，旨在帮助将本地 SSL 证书上传到 Cachefly 服务中。以下是使用该工具的方式及配置说明：

## 使用方式

1. 首先，您需要在 [https://github.com/biliblihuorong/CacheflySSLUpDate/releases](https://github.com/biliblihuorong/CacheflySSLUpDate/releases) 页面下载适用于您设备的可执行文件。

2. 然后，在执行文件所在目录下创建一个名为 `config.json` 的文件，并按照以下示例填入相应的内容。您也可以直接下载本仓库中的 `config.json` 文件，填入您的配置信息。

```json
{
  "token": "您的 Cachefly Token，可在 https://portal.cachefly.com/app/tokens 获取",
  "certificateFile": "<证书文件路径>",
  "certificateKeyFile": "<密钥文件路径>",
  "password": "<您的 Cachefly 密码>",
  "taskInterval": 7
}
```

- `token`：您需要前往 [https://portal.cachefly.com/app/tokens](https://portal.cachefly.com/app/tokens) 获取 Cachefly Token，并将其填写在此处。
- `certificateFile` 和 `certificateKeyFile`：这两个字段应填入您本地 SSL 证书和密钥的文件路径。如果您使用宝塔面板自动签发证书，则签发地址通常在 `/www/server/panel/vhost/ssl/<您的域名>/` 下，其中公钥为 `fullchain.pem`，私钥为 `privkey.pem`。
- `password`：填入您的 Cachefly 密码，用于上传证书到 Cachefly。
- `taskInterval`：此字段定义了执行任务的间隔天数。例如，设置为 7 表示每隔 7 天执行一次上传任务。

**注意：** 请确保 `config.json` 文件的格式正确，否则工具将无法正常运行。

## 其他说明

- 本工具用于简化将本地 SSL 证书上传到 Cachefly 的流程，方便更新证书，以维护网站的安全性。
- 如果您遇到任何问题或有其他疑问，请随时提交问题到本项目的 [Issues](https://github.com/biliblihuorong/CacheflySSLUpDate/issues) 页面。

谢谢您使用 CacheflySSLUpDate！如果有任何改进或优化的建议，欢迎贡献代码或提供反馈。