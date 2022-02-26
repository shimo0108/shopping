package persistence

import (
	"fmt"

	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/repository"
	"gorm.io/gorm"
)

type orderPersistence struct {
	Conn *gorm.DB
}

func NewOrderPersistence(conn *gorm.DB) repository.OrderRepository {
	return &orderPersistence{Conn: conn}
}

func (rp *orderPersistence) Save(order *model.Order, ids []string) (*model.Order, error) {
	var lineFoods []*model.LineFood
	var sum int32
	var totalAmount int32

	db := rp.Conn

	if err := db.Find(&lineFoods, ids).Error; err != nil {
		return nil, err
	}

	for i := range lineFoods {
		db.Model(lineFoods[i]).Association("Food").Find(&lineFoods[i].Food)
		db.Model(lineFoods[i]).Association("Restaurant").Find(&lineFoods[i].Restaurant)

	}

	for _, lf := range lineFoods {
		sum += lf.Count
		totalAmount += lf.Food.Price * lf.Count
	}

	order.TotalPrice = totalAmount

	if err := orderCreateTransaction(db, lineFoods, order); err != nil {
		fmt.Printf("%v", err)
	}

	return order, nil
}

func orderCreateTransaction(db *gorm.DB, lineFoods []*model.LineFood, order *model.Order) error {

	fmt.Println("トランザクションを開始します")

	err := db.Transaction(func(tx *gorm.DB) error {
		// データベース操作をトランザクション内で行う
		if err := tx.Model(&lineFoods).Updates(map[string]interface{}{
			"active":   false,
			"order_id": order.ID,
		}).Error; err != nil {
			return err
		}

		// Example1と同じトランザクション内で行いたい処理
		if err := db.Create(&order).Error; err != nil {
			return err
		}

		// nil を返すとコミットされる
		return nil
	})

	if err != nil {
		fmt.Println("トランザクション内でエラーが発生しました")
		return err
	}

	fmt.Println("トランザクションが正常に終了しました")
	return nil
}
