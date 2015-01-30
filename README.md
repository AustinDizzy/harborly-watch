# harborly-watch
A quick and simple Go app that sends email notifications (SMS via Twilio maybe soon) on X dollar difference in Harbor.ly's bid/ask price for any Harborly supported fiat or crypto.

### Installing

**With Go**
````bash
$ go get github.com/austindizzy/harborly-watch
$ cd $GOPATH/src/github.com/austindizzy/githook
$ vim config.yaml # see "config.yaml" example below
$ make install
````

NOTE: `make install` will require sudo credentials to move files to the /opt/ and /etc/init.d directories.

**Without Go**

*Coming soon.*

### Example config.yaml
A `config.yaml` file needs to be in the same working directory as the binary. Example contents include:
````yaml
coin: btc # any crypto supported by Harborly (currently btc or ltc)
fiat: usd # any fiat type supported by Harborly
interval: 5h # strings from time.ParseDuration: http://golang.org/pkg/time/#ParseDuration
# ex: 1h15m for checking every 1 hour 15 min, 10s for every 10 seconds, 1d for every day, etc.
difference: 3.50 # dollar difference, cents supported, to trigger notification
email:
  username: test@example.com
  password: testing1234
  server: smtp.example.com
  port: 587
  recipient: user@domain.com
````

### Starting harborly-watch service
After installing:
````bash
$ sudo service harborly-watch start
--
$ sudo update-rc.d githook defaults # sets harborly-watch as a default service, runs on system boot
````
