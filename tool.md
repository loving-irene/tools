## curl时间探测命令
```curl
curl -L -w "time_namelookup: %{time_namelookup}
time_connect: %{time_connect}
time_appconnect: %{time_appconnect}
time_pretransfer: %{time_pretransfer}
time_redirect: %{time_redirect}
time_starttransfer: %{time_starttransfer}
time_total: %{time_total}
" https://example.com/
```
[参考链接](https://susam.net/blog/timing-with-curl.html)
