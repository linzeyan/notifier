# Telegram bot notifier

# Quick Start 

1. Edit `docker-compose.yml`

```
version: "3"

services:
    notifer:
        image: containerize/notifier
        environment:
            # debug mode
            - NOTIFIER_DEBUG=false
            # server port
            - NOTIFIER_PORT=8080
            # telegram bot token
            - NOTIFIER_BOTTOKEN=<telegram bot token>
            # enable send message to admin
            - NOTIFIER_ENABLEADMIN=true
            # admin ids(split with ,)
            - NOTIFIER_ADMINS=<telegram user 1 id,telegram user 2 id>
            # group ids(split with ,)
            - NOTIFIER_GROUPS=<telegram group 1 id,telegram group 2 id>
            # timezone
            - TZ=Asia/Macau
        ports:
            - 20001:8080

```
2. Run `docker-compose up -d`
3. Test 
```
curl -X "POST" "http://localhost:20001/send" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
  "message": "test123"
}'
```