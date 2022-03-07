package tmpl

type Vouches struct {
	Url  string
	Name string
}

type TemplateModel struct {
	PrintDate      string
	OrderTitle     string
	OrderInfo      OrderInfo
	CustomInfo     CustomInfo
	GoodsInfo      GoodsInfo
	AccountInfo    AccountInfo
	DeliveryInfo   DeliveryInfo
	PaybackPlan    PaybackPlan
	PaybackReceipt PaybackReceipt
	InvoiceInfo    InvoiceInfo
	OrderAudit     OrderAudit
	AftersaleInfo  AftersaleInfo
}

type Pair struct {
	Key   string
	Value string
}

type Pairs struct {
	Pairs [2]Pair
}

type Slice struct {
	Len   int
	Pairs []Pairs
}

func (s *Slice) Append(key, value string) {
	isFull := s.Len&1 == 0
	if isFull {
		pair := Pairs{
			Pairs: [2]Pair{
				{Key: key,
					Value: value,
				},
			},
		}
		s.Pairs = append(s.Pairs, pair)
	} else {
		s.Pairs[len(s.Pairs)-1].Pairs[1].Key = key
		s.Pairs[len(s.Pairs)-1].Pairs[1].Value = value
	}
	s.Len++
}

type OrderInfo struct {
	// 是否隐藏OrderInfo
	ShowOrderInfo bool
	// 订单编号
	OrderNumber string
	// 订单状态
	OrderStatus string
	// 订单ID
	OrderId string
	// 下单时间
	CreateTime string
	// 购买人
	Buyer string
	// 收货人
	Consignee string
	// 收货电话
	Mobile string
	// 收货地址
	Address string
	// 订单类型
	OrderType string
	// 订单来源
	OrderResource string
	// 客户跟进人姓名+手机号
	CustomerFollowerInfo string
	// 订单归属人姓名+手机号
	FollowerUserInfo string
	// 订单创建人姓名+手机号
	OrderCreatorInfo string
	// 备注
	Remark string
	// 支付方式
	PayType string
	// 凭证
	VouchResource     []*Vouches
	ShowVouchResource bool
	TemplateModel     Slice
}

type customInfo struct {
	Id    string
	Key   string
	Value string
}

type CustomInfo struct {
	ShowCustomInfo bool
	List           []customInfo
	TemplateModel  Slice
}

type goodsInfo struct {
	// 商品名称
	Name string
	// 单价
	Price string
	// 数量
	Num string
	// 折扣
	Discount string
	// 折扣后总价
	PaymentPrice string
}

type GoodsInfo struct {
	ShowGoodsInfo    bool
	ShowName         bool
	ShowPrice        bool
	ShowNum          bool
	ShowDiscount     bool
	ShowPaymentPrice bool
	ShowLength       int
	Goods            []goodsInfo
}

type AccountInfo struct {
	ShowAccountInfo bool
	// 商品总价
	GoodsTotalPrice string
	// 优惠总价
	DiscountTotalPrice string
	// 会员折扣
	MemberLevelDiscountAmount string
	// 优惠巻总价gg
	CouponTotalPrice string
	// 应收款
	PaymentTotalPrice string
	TemplateModel     Slice
}

type DeliveryInfo struct {
	ShowDeliveryInfo bool
	// 发货人
	Sender string
	// 发货时间
	DeliveryTime string
	// 快递公司
	ExpressCompany string
	// 快递单号
	LogisticCode  string
	TemplateModel Slice
}

type paybackPlan struct {
	// 期次
	Phase string
	// 计划回款金额
	Amount string
	// 计划回款占比
	Rate string
	// 计划回款日期
	Date string
	// 提醒日期
	RemindDate string
	// 回款状态
	Status string
}

type PaybackPlan struct {
	ShowPaybackPlan bool
	ShowPhase       bool
	ShowAmount      bool
	ShowRate        bool
	ShowDate        bool
	ShowRemindDate  bool
	ShowStatus      bool
	ShowLength      int
	PaybackPlanList []paybackPlan
}

type paybackReceipt struct {
	// 回款时间
	Date string
	// 回款金额
	Amount string
	// 打款人
	Remitter string
	// 回款方式
	Method string
}

type PaybackReceipt struct {
	ShowPaybackReceipt bool
	ShowDate           bool
	ShowAmount         bool
	ShowRemitter       bool
	ShowMethod         bool
	ShowLength         int
	PaybackReceiptList []paybackReceipt
}

type invoiceInfo struct {
	Method string
	// 关联联系人
	CusName string
	// 联系人企业
	CorpName string
	// 开票金额
	Amount string
	// 开票明细
	Detail string
	// 开票类型
	Type string
	// 审核状态
	Status string
	// 审核时间
	Date string
	// 发票编号
	SN string
}

type InvoiceSlice struct {
	Index int
	List  Slice
}

type InvoiceInfo struct {
	ShowInvoiceInfo bool
	InvoiceList     []invoiceInfo
	TemplateModel   []InvoiceSlice
}

type orderAudit struct {
	// 审核人姓名
	Name string
	// 审核时间
	Date string
	// 审核原因
	Reason string
	// 审核状态
	Status string
}

type OrderAudit struct {
	ShowOrderAudit bool
	ShowName       bool
	ShowDate       bool
	ShowReason     bool
	ShowStatus     bool
	ShowLength     int
	OrderAudits    []orderAudit
}

type AftersaleInfo struct {
	ShowAftersaleInfo bool
	// 售后单号
	Id string
	// 申请售后时间
	CreateTime string
	// 申请人
	Creator string
	// 退款类型
	Type string
	// 退款金额
	Price string
	// 售后状态
	Status string
	// 发货状态
	SendStatus string
	// 售后凭证
	Vouchers      []*Vouches
	ShowVouchers  bool
	TemplateModel Slice
}
