package bitvavo

import (
  "bytes"
  "crypto/hmac"
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "fmt"
  "github.com/gorilla/websocket"
  "io/ioutil"
  "net/http"
  "net/url"
  "reflect"
  "strconv"
  "strings"
  "sync"
  "time"
)

var rateLimitRemaining = 1000
var rateLimitReset = 0

type TimeResponse struct {
  Action   string `json:"action"`
  Response Time   `json:"response"`
}

type Time struct {
  Time int `json:"time"`
}

type MarketsResponse struct {
  Action   string    `json:"action"`
  Response []Markets `json:"response"`
}

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

type AssetsResponse struct {
  Action   string   `json:"action"`
  Response []Assets `json:"response"`
}

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

type BookResponse struct {
  Action   string `json:"action"`
  Response Book   `json:"response"`
}

type Book struct {
  Market string     `json:"market"`
  Nonce  int        `json:"nonce"`
  Bids   [][]string `json:"bids"`
  Asks   [][]string `json:"asks"`
}

type PublicTradesResponse struct {
  Action   string         `json:"action"`
  Response []PublicTrades `json:"response"`
}

type PublicTrades struct {
  Timestamp int    `json:"timestamp"`
  Id        string `json:"id"`
  Amount    string `json:"amount"`
  Price     string `json:"price"`
  Side      string `json:"side"`
}

type CandlesResponse struct {
  Action   string        `json:"action"`
  Response []interface{} `json:"response"`
}

type Candles struct {
  Candles []Candle `json:"candles"`
}

type Candle struct {
  Timestamp int
  Open      string
  High      string
  Low       string
  Close     string
  Volume    string
}

