package main

import (
	"encoding/json"
	"io/ioutil"
)

type nft struct {
	Id          string
	Description string
	OwnedBy     string
}

func getNft() (nfts []nft) {

	fileBytes, err := ioutil.ReadFile("./nft.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &nfts)

	if err != nil {
		panic(err)
	}
	return nfts
}

func saveNft(nfts []nft) {

	nftBytes, err := json.Marshal(nfts)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./nft.json", nftBytes, 0644)
	if err != nil {
		panic(err)
	}

}
