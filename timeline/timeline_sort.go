package timeline

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

var timestampDefault = "2006-01-02T15:04:05.999999999Z07:00"

type timelineSort []Timeline

func (t timelineSort) Len() int      { return len(t) }
func (t timelineSort) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t timelineSort) Less(i, j int) bool {
	ti, _ := time.Parse(timestampDefault, t[i].Timestamp)
	tj, _ := time.Parse(timestampDefault, t[j].Timestamp)
	return ti.Unix() > tj.Unix()
}

// retorna a pasição e true caso encontre
func hasTransationID(e []Timeline, id string) (int, bool) {
	for i, v := range e {
		if v.TransactionID == id {
			return i, true
		}
	}
	return 0, false
}

// TimelineEvents ...
func timelineEvents() ([]Timeline, error) {
	req, err := http.Get("https://storage.googleapis.com/dito-questions/events.json")
	if err != nil {
		return nil, errors.Wrap(err, "Falha ao fazer GET")
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Falha ao processar Body da requisição")
	}

	var events Events
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, errors.Wrap(err, "Falha ao processar json")
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
		i, bo := hasTransationID(timelines, timeline.TransactionID)
		if !bo {
			timelines = append(timelines, timeline)
		} else {
			timelines[i].Revenue = timelines[i].Revenue + timeline.Revenue
			timelines[i].Products = append(timelines[i].Products, produto)
		}
	}

	sort.Sort(timelineSort(timelines))
	return timelines, nil
}
