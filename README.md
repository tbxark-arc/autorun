# AutoRun

`autorun` is a tool that listens for file changes and automatically executes specified commands. You can use `autorun` to automatically compile and run code or update dependencies.

`autorun` 是一个监听文件变化并自动允许指定命令的工具，你可以使用`autorun`自动编译运行代码，或者再依赖配置文件发生变化时自动刷新依赖。

## Usage

```
Usage of autorun:
  -c string
        Config file path (default "autorun.config")
  -d string
        Distance dir or file path (default ".")
```

## Example

#### iOS project
```json

{
  "build": [],
  "run": {
    "name": "/usr/local/bin/pod",
    "args": [
      "update"
    ]
  },
  "include": {
    "import": [
    ],
    "pattern": [
      "Podfile"
    ]
  },
  "exclude": {
    "import": [
      ".gitignore"
    ],
    "pattern": [
      ".temp",
      ".gitignore"
    ]
  }
}

```

## Install

```shell
go install github.com/tbxark/autorun
```

## Author

tbxark, tbxark@outlook.com

## License

FlexLayout is available under the MIT license. See the LICENSE file for more info.
