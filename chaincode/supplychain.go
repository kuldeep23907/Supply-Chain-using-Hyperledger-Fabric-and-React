package main  

import ( 	
	"bytes" 	
	"encoding/json" 	
	"fmt" 	
	"strconv" 	
	"time"  	
	"github.com/hyperledger/fabric/core/chaincode/shim" 	
	pb "github.com/hyperledger/fabric/protos/peer" 
)  

type food_supplychain struct {
}

type CounterNO struct {
	Counter int `json:"counter"`
}

type User struct { 	
	Name string `json:"Name"` 	
	User_ID string `json:"UserID"` 	
	Email string  `json:"Email"` 	
	User_Type string `json:"UserType"` 
	Address string `json:"Address"`
} 


type ProductDates struct {
	ManufactureDate  string `json:"ManufactureDate"` 
	SendToWholesalerDate  string `json:"SendToWholesalerDate"` 
	SendToDistributorDate  string `json:"SendToDistributorDate"` 
	SendToRetailerDate  string `json:"SendToRetailerDate"` 
	SellToConsumerDate  string `json:"SellToConsumerDate"` 
	DeliveredDate  string `json:"DeliveredDate"` 
}

type Product struct { 	
	Product_ID string `json:"ProductID"` 	
	Order_ID string `json:"OrderID"` 	
	Name string `json:"Name"` 	
	Consumer_ID string `json:"ConsumerID"` 	
	Manufacturer_ID string `json:"ManufacturerID"` 	
	Retailer_ID string `json:"RetailerID"` 	
	Distributer_ID string `json:"DistributerID"` 	
	Wholesaler_ID string `json:"WholesalerID"` 	
	Status string `json:"Status"` 	
	Date ProductDates `json:"Date"` 	
	Price float64 `json:"Price"` 
}  

	
// =================================================================================== // Main // =================================================================================== 

func main() { 	
	err := shim.Start(new(food_supplychain)) 	
	if err != nil { 		
		fmt.Printf("Error starting Simple chaincode: %s", err) 	
		} 
} 

// Init initializes chaincode // =========================== 

func (t *food_supplychain) Init(stub shim.ChaincodeStubInterface) pb.Response 
{ 	return shim.Success(nil) }  

// Invoke - Our entry point for Invocations // ======================================== 

func (t *food_supplychain) Invoke(stub shim.ChaincodeStubInterface) pb.Response 
{ 	function, args := stub.GetFunctionAndParameters() 	
	fmt.Println("invoke is running " + function)  	
	
	// Handle different functions 
		if function == "createProduct" 
		{ //create a new product		
			return t.createProduct(stub, args) 	
		} else if function == "updateProduct"
		{ // update a product		
			return t.updateProduct(stub, args) 	
		} else if function == "sendToWholesaler"
		{ // send to wholesaler		
			return t.sendToWholesaler(stub, args) 	
		}	else if function == "sendToDistributer"
		{ // send to Distributer		
			return t.sendToDistributer(stub, args) 	
		}	else if function == "sendToRetailer"
		{ // send to Retailer		
			return t.sendToRetailer(stub, args) 	
		}	else if function == "sellToConsumer"
		{ // send to Consumer		
			return t.sellToConsumer(stub, args) 	
		} else if function == "query" {
			return t.Query(stub, args)
		} else if function == "queryAll" {
			return t.Query(stub, args)
		}
		fmt.Println("invoke did not find func: " + function) 
		//error 	
		return shim.Error("Received unknown function invocation")  
}

// Private function 

//getCounter to the latest value of the counter based on the Asset Type provided as input parameter
func getCounter(APIstub shim.ChaincodeStubInterface, AssetType string) int {
	counterAsBytes, _ := APIstub.GetState(AssetType)
	counterAsset := CounterNO{}

	json.Unmarshal(counterAsBytes, &counterAsset)
	fmt.Sprintf("Counter Current Value %d of Asset Type %s", counterAsset.Counter, AssetType)

	return counterAsset.Counter
}

