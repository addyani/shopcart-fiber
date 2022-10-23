package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	Id         int     `form:"id" json:"id" validate:"required"`
	Total      float32 `form:"total" json:"total" validate:"required"`
	UserIdCart uint    `gorm:"foreignKey:UserIdCart"`
	//ProductIdCart uint       `gorm:"many2many:cart_product;foreignKey:ProductIdCart"`
	//Products *[]Product `gorm:"many2many:cart_product;foreignKey:CartIdProduct"`
	//Products      *[]Product `gorm:"many2many:cart_products;foreignKey:ProductIdCart;joinForeignKey:CartReferID;References:CartIdProduct;joinReferences:ProductRefer"`
	//ProductIdCart uint       `gorm:"index:,unique"`
	Products []*Product `gorm:"many2many:CartProduct;"`
}

// type carts_products struct {
// 	gorm.Model
// 	cart_id    uint    `gorm:"foreignKey:cart_id"`
// 	product_id uint    `gorm:"foreignKey:product_id"`
// 	Jumlah     int     `form:"jumlah" json:"jumlah" validate:"required"`
// 	Harga      float32 `form:"harga" json:"harga" validate:"required"`
// }

type CartProduct struct {
	// cart_id    uint `gorm:"foreignKey:cart_id"`
	// product_id uint `gorm:"foreignKey:product_id"`
	IdForCart    int
	IdForProduct int
	Jumlah       int     `form:"jumlah" json:"jumlah" validate:"required"`
	Harga        float32 `form:"harga" json:"harga" validate:"required"`
}

// func ReadCart(db *gorm.DB, cart *[]Cart) (err error) {
// 	err = db.Find(cart).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// CRUD
func CreateCart(db *gorm.DB, newCart *Cart) (err error) {
	err = db.Create(newCart).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadCartById(db *gorm.DB, cart *Cart, id int) (err error) {
	err = db.Where("user_id_cart=?", id).First(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertProductToCart(db *gorm.DB, insertedCart *Cart) (err error) {
	err = db.Save(insertedCart).Error
	if err != nil {
		return err
	}
	return nil
}

func FindCartProduct(db *gorm.DB, findCart *CartProduct, cart uint, prod uint) (err error) {
	err = db.Where("cart_id=?", cart).Where("product_id=?", prod).Find(findCart).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCart(db *gorm.DB, updateCart *CartProduct, cart uint, prod uint) (err error) {
	updateCart.IdForCart = int(cart)
	updateCart.IdForProduct = int(prod)
	err = db.Where("cart_id=?", cart).Where("product_id=?", prod).Save(updateCart).Error
	if err != nil {
		return err
	}
	return nil
}

func FindCart(db *gorm.DB, findCart *[]CartProduct, cart uint) (err error) {
	err = db.Where("cart_id=?", cart).Find(findCart).Error
	if err != nil {
		return err
	}
	return nil
}
