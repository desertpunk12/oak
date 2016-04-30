package collision

import (
	//"fmt"
	"bitbucket.org/oakmoundstudio/plasticpiston/plastic/event"
	"github.com/dhconnelly/rtreego"
	"log"
)

var (
	rt *rtreego.Rtree
)

type Space struct {
	Location *rtreego.Rect
	cID      event.CID
}

func (s Space) Bounds() *rtreego.Rect {
	return s.Location
}

func Init() {
	rt = rtreego.NewTree(2, 20, 40)
}

func Clear() {
	Init()
}

func Add(sp Space) {
	rt.Insert(sp)
}

func Remove(sp Space) {
	rt.Delete(sp)
}

func UpdateSpace(x, y, w, h float64, s Space) *rtreego.Rect {
	loc := NewRect(x, y, w, h)
	Update(s, loc)
	return loc
}

func Update(s Space, loc *rtreego.Rect) {
	rt.Delete(s)
	s.Location = loc
	rt.Insert(s)
}

func Hits(sp Space) []Space {
	results := rt.SearchIntersect(sp.Bounds())
	out := make([]Space, len(results))
	for index, v := range results {
		out[index] = v.(Space)
	}
	return out
}

func NewSpace(x, y, w, h float64, cID event.CID) Space {
	rect := NewRect(x, y, w, h)
	return Space{
		rect,
		cID,
	}
}

func NewRect(x, y, w, h float64) *rtreego.Rect {
	rect, err := rtreego.NewRect(rtreego.Point{x, y}, []float64{w, h})
	if err != nil {
		log.Fatal(err)
	}
	return rect
}