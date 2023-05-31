package main

func (p *Point) Scaleby(f float64) {
	p.X *= f
	p.Y *= f
}

func (l *Line) ScaleBy(f float64) {
	l.End.Scaleby(f)
	l.Start.Scaleby(f)
}
