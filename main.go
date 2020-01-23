package main

import (
	"fmt"
	"github.com/zeebo/bencode"
	"io/ioutil"
	"os"
)

type Node string
type UrlList []string

type FileInfo struct {
	Length   int64    `bencode:"length"`
	Path     []string `bencode:"path"`
	PathUTF8 []string `bencode:"path.utf-8,omitempty"`
}

type Info struct {
	PieceLength int64      `bencode:"piece length"`
	Pieces      []byte     `bencode:"pieces"`
	Name        string     `bencode:"name"`
	Length      int64      `bencode:"length,omitempty"`
	Source      string     `bencode:"source,omitempty"`
	Files       []FileInfo `bencode:"files,omitempty"`
}

type MetaInfo struct {
	Info         Info       `bencode:"info,omitempty"`
	Announce     string     `bencode:"announce,omitempty"`
	AnnounceList [][]string `bencode:"announce-list,omitempty"`
	Nodes        []Node     `bencode:"nodes,omitempty"`
	CreationDate int64      `bencode:"creation date,omitempty,ignore_unmarshal_type_error"`
	Comment      string     `bencode:"comment,omitempty"`
	CreatedBy    string     `bencode:"created by,omitempty"`
	Encoding     string     `bencode:"encoding,omitempty"`
	UrlList      UrlList    `bencode:"url-list,omitempty"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]

	dat, err := ioutil.ReadFile(args[0])
	check(err)

	var torrent MetaInfo

	err = bencode.DecodeBytes(dat, &torrent)

	check(err)
	fmt.Printf("%+v\n", torrent)
}
