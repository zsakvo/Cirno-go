
<h1 align="center">
  <img src="./assets/cirno.png" alt="Cirno" width="150">
  <br>Cirno-go<br>
</h1>

<h4 align="center">A tool for downloading books from <a href="https://www.ciweimao.com">hbooker</a> in Go.</h4>

## Features

- Login your own account
- Search books by book name
- Download books as txt and epub files ([epub3.0](http://idpf.org/epub/30/))
- Download vip chapters
- Multi-threads support

## Usage

- `cirno login` to Log in to your account.
- `cirno search xxxx` to search books.
- `cirno download bid` to download books.
- You can add `--type` flag to specify the books type, support `txt` and `epub`, default value is `txt`.
  ```shell
  cirno -t epub download 100003327
  cirno -t epub search 幻想乡的琐碎日常
  ```

## Download

No pre-build binaries to download, please build the source by yourself.

## Notice

- This tool is for learning only. Please delete it from your computer within 24 hours after downloading.

- Please respect the copyright and do not spread the crawled books by yourself.

 