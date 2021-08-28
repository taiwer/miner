package httpMaster

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"goplot/common/dbPlotOpr"
	"reflect"
	"strconv"
	"strings"
)

func PrintOperation(surl string, acticon string, args map[string]interface{}, username string, ip string, ret string, err error) {
	zap.L().Info(fmt.Sprintf("PrintOperation surl:%s acticon:%s args:%v username:%s ip:%s ret:%s err:%v",
		surl, acticon, args, username, ip, ret, err))

}

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func snakeField(args map[string]interface{}) map[string]interface{} {
	newArgs := make(map[string]interface{})
	for k, v := range args {
		snakeK := snakeString(k)
		newArgs[snakeK] = v
	}
	return newArgs
}

func WebDataBaseOpr(surl string, acticon string, args map[string]interface{}, username string, ip string) (ret string, err error) {
	args = snakeField(args)
	delete(args, "updated_at")
	switch surl {
	case "plotPcOpr":
		{
			ret, err = plotPcOpr(acticon, args, username, ip)
		}
	case "plotDiskOpr":
		{
			ret, err = plotDiskOpr(acticon, args, username, ip)
		}
	case "plotKeyOpr":
		{
			ret, err = plotKeyOpr(acticon, args, username, ip)
		}
	case "globalList": //分配规则
		{
			ret, err = globalListOpr(acticon, args, username, ip)
		}
	case "upLoadFile": //分配规则
		{
			ret, err = upLoadFileOpr(acticon, args, username, ip)
		}
	default:
		{
			ret, err = fmt.Sprintf("default surl:%s acticon:%s", surl, acticon), fmt.Errorf("default surl")
		}
	}
	PrintOperation(surl, acticon, args, username, ip, ret, err)
	return ret, err
}
func atoi64(text string) (int64, error) {
	if val, err := strconv.Atoi(text); err == nil {
		return int64(val), err
	} else {
		return 0, err
	}
}

func parseArgs(args map[string]interface{}, idKey []string, effectiveKey []string) (isHaveId bool, keyValueInt64 int64, keyValueText string, newArgs map[string]interface{}, err error) {
	isHaveId = false
	newArgs = map[string]interface{}{}
	for k, v := range args {
		isKey := false
		for i := 0; i < len(idKey); i++ {
			if idKey[i] == k {
				val := reflect.ValueOf(v)
				switch val.Kind() {
				case reflect.String:
					keyValueText = v.(string)
					isHaveId = true
				case reflect.Float64:
					keyValueInt64 = int64(v.(float64))
					isHaveId = true
				case reflect.Uint8:
					keyValueInt64 = int64(v.(uint8))
					isHaveId = true
				case reflect.Uint16:
					keyValueInt64 = int64(v.(uint16))
					isHaveId = true
				case reflect.Uint32:
					keyValueInt64 = int64(v.(uint32))
					isHaveId = true
				case reflect.Uint64:
					keyValueInt64 = int64(v.(uint64))
					isHaveId = true
				case reflect.Int:
					keyValueInt64 = int64(v.(int))
					isHaveId = true
				case reflect.Int64:
					keyValueInt64 = v.(int64)
					isHaveId = true
				}
				isKey = true
				break
			}
		}
		if !isKey && (len(effectiveKey) == 0 || StringInList(k, effectiveKey)) {
			newArgs[k] = v
		}
	}
	return isHaveId, keyValueInt64, keyValueText, newArgs, err
}

