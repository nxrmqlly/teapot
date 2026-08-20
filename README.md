# teapot

HTTP/1.1 418 

```sh
docker build -t teapot .
docker run -d -p 4180:4180 --restart=unless-stopped --name running-teapot teapot
```
