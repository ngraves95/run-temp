# run-temp
run-temp adds temperature and humidity information to your Strava activities. After installation, run-temp runs in the background, checking your Strava account periodically. When it finds an activity with no temperature data, it will add the current temperature and humidity to that activity based on the location of the activity.

To install run-temp:

```shell
~ $ go get github.com/ngraves95/run-temp
~ $ cd /your/go/path/run-temp
~/your/go/path/run-temp $ ./install.sh
```

