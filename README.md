<h1 align="center">
  <br>she<br>
</h1>

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

## 功能

- 根据定义文件自动对csv数据构建索引
- 快速并行的在多个数据源中进行查询
- 自动监视文件夹并对其中的文本库建立索引

## 安装

### 从源码编译

```sh
$ go get -u -v github.com/zu1k/she
```

需要下载字典

### 下载预编译的可执行程序

[release](https://github.com/zu1k/she/releases)

### Docker

使用预编译的Docker镜像

```sh
docker pull zu1k/she
docker run -itd -p 80:10086 -v /path/to/she/data:/root/she zu1k/she serve -m auto
```

然后将你的各种文本库扔 /path/to/she/data/origin 目录下，然后等待索引就行了

## 说明

Mit协议开源

课程作业，未完工，请勿用来做违法事情
