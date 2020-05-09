<h1 align="center">
  <br>she<br>
</h1>

<h4 align="center">Something</h4>

<p align="center">
  <a href="https://github.com/zu1k/she/actions">
    <img src="https://img.shields.io/github/workflow/status/zu1k/she/Go?style=flat-square" alt="Github Actions">
  </a>
  <a href="https://goreportcard.com/report/github.com/zu1k/she">
    <img src="https://goreportcard.com/badge/github.com/zu1k/she?style=flat-square">
  </a>
  <a href="https://github.com/zu1k/she/releases">
    <img src="https://img.shields.io/github/release/zu1k/she/all.svg?style=flat-square">
  </a>
</p>

## Features

- 根据定义文件自动对csv数据构建索引
- 快速并行的在多个数据源中进行查询


## Source

- QQGroup 99G 已导入mssql数据库，有索引
- 163 51G 还未导入数据库
- 12306 200M 遍历搜索，极快
- 2000W酒店数据 3G 使用bleve+sego分词索引，索引膨胀率爆炸，但是查询速度快


## Install

she Requires Go >= 1.13. You can build it from source:

```sh
$ go get -u -v github.com/zu1k/she
```

Download dictionary.txt to the same dir with she

Pre-built binaries are available here: [release](https://github.com/zu1k/she/releases)

## Use

### Docker

使用预编译的Docker镜像

```sh
docker pull zu1k/she
docker run -itd -p 80:10086 -v /path/to/she/data:/root/she zu1k/she serve -m auto
```

然后将你的各种文本库扔 /path/to/she/data/origin 目录下，然后等待索引就行了
