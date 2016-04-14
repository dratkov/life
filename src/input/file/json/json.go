package json

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"../../../build"
    "../../../cell/fish/clawn"
    "../../../cell/fish/predator/shark"
	//"../../../area"
)

type JSON struct {
    Life LifeType
	file_name *string
}

type AreaType struct {
	Width, Height uint
}

type LifeType struct {
	Area AreaType
	Cells CellsType
}

type CellsType struct {
	Clawn ClawnType
	Shark ShartType
}

type CellCount struct {
	Count int
}

type ClawnType struct {
	CellCount
}

type ShartType struct {
	CellCount
}

func New(file_name *string) *JSON {
	j := &JSON{}
	j.file_name = file_name
	file, e := ioutil.ReadFile( *file_name )
	if e != nil {
		fmt.Println( "File error: %v\n", e )
		return nil
	}
	json.Unmarshal( file, &j )

	return j
}

func ( jsontype *JSON ) BuildAreaAndCells( b build.Builder ) {
    b.BuildArea( jsontype.Life.Area.Width, jsontype.Life.Area.Height )
    b.BuildCell( clawn.Clawn{}, jsontype.Life.Cells.Clawn.Count )
    b.BuildCell( shark.Shark{}, jsontype.Life.Cells.Shark.Count )
}
