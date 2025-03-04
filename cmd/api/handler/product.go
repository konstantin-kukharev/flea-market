package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type GetProducts struct{}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description,omitempty"`
}

type PhotoSet map[int][]string
type ProductListResult struct {
	Products  []Product `json:"products"`
	Photo     PhotoSet  `json:"photos"`
	Paginator `json:"paginator"`
}

type Paginator struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
	Page   int `json:"page"`
}

func (s *GetProducts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	limit := 10
	offset := 0

	off := r.URL.Query().Get("offset")
	if off != "" {
		var err error
		offset, err = strconv.Atoi(off)
		if err != nil {
			http.Error(w, "Invalid offset value", http.StatusBadRequest)
			return
		}
	}

	products := []Product{
		{
			ID:          1,
			Name:        "Сумочка",
			Price:       "1800 р",
			Description: "Отличное состояние",
		},
		{
			ID:          2,
			Name:        "Бензопила",
			Price:       "12000",
			Description: "Бензопила Хусварма 142",
		},
		{
			ID:    3,
			Name:  "Покрышки",
			Price: "100000 р",
			Description: `Продаю canon 5d mark4 , куплен в фото нн, август 2021 (пробег 50тыс). Состояние идеальное. 
				Также обьектив canon 85мм 1,2 usm, тоже отличное состояние -80 тр. Зарядники, коробки, чехлы, чеки,
				флешки, рюкзак все в комплекте. Все вопросы, фото в личку`,
		},
	}

	photos := PhotoSet{}
	//mp := "https://i.imgur.com/xdbHo4E.png"
	photos[1] = []string{
		"https://downloader.disk.yandex.ru/preview/48712f0187c2abac7c31c1e4ca1b9210c0a456d5af0790b159f0284a94c268fa/67c7883d/MsDK6LaOX69QlmM7KZJ5kQPbv4PvEko7VB6pLGdUmOcWPajzOvPdNYN_ojLTL0rWxsZkcKegrpc2AQ7dVcHRtw%3D%3D?uid=0&filename=11.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=0&tknv=v2&size=2048x2048",
		"https://downloader.disk.yandex.ru/preview/4c6d78f27f246d7c8a266e97e7dac45eb6024810fa2f86a396dad62f20467945/67c7882c/MnlL5vtHe1YY0PTsGCa26iW6HpTCA4V_iL0Chl4Doky9YlMY_M-HYINNm5eybPhh3P1r_EWR2eTD3T-JiCb8tg%3D%3D?uid=0&filename=12.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=0&tknv=v2&size=2048x2048",
		"https://downloader.disk.yandex.ru/preview/65f40d8124f5eb8ec7a0127b0a9f398249514b78a355a9a94d5258d297d820f6/67c7881a/xadFcU1c2Pu4iJeo8BIAv7s_FA3VCeC4l_-9l1Y7SE2B0EV6FScrF6xTCtk4kAbDxeAtKDwsOOuxrpqe-uaLFQ%3D%3D?uid=0&filename=13.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=0&tknv=v2&size=2048x2048",
		"https://downloader.disk.yandex.ru/preview/4a7392ea091df9b1a2d5bf4685c45cc04ff56e465e702a9aadd1f64d378245ee/67c78801/QdpbT5cD9AP04NIiOuQ5Y_-G1V3rvHYaUo6qXJxT_DKU_KqDRuz3hny1jILM1GHm4SDERppzT2rFFzt7NjsKsQ%3D%3D?uid=0&filename=14.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=0&tknv=v2&size=2048x2048",
	}
	photos[2] = []string{
		"https://downloader.disk.yandex.ru/preview/4f0db0fd5340d9a10a9b99559a2b6d77c2944c74860f98f9b928e9aa69d6811f/67c7885f/UjRtFd06dFkVqcGFTuMXt6KytUEEIDJF60gwBBe0_uPBAXTPCAduS45cJrzQxvB21s6fIsACulPUm9EwAqQrzw%3D%3D?uid=0&filename=21.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=0&tknv=v2&size=2048x2048",
		"https://downloader.disk.yandex.ru/preview/29e54094bffd3c3fc18b922cf6d18d48a4dc663369d6e6aa9885845cec110cbc/67c78878/H77UmI6lQDfvJMXj-Uk0WqKytUEEIDJF60gwBBe0_uO8HS18aRh43uJ4CAd-1T9kej83jOF5nspyuhvsUB9mzg%3D%3D?uid=0&filename=22.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=0&tknv=v2&size=2048x2048",
		"https://downloader.disk.yandex.ru/preview/56b64a7f32003586d0fa7e14c6bcde5ee3ed96f3805db07a32e992ef2539cfed/67c7888c/agGl7FsmrZi_rLwSr_YZE7s_FA3VCeC4l_-9l1Y7SE3jnRa7jO8y7whlarUYm3P8w1f0HqEG6KSZhfagqwHxuQ%3D%3D?uid=0&filename=23.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=0&tknv=v2&size=2048x2048",
	}

	Res := &ProductListResult{
		Paginator: Paginator{
			Limit:  limit,
			Offset: offset,
			Total:  2,
			Page:   1,
		},
		Photo:    photos,
		Products: products,
	}

	result, err := json.Marshal(Res)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}

func NewGetProducts() *GetProducts {
	return &GetProducts{}
}
