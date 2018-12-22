# ntp
Small NTP client written in GO. Used to schedule reboots on server's whos time is running in a loop, and do not contain their own NTP service.

## Usage
Place somewhere in your bath and then invoke without any arguments:
```
gntp

----
15 23 55
```

Parse with awk to get hours:
```
gntp | awk '{print $1}'
```
