package controllers

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/go-admin/db"
	"github.com/radish-miyazaki/go-admin/models"
	"io"
	"net/http"
	"os"
	"strconv"
)

func AllOrders(c *gin.Context) {
	// ページネーション用変数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// INFO: entityにはモデルのインスタンスを代入
	data := models.Paginate(db.DB, &models.Order{}, page)

	c.JSON(http.StatusOK, data)
}

func Export(c *gin.Context) {
	fp := "./csv/order.csv"

	// CSVファイルの生成
	if err := CreateFile(fp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	// CSVファイルの出力
	f, _ := os.Open(fp)
	defer f.Close()
	c.Writer.Header().Add("Content-type", "application/octet-stream")
	io.Copy(c.Writer, f)
}

func CreateFile(fp string) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	var os []models.Order
	db.DB.Preload("OrderItems").Find(&os)

	// CSV Header
	w.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	// CSV Body
	for _, o := range os {
		data := []string{
			strconv.Itoa(int(o.ID)),
			o.FirstName + " " + o.LastName,
			o.Email,
			"",
			"",
			"",
		}

		// ファイルへの書き込み
		if err := w.Write(data); err != nil {
			return err
		}

		for _, oi := range o.OrderItems {
			data := []string{
				"",
				"",
				"",
				oi.ProductTitle,
				strconv.Itoa(int(oi.Price)),
				strconv.Itoa(int(oi.Quantity)),
			}

			// ファイルへの書き込み
			if err := w.Write(data); err != nil {
				return err
			}
		}
	}
	return nil
}

type Sales struct {
	Date string `json:"date"`
	Sum  string `json:"sum"`
}

// Chart チャート出力用の日付別の売上を返す。raw sqlを実行。
func Chart(c *gin.Context) {
	var sales []Sales

	db.DB.Raw(`
		SELECT DATE_FORMAT(o.created_at, '%Y-%m-%d') AS date, sum(oi.price * oi.quantity) as sum
		FROM orders o
		INNER JOIN order_items oi ON o.id = oi.order_id
		GROUP BY date
	`).Scan(&sales)

	c.JSON(http.StatusOK, sales)
}
