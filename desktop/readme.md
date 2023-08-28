# 🛠️ OPENCNC-J-DESKTOP

⚡ Desktop Application for LinuxCNC. ⚡

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Install

Please make sure that you have installed the GoLang and Node development environments, ensuring that the Go version is 1.18 or higher and the Node version is 18.13 or higher

```shell
# Query Go version
go version
# Query Node version
node --version
```

```shell
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## Live Development

You can run your application in development mode by running `wails dev` from your project directory

```
wails dev
```

## Building

Build the installation package for the application using the following command script

```[readme.md](readme.md)
wails build -webview2 embed -nsis
```























