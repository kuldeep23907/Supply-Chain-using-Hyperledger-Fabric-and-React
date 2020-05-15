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

func (t *food_supplychain) Init(stub shim.ChaincodeStubInterface) pb.Response{ 	
	// Initializing Product Counter
	ProductCounterBytes, _ := APIstub.GetState("ProductCounterNO")
	if ProductCounterBytes == nil {
		var ProductCounter = CounterNO{Counter: 0}
		ProductCounterBytes, _ := json.Marshal(ProductCounter)
		err := APIstub.PutState("ProductCounterNO", ProductCounterBytes)
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to Intitate Product Counter"))
		}
	}
  // Initializing Order Counter
	OrderCounterBytes, _ := APIstub.GetState("OrderCounterNO")
	if OrderCounterBytes == nil {
		var OrderCounter = CounterNO{Counter: 0}
		OrderCounterBytes, _ := json.Marshal(OrderCounter)
		err := APIstub.PutState("OrderCounterNO", OrderCounterBytes)
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to Intitate Order Counter"))
		} 
	}
	UserCounterBytes, _ := APIstub.GetState("UserCounterNO")
	if UserCounterBytes == nil {
		var UserCounter = CounterNO{Counter: 0}
		UserCounterBytes, _ := json.Marshal(UserCounter)
		err := APIstub.PutState("UserCounterNO", UserCounterBytes)
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to Intitate User Counter"))
		} 
	}
}
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

func (t *food_supplychain) createProduct(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
 
	//To check number of arguments are 11
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments, Required 4 arguments")
	}

	if len(args[1]) == 0 {
		return shim.Error("Name must be provided to create a product")
	}

	if len(args[2]) == 0 {
		return shim.Error("Manufacturer_ID must be provided")
	}
  
	if len(args[3]) == 0 {
		return shim.Error("Price must be non-empty ")
	}

	//Price conversion - Error handeling 
	i1, errPrice := strconv.ParseFloat(args[3], 64)
	if errPrice != nil {
		return shim.Error(fmt.Sprintf("Failed to Convert Price: %s", errPrice))
		//fmt.Println(i1)
	}

	productCounter := getCounter(APIstub,"ProductCounterNO")
	productCounter++
	
	// DATES
	dates := ProductDates{}
	json.Unmarshal(product.Date, &dates)

	//To Get the transaction TimeStamp from the Channel Header
	txTimeAsPtr, errTx := t.GetTxTimestampChannel(APIstub)
	if errTx != nil {
		return shim.Error("Returning error in Transaction TimeStamp")
	}

	dates.ManufactureDate = txTimeAsPtr

	var comAsset = Product{Product_Id: "Product" + strconv.Itoa(productCounter), Order_ID: nil, Name: args[1], Consumer_ID: nil, Manufacturer_ID: args[2], Retailer_ID: nil, Distributer_ID: nil, Wholesaler_ID: nil, Status: "Available", Date: dates.ManufactureDate, Product_Price: i1}
	
	comAssetAsBytes, errMarshal := json.Marshal(comAsset)
   
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error in Product: %s", errMarshal))
	}
  
	errPut := APIstub.PutState(comAsset.Product_Id, comAssetAsBytes)
  
	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to create Product Asset: %s", comAsset.Product_Id))
	}
  
	//TO Increment the Product Counter
	incrementCounter(APIstub,"ProductCounterNO")
  
	fmt.Println("Success in creating Product Asset %v",comAsset)
  
	return shim.Success(nil)
  
}

// send to Wholesaler
func (t *food_supplychain) sendToWholesaler(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Less no of arguements provided");
	}

	if len(args[1]) == 0 {
		return shim.Error("Product Id must be provided");
	}

	if len(args[2]) == 0 {
		return shim.Error("Wholesaler Id must be provided");
	}

	userBytes, _ := APIstub.GetState(args[2])

	if userBytes == nil {
		return shim.Error("Cannot Find Wholesaler user")
	}

	user := User{}

	json.Unmarshal(userBytes, &user)

	if user.User_Type != "wholesaler" {
		return shim.Error("User type must be Wholesaler")
	}


	productBytes, _ := APIstub.GetState(args[1])

	if productBytes == nil {
		return shim.Error("Cannot Find Product")
	}

	product := Product{}

	json.Unmarshal(productBytes, &product)

	if product.Wholesaler_ID != nil {
		return shim.Error("Product is send to Wholesaler already")
	}

	dates := ProductDates{}
	json.Unmarshal(product.Date, &dates)

	//To Get the transaction TimeStamp from the Channel Header
	txTimeAsPtr, errTx := t.GetTxTimestampChannel(APIstub)
	if errTx != nil {
		return shim.Error("Returning error in Transaction TimeStamp")
	}

	dates.SendToWholesalerDate = txTimeAsPtr
	datesAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}

	product.Wholesaler_ID = user.User_ID
	product.Date = datesAsBytes
	updatedProductAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}

	errPut := APIstub.PutState(product.Product_ID, updatedProductAsBytes)
	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to Send to Wholesaler: %s", product.Product_ID))
	}

	fmt.Println("Success in sending Product %v ", product.Product_ID)
	return shim.Success(nil)
}



