package dbManager

type UpdateLog struct {
	CommonItem
	UpdateTable string
}

func NewUpdateLog() *UpdateLog {
	return &UpdateLog{}
}

func (s *UpdateLog) GetPK() interface{} {
	return s.UpdateTable
}

func (s *UpdateLog) SetValue(obj interface{}) bool {
	ret, _ := s.CommonItem.SetValue(s, obj)
	return ret
}
