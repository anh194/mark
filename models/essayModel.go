package models

import(
	"mark/entity"
	"errors"
	"golang.org/x/exp/slices"
)

var listEssay = make([]*entity.Essay, 0)

func CreateEssay(essay *entity.Essay) bool {
	if essay.ID != "" && essay.Body != "" && essay.Header != "" {
		if essayF, _ := FindEssay(essay.ID); essayF == nil {
			listEssay = append(listEssay, essay)
			return true
		}
	}
	return false
}

func UpdateEssay(newEssay *entity.Essay) bool {
	for index, essay := range listEssay {
		if essay.ID == newEssay.ID {
			listEssay[index] = newEssay
			return true
		}
	}
	return false
}

func FindEssay(id string) (*entity.Essay, error){
	for _, essay := range listEssay {
		if essay.ID == id {
			return essay, nil
		}
	}
	return nil, errors.New("Essay not found")
}

func DeleteEssay(id string) bool{
	for index, essay := range listEssay{
		if essay.ID == id {
			listEssay = slices.Delete(listEssay, index, index+1)
			return true
		}
	}
	return false
}

func GetAllEssay() []*entity.Essay {
    return listEssay
}