// send to retailer
func (t *food_supplychain) sendToRetailer(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Less no of arguments provided")
	}

	if len(args[1]) == 0 {
		return shim.Error("ProductId must be specified")
	}

	if len(args[2]) == 0 {
		return shim.Error("RetailerId must be specified")
	}

	userBytes, _ := APIstub.GetState(args[2])
	if userBytes == nil {
		return shim.Error("Could not find the retailer")
	}

	user := User{}
	json.Unmarshal(userBytes, &user)
	if user.User_Type != "retailer" {
		return shim.Error("User must be a retailer")
	}

	productBytes, _ := APIstub.GetState(args[1])
	if productBytes == nil {
		return shim.Error("Could not find the product")
	}

	product := Product{}
	json.Unmarshal(productBytes, &product)
	if product.Retailer_ID != nil {
		return shim.Error("Product has already been sent to retailer")
	}

	dates := ProductDates{}
	json.Unmarshal(product.Date, &dates)

	// get the transaction timestamp from the channel header
	txTimeAsPtr, errTx := t.GetTxTimestampChannel(APIstub)
	if errTx != nil {
		return shim.Error("Error while retrieving transaction timestamp")
	}

	dates.SendToRetailerDate = txTimeAsPtr
	datesAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal error: %s", errMarshal))
	}

	product.Retailer_ID = User_ID
	product.Date = datesAsBytes
	updatedProductAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal error: %s", errMarshal))
	}

	errPut := APIstub.PutState(product.Product_ID, updatedProductAsBytes)
	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to send to retailer: %s", product.Product_ID))
	}

	fmt.Println("Sent product %v to retailer successfully", product.Product_ID)
	return shim.Success(nil)
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

	errPut := APIstub.PutState(product.Product_ID, updatedProductAsBytes)
	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to Send to Distributor: %s", product.Product_ID))
	}

	fmt.Println("Success in sending Product %v ", product.Product_ID)
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

func (t *food_supplychain) queryProduct(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expected 1 argument")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(productAsBytes)
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

// function to sell the product to consumer
// Input params , product id  consumer id
func sellToConsumer(APIstub shim.ChaincodeStubInterface,args[] string) pb.Response{

	// parameter length check
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments, Required 2")
	}

	// parameter null check
	if len(args[0]) == 0 {
		return shim.Error("Product Id must be provided");
	}

	if len(args[1]) == 0 {
		return shim.Error("Consumer Id must be provided");
	}

	// get user details from the stub ie. Chaincode stub in network using the user id passed
	userBytes, _ := APIstub.GetState(args[1])

	if userBytes == nil {
		return shim.Error("Cannot Find Consumer user")
	}

	user := User{}
	
	// unmarsahlling product the data from API
	json.Unmarshal(userBytes, &user)

	// User type check for the function
	if user.User_Type != "consumer" {
		return shim.Error("User type must be Consumer")
	}

	// get product details from the stub ie. Chaincode stub in network using the product id passed
	productBytes, _ := APIstub.GetState(args[0])


	if productBytes == nil {
		return shim.Error("Cannot Find Product")
	}

	product := Product{}

	// unmarsahlling product the data from API
	json.Unmarshal(productBytes, &product)

	// check if the product is sold to consumer already
	if product.Consumer_ID != nil {
		return shim.Error("Product is sold to customer already")
	}

	// Updating date for the transaction
	dates := ProductDates{}
	json.Unmarshal(product.Date, &dates)

	//To Get the transaction TimeStamp from the Channel Header
	txTimeAsPtr, errTx := t.GetTxTimestampChannel(APIstub)
	if errTx != nil {
		return shim.Error("Returning error in Transaction TimeStamp")
	}

	dates.SellToConsumerDate = txTimeAsPtr
	datesAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}

	// Updating the product values to be updated after the function
	product.Consumer_ID = user.User_ID
	product.Date = datesAsBytes
	product.Status = "Delivered"
	updatedProductAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}

	errPut := APIstub.PutState(product.Product_ID, updatedProductAsBytes)
	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to Sell To Cosumer : %s", product.Product_ID))
	}

	fmt.Println("Success in sending Product %v ", product.Product_ID)
	return shim.Success(nil)
}


