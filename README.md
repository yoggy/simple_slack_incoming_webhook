# simple_slack_incoming_webhook

```
$ go build 
$ ./simple_slack_incoming_webhook
usage: ./simple_slack_incoming_webhook slack_incoming_webhook_url

example:
    $ somecommand | ./simple_slack_incoming_webhook https://hooks.slack.com/services/xxxxxxxxxxxx


$ date | ./simple_slack_incoming_webhook https://hooks.slack.com/services/xxxxxxxxxxxx
2022/01/25 18:43:21 [+] url=https://hooks.slack.com/services/xxxxxxxxxxxx
2022/01/25 18:43:21 [+] {"text":"2022年  1月 25日 火曜日 18:43:21 JST\n"}
2022/01/25 18:43:22 [*] 200 OK


$ crontab -e
0 10 * * *  date | ./simple_slack_incoming_webhook https://hooks.slack.com/services/xxxxxxxxxxxx 2>/dev/null

```

## Copyright and license
Copyright (c) 2022 yoggy

Released under the [MIT license](LICENSE.txt)
