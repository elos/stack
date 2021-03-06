package stack

import (
	"encoding/json"
	"log"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/class"
	"github.com/elos/models/event"
	"github.com/elos/models/object"
	"github.com/elos/models/ontology"
	"github.com/elos/models/routine"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
)

func Sandbox(s data.Store) {
	/* free sandbox at the beginning of server,
	nice to test eventual functionality */

	if s == nil {
		return
	}

	u, _ := user.CreateAttrs(s, data.AttrMap{"name": "Sandy Sandbox"})

	e, _ := event.New(s)

	e.SetID(s.NewID())

	u.SetName("Sandy Sandbox")
	e.SetName("Sandy's Party")

	err := e.SetUser(u)
	if err != nil {
		log.Fatal(err)
	}

	t, _ := task.New(s)
	t.SetID(s.NewID())
	t.SetName("Sandy's Parent Task")

	t1, _ := task.New(s)
	t2, _ := task.New(s)

	t1.SetName("Sandy's Child Task 1")
	t2.SetName("Sandy's Child Task 2")
	t1.SetID(s.NewID())
	t2.SetID(s.NewID())

	u.IncludeTask(t)
	u.IncludeTask(t1)
	u.IncludeTask(t2)

	t1.AddDependency(t)
	t2.AddDependency(t)

	s.Save(t)
	s.Save(t1)
	s.Save(t2)

	if err = s.Save(u); err != nil {
		log.Fatal(err)
	}
	if err = s.Save(e); err != nil {
		log.Fatal(err)
	}

	r, _ := routine.New(s)
	r.SetID(s.NewID())
	r.IncludeTask(t)
	r.IncludeTask(t1)
	r.IncludeTask(t2)

	log.Print("ActionCount: ", r.ActionCount())
	access := data.NewAccess(u, s)

	nextAction, _ := r.NextAction(access)
	u.SetCurrentAction(nextAction)
	nextAction.SetStartTime(time.Now())
	u.SetCurrentActionable(r)

	access.Save(r)
	access.Save(nextAction)
	access.Save(u)

	o, _ := ontology.New(s)
	o.SetID(s.NewID())
	s.Save(o)
	o.SetUser(u)
	s.Save(u)

	c, _ := class.New(s)
	c.SetName("Measurement")
	s.Save(c)

	tr := &models.Trait{
		Name: "quantity",
		Type: "string",
	}

	c.IncludeTrait(tr)
	s.Save(c)

	o.IncludeClass(c)
	s.Save(c)
	s.Save(o)

	obj, _ := object.New(s)
	log.Print(obj.SetClass(c))
	log.Print(s.Save(obj))
	log.Print(obj.SetTrait(access, "quantity", "1000 ml"))
	log.Printf("%+v", obj)
	log.Print(s.Save(obj))
	o.IncludeObject(obj)
	s.Save(obj)
	s.Save(o)

	s.Save(o)
	s.Save(c)
	s.Save(obj)

	log.Printf("User id: %s", u.ID())
	log.Printf("Event id: %s", e.ID())
	bytes, err := json.Marshal(u)
	log.Print(string(bytes))
	bytes, err = json.Marshal(e)
	log.Print(string(bytes))
	bytes, err = json.Marshal(t1)
	log.Print(string(bytes))
}
