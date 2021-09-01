package httpMaster

import (
	"encoding/json"
	"fmt"
	"github.com/taiwer/miner/common/dbJd"
)

func WdbOprPanicBuyingList(acticon string, args map[string]interface{}, username string, ip string) (ret string, err error) {
	args = snakeField(args)
	delete(args, "updated_at")
	defer PrintOperation("WdbOprPanicBuyingList", acticon, args, username, ip, ret, err)
	switch acticon {
	case "add":
		{
			obj := &dbJd.PanicBuying{}
			SetObjValue(obj, args)
			if _, err := obj.InserTodb(nil); err == nil {
				return "0", nil
			} else {
				return err.Error(), err
			}
		}
	case "del":
		{
			if iids, ok := args["ids"]; ok {
				ids := make([]int64, 0)
				szids, _ := iids.(string)
				if err := json.Unmarshal([]byte(szids), &ids); err == nil {
					for _, v := range ids {
						if _, err := dbJd.DelPanicBuyingById(v); err != nil {
							return "1", err
						}
					}
					return fmt.Sprintf("%d", len(ids)), err
				} else {
					return "1", err
				}
			} else {
				ishaveId, id, _, _, _ := parseArgs(args, []string{"id"}, nil)
				if ishaveId {
					if _, err := dbJd.DelPanicBuyingById(int64(id)); err == nil {
						return "0", nil
					} else {
						return err.Error(), err
					}
				} else {

					return "not find id", fmt.Errorf("not find id or ids")
				}
			}
		}
	case "mod":
		{
			effectiveKey := []string{"user_name", "item_id", "item_name", "enable", "num", "limit_price", "start_at", "stop_at"}
			isHaveId, id, _, newArgs, _ := parseArgs(args, []string{"id"}, effectiveKey)
			if isHaveId {
				if _, err := dbJd.UpdatePanicBuyingById(id, newArgs); err == nil {
					return "0", nil
				} else {
					return err.Error(), err
				}
			} else {
				errInfo := fmt.Sprintf("Err not acticon:%s not have pk", acticon)
				return errInfo, fmt.Errorf(errInfo)
			}
		}
	default:
		return fmt.Sprintf("Err not acticon:%s", acticon), fmt.Errorf("Err not acticon:%s", acticon)
	}
}
