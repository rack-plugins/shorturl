package shorturl

import (
	"path/filepath"

	"github.com/fimreal/goutils/ezap"
	"github.com/spf13/viper"
	"github.com/syndtr/goleveldb/leveldb"
)

var DB *leveldb.DB

func initDB() (err error) {
	dbPath := filepath.Join(viper.GetString("workdir"), viper.GetString("shorturl.dbpath"))
	ezap.Debug("[module] shorturl db init at ", dbPath)
	DB, err = leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return
	}
	// keep db connect
	// defer DB.Close()
	return
}
