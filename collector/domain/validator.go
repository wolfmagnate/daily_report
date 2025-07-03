package domain

import "github.com/go-playground/validator/v10"

// New系統の状態として存在する
// 複数の意味的なまとまりが単一の関数の状態になりうることを示す例
var validate = validator.New()
