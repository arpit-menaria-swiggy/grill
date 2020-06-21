package grilldynamo

import (
	"fmt"
	"reflect"

	"bitbucket.org/swigy/grill"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (grilldynamo *GrillDynamo) AssertScanCount(input *dynamodb.ScanInput, expectedCount int) grill.Assertion {
	return grill.AssertionFunc(func() error {

		output, err := grilldynamo.dynamo.Client.Scan(input)
		if err != nil {
			return err
		}

		if len(output.Items) != expectedCount {
			return fmt.Errorf("invalid number of items, got=%v, want=%v", len(output.Items), expectedCount)
		}

		return nil
	})
}

func (grilldynamo *GrillDynamo) AssertItem(input *dynamodb.GetItemInput, expected interface{}) grill.Assertion {
	return grill.AssertionFunc(func() error {
		output, err := grilldynamo.dynamo.Client.GetItem(input)
		if err != nil {
			return err
		}

		want, err := dynamodbattribute.MarshalMap(expected)
		if err != nil {
			return err
		}

		if !reflect.DeepEqual(output.Item, want) {
			return fmt.Errorf("invalid item, got=%v, want=%v", output.Item, want)
		}

		return nil
	})
}
