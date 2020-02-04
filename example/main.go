package main

import (
  "encoding/json"
  "fmt"
  "github.com/bitvavo/go-bitvavo-api"
)

/*
 * This is an example utilising all functions of the GO Bitvavo API wrapper.
 * The APIKEY and APISECRET should be replaced by your own key and secret.
 * For public functions the APIKEY and SECRET can be removed.
 * Documentation: https://docs.bitvavo.com
 * Bitvavo: https://bitvavo.com
 * README: https://github.com/bitvavo/go-bitvavo-api
 */

// Use this definition to make passing optionals easier.
// e.g. bitvavo.Markets(Options{ "market": "BTC-EUR" })
type Options map[string]string

// Use this function to print a human readable version of the returned struct.
func PrettyPrint(v interface{}) (err error) {
  b, err := json.MarshalIndent(v, "", "  ")
  if err == nil {
    fmt.Println(string(b))
  }
  return
}

func main() {
  bitvavo := bitvavo.Bitvavo{
    ApiKey:       "<APIKEY>",
    ApiSecret:    "<APISECRET>",
    RestUrl:      "https://api.bitvavo.com/v2",
    WsUrl:        "wss://ws.bitvavo.com/v2/",
    AccessWindow: 10000,
    Debugging:    false}

  var limit = bitvavo.GetRemainingLimit()
  fmt.Println("The remaining rate limit is", limit)

  testREST(bitvavo)
  testWebsocket(bitvavo)
}

