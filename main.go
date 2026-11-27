package main

import (
	"financialExchange/storage"
	"financialExchange/enum"
	"fmt"
	"log/slog"
	"os"
	"uuid"
	"time"
)

func main(){
	options := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	logger := slog.New(handler)

	orderBook := storage.NewIndexedStorage[uuid.UUID, enum.Order](10)

	uuidOrder1 := uuid.NewV7()
	uuidOrder2 := uuid.NewV7()

	now := time.Now().UTC()

	orderBuy := 	enum.Order{
						ID: uuidOrder1,
						TraderID: uuid.NewV4(),
						Side: enum.SideBuy,
						Type: enum.OrderTypeLimit,
						Symbo: "USD",
						Price: 1000.00,
						Quantity: 2.00,
						Filled: 0.00,
						Status: enum.OrderStatusNew,
						CreatedAt: now,
						UpdatedAt: now,
							}

	orderSell := 	enum.Order{
						ID: uuidOrder2,
						TraderID: uuid.NewV4(),
						Side: enum.SideSell,
						Type: enum.OrderTypeLimit,
						Symbo: "USD",
						Price: 1000.00,
						Quantity: 1.00,
						Filled: 0.00,
						Status: enum.OrderStatusNew,
						CreatedAt: now,
						UpdatedAt: now,
						}

	fmt.Scanln()

	logger.Info("Заявки созданы", "buy_order", orderBuy, "sell_order", orderSell)
	orderBook.Put(uuidOrder1, orderBuy)
	orderBook.Put(uuidOrder2, orderSell)

	fmt.Scanln()

	matchOrders(orderBook, orderBuy.ID, orderSell.ID, logger)

	// logger.Info("Зарегистрирована сделка",
	// 		slog.String("Type", orderBuy.Type.String()),
	// 		slog.Float64("Value:", orderBuy.Price),
	// 		slog.Float64("Quantity", orderBuy.Quantity),
	// 		slog.String("Symbo", orderBuy.Symbo),
	// 		slog.String("TraderID", orderBuy.TraderID.String()))
					

	fmt.Scanln()

}

func matchorders(storage storage.Store[uuid.UUID, enum.Order], buy ID, sellID uuid. UUID, logger *slog.Logger) {
	buyOrder, buyok: storage.Get(buy ID)
	sellorder, sellok: storage.Get(sellID)

	if !buyOk || !sellOk{
		logger.Error("Возникли ошибки с поиском заказов в хранилище")
		return
	}

	if buyOrder.Symbo != sellOrder.Symbo{
		logger.Warn("Сделка не может быть осуществлена в связи несовпадением валют")
		return
	}

	if buyOrder.Status == enum.OrderStatus Filled || sellorder. Status == enum.OrderStatus Filled{
		logger.Warn("Заказы уже завершены")
		return
	}

	if buyOrder. Price < sellorder.Price{
		logger.Info("Сделка не может быть совершена из-за несовпадения цен")
		return
	}

	quantityTrade:= min(buyOrder.Remaining(), sellorder.Remaining())

	updatedBuyOrder, ok:= storage.Update(buyID, modifiers.ModifierForOrder(quantityTrade))

	if !ok{
		logger.Error("Произошла ошибка при обновлении buyOrder")
	}

	updatedSell0rder, ok:= storage.Update(sellID, modifiers.ModifierForOrder (quantity Trade))

	if !ok{
		logger.Error("Произошла ошибка при обновлении sellOrder")
	}

	logger.Info("Заявки обновлены", "buy_order", updatedBuyOrder, "sell_order", updatedSell0rder)

	logger.Info("Сделка успешно проведена",
				slog.Float64("quantity", quantityTrade),
				slog.Any("buy_order", updatedBuyOrder),
				slog.Any("sell_order", updatedSellOrder)
	)
}









	// currencies := storage.NewStorage[string](5)
	// currencies.Add("BTC-RUB")
	// currencies.Add("ETH-RUB")
	// currencies.Add("USD-RUB")
	// currencies.Add("USD-RUB")

	// lengthOfCurrencies := currencies.Length()
	// logger.Info("Работа с валютами:",
	//  slog.Int("Количество уникальных комбинаций", lengthOfCurrencies))

	// traderStore := storage.NewStore[uuid.UUID, float64](10)
	// traderID := uuid.NewV7()

	// traderStore.Put(traderID, 50000.0004)
	// logger.Info("Новый трейдер добавлен в хранилище:",
	//  slog.String("UUID", traderID.String()),
	// slog.Float64("Начальный баланс", 50000.0004))


	// if balance, ok := traderStore.Get(traderID); ok{
	// 	logger.Info("поиск в хранилище:",
	//  	slog.String("найден трейдер с UUID", traderID.String()),
	// 	slog.Float64("его текущий баланс", balance))
	// }else{
	// 	logger.Error("элемент не найден в хранилище:",
	//  	slog.String("проверьте UUID", traderID.String()),)
	// }