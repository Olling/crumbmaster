package main

import (
	"os"
	"sync"
	"io/ioutil"
	"encoding/json"
	"github.com/olling/logger"
)


var (
        outputMutex sync.Mutex
)


func ReadJsonFile (path string, output interface{}) (error) {
	logger.Debug("Reading JSON file:", path)
        file,fileerr := os.Open(path)
	if fileerr != nil {
		logger.Debug("Error opening JSON file for reading: ", fileerr)
		return fileerr
	}

        decoder := json.NewDecoder(file)
        decodererr := decoder.Decode(&output)
	if decodererr != nil {
		logger.Debug("Error decoding JSON file", decodererr)
		return decodererr
	}

	logger.Debug("Done reading JSON file:", path, output)
	return nil
}


func WriteJsonFile(s interface{}, path string) (err error){
	logger.Debug("Writing to path: " + path, "content:", s)
        outputMutex.Lock()
        defer outputMutex.Unlock()

        bytes, marshalErr := json.MarshalIndent(s,"","\t")
        if marshalErr != nil {
                logger.Error("Could not convert struct to bytes", marshalErr)
                return marshalErr
        }

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		logger.Error("Could not write file: " + path + ". Error:", err)
		return err
	}
	return nil
}
