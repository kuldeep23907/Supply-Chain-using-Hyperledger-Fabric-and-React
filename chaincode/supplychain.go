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
} 

type Product struct { 	
	//ProductId, OrderId, Name,ConsumerId,ManufactureId, 	
	//WholesalerId, RetailerId,status,date,Price. 	
	Product_ID string `json:"ProductID"` 	
	Order_ID string `json:"OrderID"` 	
	Name string `json:"Name"` 	
	Consumer_ID string `json:"ConsumerID"` 	
	Manufacturer_ID string `json:"ManufacturerID"` 	
	Retailer_ID string `json:"RetailerID"` 	
	Distributer_ID string `json:"DistributerID"` 	
	Wholesaler_ID string `json:"WholesalerID"` 	
	Status string `json:"Status"` 	
	Date string `json:"Date"` 	
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
