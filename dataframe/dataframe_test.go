package dataframe

import (
	"fmt"
	"github.com/xinzf/dataframe/series"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	values := make([]interface{}, 0)
	{
		values = append(values, []interface{}{1, 2, 3})
		//values = append(values, []interface{}{"11", "22", "33"})
	}
	df := New(series.New(values, series.Array, "shuzhu"))
	fmt.Println(df.Filter(F{
		Colidx:     0,
		Colname:    "shuzhu",
		Comparator: series.CompFunc,
		Comparando: func(el series.Element) bool {
			//vals:=[]interface{}{1,2,3}
			//var flag = true
			//for _,v:=range vals{
			//	el.
			//}
			log.Println("el", el.Val())
			return true
		},
	}))
}