func plotPcOpr(acticon string, args map[string]interface{}, username string, ip string) (ret string, err error) {
	switch acticon {
	case "add":
		{
			obj := &dbPlotOpr.Node{}
			SetObjValue(obj, args)
			if _, err := obj.InserTodb(nil); err == nil {
				return "0", nil
			} else {
				return err.Error(), err
			}
		}
	case "del":
		{
			if _, ok := args["ids"]; ok {
				return "1", fmt.Errorf("cat use del ids handle")
			} else {
				ishaveId, _, name, _, _ := parseArgs(args, []string{"name"}, nil)
				if ishaveId {
					if _, err := dbPlotOpr.DelNodeByName(name); err == nil {
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
			effectiveKey := []string{"name", "descs", "ip", "user_name", "group_id", "online", "version", "start_at", "update_proxy_version", "update_proxy_status", "update_proxy_err_info",
				"off", "plot_size", "plot_thread", "plot_bucket", "plot_memory", "plot_job_tmp_size", "plot_interval", "plot_dst_work_num", "plot_max_job", "plot_farmer_public_key", "plot_pool_public_key"}
			isHaveId, _, name, newArgs, _ := parseArgs(args, []string{"name"}, effectiveKey)
			if isHaveId {
				if _, err := dbPlotOpr.UpdateNodeByName(name, newArgs); err == nil {
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

func plotDiskOpr(acticon string, args map[string]interface{}, username string, ip string) (ret string, err error) {
	switch acticon {
	case "add":
		{
			obj := &dbPlotOpr.PlotDisk{}
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
						if _, err := dbPlotOpr.DelPlotDiskById(v); err != nil {
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
					if _, err := dbPlotOpr.DelPlotDiskById(int64(id)); err == nil {
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
			effectiveKey := []string{"priority", "disk_type", "plot_off", "plot_size", "plot_thread", "plot_bucket", "plot_memory",
				"plot_job_tmp_size", "plot_interval", "plot_max_job", "plot_farmer_public_key", "plot_pool_public_key",
			}
			isHaveId, id, _, newArgs, _ := parseArgs(args, []string{"id"}, effectiveKey)
			if isHaveId {
				if _, err := dbPlotOpr.UpdatePlotDiskById(id, newArgs); err == nil {
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

func plotKeyOpr(acticon string, args map[string]interface{}, username string, ip string) (ret string, err error) {
	switch acticon {
	case "add":
		{
			obj := &dbPlotOpr.PlotKey{}
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
						if _, err := dbPlotOpr.DelPlotKeyById(v); err != nil {
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
					if _, err := dbPlotOpr.DelPlotKeyById(int64(id)); err == nil {
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
			effectiveKey := []string{"priority", "disk_type", "plot_off", "plot_size", "plot_thread", "plot_bucket", "plot_memory",
				"plot_job_tmp_size", "plot_interval", "plot_max_job", "plot_farmer_public_key", "plot_pool_public_key",
			}
			isHaveId, id, _, newArgs, _ := parseArgs(args, []string{"id"}, effectiveKey)
			if isHaveId {
				if _, err := dbPlotOpr.UpdatePlotKeyById(id, newArgs); err == nil {
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

func globalListOpr(acticon string, args map[string]interface{}, username string, ip string) (ret string, err error) {
	switch acticon {
	case "add":
		{
			obj := &dbPlotOpr.Global{}
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
						if _, err := dbPlotOpr.DelGlobalById(v); err != nil {
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
					if _, err := dbPlotOpr.DelGlobalById(int64(id)); err == nil {
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
			effectiveKey := []string{"name", "val", "text", "descs"}
			isHaveId, _, name, newArgs, _ := parseArgs(args, []string{"name"}, effectiveKey)
			if isHaveId {
				if _, err := dbPlotOpr.UpdateGlobalByName(name, newArgs); err == nil {
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

func upLoadFileOpr(acticon string, args map[string]interface{}, username string, ip string) (ret string, err error) {
	switch acticon {
	case "add":
		{
			obj := &dbPlotOpr.UpLoadFile{}
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
						if _, err := dbPlotOpr.DelUpLoadFileById(v); err != nil {
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
					if _, err := dbPlotOpr.DelUpLoadFileById(int64(id)); err == nil {
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
			effectiveKey := []string{"file_name", "file_class", "version", "descs"}
			isHaveId, id, _, newArgs, _ := parseArgs(args, []string{"id"}, effectiveKey)
			if isHaveId {
				if _, err := dbPlotOpr.UpdateUpLoadFileById(id, newArgs); err == nil {
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