func testREST(bitvavo bitvavo.Bitvavo) {
  timeResponse, timeErr := bitvavo.Time()
  if timeErr != nil {
    fmt.Println(timeErr)
  } else {
    PrettyPrint(timeResponse)
  }

  // marketsResponse, marketsErr := bitvavo.Markets(map[string]string{})
  // if marketsErr != nil {
  //   fmt.Println(marketsErr)
  // } else {
  //   for _, value := range marketsResponse {
  //     PrettyPrint(value)
  //   }
  // }

  // assetsResponse, assetsErr := bitvavo.Assets(map[string]string{})
  // if assetsErr != nil {
  //   fmt.Println(assetsErr)
  // } else {
  //   for _, value := range assetsResponse {
  //     PrettyPrint(value)
  //   }
  // }

  // bookResponse, bookErr := bitvavo.Book("BTC-EUR", map[string]string{})
  // if bookErr != nil {
  //   fmt.Println(bookErr)
  // } else {
  //   PrettyPrint(bookResponse)
  // }

  // publicTradesResponse, publicTradesErr := bitvavo.PublicTrades("BTC-EUR", map[string]string{})
  // if publicTradesErr != nil {
  //   fmt.Println(publicTradesErr)
  // } else {
  //   for _, trade := range publicTradesResponse {
  //     PrettyPrint(trade)
  //   }
  // }

  // candlesResponse, candlesErr := bitvavo.Candles("BTC-EUR", "1h", map[string]string{})
  // if candlesErr != nil {
  //   fmt.Println(candlesErr)
  // } else {
  //   for _, candle := range candlesResponse {
  //     PrettyPrint(candle)
  //   }
  // }

  // tickerPriceResponse, tickerPriceErr := bitvavo.TickerPrice(map[string]string{})
  // if tickerPriceErr != nil {
  //   fmt.Println(tickerPriceErr)
  // } else {
  //   for _, price := range tickerPriceResponse {
  //     PrettyPrint(price)
  //   }
  // }

  // tickerBookResponse, tickerBookErr := bitvavo.TickerBook(map[string]string{})
  // if tickerBookErr != nil {
  //   fmt.Println(tickerBookErr)
  // } else {
  //   for _, book := range tickerBookResponse {
  //     PrettyPrint(book)
  //   }
  // }

  // ticker24hResponse, ticker24hErr := bitvavo.Ticker24h(map[string]string{})
  // if ticker24hErr != nil {
  //   fmt.Println(ticker24hErr)
  // } else {
  //   for _, ticker := range ticker24hResponse {
  //     PrettyPrint(ticker)
  //   }
  // }

  // placeOrderResponse, placeOrderErr := bitvavo.PlaceOrder(
  //   "BTC-EUR",
  //   "buy",
  //   "limit",
  //   map[string]string{"amount": "0.3", "price": "2000"})
  // if placeOrderErr != nil {
  //   fmt.Println(placeOrderErr)
  // } else {
  //   PrettyPrint(placeOrderResponse)
  // }

  // updateOrderResponse, updateOrderErr := bitvavo.UpdateOrder("BTC-EUR", "68c72b7a-2cf5-4516-8915-703a5d38c77e", map[string]string{"amount": "0.4"})
  // if updateOrderErr != nil {
  //   fmt.Println(updateOrderErr)
  // } else {
  //   PrettyPrint(updateOrderResponse)
  // }

  // getOrderResponse, getOrderErr := bitvavo.GetOrder("BTC-EUR", "68c72b7a-2cf5-4516-8915-703a5d38c77e")
  // if getOrderErr != nil {
  //   fmt.Println(getOrderErr)
  // } else {
  //   PrettyPrint(getOrderResponse)
  // }

  // cancelOrderResponse, cancelOrderErr := bitvavo.CancelOrder("BTC-EUR", "68c72b7a-2cf5-4516-8915-703a5d38c77e")
  // if cancelOrderErr != nil {
  //   fmt.Println(cancelOrderErr)
  // } else {
  //   PrettyPrint(cancelOrderResponse)
  // }

  // getOrdersResponse, getOrdersErr := bitvavo.GetOrders("BTC-EUR", map[string]string{})
  // if getOrdersErr != nil {
  //   fmt.Println(getOrdersErr)
  // } else {
  //   for _, order := range getOrdersResponse {
  //     PrettyPrint(order)
  //   }
  // }

  // cancelOrdersResponse, cancelOrdersErr := bitvavo.CancelOrders(map[string]string{"market": "BTC-EUR"})
  // if cancelOrdersErr != nil {
  //   fmt.Println(cancelOrdersErr)
  // } else {
  //   for _, order := range cancelOrdersResponse {
  //     PrettyPrint(order)
  //   }
  // }

  // ordersOpenResponse, ordersOpenErr := bitvavo.OrdersOpen(map[string]string{"market": "BTC-EUR"})
  // if ordersOpenErr != nil {
  //   fmt.Println(ordersOpenErr)
  // } else {
  //   for _, order := range ordersOpenResponse {
  //     PrettyPrint(order)
  //   }
  // }

  // tradesResponse, tradesErr := bitvavo.Trades("BTC-EUR", map[string]string{})
  // if tradesErr != nil {
  //   fmt.Println(tradesErr)
  // } else {
  //   for _, trade := range tradesResponse {
  //     PrettyPrint(trade)
  //   }
  // }

  // balanceResponse, balanceErr := bitvavo.Balance(map[string]string{})
  // if balanceErr != nil {
  //   fmt.Println(balanceErr)
  // } else {
  //   for _, balance := range balanceResponse {
  //     PrettyPrint(balance)
  //   }
  // }

  // depositAssetsResponse, depositAssetsErr := bitvavo.DepositAssets("BTC")
  // if depositAssetsErr != nil {
  //   fmt.Println(depositAssetsErr)
  // } else {
  //   PrettyPrint(depositAssetsResponse)
  // }

  // withdrawAssetsResponse, withdrawAssetsErr := bitvavo.WithdrawAssets("BTC", "1", "BitcoinAddress", map[string]string{})
  // if withdrawAssetsErr != nil {
  //   fmt.Println(withdrawAssetsErr)
  // } else {
  //   PrettyPrint(withdrawAssetsResponse)
  // }

  // depositHistoryResponse, depositHistoryErr := bitvavo.DepositHistory(map[string]string{})
  // if depositHistoryErr != nil {
  //   fmt.Println(depositHistoryErr)
  // } else {
  //   for _, deposit := range depositHistoryResponse {
  //     PrettyPrint(deposit)
  //   }
  // }

  // withdrawalHistoryResponse, withdrawalHistoryErr := bitvavo.WithdrawalHistory(map[string]string{})
  // if withdrawalHistoryErr != nil {
  //   fmt.Println(withdrawalHistoryErr)
  // } else {
  //   for _, withdrawal := range withdrawalHistoryResponse {
  //     PrettyPrint(withdrawal)
  //   }
  // }
}

