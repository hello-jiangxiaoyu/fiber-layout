package test

import (
	"context"
	"encoding/json"
	"fiber/pkg/utils"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/open-policy-agent/opa/rego"
)

func TestRego(t *testing.T) {
	var input any
	content, err := os.ReadFile("./input.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	if err = json.Unmarshal(content, &input); err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	ctx := context.Background()
	query, err := rego.New(rego.Query("data"), rego.Load([]string{"opa.rego"}, nil)).PrepareForEval(ctx)
	if err != nil {
		log.Fatal(err)
	}

	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(utils.StructToString(rs))
}
