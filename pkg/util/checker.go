package util

// 構造体に期待する形で値が存在するか確認します。
// targetは構造体に取得した値を指定します。
// objはモデルで指定しているエンテティの構造体を指定します。
func IsEmpty(target interface{}, obj interface{}) bool {
	return target == (obj)
}
