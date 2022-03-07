package tmpl

var Html = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <title>Document</title>
    <style type="text/css">
        body {
            width: 888px;
            margin: 0 auto;
            font-size: 12px;
        }

        .template-page {
            color: #262626;
            padding: 24px;
        }

        .time {
            margin-bottom: 10px;
        }

        .title {
            font-size: 24px;
            line-height: 33px;
            font-weight: bolder;
            text-align: center;
        }

        .info-title {
            font-size: 17px;
            line-height: 24px;
            margin-bottom: 10px;
            font-weight: bolder;
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
        }

        .info-box {
            margin-bottom: 20px;
        }

        table {
            border-collapse: collapse;
            border: 1px solid #e5e5e5;
            width: 100%;
        }

        table tr {
            border-bottom: 1px solid #e5e5e5;
			page-break-inside: avoid; 
        }

        table tr th,
        table tr td {
            text-align: left;
            padding: 12px;
        }

        table tr th {
            background-color: #fafafa;
        }

        table tr td span {
            display: inline-block;
            line-height: 16px;
            word-break: break-all;
            vertical-align: bottom;
        }

        /* 二等分表格 */
        .halve {
            width: 100%;
            display: flex;
            flex-wrap: wrap;
            align-content: center;
            border: 1px solid #e5e5e5;
            box-sizing: border-box;
            border-bottom: none;
        }

        .header {
            width: 100%;
            background-color: #f0f0f0;
            box-sizing: border-box;
            padding: 12px;
            font-weight: bold;
            border: 1px solid #e5e5e5;
            border-bottom: none;
            position: relative;
        }

        .list-box {
            width: 50%;
            box-sizing: border-box;
            padding: 12px;
            border-bottom: 1px solid #e5e5e5;
            position: relative;
        }

        .list-box:nth-child(odd) {
            border-right: 1px solid #e5e5e5;
        }

        .entireLine {
            width: 100%;
            border-right: none !important;
            position: relative;
        }

        .entireLine:before {
            content: '';
            position: absolute;
            top: -1px;
            left: 0;
            right: 0;
            height: 1px;
            background-color: #e5e5e5;
        }

        .list-box img {
            display: block;
            max-width: 100%;
            max-height: 453px;
            overflow: hidden;
            background-color: #f2f2f2;
            margin: 10px auto 0;
        }

        .halve-table tr td span {
            width: 396px;
        }

        .halve-table img {
            display: block;
            max-width: 100%;
            max-height: 453px;
            overflow: hidden;
            background-color: #f2f2f2;
            margin: 10px auto 0;
        }

        .halve-table tr td:nth-child(2n - 1) {
            border-right: 1px solid #e5e5e5;
        }

        .halve-table:last-child {
            margin-top: 10px;
        }

        /* 商品基础信息表格 */
        .basic tr td span {
            width: 111px;
        }
        .basic .first-child span {
            width: 276px;
            word-break: break-all;
        }

         /*单列表格*/
        .singleRow-table tr td span{
            width: 816px;
        }
        /*新二等分表格*/
        .halve-table-new tr td span {
            width: 396px;
        }
        /*三等分表格*/
        .trisection-table tr td span {
            width: 256px;
        }
        /* 四等分表格 */
        .quartering-table tr td span {
            width: 186px;
        }
        /* 五等分表格 */
        .quintile-table tr td span {
            width: 144px;
        }
        /* 六等分表格 */
        .six-equal-table tr td span {
            width: 116px;
        }
    </style>
</head>

