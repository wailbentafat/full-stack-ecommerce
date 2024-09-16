package cach

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

)

type Cach struct {
	Data []byte
	expiration time.Time
}

type Filecach struct {
	cachdir string
}

func NewFilecach(cachdir string) (*Filecach, error) {
    err:=os.MkdirAll(cachdir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &Filecach{ cachdir: cachdir}, nil
	
	
	}
func (f *Filecach)getfilepath(key string) string{
	return filepath.Join(f.cachdir, key+".cacked")

}
func (f *Filecach)GEtkey(key string) ([]byte,bool) {

	filepath:=f.getfilepath(key)
	_,err:= os.Stat(filepath)
	if err != nil {
		return nil,false}
		
	file, err := os.Open(filepath)
	if err != nil {
		return nil,false}
	defer file.Close()
	decoder:=json.NewDecoder(file)
	var cach Cach
	err=decoder.Decode(&cach)
	if err != nil {
		return nil,false
	}
	if cach.expiration.Before(time.Now()) {
		return nil,false
	}
	return cach.Data, true
}
func ( f *Filecach)Setkey(key string, data []byte,duration time.Duration) error {
	filepath:=f.getfilepath(key)
	item:=Cach{
		Data: data,
		expiration: time.Now().Add(duration),
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder:=json.NewEncoder(file)
	return encoder.Encode(&item)
	
}