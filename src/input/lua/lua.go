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

    if LClawnFactory := L.GetGlobal("clawn_factory"); LClawnFactory.Type() == lua.LTTable {
        lTClawnFactory := LClawnFactory.(*lua.LTable)
        for i := 1; i < lTClawnFactory.Len(); i++ {
            c := b.BuildCell(clawn.Clawn{})
            LClawn := lTClawnFactory.RawGetInt(i)
            if LClawn.Type() != lua.LTTable {
                continue
            }
            lMaxLiveCycle := L.GetField(LClawn, "max_live_cycle")
            if max_live_cycle := int(lua.LVAsNumber(lMaxLiveCycle)); max_live_cycle > 0 {
                c.SetMaxLiveCycle(max_live_cycle, clawn.MaxLiveCycleDeltaPercent)
            }
        }
    }

    if LSharkFactory := L.GetGlobal("shark_factory"); LSharkFactory.Type() == lua.LTTable {
        lTSharkFactory := LSharkFactory.(*lua.LTable)
        for i := 1; i < lTSharkFactory.Len(); i++ {
            s := b.BuildCell(shark.Shark{})
            LShark := lTSharkFactory.RawGetInt(i)
            if LShark.Type() != lua.LTTable {
                continue
            }
            lMaxLiveCycle := L.GetField(LShark, "max_live_cycle")
            if max_live_cycle := int(lua.LVAsNumber(lMaxLiveCycle)); max_live_cycle > 0 {
                s.SetMaxLiveCycle(max_live_cycle, shark.MaxLiveCycleDeltaPercent)
            }
        }
    }
}