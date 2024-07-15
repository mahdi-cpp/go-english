package repository

import (
	"english/config"
	"english/models"
	"fmt"
)

func CreatPasswords() {
	config.DB.Create(&models.Password{Name: "twitter.com", Username: "mahdi.cpp@gmail.com", Password: "$Mahdi@3327#&", Link: "http://twitter.com/"})
	config.DB.Create(&models.Password{Name: "github.com", Username: "mahdi-cpp", Password: "$Mahdi@4172$", Link: ""})
	config.DB.Create(&models.Password{Name: "hub.docker.com", Username: "", Password: "Mahdi@1400#", Link: ""})
	config.DB.Create(&models.Password{Name: "instagram", Username: "mahdi.tinyhome", Password: "Mahdi@1234", Link: "https://www.instagram.com/"})

	config.DB.Create(&models.Password{Name: "novinhost.org", Username: "", Password: "Mahdi@1400#", Link: "https://www.novinhost.org/"})
	config.DB.Create(&models.Password{Name: "www.gate.io", Username: "", Password: "Mahdi@1366#", Link: ""})
	config.DB.Create(&models.Password{Name: "panel.kavenegar.com/", Username: "", Password: "", Link: "panel.kavenegar.com/"})
	config.DB.Create(&models.Password{Name: "panel.kavenegar.com/", Username: "", Password: "itc372669", Link: ""})
	config.DB.Create(&models.Password{Name: "kucoin TRADE PASSWORD", Username: "", Password: "332733", Link: ""})
	config.DB.Create(&models.Password{Name: "Telegram", Username: "", Password: "4172", Link: ""})
}

func Add(password models.Password) error {
	err := config.DB.Debug().Create(&password).Error
	return err
}

func GetPasswords() ([]models.Password, error) {
	var passwords []models.Password

	err := config.DB.Debug().
		Offset(int(0)).
		Limit(10).
		Order("id ASC").
		Find(&passwords).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(len(passwords))
	return passwords, nil
}
