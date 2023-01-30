package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30
	fmt.Println(rand.Intn(max-min+1) + min)

	var myArray [100]ObdSignal

	signaljson := `[
						{"SignalName":"canbus_OBDEngineSpeed_FORD201", "SignalType": "OBDEngineSpeed" , "DataType": "int", "MinVal": 0, "MaxVal": 5000},
						{"SignalName":"canbus_OBDThrottlePosition_FORD201", "SignalType": "OBDThrottlePosition" , "DataType": "int", "MinVal": 0, "MaxVal": 100},
						{"SignalName":"canbus_odometer_CHRYSLER310", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_FORD159", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_FORD179", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_FORD3A0", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_FORD3A0", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_FORD430", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_GM120", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_GM120", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_GM2154561536", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_HONDA1F4", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_HONDA516", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_HYUNDAIKIA4F0", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_HYUNDAIKIA56E", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_HYUNDAIKIA5B0", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_HYUNDAIKIA5D2", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_HYUNDAIKIA5D7", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_NISSAN5C5", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_TOYOTA611", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_VW520", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_VW6B2", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometer_VW6B7", "SignalType": "odometer" , "DataType": "int", "MinVal": 0, "MaxVal": 350000},
						{"SignalName":"canbus_odometerLEFT_HYUNDAIKIA1F1", "SignalType": "odometerLEFT" , "DataType": "int", "MinVal": 0, "MaxVal": 5000}
						]`

	//var mysignals []OBD_Signal
	//var mysignals = myArray[:]
	//mysignals = myArray[0:27]
	var mysignals = make([]ObdSignal, 1)

	//var mysignals = &myArray
	//json.Unmarshal([]byte(signaljson), &mysignals)

	//var mysignals = &myArray
	json.Unmarshal([]byte(signaljson), &mysignals)

	//json.Marshal()

	fmt.Println("test")
	fmt.Println(myArray[0].SignalName)
	fmt.Println(mysignals[0].SignalName)

	//fmt.Println(mysignals[0])
	fmt.Println(signaljson)
	//fmt.Println(GetRandomCanSignal(&mysignals))

	var x = map[string]map[string]map[string]any{}

	x["signals"] = map[string]map[string]any{}

	for i := 0; i < len(mysignals); i++ {
		x["signals"][mysignals[i].SignalName] = map[string]any{}

		x["signals"][mysignals[i].SignalName]["value"] = mysignals[i].GetRandomValue()
		x["signals"][mysignals[i].SignalName]["_stamp"] = time.Now().Format(time.RFC3339Nano)
	}
	jsonStr, err := json.Marshal(x)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		os.WriteFile("temp.json", jsonStr, 0666)
		fmt.Print(string(jsonStr))
		fmt.Print(string("\n"))
	}

	for i := 0; i < len(mysignals); i++ {
		//fmt.Println(i, ":   ", v)
		fmt.Println(i, ":   ")
		GetRandomCanSignal(mysignals)
	}
	fmt.Println(len(mysignals))

}

type signals struct {
}

/*
type ObdSignal_Val_and_TS struct {
	Stamp time.Time `json:"_stamp"`
	Value any       `json:"value"`
}

func NewObdSignal_Val_and_TS(stamp time.Time, value ObdSignalValue) *ObdSignal_Val_and_TS {
	return &ObdSignal_Val_and_TS{Stamp: stamp, Value: value}
}*/

type ObdSignalValue interface {
	int | float64 | string
}

func RandomOdometerValue() (v int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(200 - 1 + 1)
}

type ObdSignal struct {
	SignalName string   `json:"SignalName"`
	SignalType string   `json:"SignalType"` // odometer, vin, fuelevel, etc
	DataType   string   `json:"DataType"`   // int, float, string
	MinVal     int      `json:"MinVal"`     // max value for this signal
	MaxVal     int      `json:"MaxVal"`     // min value for this signal
	Options    []string `json:"Options"`    // available choices for this signal value
}

/*
func ReturnPopulatedMap() map[string]string {
	mymap map[string]string =  map[string]string{}
}*/

// This method reads an ObdSignal struct,
// and generates an appropriate value for the signal type (within a range)
// odometer is first up, will expand into processing other signal types after
func (thisSignal ObdSignal) GetRandomValue() int {
	if thisSignal.SignalType == "odometer" {
		rand.Seed(time.Now().UnixNano())
		min := thisSignal.MinVal
		max := thisSignal.MaxVal
		returnVal := rand.Intn((max - min + 1) + min)
		fmt.Println("getrandomvalue: ", returnVal)
		return returnVal
	} else {
		return 999
	}
}

func GetRandomCanSignal(inputArray []ObdSignal) ObdSignal {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 23
	fmt.Println("random index: ", rand.Intn(max-min+1)+min)
	var returnVal ObdSignal = inputArray[rand.Intn(max-min+1)+min]
	fmt.Println("returnVal:", returnVal)
	return returnVal
}

