package ItemDetails

import (
	"math"
	"testing"
)

type TaxCalulationLogicTest struct {
	i        Item
	expected float64
}

var TaxTests = []TaxCalulationLogicTest{
	TaxCalulationLogicTest{i: Item{ItemName: "Pen", ItemPrice: 45, ItemQuantity: 5, ItemType: "raw"}, expected: 253.120000},
	TaxCalulationLogicTest{i: Item{ItemName: "Copy", ItemPrice: 75, ItemQuantity: 15, ItemType: "manufactured"}, expected: 1267.310000},
	TaxCalulationLogicTest{i: Item{ItemName: "Eraser", ItemPrice: 85, ItemQuantity: 16, ItemType: "raw"}, expected: 1530.000000},
	TaxCalulationLogicTest{i: Item{ItemName: "Tool", ItemPrice: 105, ItemQuantity: 18, ItemType: "imported"}, expected: 253.120000},
	TaxCalulationLogicTest{i: Item{ItemName: "Book", ItemPrice: 85.4, ItemQuantity: 29, ItemType: "imported"}, expected: 2609.390000},
}

func TestTaxCalulationLogic(t *testing.T) {
	for _, test := range TaxTests {
		output := test.i.TaxCalulationLogic(test.i.ItemPrice * float64(test.i.ItemQuantity))
		output = math.Floor(output*100) / 100
		expected := math.Floor(test.expected*100) / 100
		if output != expected {
			t.Errorf("Output %f not equal to expected %f", output, expected)
		}
	}

}
