package lua

import (
    "github.com/yuin/gopher-lua"
    "../../build"
    "../../cell/fish/clawn"
    "../../cell/fish/predator/shark"
)

type LUA struct {
    lua_state *lua.LState
    script_name *string
}

func NewLua(script_name *string) *LUA  {
    l := &LUA{}
    l.lua_state = lua.NewState()
    l.script_name = script_name

    return l
}


func (l *LUA) BuildAreaAndCells(b build.Builder) {
    L := l.lua_state
    if err := L.DoFile(*l.script_name); err != nil {
        panic(err)
    }

    if LArea := L.GetGlobal("area"); LArea.Type() == lua.LTTable {
        lTArea := LArea.(*lua.LTable)
        lWidth := L.GetField(lTArea, "width")
        lHeight := L.GetField(lTArea, "height")
        b.BuildArea(uint(lua.LVAsNumber(lWidth)), uint(lua.LVAsNumber(lHeight)))
    }

    if LClawn := L.GetGlobal("clawn"); LClawn.Type() == lua.LTTable {
        lTClawn := LClawn.(*lua.LTable)
        lCount := L.GetField(lTClawn, "count")
        b.BuildCell(clawn.Clawn{}, int(lua.LVAsNumber(lCount)))
        //lMaxLiveCycle := L.GetField(lTClawn, "max_live_cycle")
        //if max_live_cycle := int(lua.LVAsNumber(lMaxLiveCycle)); max_live_cycle > 0 {
            //cell.SetMaxLiveCycle(max_live_cycle, 0)
        //}
    }

    if LShark := L.GetGlobal("shark"); LShark.Type() == lua.LTTable {
        lTShark := LShark.(*lua.LTable)
        lCount := L.GetField(lTShark, "count")
        b.BuildCell(shark.Shark{}, int(lua.LVAsNumber(lCount)))
    }
}