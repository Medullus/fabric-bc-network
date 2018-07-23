package main

import (
	"log"
	"net/http"
	"bytes"
	"os"
	"bufio"
	"fmt"
	"time"
	"encoding/json"
	"strconv"
)


const(
	ENTITY = "EntityMaster:"
	USER = "User:"
	PO = "PurchaseOrder:"
	INVOICE = "Invoice:"

	//URL = "http://leanblocks.eastus.cloudapp.azure.com"
	URL = "http://localhost:3000"

	ADDUSER = URL+"/users"
	ADDINVOICE = URL+"/invoices"
	ADDENTITY = URL+"/entitymasters"
	ADDPO = URL+"/purchaseorders"
	DOC = URL+"/documents/123"
)

var badList = make([]string, 1)

func main(){
	//files :=  []string{"Users.json", "PurchaseOrder.json", "Invoices.json"}
	//files :=  []string{"Invoices.json"}
	files := []string{"document.txt"}

	start := time.Now()

	for _,v := range files{
		runFile(v)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Completion time : "+elapsed.String())
	fmt.Println(badList)
}

func runFile(fileStr string){
	file, err := os.Open(fileStr)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//line := scanner.Text()
		//fmt.Println(line)
		switch line:= scanner.Text(); line{
		case USER:
			scanner.Scan()
			nxtLine := scanner.Text()
			//fmt.Println(os.Args[0]+ ":"+os.Args[1])
			if len(os.Args) > 1 {
				var userOb UserRegister
				err := json.Unmarshal([]byte(nxtLine), &userOb)
				if err != nil{
					fmt.Errorf("Error unmarshaling "+nxtLine)
				}
				index := strconv.Itoa(i)

				userOb.UserRegister.UserName = "randomUser"+index+os.Args[1]
				i++
				fmt.Println("generating random user "+index)
				byte, err := json.Marshal(userOb)
				nxtLine = string(byte)
			}

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
		default:
			var doc Doc
			err := json.Unmarshal([]byte(line), &doc)
			if err !=nil{
				fmt.Errorf(err.Error())
			}

			for i := 0; i < len(doc.Documents); i++{
				doc.Documents[i].DocumentPK = "docPK"+strconv.Itoa(i)
			}
			b, err := json.Marshal(doc)
			if err!=nil{
				fmt.Println(err.Error())
			}
			line = string(b)

			callApi(DOC, line)
			fmt.Println("Doc number :"+ strconv.Itoa(len(doc.Documents)))
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


type UserRegister struct {
	RequestHeader struct {
		Caller string `json:"caller"`
		Org    string `json:"org"`
	} `json:"requestHeader"`
	UserRegister struct {
		Secret   string `json:"secret"`
		UserName string `json:"userName"`
	} `json:"userRegister"`
}

type Doc struct {
	Documents []struct {
		DocumentPK string `json:"documentPK"`
		AnyKey1 string `json:"anyKey1"`
		AnyKey2 string `json:"anyKey2"`
	} `json:"documents"`
	RequestHeader struct {
		Caller string `json:"caller"`
		Org    string `json:"org"`
	} `json:"requestHeader"`
}