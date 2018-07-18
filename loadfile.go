package main

import (
	"log"
	"net/http"
	"bytes"
	"os"
	"bufio"
	"fmt"
)


const(
	ENTITY = "EntityMaster:"
	USER = "User:"
	PO = "PurchaseOrder:"
	INVOICE = "Invoice:"

	ADDUSER = "http://localhost:3000/users"
	ADDINVOICE = "http://localhost:3000/invoices"
	ADDENTITY = "http://localhost:3000/entitymasters"
	ADDPO = "http://localhost:3000/purchaseorders"
)

func main(){
	files :=  []string{"Users.json", "PurchaseOrder.json", "Invoices.json"}

	for _,v := range files{
		runFile(v)
	}

}

func runFile(fileStr string){
	file, err := os.Open(fileStr)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//line := scanner.Text()
		//fmt.Println(line)
		switch line:= scanner.Text(); line{
		case USER:
			scanner.Scan()
			nxtLine := scanner.Text()
			//fmt.Println("USER JSON :"+nxtLine)

			callApi(ADDUSER, nxtLine)
		case ENTITY:
			scanner.Scan()
			nxtLine:= scanner.Text()
			//fmt.Println("ENTITY JSON :"+nxtLine)
			callApi(ADDENTITY, nxtLine)
		case PO:
			scanner.Scan()
			nxtLine:= scanner.Text()
			callApi(ADDPO, nxtLine)
		case INVOICE:
			scanner.Scan()
			nxtLine:= scanner.Text()
			callApi(ADDINVOICE, nxtLine)

		}
	}
}



func callApi(endpoint string, str string ){
	var jsonStr = []byte(str)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		panic(err)

	}

	defer resp.Body.Close()

	fmt.Printf("Code : %s for %s\n %s \n\n", resp.Status, endpoint, str)
}



type EntityMaster struct {
	SubName           string `json:"subName"`
	GlEntityCode      string `json:"glEntityCode"`
	Group             string `json:"group"`
	FnlCurr           string `json:"fnlCurr"`
	Account           string    `json:"account"`
	AdditionalReview  string `json:"additionalReview"`
	Bank              string `json:"bank"`
	Country           string `json:"country"`
	NettingSettRules  string `json:"nettingSettRules"`
	Paymaster         string `json:"paymaster"`
	PaymasterEligible string `json:"paymasterEligible"`
	Wht               string `json:"wht"`
}