package repository

import (
	"english/config"
	"english/models"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

func InitUser() {
	config.DB.Create(&models.User{Username: "mahdiabdolmaleki", Email: "mahdi.cpp@gmail.com", Phone: "09355512619", Avatar: "2018-10-23_13-55-58_UTC_profile_pic.jpg", Biography: "go lang programmer"})
}

func InitCategory() {
	config.DB.Create(&models.Category{Hashtag: "All"})
	config.DB.Create(&models.Category{Hashtag: "University"})
	config.DB.Create(&models.Category{Hashtag: "American English File"})
	config.DB.Create(&models.Category{Hashtag: "Oxford"})
	config.DB.Create(&models.Category{Hashtag: "Google"})
	config.DB.Create(&models.Category{Hashtag: "Youtube"})
	config.DB.Create(&models.Category{Hashtag: "Medium"})
	config.DB.Create(&models.Category{Hashtag: "Instagram"})
	config.DB.Create(&models.Category{Hashtag: "Podcast"})
	config.DB.Create(&models.Category{Hashtag: "Electronic"})
	config.DB.Create(&models.Category{Hashtag: "Finance"})
	config.DB.Create(&models.Category{Hashtag: "Programing"})
	config.DB.Create(&models.Category{Hashtag: "Word504"})
	config.DB.Create(&models.Category{Hashtag: "Movie"})
	config.DB.Create(&models.Category{Hashtag: "WestWorld"})
}

func AddWord(word models.Word) error {
	err := config.DB.Debug().Create(&word).Error
	return err
}

func EditWord(word models.Word) error {

	var data = map[string]interface{}{
		"English":  strings.ToLower(word.English),
		"Hashtags": word.Hashtags,
		"Learned":  word.Learned,
	}

	err := config.DB.Debug().Where("id", word.ID).Model(models.Word{}).Updates(data).Error
	if err != nil {
		return err
	}

	fmt.Println("Translation 1: ", word.Translations)

	t := word.Translations[0]
	err = config.DB.Where("id", t.ID).
		Updates(&models.Translation{Persians: t.Persians, Type: t.Type}).Error
	if err != nil {
		return err
	}

	//init the loc
	loc, _ := time.LoadLocation("Asia/Tehran")

	//set timezone,
	now := time.Now().In(loc)

	if len(word.Translations) > 1 {
		fmt.Println("Translation 1: ", t.Persians)
		t = word.Translations[1]
		var data = map[string]interface{}{
			"WordRefer": word.ID,
			"Type":      t.Type,
			"Persians":  t.Persians,
			"CreatedAt": now,
		}
		if t.ID == 0 { //Add New
			err = config.DB.Debug().Model(models.Translation{}).Create(data).Error
			if err != nil {
				return err
			}
		} else { //Edit
			var updateData = map[string]interface{}{
				"Type":     t.Type,
				"Persians": t.Persians,
			}
			err = config.DB.Debug().Model(models.Translation{}).Where("id", t.ID).Updates(updateData).Error
			if err != nil {
				return err
			}
		}
	}

	return err
}

func EditLearned(id string, checked bool) error {
	var update = map[string]interface{}{"learned": checked}
	return config.DB.Debug().Model(models.Word{}).Where("id", id).Updates(update).Error
}

func DeleteByWord(word models.Word) error {
	err := config.DB.Debug().Where("id", word.ID).Delete(&models.Word{}).Error
	return err
}

func EditLearn(hashtag string, learn string) error {
	var update = map[string]interface{}{
		"Learn": learn,
		"Page":  0,
	}
	err := config.DB.Model(models.Category{}).Where("hashtag", hashtag).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}
func EditOrder(hashtag string, order string) error {
	var update = map[string]interface{}{
		"Order": order,
	}
	err := config.DB.Model(models.Category{}).Where("hashtag", hashtag).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}
func EditType(hashtag string, kind string) error {
	var update = map[string]interface{}{
		"Type": kind,
		"Page": 0,
	}
	err := config.DB.Model(models.Category{}).Where("hashtag", hashtag).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}
func EditPage(hashtag string, page string) error {

	if hashtag == "All" { // Always show first page for search result
		page = "0"
	}

	var update = map[string]interface{}{
		"Page": page,
	}
	err := config.DB.Model(models.Category{}).Where("hashtag", hashtag).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategory(hashtag string) (models.Category, error) {
	var category models.Category
	err := config.DB.Debug().Where("hashtag", hashtag).Find(&category).Error
	if err != nil {
		return category, err
	}
	fmt.Println(category)
	return category, nil
}
func GetWords(hashtag string, search string) (models.EnglishEntity, error) {

	var entity models.EnglishEntity
	var category models.Category
	var words []models.Word
	var where = ""
	var count int64

	fmt.Println("search: " + search)

	err := config.DB.Debug().Where("hashtag", hashtag).Find(&category).Error
	if err != nil {
		return models.EnglishEntity{}, err
	}

	println("Category:", category.Hashtag)

	if strings.Contains(category.Learn, "learned") {
		where += "english.words.learned = true "
	} else if strings.Contains(category.Learn, "tutorial") {
		where += "english.words.learned = false"
	}

	if len(search) > 0 {
		if len(where) > 0 {
			where += " AND english.words.english like '" + search + "%'"
		} else {
			where += "english.words.english like '" + search + "%'"
		}
	}

	if hashtag != "All" {
		if len(where) > 0 {
			where += " AND  '" + hashtag + "'=ANY(english.words.hashtags)"
		} else {
			where = "'" + hashtag + "'=ANY(english.words.hashtags)"
		}
	}

	//if category.Type != "" {
	//	if len(where) > 0 {
	//		where += " AND english.translations.type = '" + category.Type + "'"
	//	} else {
	//		where = "english.translations.type = '" + category.Type + "'"
	//	}
	//}

	config.DB.Debug().
		Preload("Translations", func(db *gorm.DB) *gorm.DB { return db.Order("id ASC") }).
		Where(where).
		Offset(int(category.Page * 5)).
		Limit(5).
		Order("id " + category.Order).
		Find(&words)

	config.DB.Debug().Model(&models.Word{}).Where(where).Count(&count)

	entity.Words = words
	entity.Count = count

	println("......Count:", count)

	return entity, nil
}

func GetById(id string) (models.Word, error) {
	var word models.Word
	err := config.DB.Debug().Where("id", id).
		Preload("Translations", func(db *gorm.DB) *gorm.DB { return db.Order("id ASC") }).
		Find(&word).Error
	return word, err
}

func GetIsWordAvailable(english string) int64 {
	var word models.Word
	result := config.DB.Debug().Where("english", english).Find(&word)
	return result.RowsAffected
}
