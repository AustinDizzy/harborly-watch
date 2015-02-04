package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"path"
	"strconv"
)

func main() {

	LoadConfig()

	db, err := bolt.Open(path.Join(GetDir(), "harborly.db"), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	c := cron.New()
	c.AddFunc("@every "+Config.Interval, func() {
		resp, _ := http.Get("https://harbor.ly/ticker/" + Config.Coin + "/" + Config.Fiat)
		body, _ := ioutil.ReadAll(resp.Body)
		var r map[string]interface{}
		json.Unmarshal(body, &r)

		db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("BTCTicker"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}

			for _, v := range []string{"bid", "ask"} {
				val := b.Get([]byte(v))
				bytesVal := []byte(r[v].(string))

				if val != nil {
					checkPrice(b, v, bytesVal)
				}

				updateField(b, v, bytesVal)
			}

			return nil
		})
	})

	c.Start()
	select {}
}

func checkPrice(b *bolt.Bucket, key string, val []byte) {
	v, err := strconv.ParseFloat(string(val), 64)
	bytesOrigVal := b.Get([]byte(key))
	LogErr(err)

	if bytesOrigVal == nil {
		return
	}

	origVal, err := strconv.ParseFloat(string(bytesOrigVal), 64)
	LogErr(err)

	diff := origVal - v
	if math.Abs(diff) >= Config.Difference {
		SendEmail(diff)
	}
}

func updateField(b *bolt.Bucket, key string, newVal []byte) {
	b.Put([]byte(key), newVal)
}

//LogErr only logs an error to the console
//if and only if the error is not nil.
func LogErr(err error) {
	if err != nil {
		log.Printf("got error: %v", err.Error())
	}
}
