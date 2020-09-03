package main

import (
       "encoding/json"
        "fmt"
        "bytes"
        "strconv"
        //"strings"
        "time"

        "github.com/hyperledger/fabric/core/chaincode/shim"
        pb "github.com/hyperledger/fabric/protos/peer"
)

//      "strings"

type Certificate struct {
        CertificateId           string  `json:"CertificateId"`
        FileHash                string  `json:"FileHash"`
        FilePath                string  `json:"FilePath"`
        DataHash                string  `json:"DataHash"`
        CertType                string  `json:"CertType"`
        Status                  string  `json:"Status"`
        Certifier1              string  `json:"Certifier1"`
        Certifier2              string  `json:"Certifier2"`
        Certifier3              string  `json:"Certifier3"`
        StudentAck              string  `json:"StudentAck"`
        AllValues               string  `json:"AllValues"`
        CurrentOwner            string  `json:"CurrentOwner"`
        TransferTo              string  `json:"TransferTo"`
}

type CertificateChaincode struct {
}
/*
 * The Init method is called when the Smart Contract "certificate" is instantiated by the blockchain network
 */
func (t *CertificateChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
        return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "certificate"
 */
func (t *CertificateChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
        fmt.Println("chaincode_custom Invoke")
        function, args := stub.GetFunctionAndParameters()

        if function == "query" {
                return t.query(stub, args)
        } else if function == "addActive" {
                return t.addActive(stub, args)
        } else if function == "addPassedOut" {
                        return t.addPassedOut(stub, args)
        } else if function == "studentAcknowledgement" {
                return t.studentAcknowledgement(stub, args)
        } else if function == "updateCertificate" {
                return t.updateCertificate(stub, args)
        } else if function == "updateOwner" {
                return t.updateOwner(stub, args)
        } else if function == "getHistory" {
                return t.getHistory(stub, args)
        }

        return shim.Error("Invalid invoke function name")
}


func (t *CertificateChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        var A string // Entities
        var err error

        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting certificate Id to query")
        }

        A = args[0]

        // Get the state from the ledger
        Avalbytes, err := stub.GetState(A)
        if err != nil {
                jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
                return shim.Error(jsonResp)
        }

        if Avalbytes == nil {
                jsonResp := "{\"Error\":\"Nil amount\"}"
                return shim.Error(jsonResp)
        }

        return shim.Success(Avalbytes)
}

func (t *CertificateChaincode) addActive(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        fmt.Println("Add Certificate.. ")
        var certificate Certificate
        var err error

        isExistAsBytes, _ := stub.GetState(args[0])
        if isExistAsBytes != nil {
                return shim.Error(err.Error())
        }

        certificate.CertificateId = args[0]
        certificate.FileHash      = args[1]
        certificate.FilePath      = args[2]
        certificate.Certifier1    = args[3]
        certificate.Certifier2    = args[4]
        certificate.Certifier3    = args[5]
        certificate.AllValues     = args[6]
        certificate.CurrentOwner  = args[7]

        certificateAsBytes, _ := json.Marshal(certificate)
        err = stub.PutState(certificate.CertificateId, certificateAsBytes)
        if err != nil {
                return shim.Error(err.Error())
        }

        return shim.Success(nil)
}

func (t *CertificateChaincode) studentAcknowledgement(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        fmt.Println("Student Acknowledgement...")
        var certificate Certificate
        var err error

        certificate.CertificateId = args[0]
        certificate.StudentAck = args[1]

        certificateAsBytes, _ := json.Marshal(certificate)
        err = stub.PutState(certificate.CertificateId, certificateAsBytes)

        if err != nil {
                return shim.Error(err.Error())
        }

        return shim.Success(nil)
}

func (t *CertificateChaincode) addPassedOut(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        fmt.Println("Add Certificate.. ")
        var certificate Certificate
        var err error

        isExistAsBytes, _ := stub.GetState(args[0])
        if isExistAsBytes != nil {
                return shim.Error(err.Error())
        }

        certificate.CertificateId = args[0]
        certificate.FileHash = args[1]
        certificate.FilePath = args[2]
        certificate.DataHash = args[3]
        certificate.CertType = args[4]
        certificate.Status   = args[5]

        certificateAsBytes, _ := json.Marshal(certificate)
        err = stub.PutState(certificate.CertificateId, certificateAsBytes)
        if err != nil {
                return shim.Error(err.Error())
        }

        return shim.Success(nil)
}