/*	json.Unmarshal( [
		{"signal_name":"canbus_OBDEngineSpeed_FORD201", "signal_type": "int" , "dataType": "OBDEngineSpeed", "minVal": 7, "maxVal": 22},
		{"signal_name":"canbus_OBDEngineSpeed_FORD201", "signal_type": "int" , "dataType": "OBDEngineSpeed", "minVal": 7, "maxVal": 22},
		{"signal_name":"canbus_OBDEngineSpeed_FORD201", "signal_type": "int" , "dataType": "OBDEngineSpeed", "minVal": 7, "maxVal": 22},
		{"signal_name":"canbus_OBDEngineSpeed_FORD201", "signal_type": "int" , "dataType": "OBDEngineSpeed", "minVal": 7, "maxVal": 22}
	]
		{"canbus_odometer_CHRYSLER310", "int" ,"odometer", 7,22},
		{"canbus_OBDThrottlePosition_FORD201", "string" ,"OBDThrottlePosition", 7,22},
		{"canbus_odometer_FORD159", "int" ,"odometer", 7,22},
		{"canbus_odometer_FORD179", "int" ,"odometer", 7,22},
		{"canbus_odometer_FORD3A0", "int" ,"odometer", 7,22},
	], &signaldata)


data_properties_row := {"canbus_OBDEngineSpeed_FORD201" string ,"OBDEngineSpeed"string ,7,22,"int"}


)data_properties := {"canbus_OBDEngineSpeed_FORD201"  ,"OBDEngineSpeed","int",,)
"canbus_OBDThrottlePosition_FORD201"	,OBDThrottlePosition,7,27,int,,
canbus_odometer_CHRYSLER310		,odometer,7,16,int,,
canbus_odometer_FORD159			,odometer,7,16,int,,
canbus_odometer_FORD179			,odometer,7,16,int,,
canbus_odometer_FORD3A0			,odometer,7,16,int,,
canbus_odometer_FORD3A0			,odometer,7,16,int,,
canbus_odometer_FORD430			,odometer,7,16,int,,
canbus_odometer_GM120				,odometer,7,16,int,,
canbus_odometer_GM120				,odometer,7,16,int,,
canbus_odometer_GM2154561536		,odometer,7,16,int,,
canbus_odometer_HONDA1F4			,odometer,7,16,int,,
canbus_odometer_HONDA516			,odometer,7,16,int,,
canbus_odometer_HYUNDAIKIA4F0		,odometer,7,16,int,,
canbus_odometer_HYUNDAIKIA56E		,odometer,7,16,int,,
canbus_odometer_HYUNDAIKIA5B0		,odometer,7,16,int,,
canbus_odometer_HYUNDAIKIA5D2		,odometer,7,16,int,,
canbus_odometer_HYUNDAIKIA5D7		,odometer,7,16,int,,
canbus_odometer_NISSAN5C5			,odometer,7,16,int,,
canbus_odometer_TOYOTA611			,odometer,7,16,int,,
canbus_odometer_VW520	,odometer,7,16,int,,
canbus_odometer_VW6B2				,odometer,7,16,int,,
canbus_odometer_VW6B7				,odometer,7,16,int,,
canbus_odometerLEFT_HYUNDAIKIA1F1 ,odometerLEFT,7,20,int,,
canbus_odometerRIGHT_HYUNDAIKIA1F1 ,odometerRIGHT,7,21,int,,
canbus_odometerUnit_TOYOTA611		,odometerUnit,7,20,string,,
canbus_VehicleSpeed_FORD201		,VehicleSpeed,7,20,int,,
canbus_VehicleSpeed_HYUNDAIKIA56E	,VehicleSpeed,7,20,int,,
canbus_vin_TOYOTA580			,vin,7,11,string,,
canbus_vin_TOYOTA580v1				,vin,7,11,string,,
canbus_vin_TOYOTA580v2				,vin,7,11,string,,
canbus_vin_TOYOTA580v3				,vin,7,11,string,,
canbus_vin_TOYOTA580v4				,vin,7,11,string,,
canbus_vin_TOYOTA580v5				,vin,7,11,string,,
canbus_vin_TOYOTA580v6				,vin,7,11,string,,
canbus_vin_TOYOTA580v7				,vin,7,11,string,,
canbus_vin_TOYOTA580v8				,vin,7,11,string,,
canbus_vin_TOYOTA581v10			,vin,7,11,string,,
canbus_vin_TOYOTA581v11			,vin,7,11,string,,
canbus_vin_TOYOTA581v12			,vin,7,11,string,,
canbus_vin_TOYOTA581v13			,vin,7,11,string,,
canbus_vin_TOYOTA581v14			,vin,7,11,string,,
canbus_vin_TOYOTA581v15			,vin,7,11,string,,
canbus_vin_TOYOTA581v16			,vin,7,11,string,,
canbus_vin_TOYOTA581v9				,vin,7,11,string,,
canbus_vin_TOYOTA582v17			,vin,7,11,string,,
canbus_vin_VW6B4vKS_Geheimnis_1 m0	,vin,7,11,string,,
canbus_vin_VW6B4vKS_Geheimnis_2 m0	,vin,7,11,string,,
canbus_vin_VW6B4vKS_Geheimnis_3 m0	,vin,7,11,string,,
canbus_vin_VW6B4vKS_Geheimnis_4 m0	,vin,7,11,string,,
canbus_vin_VW6B4vVIN_				,vin,7,11,string,,
canbus_vin_VW6B4vVIN_01_MUX M		,vin,7,11,string,,
canbus_vin_VW6B4vVIN_1 m0			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_10 m1			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_11 m2			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_12 m2			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_13 m2			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_14 m2			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_15 m2			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_16 m2			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_17 m2			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_2 m0			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_3 m0			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_4 m1			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_5 m1			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_6 m1			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_7 m1			,vin,7,11,string,,
canbus_vin_VW6B4vVIN_9 m1			,vin,7,11,string,,
canbus_vinP1_GM514					,vinP1,7,13,string,,
canbus_vinP2_GM4E1					,vinP2,7,13,string,,
}*/
