# CacheflySSLUpDate

[中文版本](/README_CN.md) | English Version

CacheflySSLUpDate is a Go language project designed to facilitate the process of uploading local SSL certificates to Cachefly. Below are instructions on how to use the tool and configure it properly:

## Usage

1. First, download the executable file for your device from [https://github.com/biliblihuorong/CacheflySSLUpDate/releases](https://github.com/biliblihuorong/CacheflySSLUpDate/releases).

2. Create a file named `config.json` in the same directory as the executable and fill it with the required information following the example below. Alternatively, you can download the `config.json` file directly from this repository and fill it with your own configuration details.

```json
{
  "token": "Your Cachefly Token, which can be obtained from https://portal.cachefly.com/app/tokens",
  "certificateFile": "<Certificate file path>",
  "certificateKeyFile": "<Certificate key file path>",
  "password": "<Your Cachefly Password>",
  "taskInterval": 7
}
```

- `token`: Obtain your Cachefly Token from [https://portal.cachefly.com/app/tokens](https://portal.cachefly.com/app/tokens) and enter it here.
- `certificateFile` and `certificateKeyFile`: Provide the file paths to your local SSL certificate and key, respectively. If you are using the Baota panel for automatic certificate issuance, the paths are usually located at `/www/server/panel/vhost/ssl/<your_domain>/`, with the public key named `fullchain.pem` and the private key named `privkey.pem`.
- `password`: Enter your Cachefly password required for certificate upload.
- `taskInterval`: This field defines the interval in days for executing the upload task. For example, setting it to 7 will perform the upload task every 7 days.

**Note:** Ensure that the `config.json` file is properly formatted; otherwise, the tool will not function correctly.

## Additional Information

- This tool simplifies the process of uploading local SSL certificates to Cachefly, making it convenient to update certificates and maintain website security.
- If you encounter any issues or have questions, please feel free to submit them on the [Issues](https://github.com/biliblihuorong/CacheflySSLUpDate/issues) page of this project.

Thank you for using CacheflySSLUpDate! If you have any suggestions for improvement or optimizations, we welcome code contributions and feedback.