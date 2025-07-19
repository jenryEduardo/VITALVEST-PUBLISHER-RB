package domain



type Idata interface{
	DataMqtt(data DatosSensor)error
}