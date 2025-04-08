package enums

type StockAction string

const (
	IncreaseStock StockAction = "inc"
	DecreaseStock StockAction = "dec"
)

func (s StockAction) ToString() string {
	switch s {
	case IncreaseStock:
		return "inc"
	case DecreaseStock:
		return "dec"
	default:
		return ""
	}
}

func (s StockAction) IsValid() bool {
	switch s {
	case IncreaseStock, DecreaseStock:
		return true
	default:
		return false
	}
}
