<p align="center">
  <br>
  <a href="https://bitvavo.com"><img src="https://bitvavo.com/assets/static/ext/logo-shape.svg" width="100" title="Bitvavo Logo"></a>
</p>


# Go Bitvavo API
This is the Go wrapper for the Bitvavo API. This project can be used to build your own projects which interact with the Bitvavo platform. Every function available on the API can be called through a REST request or over websockets. For info on the specifics of every parameter consult the [Bitvavo API documentation](https://docs.bitvavo.com/)

* Getting started       [REST](https://github.com/bitvavo/go-bitvavo-api#getting-started) [Websocket](https://github.com/bitvavo/go-bitvavo-api#getting-started-1)
* General
  * Time                [REST](https://github.com/bitvavo/go-bitvavo-api#get-time) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-time-1)
  * Markets             [REST](https://github.com/bitvavo/go-bitvavo-api#get-markets) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-markets-1)
  * Assets              [REST](https://github.com/bitvavo/go-bitvavo-api#get-assets) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-assets-1)
* Market Data
  * Book           [REST](https://github.com/bitvavo/go-bitvavo-api#get-book-per-market) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-book-per-market-1)
  * Public Trades         [REST](https://github.com/bitvavo/go-bitvavo-api#get-trades-per-market) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-trades-per-market-1)
  * Candles        [REST](https://github.com/bitvavo/go-bitvavo-api#get-candles-per-market) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-candles-per-market-1)
  * Price Ticker        [REST](https://github.com/bitvavo/go-bitvavo-api#get-price-ticker) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-price-ticker-1)
  * Book Ticker         [REST](https://github.com/bitvavo/go-bitvavo-api#get-book-ticker) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-book-ticker-1)
  * 24 Hour Ticker      [REST](https://github.com/bitvavo/go-bitvavo-api#get-24-hour-ticker) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-24-hour-ticker-1)
* Private
  * Place Order         [REST](https://github.com/bitvavo/go-bitvavo-api#place-order) [Websocket](https://github.com/bitvavo/go-bitvavo-api#place-order-1)
  * Update Order        [REST](https://github.com/bitvavo/go-bitvavo-api#update-order) [Websocket](https://github.com/bitvavo/go-bitvavo-api#update-order-1)
  * Get Order           [REST](https://github.com/bitvavo/go-bitvavo-api#get-order) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-order-1)
  * Cancel Order        [REST](https://github.com/bitvavo/go-bitvavo-api#cancel-order) [Websocket](https://github.com/bitvavo/go-bitvavo-api#cancel-order-1)
  * Get Orders          [REST](https://github.com/bitvavo/go-bitvavo-api#get-orders) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-orders-1)
  * Cancel Orders       [REST](https://github.com/bitvavo/go-bitvavo-api#cancel-orders) [Websocket](https://github.com/bitvavo/go-bitvavo-api#cancel-orders-1)
  * Orders Open         [REST](https://github.com/bitvavo/go-bitvavo-api#get-orders-open) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-orders-open-1)
  * Trades              [REST](https://github.com/bitvavo/go-bitvavo-api#get-trades) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-trades-1)
  * Balance             [REST](https://github.com/bitvavo/go-bitvavo-api#get-balance) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-balance-1)
  * Deposit Assets     [REST](https://github.com/bitvavo/go-bitvavo-api#deposit-assets) [Websocket](https://github.com/bitvavo/go-bitvavo-api#deposit-assets-1)
  * Withdraw Assets   [REST](https://github.com/bitvavo/go-bitvavo-api#withdraw-assets) [Websocket](https://github.com/bitvavo/go-bitvavo-api#withdraw-assets-1)
  * Deposit History     [REST](https://github.com/bitvavo/go-bitvavo-api#get-deposit-history) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-deposit-history-1)
  * Withdrawal History  [REST](https://github.com/bitvavo/go-bitvavo-api#get-withdrawal-history) [Websocket](https://github.com/bitvavo/go-bitvavo-api#get-withdrawal-history-1)
* [Subscription](https://github.com/bitvavo/go-bitvavo-api#subscriptions)
  * [Ticker Subscription](https://github.com/bitvavo/go-bitvavo-api#ticker-subscription)
  * [Ticker 24 Hour Subscription](https://github.com/bitvavo/go-bitvavo-api#ticker-24-hour-subscription)
  * [Account Subscription](https://github.com/bitvavo/go-bitvavo-api#account-subscription)
  * [Candles Subscription](https://github.com/bitvavo/go-bitvavo-api#candles-subscription)
  * [Trades Subscription](https://github.com/bitvavo/go-bitvavo-api#trades-subscription)
  * [Book Subscription](https://github.com/bitvavo/go-bitvavo-api#book-subscription)
  * [Book subscription with local copy](https://github.com/bitvavo/go-bitvavo-api#book-subscription-with-local-copy)

## Installation
```
Get:
go get github.com/bitvavo/go-bitvavo-api

Import:
import("github.com/bitvavo/go-bitvavo-api")

Run test:
cd $GOPATH/src/github.com/bitvavo/go-bitvavo-api/example
go run main.go
```

## Rate Limiting

Bitvavo uses a weight based rate limiting system, with an allowed limit of 1000 per IP or API key each minute. Please inspect each endpoint in the [Bitvavo API documentation](https://docs.bitvavo.com/) to see the weight. Failure to respect the rate limit will result in an IP or API key ban.
Since the remaining limit is returned in the header on each REST request, the remaining limit is tracked locally and can be requested through:
```
var limit = bitvavo.GetRemainingLimit()
fmt.Println("The remaining rate limit is", limit)
```
The websocket functions however do not return a remaining limit, therefore the limit is only updated locally once a ban has been issued.


## REST requests

The general convention used in all functions (both REST and websockets), is that all optional parameters are passed as `map[string]string`, while required parameters are passed as separate values. Only when placing orders some of the optional parameters are required, since a limit order requires more information than a market order. The returned responses are all converted to an object, such that `response.<key> = '<value>'`. The definition of the structs is supplied for every function, but the general convention used is that all variables in the struct are defined in UpperCamelCase (PascalCase). On top of the object we also return an error, this error should be checked before handling the response in the following manner:

```go
response, err := bitvavo.Time()
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.Time)
}
```

### Getting started

The API key and secret are required for private calls and optional for public calls. The access window and debugging parameter are optional for all calls. The access window is used to determine whether the request arrived within time, the value is specified in milliseconds. You can use the [time](https://github.com/bitvavo/go-bitvavo-api#get-time) function to synchronize your time to our server time if errors arise. Debugging should be set to true when you want to log additional information and full responses. Any parameter can be omitted, private functions will return an error when the api key and secret have not been set.

```go
import . "bitvavo"

bitvavo := Bitvavo{
    ApiKey: "<APIKEY>",
    ApiSecret: "<APISECRET>",
    AccessWindow: 10000,
    Debugging: true
  }
```

### General

#### Get time
```go
response, err := bitvavo.Time()
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
}
```
<details>
 <summary>View Response</summary>

```go
{
  "time": 1548679180309
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Time struct {
  Time int `json:"time"`
}
```
</details>

#### Get markets
```go
// options: market
response, err := bitvavo.Markets(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, value := range response {
    fmt.Printf("%+v\n", value)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "status": "trading",
  "base": "ADA",
  "quote": "BTC",
  "market": "ADA-BTC",
  "pricePrecision": 5,
  "minOrderInQuoteAsset": "0.001",
  "minOrderInBaseAsset": "100",
  "orderTypes": [
    "market",
    "limit"
  ]
}
{
  "status": "trading",
  "base": "ADA",
  "quote": "EUR",
  "market": "ADA-EUR",
  "pricePrecision": 5,
  "minOrderInQuoteAsset": "10",
  "minOrderInBaseAsset": "100",
  "orderTypes": [
    "market",
    "limit"
  ]
}
{
  "status": "trading",
  "base": "AE",
  "quote": "BTC",
  "market": "AE-BTC",
  "pricePrecision": 5,
  "minOrderInQuoteAsset": "0.001",
  "minOrderInBaseAsset": "10",
  "orderTypes": [
    "market",
    "limit"
  ]
}
{
  "status": "trading",
  "base": "AE",
  "quote": "EUR",
  "market": "AE-EUR",
  "pricePrecision": 5,
  "minOrderInQuoteAsset": "10",
  "minOrderInBaseAsset": "10",
  "orderTypes": [
    "market",
    "limit"
  ]
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Markets struct {
  Status               string   `json:"status"`
  Base                 string   `json:"base"`
  Quote                string   `json:"quote"`
  Market               string   `json:"market"`
  PricePrecision       int      `json:"pricePrecision"`
  MinOrderInQuoteAsset string   `json:"minOrderInQuoteAsset"`
  MinOrderInBaseAsset  string   `json:"minOrderInBaseAsset"`
  OrderTypes           []string `json:"orderTypes"`
}
```
</details>

#### Get assets
```go
// options: symbol
response, err := bitvavo.Assets(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, value := range response {
    fmt.Printf("%+v\n", value)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "ADA",
  "name": "Cardano",
  "decimals": 6,
  "depositFee": "0",
  "depositConfirmations": 20,
  "depositStatus": "OK",
  "withdrawalFee": "0.2",
  "withdrawalMinAmount": "0.2",
  "withdrawalStatus": "OK",
  "networks": [
    "Mainnet"
  ],
  "message": ""
}
{
  "symbol": "AE",
  "name": "Aeternity",
  "decimals": 8,
  "depositFee": "0",
  "depositConfirmations": 30,
  "depositStatus": "OK",
  "withdrawalFee": "2",
  "withdrawalMinAmount": "2",
  "withdrawalStatus": "OK",
  "networks": [
    "Mainnet"
  ],
  "message": ""
}
{
  "symbol": "AION",
  "name": "Aion",
  "decimals": 8,
  "depositFee": "0",
  "depositConfirmations": 0,
  "depositStatus": "",
  "withdrawalFee": "3",
  "withdrawalMinAmount": "3",
  "withdrawalStatus": "",
  "networks": [
    "Mainnet"
  ],
  "message": ""
}
{
  "symbol": "ANT",
  "name": "Aragon",
  "decimals": 8,
  "depositFee": "0",
  "depositConfirmations": 30,
  "depositStatus": "OK",
  "withdrawalFee": "2",
  "withdrawalMinAmount": "2",
  "withdrawalStatus": "OK",
  "networks": [
    "Mainnet"
  ],
  "message": ""
}
 ...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Assets struct {
  Symbol               string   `json:"symbol"`
  Name                 string   `json:"name"`
  Decimals             int      `json:"decimals"`
  DepositFee           string   `json:"depositFee"`
  DepositConfirmations int      `json:"depositConfirmations"`
  DepositStatus        string   `json:"depositStatus"`
  WithdrawalFee        string   `json:"withdrawalFee"`
  WithdrawalMinAmount  string   `json:"withdrawalMinAmount"`
  WithdrawalStatus     string   `json:"withdrawalStatus"`
  Networks             []string `json:"networks"`
  Message              string   `json:"message"`
}
```
</details>

### Market Data

#### Get book per market
```go
// options: depth
response, err := bitvavo.Book("BTC-EUR", map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
}
```
<details>
 <summary>View Response</summary>

```go
{
  "market": "BTC-EUR",
  "nonce": 7831,
  "bids": [
    [
      "2992.4",
      "0.11023153"
    ],
    [
      "2991.5",
      "0.50272746"
    ],
    [
      "2990.6",
      "1.27126107"
    ],
    [
      "2989.6",
      "3.0301821"
    ],
    [
      "2988.6",
      "3.2848159"
    ],
    ...
  ],
  "asks": [
    [
      "2993.3",
      "1.1852825"
    ],
    [
      "2994.1",
      "0.401334"
    ],
    [
      "2994.7",
      "0.31577418"
    ],
    [
      "2995.5",
      "2.30306344"
    ],
    [
      "2996.1",
      "2.63275436"
    ],
    ...
  ]
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Book struct {
  Market string     `json:"market"`
  Nonce  int        `json:"nonce"`
  Bids    [][]string `json:"bids"`
  Asks   [][]string `json:"asks"`
}
```
</details>

#### Get trades per market
```go
// options: limit, start, end, tradeIdFrom, tradeIdTo
response, err := bitvavo.PublicTrades("BTC-EUR", map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, trade := range response {
    fmt.Printf("%+v\n", trade)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "timestamp": 1548678622527,
  "id": "d4e1c700-8432-4ec7-b141-cb6bcee7a613",
  "amount": "2.36697512",
  "price": "3001.8",
  "side": "buy"
}
{
  "timestamp": 1548678622520,
  "id": "17c4ec15-7806-451e-8cc6-8ca48f1eef66",
  "amount": "3.37614478",
  "price": "3001.2",
  "side": "buy"
}
{
  "timestamp": 1548678622514,
  "id": "eb9d8230-6c55-424b-be93-2d41d39ab092",
  "amount": "3.27104588",
  "price": "3000.2",
  "side": "buy"
}
{
  "timestamp": 1548678622506,
  "id": "9ccc74d6-47d9-44f1-a51a-7dbbb3af5373",
  "amount": "0.42395607",
  "price": "2999.6",
  "side": "buy"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type PublicTrades struct {
  Timestamp int    `json:"timestamp"`
  Id        string `json:"id"`
  Amount    string `json:"amount"`
  Price     string `json:"price"`
  Side      string `json:"side"`
}
```
</details>

#### Get candles per market
```go
// options: limit, start, end
response, err := bitvavo.candles("BTC-EUR", "1h", map[string]string{})

if err != nil {
  fmt.Println(err)
} else {
  for _, candle := range response {
    fmt.Printf("%+v\n", candle)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "Timestamp": 1548680400000,
  "Open": "2989",
  "High": "2991.1",
  "Low": "2989",
  "Close": "2989",
  "Volume": "0.9"
}
{
  "Timestamp": 1548676800000,
  "Open": "2999.3",
  "High": "3002.6",
  "Low": "2989.2",
  "Close": "2999.3",
  "Volume": "63.00046504"
}
{
  "Timestamp": 1548669600000,
  "Open": "3012.9",
  "High": "3015.8",
  "Low": "3000",
  "Close": "3012.9",
  "Volume": "8"
}
{
  "Timestamp": 1548417600000,
  "Open": "3124",
  "High": "3125.1",
  "Low": "3124",
  "Close": "3124",
  "Volume": "0.1"
}
{
  "Timestamp": 1548237600000,
  "Open": "3143",
  "High": "3143.3",
  "Low": "3141.1",
  "Close": "3143",
  "Volume": "60.67250851"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Candle struct {
  Timestamp int
  Open      string
  High      string
  Low       string
  Close     string
  Volume    string
}
```
</details>

#### Get price ticker
```go
// options: market
response, err := bitvavo.TickerPrice(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, price := range response {
    fmt.Printf("%+v\n", price)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "market": "EOS-EUR",
  "price": "2.0142"
}
{
  "market": "XRP-EUR",
  "price": "0.25193"
}
{
  "market": "ETH-EUR",
  "price": "99.877"
}
{
  "market": "IOST-EUR",
  "price": "0.005941"
}
{
  "market": "BCH-EUR",
  "price": "106.57"
}
{
  "market": "BTC-EUR",
  "price": "2991.1"
}
{
  "market": "STORM-EUR",
  "price": "0.0025672"
}
{
  "market": "EOS-BTC",
  "price": "0.00066289"
}
{
  "market": "BSV-EUR",
  "price": "57.6"
}
{
  "market": "ETH-BTC",
  "price": "0.032373"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type TickerPrice struct {
  Market string `json:"market"`
  Price  string `json:"price"`
}
```
</details>

#### Get book ticker
```go
// options: market
response, err := bitvavo.TickerBook(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, book := range response {
    fmt.Printf("%+v\n", book)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "market": "XVG-BTC",
  "bid": "0.00000045",
  "ask": "0.00000046",
  "bidSize": "28815.01275017",
  "askSize": "38392.85089495"
}
{
  "market": "XVG-EUR",
  "bid": "0.004213",
  "ask": "0.0043174",
  "bidSize": "1699002.27041019",
  "askSize": "638327.18139947"
}
{
  "market": "ZIL-BTC",
  "bid": "0.00000082",
  "ask": "0.00000083",
  "bidSize": "140980.13397262",
  "askSize": "98568.18059098"
}
{
  "market": "ZIL-EUR",
  "bid": "0.0076771",
  "ask": "0.0077744",
  "bidSize": "320216.82744213",
  "askSize": "157923.96870507"
}
{
  "market": "ZRX-BTC",
  "bid": "0.00001684",
  "ask": "0.000016898",
  "bidSize": "631.2417155",
  "askSize": "1551.90609202"
}
{
  "market": "ZRX-EUR",
  "bid": "0.15766",
  "ask": "0.15826",
  "bidSize": "873.64460869",
  "askSize": "1010.23609323"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type TickerBook struct {
  Market string `json:"market"`
  Bid    string `json:"bid"`
  Ask    string `json:"ask"`
  BidSize string `json:"bidSize"`
  AskSize string `json:"askSize"`
}
```
</details>

#### Get 24 hour ticker
```go
// options: market
response, err := bitvavo.Ticker24h(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, ticker := range response {
    fmt.Printf("%+v\n", ticker)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "market": "XTZ-EUR",
  "open": "1.21",
  "high": "1.2114",
  "low": "1.0974",
  "last": "1.1096",
  "volume": "41994.8008",
  "volumeQuote": "48041.67",
  "bid": "1.1088",
  "ask": "1.1155",
  "timestamp": 1565775776174,
  "bidSize": "175.05128",
  "askSize": "138.519066"
}
{
  "market": "XVG-EUR",
  "open": "0.0043222",
  "high": "0.0044139",
  "low": "0.0040849",
  "last": "0.0041952",
  "volume": "1237140.82971657",
  "volumeQuote": "5267.56",
  "bid": "0.0042134",
  "ask": "0.0043193",
  "timestamp": 1565775776103,
  "bidSize": "1698875.30729496",
  "askSize": "638047.77525823"
}
{
  "market": "ZIL-EUR",
  "open": "0.0081618",
  "high": "0.0082359",
  "low": "0.0076094",
  "last": "0.0077285",
  "volume": "774485.99486622",
  "volumeQuote": "6015.82",
  "bid": "0.0076778",
  "ask": "0.0077779",
  "timestamp": 1565775776160,
  "bidSize": "320186.06168593",
  "askSize": "158553.66311916"
}
...
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Ticker24h struct {
  Market      string `json:"market"`
  Open        string `json:"open"`
  High        string `json:"high"`
  Low         string `json:"low"`
  Last        string `json:"last"`
  Volume      string `json:"volume"`
  VolumeQuote string `json:"volumeQuote"`
  Bid         string `json:"bid"`
  Ask         string `json:"ask"`
  Timestamp   int    `json:"timestamp"`
  BidSize     string `json:"bidSize"`
  AskSize     string `json:"askSize"`
}
```
</details>

### Private

#### Place order
When placing an order, make sure that the correct optional parameters are set. For a limit order it is required to set both the amount and price. A market order is valid if either the amount or amountQuote has been set.

```go
// optional parameters: limit:(amount, price, postOnly), market:(amount, amountQuote, disableMarketProtection),
// both: timeInForce, selfTradePrevention, responseRequired
response, err := bitvavo.PlaceOrder("BTC-EUR", "buy", "market", map[string]string{"amount": "0.3"})
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "2261cb7b-346f-4d81-ae36-189dc44f8a87",
  "market": "BTC-EUR",
  "created": 1548681232827,
  "updated": 1548681232827,
  "status": "filled",
  "side": "sell",
  "orderType": "market",
  "amount": "0.3",
  "amountRemaining": "0",
  "price": "",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "0",
  "onHoldCurrency": "BTC",
  "filledAmount": "0.3",
  "filledAmountQuote": "898.107083735",
  "filledPrice": "",
  "feePaid": "2.247083735",
  "feeCurrency": "EUR",
  "fills": [
    {
      "id": "4ca1a839-2fe1-46d0-86dc-6582cf8d7456",
      "timestamp": 1548681232832,
      "amount": "0.17416747",
      "price": "2993.9",
      "taker": true,
      "fee": "1.309988433",
      "feeCurrency": "EUR",
      "settled": true
    },
    {
      "id": "99bc3431-604e-41ca-a4e6-610a7b0d1e7e",
      "timestamp": 1548681232840,
      "amount": "0.12583253",
      "price": "2993.4",
      "taker": true,
      "fee": "0.937095302",
      "feeCurrency": "EUR",
      "settled": true
    }
  ],
  "selfTradePrevention": "decrementAndCancel",
  "visible": false,
  "disableMarketProtection": false,
  "timeInForce": "",
  "postOnly": false
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Order struct {
  OrderId                 string `json:"orderId"`
  Market                  string `json:"market"`
  Created                 int    `json:"created"`
  Updated                 int    `json:"updated"`
  Status                  string `json:"status"`
  Side                    string `json:"side"`
  OrderType               string `json:"orderType"`
  Amount                  string `json:"amount"`
  AmountRemaining         string `json:"amountRemaining"`
  Price                   string `json:"price"`
  AmountQuote             string `json:"amountQuote"`
  AmountQuoteRemaining    string `json:"amountQuoteRemaining"`
  OnHold                  string `json:"onHold"`
  OnHoldCurrency          string `json:"onHoldCurrency"`
  FilledAmount            string `json:"filledAmount"`
  FilledAmountQuote       string `json:"filledAmountQuote"`
  FeePaid                 string `json:"feePaid"`
  FeeCurrency             string `json:"feeCurrency"`
  Fills                   []Fill `json:"fills"`
  SelfTradePrevention     string `json:"selfTradePrevention"`
  Visible                 bool   `json:"visible"`
  DisableMarketProtection bool   `json:"disableMarketProtection"`
  TimeInForce             string `json:"timeInForce"`
  PostOnly                bool   `json:"postOnly"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Update order
When updating an order make sure that at least one of the optional parameters has been set. Otherwise nothing can be updated.

```go
// Optional parameters: limit:(amount, amountRemaining, price, timeInForce, selfTradePrevention, postOnly)
// (set at least 1) (responseRequired can be set as well, but does not update anything)
response, err := bitvavo.UpdateOrder("BTC-EUR", "c2aa3b68-d72f-4a02-bb3d-30401f7d43ed",
                                      map[string]string{"price": "4000"})
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11",
  "market": "BTC-EUR",
  "created": 1548681391351,
  "updated": 1548681420120,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.4",
  "amountRemaining": "0.4",
  "price": "2000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "802",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Order struct {
  OrderId                 string `json:"orderId"`
  Market                  string `json:"market"`
  Created                 int    `json:"created"`
  Updated                 int    `json:"updated"`
  Status                  string `json:"status"`
  Side                    string `json:"side"`
  OrderType               string `json:"orderType"`
  Amount                  string `json:"amount"`
  AmountRemaining         string `json:"amountRemaining"`
  Price                   string `json:"price"`
  AmountQuote             string `json:"amountQuote"`
  AmountQuoteRemaining    string `json:"amountQuoteRemaining"`
  OnHold                  string `json:"onHold"`
  OnHoldCurrency          string `json:"onHoldCurrency"`
  FilledAmount            string `json:"filledAmount"`
  FilledAmountQuote       string `json:"filledAmountQuote"`
  FeePaid                 string `json:"feePaid"`
  FeeCurrency             string `json:"feeCurrency"`
  Fills                   []Fill `json:"fills"`
  SelfTradePrevention     string `json:"selfTradePrevention"`
  Visible                 bool   `json:"visible"`
  DisableMarketProtection bool   `json:"disableMarketProtection"`
  TimeInForce             string `json:"timeInForce"`
  PostOnly                bool   `json:"postOnly"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Get order
```go
response, err := bitvavo.GetOrder("BTC-EUR", "4729e0c3-fb21-41cf-957c-4c406ea78b11")
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11",
  "market": "BTC-EUR",
  "created": 1548681391351,
  "updated": 1548681420120,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.4",
  "amountRemaining": "0.4",
  "price": "2000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "802",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Order struct {
  OrderId                 string `json:"orderId"`
  Market                  string `json:"market"`
  Created                 int    `json:"created"`
  Updated                 int    `json:"updated"`
  Status                  string `json:"status"`
  Side                    string `json:"side"`
  OrderType               string `json:"orderType"`
  Amount                  string `json:"amount"`
  AmountRemaining         string `json:"amountRemaining"`
  Price                   string `json:"price"`
  AmountQuote             string `json:"amountQuote"`
  AmountQuoteRemaining    string `json:"amountQuoteRemaining"`
  OnHold                  string `json:"onHold"`
  OnHoldCurrency          string `json:"onHoldCurrency"`
  FilledAmount            string `json:"filledAmount"`
  FilledAmountQuote       string `json:"filledAmountQuote"`
  FeePaid                 string `json:"feePaid"`
  FeeCurrency             string `json:"feeCurrency"`
  Fills                   []Fill `json:"fills"`
  SelfTradePrevention     string `json:"selfTradePrevention"`
  Visible                 bool   `json:"visible"`
  DisableMarketProtection bool   `json:"disableMarketProtection"`
  TimeInForce             string `json:"timeInForce"`
  PostOnly                bool   `json:"postOnly"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Cancel order
```go
response, err := bitvavo.CancelOrder("BTC-EUR", "c2aa3b68-d72f-4a02-bb3d-30401f7d43ed")
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "c2aa3b68-d72f-4a02-bb3d-30401f7d43ed"
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type CancelOrder struct {
  OrderId string `json:"orderId"`
}
```
</details>

#### Get orders
Returns the same as get order, but can be used to return multiple orders at once.

```go
// options: limit, start, end, orderIdFrom, orderIdTo
response, err := bitvavo.GetOrders("BTC-EUR", map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, order := range response {
    fmt.Printf("%+v\n", order)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11",
  "market": "BTC-EUR",
  "created": 1548681391351,
  "updated": 1548681420120,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.4",
  "amountRemaining": "0.4",
  "price": "2000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "802",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
{
  "orderId": "2261cb7b-346f-4d81-ae36-189dc44f8a87",
  "market": "BTC-EUR",
  "created": 1548681232827,
  "updated": 1548681232827,
  "status": "filled",
  "side": "sell",
  "orderType": "market",
  "amount": "0.3",
  "amountRemaining": "0",
  "price": "",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "0",
  "onHoldCurrency": "BTC",
  "filledAmount": "0.3",
  "filledAmountQuote": "898.107083735",
  "filledPrice": "",
  "feePaid": "2.247083735",
  "feeCurrency": "EUR",
  "fills": [
    {
      "id": "4ca1a839-2fe1-46d0-86dc-6582cf8d7456",
      "timestamp": 1548681232832,
      "amount": "0.17416747",
      "price": "2993.9",
      "taker": true,
      "fee": "1.309988433",
      "feeCurrency": "EUR",
      "settled": true
    },
    {
      "id": "99bc3431-604e-41ca-a4e6-610a7b0d1e7e",
      "timestamp": 1548681232840,
      "amount": "0.12583253",
      "price": "2993.4",
      "taker": true,
      "fee": "0.937095302",
      "feeCurrency": "EUR",
      "settled": true
    }
  ],
  "selfTradePrevention": "decrementAndCancel",
  "visible": false,
  "disableMarketProtection": false,
  "timeInForce": "",
  "postOnly": false
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Order struct {
  OrderId                 string `json:"orderId"`
  Market                  string `json:"market"`
  Created                 int    `json:"created"`
  Updated                 int    `json:"updated"`
  Status                  string `json:"status"`
  Side                    string `json:"side"`
  OrderType               string `json:"orderType"`
  Amount                  string `json:"amount"`
  AmountRemaining         string `json:"amountRemaining"`
  Price                   string `json:"price"`
  AmountQuote             string `json:"amountQuote"`
  AmountQuoteRemaining    string `json:"amountQuoteRemaining"`
  OnHold                  string `json:"onHold"`
  OnHoldCurrency          string `json:"onHoldCurrency"`
  FilledAmount            string `json:"filledAmount"`
  FilledAmountQuote       string `json:"filledAmountQuote"`
  FeePaid                 string `json:"feePaid"`
  FeeCurrency             string `json:"feeCurrency"`
  Fills                   []Fill `json:"fills"`
  SelfTradePrevention     string `json:"selfTradePrevention"`
  Visible                 bool   `json:"visible"`
  DisableMarketProtection bool   `json:"disableMarketProtection"`
  TimeInForce             string `json:"timeInForce"`
  PostOnly                bool   `json:"postOnly"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Cancel orders
Cancels all orders in a market. If no market is specified, all orders of an account will be canceled.

```go
// options: market
response, err := bitvavo.CancelOrders(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, order := range response {
    fmt.Printf("%+v\n", order)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11"
}
{
  "orderId": "5850bca1-8bbb-470b-80a3-06ca7ec0e4c9"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type CancelOrder struct {
  OrderId string `json:"orderId"`
}
```
</details>

#### Get orders open
Returns all orders which are not filled or canceled.

```go
// options: market
response, err := bitvavo.OrdersOpen(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, order := range response {
    fmt.Printf("%+v\n", order)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11",
  "market": "BTC-EUR",
  "created": 1548681391351,
  "updated": 1548681420120,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.4",
  "amountRemaining": "0.4",
  "price": "2000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "802",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
{
  "orderId": "130e264c-dbb2-4535-9b5d-850498c2b923",
  "market": "BTC-EUR",
  "created": 1548681473272,
  "updated": 1548681484937,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.5",
  "amountRemaining": "0.5",
  "price": "3000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "1504",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Order struct {
  OrderId                 string `json:"orderId"`
  Market                  string `json:"market"`
  Created                 int    `json:"created"`
  Updated                 int    `json:"updated"`
  Status                  string `json:"status"`
  Side                    string `json:"side"`
  OrderType               string `json:"orderType"`
  Amount                  string `json:"amount"`
  AmountRemaining         string `json:"amountRemaining"`
  Price                   string `json:"price"`
  AmountQuote             string `json:"amountQuote"`
  AmountQuoteRemaining    string `json:"amountQuoteRemaining"`
  OnHold                  string `json:"onHold"`
  OnHoldCurrency          string `json:"onHoldCurrency"`
  FilledAmount            string `json:"filledAmount"`
  FilledAmountQuote       string `json:"filledAmountQuote"`
  FeePaid                 string `json:"feePaid"`
  FeeCurrency             string `json:"feeCurrency"`
  Fills                   []Fill `json:"fills"`
  SelfTradePrevention     string `json:"selfTradePrevention"`
  Visible                 bool   `json:"visible"`
  DisableMarketProtection bool   `json:"disableMarketProtection"`
  TimeInForce             string `json:"timeInForce"`
  PostOnly                bool   `json:"postOnly"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Get trades
Returns all trades within a market for this account.

```go
// options: limit, start, end, tradeIdFrom, tradeIdTo
response, err := bitvavo.Trades("BTC-EUR", map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, trade := range response {
    fmt.Printf("%+v\n", trade)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "timestamp": 1548681232840,
  "market": "BTC-EUR",
  "amount": "0.12583253",
  "side": "sell",
  "price": "2993.4",
  "taker": true,
  "fee": "0.937095302",
  "feeCurrency": "EUR",
  "settled": true
}
{
  "timestamp": 1548681232832,
  "market": "BTC-EUR",
  "amount": "0.17416747",
  "side": "sell",
  "price": "2993.9",
  "taker": true,
  "fee": "1.309988433",
  "feeCurrency": "EUR",
  "settled": true
}
{
  "timestamp": 1548679666586,
  "market": "BTC-EUR",
  "amount": "2.59562705",
  "side": "buy",
  "price": "2995.8",
  "taker": true,
  "fee": "19.43048361",
  "feeCurrency": "EUR",
  "settled": true
}
{
  "timestamp": 1548679666579,
  "market": "BTC-EUR",
  "amount": "1.42331069",
  "side": "buy",
  "price": "2994.8",
  "taker": true,
  "fee": "10.659145588",
  "feeCurrency": "EUR",
  "settled": true
}
{
  "timestamp": 1548679666574,
  "market": "BTC-EUR",
  "amount": "2.79772289",
  "side": "buy",
  "price": "2993.7",
  "taker": true,
  "fee": "20.936984207",
  "feeCurrency": "EUR",
  "settled": true
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Trades struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Market      string `json:"market"`
  Amount      string `json:"amount"`
  Side        string `json:"side"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Get balance
Returns the balance for this account.

```go
// options: symbol
response, err := bitvavo.Balance(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, balance := range response {
    fmt.Printf("%+v\n", balance)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "EUR",
  "available": "5003.5",
  "inOrder": "7963.43"
}
{
  "symbol": "BTC",
  "available": "0.02605972",
  "inOrder": "0.079398"
}
{
  "symbol": "ADA",
  "available": "3.8",
  "inOrder": "1"
}
{
  "symbol": "BCH",
  "available": "0.00952811",
  "inOrder": "0"
}
{
  "symbol": "BSV",
  "available": "0.00952811",
  "inOrder": "0"
}
{
  "symbol": "DASH",
  "available": "22.36624125",
  "inOrder": "0.63375875"
}
 ...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Balance struct {
  Symbol    string `json:"symbol"`
  Available string `json:"available"`
  InOrder   string `json:"inOrder"`
}
```
</details>

#### Deposit assets
Returns the address which can be used to deposit funds.

```go
response, err := bitvavo.DepositAssets("BTC")
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
}
```
<details>
 <summary>View Response</summary>

```go
{
  "address": "BitcoinAddress",
  "iban": "",
  "bic": "",
  "description": "",
  "paymentId": ""
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type DepositAssets struct {
  Address     string `json:"address"`
  Iban        string `json:"iban"`
  Bic         string `json:"bic"`
  Description string `json:"description"`
  PaymentId   string `json:"paymentId"`
}
```
</details>

#### Withdraw assets
Can be used to withdraw funds from Bitvavo.

```go
// optional parameters: paymentId, internal, addWithdrawalFee
response, err := bitvavo.WithdrawAssets("BTC", "1", "BitcoinAddress", map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "BTC",
  "amount": "1",
  "success": true
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type WithdrawAssets struct {
  Symbol  string `json:"symbol"`
  Amount  string `json:"amount"`
  Success bool   `json:"success"`
}
```
</details>

#### Get deposit history
Returns the deposit history of your account.

```go
// options: symbol, limit, start, end
response, err := bitvavo.DepositHistory(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, deposit := range response {
    fmt.Printf("%+v\n", deposit)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "EUR",
  "amount": "1",
  "address": "NL12RABO324234234",
  "paymentId": "",
  "fee": "0",
  "txId": "",
  "timestamp": 1521550025000,
  "status": "completed"
}
{
  "symbol": "BTC",
  "amount": "0.099",
  "address": "",
  "paymentId": "",
  "fee": "0",
  "txId": "0c6497e608212a516b8218674cb0ca04f65b67a00fe8bddaa1ecb03e9b029255",
  "timestamp": 1511873910000,
  "status": "completed"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type History struct {
  Symbol    string `json:"symbol"`
  Amount    string `json:"amount"`
  Address   string `json:"address"`
  PaymentId string `json:"paymentId"`
  Fee       string `json:"fee"`
  TxId      string `json:"txId"`
  Timestamp int    `json:"timestamp"`
  Status    string `json:"status"`
}
```
</details>

#### Get withdrawal history
Returns the withdrawal history of an account.

```go
// options: symbol, limit, start, end
response, err := bitvavo.WithdrawalHistory(map[string]string{})
if err != nil {
  fmt.Println(err)
} else {
  for _, withdrawal := range response {
    fmt.Printf("%+v\n", withdrawal)
  }
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "BTC",
  "amount": "0.99994",
  "address": "1CqtG5z55x7bYD5GxsAXPx59DEyujs4bjm",
  "paymentId": "",
  "fee": "0.00006",
  "txId": "",
  "timestamp": 1548682993000,
  "status": "awaiting_processing"
}
{
  "symbol": "BTC",
  "amount": "0.09994",
  "address": "1CqtG5z55x7bYD5GxsAXPx59DEyujs4bjm",
  "paymentId": "",
  "fee": "0.00006",
  "txId": "",
  "timestamp": 1548425559000,
  "status": "awaiting_processing"
}
{
  "symbol": "EUR",
  "amount": "50",
  "address": "NL123BIM",
  "paymentId": "",
  "fee": "0",
  "txId": "",
  "timestamp": 1548409721000,
  "status": "completed"
}
{
  "symbol": "BTC",
  "amount": "0.01939",
  "address": "3QpyxeA7yWWsSURXEmuBBzHpxjqn7Rbyme",
  "paymentId": "",
  "fee": "0.00002",
  "txId": "da2299c86fce67eb899aeaafbe1f81cf663a3850cf9f3337c92b2d87945532db",
  "timestamp": 1537803091000,
  "status": "completed"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type History struct {
  Symbol    string `json:"symbol"`
  Amount    string `json:"amount"`
  Address   string `json:"address"`
  PaymentId string `json:"paymentId"`
  Fee       string `json:"fee"`
  TxId      string `json:"txId"`
  Timestamp int    `json:"timestamp"`
  Status    string `json:"status"`
}
```
</details>

## Websockets

All requests which can be done through REST requests can also be performed over websockets. Bitvavo also provides six [subscriptions](https://github.com/bitvavo/go-bitvavo-api#subscriptions). If subscribed to these, updates specific for that type/market are pushed immediately.

### Getting started

The websocket object should be intialised through the `bitvavo.NewWebsocket()` function. This function returns the websocket object as the first parameter and the error channel as second parameter. All other functions return a single [channel](https://golang.org/doc/effective_go.html#channels) which will be used to return all responses over. The only exception is the [`websocket.SubscriptionAccount()`](https://github.com/bitvavo/go-bitvavo-api#account-subscription) function, which returns two separate channels. This is because the account channel might receive either an order event or a fill event. Each different response can be handled separately this way.
Best practice would be to listen to all channels in a single for loop and setup your response handling there.


```go
websocket, errChannel := bitvavo.NewWebsocket()

timeChannel := websocket.Time()

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle your errors here", errorMsg)
  case response := <-timeChannel:
    fmt.Println("Handle your response here", response)
  }
}
```

The api key and secret are copied from the bitvavo object. Therefore if you want to use the private portion of the websockets API, you should set both the key and secret as specified in [REST requests](https://github.com/bitvavo/go-bitvavo-api#rest-requests).

### Public

#### Get time
```go
channel := websocket.Time()

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "time": 1548679180309
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Time struct {
  Time int `json:"time"`
}
```
</details>

#### Get markets
```go
// options: market
channel := websocket.Markets(map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "status": "trading",
  "base": "ADA",
  "quote": "BTC",
  "market": "ADA-BTC",
  "pricePrecision": 5,
  "minOrderInQuoteAsset": "0.001",
  "minOrderInBaseAsset": "100",
  "orderTypes": [
    "market",
    "limit"
  ]
}
{
  "status": "trading",
  "base": "ADA",
  "quote": "EUR",
  "market": "ADA-EUR",
  "pricePrecision": 5,
  "minOrderInQuoteAsset": "10",
  "minOrderInBaseAsset": "100",
  "orderTypes": [
    "market",
    "limit"
  ]
}
{
  "status": "trading",
  "base": "AE",
  "quote": "BTC",
  "market": "AE-BTC",
  "pricePrecision": 5,
  "minOrderInQuoteAsset": "0.001",
  "minOrderInBaseAsset": "10",
  "orderTypes": [
    "market",
    "limit"
  ]
}
{
  "status": "trading",
  "base": "AE",
  "quote": "EUR",
  "market": "AE-EUR",
  "pricePrecision": 5,
  "minOrderInQuoteAsset": "10",
  "minOrderInBaseAsset": "10",
  "orderTypes": [
    "market",
    "limit"
  ]
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Markets struct {
  Status               string   `json:"status"`
  Base                 string   `json:"base"`
  Quote                string   `json:"quote"`
  Market               string   `json:"market"`
  PricePrecision       int      `json:"pricePrecision"`
  MinOrderInQuoteAsset string   `json:"minOrderInQuoteAsset"`
  MinOrderInBaseAsset  string   `json:"minOrderInBaseAsset"`
  OrderTypes           []string `json:"orderTypes"`
}
```
</details>

#### Get assets
```go
// options: symbol
channel := websocket.Assets(map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "ADA",
  "name": "Cardano",
  "decimals": 6,
  "depositFee": "0",
  "depositConfirmations": 20,
  "depositStatus": "OK",
  "withdrawalFee": "0.2",
  "withdrawalMinAmount": "0.2",
  "withdrawalStatus": "OK",
  "networks": [
    "Mainnet"
  ],
  "message": ""
}
{
  "symbol": "AE",
  "name": "Aeternity",
  "decimals": 8,
  "depositFee": "0",
  "depositConfirmations": 30,
  "depositStatus": "OK",
  "withdrawalFee": "2",
  "withdrawalMinAmount": "2",
  "withdrawalStatus": "OK",
  "networks": [
    "Mainnet"
  ],
  "message": ""
}
{
  "symbol": "AION",
  "name": "Aion",
  "decimals": 8,
  "depositFee": "0",
  "depositConfirmations": 0,
  "depositStatus": "",
  "withdrawalFee": "3",
  "withdrawalMinAmount": "3",
  "withdrawalStatus": "",
  "networks": [
    "Mainnet"
  ],
  "message": ""
}
{
  "symbol": "ANT",
  "name": "Aragon",
  "decimals": 8,
  "depositFee": "0",
  "depositConfirmations": 30,
  "depositStatus": "OK",
  "withdrawalFee": "2",
  "withdrawalMinAmount": "2",
  "withdrawalStatus": "OK",
  "networks": [
    "Mainnet"
  ],
  "message": ""
}
 ...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Assets struct {
  Symbol               string   `json:"symbol"`
  Name                 string   `json:"name"`
  Decimals             int      `json:"decimals"`
  DepositFee           string   `json:"depositFee"`
  DepositConfirmations int      `json:"depositConfirmations"`
  DepositStatus        string   `json:"depositStatus"`
  WithdrawalFee        string   `json:"withdrawalFee"`
  WithdrawalMinAmount  string   `json:"withdrawalMinAmount"`
  WithdrawalStatus     string   `json:"withdrawalStatus"`
  Networks             []string `json:"networks"`
  Message              string   `json:"message"`
}
```
</details>

#### Get book per market
```go
// options: depth
channel := websocket.book("BTC-EUR", map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "market": "BTC-EUR",
  "nonce": 7831,
  "bids": [
    [
      "2992.4",
      "0.11023153"
    ],
    [
      "2991.5",
      "0.50272746"
    ],
    [
      "2990.6",
      "1.27126107"
    ],
    [
      "2989.6",
      "3.0301821"
    ],
    [
      "2988.6",
      "3.2848159"
    ],
    ...
  ],
  "asks": [
    [
      "2993.3",
      "1.1852825"
    ],
    [
      "2994.1",
      "0.401334"
    ],
    [
      "2994.7",
      "0.31577418"
    ],
    [
      "2995.5",
      "2.30306344"
    ],
    [
      "2996.1",
      "2.63275436"
    ],
    ...
  ]
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Book struct {
  Market string     `json:"market"`
  Nonce  int        `json:"nonce"`
  Bids    [][]string `json:"bids"`
  Asks   [][]string `json:"asks"`
}
```
</details>

#### Get trades per market
```go
// options: limit, start, end, tradeIdFrom, tradeIdTo
channel := websocket.PublicTrades("BTC-EUR", map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "timestamp": 1548678622527,
  "id": "d4e1c700-8432-4ec7-b141-cb6bcee7a613",
  "amount": "2.36697512",
  "price": "3001.8",
  "side": "buy"
}
{
  "timestamp": 1548678622520,
  "id": "17c4ec15-7806-451e-8cc6-8ca48f1eef66",
  "amount": "3.37614478",
  "price": "3001.2",
  "side": "buy"
}
{
  "timestamp": 1548678622514,
  "id": "eb9d8230-6c55-424b-be93-2d41d39ab092",
  "amount": "3.27104588",
  "price": "3000.2",
  "side": "buy"
}
{
  "timestamp": 1548678622506,
  "id": "9ccc74d6-47d9-44f1-a51a-7dbbb3af5373",
  "amount": "0.42395607",
  "price": "2999.6",
  "side": "buy"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type PublicTrades struct {
  Timestamp int    `json:"timestamp"`
  Id        string `json:"id"`
  Amount    string `json:"amount"`
  Price     string `json:"price"`
  Side      string `json:"side"`
}
```
</details>

#### Get candles per market
```go
// options: limit, start, end
channel := websocket.candles("BTC-EUR", "1h", map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "Timestamp": 1548680400000,
  "Open": "2989",
  "High": "2991.1",
  "Low": "2989",
  "Close": "2989",
  "Volume": "0.9"
}
{
  "Timestamp": 1548676800000,
  "Open": "2999.3",
  "High": "3002.6",
  "Low": "2989.2",
  "Close": "2999.3",
  "Volume": "63.00046504"
}
{
  "Timestamp": 1548669600000,
  "Open": "3012.9",
  "High": "3015.8",
  "Low": "3000",
  "Close": "3012.9",
  "Volume": "8"
}
{
  "Timestamp": 1548417600000,
  "Open": "3124",
  "High": "3125.1",
  "Low": "3124",
  "Close": "3124",
  "Volume": "0.1"
}
{
  "Timestamp": 1548237600000,
  "Open": "3143",
  "High": "3143.3",
  "Low": "3141.1",
  "Close": "3143",
  "Volume": "60.67250851"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Candle struct {
  Timestamp int
  Open      string
  High      string
  Low       string
  Close     string
  Volume    string
}
```
</details>

#### Get price ticker
```go
// options: market
channel := websocket.TickerPrice(map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "market": "EOS-EUR",
  "price": "2.0142"
}
{
  "market": "XRP-EUR",
  "price": "0.25193"
}
{
  "market": "ETH-EUR",
  "price": "99.877"
}
{
  "market": "IOST-EUR",
  "price": "0.005941"
}
{
  "market": "BCH-EUR",
  "price": "106.57"
}
{
  "market": "BTC-EUR",
  "price": "2991.1"
}
{
  "market": "STORM-EUR",
  "price": "0.0025672"
}
{
  "market": "EOS-BTC",
  "price": "0.00066289"
}
{
  "market": "BSV-EUR",
  "price": "57.6"
}
{
  "market": "ETH-BTC",
  "price": "0.032373"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type TickerPrice struct {
  Market string `json:"market"`
  Price  string `json:"price"`
}
```
</details>

#### Get book ticker
```go
// options: market
channel := websocket.TickerBook(map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "market": "XVG-BTC",
  "bid": "0.00000045",
  "ask": "0.00000046",
  "bidSize": "28815.01275017",
  "askSize": "38392.85089495"
}
{
  "market": "XVG-EUR",
  "bid": "0.004213",
  "ask": "0.0043174",
  "bidSize": "1699002.27041019",
  "askSize": "638327.18139947"
}
{
  "market": "ZIL-BTC",
  "bid": "0.00000082",
  "ask": "0.00000083",
  "bidSize": "140980.13397262",
  "askSize": "98568.18059098"
}
{
  "market": "ZIL-EUR",
  "bid": "0.0076771",
  "ask": "0.0077744",
  "bidSize": "320216.82744213",
  "askSize": "157923.96870507"
}
{
  "market": "ZRX-BTC",
  "bid": "0.00001684",
  "ask": "0.000016898",
  "bidSize": "631.2417155",
  "askSize": "1551.90609202"
}
{
  "market": "ZRX-EUR",
  "bid": "0.15766",
  "ask": "0.15826",
  "bidSize": "873.64460869",
  "askSize": "1010.23609323"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type TickerBook struct {
  Market string `json:"market"`
  Bid    string `json:"bid"`
  Ask    string `json:"ask"`
  BidSize string `json:"bidSize"`
  AskSize string `json:"askSize"`
}
```
</details>


#### Get 24 hour ticker
```go
// options: market
channel := websocket.Ticker24h(map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "market": "XTZ-EUR",
  "open": "1.21",
  "high": "1.2114",
  "low": "1.0974",
  "last": "1.1096",
  "volume": "41994.8008",
  "volumeQuote": "48041.67",
  "bid": "1.1088",
  "ask": "1.1155",
  "timestamp": 1565775776174,
  "bidSize": "175.05128",
  "askSize": "138.519066"
}
{
  "market": "XVG-EUR",
  "open": "0.0043222",
  "high": "0.0044139",
  "low": "0.0040849",
  "last": "0.0041952",
  "volume": "1237140.82971657",
  "volumeQuote": "5267.56",
  "bid": "0.0042134",
  "ask": "0.0043193",
  "timestamp": 1565775776103,
  "bidSize": "1698875.30729496",
  "askSize": "638047.77525823"
}
{
  "market": "ZIL-EUR",
  "open": "0.0081618",
  "high": "0.0082359",
  "low": "0.0076094",
  "last": "0.0077285",
  "volume": "774485.99486622",
  "volumeQuote": "6015.82",
  "bid": "0.0076778",
  "ask": "0.0077779",
  "timestamp": 1565775776160,
  "bidSize": "320186.06168593",
  "askSize": "158553.66311916"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Ticker24h struct {
  Market      string `json:"market"`
  Open        string `json:"open"`
  High        string `json:"high"`
  Low         string `json:"low"`
  Last        string `json:"last"`
  Volume      string `json:"volume"`
  VolumeQuote string `json:"volumeQuote"`
  Bid         string `json:"bid"`
  Ask         string `json:"ask"`
  Timestamp   int    `json:"timestamp"`
  BidSize     string `json:"bidSize"`
  AskSize     string `json:"askSize"`
}
```
</details>

### Private

#### Place order
When placing an order, make sure that the correct optional parameters are set. For a limit order it is required to set both the amount and price. A market order is valid if either the amount or the amountQuote has been set.

```go
// optional parameters: limit:(amount, price, postOnly), market:(amount, amountQuote, disableMarketProtection),
// both: timeInForce, selfTradePrevention, responseRequired
channel := websocket.PlaceOrder("BTC-EUR", "sell", "market", map[string]string{"amount": "0.3"})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "2261cb7b-346f-4d81-ae36-189dc44f8a87",
  "market": "BTC-EUR",
  "created": 1548681232827,
  "updated": 1548681232827,
  "status": "filled",
  "side": "sell",
  "orderType": "market",
  "amount": "0.3",
  "amountRemaining": "0",
  "price": "",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "0",
  "onHoldCurrency": "BTC",
  "filledAmount": "0.3",
  "filledAmountQuote": "898.107083735",
  "filledPrice": "",
  "feePaid": "2.247083735",
  "feeCurrency": "EUR",
  "fills": [
    {
      "id": "4ca1a839-2fe1-46d0-86dc-6582cf8d7456",
      "timestamp": 1548681232832,
      "amount": "0.17416747",
      "price": "2993.9",
      "taker": true,
      "fee": "1.309988433",
      "feeCurrency": "EUR",
      "settled": true
    },
    {
      "id": "99bc3431-604e-41ca-a4e6-610a7b0d1e7e",
      "timestamp": 1548681232840,
      "amount": "0.12583253",
      "price": "2993.4",
      "taker": true,
      "fee": "0.937095302",
      "feeCurrency": "EUR",
      "settled": true
    }
  ],
  "selfTradePrevention": "decrementAndCancel",
  "visible": false,
  "disableMarketProtection": false,
  "timeInForce": "",
  "postOnly": false
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Order struct {
  OrderId                 string `json:"orderId"`
  Market                  string `json:"market"`
  Created                 int    `json:"created"`
  Updated                 int    `json:"updated"`
  Status                  string `json:"status"`
  Side                    string `json:"side"`
  OrderType               string `json:"orderType"`
  Amount                  string `json:"amount"`
  AmountRemaining         string `json:"amountRemaining"`
  Price                   string `json:"price"`
  AmountQuote             string `json:"amountQuote"`
  AmountQuoteRemaining    string `json:"amountQuoteRemaining"`
  OnHold                  string `json:"onHold"`
  OnHoldCurrency          string `json:"onHoldCurrency"`
  FilledAmount            string `json:"filledAmount"`
  FilledAmountQuote       string `json:"filledAmountQuote"`
  FeePaid                 string `json:"feePaid"`
  FeeCurrency             string `json:"feeCurrency"`
  Fills                   []Fill `json:"fills"`
  SelfTradePrevention     string `json:"selfTradePrevention"`
  Visible                 bool   `json:"visible"`
  DisableMarketProtection bool   `json:"disableMarketProtection"`
  TimeInForce             string `json:"timeInForce"`
  PostOnly                bool   `json:"postOnly"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Update order
When updating an order make sure that at least one of the optional parameters has been set. Otherwise nothing can be updated.

```go
// Optional parameters: limit:(amount, amountRemaining, price, timeInForce, selfTradePrevention, postOnly)
// (set at least 1) (responseRequired can be set as well, but does not update anything)
channel := websocket.UpdateOrder("BTC-EUR", "4729e0c3-fb21-41cf-957c-4c406ea78b11", 
                                  map[string]string{"amount":"0.4"})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11",
  "market": "BTC-EUR",
  "created": 1548681391351,
  "updated": 1548681420120,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.4",
  "amountRemaining": "0.4",
  "price": "2000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "802",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type UpdateOrder struct {
  OrderId              string `json:"orderId"`
  Market               string `json:"market"`
  Created              int    `json:"created"`
  Updated              int    `json:"updated"`
  Status               string `json:"status"`
  Side                 string `json:"side"`
  OrderType            string `json:"orderType"`
  Amount               string `json:"amount"`
  AmountRemaining      string `json:"amountRemaining"`
  Price                string `json:"price"`
  AmountQuote          string `json:"amountQuote"`
  AmountQuoteRemaining string `json:"amountQuoteRemaining"`
  OnHold               string `json:"onHold"`
  OnHoldCurrency       string `json:"onHoldCurrency"`
  FilledAmount         string `json:"filledAmount"`
  FilledAmountQuote    string `json:"filledAmountQuote"`
  FeePaid              string `json:"feePaid"`
  FeeCurrency          string `json:"feeCurrency"`
  Fills                []Fill `json:"fills"`
  TimeInForce          string `json:"timeInForce"`
  PostOnly             bool   `json:"postOnly"`
  SelfTradePrevention  string `json:"selfTradePrevention"`
  Visible              bool   `json:"visible"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Get order
```go
channel := websocket.GetOrder("BTC-EUR", "c5f419b3-65c8-4a4a-8f5f-4dd7c24c5172")

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11",
  "market": "BTC-EUR",
  "created": 1548681391351,
  "updated": 1548681420120,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.4",
  "amountRemaining": "0.4",
  "price": "2000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "802",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Order struct {
  OrderId                 string `json:"orderId"`
  Market                  string `json:"market"`
  Created                 int    `json:"created"`
  Updated                 int    `json:"updated"`
  Status                  string `json:"status"`
  Side                    string `json:"side"`
  OrderType               string `json:"orderType"`
  Amount                  string `json:"amount"`
  AmountRemaining         string `json:"amountRemaining"`
  Price                   string `json:"price"`
  AmountQuote             string `json:"amountQuote"`
  AmountQuoteRemaining    string `json:"amountQuoteRemaining"`
  OnHold                  string `json:"onHold"`
  OnHoldCurrency          string `json:"onHoldCurrency"`
  FilledAmount            string `json:"filledAmount"`
  FilledAmountQuote       string `json:"filledAmountQuote"`
  FeePaid                 string `json:"feePaid"`
  FeeCurrency             string `json:"feeCurrency"`
  Fills                   []Fill `json:"fills"`
  SelfTradePrevention     string `json:"selfTradePrevention"`
  Visible                 bool   `json:"visible"`
  DisableMarketProtection bool   `json:"disableMarketProtection"`
  TimeInForce             string `json:"timeInForce"`
  PostOnly                bool   `json:"postOnly"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Cancel order
```go
channel := websocket.CancelOrder("BTC-EUR", "c2aa3b68-d72f-4a02-bb3d-30401f7d43ed")

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "c2aa3b68-d72f-4a02-bb3d-30401f7d43ed"
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type CancelOrder struct {
  OrderId string `json:"orderId"`
}
```
</details>

#### Get orders
Returns the same as get order, but can be used to return multiple orders at once.

```go
// options: limit, start, end, orderIdFrom, orderIdTo
channel := websocket.GetOrders("BTC-EUR", map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11",
  "market": "BTC-EUR",
  "created": 1548681391351,
  "updated": 1548681420120,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.4",
  "amountRemaining": "0.4",
  "price": "2000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "802",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
{
  "orderId": "2261cb7b-346f-4d81-ae36-189dc44f8a87",
  "market": "BTC-EUR",
  "created": 1548681232827,
  "updated": 1548681232827,
  "status": "filled",
  "side": "sell",
  "orderType": "market",
  "amount": "0.3",
  "amountRemaining": "0",
  "price": "",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "0",
  "onHoldCurrency": "BTC",
  "filledAmount": "0.3",
  "filledAmountQuote": "898.107083735",
  "filledPrice": "",
  "feePaid": "2.247083735",
  "feeCurrency": "EUR",
  "fills": [
    {
      "id": "4ca1a839-2fe1-46d0-86dc-6582cf8d7456",
      "timestamp": 1548681232832,
      "amount": "0.17416747",
      "price": "2993.9",
      "taker": true,
      "fee": "1.309988433",
      "feeCurrency": "EUR",
      "settled": true
    },
    {
      "id": "99bc3431-604e-41ca-a4e6-610a7b0d1e7e",
      "timestamp": 1548681232840,
      "amount": "0.12583253",
      "price": "2993.4",
      "taker": true,
      "fee": "0.937095302",
      "feeCurrency": "EUR",
      "settled": true
    }
  ],
  "selfTradePrevention": "decrementAndCancel",
  "visible": false,
  "disableMarketProtection": false,
  "timeInForce": "",
  "postOnly": false
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Order struct {
  OrderId                 string `json:"orderId"`
  Market                  string `json:"market"`
  Created                 int    `json:"created"`
  Updated                 int    `json:"updated"`
  Status                  string `json:"status"`
  Side                    string `json:"side"`
  OrderType               string `json:"orderType"`
  Amount                  string `json:"amount"`
  AmountRemaining         string `json:"amountRemaining"`
  Price                   string `json:"price"`
  AmountQuote             string `json:"amountQuote"`
  AmountQuoteRemaining    string `json:"amountQuoteRemaining"`
  OnHold                  string `json:"onHold"`
  OnHoldCurrency          string `json:"onHoldCurrency"`
  FilledAmount            string `json:"filledAmount"`
  FilledAmountQuote       string `json:"filledAmountQuote"`
  FeePaid                 string `json:"feePaid"`
  FeeCurrency             string `json:"feeCurrency"`
  Fills                   []Fill `json:"fills"`
  SelfTradePrevention     string `json:"selfTradePrevention"`
  Visible                 bool   `json:"visible"`
  DisableMarketProtection bool   `json:"disableMarketProtection"`
  TimeInForce             string `json:"timeInForce"`
  PostOnly                bool   `json:"postOnly"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Cancel orders
Cancels all orders in a market. If no market is specified, all orders of an account will be canceled.

```go
channel := websocket.CancelOrders(map[string]string{"market": "BTC-EUR"})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11"
}
{
  "orderId": "5850bca1-8bbb-470b-80a3-06ca7ec0e4c9"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type CancelOrder struct {
  OrderId string `json:"orderId"`
}
```
</details>

#### Get orders open
Returns all orders which are not filled or canceled.

```go
// options: market
channel := websocket.OrdersOpen(map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "orderId": "4729e0c3-fb21-41cf-957c-4c406ea78b11",
  "market": "BTC-EUR",
  "created": 1548681391351,
  "updated": 1548681420120,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.4",
  "amountRemaining": "0.4",
  "price": "2000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "802",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
{
  "orderId": "130e264c-dbb2-4535-9b5d-850498c2b923",
  "market": "BTC-EUR",
  "created": 1548681473272,
  "updated": 1548681484937,
  "status": "new",
  "side": "buy",
  "orderType": "limit",
  "amount": "0.5",
  "amountRemaining": "0.5",
  "price": "3000",
  "amountQuote": "",
  "amountQuoteRemaining": "",
  "onHold": "1504",
  "onHoldCurrency": "EUR",
  "filledAmount": "0",
  "filledAmountQuote": "0",
  "filledPrice": "",
  "feePaid": "0",
  "feeCurrency": "EUR",
  "fills": [],
  "selfTradePrevention": "decrementAndCancel",
  "visible": true,
  "disableMarketProtection": false,
  "timeInForce": "GTC",
  "postOnly": false
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Order struct {
  OrderId                 string `json:"orderId"`
  Market                  string `json:"market"`
  Created                 int    `json:"created"`
  Updated                 int    `json:"updated"`
  Status                  string `json:"status"`
  Side                    string `json:"side"`
  OrderType               string `json:"orderType"`
  Amount                  string `json:"amount"`
  AmountRemaining         string `json:"amountRemaining"`
  Price                   string `json:"price"`
  AmountQuote             string `json:"amountQuote"`
  AmountQuoteRemaining    string `json:"amountQuoteRemaining"`
  OnHold                  string `json:"onHold"`
  OnHoldCurrency          string `json:"onHoldCurrency"`
  FilledAmount            string `json:"filledAmount"`
  FilledAmountQuote       string `json:"filledAmountQuote"`
  FeePaid                 string `json:"feePaid"`
  FeeCurrency             string `json:"feeCurrency"`
  Fills                   []Fill `json:"fills"`
  SelfTradePrevention     string `json:"selfTradePrevention"`
  Visible                 bool   `json:"visible"`
  DisableMarketProtection bool   `json:"disableMarketProtection"`
  TimeInForce             string `json:"timeInForce"`
  PostOnly                bool   `json:"postOnly"`
}

type Fill struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>


#### Get trades
Returns all trades within a market for this account.

```go
// options: limit, start, end, tradeIdFrom, tradeIdTo
channel := websocket.Trades("BTC-EUR", map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "timestamp": 1548681232840,
  "market": "BTC-EUR",
  "amount": "0.12583253",
  "side": "sell",
  "price": "2993.4",
  "taker": true,
  "fee": "0.937095302",
  "feeCurrency": "EUR",
  "settled": true
}
{
  "timestamp": 1548681232832,
  "market": "BTC-EUR",
  "amount": "0.17416747",
  "side": "sell",
  "price": "2993.9",
  "taker": true,
  "fee": "1.309988433",
  "feeCurrency": "EUR",
  "settled": true
}
{
  "timestamp": 1548679666586,
  "market": "BTC-EUR",
  "amount": "2.59562705",
  "side": "buy",
  "price": "2995.8",
  "taker": true,
  "fee": "19.43048361",
  "feeCurrency": "EUR",
  "settled": true
}
{
  "timestamp": 1548679666579,
  "market": "BTC-EUR",
  "amount": "1.42331069",
  "side": "buy",
  "price": "2994.8",
  "taker": true,
  "fee": "10.659145588",
  "feeCurrency": "EUR",
  "settled": true
}
{
  "timestamp": 1548679666574,
  "market": "BTC-EUR",
  "amount": "2.79772289",
  "side": "buy",
  "price": "2993.7",
  "taker": true,
  "fee": "20.936984207",
  "feeCurrency": "EUR",
  "settled": true
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Trades struct {
  Id          string `json:"id"`
  Timestamp   int    `json:"timestamp"`
  Market      string `json:"market"`
  Amount      string `json:"amount"`
  Side        string `json:"side"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
  Settled     bool   `json:"settled"`
}
```
</details>

#### Get balance
Returns the balance for this account.

```go
// options: symbol
channel := websocket.Balance(map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "EUR",
  "available": "5003.5",
  "inOrder": "7963.43"
}
{
  "symbol": "BTC",
  "available": "0.02605972",
  "inOrder": "0.079398"
}
{
  "symbol": "ADA",
  "available": "3.8",
  "inOrder": "1"
}
{
  "symbol": "BCH",
  "available": "0.00952811",
  "inOrder": "0"
}
{
  "symbol": "BSV",
  "available": "0.00952811",
  "inOrder": "0"
}
{
  "symbol": "DASH",
  "available": "22.36624125",
  "inOrder": "0.63375875"
}
 ...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type Balance struct {
  Symbol    string `json:"symbol"`
  Available string `json:"available"`
  InOrder   string `json:"inOrder"`
}
```
</details>

#### Deposit assets
Returns the address which can be used to deposit funds.

```go
channel := websocket.DepositAssets("BTC")

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "address": "BitcoinAddress",
  "iban": "",
  "bic": "",
  "description": "",
  "paymentId": ""
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type DepositAssets struct {
  Address     string `json:"address"`
  Iban        string `json:"iban"`
  Bic         string `json:"bic"`
  Description string `json:"description"`
  PaymentId   string `json:"paymentId"`
}
```
</details>

#### Withdraw Assets
Can be used to withdraw funds from Bitvavo.

```go
// optional parameters: paymentId, internal, addWithdrawalFee
channel := websocket.WithdrawAssets("BTC", "1", "BitcoinAddress", map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "BTC",
  "amount": "1",
  "success": true
}
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type WithdrawAssets struct {
  Symbol  string `json:"symbol"`
  Amount  string `json:"amount"`
  Success bool   `json:"success"`
}
```
</details>

#### Get deposit history
Returns the deposit history of your account.

```go
// options: symbol, limit, start, end
channel := websocket.DepositHistory(map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "EUR",
  "amount": "1",
  "address": "NL12RABO324234234",
  "paymentId": "",
  "fee": "0",
  "txId": "",
  "timestamp": 1521550025000,
  "status": "completed"
}
{
  "symbol": "BTC",
  "amount": "0.099",
  "address": "",
  "paymentId": "",
  "fee": "0",
  "txId": "0c6497e608212a516b8218674cb0ca04f65b67a00fe8bddaa1ecb03e9b029255",
  "timestamp": 1511873910000,
  "status": "completed"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type History struct {
  Symbol    string `json:"symbol"`
  Amount    string `json:"amount"`
  Address   string `json:"address"`
  PaymentId string `json:"paymentId"`
  Fee       string `json:"fee"`
  TxId      string `json:"txId"`
  Timestamp int    `json:"timestamp"`
  Status    string `json:"status"`
}
```
</details>

#### Get withdrawal history
Returns the withdrawal history of an account.

```go
// options: symbol, limit, start, end
channel := websocket.WithdrawalHistory(map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    for _, entry := range result {
      fmt.Printf("%+v\n", entry)
    }
  } 
}
```
<details>
 <summary>View Response</summary>

```go
{
  "symbol": "BTC",
  "amount": "0.99994",
  "address": "1CqtG5z55x7bYD5GxsAXPx59DEyujs4bjm",
  "paymentId": "",
  "fee": "0.00006",
  "txId": "",
  "timestamp": 1548682993000,
  "status": "awaiting_processing"
}
{
  "symbol": "BTC",
  "amount": "0.09994",
  "address": "1CqtG5z55x7bYD5GxsAXPx59DEyujs4bjm",
  "paymentId": "",
  "fee": "0.00006",
  "txId": "",
  "timestamp": 1548425559000,
  "status": "awaiting_processing"
}
{
  "symbol": "EUR",
  "amount": "50",
  "address": "NL123BIM",
  "paymentId": "",
  "fee": "0",
  "txId": "",
  "timestamp": 1548409721000,
  "status": "completed"
}
{
  "symbol": "BTC",
  "amount": "0.01939",
  "address": "3QpyxeA7yWWsSURXEmuBBzHpxjqn7Rbyme",
  "paymentId": "",
  "fee": "0.00002",
  "txId": "da2299c86fce67eb899aeaafbe1f81cf663a3850cf9f3337c92b2d87945532db",
  "timestamp": 1537803091000,
  "status": "completed"
}
...
```
</details>

<details>
 <summary>Struct Definition</summary>

```go
type History struct {
  Symbol    string `json:"symbol"`
  Amount    string `json:"amount"`
  Address   string `json:"address"`
  PaymentId string `json:"paymentId"`
  Fee       string `json:"fee"`
  TxId      string `json:"txId"`
  Timestamp int    `json:"timestamp"`
  Status    string `json:"status"`
}
```
</details>

### Subscriptions

#### Ticker subscription
Sends an update every time the best bid, best ask or last price changed.

```go
channel := websocket.SubscriptionTicker("BTC-EUR")

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```

<details>
 <summary>View Response</summary>

```go
{
  "event": "ticker",
  "market": "BTC-EUR",
  "bestBid": "9370.4",
  "bestBidSize": "0.3080136",
  "bestAsk": "9369.3",
  "bestAskSize": "0.10681936",
  "lastPrice": "9369.3"
}
```
</details>
<details>
 <summary>Struct Definition</summary>

```go
type SubscriptionTicker struct {
  Event       string `json:"event"`
  Market      string `json:"market"`
  BestBid     string `json:"bestBid"`
  BestBidSize string `json:"bestBidSize"`
  BestAsk     string `json:"bestAsk"`
  BestAskSize string `json:"bestAskSize"`
  LastPrice   string `json:"lastPrice"`
}
```
</details>

#### Ticker 24 hour subscription
Updated ticker24h objects are sent on this channel once per second. A ticker24h object is considered updated if one of the values besides timestamp has changed.

```go
channel := websocket.SubscriptionTicker24h("BTC-EUR")

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```

<details>
 <summary>View Response</summary>

```go
{
  "market": "BTC-EUR",
  "open": "10065",
  "high": "10082",
  "low": "9265.4",
  "last": "9371.8",
  "volume": "308.06115709",
  "volumeQuote": "2983934.09",
  "bid": "9365.3",
  "ask": "9369",
  "timestamp": 1565776275740,
  "bidSize": "0.10615797",
  "askSize": "0.10895214"
}
```
</details>
<details>
 <summary>Struct Definition</summary>

```go
type SubscriptionTicker24h struct {
  Event string      `json:"event"`
  Data  []Ticker24h `json:"data"`
}

type Ticker24h struct {
  Market      string `json:"market"`
  Open        string `json:"open"`
  High        string `json:"high"`
  Low         string `json:"low"`
  Last        string `json:"last"`
  Volume      string `json:"volume"`
  VolumeQuote string `json:"volumeQuote"`
  Bid         string `json:"bid"`
  Ask         string `json:"ask"`
  Timestamp   int    `json:"timestamp"`
  BidSize     string `json:"bidSize"`
  AskSize     string `json:"askSize"`
}
```
</details>

#### Account subscription
Sends an update whenever an event happens which is related to the account. These are ‘order’ events (create, update, cancel) or ‘fill’ events (a trade occurred).

```go
orderChannel, fillChannel := websocket.SubscriptionAccount("BTC-EUR")

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-orderChannel:
    fmt.Printf("%+v\n", result)
  case result := <-fillChannel:
    fmt.Printf("%+v\n", result)
  } 
}
```

<details>
 <summary>View Response</summary>

```go
Order:
{
  "event": "order",
  "orderId": "80b5f04d-21fc-4ebe-9c5f-6d34f78ee477",
  "market": "BTC-EUR",
  "created": 1548684420771,
  "updated": 1548684420771,
  "status": "filled",
  "side": "buy",
  "orderType": "limit",
  "amount": "1",
  "amountRemaining": "0",
  "price": "3000",
  "onHold": "0.83",
  "onHoldCurrency": "EUR",
  "timeInForce": "GTC",
  "postOnly": false,
  "selfTradePrevention": "decrementAndCancel",
  "visible": true
}

Fill:
{
  "event": "fill",
  "timestamp": 1548684420790,
  "market": "BTC-EUR",
  "orderId": "80b5f04d-21fc-4ebe-9c5f-6d34f78ee477",
  "fillId": "64cc0e3d-6e7b-451c-9034-9a6dc6c4665a",
  "amount": "0.17228569",
  "price": "2995.3",
  "taker": true,
  "fee": "1.282672743",
  "feeCurrency": "EUR"
}
```
</details>
<details>
 <summary>Struct Definition</summary>

```go
type SubscriptionAccountOrder struct {
  Event                string `json:"event"`
  OrderId              string `json:"orderId"`
  Market               string `json:"market"`
  Created              int    `json:"created"`
  Updated              int    `json:"updated"`
  Status               string `json:"status"`
  Side                 string `json:"side"`
  OrderType            string `json:"orderType"`
  Amount               string `json:"amount"`
  AmountRemaining      string `json:"amountRemaining"`
  AmountQuote          string `json:"amountQuote"`
  AmountQuoteRemaining string `json:"amountQuoteRemaining"`
  Price                string `json:"price"`
  OnHold               string `json:"onHold"`
  OnHoldCurrency       string `json:"onHoldCurrency"`
  TimeInForce          string `json:"timeInForce"`
  PostOnly             bool   `json:"postOnly"`
  SelfTradePrevention  string `json:"selfTradePrevention"`
  Visible              bool   `json:"visible"`
}

type SubscriptionAccountFill struct {
  Event       string `json:"event"`
  Timestamp   int    `json:"timestamp"`
  Market      string `json:"market"`
  OrderId     string `json:"orderId"`
  FillId      string `json:"fillId"`
  Amount      string `json:"amount"`
  Price       string `json:"price"`
  Taker       bool   `json:"taker"`
  Fee         string `json:"fee"`
  FeeCurrency string `json:"feeCurrency"`
}
```
</details>

#### Candles subscription
Sends an updated candle after each trade for the specified interval and market.

```go
channel := websocket.SubscriptionCandles("BTC-EUR", "1h")

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```

<details>
 <summary>View Response</summary>

```go
{
  "event": "candle",
  "market": "BTC-EUR",
  "interval": "1h",
  "candle": [
    {
      "Timestamp": 1548684000000,
      "Open": "2993.7",
      "High": "2996.9",
      "Low": "2992.5",
      "Close": "2993.7",
      "Volume": "8"
    }
  ]
}
```
</details>
<details>
 <summary>Struct Definition</summary>

```go
type SubscriptionCandles struct {
  Event    string   `json:"event"`
  Market   string   `json:"market"`
  Interval string   `json:"interval"`
  Candle   []Candle `json:"candle"`
}

type Candle struct {
  Timestamp int
  Open      string
  High      string
  Low       string
  Close     string
  Volume    string
}
```
</details>

#### Trades subscription
Sends an update whenever a trade has happened on this market. For your own trades, please subscribe to account.

```go
channel := websocket.SubscriptionTrades("BTC-EUR")

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```

<details>
 <summary>View Response</summary>

```go
{
  "event": "trade",
  "timestamp": 1548685870299,
  "market": "BTC-EUR",
  "id": "616bfa4e-b3ff-4b3f-a394-1538a49eb9bc",
  "amount": "1",
  "price": "2996",
  "side": "buy"
}
```
</details>
<details>
 <summary>Struct Definition</summary>

```go
type SubscriptionTrades struct {
  Event     string `json:"event"`
  Timestamp int    `json:"timestamp"`
  Market    string `json:"market"`
  Id        string `json:"id"`
  Amount    string `json:"amount"`
  Price     string `json:"price"`
  Side      string `json:"side"`
}
```
</details>

#### Book subscription
Sends an update whenever the order book for this specific market has changed. A list of tuples ([price, amount]) are returned, where amount ‘0’ means that there are no more orders at this price. If you wish to maintain your own copy of the order book, consider using the book subscription with local copy function.

```go
channel := websocket.SubscriptionBookUpdate("BTC-EUR")

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```

<details>
 <summary>View Response</summary>

```go
{
  "event": "book",
  "market": "BTC-EUR",
  "nonce": 484,
  "bids": [],
  "asks": [
    [
      "3010.5",
      "0"
    ],
    [
      "3010.2",
      "2.5018209"
    ]
  ]
}
```
</details>
<details>
 <summary>Struct Definition</summary>

```go
type SubscriptionBookUpdate struct {
  Event  string     `json:"event"`
  Market string     `json:"market"`
  Nonce  int        `json:"nonce"`
  Bids   [][]string `json:"bids"`
  Asks   [][]string `json:"asks"`
}
```
</details>

#### Book subscription with local copy
This is a combination of get book per market and the book subscription which maintains a local copy. On every update to the order book, the entire order book is returned to the callback, while the book subscription will only return updates to the book.

```go
channel := websocket.SubscriptionBook("BTC-EUR", map[string]string{})

for {
  select {
  case errorMsg := <-errChannel:
    fmt.Println("Handle error", errorMsg)
  case result := <-channel:
    fmt.Printf("%+v\n", result)
  } 
}
```

<details>
 <summary>View Response</summary>

```go
{
  "market": "BTC-EUR",
  "nonce": 484,
  "bids": [
    [
      "2995.2",
      "0.61519665"
    ],
    [
      "2994.7",
      "0.21423785"
    ],
    [
      "2994.1",
      "0.25726135"
    ],
    [
      "2993.6",
      "0.42691453"
    ],
    [
      "2993.1",
      "2.15710952"
    ],
    ...
  ],
  "asks": [
    [
      "2996.1",
      "1.16874422"
    ],
    [
      "2996.6",
      "0.7197304"
    ],
    [
      "2997.1",
      "1.64062477"
    ],
    [
      "2998",
      "2.15826937"
    ],
    [
      "2999.8",
      "3.51648046"
    ],
    ...
  ]
}
```
</details>
<details>
 <summary>Struct Definition</summary>

```go
type LocalBook struct {
  Book map[string]Book `json:"book"`
}

type Book struct {
  Market string     `json:"market"`
  Nonce  int        `json:"nonce"`
  Bids   [][]string `json:"bids"`
  Asks   [][]string `json:"asks"`
}
```
</details>