package main

import (
    "log"
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "github.com/robfig/cron"
    "github.com/boltdb/bolt"
    "strconv"
    "math"
)

func main() {

    LoadConfig()

    db, err := bolt.Open("harborly.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    c := cron.New()
    c.AddFunc("@every " + Config.Interval, func(){
        resp, _ := http.Get("https://harbor.ly/ticker/" + Config.Coin + "/" + Config.Fiat)
        body, _ := ioutil.ReadAll(resp.Body)
        var r map[string]interface{}
        json.Unmarshal(body, &r)

        db.Update(func(tx *bolt.Tx) error {
            b, err := tx.CreateBucketIfNotExists([]byte("BTCTicker"))
            if err != nil {
                return fmt.Errorf("create bucket: %s", err)
            }
            askVal := b.Get([]byte("ask"))
            bidVal := b.Get([]byte("bid"))

            if askVal == nil {
                bytesAsk := []byte(r["ask"].(string))
                b.Put([]byte("ask"), bytesAsk)
            } else {
                checkPrice(b, "ask", askVal)
            }

            if bidVal == nil {
                bytesBid := []byte(r["bid"].(string))
                b.Put([]byte("bid"), bytesBid)
            } else {
                checkPrice(b, "bid", bidVal)
            }

            return nil
        })
    })

    c.Start()
    select{}
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

    diff := math.Abs(origVal - v)
    if diff >= 5 {
        //todo: send email
    }
}

func LogErr(err error) {
    if err != nil {
        log.Printf("got error: %v", err.Error())
    }
}