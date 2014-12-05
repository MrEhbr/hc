package model

// Base interface for all accessories
type Accessory interface {
    Compareable
    
    // Returns the accessory id
    GetId() int64
    
    // Returns the services which represent the accessory 
    GetServices()[]Service
    
    // Returns the name
    Name() string
    
    // Returns the serial number
    SerialNumber() string
    
    // Returns the manufacturer name
    Manufacturer() string
    
    // Returns the model description
    Model() string
}