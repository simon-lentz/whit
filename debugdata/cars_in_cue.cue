import (
    "strings" 
    "time"
)
#Car: {
    #RegisteredVehicle
}
#Person: {
    #Entity
    birthday: #timeStamp
    name: string
    Retired?: #Retired
}
#Retirement: {
    id: string
}
#Retired: {
    HAS_RETIRED_Retirement?: [...#EDGE_HAS_RETIRED_Retirement]
}
#MC: {
    #RegisteredVehicle
}
#Registered: {
    regNbr: string
    registrationDate: #timeStamp
    deregisterationDate?: #timeStamp
    REGISTERED_WITH_Registrar?: {
        WHERE: #REF_TO_REGISTRAR
    }
}
#Registrar: {
    name: "DMV"|"ABC"
}
#Vehicle: {
    color?: string
    model?: string
}
#RegisteredVehicle: {
    #Vehicle
    #Registered
}
#Bicycle: {
    #Vehicle
    bikeID: string
}
#MotorVehicleOwner: {
    OWNS_VEHICLE_RegisteredVehicle?: [...#EDGE_OWNS_VEHICLE_RegisteredVehicle]
}
#Entity: {
    MotorVehicleOwner?: #MotorVehicleOwner
}
#Company: {
    #Entity
}
#Graph: {
Cars?: [...#Car]
People?: [...#Person]
Retirements?: [...#Retirement]
Retireds?: [...#Retired]
MCS?: [...#MC]
Registrars?: [...#Registrar]
Bicycles?: [...#Bicycle]
Companies?: [...#Company]
}
#timeStamp: string
#EDGE_HAS_RETIRED_Retirement: {
    fromDate: #timeStamp
    WHERE: #REF_TO_RETIREMENT
}
#EDGE_REGISTERED_WITH_Registrar: {
    WHERE: #REF_TO_REGISTRAR
}
#EDGE_OWNS_VEHICLE_RegisteredVehicle: {
    fromDate: #timeStamp
    toDate?: #timeStamp
    WHERE: #REF_TO_REGISTERED_VEHICLE
}
#REF_TO_REGISTERED_VEHICLE: {
    regNbr: string
}
#REF_TO_REGISTRAR: {
    name: "DMV"|"ABC"
}
#REF_TO_RETIREMENT: {
    id: string
}
