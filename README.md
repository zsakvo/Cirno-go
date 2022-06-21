
<h1 align="center">
  <img src="./assets/cirno.png" alt="Cirno" width="150">
  <br>Cirno-go<br>
</h1>

<h4 align="center">A tool for downloading books from <a href="https://www.ciweimao.com">hbooker</a> in Go.</h4>

## Features

- ~~Login your own account~~
- Search books by book name
- Download books as txt and epub files ([epub3.0](http://idpf.org/epub/30/))
- Download vip chapters
- Multi-threads support
- Cache images in chapters

## ⚠️ Notice

- Now, a third-party captcha([geetest](https://www.geetest.com/)) has been add to the official server. If you login failed with the right account and password, please capture packets of the app to get the token and account, then, locate to `$HOME/cirno` dir, create a file named `config.yaml`:
```yaml
app:
  account: 
  login_token: 
```

## Usage

- ~~`cirno login` to Log in to your account.~~
- `cirno search xxxx` to search books.
- `cirno download bid` to download books.
- You can add `--type` flag to specify the books type, support `txt` and `epub`, default value is `txt`.
  
  ```shell
  cirno -t epub download 100003327
  cirno -t epub search happyend
  ```

## Config

- All files about `Cirno-go` are located in `$HOME/Cirno/`
- Do not delete `config.yaml`, otherwise you need to do `cirno login` again.
- Cache chapter images is an extra feature, you need open it manually, Such as:

  ```yaml
  app:
    account: 
    login_token: 
    user_name: 
    host_url:  # you can use another official api https://app.happybooker.cn if you can't visit the default one.

  extra:
    cpic: true  #set true to cache images automatically.
  ```
- There is a little extra properties waiting for you to discover.



## Download

No pre-build binaries to download, please build the source by yourself.

- if you want to run it on Android devices, please use `termux`, and run 

  `pkg install proot resolv-conf && proot -b $PREFIX/etc/resolv.conf:/etc/resolv.conf` 
  

## Notice

- This tool is for learning only. Please delete it from your computer within 24 hours after downloading.

- Please respect the copyright and do not spread the crawled books by yourself.

 