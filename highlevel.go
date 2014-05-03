package goleg

/*
#cgo CFLAGS: -I../oleg/include/
#cgo LDFLAGS: -L../oleg/build/lib/ -loleg
#include <oleg.h>
*/
import "C"
import "time"

type Database struct {
	db          *C.ol_database
	RecordCount *C.int
}

func Open(path, name string, features int) Database {
	var database Database
	database.db = COpen(path, name, features)
	database.RecordCount = &database.db.rcrd_cnt
	return database
}

func (d Database) Close() int {
	return CClose(d.db)
}

func (d Database) CloseSave() int {
	return CCloseSave(d.db)
}

func (d Database) Unjar(key string) []byte {
	var dsize uintptr
	return CUnjarDs(d.db, key, uintptr(len(key)), &dsize)
}

func (d Database) Jar(key string, value []byte) int {
	return CJar(d.db, key, uintptr(len(key)), value, uintptr(len(value)))
}

func (d Database) JarCt(key string, value []byte, content_type string) int {
	return CJarCt(d.db, key, uintptr(len(key)), value, uintptr(len(value)), content_type, uintptr(len(content_type)))
}

func (d Database) ContentType(key string) string {
	return CContentType(d.db, key, uintptr(len(key)))
}

func (d Database) Scoop(key string) int {
	return CScoop(d.db, key, uintptr(len(key)))
}

func (d Database) Uptime() int {
	return CUptime(d.db)
}

func (d Database) Expiration(key string) (time.Time, bool) {
	return CExpirationTime(d.db, key, uintptr(len(key)))
}

func (d Database) Spoil(key string, expiration time.Time) int {
	return CSpoil(d.db, key, uintptr(len(key)), expiration)
}
