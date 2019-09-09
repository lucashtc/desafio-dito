package timeline

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
)

//HasElem retorna true caso valor exista no slice
func HasElem(find interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(find)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

type TimelineSort []Timeline

func (t TimelineSort) Len() int           { return len(t) }
func (t TimelineSort) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TimelineSort) Less(i, j int) bool { return t[i].Timestamp < t[j].Timestamp }

// TimelineOrder ...
func TimelineOrder() error {
	req, err := http.Get("https://storage.googleapis.com/dito-questions/events.json")
	if err != nil {
		return errors.Wrap(err, "Falha ao fazer GET")
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return errors.Wrap(err, "Falha ao processar Body da requisição")
	}

	var events Events
	err = json.Unmarshal(body, &events)
	if err != nil {
		return errors.Wrap(err, "Falha ao processar json")
	}

	var timelines []Timeline

	for _, v := range events.Events {
		var timeline Timeline
		var produto Products
		timeline.Timestamp = v.Timestamp
		for _, t := range v.CustomData {
			if t.Key == "store_name" {
				timeline.StoreName = fmt.Sprint(t.Value)
			}
			if t.Key == "product_name" {
				produto.Name = fmt.Sprint(t.Value)
			}
			if t.Key == "product_price" {
				price, _ := strconv.ParseFloat(fmt.Sprint(t.Value), 64)
				produto.Price, _ = strconv.Atoi(fmt.Sprint(t.Value))
				timeline.Revenue = timeline.Revenue + price
			}
			if t.Key == "transaction_id" {
				timeline.TransactionID = fmt.Sprint(t.Value)
			}
		}
		timeline.Products = append(timeline.Products, produto)
		timelines = append(timelines, timeline)
	}

	sort.Sort(TimelineSort(timelines))

	fmt.Printf("%+v\n", timelines)
	return nil
}
