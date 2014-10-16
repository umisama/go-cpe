# umisama/go-cpe [![Build Status](https://travis-ci.org/umisama/go-cpe.svg)](https://travis-ci.org/umisama/go-cpe)
A Common Platform Enumeration 2.3 implementation for golang.

## installation
```
go get github.com/umisama/go-cpe
```

## usage
Simple example is here:
```go
// Create new item with Setter functions.
item1 := cpe.NewItem(
item1.SetPart(cpe.Application)
item1.SetVendor(cpe.NewStringAttr("microsoft"))
item1.SetProduct(cpe.NewStringAttr("internet_explorer"))

// create new item with WFN.  You can use also other formats(url, formatted string and WFN).
item2, err := cpe.NewItemFromWfn(`wfn:[part="a",vendor="microsoft",product="internet_explorer",version="8\.0\.6001",update="beta"]`)
if err != nil {
        panic(err)
}
fmt.Println("Vendor :", item2.Vendor()

// Compare functions
fmt.Println("is relation superset between item1 and item2? : ", cpe.CheckSuperset(item1, item2))
fmt.Println("is relation equal between item1 and item2? : ", cpe.CheckEqual(item1, item2))
))
```

## document
[godoc.org](http://godoc.org/github.com/umisama/go-cpe)


## reference
 * [NIST IR 7695 — Common Platform Enumeration: Naming Specification Version 2.3](http://csrc.nist.gov/publications/nistir/ir7695/NISTIR-7695-CPE-Naming.pdf)
 * [NIST IR 7696 — Common Platform Enumeration: Name Matching Specification Version 2.3](http://csrc.nist.gov/publications/nistir/ir7696/NISTIR-7696-CPE-Matching.pdf)

## license 
under the MIT License.
