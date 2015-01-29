# harborly-watch
A quick and simple Go app that sends email notifications (SMS via Twilio maybe soon) on X dollar difference in Harbor.ly's bid/ask price for any Harborly supported fiat or crypto.

### Installing

Full instructions coming soon. Essentially, do the following: `git clone` this repo, `go build` it, make the config.yaml file, and daemonize the binary.

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
