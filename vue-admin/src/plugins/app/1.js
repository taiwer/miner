Date.prototype.Format = function(fmt) { //author: meizz 
	var o = {
		"M+": this.getMonth() + 1, //月份 
		"d+": this.getDate(), //日 
		"H+": this.getHours(), //小时 
		"m+": this.getMinutes(), //分 
		"s+": this.getSeconds(), //秒 
		"q+": Math.floor((this.getMonth() + 3) / 3), //季度 
		"S": this.getMilliseconds() //毫秒 
	};
	if(/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
	for(var k in o)
		if(new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
	return fmt;
}

function formatDate(time) {
	var date = new Date(time);

	var year = date.getFullYear(),
		month = date.getMonth() + 1, //月份是从0开始的
		day = date.getDate(),
		hour = date.getHours(),
		min = date.getMinutes(),
		sec = date.getSeconds();
	var newTime = year + '-' +
		(month < 10 ? '0' + month : month) + '-' +
		(day < 10 ? '0' + day : day) + ' ' +
		(hour < 10 ? '0' + hour : hour) + ':' +
		(min < 10 ? '0' + min : min) + ':' +
		(sec < 10 ? '0' + sec : sec);

	return newTime;
}

function add0(m) {
	return m < 10 ? '0' + m : m
}

function format(shijianchuo) {
	//shijianchuo是整数，否则要parseInt转换
	var time = new Date(shijianchuo * 1000);
	var y = time.getFullYear();
	var m = time.getMonth() + 1;
	var d = time.getDate();
	var h = time.getHours();
	var mm = time.getMinutes();
	var s = time.getSeconds();
	return y + '-' + add0(m) + '-' + add0(d) + ' ' + add0(h) + ':' + add0(mm) + ':' + add0(s);
}

function formatTime(shijianchuo) {
	//shijianchuo是整数，否则要parseInt转换
	var time = new Date(shijianchuo * 1000);
	var y = time.getFullYear();
	var m = time.getMonth() + 1;
	var d = time.getDate();
	var h = time.getHours();
	var mm = time.getMinutes();
	var s = time.getSeconds();
	return add0(h) + ':' + add0(mm) + ':' + add0(s);
}

function formatUtc(value){
	timeinterval = value;
	//计算出相差天数
	var days = parseInt(timeinterval / (24 * 3600 ));
	//计算出小时数
	var leave1 = timeinterval % (24 * 3600 ); //计算天数后剩余的毫秒数
	var hours = parseInt(leave1 / (3600));
	//计算相差分钟数
	var leave2 = leave1 % (3600 ); //计算小时数后剩余的毫秒数
	var minutes = parseInt(leave2 / (60 ));
	var seconds = leave2 % (60);        //计算小时数后剩余的毫秒数
	txt = "";
	if(days > 0) {
		txt += days + "天"
	}
	if(hours > 0) {
		txt += hours + "时"
	}
	if (minutes > 0) {
		txt += minutes + "分"
	}
	txt += seconds + "秒"
	ret = "";
	return txt;
}

function GetTimeUtc(value) {
    var value2 = value.replace("+08:00", 'Z');
    t = new Date(value2);
    timeinterval = t.getTime();//转化为时间戳毫秒数;
    return timeinterval/1000;
}

function formatTimeAgo(value) {
	var t = new Date(); //你已知的时间
	var t_s = t.getTime(); //转化为时间戳毫秒数
	//t.setTime(t_s - 1000 * 60 * 60 * 8); //设置新时间比旧时间多一分钟
	var value2 = value.replace("+08:00", 'Z');
	oldt = new Date(value2);
	timeinterval = t.getTime() - oldt.getTime();
	return formatUtc(timeinterval/1000);
}

function check_ip(objvalue) {
	var b = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
	if(!b.test(objvalue)) {
		return false;
	} else {
		return true;
	}
}

function PostUp(url, variable) {
	$.ajaxSetup({
		async: false
	});
	var ret = new Object;
	$.post(url, variable, function(data) {
		if(data.status == true) {
			ret.result = "";
		} else {
			ret.result = data.info;
		}
	}).error(function(err) {
		ret.result = "异常";
	});
	return ret.result;
}

//初始化fileinput
var FileInput = function() {
	var oFile = new Object();

	//初始化fileinput控件（第一次初始化）
	oFile.Init = function(ctrlName, uploadUrl) {
		var control = $('#' + ctrlName);

		//初始化上传控件的样式
		control.fileinput({
			language: 'zh', //设置语言
			uploadUrl: uploadUrl, //上传的地址
			//allowedFileExtensions: ['jpg', 'gif', 'png'],//接收的文件后缀
			showUpload: true, //是否显示上传按钮
			//showCaption: false,//是否显示标题
			browseClass: "btn btn-primary", //按钮样式 
			//dropZoneEnabled: false,//是否显示拖拽区域
			//minImageWidth: 50, //图片的最小宽度
			//minImageHeight: 50,//图片的最小高度
			//maxImageWidth: 1000,//图片的最大宽度
			//maxImageHeight: 1000,//图片的最大高度
			//maxFileSize: 0,//单位为kb，如果为0表示不限制文件大小
			//minFileCount: 0,
			maxFileCount: 10, //表示允许同时上传的最大文件个数
			enctype: 'multipart/form-data',
			//validateInitialCount:true,
			previewFileIcon: "<i class='glyphicon glyphicon-king'></i>",
			msgFilesTooMany: "选择上传的文件数量({n}) 超过允许的最大数值{m}！",
		});

		//导入文件上传完成之后的事件
		$('#file-zh').on('fileuploaded', function(event, data, previewId, index) {
			var form = data.form,
				files = data.files,
				extra = data.extra,
				response = data.response,
				reader = data.reader;
			console.log(response); //打印出返回的json
			console.log(response.paths); //打印出路径
		});
	}
	return oFile;
};

function conver(limit) {
	var size = "";
	if(limit < 0.1 * 1024) { //如果小于0.1KB转化成B
		size = limit.toFixed(2) + "B";
	} else if(limit < 0.1 * 1024 * 1024) { //如果小于0.1MB转化成KB
		size = (limit / 1024).toFixed(2) + "KB";
	} else if(limit < 0.1 * 1024 * 1024 * 1024) { //如果小于0.1GB转化成MB
		size = (limit / (1024 * 1024)).toFixed(2) + "MB";
	} else { //其他转化成GB
		size = (limit / (1024 * 1024 * 1024)).toFixed(2) + "GB";
	}

	var sizestr = size + "";
	var len = sizestr.indexOf("\.");
	var dec = sizestr.substr(len + 1, 2);
	if(dec == "00") { //当小数点后为00时 去掉小数部分
		return sizestr.substring(0, len) + sizestr.substr(len + 3, 2);
	}
	return sizestr;
}

function getQueryString(name) { 
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i"); 
        var r = window.location.search.substr(1).match(reg); 
        if (r != null) {
        		//ret = unescape(r[2]);
        		ret = decodeURI(r[2]);
        		return ret;
        }
        return null; 
    }

function delete_table_select_row(btn, table_list){
    var slt = table_list.bootstrapTable("getSelections");
    var slt_ids = $.map(slt, function (row){ return row.Id; }); // 提取选中行的id
    var variable= new Object;
    variable.ids = JSON.stringify(slt_ids);
    $.post(btn.attr('url'), variable, function(data) {
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


function put_mode_info(btn) {
    var modal_id = btn.parents('.modal').attr('id');
    var form = $('#' + modal_id).find('.modal-body').find('form');
    var variable = form.serialize();
    $.post(btn.attr('putkey'), variable, function(data) {
        if(data.status == true) {
            //alert('成功');
            $('#' + modal_id).modal('hide')
            window.location.reload()
        } else {
            alert(data.info);
        }
    }).error(function(err) {
        alert("异常");
    });
    return false
}

function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    var r = window.location.search.substr(1).match(reg);  //匹配目标参数
    if (r != null) return unescape(r[2]); return null; //返回参数值
}

function get_data(url,variable) {
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