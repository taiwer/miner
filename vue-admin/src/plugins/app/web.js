function get_data_table(url,variable) {
    //var variable = new Object;
    $.ajaxSetup({
        async: false
    });
    var ret =new Object;
    $.post(url, variable, function(result) {
        ret = result;
    }).error(function(err) {
        ret = null;
    });
    return ret;
}


function post_info(url,variable) {
    $.ajaxSetup({
        async: false
    });
    var err =false;
    $.post(url, variable, function(data) {
        if(data.status == true) {
        } else {
            alert(data.info);
            err = true;
        }
    }).error(function(err) {
        //alert("异常");
        err = true;
    });
    return err;
}

function delete_select_row(btn){
    var slt = $("#table_list").bootstrapTable("getSelections");
    var slt_ids = $.map(slt, function (row){ return row.Id; });// 提取选中行的id
    var variable= new Object;
    variable.ids = JSON.stringify(slt_ids);
    url = btn.attr('url');
    $.post(url, variable, function(data) {
        if(data.status == true) {
            window.location.reload()
        } else {
            alert(data.info);
        }
    }).error(function(err) {
        alert("异常");
    });
    return false
}

function add_select_row(btn, add_id) {
    var variable = new Object;
    variable.name_id = btn.attr("name_id");
    variable.add_id =add_id;
    url = btn.attr('url')
    if (post_info(url,variable)){
        return true;
    }else{
        window.location.reload()
    }
    return false
}

var BootStrapTableQueryParams = function(params){ // 请求服务器数据时发送的参数，可以在这里添加额外的查询参数，返回false则终止请求
    return {
        pageSize: params.limit, // 每页要显示的数据条数
        offset: params.offset, // 每页显示数据的开始行号
        sort: params.sort, // 要排序的字段
        sortOrder: params.order, // 排序规则
        search: params.search,// 搜索
    }
}

 var BootStrapTableResponseHandler = function(result) {
    //alert(JSON.stringify(result));
    var errcode = result.Errcode; //在此做了错误代码的判断
    if (errcode != 0) {
        alert("错误代码" + errcode);
        return;
    }
    //如果没有错误则返回数据，渲染表格
    if (result.Total==0){
        return {
            total: result.Total,
            rows: {}
        };
    }else{
        return {
            total: result.Total,
            rows: result.Rows
        };
    }
} //请求数据成功后，渲染表格前的方法