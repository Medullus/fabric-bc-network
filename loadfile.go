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

	URL = "http://localhost:3000"

	ADDUSER = URL+"/users"
	ADDINVOICE = URL+"/invoices"
	ADDENTITY = URL+"/entitymasters"
	ADDPO = URL+"/purchaseorders"
)

var badList = make([]string, 1)

func main(){
	files :=  []string{"Users.json", "PurchaseOrder.json", "Invoices.json"}
	//files :=  []string{"Invoices.json"}


	for _,v := range files{
		runFile(v)
	}

	fmt.Println(badList)

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

	if resp.StatusCode != 200{
		badList = append(badList, endpoint+"----"+str)
	}

	fmt.Printf("Code : %s for %s\n %s \n\n", resp.Status, endpoint, str)
}

