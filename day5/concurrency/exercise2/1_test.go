package operators

import(
	"net/http"
	"testing"
)

const success = "\u2713"
const failure = "\u2717"

func TestAddition(t *testing.T){
	testcases := struct[] {
		x: 5,
		y: 5,
		ans: 10
	},

}

t.log("add given numbers"){
	for i, tt := range testcases{
		t.Logf("%s ")
	}
}