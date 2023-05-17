package initialize

import (
	"crypto/tls"
	"fmt"
	"mall/common"
	"mall/global"
	"mall/models/app"
	"mall/models/web"
	"mall/service"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

// 定时任务
func Cron() {

	if !global.Config.Cron.Enable {
		return
	}

	ticker := time.NewTicker(3600 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var c service.AppCartService
		var o service.AppOrderService
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		openId := "oUT385ZLmRr6R_a9xKSfSW9SekYI"

		goodsId := make([]uint64, 3)
		global.Db.Select("id").Model(&web.Goods{}).Limit(3).Find(&goodsId)

		t1 := r.Intn(10) + 1
		t2 := r.Intn(3) + 1

		for i := 0; i < t1; i++ {
			// 模拟添加购物车
			for i := 0; i < t2; i++ {
				c.Add(app.CartAddParam{
					GoodsId:    uint(goodsId[r.Intn(3)]),
					GoodsCount: uint(r.Intn(5) + 1),
					OpenId:     openId,
				})
			}
			// 模拟提交订单
			o.Submit(app.OrderSubmitParam{
				OpenId: openId,
				Sid:    100001,
			})
		}

		// 模拟更新订单
		var oc int64

		orderId := make([]int, 0)
		global.Db.Select("id").Model(&web.Order{}).Find(&orderId).Count(&oc)
		status := r.Intn(4) + 2

		order := web.Order{
			Status:  status,
			Updated: common.NowTime(),
		}
		global.Db.Model(&order).Where("id IN ?", orderId[oc-3:oc]).Updates(order)
	}
}

func SaleWatcher() {
	ticker := time.NewTicker(600 * time.Second)
	var goods []web.Goods
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("库存监控")
			global.Db.Model(&goods).Where("status = ?", 1).Where("quantity < ?", 10).Find(&goods)
			if len(goods) > 0 {
				goodsName := ""
				for _, v := range goods {
					goodsName += v.Name + " "
				}
				smtp := global.Config.Feedback.QqSmtp
				email := global.Config.Feedback.QqEmail
				secret := global.Config.Feedback.QqEmailSecret
				m := gomail.NewMessage()
				m.SetHeader("From", email)                          // 发件人
				m.SetHeader("To", email)                            // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
				m.SetHeader("Cc", email)                            // 抄送，可以多个
				m.SetHeader("Bcc", email)                           // 暗送，可以多个
				m.SetHeader("Subject", "商城后台-问题反馈")                 // 邮件主题
				m.SetBody("text/html", "商品"+goodsName+"库存不足，请及时补充") // 邮件正文
				d := gomail.NewDialer(smtp, 25, email, secret)
				// 关闭SSL协议认证
				d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
				if err := d.DialAndSend(m); err != nil {
					fmt.Println("发送失败", err)
				}
			}

		}
	}
}
