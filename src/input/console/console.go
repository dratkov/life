package console

import (
    "os"
    "fmt"
    "strconv"
    "../../build"
    //"../../cell/fish/clawn"
    //"../../cell/fish/predator/shark"
)

type Console struct {
    
}

func New() *Console {
    return &Console{}
}

func ( c *Console ) IsSourceInit() bool {
    return len( os.Args ) == 1
}

func ( c *Console ) GetDataAndBuildAreaAndCell( b build.Builder ) {
    c.GetAreaDataAndBuild( b )
    c.GetFishDataAndBuild( b )
}

func ( c *Console ) GetAreaDataAndBuild( b build.Builder ) {
    var area_width_str, area_height_str string
    var area_width_i, area_height_i int
    fmt.Print( "Enter width of area: " )
    fmt.Scan( &area_width_str )
    area_width_i, _ = strconv.Atoi( area_width_str )
    area_width_u := uint(area_width_i)
    fmt.Print( "Enter height of area: " )
    fmt.Scan( &area_height_str )
    area_height_i, _ = strconv.Atoi( area_height_str )
    area_height_u := uint(area_height_i)
    b.BuildArea( area_width_u, area_height_u )
}

func ( c *Console ) GetFishDataAndBuild( b build.Builder ) {
    //var fish_type string
    /*
    var clawn_count, shark_count int
    var clawn_count_str, shark_count_str string
    //fmt.Print( "Enter fish type: " )
    //fmt.Scan( &fish_type )
    fmt.Print( "Enter fish clawn count: " )
    fmt.Scan( &clawn_count_str )
    clawn_count, _ = strconv.Atoi( clawn_count_str )
    //b.BuildCell( &clawn.Clawn{}, clawn_count )
    fmt.Print( "Enter fish shark count: " )
    fmt.Scan( &shark_count_str )
    shark_count, _ = strconv.Atoi( shark_count_str )
    //b.BuildCell( &shark.Shark{}, shark_count )
    */
}

func (c *Console) BuildAreaAndCells(b build.Builder) {

}