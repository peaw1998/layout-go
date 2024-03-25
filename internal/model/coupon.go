package model

type Coupon struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Value int    `json:"value"`
}