type Ticker24hResponse struct {
  Action   string      `json:"action"`
  Response []Ticker24h `json:"response"`
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

type TickerPriceResponse struct {
  Action   string        `json:"action"`
  Response []TickerPrice `json:"response"`
}

type TickerPrice struct {
  Market string `json:"market"`
  Price  string `json:"price"`
}

type TickerBookResponse struct {
  Action   string       `json:"action"`
  Response []TickerBook `json:"response"`
}

type TickerBook struct {
  Market string `json:"market"`
  Bid    string `json:"bid"`
  Ask    string `json:"ask"`
  BidSize string `json:"bidSize"`
  AskSize string `json:"askSize"`
}

type PlaceOrderResponse struct {
  Action   string `json:"action"`
  Response Order  `json:"response"`
}

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
  TriggerAmount           string `json:"triggerAmount"`
  TriggerPrice            string `json:"triggerPrice"`
  TriggerType             string `json:"triggerType"`
  TriggerReference        string `json:"triggerReference"`
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

type GetOrderResponse struct {
  Action   string `json:"action"`
  Response Order  `json:"response"`
}

type UpdateOrderResponse struct {
  Action   string `json:"action"`
  Response Order  `json:"response"`
}

type CancelOrderResponse struct {
  Action   string      `json:"action"`
  Response CancelOrder `json:"response"`
}

type CancelOrder struct {
  OrderId string `json:"orderId"`
}

type GetOrdersResponse struct {
  Action   string  `json:"action"`
  Response []Order `json:"response"`
}

type CancelOrdersResponse struct {
  Action   string        `json:"action"`
  Response []CancelOrder `json:"response"`
}

type OrdersOpenResponse struct {
  Action   string  `json:"action"`
  Response []Order `json:"response"`
}

type TradesResponse struct {
  Action   string   `json:"action"`
  Response []Trades `json:"response"`
}

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

type AccountResponse struct {
  Action   string  `json:"action"`
  Response Account `json:"response"`
}

type Account struct {
  Fees  FeeObject `json:"fees"`
}

type FeeObject struct {
  Taker  string `json:"taker"`
  Maker  string `json:"maker"`
  Volume string `json:"volume"`
}

type BalanceResponse struct {
  Action   string    `json:"action"`
  Response []Balance `json:"response"`
}

type Balance struct {
  Symbol    string `json:"symbol"`
  Available string `json:"available"`
  InOrder   string `json:"inOrder"`
}

type DepositAssetsResponse struct {
  Action   string        `json:"action"`
  Response DepositAssets `json:"response"`
}

type DepositAssets struct {
  Address     string `json:"address"`
  Iban        string `json:"iban"`
  Bic         string `json:"bic"`
  Description string `json:"description"`
  PaymentId   string `json:"paymentId"`
}

type WithdrawAssetsResponse struct {
  Action   string         `json:"action"`
  Response WithdrawAssets `json:"response"`
}

type WithdrawAssets struct {
  Symbol  string `json:"symbol"`
  Amount  string `json:"amount"`
  Success bool   `json:"success"`
}

type HistoryResponse struct {
  Action   string    `json:"action"`
  Response []History `json:"response"`
}

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

type SubscriptionTickerResponse struct {
  Action   string             `json:"action"`
  Response SubscriptionTicker `json:"response"`
}

type SubscriptionTicker struct {
  Event       string `json:"event"`
  Market      string `json:"market"`
  BestBid     string `json:"bestBid"`
  BestBidSize string `json:"bestBidSize"`
  BestAsk     string `json:"bestAsk"`
  BestAskSize string `json:"bestAskSize"`
  LastPrice   string `json:"lastPrice"`
}

type SubscriptionTicker24h struct {
  Event string      `json:"event"`
  Data  []Ticker24h `json:"data"`
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
  TriggerAmount        string `json:"triggerAmount"`
  TriggerPrice         string `json:"triggerPrice"`
  TriggerType          string `json:"triggerType"`
  TriggerReference     string `json:"triggerReference"`
}

type SubscriptionCandlesResponse struct {
  Action   string             `json:"action"`
  Response SubscriptionTicker `json:"response"`
}

type SubscriptionCandles struct {
  Event    string   `json:"event"`
  Market   string   `json:"market"`
  Interval string   `json:"interval"`
  Candle   []Candle `json:"candle"`
}

type PreCandle struct {
  Event    string        `json:"event"`
  Market   string        `json:"market"`
  Interval string        `json:"interval"`
  Candle   []interface{} `json:"candle"`
}

type SubscriptionTrades struct {
  Event     string `json:"event"`
  Timestamp int    `json:"timestamp"`
  Market    string `json:"market"`
  Id        string `json:"id"`
  Amount    string `json:"amount"`
  Price     string `json:"price"`
  Side      string `json:"side"`
}

type SubscriptionBookUpdate struct {
  Event  string     `json:"event"`
  Market string     `json:"market"`
  Nonce  int        `json:"nonce"`
  Bids   [][]string `json:"bids"`
  Asks   [][]string `json:"asks"`
}

type SubscriptionTickAccObject struct {
  Action   string   `json:"action"`
  Channels []string `json:"channels"`
}

type SubscriptionTickerObject struct {
  Action   string                         `json:"action"`
  Channels []SubscriptionTickAccSubObject `json:"channels"`
}

type SubscriptionTickAccSubObject struct {
  Name    string   `json:"name"`
  Markets []string `json:"markets"`
}

type SubscriptionTradesBookObject struct {
  Action   string                            `json:"action"`
  Channels []SubscriptionTradesBookSubObject `json:"channels"`
}

type SubscriptionTradesBookSubObject struct {
  Name    string   `json:"name"`
  Markets []string `json:"markets"`
}

type SubscriptionCandlesObject struct {
  Action   string                         `json:"action"`
  Channels []SubscriptionCandlesSubObject `json:"channels"`
}

type SubscriptionCandlesSubObject struct {
  Name     string   `json:"name"`
  Interval []string `json:"interval"`
  Markets  []string `json:"markets"`
}

type LocalBook struct {
  Book map[string]Book `json:"book"`
}

type MyError struct {
  Err         error
  CustomError CustomError
}

func (e MyError) Error() string {
  if e.Err != nil {
    errorString := e.Err.Error()
    return errorString
  } else {
    return fmt.Sprintf("Error returned by API: errorCode:%d, Message: %s", e.CustomError.Code, e.CustomError.Message)
  }
}

type CustomError struct {
  Code    int    `json:"errorCode"`
  Message string `json:"error"`
  Action  string `json:"action"`
}

type Bitvavo struct {
  ApiKey, ApiSecret string
  RestUrl           string
  WsUrl             string
  AccessWindow      int
  WS                Websocket
  reconnectTimer    int
  Debugging         bool
}

type Websocket struct {
  ApiKey                   string
  WsUrl                    string
  Debugging                bool
  BookLock                 sync.Mutex
  sendLock                 sync.Mutex
  conn                     *websocket.Conn
  localBook                LocalBook
  reconnectOnError         bool
  authenticated            bool
  authenticationFailed     bool
  keepLocalBook            bool
  errChannel               chan MyError
  timeChannel              chan Time
  marketsChannel           chan []Markets
  assetsChannel            chan []Assets
  bookChannel              chan Book
  publicTradesChannel      chan []PublicTrades
  candlesChannel           chan []Candle
  ticker24hChannel         chan []Ticker24h
  tickerPriceChannel       chan []TickerPrice
  tickerBookChannel        chan []TickerBook
  placeOrderChannel        chan Order
  getOrderChannel          chan Order
  updateOrderChannel       chan Order
  cancelOrderChannel       chan CancelOrder
  getOrdersChannel         chan []Order
  cancelOrdersChannel      chan []CancelOrder
  ordersOpenChannel        chan []Order
  tradesChannel            chan []Trades
  accountChannel           chan Account
  balanceChannel           chan []Balance
  depositAssetsChannel     chan DepositAssets
  withdrawAssetsChannel    chan WithdrawAssets
  depositHistoryChannel    chan []History
  withdrawalHistoryChannel chan []History

  subscriptionTickerChannelMap map[string]chan SubscriptionTicker
  subscriptionTickerOptionsMap map[string]SubscriptionTickerObject

  subscriptionTicker24hChannelMap map[string]chan Ticker24h
  subscriptionTicker24hOptionsMap map[string]SubscriptionTickerObject

  subscriptionAccountFillChannelMap  map[string]chan SubscriptionAccountFill
  subscriptionAccountOrderChannelMap map[string]chan SubscriptionAccountOrder
  subscriptionAccountOptionsMap      map[string]SubscriptionTickerObject

  subscriptionCandlesOptionsMap map[string]map[string]SubscriptionCandlesObject
  subscriptionCandlesChannelMap map[string]map[string]chan SubscriptionCandles

  subscriptionTradesChannelMap map[string]chan SubscriptionTrades
  subscriptionTradesOptionsMap map[string]SubscriptionTradesBookObject

  subscriptionBookUpdateChannelMap map[string]chan SubscriptionBookUpdate
  subscriptionBookUpdateOptionsMap map[string]SubscriptionTradesBookObject

  subscriptionBookChannelMap       map[string]chan Book
  subscriptionBookOptionsFirstMap  map[string]map[string]string
  subscriptionBookOptionsSecondMap map[string]SubscriptionTradesBookObject
}

func (bitvavo Bitvavo) NewWebsocket() (*Websocket, chan MyError) {
  ws := Websocket{}
  ws.Debugging = bitvavo.Debugging
  ws.ApiKey = bitvavo.ApiKey
  ws.conn = bitvavo.InitWS()
  ws.reconnectOnError = true
  ws.authenticated = false
  ws.authenticationFailed = false
  ws.keepLocalBook = false
  errChannel := make(chan MyError)
  ws.errChannel = errChannel
  go bitvavo.handleMessage(&ws)
  return &ws, errChannel
}

func (bitvavo Bitvavo) createSignature(timestamp string, method string, url string, body map[string]string, ApiSecret string) string {
  result := timestamp + method + "/v2" + url
  if len(body) != 0 {
    bodyString, err := json.Marshal(body)
    if err != nil {
      errorToConsole("Converting map to string went wrong!")
    }
    result = result + string(bodyString)
  }
  h := hmac.New(sha256.New, []byte(ApiSecret))
  h.Write([]byte(result))
  sha := hex.EncodeToString(h.Sum(nil))
  return sha
}

func (bitvavo Bitvavo) sendPublic(endpoint string) []byte {
  client := &http.Client{}
  req, err := http.NewRequest("GET", endpoint, bytes.NewBuffer(nil))
  if err != nil {
    errorToConsole("We caught error " + err.Error())
  }
  if bitvavo.ApiKey != "" {
    millis := time.Now().UnixNano() / 1000000
    timeString := strconv.FormatInt(millis, 10)
    sig := bitvavo.createSignature(timeString, "GET", strings.Replace(endpoint, bitvavo.RestUrl, "", 1), map[string]string{}, bitvavo.ApiSecret)
    req.Header.Set("bitvavo-access-key", bitvavo.ApiKey)
    req.Header.Set("bitvavo-access-signature", sig)
    req.Header.Set("bitvavo-access-timestamp", timeString)
    req.Header.Set("bitvavo-access-window", strconv.Itoa(bitvavo.AccessWindow))
  }
  req.Header.Set("Content-Type", "application/json")
  resp, err := client.Do(req)
  if err != nil {
    errorToConsole("Caught error " + err.Error())
    return []byte("caught error")
  } else {
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      errorToConsole("Caught error " + err.Error())
      return []byte("caught error")
    }
    updateRateLimit(resp.Header)
    return body
  }
}

