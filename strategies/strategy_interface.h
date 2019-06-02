#ifndef CRYPTO-TRADER_STRATEGIES_STRATEGY_INTERFACE_H
#define CRYPTO-TRADER_STRATEGIES_STRATEGY_INTERFACE_H

#include <utility>
#include <memory>

class StrategyInterface {
  public:
    StrategyInterface(std::shared_ptr<StockExchange>);
    virtual ~StrategyInterface();

  private:
    std::shared_ptr<StockExchange> exchange_;
};

#endif