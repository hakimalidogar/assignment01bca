package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	transactions     [3]string
	nounce           int
	prevblockhash    string
	currentblockhash    string
	timestamp   	 string
}

type Blockchain struct{
	list     []*Block
}

func NewBlock(transaction [3]string, nonce int, previousHash string) *Block {
	b := new(Block)
	b.nounce = nonce
	b.transactions = transaction
	b.prevblockhash = previousHash
	comtransactions := transaction[0]+transaction[1]+transaction[2]
	b.currentblockhash = CalculateHash(comtransactions)
	dt := time. Now()
	b.timestamp = dt.String()
	return b
}

func (obj *Blockchain) Addblocktolist(transaction [3]string, nonce int) *Block {
	lenthoflist := len(obj.list)
	prevblockhash := ""

	if lenthoflist > 0{
		prevblockhash = obj.list[lenthoflist-1].transactions[0]+obj.list[lenthoflist-1].transactions[1]+
						obj.list[lenthoflist-1].transactions[2]+strconv.Itoa(obj.list[lenthoflist-1].nounce)+
						obj.list[lenthoflist-1].prevblockhash+obj.list[lenthoflist-1].currentblockhash+
						obj.list[lenthoflist-1].timestamp
	}
	previousHash := CalculateHash(prevblockhash)
	as1 := NewBlock(transaction, nonce, previousHash)
	obj.list = append(obj.list, as1)
	return as1
}

func ListBlocks(obj *Blockchain) { 
	for i := 0; i < len(obj.list); i++ {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))

		for j := 0; j < len(obj.list[i].transactions); j++ {
			fmt.Printf("Transactions %d: %s", j+1,obj.list[i].transactions[j])
			fmt.Printf("\n")
		}
		fmt.Println("Nounce: ", obj.list[i].nounce)
		fmt.Println("Previous Block Hash: ", obj.list[i].prevblockhash)
		fmt.Println("Current Block Transactions Hash: ", obj.list[i].currentblockhash)
		fmt.Println("Timestamp of Block: ", obj.list[i].timestamp)
		fmt.Printf("\n\n")
	}
}

func ChangeBlock(obj *Blockchain, nounce int, newtransaction [3]string) {
	for i := 0; i < len(obj.list); i++ {
		if obj.list[i].nounce == nounce {
		
			obj.list[i].transactions[0] = newtransaction[0]
			obj.list[i].transactions[1] = newtransaction[1]
			obj.list[i].transactions[2] = newtransaction[2]
			fmt.Printf("Changes completed\n")
			return
		}
	}
		
	fmt.Printf("block not found!\n\n")
}

func VerifyChain(obj *Blockchain) bool {
	var checking = ""
	for i := 0; i < len(obj.list); i++ {
		checking = obj.list[i].transactions[0]+obj.list[i].transactions[1]+
				   obj.list[i].transactions[2]
		sum := CalculateHash(checking)

		if sum != obj.list[i].currentblockhash {
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

	// var trans1 = [3]string{"Alice to Bob", "Sohaib to Umar", "Huma to Usama"}
	// var trans2 = [3]string{"Hakim to Hamnah", "Abdullah to Ammar", "Khalil to Gujjar"}
	// var trans3 = [3]string{"Charlie to Merry", "Ahad to Umar", "Khizer to Huzaifa"}
	// b1.Addblocktolist(trans1, 2781)
	// b1.Addblocktolist(trans2, 2032)
	// b1.Addblocktolist(trans3, 48876)

	// ListBlocks(b1)

	// VerifyChain(b1)

	// ChangeBlock(b1,2032,trans3)

	// VerifyChain(b1)

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

	fmt.Printf("Exited")
	
}