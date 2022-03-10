package keyHandler

import (
	"database/sql"
	"fmt"
	"time"

	"src/mcl_server/pkg/mcl"

	"github.com/go-gorp/gorp"

	_ "github.com/lib/pq"
)

func getMasterKey(time string) *mcl.Fr {
	key := &mcl.Fr{}

	db, err := sql.Open("postgres", "user=pguser dbname=db password=password sslmode=disable?parseTime=true")
	if err != nil {
		return nil
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	defer dbmap.Db.Close()

	var keys []string
	_, err = dbmap.Select(&keys, fmt.Sprintf("select key from MasterKey where end > %d and start <= %d", time, time))

	if len(keys) != 1 || err != nil {
		return nil
	}
	key.SetString(keys[0], 10)
	return key
}

func getSecretKey(time string, email string) string {
	var g, out mcl.G1
	master := getMasterKey(time)
	g.HashAndMapTo([]byte(email))
	mcl.G1Mul(&out, &g, master)
	return out.GetString(16)
}

func getSecretKey2(time string, email string) string {
	var g, out mcl.G2
	master := getMasterKey(time)
	g.HashAndMapTo([]byte(email))
	mcl.G2Mul(&out, &g, master)
	return out.GetString(16)

}

func getPublicKey(param string) string {
	var p1, out mcl.G1
	p1.SetString(param, 16)
	master := getMasterKey(time.Now().Format("2006-01-02 15:04:05"))
	mcl.G1Mul(&out, &p1, master)
	return out.GetString(16)
}

func getPublicKey2(time string, param string) string {
	var p2, out mcl.G2
	p2.SetString(param, 16)
	master := getMasterKey(time)
	mcl.G2Mul(&out, &p2, master)
	return out.GetString(16)
}

func getKtTimeKey(time string, param string) string {
	var g, out mcl.G1
	master := getMasterKey(time)
	g.HashAndMapTo([]byte(param))
	mcl.G1Mul(&out, &g, master)
	return out.GetString(16)
}

func getTtTimeKey(time string, dectime string) string {
	var p1, out mcl.G1
	p1.SetString(dectime, 16)
	master := getMasterKey(time)
	mcl.G1Mul(&out, &p1, master)
	return out.GetString(16)
}
