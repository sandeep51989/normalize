package normalize

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

type Message struct {
	BaseName  string  `json:"bn,omitempty"`
	BaseTime  int64   `json:"bt,omitempty"`
	BaseUnits string  `json:"bu,omitempty"`
	Version   int     `json:"ver"`
	Entries   []Entry `json:"e"`
}

type Entry struct {
	Name  string   `json:"n,omitempty"`
	Value *float64 `json:"v,omitempty"`
}
type OutRecord struct {
	Name     string   `json:"n,omitempty"`
	Value    *float64 `json:"v,omitempty"`
	BaseTime int64    `json:"bt,omitempty"`
}

// type Queue struct{

// 	sync.Mutex
//     Items  []interface {}

// }
// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	//var entry Message
	raw := context.GetInput("data").(Message)
	//Buffer := context.GetInput("buffer").(int)
	//var queue = make(chan Message, Buffer)
	var data []OutRecord
	for _, items := range raw.Entries {
		data = append(data, OutRecord{
			Name:     items.Name,
			Value:    items.Value,
			BaseTime: raw.BaseTime})

	}
	context.SetOutput("result", data)
	//err = json.Unmarshal([]byte(data), &entry)
	// do eval
	return true, nil
}

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-Normalizer")

// init create & register activity
func init() {
	//md := activity.NewMetadata(jsonMetadata)
	//activity.Register(NewActivity(md))
	activityLog.SetLogLevel(logger.InfoLevel)
	//act := NewActivity(getActivityMetadata())

}