func (bitvavo Bitvavo) sendPrivate(endpoint string, postfix string, body map[string]string, method string) []byte {
  millis := time.Now().UnixNano() / 1000000
  timeString := strconv.FormatInt(millis, 10)
  sig := bitvavo.createSignature(timeString, method, (endpoint + postfix), body, bitvavo.ApiSecret)
  url := bitvavo.RestUrl + endpoint + postfix
  client := &http.Client{}
  byteBody := []byte{}
  if len(body) != 0 {
    bodyString, err := json.Marshal(body)
    if err != nil {
      errorToConsole("We caught error " + err.Error())
    }
    byteBody = []byte(bodyString)
  } else {
    byteBody = nil
  }
  req, err := http.NewRequest(method, url, bytes.NewBuffer(byteBody))
  req.Header.Set("bitvavo-access-key", bitvavo.ApiKey)
  req.Header.Set("bitvavo-access-signature", sig)
  req.Header.Set("bitvavo-access-timestamp", timeString)
  req.Header.Set("bitvavo-access-window", strconv.Itoa(bitvavo.AccessWindow))
  req.Header.Set("content-type", "application/json")
  resp, err := client.Do(req)
  if err != nil {
    errorToConsole("We caught an error " + err.Error())
    return nil
  }
  defer resp.Body.Close()
  respBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    errorToConsole("Caught error " + err.Error())
    return nil
  }
  updateRateLimit(resp.Header)
  return respBody
}

func checkLimit() {
  now := int(time.Nanosecond * time.Duration(time.Now().UnixNano()) / time.Millisecond)
  if rateLimitReset <= now {
    rateLimitRemaining = 1000
  }
}

func updateRateLimit(response http.Header) {
  for key, value := range response {
    if key == "bitvavo-ratelimit-remaining" {
      rateLimitRemaining, _ = strconv.Atoi(value[0])
    }
    if key == "bitvavo-ratelimit-resetat" {
      rateLimitReset, _ = strconv.Atoi(value[0])
      now := int(time.Nanosecond * time.Duration(time.Now().UnixNano()) / time.Millisecond)
      var timeToWait = rateLimitReset - now
      time.AfterFunc(time.Duration(timeToWait)*time.Millisecond, checkLimit)
    }
  }
}

func (bitvavo Bitvavo) GetRemainingLimit() int {
  return rateLimitRemaining
}

func (bitvavo Bitvavo) createPostfix(options map[string]string) string {
  result := []string{}
  for k := range options {
    result = append(result, (k + "=" + options[k]))
  }
  params := strings.Join(result, "&")
  if len(params) != 0 {
    params = "?" + params
  }
  return params
}

func handleAPIError(jsonResponse []byte) error {
  var e CustomError
  err := json.Unmarshal(jsonResponse, &e)
  if err != nil {
    errorToConsole("error casting")
    return MyError{Err: err}
  }
  if e.Code == 105 {
    rateLimitRemaining = 0
    rateLimitReset, _ = strconv.Atoi(strings.Split(strings.Split(e.Message, " at ")[1], ".")[0])
    now := int(time.Nanosecond * time.Duration(time.Now().UnixNano()) / time.Millisecond)
    var timeToWait = rateLimitReset - now
    time.AfterFunc(time.Duration(timeToWait)*time.Millisecond, checkLimit)
  }
  return MyError{CustomError: e}
}

