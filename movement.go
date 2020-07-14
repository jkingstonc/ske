package ske

import "reflect"

type SimpleMoveController struct {
	Component
	Speed     float64
	Transform *TransformComponent
}

func (m*SimpleMoveController) OnLoad(){
	m.Transform = m.Entity.GetComponent(reflect.TypeOf(&TransformComponent{})).(*TransformComponent)
}


func (m*SimpleMoveController) Update(){
	if m.Transform != nil{
		if Inputs.Button("w").Held(){
			m.Transform.Translate(V2(0,1*m.Speed*DT))
		}else if Inputs.Button("s").Held(){
			m.Transform.Translate(V2(0,-1*m.Speed*DT))
		}
		if Inputs.Button("a").Held(){
			m.Transform.Translate(V2(-1*m.Speed*DT,0))
		}else if Inputs.Button("d").Held(){
			m.Transform.Translate(V2(1*m.Speed*DT,0))
		}
	}
}