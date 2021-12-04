package models

type ICompoundGraphic interface {
	Graphic
	Add(child Graphic)
	Remove(child Graphic)
}

type CompoundGraphic struct {
	children []Graphic
}

func (c *CompoundGraphic) Add(child Graphic) {
	c.children = append(c.children, child)
}

func (c *CompoundGraphic) Remove(child Graphic) {

}

func (c *CompoundGraphic) Move(dx, dy int) {
	for _, child := range c.children {
		child.Move(dx, dy)
	}
}

func (c *CompoundGraphic) Draw() {
	for _, child := range c.children {
		child.Draw()
	}
}

func NewCompoundGraphic() ICompoundGraphic {
	return &CompoundGraphic{}
}
