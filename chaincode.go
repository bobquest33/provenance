/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
  "errors"
  "fmt"

  "github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
  err := shim.Start(new(SimpleChaincode))
  if err != nil {
    fmt.Printf("Error starting Simple chaincode: %s.", err)
  }
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  if len(args) != 1 {
    return nil, errors.New("Incorrect number of arguments. Expecting 1.")
  }
  return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("Invoke is running " + function + ".")
  // Handle different functions
  if function == "init" {
    return t.Init(stub, "init", args)
  } else if function == "write" {
    return t.write(stub, args)
  }
  if function == "init" {
    return t.Init(stub, "init", args)
  }
  fmt.Println("Invoke did not find function: " + function + ".")		//error
  return nil, errors.New("Received unknown function invocation.")
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("Query is running " + function + ".")
  // Handle different functions
  if function == "dummy_query" {
    fmt.Println("Hi there " + function + "!")
    return nil, nil;
  } else if function == "read" {
    return t.read(stub, args)
  }
  fmt.Println("Query did not find function: " + function + ".")			//error
  return nil, errors.New("Received unknown function query.")
}

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var key, value string
  var err error
  fmt.Println("Running write().")
  if len(args) != 2 {
    return nil, errors.New("Incorrect number of arguments. Expecting 2. The key and the value to set.")
  }
  key = args[0]
  value = args[1]
  err = stub.PutState(key, []byte(value))					//write the variable into the chaincode state
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var key, jsonResp string
  var err error
  if len(args) != 1 {
    return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query.")
  }
  key = args[0]
  valAsbytes, err := stub.GetState(key)
  if err != nil {
    jsonResp = "{\"Error\":\"Failed to get state for key: " + key + ".\"}"
    return nil, errors.New(jsonResp)
  }
  if valAsbytes == nil {
    jsonResp = "{\"Error\":\"Cannot find BlockChain key: " + key + ".\"}"
    return nil, errors.New(jsonResp)
  }
  return valAsbytes, nil
}
