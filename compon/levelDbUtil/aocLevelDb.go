package levelDbUtil

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"sync"
)

type AocLevelDb struct {
	dbc      *leveldb.DB
	filePath string
	//wo *opt.WriteOptions
	locker sync.RWMutex //leveldb一般并发读写操作并不需要锁，但是在某些特殊的操作的时候加锁，比如并发[读取内容，对读取的内容修改然后写入]，并发这个事务需要锁
}

func NewAocLevelDb(url string) (*AocLevelDb, error) {
	resModel := &AocLevelDb{}
	resModel.filePath = url
	dbc, err := leveldb.OpenFile(resModel.filePath, nil)
	if err != nil {
		return nil, err
	}

	resModel.dbc = dbc
	return resModel, nil
}

func (this *AocLevelDb) GetRwlocker() *sync.RWMutex {
	return &this.locker
}

func (this *AocLevelDb) Close() {
	if this.dbc == nil {
		return
	}

	this.dbc.Close()
}

type KeyValuePair struct {
	Key   []byte
	Value []byte
}

/*sync bool表示同步操作，false表示异步*/
func (this *AocLevelDb) BatchPut(sync bool, array []KeyValuePair) error {
	if array == nil || len(array) == 0 {
		return nil
	}

	wo := new(opt.WriteOptions)
	wo.Sync = sync

	batch := new(leveldb.Batch)
	for idx, _ := range array {
		batch.Put(array[idx].Key, array[idx].Value)
	}
	return this.dbc.Write(batch, wo)
}
func (this *AocLevelDb) Put(key, value []byte) error {
	return this.dbc.Put(key, value, nil)
}

func (this *AocLevelDb) Get(key []byte) ([]byte, error) {
	return this.dbc.Get(key, nil)
}

func (this *AocLevelDb) Delete(key []byte) error {
	return this.dbc.Delete(key, nil)
}

/*sync bool表示同步操作，false表示异步*/
func (this *AocLevelDb) BatchDel(sync bool, array [][]byte) error {
	if array == nil || len(array) == 0 {
		return nil
	}

	wo := new(opt.WriteOptions)
	wo.Sync = sync

	batch := new(leveldb.Batch)
	for idx, _ := range array {
		batch.Delete(array[idx])
	}
	return this.dbc.Write(batch, wo)
}

func (this *AocLevelDb) CompactRange(startKey, limit []byte) error {
	return this.dbc.CompactRange(util.Range{Start: startKey, Limit: limit})
}

func (this *AocLevelDb) GetDbc() *leveldb.DB {
	return this.dbc
}
func (this *AocLevelDb) GetDbUrl() string {
	return this.filePath
}