// function to update the product name and price
// Input params : product id , user id , product name , preduct price
func updateProduct(APIstub shim.ChaincodeStubInterface,args[] string) pb.Response{

	// parameter length check
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments, Required 4")
	}

	// parameter null check
	if len(args[0]) == 0 {
		return shim.Error("Product Id must be provided");
	}

	if len(args[1]) == 0 {
		return shim.Error("User Id must be provided");
	}

	if len(args[2]) == 0 {
		return shim.Error("Product Name must be provided");
	}

	if len(args[3]) == 0 {
		return shim.Error("Product Price must be provided");
	}


	// get user details from the stub ie. Chaincode stub in network using the user id passed
	userBytes, _ := APIstub.GetState(args[1])

	if userBytes == nil {
		return shim.Error("Cannot Find User")
	}

	user := User{}

	// unmarsahlling product the data from API
	json.Unmarshal(userBytes, &user)

	// User type check for the function
	if user.User_Type == "consumer" {
		return shim.Error("User type cannot be Consumer")
	}

	// get product details from the stub ie. Chaincode stub in network using the product id passed
	productBytes, _ := APIstub.GetState(args[0])
	if productBytes == nil {
		return shim.Error("Cannot Find Product")
	}
	product := Product{}

	// unmarsahlling product the data from API
	json.Unmarshal(productBytes, &product)

	// Updating the product values withe the new values
	product.Name = args[2] // product name from UI for the update
	product.Price = args[3] // product value from UI for the update

	updatedProductAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}

	errPut := APIstub.PutState(product.Product_ID, updatedProductAsBytes)
	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to Sell To Cosumer : %s", product.Product_ID))
	}

	fmt.Println("Success in sending Product %v ", product.Product_ID)
	return shim.Success(nil)

}

func orderProduct(APIstub shim.ChaincodeStubInterface,args[] string) pb.Response{
	// parameter length check
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments, Required 4")
	}

	// parameter null check
	if len(args[0]) == 0 {
		return shim.Error("Consumer Id must be provided");
	}

	if len(args[1]) == 0 {
		return shim.Error("Product Id must be provided");
	}

	userBytes, _ := APIstub.GetState(args[0])

	if userBytes == nil {
		return shim.Error("Cannot Find Consumer")
	}

	user := User{}

	// unmarsahlling product the data from API
	json.Unmarshal(userBytes, &user)

	// User type check for the function
	if user.User_Type == "consumer" {
		return shim.Error("User type cannot be Consumer")
	}

	productBytes, _ := APIstub.GetState(args[1])
	if productBytes == nil {
		return shim.Error("Cannot Find Product")
	}
	product := Product{}

	// unmarsahlling product the data from API
	json.Unmarshal(productBytes, &product)


	orderCounter := getCounter(APIstub,"OrderCounterNO")
	orderCounter++

	product.Order_ID = "Order" + strconv.Itoa(orderCounter) 
	product.Consumer_ID = args[1] 
	product.Status = "Ordered"

	updatedProductAsBytes, errMarshal := json.Marshal(dates)
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}

	incrementCounter(APIstub,"OrderCounterNO")

	errPut := APIstub.PutState(product.Product_ID, updatedProductAsBytes)
	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to place the order : %s", product.Product_ID))
	}

	fmt.Println("Order placed successfuly %v ", product.Product_ID)
	return shim.Success(nil)
}

func (t *food_supplychain) createUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments, Required 4 arguments")
	}

	if len(args[0]) == 0 {
		return shim.Error("Name must be provided to register user")
	}

	if len(args[1]) == 0 {
		return shim.Error("Email is mandatory")
	}
  
	if len(args[2]) == 0 {
		return shim.Error("User type must be specified")
	}

	if len(args[3]) == 0 {
		return shim.Error("Address must be non-empty ")
	}

	userCounter := getCounter(APIstub,"UserCounterNO")
	userCounter++

	var comAsset = User{Name: args[0], User_ID: "User" + strconv.Itoa(userCounter), Email: args[1], User_Type: args[2], Address: args[3]}
	
	comAssetAsBytes, errMarshal := json.Marshal(comAsset)
   
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error in Product: %s", errMarshal))
	}
  
	errPut := APIstub.PutState(comAsset.Product_Id, comAssetAsBytes)
  
	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to register user: %s", comAsset.Product_Id))
	}
  
	//TO Increment the Product Counter
	incrementCounter(APIstub,"UserCounterNO")
  
	fmt.Println("User register successfully %v",comAsset)
  
	return shim.Success(nil)
  
}
