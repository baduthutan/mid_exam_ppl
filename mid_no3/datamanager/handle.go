package datamanager

import (
    "encoding/json"
    "net/http" 
)

// Structure of data to be handled in the application.
type Data struct {
    Nama string 
    Nim string 
    Alamat string 
}

// Slice of Data type used to store all the data entered in the application.
var Datas []Data

// Display all the data present in the application.
func DisplayEveryData(w http.ResponseWriter, r *http.Request) {
    // Convert Datas slice to JSON format
    data, err := json.Marshal(Datas)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // Set header of the response to "application/json"
    w.Header().Set("Content-Type", "application/json")
    // Write JSON encoded data to response
    w.Write(data)
}

// Append new data to the Datas slice.
func AppendData(w http.ResponseWriter, r *http.Request) {
    // Create a new Data object to store incoming data
    var data Data
    // Decode JSON data from the request body and store it in data object
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Append the new data to the Datas slice
    Datas = append(Datas,data)
    // Set the response status code to 201 (Created)
    w.WriteHeader(http.StatusCreated)
}