func testWebsocket(bitvavo bitvavo.Bitvavo) {
  websocket, errChannel := bitvavo.NewWebsocket()

  timeChannel := websocket.Time()
  // marketsChannel := websocket.Markets(map[string]string{})
  // assetsChannel := websocket.Assets(map[string]string{})

  // bookChannel := websocket.Book("BTC-EUR", map[string]string{})
  // publicTradesChannel := websocket.PublicTrades("BTC-EUR", map[string]string{})
  // candlesChannel := websocket.Candles("LTC-EUR", "1h", map[string]string{})

  // tickerPriceChannel := websocket.TickerPrice(map[string]string{})
  // tickerBookChannel := websocket.TickerBook(map[string]string{})
  // ticker24hChannel := websocket.Ticker24h(map[string]string{})

  // placeOrderChannel := websocket.PlaceOrder("BTC-EUR", "buy", "limit", map[string]string{"amount": "0.1", "price": "2000"})
  // updateOrderChannel := websocket.UpdateOrder("BTC-EUR", "556314b8-f719-466f-b63d-bf429b724ad2", map[string]string{"amount": "0.2"})
  // getOrderChannel := websocket.GetOrder("BTC-EUR", "556314b8-f719-466f-b63d-bf429b724ad2")
  // cancelOrderChannel := websocket.CancelOrder("BTC-EUR", "556314b8-f719-466f-b63d-bf429b724ad2")
  // getOrdersChannel := websocket.GetOrders("BTC-EUR", map[string]string{})
  // cancelOrdersChannel := websocket.CancelOrders(map[string]string{"market": "BTC-EUR"})
  // ordersOpenChannel := websocket.OrdersOpen(map[string]string{})

  // tradesChannel := websocket.Trades("BTC-EUR", map[string]string{})

  // balanceChannel := websocket.Balance(map[string]string{})
  // depositAssetsChannel := websocket.DepositAssets("BTC")
  // withdrawAssetsChannel := websocket.WithdrawAssets("EUR", "50", "NL123BIM", map[string]string{})
  // depositHistoryChannel := websocket.DepositHistory(map[string]string{})
  // withdrawalHistoryChannel := websocket.WithdrawalHistory(map[string]string{})

  // subscriptionTickerChannel := websocket.SubscriptionTicker("BTC-EUR")
  // subscriptionTicker24hChannel := websocket.SubscriptionTicker24h("BTC-EUR")
  // subscriptionAccountOrderChannel, subscriptionAccountFillChannel := websocket.SubscriptionAccount("BTC-EUR")
  // subscriptionCandlesChannel := websocket.SubscriptionCandles("BTC-EUR", "1h")
  // subscriptionTradesChannel := websocket.SubscriptionTrades("BTC-EUR")
  // subscriptionBookUpdateChannel := websocket.SubscriptionBookUpdate("BTC-EUR")
  // subscriptionBookChannel := websocket.SubscriptionBook("BTC-EUR", map[string]string{})

  // Keeps program running
  for {
    select {
    case result := <-errChannel:
      fmt.Println("Error received", result)
    case result := <-timeChannel:
      PrettyPrint(result)
    // case result := <-marketsChannel:
    //   PrettyPrint(result)
    // case result := <-assetsChannel:
    //   PrettyPrint(result)
    // case result := <-bookChannel:
    //   PrettyPrint(result)
    // case result := <-publicTradesChannel:
    //   PrettyPrint(result)
    // case result := <-candlesChannel:
    //   PrettyPrint(result)
    // case result := <-tickerPriceChannel:
    //   PrettyPrint(result)
    // case result := <-tickerBookChannel:
    //   PrettyPrint(result)
    // case result := <-ticker24hChannel:
    //   PrettyPrint(result)
    // case result := <-placeOrderChannel:
    //   PrettyPrint(result)
    // case result := <-getOrderChannel:
    //   PrettyPrint(result)
    // case result := <-updateOrderChannel:
    //   PrettyPrint(result)
    // case result := <-cancelOrderChannel:
    //   PrettyPrint(result)
    // case result := <-getOrdersChannel:
    //   PrettyPrint(result)
    // case result := <-cancelOrdersChannel:
    //   PrettyPrint(result)
    // case result := <-ordersOpenChannel:
    //   PrettyPrint(result)
    // case result := <-tradesChannel:
    //   PrettyPrint(result)
    // case result := <-balanceChannel:
    //   PrettyPrint(result)
    // case result := <-depositAssetsChannel:
    //   PrettyPrint(result)
    // case result := <-withdrawAssetsChannel:
    //   PrettyPrint(result)
    // case result := <-depositHistoryChannel:
    //   PrettyPrint(result)
    // case result := <-withdrawalHistoryChannel:
    //   PrettyPrint(result)
    // case result := <-subscriptionTickerChannel:
    //   PrettyPrint(result)
    // case result := <-subscriptionTicker24hChannel:
    //     PrettyPrint(result)
    // case result := <-subscriptionAccountOrderChannel:
    //   PrettyPrint(result)
    // case result := <-subscriptionAccountFillChannel:
    //   PrettyPrint(result)
    // case result := <-subscriptionCandlesChannel:
    //   PrettyPrint(result)
    // case result := <-subscriptionTradesChannel:
    //   PrettyPrint(result)
    // case result := <-subscriptionBookUpdateChannel:
    // PrettyPrint(result)
    // case result := <-subscriptionBookChannel:
    //   PrettyPrint(result)
    }
  }

  // Once close is called on the websocket, nothing will be received until bitvavo.NewWebsocket() is called again.
  websocket.Close()
}
