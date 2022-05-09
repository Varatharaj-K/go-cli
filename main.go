package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	getAll := getCmd.Bool("all", false, "Get all NFTs")
	getID := getCmd.String("id", "", "NFT ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	addID := addCmd.String("id", "", "nft ID")
	addTitle := addCmd.String("description", "", "Description of the NFT")
	addOwner := addCmd.String("ownedBy", "", "Owner of the NFT")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getID)
	case "add":
		HandleAdd(addCmd, addID, addTitle, addOwner)
	default:
	}

}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all for all NFTs")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		nfts := getNft()

		fmt.Printf("ID \t Description \t Owner \n")
		for _, nft := range nfts {
			fmt.Print(nft.Id)
			fmt.Printf("%v \t %v \t %v \n", nft.Id, nft.Description, nft.OwnedBy)
		}

		return
	}

	if *id != "" {
		nfts := getNft()
		id := *id
		for _, nft := range nfts {
			if id == nft.Id {
				fmt.Printf("ID \t Description \t Owner \n")
				fmt.Printf("%v \t %v \t %v \n", nft.Id, nft.Description, nft.OwnedBy)
			}
		}
	}

}

func ValidateNFT(addCmd *flag.FlagSet, id *string, description *string, ownedBy *string) {

	addCmd.Parse(os.Args[2:])

	if *id == "" || *description == "" || *ownedBy == "" {
		fmt.Print("all fields are required for adding an NFT")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet, id *string, description *string, ownedBy *string) {

	ValidateNFT(addCmd, id, description, ownedBy)

	nft := nft{
		Id:          *id,
		Description: *description,
		OwnedBy:     *ownedBy,
	}

	nfts := getNft()
	nfts = append(nfts, nft)

	saveNft(nfts)

}