//incrementCounter to the increase value of the counter based on the Asset Type provided as input parameter by 1
func incrementCounter(APIstub shim.ChaincodeStubInterface, AssetType string) int {
	counterAsBytes, _ := APIstub.GetState(AssetType)
	counterAsset := CounterNO{}

	json.Unmarshal(counterAsBytes, &counterAsset)
	counterAsset.Counter++
	counterAsBytes, _ = json.Marshal(counterAsset)

	err := APIstub.PutState(AssetType, counterAsBytes)
	if err != nil {

		fmt.Sprintf("Failed to Increment Counter")

	}
	return counterAsset.Counter
}

// send to distributor
func (t *food_supplychain) sendToDistributer(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Less no of arguements provided");
	}

	if len(args[1]) == 0 {
		return shim.Error("Product Id must be provided");
	}

	if len(args[2]) == 0 {
		return shim.Error("Distributer Id must be provided");
	}

	userBytes, _ := APIstub.GetState(args[2])

	if userBytes == nil {
		return shim.Error("Cannot Find Distributer user")
	}

	user := User{}

	json.Unmarshal(userBytes, &user)

	if user.User_Type != "distributor" {
		return shim.Error("User type must be distributor")
	}


	productBytes, _ := APIstub.GetState(args[1])

	if productBytes == nil {
		return shim.Error("Cannot Find Product")
	}

	product := Product{}

	json.Unmarshal(productBytes, &product)

	if product.Distributer_ID != nil {
		return shim.Error("Product is send to distributer already")
	}

	dates := ProductDates{}
	json.Unmarshal(product.Date, &dates)

	//To Get the transaction TimeStamp from the Channel Header
	txTimeAsPtr, errTx := t.GetTxTimestampChannel(APIstub)
	if errTx != nil {
		return shim.Error("Returning error in Transaction TimeStamp")
	}

	dates.SendToDistributorDate = txTimeAsPtr
	datesAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}

	product.Distributer_ID = user.User_ID
	product.Date = datesAsBytes
	updatedProductAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}

	// var newProductDetails = Product{ 
	// 	Product_Id: product.product_ID, 
	// 	Order_ID: product.Order_ID, 
	// 	Name: product.Name, 
	// 	Status: product.Status, 
	// 	Price: product.Price,
	// 	Manufacturer_ID: product.Manufacturer_ID,
	// 	Wholesaler_ID: product.Wholesaler_ID,
	// 	Distributer_ID:user.User_ID,
	// 	Retailer_ID: product.Retailer_ID,
	// 	Consumer_ID: product.Consumer_ID,
	// 	Date:
	// }

	errPut := APIstub.PutState(product.Product_Id, updatedProductAsBytes)
	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to Send to Distributor: %s", product.Product_ID.Product_Id))
	}

	fmt.Println("Success in updating Product Asset %v ", product.Product_ID)
	return shim.Success(nil)
}

// GetTxTimestampChannel Function gets the Transaction time when the chain code was executed it remains same on all the peers where chaincode executes
func (t *food_supplychain) GetTxTimestampChannel(APIstub shim.ChaincodeStubInterface) (string, error) {
	txTimeAsPtr, err := APIstub.GetTxTimestamp()
	if err != nil {
		fmt.Printf("Returning error in TimeStamp \n")
		return "Error", err
	}
	fmt.Printf("\t returned value from APIstub: %v\n", txTimeAsPtr)
	timeStr := time.Unix(txTimeAsPtr.Seconds, int64(txTimeAsPtr.Nanos)).String()

	return timeStr, nil
}


// query all 
func (t *food_supplychain) queryAll(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	startKey := ""

	endKey := ""

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)

	if err != nil {

		return shim.Error(err.Error())

	}

	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults

	var buffer bytes.Buffer

	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {

		queryResponse, err := resultsIterator.Next()
		// respValue := string(queryResponse.Value)
		if err != nil {

			return shim.Error(err.Error())

		}

		// Add a comma before array members, suppress it for the first array member

		if bArrayMemberAlreadyWritten == true {

			buffer.WriteString(",")

		}

		buffer.WriteString("{\"Key\":")

		buffer.WriteString("\"")

		buffer.WriteString(queryResponse.Key)

		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")

		// Record is a JSON object, so we write as-is

		buffer.WriteString(string(queryResponse.Value))

		buffer.WriteString("}")

		bArrayMemberAlreadyWritten = true

	}

	buffer.WriteString("]")

	fmt.Printf("- queryAllAssets:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}
