package dto

import (
	"fmt"
)

type StateType struct {
	Default        int
	Interface      int
	Connect        int
	Relay          int
	Parking        int
	Charger        int
	StateTypeFault int
	BmsFault       int
	Reason         int
	Strategy       int
	AuxType        int
	ControlModel   int
}

func (s *StateType) SetValue(index, val int) error {
	switch index {
	case 0:
		s.Default = val
	case 1:
		s.Interface = val
	case 2:
		s.Connect = val
	case 3:
		s.Relay = val
	case 4:
		s.Parking = val
	case 5:
		s.Charger = val
	case 6:
		s.StateTypeFault = val
	case 7:
		s.BmsFault = val
	case 8:
		s.Reason = val
	case 9:
		s.Strategy = val
	case 10:
		s.AuxType = val
	case 11:
		s.ControlModel = val
	default:
		return fmt.Errorf("无效的状态索引")
	}
	return nil
}
