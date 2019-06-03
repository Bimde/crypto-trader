#include <mutex>
#include "crypto-trader/client/exchange/stock_exchange.h"

std::mutex mtx;

StockExchange::StockExchange() {}

void StockExchange::runExchange() {

	// Implement running in a separate thread, updating stock values
}

void StockExchange::getStock(const string& ticker) {
	mtx.lock();
	auto stock = tickerToStock[ticker];
	mtx.unlock();
	
	return stock;
}