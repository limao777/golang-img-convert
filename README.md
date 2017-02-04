# golang-img-convert
golang写的图片转换程序，可以将图片转换成任一尺寸，基于beego的web应用

## nginx需要做优先的配置
```text
    location / {
        root   /xxx;
        index  index.html index.htm;
        try_files $uri @imghandle;
     }

 location @imghandle {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:9080;
    }
```

## 然后部署本程序，端口随配置文件改变而改变