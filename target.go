package main

import (
  "errors"
  "net/http"
  "fmt"
)

// Object return when a Target is checked
type Result struct {
    Status bool
    AvgTime int
    NbRetry int
    Reason string
  }

// Interface that must be implement by each Target types
type Target interface {
    Register() error
    Check() (*Result, error)
    Remove() error
  }

// TCPTarget: looking for open TCP port on a given IP address
type TCPTarget struct {
    IP string
    Port int
  }

  // UDPTarget: looking for open UDP port on a given IP address
type UDPTarget struct {
    IP string
    Port int
  }

  // HTTPTarget: looking for HTTP service liveness (URL + HTTP response code)
type HTTPTarget struct {
    Addr string
    StatusCode int
    HTTPMethod string
  }

func (ht *HTTPTarget) Check() (*Result, error) {
  tRes := new(Result)
  switch ht.HTTPMethod {
  case "GET":
    res, err := http.Get(ht.Addr) 
    if err != nil {
      return tRes, errors.New(fmt.Sprintf("Error during HTTP GET on taget %s", ht.Addr))
    }
    if ht.StatusCode != res.StatusCode {
      tRes.Status = false
      tRes.Reason = fmt.Sprintf("Expected %d got %d status code", ht.StatusCode, res.StatusCode) 
    }else{
      tRes.Status = true
    }
    return tRes, nil
  //case "POST":
  //case "UPDATE":
  //case "DELETE":
  default:
    return tRes, errors.New("Unsuppored method") 
  }
}  
