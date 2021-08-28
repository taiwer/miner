/**
 * @author zhixin wen <wenzhixin2010@gmail.com>
 * extensions: https://github.com/vitalets/x-editable
 */

!function ($) {

    'use strict';

    $.extend($.fn.bootstrapTable.defaults, {
        editable: true,
        onEditableInit: function () {
            return false;
        },
        onEditableSave: function (field, row, oldValue, $el) {
            return false;
        },
        onEditableShown: function (field, row, $el, editable) {
            return false;
        },
        onEditableHidden: function (field, row, $el, reason) {
            return false;
        }
    })

    $.extend($.fn.bootstrapTable.Constructor.EVENTS, {
        'editable-init.bs.table': 'onEditableInit',
        'editable-save.bs.table': 'onEditableSave',
        'editable-shown.bs.table': 'onEditableShown',
        'editable-hidden.bs.table': 'onEditableHidden'
    })

  const BootstrapTable = $.fn.bootstrapTable.Constructor,
    _initTable = BootstrapTable.prototype.initTable,
    _initBody = BootstrapTable.prototype.initBody

  BootstrapTable.prototype.initTable = function () {
        var that = this;
        _initTable.apply(this, Array.prototype.slice.apply(arguments))
        if (!this.options.editable) {
            return;
        }

        $.each(this.columns, function (i, column) {
        		if (column.switch) {
        			column.formatter = function (value, row, index) {
        				if (value==0){
        					return '<button class="mod-btn-off"  data-name="' + column.field + '"><span></span></button>';
    					}else{
    						return '<button class="mod-btn-on"  data-name="' + column.field + '"><span></span></button>';
    					}
        			}
        			return;
        		}
            if (!column.editable) {
                return;
            }

            var editableOptions = {}, editableDataMarkup = [], editableDataPrefix = 'editable-';

            var processDataOptions = function(key, value) {
              // Replace camel case with dashes.
              var dashKey = key.replace(/([A-Z])/g, function($1){return "-"+$1.toLowerCase()})
              if (dashKey.slice(0, editableDataPrefix.length) == editableDataPrefix) {
                var dataKey = dashKey.replace(editableDataPrefix, 'data-')
                editableOptions[dataKey] = value;
              }
            };

            $.each(that.options, processDataOptions)

            var _formatter = column.formatter;
            column.formatter = function (value, row, index) {
                var result = _formatter ? _formatter(value, row, index) : value;
                $.each(column, processDataOptions)

                $.each(editableOptions, function (key, value) {
                    editableDataMarkup.push(' ' + key + '="' + value + '"')
                })

                return ['<a href="javascript:void(0)"',
                    ' data-name="' + column.field + '"',
                    ' data-pk="' + row[that.options.idField] + '"',
                    ' data-value="' + result + '"',
                    editableDataMarkup.join(''),
                    '>' + '</a>'
                ].join('')
            };
        })
    };

    BootstrapTable.prototype.initBody = function () {
        var that = this;
        _initBody.apply(this, Array.prototype.slice.apply(arguments))

        if (!this.options.editable) {
            return;
        }

        $.each(this.columns, function (i, column) {
        		if (column.switch) {
        			that.$body.find('button[data-name="' + column.field + '"]').on('click', function () {
                    var data = that.getData(),
                    		index = $(this).parents('tr[data-index]').data('index'),
                        row = data[index],
                        oldValue = row[column.field];

                    row[column.field] = oldValue==0 ? 1: 0;
                    that.trigger('editable-save', column.field, row, oldValue, $(this))
                    that.resetFooter()
                    that.updateCell({index:index,field:column.field,value:row[column.field]})
                })
        			return;
        		}
            if (!column.editable) {
                return;
            }

            that.$body.find('a[data-name="' + column.field + '"]').editable(column.editable)
                .off('save').on('save', function (e, params) {
                    var data = that.getData(),
                        index = $(this).parents('tr[data-index]').data('index'),
                        row = data[index],
                        oldValue = row[column.field];

                    $(this).data('value', params.submitValue)
                    row[column.field] = params.submitValue;
                    that.trigger('editable-save', column.field, row, oldValue, $(this))
                    that.resetFooter()
                    that.updateCell({index:index,field:column.field,value:row[column.field]})
                })
            that.$body.find('a[data-name="' + column.field + '"]').editable(column.editable)
                .off('shown').on('shown', function (e, editable) {
                    var data = that.getData(),
                        index = $(this).parents('tr[data-index]').data('index'),
                        row = data[index];

                    that.trigger('editable-shown', column.field, row, $(this), editable)
                })
            that.$body.find('a[data-name="' + column.field + '"]').editable(column.editable)
                .off('hidden').on('hidden', function (e, reason) {
                    var data = that.getData(),
                        index = $(this).parents('tr[data-index]').data('index'),
                        row = data[index];

                    that.trigger('editable-hidden', column.field, row, $(this), reason)
                })
        })
        this.trigger('editable-init')
    };

}(jQuery)
