package enums

type StockAction string

const (
	IncreaseStock StockAction = "INC"
	DecreaseStock StockAction = "DESC"
)

func (s StockAction) ToString() string {
	switch s {
	case IncreaseStock:
		return "INC"
	case DecreaseStock:
		return "DESC"
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
