package food


import (
	"gorm.io/gorm"

)


type FoodRepository struct {
	db *gorm.DB
}



func NewFoodRepository(db *gorm.DB) *FoodRepository {
	return &FoodRepository{db: db}
}


func (fd *FoodRepository) Insert(food  Food)  Food {
	fd.db.Create(&food)
	return food
}

func (fd *FoodRepository) Update(food  Food)  Food {
	fd.db.Save(&food)
	return food
}

func (fd *FoodRepository) Delete(food  Food)  Food {
	fd.db.Delete(&food)
	return food
}

func (fd *FoodRepository) Select(columns string, condition string) []Food {
	var foods []Food
	query := fd.db.Select(columns).Where(condition).Find(&foods)
	if query.Error != nil {
		return nil
	}
	return foods
}
	

func (fd *FoodRepository) ExecSql(sql string)  {
	fd.db.Exec(sql)
}