func (bitvavo Bitvavo) Time() (Time, error) {
  jsonResponse := bitvavo.sendPublic(bitvavo.RestUrl + "/time")
  var t Time

  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return Time{}, MyError{Err: err}
  }
  if t.Time == 0.0 {
    return Time{}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: market
func (bitvavo Bitvavo) Markets(options map[string]string) ([]Markets, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPublic(bitvavo.RestUrl + "/markets" + postfix)
  t := make([]Markets, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []Markets{Markets{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: symbol
func (bitvavo Bitvavo) Assets(options map[string]string) ([]Assets, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPublic(bitvavo.RestUrl + "/assets" + postfix)
  t := make([]Assets, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []Assets{Assets{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: depth
func (bitvavo Bitvavo) Book(symbol string, options map[string]string) (Book, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPublic(bitvavo.RestUrl + "/" + symbol + "/book" + postfix)
  var t Book
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return Book{}, MyError{Err: err}
  }
  if t.Market == "" {
    return Book{}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: limit, start, end, tradeIdFrom, tradeIdTo
func (bitvavo Bitvavo) PublicTrades(symbol string, options map[string]string) ([]PublicTrades, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPublic(bitvavo.RestUrl + "/" + symbol + "/trades" + postfix)
  t := make([]PublicTrades, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []PublicTrades{PublicTrades{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: limit, start, end
func (bitvavo Bitvavo) Candles(symbol string, interval string, options map[string]string) ([]Candle, error) {
  options["interval"] = interval
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPublic(bitvavo.RestUrl + "/" + symbol + "/candles" + postfix)
  var t []interface{}
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []Candle{Candle{}}, MyError{Err: err}
  }
  var candles []Candle
  for i := 0; i < len(t); i++ {
    entry := reflect.ValueOf(t[i])
    candles = append(candles, Candle{Timestamp: int(entry.Index(0).Interface().(float64)), Open: entry.Index(1).Interface().(string), High: entry.Index(2).Interface().(string), Low: entry.Index(3).Interface().(string), Close: entry.Index(4).Interface().(string), Volume: entry.Index(5).Interface().(string)})
  }
  return candles, nil
}

// options: market
func (bitvavo Bitvavo) TickerPrice(options map[string]string) ([]TickerPrice, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPublic(bitvavo.RestUrl + "/ticker/price" + postfix)
  t := make([]TickerPrice, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    var t TickerPrice
    err = json.Unmarshal(jsonResponse, &t)
    if err != nil {
      return []TickerPrice{TickerPrice{}}, MyError{Err: err}
    }
    if t.Market == "" {
      return []TickerPrice{TickerPrice{}}, handleAPIError(jsonResponse)
    }
    return []TickerPrice{t}, nil
  }
  return t, nil
}

// options: market
func (bitvavo Bitvavo) TickerBook(options map[string]string) ([]TickerBook, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPublic(bitvavo.RestUrl + "/ticker/book" + postfix)
  t := make([]TickerBook, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    var t TickerBook
    err = json.Unmarshal(jsonResponse, &t)
    if err != nil {
      return []TickerBook{TickerBook{}}, MyError{Err: err}
    }
    if t.Market == "" {
      return []TickerBook{TickerBook{}}, handleAPIError(jsonResponse)
    }
    return []TickerBook{t}, nil
  }
  return t, nil
}

// options: market
func (bitvavo Bitvavo) Ticker24h(options map[string]string) ([]Ticker24h, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPublic(bitvavo.RestUrl + "/ticker/24h" + postfix)
  t := make([]Ticker24h, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    var t Ticker24h
    err = json.Unmarshal(jsonResponse, &t)
    if err != nil {
      return []Ticker24h{Ticker24h{}}, MyError{Err: err}
    }
    if t.Market == "" {
      return []Ticker24h{Ticker24h{}}, handleAPIError(jsonResponse)
    }
    return []Ticker24h{t}, nil
  }
  return t, nil
}

// optional body parameters: limit:(amount, price, postOnly), market:(amount, amountQuote, disableMarketProtection)
//                           stopLoss/takeProfit:(amount, amountQuote, disableMarketProtection, triggerType, triggerReference, triggerAmount)
//                           stopLossLimit/takeProfitLimit:(amount, price, postOnly, triggerType, triggerReference, triggerAmount)
//                           all orderTypes: timeInForce, selfTradePrevention, responseRequired
func (bitvavo Bitvavo) PlaceOrder(market string, side string, orderType string, body map[string]string) (Order, error) {
  body["market"] = market
  body["side"] = side
  body["orderType"] = orderType
  jsonResponse := bitvavo.sendPrivate("/order", "", body, "POST")
  var t Order
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return Order{}, MyError{Err: err}
  }
  if t.OrderId == "" {
    return Order{}, handleAPIError(jsonResponse)
  }
  return t, nil
}

func (bitvavo Bitvavo) GetOrder(market string, orderId string) (Order, error) {
  options := map[string]string{"market": market, "orderId": orderId}
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/order", postfix, map[string]string{}, "GET")
  var t Order
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return Order{}, MyError{Err: err}
  }
  if t.OrderId == "" {
    return Order{}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// Optional body parameters: limit:(amount, amountRemaining, price, timeInForce, selfTradePrevention, postOnly)
//               untriggered stopLoss/takeProfit:(amount, amountQuote, disableMarketProtection, triggerType, triggerReference, triggerAmount)
//                           stopLossLimit/takeProfitLimit: (amount, price, postOnly, triggerType, triggerReference, triggerAmount)
func (bitvavo Bitvavo) UpdateOrder(market string, orderId string, body map[string]string) (Order, error) {
  body["market"] = market
  body["orderId"] = orderId
  jsonResponse := bitvavo.sendPrivate("/order", "", body, "PUT")
  var t Order
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return Order{}, MyError{Err: err}
  }
  if t.OrderId == "" {
    return Order{}, handleAPIError(jsonResponse)
  }
  return t, nil
}

func (bitvavo Bitvavo) CancelOrder(market string, orderId string) (CancelOrder, error) {
  options := map[string]string{"market": market, "orderId": orderId}
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/order", postfix, map[string]string{}, "DELETE")
  var t CancelOrder
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return CancelOrder{}, MyError{Err: err}
  }
  if t.OrderId == "" {
    return CancelOrder{}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: limit, start, end, orderIdFrom, orderIdTo
func (bitvavo Bitvavo) GetOrders(market string, options map[string]string) ([]Order, error) {
  options["market"] = market
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/orders", postfix, map[string]string{}, "GET")
  t := make([]Order, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []Order{Order{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: market
func (bitvavo Bitvavo) CancelOrders(options map[string]string) ([]CancelOrder, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/orders", postfix, map[string]string{}, "DELETE")
  t := make([]CancelOrder, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []CancelOrder{CancelOrder{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: market
func (bitvavo Bitvavo) OrdersOpen(options map[string]string) ([]Order, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/ordersOpen", postfix, map[string]string{}, "GET")
  t := make([]Order, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []Order{Order{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: limit, start, end, tradeIdFrom, tradeIdTo
func (bitvavo Bitvavo) Trades(market string, options map[string]string) ([]Trades, error) {
  options["market"] = market
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/trades", postfix, map[string]string{}, "GET")
  t := make([]Trades, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []Trades{Trades{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

func (bitvavo Bitvavo) Account() (Account, error) {
  jsonResponse := bitvavo.sendPrivate("/account", "", map[string]string{}, "GET")
  var t Account
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return Account{}, MyError{Err: err}
  }
  return t, nil
}

// options: symbol
func (bitvavo Bitvavo) Balance(options map[string]string) ([]Balance, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/balance", postfix, map[string]string{}, "GET")
  t := make([]Balance, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []Balance{Balance{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

func (bitvavo Bitvavo) DepositAssets(symbol string) (DepositAssets, error) {
  options := map[string]string{"symbol": symbol}
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/deposit", postfix, map[string]string{}, "GET")
  var t DepositAssets
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return DepositAssets{}, MyError{Err: err}
  }
  if t.Address == "" {
    return DepositAssets{}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// optional body parameters: paymentId, internal, addWithdrawalFee
func (bitvavo Bitvavo) WithdrawAssets(symbol string, amount string, address string, body map[string]string) (WithdrawAssets, error) {
  body["symbol"] = symbol
  body["amount"] = amount
  body["address"] = address
  jsonResponse := bitvavo.sendPrivate("/withdrawal", "", body, "POST")
  var t WithdrawAssets
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return WithdrawAssets{}, MyError{Err: err}
  }
  if t.Symbol == "" {
    return WithdrawAssets{}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: symbol, limit, start, end
func (bitvavo Bitvavo) DepositHistory(options map[string]string) ([]History, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/depositHistory", postfix, map[string]string{}, "GET")
  t := make([]History, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []History{History{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

// options: symbol, limit, start, end
func (bitvavo Bitvavo) WithdrawalHistory(options map[string]string) ([]History, error) {
  postfix := bitvavo.createPostfix(options)
  jsonResponse := bitvavo.sendPrivate("/withdrawalHistory", postfix, map[string]string{}, "GET")
  t := make([]History, 0)
  err := json.Unmarshal(jsonResponse, &t)
  if err != nil {
    return []History{History{}}, handleAPIError(jsonResponse)
  }
  return t, nil
}

func handleError(err error) bool {
  if err != nil {
    errorToConsole(err.Error())
    return true
  }
  return false
}

func sortAndInsert(update [][]string, book [][]string, asksCompare bool) [][]string {
  for i := 0; i < len(update); i++ {
    entrySet := false
    updateEntry := update[i]
    for j := 0; j < len(book); j++ {
      bookItem := book[j]
      updatePrice, _ := strconv.ParseFloat(updateEntry[0], 64)
      bookPrice, _ := strconv.ParseFloat(bookItem[0], 64)
      if asksCompare {
        if updatePrice < bookPrice {
          book = append(book, make([]string, 2))
          copy(book[j+1:], book[j:])
          book[j] = updateEntry
          entrySet = true
          break
        }
      } else {
        if updatePrice > bookPrice {
          book = append(book, make([]string, 2))
          copy(book[j+1:], book[j:])
          book[j] = updateEntry
          entrySet = true
          break
        }
      }
      if updatePrice == bookPrice {
        updateAmount, _ := strconv.ParseFloat(updateEntry[1], 64)
        if updateAmount > 0.0 {
          book[j] = updateEntry
          entrySet = true
          break
        } else {
          book = append(book[:j], book[j+1:]...)
          entrySet = true
          break
        }
      }
    }
    if entrySet == false {
      book = append(book, updateEntry)
    }
  }
  return book
}

func addToBook(t SubscriptionBookUpdate, ws *Websocket) {
  ws.BookLock.Lock()
  defer ws.BookLock.Unlock()
  var book = ws.localBook.Book[t.Market]
  book.Bids = sortAndInsert(t.Bids, ws.localBook.Book[t.Market].Bids, false)
  book.Asks = sortAndInsert(t.Asks, ws.localBook.Book[t.Market].Asks, true)
  if book.Nonce != (t.Nonce - 1) {
    ws.SubscriptionBook(t.Market, ws.subscriptionBookOptionsFirstMap[t.Market])
    return
  }
  book.Nonce = t.Nonce
  ws.localBook.Book[t.Market] = book
  ws.subscriptionBookChannelMap[t.Market] <- ws.localBook.Book[t.Market]
}

func (bitvavo Bitvavo) DebugToConsole(message string) {
  if bitvavo.Debugging {
    fmt.Println(time.Now().Format("15:04:05") + " DEBUG: " + message)
  }
}

func (ws *Websocket) DebugToConsole(message string) {
  if ws.Debugging {
    fmt.Println(time.Now().Format("15:04:05") + " DEBUG: " + message)
  }
}

func errorToConsole(message string) {
  fmt.Println(time.Now().Format("15:04:05") + " ERROR: " + message)
}

func (bitvavo Bitvavo) reconnect(ws *Websocket) {
  bitvavo.DebugToConsole("Reconnecting")
  time.Sleep(500 * time.Millisecond)
  ws.authenticated = false
  ws.authenticationFailed = false
  ws.keepLocalBook = false
  ws.conn = bitvavo.InitWS()

  go bitvavo.handleMessage(ws)

  for market := range ws.subscriptionTickerOptionsMap {
    myMessage, _ := json.Marshal(ws.subscriptionTickerOptionsMap[market])
    ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  }
  for market := range ws.subscriptionTicker24hOptionsMap {
    myMessage, _ := json.Marshal(ws.subscriptionTicker24hOptionsMap[market])
    ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  }
  for market := range ws.subscriptionAccountOptionsMap {
    myMessage, _ := json.Marshal(ws.subscriptionAccountOptionsMap[market])
    ws.sendPrivate(myMessage)
  }
  for market := range ws.subscriptionTradesOptionsMap {
    myMessage, _ := json.Marshal(ws.subscriptionTradesOptionsMap[market])
    ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  }
  for market := range ws.subscriptionCandlesOptionsMap {
    for interval := range ws.subscriptionCandlesOptionsMap[market] {
      myMessage, _ := json.Marshal(ws.subscriptionCandlesOptionsMap[market][interval])
      ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
    }
  }
  for market := range ws.subscriptionBookUpdateOptionsMap {
    myMessage, _ := json.Marshal(ws.subscriptionBookUpdateOptionsMap[market])
    ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  }
  for market := range ws.subscriptionBookOptionsFirstMap {
    ws.keepLocalBook = true
    myMessage, _ := json.Marshal(ws.subscriptionBookOptionsFirstMap[market])
    ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  }
  for market := range ws.subscriptionBookOptionsSecondMap {
    myMessage, _ := json.Marshal(ws.subscriptionBookOptionsSecondMap[market])
    ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  }
}

func (bitvavo Bitvavo) handleMessage(ws *Websocket) {
  for {
    bitvavo.reconnectTimer = 100
    _, message, err := ws.conn.ReadMessage()
    if handleError(err) {
      if (ws.reconnectOnError) {
        err = ws.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
        bitvavo.reconnect(ws)
      }
      return
    }
    bitvavo.DebugToConsole("FULL RESPONSE: " + string(message))
    var x map[string]interface{}
    err = json.Unmarshal(message, &x)
    if handleError(err) {
      errorToConsole("We are returning, this should not happen...")
      return
    }
    if _, ok := x["error"]; ok {
      if x["action"] == "authenticate" {
        ws.authenticationFailed = true
      }
      var t CustomError
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        errorToConsole("Something failed during reception of the authentication response.")
        return
      }
      ws.errChannel <- MyError{CustomError: t}
    }
    if x["event"] == "authenticate" {
      ws.authenticated = true
    } else if x["event"] == "book" {
      var t SubscriptionBookUpdate
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      market, _ := x["market"].(string)
      if ws.subscriptionBookUpdateChannelMap[market] != nil {
        ws.subscriptionBookUpdateChannelMap[market] <- t
      }
      if ws.keepLocalBook {
        addToBook(t, ws)
      }
    } else if x["event"] == "trade" {
      var t SubscriptionTrades
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      market, _ := x["market"].(string)
      if ws.subscriptionTradesChannelMap[market] != nil {
        ws.subscriptionTradesChannelMap[market] <- t
      }
    } else if x["event"] == "fill" {
      var t SubscriptionAccountFill
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      market, _ := x["market"].(string)
      if ws.subscriptionAccountFillChannelMap[market] != nil {
        ws.subscriptionAccountFillChannelMap[market] <- t
      }
    } else if x["event"] == "order" {
      var t SubscriptionAccountOrder
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      market, _ := x["market"].(string)
      if ws.subscriptionAccountOrderChannelMap[market] != nil {
        ws.subscriptionAccountOrderChannelMap[market] <- t
      }
    } else if x["event"] == "ticker" {
      var t SubscriptionTicker
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      market, _ := x["market"].(string)
      ws.subscriptionTickerChannelMap[market] <- t
    } else if x["event"] == "ticker24h" {
      var t SubscriptionTicker24h
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      for i := 0; i < len(t.Data); i++ {
        ws.subscriptionTicker24hChannelMap[t.Data[i].Market] <- t.Data[i]
      }
    } else if x["event"] == "candle" {
      var t PreCandle
      err := json.Unmarshal(message, &t)
      if err != nil {
        return
      }
      var candles []Candle
      for i := 0; i < len(t.Candle); i++ {
        entry := reflect.ValueOf(t.Candle[i])
        candles = append(candles, Candle{Timestamp: int(entry.Index(0).Interface().(float64)), Open: entry.Index(1).Interface().(string), High: entry.Index(2).Interface().(string), Low: entry.Index(3).Interface().(string), Close: entry.Index(4).Interface().(string), Volume: entry.Index(5).Interface().(string)})
      }
      market, _ := x["market"].(string)
      interval, _ := x["interval"].(string)
      ws.subscriptionCandlesChannelMap[market][interval] <- SubscriptionCandles{Event: t.Event, Market: t.Market, Interval: t.Interval, Candle: candles}
    }
    if x["action"] == "getTime" {
      var t TimeResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.timeChannel <- t.Response

    } else if x["action"] == "getMarkets" {
      var t MarketsResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.marketsChannel <- t.Response
    } else if x["action"] == "getAssets" {
      var t AssetsResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.assetsChannel <- t.Response
    } else if x["action"] == "getBook" {
      var t BookResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      if ws.bookChannel != nil {
        ws.bookChannel <- t.Response
      }
      if ws.keepLocalBook {
        ws.localBook.Book[t.Response.Market] = t.Response
        ws.subscriptionBookChannelMap[t.Response.Market] <- ws.localBook.Book[t.Response.Market]
      }
    } else if x["action"] == "getTrades" {
      var t PublicTradesResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.publicTradesChannel <- t.Response
    } else if x["action"] == "getCandles" {
      var t CandlesResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      var candles []Candle
      for i := 0; i < len(t.Response); i++ {
        entry := reflect.ValueOf(t.Response[i])
        candles = append(candles, Candle{Timestamp: int(entry.Index(0).Interface().(float64)), Open: entry.Index(1).Interface().(string), High: entry.Index(2).Interface().(string), Low: entry.Index(3).Interface().(string), Close: entry.Index(4).Interface().(string), Volume: entry.Index(5).Interface().(string)})
      }
      ws.candlesChannel <- candles
    } else if x["action"] == "getTicker24h" {
      var t Ticker24hResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.ticker24hChannel <- t.Response
    } else if x["action"] == "getTickerPrice" {
      var t TickerPriceResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.tickerPriceChannel <- t.Response
    } else if x["action"] == "getTickerBook" {
      var t TickerBookResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.tickerBookChannel <- t.Response
    } else if x["action"] == "privateCreateOrder" {
      var t PlaceOrderResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.placeOrderChannel <- t.Response
    } else if x["action"] == "privateUpdateOrder" {
      var t UpdateOrderResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      if t.Response.OrderId != "" {
        ws.updateOrderChannel <- t.Response
      }
    } else if x["action"] == "privateGetOrder" {
      var t GetOrderResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.getOrderChannel <- t.Response
    } else if x["action"] == "privateCancelOrder" {
      var t CancelOrderResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.cancelOrderChannel <- t.Response
    } else if x["action"] == "privateGetOrders" {
      var t GetOrdersResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.getOrdersChannel <- t.Response
    } else if x["action"] == "privateCancelOrders" {
      var t CancelOrdersResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.cancelOrdersChannel <- t.Response
    } else if x["action"] == "privateGetOrdersOpen" {
      var t OrdersOpenResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.ordersOpenChannel <- t.Response
    } else if x["action"] == "privateGetTrades" {
      var t TradesResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.tradesChannel <- t.Response
    } else if x["action"] == "privateGetAccount" {
      var t AccountResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.accountChannel <- t.Response
    } else if x["action"] == "privateGetBalance" {
      var t BalanceResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.balanceChannel <- t.Response
    } else if x["action"] == "privateDepositAssets" {
      var t DepositAssetsResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.depositAssetsChannel <- t.Response
    } else if x["action"] == "privateWithdrawAssets" {
      var t WithdrawAssetsResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.withdrawAssetsChannel <- t.Response
    } else if x["action"] == "privateGetDepositHistory" {
      var t HistoryResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.depositHistoryChannel <- t.Response
    } else if x["action"] == "privateGetWithdrawalHistory" {
      var t HistoryResponse
      err = json.Unmarshal(message, &t)
      if handleError(err) {
        return
      }
      ws.withdrawalHistoryChannel <- t.Response
    }
  }
  errorToConsole("HandleMessage has ended, messages will no longer be received, please restart.")
}

func (bitvavo Bitvavo) InitWS() *websocket.Conn {
  bitvavo.reconnectTimer = 100
  uri, _ := url.Parse(bitvavo.WsUrl)
  c, _, err := websocket.DefaultDialer.Dial(uri.String(), nil)
  if err != nil {
    errorToConsole("Caught error " + err.Error())
    c = bitvavo.retryReconnect()
  }
  now := time.Now()
  nanos := now.UnixNano()
  millis := nanos / 1000000
  timestamp := strconv.FormatInt(millis, 10)
  if bitvavo.ApiKey != "" {
    authenticate := map[string]string{"action": "authenticate", "key": bitvavo.ApiKey, "signature": bitvavo.createSignature(timestamp, "GET", "/websocket", map[string]string{}, bitvavo.ApiSecret), "timestamp": timestamp, "window": strconv.Itoa(bitvavo.AccessWindow)}
    myMessage, _ := json.Marshal(authenticate)
    bitvavo.DebugToConsole("SENDING: " + string(myMessage))
    c.WriteMessage(websocket.TextMessage, []byte(myMessage))
  }
  return c
}

func (bitvavo Bitvavo) retryReconnect() *websocket.Conn {
  time.Sleep(time.Duration(bitvavo.reconnectTimer) * time.Millisecond)
  bitvavo.reconnectTimer = bitvavo.reconnectTimer * 2
  bitvavo.DebugToConsole("We waited for " + strconv.Itoa(bitvavo.reconnectTimer) + " seconds to reconnect")
  uri, _ := url.Parse(bitvavo.WsUrl)
  c, _, err := websocket.DefaultDialer.Dial(uri.String(), nil)
  if err != nil {
    errorToConsole("Caught error " + err.Error())
    c = bitvavo.retryReconnect()
    return c
  }
  return c
}

func (ws *Websocket) Close() {
  ws.reconnectOnError = false
  ws.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func (ws *Websocket) Time() chan Time {
  ws.timeChannel = make(chan Time, 100)
  myMessage, _ := json.Marshal(map[string]string{"action": "getTime"})
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.timeChannel
}

// options: market
func (ws *Websocket) Markets(options map[string]string) chan []Markets {
  ws.marketsChannel = make(chan []Markets, 100)
  options["action"] = "getMarkets"
  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.marketsChannel
}

// options: symbol
func (ws *Websocket) Assets(options map[string]string) chan []Assets {
  ws.assetsChannel = make(chan []Assets, 100)
  options["action"] = "getAssets"
  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.assetsChannel
}

// options: depth
func (ws *Websocket) Book(market string, options map[string]string) chan Book {
  ws.bookChannel = make(chan Book, 100)
  options["market"] = market
  options["action"] = "getBook"
  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.bookChannel
}

// options: limit, start, end, tradeIdFrom, tradeIdTo
func (ws *Websocket) PublicTrades(market string, options map[string]string) chan []PublicTrades {
  ws.publicTradesChannel = make(chan []PublicTrades, 100)
  options["market"] = market
  options["action"] = "getTrades"
  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.publicTradesChannel
}

// options: limit, start, end
func (ws *Websocket) Candles(market string, interval string, options map[string]string) chan []Candle {
  ws.candlesChannel = make(chan []Candle, 100)
  options["market"] = market
  options["interval"] = interval
  options["action"] = "getCandles"
  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.candlesChannel
}

// options: market
func (ws *Websocket) Ticker24h(options map[string]string) chan []Ticker24h {
  ws.ticker24hChannel = make(chan []Ticker24h, 100)
  options["action"] = "getTicker24h"
  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.ticker24hChannel
}

// options: market
func (ws *Websocket) TickerPrice(options map[string]string) chan []TickerPrice {
  ws.tickerPriceChannel = make(chan []TickerPrice, 100)
  options["action"] = "getTickerPrice"
  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.tickerPriceChannel
}

// options: market
func (ws *Websocket) TickerBook(options map[string]string) chan []TickerBook {
  ws.tickerBookChannel = make(chan []TickerBook, 100)
  options["action"] = "getTickerBook"
  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.tickerBookChannel
}

func (ws *Websocket) sendPrivate(msg []byte) {
  if ws.ApiKey == "" {
    errorToConsole("You did not set the API key, but requested a private function.")
    return
  }
  if ws.authenticated == true {
    ws.DebugToConsole("SENDING: " + string(msg))
    ws.conn.WriteMessage(websocket.TextMessage, msg)
  } else {
    if ws.authenticationFailed == false {
      ws.DebugToConsole("Waiting 100 milliseconds till authenticated is received")
      time.Sleep(100 * time.Millisecond)
      ws.sendPrivate(msg)
    } else {
      errorToConsole("Authentication is required for sending this message, but authentication failed.")
    }
  }
}

// optional body parameters: limit:(amount, price, postOnly), market:(amount, amountQuote, disableMarketProtection)
//                           stopLoss/takeProfit:(amount, amountQuote, disableMarketProtection, triggerType, triggerReference, triggerAmount)
//                           stopLossLimit/takeProfitLimit:(amount, price, postOnly, triggerType, triggerReference, triggerAmount)
//                           all orderTypes: timeInForce, selfTradePrevention, responseRequired
func (ws *Websocket) PlaceOrder(market string, side string, orderType string, body map[string]string) chan Order {
  body["market"] = market
  body["side"] = side
  body["orderType"] = orderType
  ws.placeOrderChannel = make(chan Order, 100)
  body["action"] = "privateCreateOrder"
  myMessage, _ := json.Marshal(body)
  go ws.sendPrivate(myMessage)
  return ws.placeOrderChannel
}

func (ws *Websocket) GetOrder(market string, orderId string) chan Order {
  ws.getOrderChannel = make(chan Order, 100)
  options := map[string]string{"action": "privateGetOrder", "market": market, "orderId": orderId}
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.getOrderChannel
}

// Optional body parameters: limit:(amount, amountRemaining, price, timeInForce, selfTradePrevention, postOnly)
//               untriggered stopLoss/takeProfit:(amount, amountQuote, disableMarketProtection, triggerType, triggerReference, triggerAmount)
//                           stopLossLimit/takeProfitLimit: (amount, price, postOnly, triggerType, triggerReference, triggerAmount)
func (ws *Websocket) UpdateOrder(market string, orderId string, body map[string]string) chan Order {
  ws.updateOrderChannel = make(chan Order, 100)
  body["market"] = market
  body["orderId"] = orderId
  body["action"] = "privateUpdateOrder"
  myMessage, _ := json.Marshal(body)
  go ws.sendPrivate(myMessage)
  return ws.updateOrderChannel
}

func (ws *Websocket) CancelOrder(market string, orderId string) chan CancelOrder {
  ws.cancelOrderChannel = make(chan CancelOrder, 100)
  options := map[string]string{"action": "privateCancelOrder", "market": market, "orderId": orderId}
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.cancelOrderChannel
}

// options: limit, start, end, orderIdFrom, orderIdTo
func (ws *Websocket) GetOrders(market string, options map[string]string) chan []Order {
  ws.getOrdersChannel = make(chan []Order, 100)
  options["action"] = "privateGetOrders"
  options["market"] = market
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.getOrdersChannel
}

func (ws *Websocket) CancelOrders(options map[string]string) chan []CancelOrder {
  ws.cancelOrdersChannel = make(chan []CancelOrder, 100)
  options["action"] = "privateCancelOrders"
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.cancelOrdersChannel
}

// options: market
func (ws *Websocket) OrdersOpen(options map[string]string) chan []Order {
  ws.ordersOpenChannel = make(chan []Order, 100)
  options["action"] = "privateGetOrdersOpen"
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.ordersOpenChannel
}

// options: limit, start, end, tradeIdFrom, tradeIdTo
func (ws *Websocket) Trades(market string, options map[string]string) chan []Trades {
  ws.tradesChannel = make(chan []Trades, 100)
  options["action"] = "privateGetTrades"
  options["market"] = market
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.tradesChannel
}

func (ws *Websocket) Account() chan Account {
  ws.accountChannel = make(chan Account, 100)
  options := map[string]string{"action": "privateGetAccount"}
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.accountChannel
}

// options: symbol
func (ws *Websocket) Balance(options map[string]string) chan []Balance {
  ws.balanceChannel = make(chan []Balance, 100)
  options["action"] = "privateGetBalance"
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.balanceChannel
}

func (ws *Websocket) DepositAssets(symbol string) chan DepositAssets {
  ws.depositAssetsChannel = make(chan DepositAssets, 100)
  options := map[string]string{"action": "privateDepositAssets", "symbol": symbol}
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.depositAssetsChannel
}

// optional body parameters: paymentId, internal, addWithdrawalFee
func (ws *Websocket) WithdrawAssets(symbol string, amount string, address string, body map[string]string) chan WithdrawAssets {
  ws.withdrawAssetsChannel = make(chan WithdrawAssets, 100)
  body["symbol"] = symbol
  body["amount"] = amount
  body["address"] = address
  body["action"] = "privateWithdrawAssets"
  myMessage, _ := json.Marshal(body)
  go ws.sendPrivate(myMessage)
  return ws.withdrawAssetsChannel
}

// options: symbol, limit, start, end
func (ws *Websocket) DepositHistory(options map[string]string) chan []History {
  ws.depositHistoryChannel = make(chan []History, 100)
  options["action"] = "privateGetDepositHistory"
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.depositHistoryChannel
}

// options: symbol, limit, start, end
func (ws *Websocket) WithdrawalHistory(options map[string]string) chan []History {
  ws.withdrawalHistoryChannel = make(chan []History, 100)
  options["action"] = "privateGetWithdrawalHistory"
  myMessage, _ := json.Marshal(options)
  go ws.sendPrivate(myMessage)
  return ws.withdrawalHistoryChannel
}

func (ws *Websocket) SubscriptionTicker(market string) chan SubscriptionTicker {
  options := SubscriptionTickerObject{Action: "subscribe", Channels: []SubscriptionTickAccSubObject{SubscriptionTickAccSubObject{Name: "ticker", Markets: []string{market}}}}
  if ws.subscriptionTickerChannelMap == nil {
    ws.subscriptionTickerChannelMap = map[string]chan SubscriptionTicker{}
  }
  if ws.subscriptionTickerOptionsMap == nil {
    ws.subscriptionTickerOptionsMap = map[string]SubscriptionTickerObject{}
  }
  ws.subscriptionTickerChannelMap[market] = make(chan SubscriptionTicker, 100)
  ws.subscriptionTickerOptionsMap[market] = options

  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.subscriptionTickerChannelMap[market]
}

func (ws *Websocket) SubscriptionTicker24h(market string) chan Ticker24h {
  options := SubscriptionTickerObject{Action: "subscribe", Channels: []SubscriptionTickAccSubObject{SubscriptionTickAccSubObject{Name: "ticker24h", Markets: []string{market}}}}
  if ws.subscriptionTicker24hChannelMap == nil {
    ws.subscriptionTicker24hChannelMap = map[string]chan Ticker24h{}
  }
  if ws.subscriptionTicker24hOptionsMap == nil {
    ws.subscriptionTicker24hOptionsMap = map[string]SubscriptionTickerObject{}
  }
  ws.subscriptionTicker24hChannelMap[market] = make(chan Ticker24h, 100)
  ws.subscriptionTicker24hOptionsMap[market] = options

  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.subscriptionTicker24hChannelMap[market]
}

func (ws *Websocket) SubscriptionAccount(market string) (chan SubscriptionAccountOrder, chan SubscriptionAccountFill) {
  options := SubscriptionTickerObject{Action: "subscribe", Channels: []SubscriptionTickAccSubObject{SubscriptionTickAccSubObject{Name: "account", Markets: []string{market}}}}
  if ws.subscriptionAccountOrderChannelMap == nil {
    ws.subscriptionAccountOrderChannelMap = map[string]chan SubscriptionAccountOrder{}
  }
  if ws.subscriptionAccountFillChannelMap == nil {
    ws.subscriptionAccountFillChannelMap = map[string]chan SubscriptionAccountFill{}
  }
  ws.subscriptionAccountOrderChannelMap[market] = make(chan SubscriptionAccountOrder, 100)
  ws.subscriptionAccountFillChannelMap[market] = make(chan SubscriptionAccountFill, 100)

  if ws.subscriptionAccountOptionsMap == nil {
    ws.subscriptionAccountOptionsMap = map[string]SubscriptionTickerObject{}
  }
  ws.subscriptionAccountOptionsMap[market] = options

  myMessage, _ := json.Marshal(options)

  ws.sendPrivate(myMessage)
  return ws.subscriptionAccountOrderChannelMap[market], ws.subscriptionAccountFillChannelMap[market]
}

func (ws *Websocket) SubscriptionCandles(market string, interval string) chan SubscriptionCandles {
  options := SubscriptionCandlesObject{Action: "subscribe", Channels: []SubscriptionCandlesSubObject{SubscriptionCandlesSubObject{Name: "candles", Interval: []string{interval}, Markets: []string{market}}}}
  if ws.subscriptionCandlesChannelMap == nil {
    ws.subscriptionCandlesChannelMap = map[string]map[string]chan SubscriptionCandles{}
  }
  if ws.subscriptionCandlesChannelMap[market] == nil {
    ws.subscriptionCandlesChannelMap[market] = map[string]chan SubscriptionCandles{}
  }
  if ws.subscriptionCandlesOptionsMap == nil {
    ws.subscriptionCandlesOptionsMap = map[string]map[string]SubscriptionCandlesObject{}
  }
  if ws.subscriptionCandlesOptionsMap[market] == nil {
    ws.subscriptionCandlesOptionsMap[market] = map[string]SubscriptionCandlesObject{}
  }
  ws.subscriptionCandlesChannelMap[market][interval] = make(chan SubscriptionCandles, 100)
  ws.subscriptionCandlesOptionsMap[market][interval] = options

  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))

  return ws.subscriptionCandlesChannelMap[market][interval]
}

func (ws *Websocket) SubscriptionTrades(market string) chan SubscriptionTrades {
  options := SubscriptionTradesBookObject{Action: "subscribe", Channels: []SubscriptionTradesBookSubObject{SubscriptionTradesBookSubObject{Name: "trades", Markets: []string{market}}}}
  if ws.subscriptionTradesChannelMap == nil {
    ws.subscriptionTradesChannelMap = map[string]chan SubscriptionTrades{}
  }
  if ws.subscriptionTradesOptionsMap == nil {
    ws.subscriptionTradesOptionsMap = map[string]SubscriptionTradesBookObject{}
  }
  ws.subscriptionTradesChannelMap[market] = make(chan SubscriptionTrades, 100)
  ws.subscriptionTradesOptionsMap[market] = options

  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.subscriptionTradesChannelMap[market]
}

func (ws *Websocket) SubscriptionBookUpdate(market string) chan SubscriptionBookUpdate {
  options := SubscriptionTradesBookObject{Action: "subscribe", Channels: []SubscriptionTradesBookSubObject{SubscriptionTradesBookSubObject{Name: "book", Markets: []string{market}}}}
  if ws.subscriptionBookUpdateChannelMap == nil {
    ws.subscriptionBookUpdateChannelMap = map[string]chan SubscriptionBookUpdate{}
  }
  if ws.subscriptionBookUpdateOptionsMap == nil {
    ws.subscriptionBookUpdateOptionsMap = map[string]SubscriptionTradesBookObject{}
  }
  ws.subscriptionBookUpdateChannelMap[market] = make(chan SubscriptionBookUpdate, 100)
  ws.subscriptionBookUpdateOptionsMap[market] = options
  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))
  return ws.subscriptionBookUpdateChannelMap[market]
}

func (ws *Websocket) SubscriptionBook(market string, options map[string]string) chan Book {
  ws.keepLocalBook = true
  options["action"] = "getBook"
  options["market"] = market
  secondOptions := SubscriptionTradesBookObject{Action: "subscribe", Channels: []SubscriptionTradesBookSubObject{SubscriptionTradesBookSubObject{Name: "book", Markets: []string{market}}}}

  if ws.subscriptionBookChannelMap == nil {
    ws.subscriptionBookChannelMap = map[string]chan Book{}
  }
  if ws.subscriptionBookOptionsFirstMap == nil {
    ws.subscriptionBookOptionsFirstMap = map[string]map[string]string{}
  }
  if ws.subscriptionBookOptionsSecondMap == nil {
    ws.subscriptionBookOptionsSecondMap = map[string]SubscriptionTradesBookObject{}
  }
  if ws.localBook.Book == nil {
    ws.localBook.Book = map[string]Book{}
  }
  ws.subscriptionBookChannelMap[market] = make(chan Book, 100)
  ws.subscriptionBookOptionsFirstMap[market] = options
  ws.subscriptionBookOptionsSecondMap[market] = secondOptions

  myMessage, _ := json.Marshal(options)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(myMessage))

  mySecondMessage, _ := json.Marshal(secondOptions)
  ws.conn.WriteMessage(websocket.TextMessage, []byte(mySecondMessage))
  return ws.subscriptionBookChannelMap[market]
}
