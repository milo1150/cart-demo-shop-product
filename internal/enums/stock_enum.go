package enums

type StockAction string

const (
	IncreaseStock StockAction = "inc"
	DecreaseStock StockAction = "dec"
	UpdateStock   StockAction = "update"
)

func (s StockAction) ToString() string {
	switch s {
	case IncreaseStock:
		return "inc"
	case DecreaseStock:
		return "dec"
	case UpdateStock:
		return "update"
	default:
		return ""
	}
}

func (s StockAction) IsValid() bool {
	switch s {
	case IncreaseStock, DecreaseStock, UpdateStock:
		return true
	default:
		return false
	}
}
