package forms

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type Form struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Live     bool   `json:"live"`
}

type Element struct {
	ID       int64  `json:"id"`
	FormID   int64  `json:"formID"`
	Label    string `json:"label"`
	Type     string `json:"type"`
	Position int    `json:"position"` // index
	Required bool   `json:"required"`
	Priority int    `json:"priority"`
	Search   bool   `json:"search"`
}

type Option struct {
	ID        int64  `json:"id"`
	ElementID int64  `json:"elementID"`
	Name      string `json:"name"`
	Position  int    `json:"position"` // index
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ElementCategory struct {
	ID         int64 `json:"id"`
	ElementID  int64 `json:"elementID"`
	CategoryID int64 `json:"categoryID"`
}

// TODO
func handler(ctx context.Context) error {
	sess := session.Must(session.NewSession())
	svc := ssm.New(sess)
	path := fmt.Sprintf("/icc/%s/database/", os.Getenv("APP_ENV"))
	input := ssm.GetParametersByPathInput{
		Path: &path,
	}
	out, err := svc.GetParametersByPath(&input)
	if err != nil {
		return err
	}
	params := out.Parameters
	for i := 0; i < len(params); i++ {
		fmt.Println(*params[i].Name + ": " + *params[i].Value)
	}
	// connection := database.SqlConnection{
	// 	Host: os.Getenv("DB_HOST"),
	// 	Port: os.Getenv("DB_PORT"),
	// 	User: os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// }
	// database.Connect()
	return nil
}

func main() {
	lambda.Start(handler)
}
