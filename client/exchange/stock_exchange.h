#ifndef CRYPTO-TRADER_CLIENT_EXCHANGE_STOCK_EXCHANGE_H
#define CRYPTO-TRADER_CLIENT_EXCHANGE_STOCK_EXCHANGE_H

#include <unordered_map>

class StockExchange {
  public:
    StockExchange(/*Put deps here*/);
    static StockExchange* CreateExchange();

    void runExchange();

    const Stock getStock(const string& ticker);
  private:
    std::unordered_map<string, Stock> tickerToStock;
};

#endif