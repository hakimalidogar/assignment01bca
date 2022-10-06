package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Transactions     [3]string
	Nounce           int
	Prevblockhash    string
	Currentblockhash    string
	Timestamp   	 string
}

type Blockchain struct{
	List     []*Block
}

func NewBlock(transaction [3]string, nonce int, previousHash string) *Block {
	b := new(Block)
	b.Nounce = nonce
	b.Transactions = transaction
	b.Prevblockhash = previousHash
	comtransactions := transaction[0]+transaction[1]+transaction[2]
	b.Currentblockhash = CalculateHash(comtransactions)
	dt := time. Now()
	b.Timestamp = dt.String()
	return b
}

func (obj *Blockchain) Addblocktolist(transaction [3]string, nonce int) *Block {
	lenthoflist := len(obj.List)
	prevblockhash := ""

	if lenthoflist > 0{
		prevblockhash = obj.List[lenthoflist-1].Transactions[0]+obj.List[lenthoflist-1].Transactions[1]+
						obj.List[lenthoflist-1].Transactions[2]+strconv.Itoa(obj.List[lenthoflist-1].Nounce)+
						obj.List[lenthoflist-1].Prevblockhash+obj.List[lenthoflist-1].Currentblockhash+
						obj.List[lenthoflist-1].Timestamp
	}
	previousHash := CalculateHash(prevblockhash)
	as1 := NewBlock(transaction, nonce, previousHash)
	obj.List = append(obj.List, as1)
	return as1
}

func ListBlocks(obj *Blockchain) { 
	for i := 0; i < len(obj.List); i++ {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))

		for j := 0; j < len(obj.List[i].Transactions); j++ {
			fmt.Printf("Transactions %d: %s", j+1,obj.List[i].Transactions[j])
			fmt.Printf("\n")
		}
		fmt.Println("Nounce: ", obj.List[i].Nounce)
		fmt.Println("Previous Block Hash: ", obj.List[i].Prevblockhash)
		fmt.Println("Current Block Transactions Hash: ", obj.List[i].Currentblockhash)
		fmt.Println("Timestamp of Block: ", obj.List[i].Timestamp)
		fmt.Printf("\n\n")
	}
}

func ChangeBlock(obj *Blockchain, nounce int, newtransaction [3]string) {
	for i := 0; i < len(obj.List); i++ {
		if obj.List[i].Nounce == nounce {
		
			obj.List[i].Transactions[0] = newtransaction[0]
			obj.List[i].Transactions[1] = newtransaction[1]
			obj.List[i].Transactions[2] = newtransaction[2]
			fmt.Printf("Changes completed\n")
			return
		}
	}
		
	fmt.Printf("block not found!\n\n")
}

func VerifyChain(obj *Blockchain) bool {
	var checking = ""
	for i := 0; i < len(obj.List); i++ {
		checking = obj.List[i].Transactions[0]+obj.List[i].Transactions[1]+
				   obj.List[i].Transactions[2]
		sum := CalculateHash(checking)

		if sum != obj.List[i].Currentblockhash {
			fmt.Printf("Block is tempered, Block No. : %d\n", i+1)
			return false
		}
	}
	fmt.Printf("Blockchain is not tempered\n\n")
	return true
}

func CalculateHash(stringToHash string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func main() {
	b1 := new(Blockchain)

	var option int

	for i := 0; true; i++ {
		fmt.Printf("Press '1' to Add Block\n")
		fmt.Printf("Press '2' to change Block\n")
		fmt.Printf("Press '3' to VerifyBlockchain Block\n")
		fmt.Printf("Press '4' to Display Blockchain Blocks\n")
		fmt.Printf("Press '0' to Exit\n")
		fmt.Printf("Select your option:\n")
		fmt.Scan(&option)

		if option == 1 {
			var transaction [3]string
			var nonce = 0

			fmt.Printf("Enter transaction in fortmat: 'sender_to_reciever'\n\n");
			fmt.Printf("Enter the 1 transaction: ");
			fmt.Scan(&transaction[0])
			fmt.Printf("\nEnter the 2 transaction: ");
			fmt.Scan(&transaction[1])
			fmt.Printf("\nEnter the 3 transaction: ");
			fmt.Scan(&transaction[2])

			fmt.Printf("\nEnter the nonce : ");
			fmt.Scan(&nonce)

			b1.Addblocktolist(transaction, nonce)
			fmt.Printf("Block Added\n\n");
		}else if option == 2 {
			var transaction  = [3]string{}
			var nonce = 0

			fmt.Printf("Enter transaction in fortmat: 'sender_to_reciever'\n\n");
			fmt.Printf("Enter the 1 new transaction: ");
			fmt.Scan(&transaction[0])
			fmt.Printf("\nEnter the 2 new transaction: ");
			fmt.Scan(&transaction[1])
			fmt.Printf("\nEnter the 3 new transaction: ");
			fmt.Scan(&transaction[2])

			fmt.Printf("\nEnter the nonce of desired block: ");
			fmt.Scan(&nonce)

			ChangeBlock(b1,nonce,transaction)
			fmt.Printf("Block changed\n\n");
		}else if option == 3 {
			fmt.Printf("\n\n")
			VerifyChain(b1)
		}else if option == 4 {
			fmt.Printf("\n\n")
			ListBlocks(b1)
		}else if option == 0 {
			break
		}
	}

	fmt.Printf("Exited\n\n")
	
}