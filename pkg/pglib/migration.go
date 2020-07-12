package pglib

import (
	"context"
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
)

var dir = "./migration/postgres/"

type table struct {
	TableName *string `json:"table_name" db:"table_name"`
}

type metadata struct {
	Key   string `json:"key" db:"key"`
	Value string `json:"value" db:"value"`
}

//Migrate func to create database DDL and seeder
func Migrate(db *sqlx.DB) error {
	ctx := context.Background()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	checktable := table{}
	emptyParams := map[string]interface{}{}
	err = Query(ctx, db, "SELECT to_regclass('metadata') as table_name", emptyParams, &checktable)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	migrationVersion := 0
	if checktable.TableName == nil {
		tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS "metadata" (
			key VARCHAR (50) PRIMARY KEY,
			value VARCHAR (50) NOT NULL
		);
		INSERT INTO public.metadata ("key",value) VALUES ('MIRAGRATION_VERSION','`+strconv.Itoa(len(files))+`');`)
	} else {
		var meta metadata
		err = Query(ctx, db, "SELECT key,value from metadata", emptyParams, &meta)
		if err != nil {
			return err
		}
		migrationVersion, _ = strconv.Atoi(meta.Value)
	}

	err = execMigration(ctx, files, tx, migrationVersion)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func execMigration(ctx context.Context, files []os.FileInfo, tx *sql.Tx, migrationVersion int) error {
	log.Printf("Migration Version %d", migrationVersion)
	log.Printf("Current Migration Version %d", len(files))
	i := 1
	for _, f := range files {
		if i > migrationVersion {
			log.Printf("Executing Migration %d: %s", i, f.Name())
			byteSQL, err := ioutil.ReadFile(dir + f.Name())
			if err != nil {
				log.Println(err)
				return err
			}
			_, err = tx.ExecContext(ctx, string(byteSQL))
			if err != nil {
				log.Println(err)
				return err
			}
		}
		i++
	}

	return nil
}
