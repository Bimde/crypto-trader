#ifndef CRYPTO-TRADER_CLIENT_EXCHANGE_STOCK_EXCHANGE_H
#define CRYPTO-TRADER_CLIENT_EXCHANGE_STOCK_EXCHANGE_H

class StockExchange {
  public:
    StockExchange(/*Put deps here*/);
    static StockExchange* CreateExchange();
};

#endif