<body>
    <div class="template-page">
        <div class="time">打印时间：{{ .PrintDate }}</div>
        <div class="title">{{ .OrderTitle }}</div>
        {{if .OrderInfo.ShowOrderInfo}}
        <div class="info-box">
            <div class="info-title">订单信息</div>
            <table class="halve-table">
                {{range .OrderInfo.TemplateModel.Pairs}}
                <tr>
                    {{range .Pairs}}
                    <td><span>{{.Key}}{{.Value}}</span></td>
                    {{end}}
                </tr>
                {{end}}
                {{if .OrderInfo.ShowVouchResource}}
                <tr>
                    <td colspan="2">
                        <span>凭证：</span>
                        {{range .OrderInfo.VouchResource}}
                        <img src={{.URL}} alt={{.Name}}>
                        {{end}}
                    </td>
                </tr>
                {{end}}
            </table>
        </div>
        {{end}}
        {{if .CustomInfo.ShowCustomInfo}}
        <div class="info-box">
            <div class="info-title">订单自定义字段</div>
            <table class="halve-table">
                {{range .CustomInfo.TemplateModel.Pairs}}
                <tr>
                    {{range .Pairs}}
                    <td><span>{{.Key}}{{.Value}}</span></td>
                    {{end}}
                </tr>
                {{end}}
            </table>
        </div>
        {{end}}
        {{if .GoodsInfo.ShowGoodsInfo}}
        <div class="info-box">
            <div class="info-title">商品基础信息</div>
            <table class="basic">
                <tr>
                    {{if .GoodsInfo.ShowName}}
                    <th><span>商品</span></th>
                    {{end}}
                    {{if .GoodsInfo.ShowPrice}}
                    <th><span>单价(元)</span></th>
                    {{end}}
                    {{if .GoodsInfo.ShowNum}}
                    <th><span>数量</span></th>
                    {{end}}
                    {{if .GoodsInfo.ShowDiscount}}
                    <th><span>折扣</span></th>
                    {{end}}
                    {{if .GoodsInfo.ShowPaymentPrice}}
                    <th><span>折后总价(元)</span></th>
                    {{end}}
                </tr>
                {{range .GoodsInfo.Goods}}
                <tr>
                    {{if $.GoodsInfo.ShowName}}
                    <td class="first-child">
                        <span>{{.Name}}</span>
                    </td>
                    {{end}}
                    {{if $.GoodsInfo.ShowPrice}}
                    <td><span>{{.Price}}</span></td>
                    {{end}}
                    {{if $.GoodsInfo.ShowNum}}
                    <td><span>{{.Num}}</span></td>
                    {{end}}
                    {{if $.GoodsInfo.ShowDiscount}}
                    <td><span>{{.Discount}}</span></td>
                    {{end}}
                    {{if $.GoodsInfo.ShowPaymentPrice}}
                    <td><span>{{.PaymentPrice}}</span></td>
                    {{end}}
                </tr>
                {{end}}
            </table>
        </div>
        {{end}}
        {{if .AccountInfo.ShowAccountInfo}}
        <div class="info-box">
            <div class="info-title">结算信息</div>
            <table class="halve-table">
                {{range .AccountInfo.TemplateModel.Pairs}}
                <tr>
                    {{range .Pairs}}
                    <td><span>{{.Key}}{{.Value}}</span></td>
                    {{end}}
                </tr>
                {{end}}
            </table>
        </div>
        {{end}}
        {{if .DeliveryInfo.ShowDeliveryInfo}}
        <div class="info-box">
            <div class="info-title">发货信息</div>
            <table class="halve-table">
                {{range .DeliveryInfo.TemplateModel.Pairs}}
                <tr>
                    {{range .Pairs}}
                    <td><span>{{.Key}}{{.Value}}</span></td>
                    {{end}}
                </tr>
                {{end}}
            </table>
        </div>
        {{end}}
        {{if .PaybackPlan.ShowPaybackPlan}}
        <div class="info-box">
            <div class="info-title">回款计划</div>
            {{if eq .PaybackPlan.ShowLength 1}}
            <table class="basic">
            {{end}}
            {{if eq .PaybackPlan.ShowLength 2}}
            <table class="halve-table-new">
            {{end}}
            {{if eq .PaybackPlan.ShowLength 3}}
            <table class="trisection-table">
            {{end}}
            {{if eq .PaybackPlan.ShowLength 4}}
            <table class="quartering-table">
            {{end}}
            {{if eq .PaybackPlan.ShowLength 5}}
            <table class="quintile-table">
            {{end}}
            {{if eq .PaybackPlan.ShowLength 6}}
            <table class="six-equal-table">
            {{end}}
                <tr>
                    {{if $.PaybackPlan.ShowPhase}}
                    <th><span>期次</span></th>
                    {{end}}
                    {{if $.PaybackPlan.ShowAmount}}
                    <th><span>计划回款金额(元)</span></th>
                    {{end}}
                    {{if $.PaybackPlan.ShowRate}}
                    <th><span>计划回款占比</span></th>
                    {{end}}
                    {{if $.PaybackPlan.ShowDate}}
                    <th><span>计划回款日期</span></th>
                    {{end}}
                    {{if $.PaybackPlan.ShowRemindDate}}
                    <th><span>提醒日期</span></th>
                    {{end}}
                    {{if $.PaybackPlan.ShowStatus}}
                    <th><span>回款状态</span></th>
                    {{end}}
                </tr>
                {{range .PaybackPlan.PaybackPlanList}}
                <tr>
                    {{if $.PaybackPlan.ShowPhase}}
                    <td><span>{{.Phase}}</span></td>
                    {{end}}
                    {{if $.PaybackPlan.ShowAmount}}
                    <td><span>{{.Amount}}</span></td>
                    {{end}}
                    {{if $.PaybackPlan.ShowRate}}
                    <td><span>{{.Rate}}</span></td>
                    {{end}}
                    {{if $.PaybackPlan.ShowDate}}
                    <td><span>{{.Date}}</span></td>
                    {{end}}
                    {{if $.PaybackPlan.ShowRemindDate}}
                    <td><span>{{.RemindDate}}</span></td>
                    {{end}}
                    {{if $.PaybackPlan.ShowStatus}}
                    <td><span>{{.Status}}</span></td>
                    {{end}}
                </tr>
                {{end}}
            </table>
        </div>
        {{end}}
        {{if .PaybackReceipt.ShowPaybackReceipt}}
        <div class="info-box">
            <div class="info-title">回款单</div>
            {{if eq .PaybackReceipt.ShowLength 1}}
            <table class="basic">
            {{end}}
            {{if eq .PaybackReceipt.ShowLength 2}}
            <table class="halve-table-new">
            {{end}}
            {{if eq .PaybackReceipt.ShowLength 3}}
            <table class="trisection-table">
            {{end}}
            {{if eq .PaybackReceipt.ShowLength 4}}
            <table class="quartering-table">
            {{end}}
            {{if eq .PaybackReceipt.ShowLength 5}}
            <table class="quintile-table">
            {{end}}
            {{if eq .PaybackReceipt.ShowLength 6}}
            <table class="six-equal-table">
            {{end}}
                <tr>
                    {{if $.PaybackReceipt.ShowDate}}
                    <th><span>回款时间</span></th>
                    {{end}}
                    {{if $.PaybackReceipt.ShowAmount}}
                    <th><span>回款金额(元)</span></th>
                    {{end}}
                    {{if $.PaybackReceipt.ShowRemitter}}
                    <th><span>打款人</span></th>
                    {{end}}
                    {{if $.PaybackReceipt.ShowMethod}}
                    <th><span>回款方式</span></th>
                    {{end}}
                </tr>
                {{range .PaybackReceipt.PaybackReceiptList}}
                <tr>
                    {{if $.PaybackReceipt.ShowDate}}
                    <td><span>{{.Date}}</span></td>
                    {{end}}
                    {{if $.PaybackReceipt.ShowAmount}}
                    <td><span>{{.Amount}}</span></td>
                    {{end}}
                    {{if $.PaybackReceipt.ShowRemitter}}
                    <td><span>{{.Remitter}}</span></td>
                    {{end}}
                    {{if $.PaybackReceipt.ShowMethod}}
                    <td><span>{{.Method}}</span></td>
                    {{end}}
                </tr>
                {{end}}
            </table>
        </div>
        {{end}}
        {{if .InvoiceInfo.ShowInvoiceInfo}}
        <div class="info-box">
            <div class="info-title">发票信息</div>
            {{range .InvoiceInfo.TemplateModel}}
            <table class="halve-table">
                <tr>
                    <th colspan="2"><span>发票 {{.Index}}</span></th>
                </tr>
                {{range .List.Pairs}}
                <tr>
                    {{range .Pairs}}
                    <td><span>{{.Key}}{{.Value}}</span></td>
                    {{end}}
                </tr>
                {{end}}
            </table>
            {{end}}
        </div>
        {{end}}
        {{if .OrderAudit.ShowOrderAudit}}
        <div class="info-box">
            <div class="info-title">订单审核</div>
            {{if eq .OrderAudit.ShowLength 1}}
            <table class="basic">
            {{end}}
            {{if eq .OrderAudit.ShowLength 2}}
            <table class="halve-table-new">
            {{end}}
            {{if eq .OrderAudit.ShowLength 3}}
            <table class="trisection-table">
            {{end}}
            {{if eq .OrderAudit.ShowLength 4}}
            <table class="quartering-table">
            {{end}}
            {{if eq .OrderAudit.ShowLength 5}}
            <table class="quintile-table">
            {{end}}
            {{if eq .OrderAudit.ShowLength 6}}
            <table class="six-equal-table">
            {{end}}
                <tr>
                    {{if $.OrderAudit.ShowName}}
                    <th><span>审核人姓名</span></th>
                    {{end}}
                    {{if $.OrderAudit.ShowDate}}
                    <th><span>审核时间</span></th>
                    {{end}}
                    {{if $.OrderAudit.ShowReason}}
                    <th><span>审核原因</span></th>
                    {{end}}
                    {{if $.OrderAudit.ShowStatus}}
                    <th><span>审核状态</span></th>
                    {{end}}
                </tr>
                {{range .OrderAudit.OrderAudits}}
                <tr>
                    {{if $.OrderAudit.ShowName}}
                    <td><span>{{.Name}}</span></td>
                    {{end}}
                    {{if $.OrderAudit.ShowDate}}
                    <td><span>{{.Date}}</span></td>
                    {{end}}
                    {{if $.OrderAudit.ShowReason}}
                    <td><span>{{.Reason}}</span></td>
                    {{end}}
                    {{if $.OrderAudit.ShowStatus}}
                    <td><span>{{.Status}}</span></td>
                    {{end}}
                </tr>
                {{end}}
            </table>
        </div>
        {{end}}
        {{if .AftersaleInfo.ShowAftersaleInfo}}
        <div class="info-box">
            <div class="info-title">售后信息</div>
            <table class="halve-table">
                {{range .AftersaleInfo.TemplateModel.Pairs}}
                <tr>
                    {{range .Pairs}}
                    <td><span>{{.Key}}{{.Value}}</span></td>
                    {{end}}
                </tr>
                {{end}}
                {{if .AftersaleInfo.ShowVouchers}}
                <tr>
                    <td colspan="2">
                        <span>售后凭证：</span>
                        {{range .AftersaleInfo.Vouchers}}
                        <img src={{.Url}} alt={{.Name}}>
                        {{end}}
                    </td>
                </tr>
                {{end}}
            </table>
        </div>
        {{end}}
    </div>
    <div style="page-break-after:always; "></div>
</body>

</html>
`
