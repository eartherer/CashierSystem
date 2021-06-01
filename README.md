# Cashier System REST API 

This is a Cashier API

# Feature
-   Purchase System
-   Add new Banknote/Coin type to Cashier Storage
-   Refill Banknote/Coin type to Cashier Storage


## Run the app

    go run .

## Run the tests

    go test .

# REST API

The REST API to support above feature

## Get Cashier Information

### Request

`GET /cashier`

### Response

    {
    "CashierName": "Cashier01",
    "Storage": {
        "StorageName": "Storage01",
        "BankNoteInfo": {
            "BankNote100": {
                "Name": "BankNote100",
                "Value": 100,
                "Quantity": 15,
                "MaxQuantity": 15
            },
            "BankNote1000": {
                "Name": "BankNote1000",
                "Value": 1000,
                "Quantity": 10,
                "MaxQuantity": 10
            },
            "BankNote20": {
                "Name": "BankNote20",
                "Value": 20,
                "Quantity": 30,
                "MaxQuantity": 30
            },
            "BankNote50": {
                "Name": "BankNote50",
                "Value": 50,
                "Quantity": 20,
                "MaxQuantity": 20
            },
            "BankNote500": {
                "Name": "BankNote500",
                "Value": 500,
                "Quantity": 20,
                "MaxQuantity": 20
            },
            "Coin0.25": {
                "Name": "Coin0.25",
                "Value": 0.25,
                "Quantity": 50,
                "MaxQuantity": 50
            },
            "Coin1": {
                "Name": "Coin1",
                "Value": 1,
                "Quantity": 20,
                "MaxQuantity": 20
            },
            "Coin10": {
                "Name": "Coin10",
                "Value": 10,
                "Quantity": 20,
                "MaxQuantity": 20
            },
            "Coin5": {
                "Name": "Coin5",
                "Value": 5,
                "Quantity": 20,
                "MaxQuantity": 20
            }
        },
        "AllBankNote": [
            "BankNote1000",
            "BankNote500",
            "BankNote100",
            "BankNote50",
            "BankNote20",
            "Coin10",
            "Coin5",
            "Coin1",
            "Coin0.25"
        ]
    }
}

## Create a new Thing

### Request

`POST /cashier/purchase?paymentAmount={amount}&netAmount={amount}}`

    /cashier/purchase?paymentAmount=250&netAmount=189

### Response

    {
        "Change": 61,
        "BanknoteChange": [
            {
                "Name": "BankNote50",
                "Value": 50,
                "Quantity": 1
            },
            {
                "Name": "Coin10",
                "Value": 10,
                "Quantity": 1
                    },
                    {
                        "Name": "Coin1",
                        "Value": 1,
                        "Quantity": 1
                    }
        ]
    }
## Refill Banknote to storage

### Request

`POST /cashier/refill?name={Banknote Name}&quantity={quantity}`

    /cashier/refill?name=Coin10&quantity=1

### Response

    {
        "Message": "Refill Banknote success"
    }
