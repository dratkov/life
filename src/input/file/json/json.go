package json

import (
	"flag"
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

func ( jsontype *JSON ) CheckFile() bool {
	file_name := flag.String( "from-json-file", "", "json file name" )
	flag.Parse()
	file, e := ioutil.ReadFile( *file_name )
	if e != nil {
		fmt.Println( "File error: %v\n", e )
		return false
	}

	json.Unmarshal( file, &jsontype )

	return true
}

func ( jsontype *JSON ) BuildAreaAndCells( b build.Builder ) {
    b.BuildArea( jsontype.Life.Area.Width, jsontype.Life.Area.Height )
    //for i := 0; i < jsontype.Life.Cells.Clawn.Count; i++ {
    b.BuildCell( clawn.Clawn{}, jsontype.Life.Cells.Clawn.Count )
    //b.BuildCellClawn( jsontype.Life.Cells.Clawn.Count )
	//}
    //for i := 0; i < jsontype.Life.Cells.Shark.Count; i++ {
    	//b.BuildCell( &shark.Shark{}, 1 )
    b.BuildCell( shark.Shark{}, jsontype.Life.Cells.Shark.Count )
    //b.BuildCellShark( jsontype.Life.Cells.Shark.Count )
	//}
}
