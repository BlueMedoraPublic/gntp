# ntp
Small NTP client written in GO. Used to schedule reboots on servers who's time is running in a fixed loop, and do not contain their own NTP service.

## Usage
Place somewhere in your path and invoke without any arguments:
```
gntp

----
15 23 55
```

Parse with awk to get hours:
```
gntp | awk '{print $1}'
```

run from cron every hour
``` E=`/usr/local/bin/gntp | awk '{print $1}'` ```
```
if [ "${E}" = "${reboot_hour}" ]; then
  # sleep for an hour and then reboot
  sleep 3600 && shutdown -r now
fi
```
