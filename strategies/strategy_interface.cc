#include "crypto-trader/strategies/strategy_interface.h"

StrategyInterface::StrategyInterface(std::shared_ptr<StockExchange> exchange) :
	exchange_{exchange} {}

StrategyInterface::~StrategyInterface() {}