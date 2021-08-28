package dbManager

import (
	"reflect"
	"sort"
	"sync"
)

type ItemObjIn interface {
	GetPK() interface{}
	SetValue(v interface{}) bool
	GetUpdateTime() int64
	SetUpdateTime(t int64)
}

type InOnChangeEvent interface {
	OnDataChange(obj interface{}) bool
	RemoveItemById(id interface{}, obj interface{}, deletetime int64) bool
}

//数据对象公共结构
type CommonItemMgr struct {
	items       []ItemObjIn
	itemsDelete []ItemObjIn
	IdxMapInt64 map[int64]int
	IdxMapStr   map[string]int

	onChangeEvent map[string]InOnChangeEvent
	lock          sync.Mutex
}

func (s *CommonItemMgr) Init() {
	s.items = make([]ItemObjIn, 0)
	s.IdxMapInt64 = make(map[int64]int)
	s.IdxMapStr = make(map[string]int)
	s.itemsDelete = make([]ItemObjIn, 0)
	s.onChangeEvent = make(map[string]InOnChangeEvent, 0)
}

func (s *CommonItemMgr) SetOnChangeEvent(key string, fun InOnChangeEvent) {
	if v, ok := s.onChangeEvent[key]; ok {
		if v == fun {
			return
		}
	}
	s.onChangeEvent[key] = fun
}

func (s *CommonItemMgr) OnUpdateEnvet(data ItemObjIn) {
	for _, v := range s.onChangeEvent {
		v.OnDataChange(data)
	}
}

func (s *CommonItemMgr) RemoveItemById(id interface{}, deletetime int64) (ItemObjIn, bool) {
	var delItem ItemObjIn
	s.lock.Lock()
	defer s.lock.Unlock()
	switch reflect.ValueOf(id).Kind() {
	case reflect.Int64:
		if idx, ok := s.IdxMapInt64[id.(int64)]; ok {
			if deletetime > 0 {
				delItem = s.items[idx]
				delItem.SetUpdateTime(deletetime)
				s.itemsDelete = append(s.itemsDelete, delItem)
			}
			delete(s.IdxMapInt64, id.(int64))
			s.items = append(s.items[:idx], s.items[idx+1:]...)
			s.toSort()
			return delItem, true
		}
	case reflect.String:
		if idx, ok := s.IdxMapStr[id.(string)]; ok {
			if deletetime > 0 {
				delItem = s.items[idx]
				delItem.SetUpdateTime(deletetime)
				s.itemsDelete = append(s.itemsDelete, delItem)
			}
			delete(s.IdxMapStr, id.(string))
			s.items = append(s.items[:idx], s.items[idx+1:]...)
			s.toSort()
			return delItem, true
		}
	default:

	}
	return nil, false
}

func (s *CommonItemMgr) AddItem(item ItemObjIn, chgItem interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	id := item.GetPK()
	switch reflect.ValueOf(id).Kind() {
	case reflect.Int64:
		if idx, ok := s.IdxMapInt64[id.(int64)]; ok {
			if chgItem == nil {
				s.items[idx].SetValue(item)
				s.OnUpdateEnvet(s.items[idx])
			} else {
				s.items[idx].SetValue(chgItem)
				s.OnUpdateEnvet(s.items[idx])
			}
		} else {
			if chgItem != nil {
				item.SetValue(chgItem)
			}
			s.items = append(s.items, item)
			s.toSort()
			s.OnUpdateEnvet(item)
		}

	case reflect.String:
		if idx, ok := s.IdxMapStr[id.(string)]; ok {
			if chgItem == nil {
				s.items[idx].SetValue(item)
				s.OnUpdateEnvet(s.items[idx])
			} else {
				s.items[idx].SetValue(chgItem)
				s.OnUpdateEnvet(s.items[idx])
			}
		} else {
			if chgItem != nil {
				item.SetValue(chgItem)
			}
			s.items = append(s.items, item)
			s.toSort()
			s.OnUpdateEnvet(item)
		}
	}
	return nil
}

//排序
func (s *CommonItemMgr) toSort() {
	if len(s.items) > 0 {
		switch reflect.ValueOf(s.items[0].GetPK()).Kind() {
		case reflect.Int64:
			sort.Slice(s.items, func(i, j int) bool {
				return s.items[i].GetPK().(int64) < s.items[j].GetPK().(int64)
			})
			for i := 0; i < len(s.items); i++ {
				s.IdxMapInt64[s.items[i].GetPK().(int64)] = i
			}
		case reflect.String:
			sort.Slice(s.items, func(i, j int) bool {
				return s.items[i].GetPK().(string) < s.items[j].GetPK().(string)
			})
			for i := 0; i < len(s.items); i++ {
				s.IdxMapStr[s.items[i].GetPK().(string)] = i
			}

		}
	}
}

func (s *CommonItemMgr) GetItems() []interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	List := make([]interface{}, 0, len(s.items))
	for i := 0; i < len(s.items); i++ {
		List = append(List, s.items[i])
	}
	return List
}

func (s *CommonItemMgr) GetItemById(id interface{}) interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	switch reflect.ValueOf(id).Kind() {
	case reflect.Int64:
		if idx, ok := s.IdxMapInt64[id.(int64)]; ok {
			return s.items[idx]
		} else {
			return nil
		}
	case reflect.String:
		if idx, ok := s.IdxMapStr[id.(string)]; ok {
			return s.items[idx]
		} else {
			return nil
		}
	}
	return nil
}

func (s *CommonItemMgr) GetItemsByUpdateTime(updatetime int64) []ItemObjIn {
	s.lock.Lock()
	defer s.lock.Unlock()
	List := make([]ItemObjIn, 0, len(s.items))
	for i := 0; i < len(s.items); i++ {
		//Println("GetItemsByUpdateTime",s.items[i])
		if s.items[i].GetUpdateTime() > updatetime {
			List = append(List, s.items[i])
		}
	}
	sort.Slice(List, func(i, j int) bool {
		return List[i].GetUpdateTime() < List[j].GetUpdateTime()
	})
	return List
}

func (s *CommonItemMgr) GetDeleteItemsByUpdateTime(updatetime int64) []ItemObjIn {
	s.lock.Lock()
	defer s.lock.Unlock()
	List := make([]ItemObjIn, 0, len(s.itemsDelete))
	for i := 0; i < len(s.itemsDelete); i++ {
		//fmt.Println("GetDeleteItemsByUpdateTime",s.itemsDelete[i])
		if s.itemsDelete[i].GetUpdateTime() > updatetime {
			List = append(List, s.itemsDelete[i])
		}
	}
	sort.Slice(List, func(i, j int) bool {
		return List[i].GetUpdateTime() < List[j].GetUpdateTime()
	})
	return List
}

func (s *CommonItemMgr) Length() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.items)
}
