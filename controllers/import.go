package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"

	"github.com/josimar-jr/go_mongo_expenses_sample/models"
)

// ImportController operations for Import
type ImportController struct {
	beego.Controller
}

// MovJSON...
type MovJSON struct {
	Date        string // time.Time
	Title       string
	Category    string
	Value       float64
	Description string
	Account     string
}

// URLMapping ...
func (c *ImportController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Import
// @Param	body		body 	models.Import	true		"body for Import content"
// @Success 201 {object} models.Import
// @Failure 403 body is empty
// @router / [post]
func (c *ImportController) Post() {
	var strJSON string
	var userProcess models.User
	var accountProcess models.Account
	var categProcess models.Category
	var value float64
	var movType string

	start := time.Now()

	colDate := 0
	colTitle := 1
	colCategory := 2
	colValue := 3
	colDescription := 4
	colAccount := 5
	first := true

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	// Openning the xlsx file
	xlsx, err := excelize.OpenFile("./tests/assets/__modelo.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Dados")
	strJSON += `[`
	for _, row := range rows {
		if row[0] == "" || row[0] == "Data" {
			continue
		} else {
			// preenche o json com os dados
			if first {
				strJSON += `{`
				first = false
			} else {
				strJSON += `,{`
			}

			strJSON += `"Date": "` + row[colDate] + `",`
			strJSON += `"Title": "` + row[colTitle] + `",`
			strJSON += `"Category": "` + row[colCategory] + `",`
			strJSON += `"Value": ` + row[colValue] + `,`
			strJSON += `"Description": "` + row[colDescription] + `",`
			strJSON += `"Account": "` + row[colAccount] + `"`
			strJSON += `}`
		}
	}
	strJSON += `]`

	jsonData := []MovJSON{}
	err = json.Unmarshal([]byte(strJSON), &jsonData)
	if err != nil {
		panic(err)
	}

	// inserting User
	userProcess, err = models.GetUserByEmail("userx@email.com")
	if err != nil {
		user := new(models.User)
		user.Email = "userx@email.com"
		user.FirstName = "User x"

		userProcess, _ = models.CreateUser(*user)
	}
	// fmt.Print("userProcess ")
	// fmt.Println(userProcess)

	for _, z := range jsonData {

		// inserting Account
		accountProcess, err = models.GetAccountByDescription("Conta Inicial") // z.Account
		if err != nil {
			acc := new(models.Account)
			acc.User = userProcess.Id
			acc.Account = "Conta Inicial" // z.Account

			accountProcess, _ = models.CreateAccount(*acc)
		}
		// fmt.Print("accountProcess ")
		// fmt.Println(accountProcess)

		// inserting Category
		categProcess, err = models.GetCategoryByDescription(z.Category) // "Categoria Inicial"
		if err != nil {
			cat := new(models.Category)
			cat.User = userProcess.Id
			cat.Description = z.Category // "Categoria Inicial"

			categProcess, _ = models.CreateCategory(*cat)
		}
		// fmt.Print("categProcess ")
		// fmt.Println(categProcess)

		// inserting Movement
		value = z.Value
		if value > 0 {
			movType = "R"
		} else {
			movType = "D"
		}
		movement := new(models.Movement)
		movement.User = userProcess.Id
		movement.AccountFrom = accountProcess.Id
		movement.Category = categProcess.Id
		movement.MovementDate = convStr2D(z.Date)
		movement.Title = z.Title
		movement.Value = value
		movement.MovementType = movType
		movement.Description = z.Description

		models.CreateMovement(*movement)
	}
	elapsed := time.Since(start)
	log.Printf("Import process took %s", elapsed)
	c.Ctx.Output.Body([]byte("import finished"))
}

func convStr2D(date string) (dateTime time.Time) {
	splits := strings.Split(date, "-")
	month, _ := strconv.Atoi(splits[0])
	day, _ := strconv.Atoi(splits[1])
	year, _ := strconv.Atoi(splits[2])
	dateTime = time.Date(2000+year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dateTime
}