func (t *CertificateChaincode) updateCertificate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        if len (args) < 2 {
                return shim.Error("Incorrect number of arguments. Expecting 2")
        }
        fmt.Println("Update Certificate")
        var err error
        certificateId := args[0]
        transferTo := args[1]

        certificateAsBytes, err := stub.GetState(certificateId)
        if err != nil {
                return shim.Error("Failed to get Certificate:" + err.Error())
        } else if certificateAsBytes == nil {
                return shim.Error("Certificate does not exist")
        }
        certificateToTransfer := Certificate{}
        err = json.Unmarshal(certificateAsBytes, &certificateToTransfer)
        if err != nil {
                return shim.Error(err.Error())
        }
        certificateToTransfer.TransferTo = transferTo
        certificateJSONAsBytes, _ := json.Marshal(certificateToTransfer)
        err = stub.PutState(certificateId, certificateJSONAsBytes)
        if err != nil {
                return shim.Error(err.Error())
        }
        return shim.Success(nil)
}

func (t *CertificateChaincode) updateOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        if len (args) < 3 {
                return shim.Error("Incorrect number of arguments. Expecting 2.")
        }
        fmt.Println("Update TransferTo")
        var err error
        certificateId := args[0]
        currentOwner := args[1]
        transferTo := args[2]

        certificateAsBytes, err := stub.GetState(certificateId)
        if err != nil {
                return shim.Error("Failed to get Certificate:" + err.Error())
        } else if certificateAsBytes == nil {
                return shim.Error("Certificate does not exist")
        }
        certificateToTransfer := Certificate{}
        err = json.Unmarshal(certificateAsBytes, &certificateToTransfer)
        if err != nil {
                return shim.Error(err.Error())
        }
        certificateToTransfer.CurrentOwner = currentOwner
        certificateToTransfer.TransferTo = transferTo
        certificateJSONAsBytes, _ := json.Marshal(certificateToTransfer)
        err = stub.PutState(certificateId, certificateJSONAsBytes)
        if err != nil {
                return shim.Error(err.Error())
        }
        return shim.Success(nil)
}

func (t *CertificateChaincode) getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        if len(args) < 1 {
                return shim.Error("Incorrect number of arguments. Expectung 1")
        }

        certificateId := args[0]
        fmt.Println("Get History")

        resultIterator, err := stub.GetHistoryForKey(certificateId)
        if err != nil {
                return shim.Error(err.Error())
        }

        defer resultIterator.Close()

        var buffer bytes.Buffer
        buffer.WriteString("[")

        bArrayMemberAlreadyWritten := false
        for resultIterator.HasNext() {
                response, err := resultIterator.Next()
                if err != nil {
                        return shim.Error(err.Error())
                }
                if bArrayMemberAlreadyWritten == true {
                        buffer.WriteString(",")
                }
                buffer.WriteString("{\"TxId\":")
                buffer.WriteString("\"")
                buffer.WriteString(response.TxId)
                buffer.WriteString("\"")
                buffer.WriteString(", \"Value\":")
                if response.IsDelete {
                        buffer.WriteString("null")
                } else {
                        buffer.WriteString(string(response.Value))
                }
                buffer.WriteString(", \"TimeStamp\":")
                buffer.WriteString("\"")
                buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
                buffer.WriteString("\"")
                buffer.WriteString(", \"IsDelete\":")
                buffer.WriteString("\"")
                buffer.WriteString(strconv.FormatBool(response.IsDelete))
                buffer.WriteString("\"")

                buffer.WriteString("}")
                bArrayMemberAlreadyWritten = true
                }
        buffer.WriteString("]")
        fmt.Println("getHistory resturning")

        return shim.Success(buffer.Bytes())
}

func main() {
        err := shim.Start(new(CertificateChaincode))
        if err != nil {
                fmt.Printf("Error starting Certificate chaincode: %s", err)
        }
}
