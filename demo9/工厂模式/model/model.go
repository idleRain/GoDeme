package model

// 这个构造体是小写开头的，因此外部无法直接引用
type student struct {
	name string
	score float64
}

// 通过工厂模式解决这个问题
func NewStudent(name string, score float64) *student {
	// 返回一个结构体的地址
	return &student{
		name: name,
		score: score,
	}
}

// 工厂模式处理构造体属性无法被外部包访问的问题
func (s *student) GetScore() float64 {
	return s.score
}
