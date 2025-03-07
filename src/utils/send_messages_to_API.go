package utils

import (
    "bytes"
    _"encoding/json"
    "net/http"
    "log"
)

func SendMessageToAPI(message string) {
    apiURL := "http://127.0.0.1:9090/v1/notifications/"
    
    // No necesitas convertir el mensaje a JSON nuevamente, ya es una cadena JSON
    jsonData := []byte(message)

    resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Printf("Error sending message to API: %s", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        log.Printf("Error: API responded with status code %d", resp.StatusCode)
    } else {
        log.Printf("Message sent to API successfully")
    }
}
