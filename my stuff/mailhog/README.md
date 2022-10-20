# Testing email with mailhog

## tutorial [link](https://www.loginradius.com/blog/engineering/sending-emails-with-golang/)

> spin-up a mailhog container

```bash
docker run -p 8025:8025 -p 1025:1025 -d mailhog/mailhog
```

> or use docker-compose

```bash
# bring up
docker-compose up -d
# bring down
docker-compose down -v
```
